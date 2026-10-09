package source

import (
	"crypto/sha256"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestValidateSemanticRejectsChangedInputSets(t *testing.T) {
	for _, test := range []struct {
		name   string
		change func(t *testing.T, sources *Sources, cloudburst, bds string)
		want   string
	}{
		{name: "unchanged"},
		{name: "modified Cloudburst", change: func(t *testing.T, _ *Sources, cloudburst, _ string) {
			if err := os.WriteFile(filepath.Join(cloudburst, "block_properties.json"), []byte("[]"), 0o644); err != nil {
				t.Fatal(err)
			}
		}, want: "digest is"},
		{name: "modified BDS", change: func(t *testing.T, _ *Sources, _, bds string) {
			if err := os.WriteFile(filepath.Join(bds, "behavior_packs/vanilla/entities/arrow.json"), []byte("{}"), 0o644); err != nil {
				t.Fatal(err)
			}
		}, want: "digest is"},
		{name: "additional entity", change: func(t *testing.T, _ *Sources, _, bds string) {
			if err := os.WriteFile(filepath.Join(bds, "behavior_packs/vanilla/entities/new.json"), []byte("{}"), 0o644); err != nil {
				t.Fatal(err)
			}
		}, want: "not in the source lock"},
		{name: "missing entity", change: func(t *testing.T, _ *Sources, _, bds string) {
			if err := os.Remove(filepath.Join(bds, "behavior_packs/vanilla/entities/arrow.json")); err != nil {
				t.Fatal(err)
			}
		}, want: "is missing"},
		{name: "unlocked file", change: func(_ *testing.T, sources *Sources, _, _ string) {
			delete(sources.Lock.Semantic.BDSFiles, "behavior_packs/vanilla/entities/arrow.json")
		}, want: "not in the source lock"},
		{name: "invalid path", change: func(_ *testing.T, sources *Sources, _, _ string) {
			sources.Lock.Semantic.BDSFiles["../outside.json"] = strings.Repeat("0", 64)
		}, want: "invalid locked path"},
		{name: "missing identity", change: func(_ *testing.T, sources *Sources, _, _ string) {
			sources.Lock.Semantic.BDSVersion = ""
		}, want: "no complete semantic source identity"},
	} {
		t.Run(test.name, func(t *testing.T) {
			cloudburst, cloudburstFiles := semanticFixtureTree(t, "cloudburst")
			bds, bdsFiles := semanticFixtureTree(t, "bds")
			sources := &Sources{Lock: Lock{Semantic: SemanticInputs{
				CloudburstRef: "fixture-cloudburst", BDSVersion: "fixture-bds",
				CloudburstFiles: cloudburstFiles, BDSFiles: bdsFiles,
			}}}
			if test.change != nil {
				test.change(t, sources, cloudburst, bds)
			}
			err := sources.ValidateSemantic(cloudburst, bds)
			if test.want == "" && err != nil {
				t.Fatal(err)
			}
			if test.want != "" && (err == nil || !strings.Contains(err.Error(), test.want)) {
				t.Fatalf("ValidateSemantic error = %v, want %q", err, test.want)
			}
		})
	}
}

// semanticFixtureTree copies a synthetic generator tree and hashes its files.
func semanticFixtureTree(t *testing.T, name string) (string, map[string]string) {
	t.Helper()
	root := t.TempDir()
	if err := os.CopyFS(root, os.DirFS(filepath.Join("..", "internal", "generator", "testdata", name))); err != nil {
		t.Fatal(err)
	}
	manifest := make(map[string]string)
	if err := filepath.WalkDir(root, func(path string, entry fs.DirEntry, err error) error {
		if err != nil || entry.IsDir() {
			return err
		}
		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		relative, err := filepath.Rel(root, path)
		if err != nil {
			return err
		}
		manifest[filepath.ToSlash(relative)] = fmt.Sprintf("%x", sha256.Sum256(data))
		return nil
	}); err != nil {
		t.Fatal(err)
	}
	return root, manifest
}
