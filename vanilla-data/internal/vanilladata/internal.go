package vanilladata

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/sandertv/gophertunnel/minecraft/nbt"
)

const internalDataManifestVersion = 1

// EndstoneSource pins the Endstone source and the small headless-export patch
// applied to it by CI. Endstone's code and patch must both match the BDS build
// before any data is accepted.
type EndstoneSource struct {
	Repository       string `json:"repository"`
	Revision         string `json:"revision"`
	BDSVersion       string `json:"bds_version"`
	HeadlessPatchSHA string `json:"headless_patch_sha256"`
}

// InternalDataFile authenticates one file written by the Endstone exporter.
type InternalDataFile struct {
	File   string `json:"file"`
	Bytes  int    `json:"bytes"`
	SHA256 string `json:"sha256"`
}

// InternalDataManifest is written by the headless Endstone exporter beside
// its outputs. It is deliberately separate from capture.json so the exporter
// can be verified before the packet bot atomically installs the final output.
type InternalDataManifest struct {
	SchemaVersion int                `json:"schema_version"`
	Target        Target             `json:"target"`
	BDSVersion    string             `json:"bds_version"`
	Endstone      EndstoneSource     `json:"endstone"`
	Files         []InternalDataFile `json:"files"`
}

// InternalDataFiles are the stable, documented files produced by Endstone's
// vanilla-data exporter. Unknown files are rejected so a source change cannot
// silently widen the checked-in format.
var internalDataFiles = map[string]struct{}{
	"block_palette.nbt":    {},
	"block_states.json":    {},
	"block_types.json":     {},
	"block_tags.json":      {},
	"items.json":           {},
	"item_tags.json":       {},
	"creative_groups.json": {},
	"biomes.json":          {},
	"item_components.nbt":  {},
	"creative_items.nbt":   {},
	"recipes.json":         {},
}

// LoadInternalArtifacts verifies an Endstone export and returns the raw
// files plus PMMP's canonical concatenated block-state representation. The
// latter is the only block-state format emitted by this package; legacy meta
// and block-item maps are intentionally not guessed from runtime hashes.
func LoadInternalArtifacts(dir string, target Target, bdsVersion string, expected EndstoneSource) (map[string][]byte, InternalDataManifest, error) {
	manifestPath := filepath.Join(dir, "endstone-export.json")
	data, err := os.ReadFile(manifestPath)
	if err != nil {
		return nil, InternalDataManifest{}, fmt.Errorf("read Endstone export manifest: %w", err)
	}
	var manifest InternalDataManifest
	if err := json.Unmarshal(data, &manifest); err != nil {
		return nil, InternalDataManifest{}, fmt.Errorf("decode Endstone export manifest: %w", err)
	}
	if err := validateInternalManifest(manifest, target, bdsVersion, expected); err != nil {
		return nil, InternalDataManifest{}, err
	}
	if len(manifest.Files) == 0 {
		return nil, InternalDataManifest{}, fmt.Errorf("Endstone export manifest has no files")
	}
	files := make(map[string][]byte, len(manifest.Files)+1)
	for _, entry := range manifest.Files {
		if err := validateInternalFileName(entry.File); err != nil {
			return nil, InternalDataManifest{}, err
		}
		if _, exists := files[entry.File]; exists {
			return nil, InternalDataManifest{}, fmt.Errorf("Endstone export manifest contains duplicate file %q", entry.File)
		}
		if entry.Bytes <= 0 || !validSHA256(entry.SHA256) {
			return nil, InternalDataManifest{}, fmt.Errorf("Endstone export file %s has incomplete digest", entry.File)
		}
		path := filepath.Join(dir, filepath.FromSlash(entry.File))
		value, err := os.ReadFile(path)
		if err != nil {
			return nil, InternalDataManifest{}, fmt.Errorf("read Endstone export %s: %w", entry.File, err)
		}
		if len(value) != entry.Bytes {
			return nil, InternalDataManifest{}, fmt.Errorf("Endstone export file %s has %d bytes, manifest says %d", entry.File, len(value), entry.Bytes)
		}
		digest := sha256.Sum256(value)
		actual := "sha256:" + hex.EncodeToString(digest[:])
		if actual != entry.SHA256 {
			return nil, InternalDataManifest{}, fmt.Errorf("Endstone export file %s SHA-256 %s, manifest says %s", entry.File, actual, entry.SHA256)
		}
		files[entry.File] = value
	}
	palette, ok := files["block_palette.nbt"]
	if !ok {
		return nil, InternalDataManifest{}, fmt.Errorf("Endstone export is missing required block_palette.nbt")
	}
	canonical, err := canonicalBlockStates(palette)
	if err != nil {
		return nil, InternalDataManifest{}, fmt.Errorf("normalize Endstone block palette: %w", err)
	}
	files["canonical_block_states.nbt"] = canonical
	return files, manifest, nil
}

func validateInternalManifest(manifest InternalDataManifest, target Target, bdsVersion string, expected EndstoneSource) error {
	if manifest.SchemaVersion != internalDataManifestVersion {
		return fmt.Errorf("unsupported Endstone export schema %d", manifest.SchemaVersion)
	}
	if manifest.Target.MinecraftVersion != target.MinecraftVersion || manifest.Target.ProtocolVersion != target.ProtocolVersion {
		return fmt.Errorf("Endstone export target %s/%d does not match capture target %s/%d", manifest.Target.MinecraftVersion, manifest.Target.ProtocolVersion, target.MinecraftVersion, target.ProtocolVersion)
	}
	if manifest.BDSVersion != bdsVersion {
		return fmt.Errorf("Endstone export BDS version %s does not match capture source %s", manifest.BDSVersion, bdsVersion)
	}
	if manifest.Endstone != expected {
		return fmt.Errorf("Endstone export source pin does not match the version-locked capture source")
	}
	if expected.Repository == "" || !validHex(expected.Revision, 40) || expected.BDSVersion != bdsVersion || !validSHA256(expected.HeadlessPatchSHA) {
		return fmt.Errorf("version-locked Endstone source has incomplete provenance")
	}
	return nil
}

func validateInternalFiles(manifest InternalDataManifest, files map[string][]byte) error {
	manifestFiles := make(map[string]InternalDataFile, len(manifest.Files))
	for _, entry := range manifest.Files {
		manifestFiles[entry.File] = entry
	}
	for _, entry := range manifest.Files {
		value, ok := files[entry.File]
		if !ok {
			return fmt.Errorf("internal data file %s is missing from the capture input", entry.File)
		}
		if len(value) != entry.Bytes {
			return fmt.Errorf("internal data file %s has %d bytes, manifest says %d", entry.File, len(value), entry.Bytes)
		}
		digest := sha256.Sum256(value)
		actual := "sha256:" + hex.EncodeToString(digest[:])
		if actual != entry.SHA256 {
			return fmt.Errorf("internal data file %s SHA-256 %s, manifest says %s", entry.File, actual, entry.SHA256)
		}
	}
	for name := range files {
		if name == "canonical_block_states.nbt" {
			continue
		}
		if _, ok := manifestFiles[name]; !ok {
			return fmt.Errorf("internal data file %s is not authenticated by the Endstone export manifest", name)
		}
	}
	if palette, ok := files["block_palette.nbt"]; ok {
		canonical, ok := files["canonical_block_states.nbt"]
		if !ok {
			return fmt.Errorf("internal data is missing normalized canonical_block_states.nbt")
		}
		want, err := canonicalBlockStates(palette)
		if err != nil {
			return fmt.Errorf("validate normalized block palette: %w", err)
		}
		if !bytes.Equal(canonical, want) {
			return fmt.Errorf("canonical_block_states.nbt does not match block_palette.nbt")
		}
	}
	return nil
}

func validateInternalFileName(name string) error {
	clean := filepath.ToSlash(filepath.Clean(name))
	if name == "" || clean != name || strings.Contains(name, "\\") || filepath.Base(clean) != clean {
		return fmt.Errorf("invalid Endstone export file name %q", name)
	}
	if _, ok := internalDataFiles[name]; !ok {
		return fmt.Errorf("unsupported Endstone export file %q", name)
	}
	return nil
}

func validSHA256(value string) bool {
	if !strings.HasPrefix(value, "sha256:") {
		return false
	}
	return validHex(strings.TrimPrefix(value, "sha256:"), 64)
}

func canonicalBlockStates(encoded []byte) ([]byte, error) {
	reader, err := gzip.NewReader(bytes.NewReader(encoded))
	if err != nil {
		return nil, fmt.Errorf("open gzip palette: %w", err)
	}
	defer reader.Close()
	var root map[string]any
	if err := nbt.NewDecoderWithEncoding(reader, nbt.BigEndian).Decode(&root); err != nil {
		return nil, fmt.Errorf("decode big-endian palette: %w", err)
	}
	blocks, ok := root["blocks"].([]any)
	if !ok || len(blocks) == 0 {
		return nil, fmt.Errorf("palette blocks has type %T or is empty", root["blocks"])
	}
	var out bytes.Buffer
	for index, raw := range blocks {
		entry, ok := raw.(map[string]any)
		if !ok {
			return nil, fmt.Errorf("palette block %d has type %T", index, raw)
		}
		name, ok := entry["name"].(string)
		if !ok || name == "" {
			return nil, fmt.Errorf("palette block %d has invalid name", index)
		}
		states, ok := entry["states"].(map[string]any)
		if !ok {
			return nil, fmt.Errorf("palette block %d has states of type %T", index, entry["states"])
		}
		version, ok := entry["version"].(int32)
		if !ok {
			return nil, fmt.Errorf("palette block %d has version of type %T", index, entry["version"])
		}
		value := map[string]any{
			"name":    name,
			"states":  states,
			"version": version,
		}
		if err := writeNBTTag(&out, 10, "", value); err != nil {
			return nil, fmt.Errorf("encode palette block %d: %w", index, err)
		}
	}
	return out.Bytes(), nil
}

// MarshalInternalManifest writes the exporter manifest with stable key order
// and a trailing newline for CI artifacts and fixture tests.
func MarshalInternalManifest(manifest InternalDataManifest) ([]byte, error) {
	data, err := json.MarshalIndent(manifest, "", "  ")
	if err != nil {
		return nil, err
	}
	return append(data, '\n'), nil
}

func DigestInternalFiles(files map[string][]byte) []InternalDataFile {
	names := make([]string, 0, len(files))
	for name := range files {
		names = append(names, name)
	}
	sort.Strings(names)
	result := make([]InternalDataFile, 0, len(names))
	for _, name := range names {
		digest := sha256.Sum256(files[name])
		result = append(result, InternalDataFile{File: name, Bytes: len(files[name]), SHA256: "sha256:" + hex.EncodeToString(digest[:])})
	}
	return result
}
