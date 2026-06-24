package driftcheck

import (
	"fmt"
	"maps"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// loadDocPackets parses every *.json file in dir into a DocPacket.
func loadDocPackets(dir string) ([]DocPacket, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("read docs dir: %w", err)
	}
	var docs []DocPacket
	for _, e := range entries {
		// Skip Mojang's "__"-prefixed non-packet files, e.g. the legacy combined
		// json/__protocoldoc.json (a JSON array, not a per-packet object).
		if e.IsDir() || !strings.HasSuffix(e.Name(), ".json") || strings.HasPrefix(e.Name(), "__") {
			continue
		}
		b, err := os.ReadFile(filepath.Join(dir, e.Name()))
		if err != nil {
			return nil, fmt.Errorf("read %s: %w", e.Name(), err)
		}
		pk, err := parseDocPacket(b)
		if err != nil {
			return nil, fmt.Errorf("parse %s: %w", e.Name(), err)
		}
		docs = append(docs, pk)
	}
	sort.Slice(docs, func(i, j int) bool { return docs[i].ID < docs[j].ID })
	return docs, nil
}

// loadSourcePackets parses every *.go file in dir, extracts packet types and
// their Marshal wire ops, and resolves numeric IDs from id.go.
func loadSourcePackets(dir string) ([]SourcePacket, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("read packets dir: %w", err)
	}

	ids := map[string]int{}
	var packets []SourcePacket
	for _, e := range entries {
		name := e.Name()
		if e.IsDir() || !strings.HasSuffix(name, ".go") || strings.HasSuffix(name, "_test.go") {
			continue
		}
		b, err := os.ReadFile(filepath.Join(dir, name))
		if err != nil {
			return nil, fmt.Errorf("read %s: %w", name, err)
		}
		if name == "id.go" {
			parsed, err := parsePacketIDs(b)
			if err != nil {
				return nil, fmt.Errorf("parse ids in %s: %w", name, err)
			}
			maps.Copy(ids, parsed)
		}
		pks, err := analyzePackets(b)
		if err != nil {
			return nil, fmt.Errorf("analyze %s: %w", name, err)
		}
		packets = append(packets, pks...)
	}

	resolved := packets[:0]
	for i := range packets {
		packets[i].ID = ids[packets[i].IDConst]
		// Drop packets with no static numeric id (e.g. gophertunnel's Unknown,
		// whose id is a runtime value) — they cannot be compared by id.
		if packets[i].ID != 0 {
			resolved = append(resolved, packets[i])
		}
	}
	sort.Slice(resolved, func(i, j int) bool { return resolved[i].ID < resolved[j].ID })
	return resolved, nil
}
