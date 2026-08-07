package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestWriteFilesSupportsModulesAndRemovesOnlyStaleGeneratedFiles(t *testing.T) {
	directory := t.TempDir()
	stale := filepath.Join(directory, "stale.rs")
	keep := filepath.Join(directory, "notes.txt")
	if err := os.WriteFile(stale, []byte("// Code generated from canonical protocol manifest v2. DO NOT EDIT.\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(keep, []byte("keep me\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := writeFiles(directory, map[string]string{"packets/login.rs": "generated\n"}); err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(filepath.Join(directory, "packets", "login.rs")); err != nil {
		t.Fatalf("nested output was not written: %v", err)
	}
	if _, err := os.Stat(stale); !os.IsNotExist(err) {
		t.Fatalf("stale generated file was not removed: %v", err)
	}
	if _, err := os.Stat(keep); err != nil {
		t.Fatalf("non-generated file was removed: %v", err)
	}
}

func TestWriteFilesRejectsParentTraversal(t *testing.T) {
	if err := writeFiles(t.TempDir(), map[string]string{"../escape.rs": "bad"}); err == nil {
		t.Fatal("writeFiles accepted parent traversal")
	}
}
