// Package vanilladata records version-locked vanilla registry packets from BDS.
package vanilladata

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"

	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

const captureFormatVersion = 4

// PacketSpec names one clientbound packet payload captured verbatim from BDS.
type PacketSpec struct {
	ID       uint32
	Name     string
	File     string
	Optional bool
}

// DefaultPacketSpecs returns the vanilla registry packets emitted by BDS during login.
func DefaultPacketSpecs() []PacketSpec {
	return []PacketSpec{
		{ID: packet.IDCraftingData, Name: "CraftingDataPacket", File: "crafting_data.dat"},
		{ID: packet.IDAvailableActorIdentifiers, Name: "AvailableActorIdentifiersPacket", File: "available_actor_identifiers.dat"},
		{ID: packet.IDBiomeDefinitionList, Name: "BiomeDefinitionListPacket", File: "biome_definition_list.dat"},
		{ID: packet.IDCreativeContent, Name: "CreativeContentPacket", File: "creative_content.dat"},
		{ID: packet.IDItemRegistry, Name: "ItemRegistryPacket", File: "item_registry.dat"},
		{ID: packet.IDDimensionData, Name: "DimensionDataPacket", File: "dimension_data.dat", Optional: true},
		{ID: packet.IDFeatureRegistry, Name: "FeatureRegistryPacket", File: "feature_registry.dat", Optional: true},
		{ID: packet.IDCameraPresets, Name: "CameraPresetsPacket", File: "camera_presets.dat"},
		{ID: packet.IDTrimData, Name: "TrimDataPacket", File: "trim_data.dat"},
		{ID: packet.IDVoxelShapes, Name: "VoxelShapesPacket", File: "voxel_shapes.dat"},
	}
}

// ValidateContract ensures the capture binary, source lock, and packet list describe one exact target.
func ValidateContract(m Manifest, source SourceConfig, minecraftVersion string, protocolVersion int, specs []PacketSpec) error {
	if m.Target.MinecraftVersion != source.MinecraftVersion || m.Target.ProtocolVersion != source.ProtocolVersion {
		return fmt.Errorf("manifest target %s/%d does not match vanilla source %s/%d", m.Target.MinecraftVersion, m.Target.ProtocolVersion, source.MinecraftVersion, source.ProtocolVersion)
	}
	if m.Target.MinecraftVersion != minecraftVersion {
		return fmt.Errorf("manifest Minecraft version %s does not match gophertunnel %s", m.Target.MinecraftVersion, minecraftVersion)
	}
	if m.Target.ProtocolVersion != protocolVersion {
		return fmt.Errorf("manifest protocol %d does not match gophertunnel protocol %d", m.Target.ProtocolVersion, protocolVersion)
	}
	var revision string
	for _, pin := range m.Sources {
		if pin.Kind == "gophertunnel-exact-codec" {
			revision = pin.Revision
			break
		}
	}
	if revision == "" || revision != source.Gophertunnel.Revision {
		return fmt.Errorf("manifest gophertunnel revision %q does not match vanilla source %q", revision, source.Gophertunnel.Revision)
	}
	packets := make(map[uint32]Packet, len(m.Packets))
	for _, manifestPacket := range m.Packets {
		packets[manifestPacket.ID] = manifestPacket
	}
	for _, spec := range specs {
		manifestPacket, ok := packets[spec.ID]
		if !ok {
			return fmt.Errorf("capture packet %s has ID %d, which is absent from the manifest", spec.Name, spec.ID)
		}
		if manifestPacket.Name != spec.Name {
			return fmt.Errorf("capture packet ID %d is %s, manifest has %s", spec.ID, spec.Name, manifestPacket.Name)
		}
		if manifestPacket.Direction != "clientbound" {
			return fmt.Errorf("capture packet %s must be clientbound, manifest direction is %s", spec.Name, manifestPacket.Direction)
		}
	}
	return nil
}

// Recorder retains the first payload observed for every configured packet ID.
type Recorder struct {
	mu       sync.Mutex
	specs    []PacketSpec
	byID     map[uint32]PacketSpec
	payloads map[string][]byte
}

// NewRecorder creates an empty packet recorder.
func NewRecorder(specs []PacketSpec) *Recorder {
	copySpecs := append([]PacketSpec(nil), specs...)
	byID := make(map[uint32]PacketSpec, len(copySpecs))
	for _, spec := range copySpecs {
		byID[spec.ID] = spec
	}
	return &Recorder{specs: copySpecs, byID: byID, payloads: make(map[string][]byte, len(copySpecs))}
}

// Observe records payload when packetID is part of the capture contract.
func (r *Recorder) Observe(packetID uint32, payload []byte) {
	r.mu.Lock()
	defer r.mu.Unlock()
	spec, ok := r.byID[packetID]
	if !ok {
		return
	}
	if _, exists := r.payloads[spec.Name]; exists {
		return
	}
	r.payloads[spec.Name] = append([]byte(nil), payload...)
}

// Complete reports whether every required packet has been captured.
func (r *Recorder) Complete() bool {
	return len(r.Missing()) == 0
}

// Missing returns uncaptured required packet names in capture-contract order.
func (r *Recorder) Missing() []string {
	r.mu.Lock()
	defer r.mu.Unlock()
	var missing []string
	for _, spec := range r.specs {
		if spec.Optional {
			continue
		}
		if _, ok := r.payloads[spec.Name]; !ok {
			missing = append(missing, spec.Name)
		}
	}
	return missing
}

// Payloads returns a copy of all recorded packet payloads keyed by manifest name.
func (r *Recorder) Payloads() map[string][]byte {
	r.mu.Lock()
	defer r.mu.Unlock()
	result := make(map[string][]byte, len(r.payloads))
	for name, payload := range r.payloads {
		result[name] = append([]byte(nil), payload...)
	}
	return result
}

// BDSProvenance fingerprints the exact server input used for the capture.
type BDSProvenance struct {
	Version          string            `json:"version"`
	ArchiveVersion   string            `json:"archive_version,omitempty"`
	ArchiveURL       string            `json:"archive_url"`
	ArchiveSHA256    string            `json:"archive_sha256"`
	BinarySHA256     string            `json:"binary_sha256"`
	ServerProperties map[string]string `json:"server_properties"`
}

// GophertunnelProvenance identifies the exact bot codec.
type GophertunnelProvenance struct {
	Repository    string `json:"repository"`
	ModulePath    string `json:"module_path"`
	Revision      string `json:"revision"`
	ModuleVersion string `json:"module_version"`
}

// ArtifactInput contains everything required to write one versioned vanilla-data directory.
type ArtifactInput struct {
	Target        Target
	BDS           BDSProvenance
	Gophertunnel  GophertunnelProvenance
	Specs         []PacketSpec
	Payloads      map[string][]byte
	InternalData  *InternalDataManifest
	InternalFiles map[string][]byte
	DerivedFiles  map[string][]byte
	Warning       string
}

// CaptureFile records the digest of one generated file.
type CaptureFile struct {
	File       string `json:"file"`
	Kind       string `json:"kind"`
	PacketID   uint32 `json:"packet_id,omitempty"`
	PacketName string `json:"packet_name,omitempty"`
	Bytes      int    `json:"bytes"`
	SHA256     string `json:"sha256"`
}

// CapturePacket records whether each contracted packet was present or absent.
type CapturePacket struct {
	ID       uint32 `json:"id"`
	Name     string `json:"name"`
	Optional bool   `json:"optional"`
	Status   string `json:"status"`
	File     string `json:"file,omitempty"`
}

// CaptureMetadata is the reproducible index written as capture.json.
type CaptureMetadata struct {
	FormatVersion int                    `json:"format_version"`
	Target        Target                 `json:"target"`
	BDS           BDSProvenance          `json:"bds"`
	Gophertunnel  GophertunnelProvenance `json:"gophertunnel"`
	InternalData  *InternalDataManifest  `json:"internal_data,omitempty"`
	Packets       []CapturePacket        `json:"packets"`
	Files         []CaptureFile          `json:"files"`
	Warning       string                 `json:"warning,omitempty"`
}

// WriteArtifacts atomically replaces out with exact payloads and their provenance index.
func WriteArtifacts(out string, input ArtifactInput) error {
	if input.BDS.Version == "" || input.BDS.ArchiveURL == "" || !validHex(input.BDS.ArchiveSHA256, 64) || !validHex(input.BDS.BinarySHA256, 64) {
		return fmt.Errorf("complete BDS provenance is required")
	}
	archiveVersion := input.BDS.ArchiveVersion
	if archiveVersion == "" {
		archiveVersion = input.BDS.Version
	}
	if !strings.HasSuffix(input.BDS.ArchiveURL, "bedrock-server-"+archiveVersion+".zip") {
		return fmt.Errorf("BDS archive version %s does not match archive URL", archiveVersion)
	}
	if input.Gophertunnel.Repository == "" || input.Gophertunnel.ModulePath == "" || input.Gophertunnel.ModuleVersion == "" || !validHex(input.Gophertunnel.Revision, 40) {
		return fmt.Errorf("complete gophertunnel provenance is required")
	}
	cleanOut := filepath.Clean(out)
	if cleanOut == "." || cleanOut == string(filepath.Separator) || filepath.Base(cleanOut) == ".." {
		return fmt.Errorf("unsafe vanilla-data output directory %q", out)
	}
	parent := filepath.Dir(cleanOut)
	if err := os.MkdirAll(parent, 0o755); err != nil {
		return fmt.Errorf("create vanilla-data parent: %w", err)
	}
	temporary, err := os.MkdirTemp(parent, ".vanilla-data-stage-*")
	if err != nil {
		return fmt.Errorf("create vanilla-data staging directory: %w", err)
	}
	defer os.RemoveAll(temporary)
	if err := os.Chmod(temporary, 0o755); err != nil {
		return fmt.Errorf("chmod vanilla-data staging directory: %w", err)
	}

	if (input.InternalData == nil) != (len(input.InternalFiles) == 0) {
		return fmt.Errorf("internal data provenance and files must be supplied together")
	}
	if input.InternalData != nil {
		if err := validateInternalManifest(*input.InternalData, input.Target, input.BDS.Version, input.InternalData.Endstone); err != nil {
			return err
		}
		if err := validateInternalFiles(*input.InternalData, input.InternalFiles); err != nil {
			return err
		}
	}
	metadata := CaptureMetadata{FormatVersion: captureFormatVersion, Target: input.Target, BDS: input.BDS, Gophertunnel: input.Gophertunnel, InternalData: input.InternalData, Warning: input.Warning}
	writtenFiles := make(map[string]struct{}, len(input.Specs)+len(input.InternalFiles)+len(input.DerivedFiles)+1)
	for _, spec := range input.Specs {
		if err := validateArtifactName(spec.File); err != nil {
			return err
		}
		payload, present := input.Payloads[spec.Name]
		status := CapturePacket{ID: spec.ID, Name: spec.Name, Optional: spec.Optional, Status: "absent"}
		if !present {
			if !spec.Optional {
				return fmt.Errorf("required packet %s is absent", spec.Name)
			}
			metadata.Packets = append(metadata.Packets, status)
			continue
		}
		if len(payload) == 0 {
			return fmt.Errorf("captured packet %s has an empty payload", spec.Name)
		}
		if _, exists := writtenFiles[spec.File]; exists {
			return fmt.Errorf("duplicate capture file %s", spec.File)
		}
		if err := os.WriteFile(filepath.Join(temporary, spec.File), payload, 0o644); err != nil {
			return fmt.Errorf("write %s: %w", spec.File, err)
		}
		status.Status = "captured"
		status.File = spec.File
		metadata.Packets = append(metadata.Packets, status)
		metadata.Files = append(metadata.Files, fileMetadata(spec.File, payload, spec))
		writtenFiles[spec.File] = struct{}{}
	}
	internalNames := make([]string, 0, len(input.InternalFiles))
	for name := range input.InternalFiles {
		internalNames = append(internalNames, name)
	}
	sort.Strings(internalNames)
	for _, name := range internalNames {
		if err := validateDerivedName(name); err != nil {
			return err
		}
		if _, exists := writtenFiles[name]; exists {
			return fmt.Errorf("duplicate capture file %s", name)
		}
		data := input.InternalFiles[name]
		if len(data) == 0 {
			return fmt.Errorf("internal data file %s is empty", name)
		}
		path := filepath.Join(temporary, filepath.FromSlash(name))
		if err := os.WriteFile(path, data, 0o644); err != nil {
			return fmt.Errorf("write %s: %w", name, err)
		}
		metadata.Files = append(metadata.Files, internalFileMetadata(name, data))
		writtenFiles[name] = struct{}{}
	}
	derivedNames := make([]string, 0, len(input.DerivedFiles))
	for name := range input.DerivedFiles {
		derivedNames = append(derivedNames, name)
	}
	sort.Strings(derivedNames)
	for _, name := range derivedNames {
		if err := validateDerivedName(name); err != nil {
			return err
		}
		data := input.DerivedFiles[name]
		if _, exists := writtenFiles[name]; exists {
			return fmt.Errorf("duplicate capture file %s", name)
		}
		if len(data) == 0 {
			return fmt.Errorf("compatibility file %s is empty", name)
		}
		path := filepath.Join(temporary, filepath.FromSlash(name))
		if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
			return fmt.Errorf("create compatibility directory for %s: %w", name, err)
		}
		if err := os.WriteFile(path, data, 0o644); err != nil {
			return fmt.Errorf("write %s: %w", name, err)
		}
		metadata.Files = append(metadata.Files, derivedFileMetadata(name, data))
		writtenFiles[name] = struct{}{}
	}
	metadataData, err := json.MarshalIndent(metadata, "", "  ")
	if err != nil {
		return fmt.Errorf("encode capture metadata: %w", err)
	}
	if err := os.WriteFile(filepath.Join(temporary, "capture.json"), append(metadataData, '\n'), 0o644); err != nil {
		return fmt.Errorf("write capture.json: %w", err)
	}
	return replaceDirectory(cleanOut, temporary)
}

func replaceDirectory(out, staged string) error {
	if info, err := os.Lstat(out); err == nil {
		if !info.IsDir() || info.Mode()&os.ModeSymlink != 0 {
			return fmt.Errorf("vanilla-data output %q is not a directory", out)
		}
	} else if !os.IsNotExist(err) {
		return fmt.Errorf("inspect vanilla-data output: %w", err)
	}
	backup := out + ".previous"
	if _, err := os.Lstat(backup); err == nil {
		if _, outErr := os.Lstat(out); os.IsNotExist(outErr) {
			if err := os.Rename(backup, out); err != nil {
				return fmt.Errorf("restore interrupted vanilla-data backup: %w", err)
			}
		} else {
			return fmt.Errorf("refusing to overwrite existing backup %q", backup)
		}
	} else if !os.IsNotExist(err) {
		return fmt.Errorf("inspect vanilla-data backup: %w", err)
	}
	hadPrevious := false
	if _, err := os.Lstat(out); err == nil {
		if err := os.Rename(out, backup); err != nil {
			return fmt.Errorf("backup previous vanilla data: %w", err)
		}
		hadPrevious = true
	}
	if err := os.Rename(staged, out); err != nil {
		if hadPrevious {
			_ = os.Rename(backup, out)
		}
		return fmt.Errorf("install vanilla data: %w", err)
	}
	if hadPrevious {
		if err := os.RemoveAll(backup); err != nil {
			return fmt.Errorf("remove previous vanilla data: %w", err)
		}
	}
	return nil
}

func fileMetadata(name string, data []byte, spec PacketSpec) CaptureFile {
	digest := sha256.Sum256(data)
	return CaptureFile{File: name, Kind: "packet_payload", PacketID: spec.ID, PacketName: spec.Name, Bytes: len(data), SHA256: "sha256:" + hex.EncodeToString(digest[:])}
}

func derivedFileMetadata(name string, data []byte) CaptureFile {
	digest := sha256.Sum256(data)
	return CaptureFile{File: name, Kind: "derived", Bytes: len(data), SHA256: "sha256:" + hex.EncodeToString(digest[:])}
}

func internalFileMetadata(name string, data []byte) CaptureFile {
	digest := sha256.Sum256(data)
	return CaptureFile{File: name, Kind: "internal_data", Bytes: len(data), SHA256: "sha256:" + hex.EncodeToString(digest[:])}
}

func validateDerivedName(name string) error {
	clean := filepath.ToSlash(filepath.Clean(name))
	if name == "" || clean != name || filepath.Base(clean) != clean || name == "capture.json" || strings.Contains(name, "\\") {
		return fmt.Errorf("invalid derived file name %q", name)
	}
	return nil
}

func validateArtifactName(name string) error {
	if name == "" || name == "." || name == ".." || filepath.Base(name) != name || strings.Contains(name, "\\") {
		return fmt.Errorf("invalid capture file name %q", name)
	}
	return nil
}
