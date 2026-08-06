package sourcelock

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestLoadRejectsMixedProtocolPins(t *testing.T) {
	path := filepath.Join(t.TempDir(), "source-lock.json")
	data := []byte(`{
  "schema_version": 1,
  "target": {"minecraft_version": "1.26.40", "protocol_version": 2168},
  "sources": [
    {"id":"mojang-2168","kind":"mojang","revision":"a","digest":"sha256:a","minecraft_version":"1.26.40","protocol_version":2168},
    {"id":"endstone-2169","kind":"endstone","revision":"b","digest":"sha256:b","minecraft_version":"1.26.50","protocol_version":2169}
  ]
}`)
	if err := os.WriteFile(path, data, 0o644); err != nil {
		t.Fatal(err)
	}
	if _, err := Load(path); err == nil || !strings.Contains(err.Error(), "mix") {
		t.Fatalf("Load error = %v, want mixed target rejection", err)
	}
}

func TestDirectoryDigestIsStableAndExcludesGitMetadata(t *testing.T) {
	root := t.TempDir()
	if err := os.Mkdir(filepath.Join(root, ".git"), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(root, "schema.json"), []byte(`{"x":1}`), 0o644); err != nil {
		t.Fatal(err)
	}
	first, err := DigestDirectory(root)
	if err != nil {
		t.Fatalf("DigestDirectory: %v", err)
	}
	if err := os.WriteFile(filepath.Join(root, ".git", "HEAD"), []byte("unrelated"), 0o644); err != nil {
		t.Fatal(err)
	}
	second, err := DigestDirectory(root)
	if err != nil {
		t.Fatalf("DigestDirectory second: %v", err)
	}
	if first != second {
		t.Fatalf("digest changed after .git metadata: %q != %q", first, second)
	}
}
