// Command protocolgen generates Go packet codecs from Mojang's
// bedrock-protocol-docs JSON.
//
// Two output formats (-format):
//
//   - gophertunnel (default): maps onto gophertunnel's existing protocol.* types
//     and conventions, for diffing against / merging into the hand-written
//     packets. Supports -overwrite (in-place merge).
//   - raw: a self-contained package — faithful structs for every packet and every
//     composite definition, marshalling against a generated IO interface, with no
//     gophertunnel dependency.
//
// Fields the docs cannot express (oneOf unions, unmapped types, NBT, optionals)
// are emitted best-effort with an explicit // TODO marker.
//
// The docs are under the Minecraft EULA; pass a local clone path via -docs and do
// not vendor their content.
package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func main() {
	docsDir := flag.String("docs", "", "path to the Mojang bedrock-protocol-docs json/ directory (local clone)")
	packetsDir := flag.String("packets", "minecraft/protocol/packet", "gophertunnel packet directory (for id.go); gophertunnel format only")
	outDir := flag.String("out", "generated", "directory to write generated files into")
	formatName := flag.String("format", "gophertunnel", "output format: gophertunnel | raw")
	rawPkg := flag.String("raw-pkg", "raw", "package name for -format=raw output")
	overwrite := flag.Bool("overwrite", false, "gophertunnel format only: overwrite existing packet files in -packets in place (matched by packet id) for git-diff review")
	preserveMatching := flag.Bool("preserve-matching", true, "overwrite mode: keep gophertunnel's existing field name/type/op when the wire encoding is unchanged (skips pure renames)")
	flag.Parse()

	if *docsDir == "" {
		fmt.Fprintln(os.Stderr, "error: -docs is required")
		flag.Usage()
		os.Exit(2)
	}
	format, ok := selectFormat(*formatName, *rawPkg)
	if !ok {
		fmt.Fprintf(os.Stderr, "error: unknown -format %q (want gophertunnel|raw)\n", *formatName)
		os.Exit(2)
	}
	if *overwrite && format.Name() != "gophertunnel" {
		fmt.Fprintln(os.Stderr, "error: -overwrite is only supported for -format=gophertunnel")
		os.Exit(2)
	}

	// gophertunnel ID constants (raw uses numeric ids directly).
	idByNum := map[int]string{}
	if format.Name() == "gophertunnel" {
		if m, err := loadIDConsts(filepath.Join(*packetsDir, "id.go")); err != nil {
			fmt.Fprintln(os.Stderr, "warning: could not load id.go, ID() will use guessed constants:", err)
		} else {
			idByNum = m
		}
	}

	idToFile := map[int]string{}
	if *overwrite {
		nameToNum := map[string]int{}
		for num, name := range idByNum {
			nameToNum[name] = num
		}
		var err error
		if idToFile, err = existingPacketFiles(*packetsDir, nameToNum); err != nil {
			fmt.Fprintln(os.Stderr, "error:", err)
			os.Exit(2)
		}
	} else if err := os.MkdirAll(*outDir, 0o755); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(2)
	}

	entries, err := os.ReadDir(*docsDir)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(2)
	}

	var total, clean, withTodos, failed int
	todoReasons := map[string]int{}
	var dirty []string
	var packets []genPacket

	for _, e := range entries {
		name := e.Name()
		if e.IsDir() || !strings.HasSuffix(name, ".json") || strings.HasPrefix(name, "__") {
			continue
		}
		b, err := os.ReadFile(filepath.Join(*docsDir, name))
		if err != nil {
			fmt.Fprintf(os.Stderr, "read %s: %v\n", name, err)
			failed++
			continue
		}
		pk, err := parseGenPacket(b)
		if err != nil {
			fmt.Fprintf(os.Stderr, "parse %s: %v\n", name, err)
			failed++
			continue
		}
		if pk.TypeName == "" {
			continue
		}
		total++
		packets = append(packets, pk)

		var (
			src     string
			todos   []string
			emitErr error
			out     = filepath.Join(*outDir, fileName(pk.TypeName))
		)
		if existing, ok := idToFile[pk.ID]; *overwrite && ok {
			b, rerr := os.ReadFile(existing)
			if rerr == nil {
				src, todos, emitErr = mergeIntoExisting(b, pk, *preserveMatching)
			}
			if rerr != nil || src == "" {
				src, todos, emitErr = emitPacket(pk, idByNum[pk.ID])
			}
			out = existing
		} else {
			src, todos, emitErr = format.EmitPacket(pk, idByNum[pk.ID])
			if *overwrite {
				out = filepath.Join(*packetsDir, fileName(pk.TypeName))
			}
		}
		if err := os.WriteFile(out, []byte(src), 0o644); err != nil {
			fmt.Fprintf(os.Stderr, "write %s: %v\n", out, err)
			failed++
			continue
		}
		if emitErr != nil {
			fmt.Fprintf(os.Stderr, "note: %s emitted but not gofmt-clean: %v\n", pk.TypeName, emitErr)
		}
		if len(todos) == 0 {
			clean++
		} else {
			withTodos++
			dirty = append(dirty, fmt.Sprintf("%-34s %d TODO", pk.TypeName, len(todos)))
			for _, td := range todos {
				if i := strings.LastIndex(td, ": "); i >= 0 {
					todoReasons[td[i+2:]]++
				}
			}
		}
	}

	// Shared files (raw: composite type definitions + the IO interface).
	comps, err := collectComposites(*docsDir)
	if err != nil {
		fmt.Fprintln(os.Stderr, "warning: collect composites:", err)
	}
	extra, eerr := format.EmitExtra(comps, packets)
	if eerr != nil {
		fmt.Fprintln(os.Stderr, "note: extra files not gofmt-clean:", eerr)
	}
	for fname, content := range extra {
		if err := os.WriteFile(filepath.Join(*outDir, fname), []byte(content), 0o644); err != nil {
			fmt.Fprintf(os.Stderr, "write %s: %v\n", fname, err)
		}
	}

	sort.Strings(dirty)
	target := *outDir + "/"
	if *overwrite {
		target = *packetsDir + "/ (in place)"
	}
	fmt.Printf("\nFormat %s: generated %d packets", format.Name(), total)
	if len(extra) > 0 {
		fmt.Printf(" + %d shared files (%d composites)", len(extra), len(comps))
	}
	fmt.Printf(" into %s (%d clean, %d need hand-finishing, %d failed)\n", target, clean, withTodos, failed)
	fmt.Println("\nTODO reasons (field count):")
	for _, r := range sortedByCount(todoReasons) {
		fmt.Printf("  %4d  %s\n", todoReasons[r], r)
	}
}

func loadIDConsts(idPath string) (map[int]string, error) {
	b, err := os.ReadFile(idPath)
	if err != nil {
		return nil, err
	}
	return packetIDsByNumber(b)
}

func sortedByCount(m map[string]int) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		if m[keys[i]] != m[keys[j]] {
			return m[keys[i]] > m[keys[j]]
		}
		return keys[i] < keys[j]
	})
	return keys
}
