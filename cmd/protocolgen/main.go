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

	"protocolgen/internal/changelog"
	"protocolgen/internal/claims"
	"protocolgen/internal/direction"
	"protocolgen/internal/emitgo"
	"protocolgen/internal/emitrust"
	"protocolgen/internal/emitter"
	"protocolgen/internal/gophertunneloracle"
	"protocolgen/internal/ingest"
	"protocolgen/internal/manifest"
	"protocolgen/internal/nbtencoding"
	"protocolgen/internal/parity"
	"protocolgen/internal/reconcile"
	"protocolgen/internal/sourcelock"
	"protocolgen/internal/updateguide"
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
	case "verify-gophertunnel":
		err = runVerifyGophertunnel(os.Args[2:])
	case "update-guide":
		err = runUpdateGuide(os.Args[2:])
	case "changelog":
		err = runChangelog(os.Args[2:])
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
	fmt.Fprintln(os.Stderr, `usage: protocolgen <reconcile|ingest|validate|emit-go|emit-rust|parity|verify-gophertunnel|changelog|update-guide|hash-source> [flags]

reconcile lowers one or both explicit source checkouts and writes manifest v2.
ingest lowers one source to auditable machine-derived claims JSON.
validate checks a canonical manifest and all fingerprint metadata.
emit-go and emit-rust consume only a canonical manifest.
parity compares a canonical manifest with Axolotl's public v1 wire manifest.
verify-gophertunnel compares a canonical manifest with a pinned gophertunnel checkout and writes a JSON report.
changelog diffs two corrected Mojang schema directories into human-readable Markdown.
update-guide turns a protocol changelog and its target corrected schemas into gophertunnel transcription snippets.
hash-source prints the deterministic source-tree digest for a lock file.`)
}

func runChangelog(args []string) error {
	fs := flag.NewFlagSet("changelog", flag.ContinueOnError)
	fromPath := fs.String("from", "", "previous corrected Mojang JSON schema directory")
	toPath := fs.String("to", "", "target corrected Mojang JSON schema directory")
	fromBranch := fs.String("from-branch", "", "previous upstream branch")
	toBranch := fs.String("to-branch", "", "target upstream branch")
	fromUpstream := fs.String("from-upstream", "", "previous upstream commit")
	toUpstream := fs.String("to-upstream", "", "target upstream commit")
	fromFixer := fs.String("from-fixer", "", "previous fixer commit")
	toFixer := fs.String("to-fixer", "", "target fixer commit")
	outPath := fs.String("out", "", "human-readable protocol changelog Markdown output")
	if err := fs.Parse(args); err != nil {
		return err
	}
	if *fromPath == "" || *toPath == "" || *outPath == "" {
		return fmt.Errorf("-from, -to, and -out are required")
	}
	if *fromBranch == "" || *toBranch == "" || *fromUpstream == "" || *toUpstream == "" || *fromFixer == "" || *toFixer == "" {
		return fmt.Errorf("provenance flags -from-branch, -to-branch, -from-upstream, -to-upstream, -from-fixer, and -to-fixer are required")
	}
	data, err := changelog.Generate(changelog.Config{
		FromDir:      *fromPath,
		ToDir:        *toPath,
		FromBranch:   *fromBranch,
		ToBranch:     *toBranch,
		FromUpstream: *fromUpstream,
		ToUpstream:   *toUpstream,
		FromFixer:    *fromFixer,
		ToFixer:      *toFixer,
	})
	if err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Dir(*outPath), 0o755); err != nil {
		return fmt.Errorf("create changelog directory: %w", err)
	}
	if err := os.WriteFile(*outPath, data, 0o644); err != nil {
		return fmt.Errorf("write changelog: %w", err)
	}
	return nil
}

func runUpdateGuide(args []string) error {
	fs := flag.NewFlagSet("update-guide", flag.ContinueOnError)
	changelogPath := fs.String("changelog", "", "human-readable protocol changelog Markdown")
	schemasPath := fs.String("schemas", "", "target corrected Mojang JSON schema directory")
	outPath := fs.String("out", "", "gophertunnel update guide Markdown output")
	if err := fs.Parse(args); err != nil {
		return err
	}
	if *changelogPath == "" || *schemasPath == "" || *outPath == "" {
		return fmt.Errorf("-changelog, -schemas, and -out are required")
	}
	changelog, err := os.ReadFile(*changelogPath)
	if err != nil {
		return fmt.Errorf("read changelog: %w", err)
	}
	guide, err := updateguide.Generate(changelog, *schemasPath)
	if err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Dir(*outPath), 0o755); err != nil {
		return fmt.Errorf("create update guide directory: %w", err)
	}
	if err := os.WriteFile(*outPath, guide, 0o644); err != nil {
		return fmt.Errorf("write update guide: %w", err)
	}
	return nil
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
	directionsPath := fs.String("directions", "", "reviewed packet-direction JSON")
	nbtEncodingsPath := fs.String("nbt-encodings", "", "reviewed per-field NBT encoding JSON")
	outPath := fs.String("out", "", "canonical manifest v2 output")
	if err := fs.Parse(args); err != nil {
		return err
	}
	if *lockPath == "" || *directionsPath == "" || *nbtEncodingsPath == "" || *outPath == "" {
		return fmt.Errorf("-lock, -directions, -nbt-encodings, and -out are required")
	}
	lock, err := sourcelock.Load(*lockPath)
	if err != nil {
		return err
	}
	directions, err := direction.Load(*directionsPath)
	if err != nil {
		return err
	}
	nbtEncodings, err := nbtencoding.Load(*nbtEncodingsPath)
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
	result, err := reconcile.ReconcileWithDirectionsAndNBT(lock.Target, results, adjudications, directions, nbtEncodings)
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
	out := fs.String("out", "", "generated Go source directory")
	namingPath := fs.String("naming", "", "reviewed naming overlay JSON; defaults to naming.json beside the manifest")
	domainsPath := fs.String("domains", "", "reviewed domain overlay JSON; defaults to domains.json beside the manifest")
	docsPath := fs.String("docs", "", "reviewed documentation overlay JSON; defaults to docs.json beside the manifest")
	protocolImport := fs.String("protocol-import", "", "import path of the generated protocol package")
	nativeTypes := fs.Bool("native-types", true, "map canonical semantic shapes to established Go types such as uuid.UUID and mgl32 vectors")
	packetRuntime := fs.Bool("packet-runtime", true, "emit the packet interface and ID methods")
	packetPools := fs.Bool("packet-pools", true, "emit packet factory pools")
	if err := fs.Parse(args); err != nil {
		return err
	}
	if *manifestPath == "" || *out == "" {
		return fmt.Errorf("-manifest and -out are required")
	}
	if *protocolImport == "" {
		return fmt.Errorf("-protocol-import is required")
	}
	result, err := emitter.Run(emitter.Config{
		ManifestPath: *manifestPath,
		NamingPath:   *namingPath,
		DomainsPath:  *domainsPath,
		DocsPath:     *docsPath,
		OutputDir:    *out,
	}, emitter.Func(func(input emitter.Input) (map[string]string, error) {
		return emitgo.GenerateWithOptions(input.Manifest, emitgo.Options{
			ProtocolImportPath: *protocolImport,
			Naming:             input.Naming,
			Domains:            input.Domains,
			Docs:               input.Docs,
			NativeTypes:        *nativeTypes,
			EmitPacketRuntime:  *packetRuntime,
			EmitPacketPools:    *packetPools,
		})
	}))
	if err != nil {
		return err
	}
	fmt.Printf("Go emitter: %d files -> %s\n", result.FileCount, *out)
	fmt.Printf("Go docs coverage: types %d/%d, fields %d/%d\n", result.Coverage.TypesDocumented, result.Coverage.TypesTotal, result.Coverage.FieldsDocumented, result.Coverage.FieldsTotal)
	return nil
}

func runEmitRust(args []string) error {
	fs := flag.NewFlagSet("emit-rust", flag.ContinueOnError)
	manifestPath := fs.String("manifest", "", "canonical manifest v2 JSON")
	out := fs.String("out", "", "generated Rust source directory")
	namingPath := fs.String("naming", "", "reviewed naming overlay JSON; defaults to naming.json beside the manifest")
	domainsPath := fs.String("domains", "", "reviewed domain overlay JSON; defaults to domains.json beside the manifest")
	docsPath := fs.String("docs", "", "reviewed documentation overlay JSON; defaults to docs.json beside the manifest")
	if err := fs.Parse(args); err != nil {
		return err
	}
	if *manifestPath == "" || *out == "" {
		return fmt.Errorf("-manifest and -out are required")
	}
	result, err := emitter.Run(emitter.Config{
		ManifestPath: *manifestPath,
		NamingPath:   *namingPath,
		DomainsPath:  *domainsPath,
		DocsPath:     *docsPath,
		OutputDir:    *out,
	}, emitter.Func(func(input emitter.Input) (map[string]string, error) {
		return emitrust.GenerateFilesWithOptions(input.Manifest, emitrust.Options{Naming: input.Naming, Domains: input.Domains, Docs: input.Docs})
	}))
	if err != nil {
		return err
	}
	fmt.Printf("Rust emitter: %d files -> %s\n", result.FileCount, *out)
	fmt.Printf("Rust docs coverage: types %d/%d, fields %d/%d\n", result.Coverage.TypesDocumented, result.Coverage.TypesTotal, result.Coverage.FieldsDocumented, result.Coverage.FieldsTotal)
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

func runVerifyGophertunnel(args []string) error {
	fs := flag.NewFlagSet("verify-gophertunnel", flag.ContinueOnError)
	manifestPath := fs.String("manifest", "", "canonical manifest v2 JSON")
	lockPath := fs.String("lock", "tools/gophertunnel-oracle/lock.json", "gophertunnel oracle lock JSON")
	acceptedPath := fs.String("accepted", "tools/gophertunnel-oracle/accepted-divergences.json", "reviewed accepted-divergences JSON")
	gophertunnelPath := fs.String("gophertunnel", "", "existing gophertunnel checkout at the locked commit")
	cacheDir := fs.String("cache-dir", "", "cache directory for a runtime clone when -gophertunnel is omitted")
	reportPath := fs.String("report", "gophertunnel-report.json", "machine-readable report output JSON")
	failOnUnaccepted := fs.Bool("fail-on-unaccepted", true, "exit non-zero when a divergence is not in the accepted baseline")
	if err := fs.Parse(args); err != nil {
		return err
	}
	if *manifestPath == "" {
		return fmt.Errorf("-manifest is required")
	}
	report, err := gophertunneloracle.CompareFile(gophertunneloracle.Options{
		ManifestPath:     *manifestPath,
		LockPath:         *lockPath,
		AcceptedPath:     *acceptedPath,
		GophertunnelPath: *gophertunnelPath,
		CacheDir:         *cacheDir,
		ReportPath:       *reportPath,
		FailOnUnaccepted: *failOnUnaccepted,
	})
	if report.SchemaVersion != 0 {
		fmt.Print(gophertunneloracle.Summary(report, *reportPath))
	}
	return err
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
