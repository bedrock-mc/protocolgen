// Command protocolgen is the v2 protocol pipeline. Sources are lowered into
// independent claims, reconciled into one version-pinned manifest, and only
// then handed to emitters or parity adapters.
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"protocolgen/internal/claims"
	"protocolgen/internal/emitgo"
	"protocolgen/internal/emitrust"
	"protocolgen/internal/ingest"
	"protocolgen/internal/manifest"
	"protocolgen/internal/parity"
	"protocolgen/internal/reconcile"
	"protocolgen/internal/sourcelock"
)

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(2)
	}
	var err error
	switch os.Args[1] {
	case "reconcile":
		err = runReconcile(os.Args[2:])
	case "ingest":
		err = runIngest(os.Args[2:])
	case "validate":
		err = runValidate(os.Args[2:])
	case "emit-go":
		err = runEmitGo(os.Args[2:])
	case "emit-rust":
		err = runEmitRust(os.Args[2:])
	case "parity":
		err = runParity(os.Args[2:])
	case "hash-source":
		err = runHashSource(os.Args[2:])
	default:
		usage()
		os.Exit(2)
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}

func usage() {
	fmt.Fprintln(os.Stderr, `usage: protocolgen <reconcile|ingest|validate|emit-go|emit-rust|parity|hash-source> [flags]

reconcile lowers one or both explicit source checkouts and writes manifest v2.
ingest lowers one source to auditable machine-derived claims JSON.
validate checks a canonical manifest and all fingerprint metadata.
emit-go and emit-rust consume only a canonical manifest.
parity compares a canonical manifest with Axolotl's public v1 wire manifest.
hash-source prints the deterministic source-tree digest for a lock file.`)
}

func runReconcile(args []string) error {
	fs := flag.NewFlagSet("reconcile", flag.ContinueOnError)
	lockPath := fs.String("lock", "", "source lock JSON")
	mojangRoot := fs.String("mojang", "", "local raw Mojang checkout or json directory")
	mojangID := fs.String("mojang-id", "mojang", "source-lock id for -mojang")
	mojangCorrections := fs.String("mojang-corrections", "", "fingerprinted correction directory for -mojang")
	endstoneRoot := fs.String("endstone", "", "local Endstone dump checkout")
	endstoneID := fs.String("endstone-id", "endstone", "source-lock id for -endstone")
	endstoneCorrections := fs.String("endstone-corrections", "", "fingerprinted correction directory for -endstone")
	adjudicationsPath := fs.String("adjudications", "", "fingerprinted adjudications JSON")
	outPath := fs.String("out", "", "canonical manifest v2 output")
	if err := fs.Parse(args); err != nil {
		return err
	}
	if *lockPath == "" || *outPath == "" {
		return fmt.Errorf("-lock and -out are required")
	}
	lock, err := sourcelock.Load(*lockPath)
	if err != nil {
		return err
	}
	var results []claims.Result
	if *mojangRoot != "" {
		result, err := loadSource(lock, "mojang", *mojangID, *mojangRoot, *mojangCorrections)
		if err != nil {
			return err
		}
		results = append(results, result)
	}
	if *endstoneRoot != "" {
		result, err := loadSource(lock, "endstone", *endstoneID, *endstoneRoot, *endstoneCorrections)
		if err != nil {
			return err
		}
		results = append(results, result)
	}
	if len(results) == 0 {
		return fmt.Errorf("at least one of -mojang or -endstone is required")
	}
	var adjudications []manifest.Adjudication
	if *adjudicationsPath != "" {
		adjudications, err = manifest.LoadAdjudications(*adjudicationsPath)
		if err != nil {
			return err
		}
	}
	result, err := reconcile.Reconcile(lock.Target, results, adjudications)
	if err != nil {
		return err
	}
	if err := manifest.Write(*outPath, result); err != nil {
		return err
	}
	fmt.Printf("manifest v2: %d packets, %d source pins, %d adjudications -> %s\n", len(result.Packets), len(result.Sources), len(result.Adjudications), *outPath)
	return nil
}

func runIngest(args []string) error {
	fs := flag.NewFlagSet("ingest", flag.ContinueOnError)
	kind := fs.String("kind", "", "mojang or endstone")
	root := fs.String("root", "", "source checkout")
	lockPath := fs.String("lock", "", "source lock JSON")
	id := fs.String("id", "", "source-lock id")
	corrections := fs.String("corrections", "", "fingerprinted correction directory")
	outPath := fs.String("out", "", "claims JSON output")
	if err := fs.Parse(args); err != nil {
		return err
	}
	if *kind == "" || *root == "" || *lockPath == "" || *id == "" || *outPath == "" {
		return fmt.Errorf("-kind, -root, -lock, -id, and -out are required")
	}
	lock, err := sourcelock.Load(*lockPath)
	if err != nil {
		return err
	}
	result, err := loadSource(lock, *kind, *id, *root, *corrections)
	if err != nil {
		return err
	}
	data, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return err
	}
	data = append(data, '\n')
	if err := os.MkdirAll(filepath.Dir(*outPath), 0o755); err != nil && filepath.Dir(*outPath) != "." {
		return err
	}
	if err := os.WriteFile(*outPath, data, 0o644); err != nil {
		return err
	}
	fmt.Printf("%s claims: %d -> %s\n", *kind, len(result.Claims), *outPath)
	return nil
}

func loadSource(lock sourcelock.Lock, kind, id, root, corrections string) (claims.Result, error) {
	pin, ok := lock.Source(id)
	if !ok {
		return claims.Result{}, fmt.Errorf("source-lock has no %q pin", id)
	}
	if pin.Kind != kind {
		return claims.Result{}, fmt.Errorf("source %q is pinned as kind %q, not %q", id, pin.Kind, kind)
	}
	if !sourcelock.IsSyntheticDigest(pin.Digest) {
		if err := sourcelock.VerifyDirectory(root, pin); err != nil {
			return claims.Result{}, err
		}
	}
	switch kind {
	case "mojang":
		return ingest.ParseMojang(root, pin, corrections)
	case "endstone":
		return ingest.ParseEndstone(root, pin, corrections)
	default:
		return claims.Result{}, fmt.Errorf("unknown source kind %q", kind)
	}
}

func runValidate(args []string) error {
	fs := flag.NewFlagSet("validate", flag.ContinueOnError)
	path := fs.String("manifest", "", "canonical manifest v2 JSON")
	if err := fs.Parse(args); err != nil {
		return err
	}
	if *path == "" {
		return fmt.Errorf("-manifest is required")
	}
	m, err := manifest.Load(*path)
	if err != nil {
		return err
	}
	fmt.Printf("valid manifest v%d: %d packets, protocol %d\n", m.SchemaVersion, len(m.Packets), m.Target.ProtocolVersion)
	return nil
}

func runEmitGo(args []string) error {
	fs := flag.NewFlagSet("emit-go", flag.ContinueOnError)
	manifestPath := fs.String("manifest", "", "canonical manifest v2 JSON")
	out := fs.String("out", "", "generated package directory")
	pkg := fs.String("pkg", "wiregen", "generated Go package name")
	if err := fs.Parse(args); err != nil {
		return err
	}
	if *manifestPath == "" || *out == "" {
		return fmt.Errorf("-manifest and -out are required")
	}
	m, err := manifest.Load(*manifestPath)
	if err != nil {
		return err
	}
	files, err := emitgo.Generate(m, *pkg)
	if err != nil {
		return err
	}
	if err := writeFiles(*out, files); err != nil {
		return err
	}
	fmt.Printf("Go emitter: %d files -> %s\n", len(files), *out)
	return nil
}

func runEmitRust(args []string) error {
	fs := flag.NewFlagSet("emit-rust", flag.ContinueOnError)
	manifestPath := fs.String("manifest", "", "canonical manifest v2 JSON")
	out := fs.String("out", "", "generated Rust source directory")
	if err := fs.Parse(args); err != nil {
		return err
	}
	if *manifestPath == "" || *out == "" {
		return fmt.Errorf("-manifest and -out are required")
	}
	m, err := manifest.Load(*manifestPath)
	if err != nil {
		return err
	}
	files, err := emitrust.GenerateFiles(m)
	if err != nil {
		return err
	}
	if err := writeFiles(*out, files); err != nil {
		return err
	}
	fmt.Printf("Rust emitter: %d files -> %s\n", len(files), *out)
	return nil
}

func runParity(args []string) error {
	fs := flag.NewFlagSet("parity", flag.ContinueOnError)
	manifestPath := fs.String("manifest", "", "canonical manifest v2 JSON")
	axolotlPath := fs.String("axolotl", "", "Axolotl public v1 wire manifest JSON")
	if err := fs.Parse(args); err != nil {
		return err
	}
	if *manifestPath == "" || *axolotlPath == "" {
		return fmt.Errorf("-manifest and -axolotl are required")
	}
	m, err := manifest.Load(*manifestPath)
	if err != nil {
		return err
	}
	if err := parity.CompareFile(m, *axolotlPath); err != nil {
		return err
	}
	fmt.Println("Axolotl v1 parity: byte-equivalence normalization matched")
	return nil
}

func runHashSource(args []string) error {
	fs := flag.NewFlagSet("hash-source", flag.ContinueOnError)
	root := fs.String("root", "", "source checkout")
	if err := fs.Parse(args); err != nil {
		return err
	}
	if *root == "" {
		return fmt.Errorf("-root is required")
	}
	digest, err := sourcelock.DigestDirectory(*root)
	if err != nil {
		return err
	}
	fmt.Println(digest)
	return nil
}

func writeFiles(directory string, files map[string]string) error {
	if err := os.MkdirAll(directory, 0o755); err != nil {
		return err
	}
	desired := make(map[string]bool, len(files))
	names := make([]string, 0, len(files))
	for name := range files {
		clean := filepath.Clean(name)
		if clean == "." || filepath.IsAbs(clean) || clean == ".." || strings.HasPrefix(clean, ".."+string(filepath.Separator)) {
			return fmt.Errorf("emitter returned unsafe filename %q", name)
		}
		desired[clean] = true
		names = append(names, name)
	}
	if err := removeStaleGeneratedFiles(directory, desired); err != nil {
		return err
	}
	sort.Strings(names)
	for _, name := range names {
		path := filepath.Join(directory, filepath.Clean(name))
		if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
			return err
		}
		if err := os.WriteFile(path, []byte(files[name]), 0o644); err != nil {
			return err
		}
	}
	return nil
}

func removeStaleGeneratedFiles(directory string, desired map[string]bool) error {
	const generatedHeader = "Code generated from canonical protocol manifest v2. DO NOT EDIT."
	return filepath.WalkDir(directory, func(path string, entry os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if entry.IsDir() {
			return nil
		}
		relative, err := filepath.Rel(directory, path)
		if err != nil || desired[relative] {
			return err
		}
		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		contents := strings.TrimLeft(string(data), "/# ")
		if strings.HasPrefix(contents, generatedHeader) {
			return os.Remove(path)
		}
		return nil
	})
}
