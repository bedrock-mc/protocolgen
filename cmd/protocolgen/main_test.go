package main

import (
	"strings"
	"testing"
)

func TestRunChangelogRequiresProvenanceFlags(t *testing.T) {
	err := runChangelog([]string{"-from", "old", "-to", "new", "-out", "changes.md"})
	if err == nil || !strings.Contains(err.Error(), "provenance") {
		t.Fatalf("runChangelog error = %v, want missing provenance error", err)
	}
}
