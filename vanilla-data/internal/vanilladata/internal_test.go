package vanilladata

import (
	"compress/gzip"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"github.com/sandertv/gophertunnel/minecraft/nbt"
)

func TestLoadInternalArtifactsNormalizesEndstonePaletteAndVerifiesProvenance(t *testing.T) {
	dir := t.TempDir()
	palette := endstonePalette(t)
	if err := os.WriteFile(filepath.Join(dir, "block_palette.nbt"), palette, 0o644); err != nil {
		t.Fatal(err)
	}
	digest := sha256.Sum256(palette)
	pins := EndstoneSource{
		Repository:       "https://github.com/EndstoneMC/endstone",
		Revision:         strings.Repeat("a", 40),
		BDSVersion:       "1.26.44.3",
		HeadlessPatchSHA: "sha256:" + strings.Repeat("b", 64),
	}
	manifest := InternalDataManifest{
		SchemaVersion: 1,
		Target:        Target{MinecraftVersion: "1.26.44", ProtocolVersion: 2168},
		BDSVersion:    "1.26.44.3",
		Endstone:      pins,
		Files:         []InternalDataFile{{File: "block_palette.nbt", Bytes: len(palette), SHA256: "sha256:" + hex.EncodeToString(digest[:])}},
	}
	writeInternalManifest(t, dir, manifest)

	artifacts, provenance, err := LoadInternalArtifacts(dir, manifest.Target, manifest.BDSVersion, pins)
	if err != nil {
		t.Fatalf("LoadInternalArtifacts: %v", err)
	}
	want := []byte{
		0x0a, 0x00, 0x00, 0x08, 0x04, 0x00, 'n', 'a', 'm', 'e', 0x0f, 0x00,
		'm', 'i', 'n', 'e', 'c', 'r', 'a', 'f', 't', ':', 's', 't', 'o', 'n', 'e',
		0x0a, 0x06, 0x00, 's', 't', 'a', 't', 'e', 's', 0x00, 0x03, 0x07, 0x00,
		'v', 'e', 'r', 's', 'i', 'o', 'n', 0x01, 0x00, 0x00, 0x00,
		0x00,
	}
	if got := artifacts["canonical_block_states.nbt"]; !reflect.DeepEqual(got, want) {
		t.Fatalf("canonical palette = %x, want %x", got, want)
	}
	if got := artifacts["block_palette.nbt"]; !reflect.DeepEqual(got, palette) {
		t.Fatal("raw Endstone palette was not preserved")
	}
	if !reflect.DeepEqual(provenance, manifest) {
		t.Fatalf("provenance = %#v, want %#v", provenance, manifest)
	}
}

func TestLoadInternalArtifactsRejectsVersionOrDigestDrift(t *testing.T) {
	dir := t.TempDir()
	palette := endstonePalette(t)
	if err := os.WriteFile(filepath.Join(dir, "block_palette.nbt"), palette, 0o644); err != nil {
		t.Fatal(err)
	}
	manifest := InternalDataManifest{
		SchemaVersion: 1,
		Target:        Target{MinecraftVersion: "1.26.44", ProtocolVersion: 2168},
		BDSVersion:    "1.26.44.3",
		Endstone: EndstoneSource{
			Repository:       "https://github.com/EndstoneMC/endstone",
			Revision:         strings.Repeat("a", 40),
			BDSVersion:       "1.26.44.3",
			HeadlessPatchSHA: "sha256:" + strings.Repeat("b", 64),
		},
		Files: []InternalDataFile{{File: "block_palette.nbt", Bytes: len(palette), SHA256: "sha256:" + strings.Repeat("c", 64)}},
	}
	writeInternalManifest(t, dir, manifest)
	if _, _, err := LoadInternalArtifacts(dir, manifest.Target, manifest.BDSVersion, manifest.Endstone); err == nil || !strings.Contains(err.Error(), "SHA-256") {
		t.Fatalf("digest drift error = %v", err)
	}
	manifest.Files[0].SHA256 = "sha256:" + strings.Repeat("d", 64)
	writeInternalManifest(t, dir, manifest)
	if _, _, err := LoadInternalArtifacts(dir, Target{MinecraftVersion: "1.26.45", ProtocolVersion: 2168}, manifest.BDSVersion, manifest.Endstone); err == nil || !strings.Contains(err.Error(), "target") {
		t.Fatalf("target drift error = %v", err)
	}
}

func endstonePalette(t *testing.T) []byte {
	t.Helper()
	value := map[string]any{"blocks": []any{
		map[string]any{
			"block_id":   int32(1),
			"name":       "minecraft:stone",
			"name_hash":  int64(2),
			"network_id": int32(3),
			"states":     map[string]any{},
			"version":    int32(1),
		},
	}}
	encoded, err := nbt.MarshalEncoding(value, nbt.BigEndian)
	if err != nil {
		t.Fatal(err)
	}
	var out strings.Builder
	writer := gzip.NewWriter(&out)
	if _, err := writer.Write(encoded); err != nil {
		t.Fatal(err)
	}
	if err := writer.Close(); err != nil {
		t.Fatal(err)
	}
	return []byte(out.String())
}

func writeInternalManifest(t *testing.T, dir string, manifest InternalDataManifest) {
	t.Helper()
	data, err := json.MarshalIndent(manifest, "", "  ")
	if err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, "endstone-export.json"), append(data, '\n'), 0o644); err != nil {
		t.Fatal(err)
	}
}
