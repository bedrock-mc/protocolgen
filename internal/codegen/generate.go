// Package codegen turns Mojang's bedrock-protocol-docs JSON into Go packet
// codecs in one of two formats (gophertunnel or raw). It is driven by the
// gophertunnel/ and raw/ commands.
package codegen

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// Options configures a Generate run.
type Options struct {
	DocsDir          string // Mojang bedrock-protocol-docs json/ directory (local clone)
	PacketsDir       string // gophertunnel packet directory (id.go + overwrite target); gophertunnel format only
	OutDir           string // directory to write generated files into
	Format           string // "gophertunnel" or "raw"
	RawPkg           string // package name for raw output
	Overwrite        bool   // gophertunnel format: overwrite the packet files in place
	PreserveMatching bool   // overwrite: keep gopher's field when the wire is unchanged
}

// Generate runs the generator and prints a summary. It returns an error for
// fatal conditions.
func Generate(o Options) error {
	if o.DocsDir == "" {
		return fmt.Errorf("-docs is required")
	}
	format, ok := selectFormat(o.Format, o.RawPkg)
	if !ok {
		return fmt.Errorf("unknown format %q (want gophertunnel|raw)", o.Format)
	}
	if o.Overwrite && format.Name() != "gophertunnel" {
		return fmt.Errorf("overwrite is only supported for the gophertunnel format")
	}

	idByNum := map[int]string{}
	if format.Name() == "gophertunnel" {
		if m, err := loadIDConsts(filepath.Join(o.PacketsDir, "id.go")); err != nil {
			fmt.Fprintln(os.Stderr, "warning: could not load id.go, ID() will use guessed constants:", err)
		} else {
			idByNum = m
		}
	}

	idToFile := map[int]string{}
	if o.Overwrite {
		nameToNum := map[string]int{}
		for num, name := range idByNum {
			nameToNum[name] = num
		}
		var err error
		if idToFile, err = existingPacketFiles(o.PacketsDir, nameToNum); err != nil {
			return err
		}
	} else if err := os.MkdirAll(o.OutDir, 0o755); err != nil {
		return err
	}

	entries, err := os.ReadDir(o.DocsDir)
	if err != nil {
		return err
	}

	var total, clean, withTodos, failed int
	todoReasons := map[string]int{}
	var packets []genPacket

	for _, e := range entries {
		name := e.Name()
		if e.IsDir() || !strings.HasSuffix(name, ".json") || strings.HasPrefix(name, "__") {
			continue
		}
		b, err := os.ReadFile(filepath.Join(o.DocsDir, name))
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
			out     = filepath.Join(o.OutDir, fileName(pk.TypeName))
		)
		if existing, ok := idToFile[pk.ID]; o.Overwrite && ok {
			eb, rerr := os.ReadFile(existing)
			if rerr == nil {
				src, todos, emitErr = mergeIntoExisting(eb, pk, o.PreserveMatching)
			}
			if rerr != nil || src == "" {
				src, todos, emitErr = emitPacket(pk, idByNum[pk.ID])
			}
			out = existing
		} else {
			src, todos, emitErr = format.EmitPacket(pk, idByNum[pk.ID])
			if o.Overwrite {
				out = filepath.Join(o.PacketsDir, fileName(pk.TypeName))
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
			for _, td := range todos {
				if i := strings.LastIndex(td, ": "); i >= 0 {
					todoReasons[td[i+2:]]++
				}
			}
		}
	}

	comps, err := collectComposites(o.DocsDir)
	if err != nil {
		fmt.Fprintln(os.Stderr, "warning: collect composites:", err)
	}
	extra, eerr := format.EmitExtra(comps, packets)
	if eerr != nil {
		fmt.Fprintln(os.Stderr, "note: extra files not gofmt-clean:", eerr)
	}
	for fname, content := range extra {
		if err := os.WriteFile(filepath.Join(o.OutDir, fname), []byte(content), 0o644); err != nil {
			fmt.Fprintf(os.Stderr, "write %s: %v\n", fname, err)
		}
	}

	target := o.OutDir + "/"
	if o.Overwrite {
		target = o.PacketsDir + "/ (in place)"
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
	return nil
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
