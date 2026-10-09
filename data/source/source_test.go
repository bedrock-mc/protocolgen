package source

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestSourceLockAuthenticatesLocalInputs(t *testing.T) {
	dir := t.TempDir()
	payload := []byte("synthetic input")
	if err := os.WriteFile(filepath.Join(dir, "input.json"), payload, 0o644); err != nil {
		t.Fatal(err)
	}
	lock := Lock{SchemaVersion: 1, MinecraftVersion: "fixture", ProtocolVersion: 1, Inputs: map[string]Input{
		"fixture": {Path: "input.json", SHA256: fmt.Sprintf("%x", sha256.Sum256(payload)), Revision: "fixture"},
	}}
	encoded, _ := json.Marshal(lock)
	path := filepath.Join(dir, "lock.json")
	if err := os.WriteFile(path, encoded, 0o644); err != nil {
		t.Fatal(err)
	}
	sources, err := Open(path)
	if err != nil {
		t.Fatal(err)
	}
	if got, err := sources.Read(context.Background(), "fixture", ""); err != nil || string(got) != string(payload) {
		t.Fatalf("verified read = %q, %v", got, err)
	}
	if err := os.WriteFile(filepath.Join(dir, "input.json"), []byte("changed"), 0o644); err != nil {
		t.Fatal(err)
	}
	if _, err := sources.Read(context.Background(), "fixture", ""); err == nil || !strings.Contains(err.Error(), "digest") {
		t.Fatalf("changed source was accepted: %v", err)
	}
	if _, err := sources.Read(context.Background(), "unknown", ""); err == nil {
		t.Fatal("unlocked source was accepted")
	}
}

func TestBundledSourceInputsAndTarget(t *testing.T) {
	sources, err := Open("")
	if err != nil {
		t.Fatal(err)
	}
	if sources.Lock.MinecraftVersion != "1.26.50" || sources.Lock.ProtocolVersion != 2193 {
		t.Fatalf("unexpected migrated target: %+v", sources.Lock)
	}
	for name, input := range sources.Lock.Inputs {
		if input.Path != "" {
			if _, err := sources.Read(context.Background(), name, ""); err != nil {
				t.Fatal(err)
			}
		}
	}
}
