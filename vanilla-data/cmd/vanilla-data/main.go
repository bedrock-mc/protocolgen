// Command vanilla-data connects an offline gophertunnel client to BDS and
// records version-locked vanilla registry packet payloads.
package main

import (
	"bufio"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"path/filepath"
	"runtime/debug"
	"sort"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/sandertv/gophertunnel/minecraft"
	gtprotocol "github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
	"protocolgen/vanilla-data/internal/vanilladata"
)

func main() {
	if err := run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}

func run(args []string) error {
	fs := flag.NewFlagSet("vanilla-data", flag.ContinueOnError)
	fs.SetOutput(io.Discard)
	manifestPath := fs.String("manifest", "", "target protocol manifest")
	sourcePath := fs.String("source", "", "version-locked vanilla capture source")
	outPath := fs.String("out", "", "versioned vanilla-data output directory")
	internalDataPath := fs.String("internal-data", "", "verified headless Endstone export directory")
	bdsBinary := fs.String("bds-binary", "", "exact bedrock_server executable used for the capture")
	address := fs.String("address", "127.0.0.1:19132", "local offline BDS RakNet address")
	timeout := fs.Duration("timeout", 90*time.Second, "maximum dial, spawn, and required-packet capture time")
	settle := fs.Duration("settle", 2*time.Second, "additional window for optional login packets")
	validateOnly := fs.Bool("validate-only", false, "validate the manifest, source lock, packet contract, and compiled codec without connecting")
	if err := fs.Parse(args); err != nil {
		return err
	}
	var missingFlags []string
	required := map[string]string{"-manifest": *manifestPath, "-source": *sourcePath}
	if !*validateOnly {
		required["-out"] = *outPath
		required["-bds-binary"] = *bdsBinary
	}
	for name, value := range required {
		if value == "" {
			missingFlags = append(missingFlags, name)
		}
	}
	if len(missingFlags) != 0 {
		sort.Strings(missingFlags)
		return fmt.Errorf("required flags: %s", strings.Join(missingFlags, ", "))
	}
	if *timeout <= 0 || *settle < 0 {
		return fmt.Errorf("-timeout must be positive and -settle must not be negative")
	}
	if err := requireLoopback(*address); err != nil {
		return err
	}

	targetManifest, err := vanilladata.LoadManifest(*manifestPath)
	if err != nil {
		return err
	}
	if err := vanilladata.ValidateGeneratedTarget(targetManifest.Target); err != nil {
		return err
	}
	source, err := vanilladata.LoadSourceConfig(*sourcePath)
	if err != nil {
		return err
	}
	specs := vanilladata.DefaultPacketSpecs()
	if err := vanilladata.ValidateContract(targetManifest, source, gtprotocol.CurrentVersion, gtprotocol.CurrentProtocol, specs); err != nil {
		return err
	}
	if err := validateAddressPort(*address, source.ServerProperties["server-port"]); err != nil {
		return err
	}
	if err := verifyGophertunnelBuild(source.Gophertunnel); err != nil {
		return err
	}
	var internalFiles map[string][]byte
	var internalManifest *vanilladata.InternalDataManifest
	if *internalDataPath != "" {
		if source.Endstone == nil {
			return fmt.Errorf("internal data requires a version-locked Endstone source lock")
		}
		files, manifest, err := vanilladata.LoadInternalArtifacts(*internalDataPath, targetManifest.Target, source.BDS.Version, *source.Endstone)
		if err != nil {
			return err
		}
		internalFiles = files
		internalManifest = &manifest
	}
	if *validateOnly {
		fmt.Printf("valid capture contract: Minecraft %s / protocol %d / gophertunnel %s\n", targetManifest.Target.MinecraftVersion, targetManifest.Target.ProtocolVersion, source.Gophertunnel.Revision)
		return nil
	}
	binarySHA256, err := hashFile(*bdsBinary)
	if err != nil {
		return fmt.Errorf("fingerprint BDS binary: %w", err)
	}
	if err := verifyServerProperties(filepath.Join(filepath.Dir(*bdsBinary), "server.properties"), source.ServerProperties); err != nil {
		return err
	}

	recorder := vanilladata.NewRecorder(specs)
	serverAddress, err := net.ResolveUDPAddr("udp", *address)
	if err != nil {
		return fmt.Errorf("resolve BDS address: %w", err)
	}
	dialer := minecraft.Dialer{
		DisconnectOnInvalidPackets: true,
		DisconnectOnUnknownPackets: true,
		PacketFunc: func(header packet.Header, payload []byte, sourceAddress, _ net.Addr) {
			if packetFromServer(sourceAddress, serverAddress) {
				recorder.Observe(header.PacketID, payload)
			}
		},
		DownloadResourcePack: func(_ uuid.UUID, _ string, _, _ int) bool { return false },
	}
	ctx, cancel := context.WithTimeout(context.Background(), *timeout)
	defer cancel()
	conn, err := dialer.DialContext(ctx, "raknet", *address)
	if err != nil {
		return fmt.Errorf("dial offline BDS: %w", err)
	}
	defer conn.Close()
	if err := conn.DoSpawnContext(ctx); err != nil {
		return fmt.Errorf("spawn BDS capture bot: %w", err)
	}
	if deadline, ok := ctx.Deadline(); ok {
		_ = conn.SetReadDeadline(deadline)
	}
	for !recorder.Complete() {
		if _, err := conn.ReadPacket(); err != nil {
			if errors.Is(ctx.Err(), context.DeadlineExceeded) || isTimeout(err) {
				return fmt.Errorf("capture timed out; missing packets: %s", strings.Join(recorder.Missing(), ", "))
			}
			return fmt.Errorf("read BDS packets; missing %s: %w", strings.Join(recorder.Missing(), ", "), err)
		}
	}
	var settleErr error
	if *settle > 0 {
		_ = conn.SetReadDeadline(time.Now().Add(*settle))
		for {
			if _, err := conn.ReadPacket(); err != nil {
				if isTimeout(err) {
					break
				}
				settleErr = err
				break
			}
		}
	}

	warning := ""
	if settleErr != nil {
		warning = "optional packet settling ended with: " + settleErr.Error()
	}
	payloads := recorder.Payloads()
	derivedFiles, err := vanilladata.BuildDerivedArtifacts(payloads)
	if err != nil {
		return fmt.Errorf("build derived vanilla data with generated protocol: %w", err)
	}
	if err := vanilladata.WriteArtifacts(*outPath, vanilladata.ArtifactInput{
		Target: targetManifest.Target,
		BDS: vanilladata.BDSProvenance{
			Version:          source.BDS.Version,
			ArchiveURL:       source.BDS.Linux.URL,
			ArchiveSHA256:    source.BDS.Linux.ArchiveSHA256,
			BinarySHA256:     binarySHA256,
			ServerProperties: source.ServerProperties,
		},
		Gophertunnel: vanilladata.GophertunnelProvenance{
			Repository:    source.Gophertunnel.Repository,
			ModulePath:    source.Gophertunnel.ModulePath,
			Revision:      source.Gophertunnel.Revision,
			ModuleVersion: source.Gophertunnel.ModuleVersion,
		},
		Specs:         specs,
		Payloads:      payloads,
		InternalData:  internalManifest,
		InternalFiles: internalFiles,
		DerivedFiles:  derivedFiles,
		Warning:       warning,
	}); err != nil {
		return err
	}
	if settleErr != nil {
		return fmt.Errorf("captured required packet evidence, but optional packet settling failed: %w", settleErr)
	}
	fmt.Printf("captured vanilla registry packets from BDS %s -> %s\n", source.BDS.Version, *outPath)
	return nil
}

func requireLoopback(address string) error {
	host, _, err := net.SplitHostPort(address)
	if err != nil {
		return fmt.Errorf("invalid -address: %w", err)
	}
	if strings.EqualFold(host, "localhost") {
		return nil
	}
	ip := net.ParseIP(host)
	if ip == nil || !ip.IsLoopback() {
		return fmt.Errorf("-address must be loopback so the fingerprinted local BDS is the capture source")
	}
	return nil
}

func validateAddressPort(address, expectedPort string) error {
	_, port, err := net.SplitHostPort(address)
	if err != nil {
		return fmt.Errorf("invalid -address: %w", err)
	}
	if port != expectedPort {
		return fmt.Errorf("BDS address port %s does not match source-locked server-port %s", port, expectedPort)
	}
	return nil
}

func packetFromServer(source net.Addr, expected *net.UDPAddr) bool {
	if source == nil {
		return false
	}
	actual, err := net.ResolveUDPAddr("udp", source.String())
	return err == nil && actual.Port == expected.Port && actual.IP.Equal(expected.IP)
}

func isTimeout(err error) bool {
	var netErr net.Error
	return errors.As(err, &netErr) && netErr.Timeout()
}

func hashFile(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()
	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}
	return hex.EncodeToString(hash.Sum(nil)), nil
}

func verifyServerProperties(path string, expected map[string]string) error {
	file, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("open BDS server properties: %w", err)
	}
	defer file.Close()
	actual := make(map[string]string)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		key, value, ok := strings.Cut(line, "=")
		if ok {
			actual[strings.TrimSpace(key)] = strings.TrimSpace(value)
		}
	}
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("read BDS server properties: %w", err)
	}
	for key, value := range expected {
		if actual[key] != value {
			return fmt.Errorf("BDS server property %s=%q, want %q", key, actual[key], value)
		}
	}
	return nil
}

func verifyGophertunnelBuild(source vanilladata.GophertunnelSource) error {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return fmt.Errorf("read Go build information")
	}
	for _, dependency := range info.Deps {
		if dependency.Path != "github.com/sandertv/gophertunnel" {
			continue
		}
		selected := dependency
		if dependency.Replace != nil {
			selected = dependency.Replace
		}
		if selected.Path != source.ModulePath || selected.Version != source.ModuleVersion || !moduleVersionMatchesRevision(selected.Version, source.Revision) {
			return fmt.Errorf("compiled gophertunnel %s@%s does not match source lock %s@%s", selected.Path, selected.Version, source.ModulePath, source.ModuleVersion)
		}
		return nil
	}
	return fmt.Errorf("compiled binary has no gophertunnel module dependency")
}

func moduleVersionMatchesRevision(version, revision string) bool {
	return len(revision) >= 12 && strings.HasSuffix(version, "-"+revision[:12])
}
