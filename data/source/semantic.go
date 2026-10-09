package source

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// SemanticInputs pins the local files read by the semantic generator. Only
// hashes are published for BDS inputs; the behavior packs stay outside the repo.
type SemanticInputs struct {
	CloudburstRef   string            `json:"cloudburst_ref"`
	BDSVersion      string            `json:"bds_version"`
	CloudburstFiles map[string]string `json:"cloudburst_files"`
	BDSFiles        map[string]string `json:"bds_files"`
}

// ValidateSemantic authenticates every local file consumed by the semantic
// generator. Added and missing JSON files in its input directories are rejected.
func (s *Sources) ValidateSemantic(cloudburstDir, bdsDir string) error {
	semantic := s.Lock.Semantic
	if semantic.CloudburstRef == "" || semantic.BDSVersion == "" {
		return fmt.Errorf("source lock has no complete semantic source identity")
	}
	if err := s.ValidateCloudburst(cloudburstDir); err != nil {
		return err
	}
	return verifySemanticTree("BDS", bdsDir, semantic.BDSFiles,
		"behavior_packs/experimental_vanilla_shapes/shapes/*.json",
		"behavior_packs/vanilla/entities/*.json",
		"behavior_packs/vanilla/items/*.json")
}

// ValidateCloudburst verifies the complete block and biome inputs independently of BDS.
func (s *Sources) ValidateCloudburst(dir string) error {
	if s.Lock.Semantic.CloudburstRef == "" {
		return fmt.Errorf("source lock has no Cloudburst revision")
	}
	return verifySemanticTree("Cloudburst", dir, s.Lock.Semantic.CloudburstFiles,
		"blocks.json", "stripped_biome_definitions.json")
}

// verifySemanticTree checks the exact input set and the content of each file.
func verifySemanticTree(name, root string, manifest map[string]string, patterns ...string) error {
	if root == "" || len(manifest) == 0 {
		return fmt.Errorf("%s directory and locked file manifest are required", name)
	}
	actual := make(map[string]bool)
	for _, pattern := range patterns {
		paths, err := filepath.Glob(filepath.Join(root, filepath.FromSlash(pattern)))
		if err != nil {
			return fmt.Errorf("list %s inputs: %w", name, err)
		}
		if len(paths) == 0 {
			return fmt.Errorf("%s input %s is missing", name, pattern)
		}
		for _, path := range paths {
			relative, err := filepath.Rel(root, path)
			if err != nil {
				return err
			}
			relative = filepath.ToSlash(relative)
			if _, ok := manifest[relative]; !ok {
				return fmt.Errorf("%s input %s is not in the source lock", name, relative)
			}
			actual[relative] = true
		}
	}
	paths := make([]string, 0, len(manifest))
	for path := range manifest {
		paths = append(paths, path)
	}
	sort.Strings(paths)
	for _, path := range paths {
		digest, err := hex.DecodeString(manifest[path])
		if !filepath.IsLocal(path) || strings.Contains(path, "\\") || err != nil || len(digest) != sha256.Size {
			return fmt.Errorf("%s input %q has an invalid locked path or digest", name, path)
		}
		if !actual[path] {
			return fmt.Errorf("%s locked input %s is missing or outside the generator's input set", name, path)
		}
		data, err := os.ReadFile(filepath.Join(root, filepath.FromSlash(path)))
		if err != nil {
			return fmt.Errorf("read %s input %s: %w", name, path, err)
		}
		if got := fmt.Sprintf("%x", sha256.Sum256(data)); got != manifest[path] {
			return fmt.Errorf("%s input %s digest is %s, want %s", name, path, got, manifest[path])
		}
	}
	return nil
}
