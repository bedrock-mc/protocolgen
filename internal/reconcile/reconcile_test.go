package reconcile

import (
	"strings"
	"testing"

	"protocolgen/internal/claims"
	"protocolgen/internal/manifest"
)

func TestReconcileRejectsUnfingerprintedSourceDisagreement(t *testing.T) {
	target := manifest.Target{MinecraftVersion: "fixture", ProtocolVersion: 2168}
	results := []claims.Result{
		{Pin: sourcePin("endstone"), Target: target, Claims: []claims.Claim{testClaim("endstone", manifest.Primitive("u8"))}},
		{Pin: sourcePin("mojang"), Target: target, Claims: []claims.Claim{testClaim("mojang", manifest.Primitive("u16le"))}},
	}
	_, err := Reconcile(target, results, nil)
	if err == nil || !strings.Contains(err.Error(), "adjudication") {
		t.Fatalf("Reconcile error = %v, want missing adjudication", err)
	}
}

func TestReconcileRejectsStaleAdjudicationFingerprint(t *testing.T) {
	target := manifest.Target{MinecraftVersion: "fixture", ProtocolVersion: 2168}
	left := testClaim("endstone", manifest.Primitive("u8"))
	right := testClaim("mojang", manifest.Primitive("u16le"))
	context, err := claims.ContextFingerprint(target, []claims.Claim{left, right})
	if err != nil {
		t.Fatalf("ContextFingerprint: %v", err)
	}
	left.Name = "MutatedAfterReview"
	adj := manifest.Adjudication{
		ID: "choose-endstone", Target: left.FieldPath, PrePatchContextSHA256: context,
		Claims:         []manifest.ClaimFingerprint{{SourceID: "endstone", Digest: mustFingerprint(t, testClaim("endstone", manifest.Primitive("u8")))}, {SourceID: "mojang", Digest: mustFingerprint(t, right)}},
		SelectedSource: "endstone", Evidence: []manifest.Evidence{{SourceID: "endstone", Locator: "fixture/evidence"}}, Reason: "fixture evidence",
	}
	_, err = Reconcile(target, []claims.Result{{Pin: sourcePin("endstone"), Target: target, Claims: []claims.Claim{left}}, {Pin: sourcePin("mojang"), Target: target, Claims: []claims.Claim{right}}}, []manifest.Adjudication{adj})
	if err == nil || !strings.Contains(err.Error(), "stale") {
		t.Fatalf("Reconcile error = %v, want stale fingerprint failure", err)
	}
}

func TestReconcileMergesIdenticalClaimsAndOnlyPinsSources(t *testing.T) {
	target := manifest.Target{MinecraftVersion: "fixture", ProtocolVersion: 2168}
	claimA := testClaim("endstone", manifest.Primitive("u8"))
	claimB := testClaim("mojang", manifest.Primitive("u8"))
	m, err := Reconcile(target, []claims.Result{{Pin: sourcePin("endstone"), Target: target, Claims: []claims.Claim{claimA}}, {Pin: sourcePin("mojang"), Target: target, Claims: []claims.Claim{claimB}}}, nil)
	if err != nil {
		t.Fatalf("Reconcile: %v", err)
	}
	field := m.Packets[0].Fields[0]
	if len(field.Provenance.Pins) != 2 || len(field.Provenance.Evidence) != 0 {
		t.Fatalf("provenance = %+v, want two pins and no detailed evidence", field.Provenance)
	}
}

func testClaim(sourceID string, node manifest.Node) claims.Claim {
	return claims.Claim{SourceID: sourceID, Locator: "fixture/Vocabulary.Value", PacketID: 1, PacketName: "Vocabulary", Direction: manifest.DirectionClientbound, FieldPath: "Vocabulary.Value", Ordinal: 0, Name: "Value", Encode: node, Symmetry: manifest.Symmetric}
}

func sourcePin(id string) manifest.SourcePin {
	return manifest.SourcePin{ID: id, Kind: id, Revision: "fixture-2168", Digest: "sha256:" + id, MinecraftVersion: "fixture", ProtocolVersion: 2168}
}

func mustFingerprint(t *testing.T, claim claims.Claim) string {
	t.Helper()
	digest, err := claims.Fingerprint(claim)
	if err != nil {
		t.Fatalf("Fingerprint: %v", err)
	}
	return digest
}
