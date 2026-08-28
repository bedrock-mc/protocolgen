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

func TestRunReconcileClaimsRequiresClaimFiles(t *testing.T) {
	err := runReconcileClaims([]string{
		"-lock", "source-lock.json",
		"-directions", "directions.json",
		"-nbt-encodings", "nbt-encodings.json",
		"-out", "manifest.json",
	})
	if err == nil || !strings.Contains(err.Error(), "-claims") {
		t.Fatalf("runReconcileClaims error = %v, want missing -claims error", err)
	}
}

func TestRunCarryAdjudicationsRequiresInputs(t *testing.T) {
	err := runCarryAdjudications(nil)
	if err == nil || !strings.Contains(err.Error(), "-base") {
		t.Fatalf("runCarryAdjudications error = %v, want required inputs error", err)
	}
}

func TestRunAdjudicateClaimsRequiresInputs(t *testing.T) {
	err := runAdjudicateClaims(nil)
	if err == nil || !strings.Contains(err.Error(), "-selections") {
		t.Fatalf("runAdjudicateClaims error = %v, want required inputs error", err)
	}
}
