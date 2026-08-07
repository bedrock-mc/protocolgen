package direction

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"protocolgen/internal/manifest"
)

func TestLoadRejectsUnknownPacket(t *testing.T) {
	table := validTable()
	table.Packets = []Entry{{ID: 2, Name: "Missing", Direction: DirectionServer, Evidence: validEvidence()}}

	loaded := writeAndLoad(t, table)
	err := loaded.Validate(table.Target, fixturePackets())
	if err == nil || !strings.Contains(err.Error(), "unknown packet") {
		t.Fatalf("Validate error = %v, want unknown packet failure", err)
	}
}

func TestLoadRejectsMissingPacketDirection(t *testing.T) {
	table := validTable()
	table.Packets = nil

	loaded := writeAndLoad(t, table)
	err := loaded.Validate(table.Target, fixturePackets())
	if err == nil || !strings.Contains(err.Error(), "missing direction") {
		t.Fatalf("Validate error = %v, want missing direction failure", err)
	}
}

func TestLoadRejectsBadDirection(t *testing.T) {
	table := validTable()
	table.Packets[0].Direction = Direction("sideways")

	path := filepath.Join(t.TempDir(), "directions.json")
	data, err := json.Marshal(table)
	if err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, data, 0o644); err != nil {
		t.Fatal(err)
	}
	if _, err := Load(path); err == nil || !strings.Contains(err.Error(), "invalid direction") {
		t.Fatalf("Load error = %v, want invalid direction failure", err)
	}
}

func TestApplySetsManifestDirectionsByPacketID(t *testing.T) {
	table := validTable()
	value := manifest.Manifest{SchemaVersion: manifest.SchemaVersion, Target: table.Target, Packets: fixturePackets()}
	if err := table.Apply(&value); err != nil {
		t.Fatalf("Apply: %v", err)
	}
	if got := value.Packets[0].Direction; got != manifest.DirectionClientbound {
		t.Fatalf("direction = %q, want %q", got, manifest.DirectionClientbound)
	}
}

func TestCheckedInDirectionTableMatchesManifest(t *testing.T) {
	table, err := Load(filepath.Join("..", "..", "generated", "1.26.40", "directions.json"))
	if err != nil {
		t.Fatalf("Load checked-in table: %v", err)
	}
	value, err := manifest.Load(filepath.Join("..", "..", "generated", "1.26.40", "manifest.json"))
	if err != nil {
		t.Fatalf("Load checked-in manifest: %v", err)
	}
	if err := table.Validate(value.Target, value.Packets); err != nil {
		t.Fatalf("Validate checked-in table: %v", err)
	}
}

func validTable() Table {
	return Table{
		SchemaVersion: SchemaVersion,
		Target:        manifest.Target{MinecraftVersion: "fixture", ProtocolVersion: 2168},
		Source: Source{
			Repository: "https://github.com/example/gophertunnel",
			Revision:   "0123456789012345678901234567890123456789",
			Locator:    "https://github.com/example/gophertunnel/blob/0123456789012345678901234567890123456789/minecraft/protocol/packet/pool.go",
			SHA256:     "sha256:0123456789012345678901234567890123456789012345678901234567890123",
		},
		Packets: []Entry{{ID: 1, Name: "Fixture", Direction: DirectionServer, Evidence: validEvidence()}},
	}
}

func validEvidence() Evidence {
	return Evidence{Locator: "https://github.com/example/gophertunnel/blob/0123456789012345678901234567890123456789/minecraft/protocol/packet/pool.go", Summary: "The packet ID is registered in the server pool."}
}

func fixturePackets() []manifest.Packet {
	return []manifest.Packet{{ID: 1, Name: "Fixture", Direction: manifest.DirectionUnknown}}
}

func writeAndLoad(t *testing.T, table Table) Table {
	t.Helper()
	path := filepath.Join(t.TempDir(), "directions.json")
	data, err := json.Marshal(table)
	if err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, data, 0o644); err != nil {
		t.Fatal(err)
	}
	loaded, err := Load(path)
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	return loaded
}
