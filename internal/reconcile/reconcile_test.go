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

func TestReconcileUsesConcreteClaimWhenAnotherSourceIsUnresolved(t *testing.T) {
	target := manifest.Target{MinecraftVersion: "fixture", ProtocolVersion: 2168}
	unresolved := testClaim("mojang", manifest.Unresolved("source omits the wire shape", true))
	concrete := testClaim("endstone", manifest.Primitive("u8"))

	m, err := Reconcile(target, []claims.Result{
		{Pin: sourcePin("mojang"), Target: target, Claims: []claims.Claim{unresolved}},
		{Pin: sourcePin("endstone"), Target: target, Claims: []claims.Claim{concrete}},
	}, nil)
	if err != nil {
		t.Fatalf("Reconcile: %v", err)
	}
	field := m.Packets[0].Fields[0]
	if field.Encode.Kind != manifest.KindPrimitive || len(field.Provenance.Pins) != 1 || field.Provenance.Pins[0] != "endstone" {
		t.Fatalf("field = %#v, want concrete Endstone-only claim", field)
	}
}

func TestReconcileUsesCompleteClaimWhenAnotherSourceHasNestedUnresolvedShape(t *testing.T) {
	target := manifest.Target{MinecraftVersion: "fixture", ProtocolVersion: 2168}
	incomplete := manifest.Node{Kind: manifest.KindStruct, Fields: []manifest.Field{{Ordinal: 0, Name: "Mode", Encode: manifest.Unresolved("missing enum values", true), Symmetry: manifest.Symmetric, Provenance: manifest.Provenance{Pins: []string{"mojang"}}}}}
	complete := manifest.Node{Kind: manifest.KindStruct, Fields: []manifest.Field{{Ordinal: 0, Name: "Mode", Encode: manifest.Primitive("u8"), Symmetry: manifest.Symmetric, Provenance: manifest.Provenance{Pins: []string{"endstone"}}}}}

	m, err := Reconcile(target, []claims.Result{
		{Pin: sourcePin("mojang"), Target: target, Claims: []claims.Claim{testClaim("mojang", incomplete)}},
		{Pin: sourcePin("endstone"), Target: target, Claims: []claims.Claim{testClaim("endstone", complete)}},
	}, nil)
	if err != nil {
		t.Fatalf("Reconcile: %v", err)
	}
	field := m.Packets[0].Fields[0]
	if field.Encode.Fields[0].Encode.Kind != manifest.KindPrimitive || len(field.Provenance.Pins) != 1 || field.Provenance.Pins[0] != "endstone" {
		t.Fatalf("field = %#v, want complete Endstone-only claim", field)
	}
}

func TestReconcileCombinesComplementaryNestedClaims(t *testing.T) {
	target := manifest.Target{MinecraftVersion: "fixture", ProtocolVersion: 2168}
	left := manifest.Node{Kind: manifest.KindStruct, Fields: []manifest.Field{
		{Ordinal: 0, Name: "First", Encode: manifest.Primitive("u8"), Symmetry: manifest.Symmetric, Provenance: manifest.Provenance{Pins: []string{"endstone"}}},
		{Ordinal: 1, Name: "Second", Encode: manifest.Unresolved("missing", true), Symmetry: manifest.Symmetric, Provenance: manifest.Provenance{Pins: []string{"endstone"}}},
	}}
	right := manifest.Node{Kind: manifest.KindStruct, Fields: []manifest.Field{
		{Ordinal: 0, Name: "First", Encode: manifest.Unresolved("missing", true), Symmetry: manifest.Symmetric, Provenance: manifest.Provenance{Pins: []string{"mojang"}}},
		{Ordinal: 1, Name: "Second", Encode: manifest.Primitive("u16le"), Symmetry: manifest.Symmetric, Provenance: manifest.Provenance{Pins: []string{"mojang"}}},
	}}

	m, err := Reconcile(target, []claims.Result{
		{Pin: sourcePin("endstone"), Target: target, Claims: []claims.Claim{testClaim("endstone", left)}},
		{Pin: sourcePin("mojang"), Target: target, Claims: []claims.Claim{testClaim("mojang", right)}},
	}, nil)
	if err != nil {
		t.Fatalf("Reconcile: %v", err)
	}
	fields := m.Packets[0].Fields[0].Encode.Fields
	if fields[0].Encode.Primitive.Code != "u8" || fields[1].Encode.Primitive.Code != "u16le" {
		t.Fatalf("merged fields = %#v", fields)
	}
}

func TestReconcilePreservesSourcePinWhenAllClaimsAreUnresolved(t *testing.T) {
	target := manifest.Target{MinecraftVersion: "fixture", ProtocolVersion: 2168}
	_, err := Reconcile(target, []claims.Result{{
		Pin: sourcePin("endstone"), Target: target,
		Claims: []claims.Claim{testClaim("endstone", manifest.Unresolved("missing selector", true))},
	}}, nil)
	if err == nil || !strings.Contains(err.Error(), "unresolved node") {
		t.Fatalf("Reconcile error = %v, want unresolved wire-shape failure", err)
	}
}

func TestReconcileMergesByteEquivalentEnumAndPrimitiveClaims(t *testing.T) {
	target := manifest.Target{MinecraftVersion: "fixture", ProtocolVersion: 2168}
	primitiveClaim := testClaim("endstone", manifest.Primitive("u8"))
	enumClaim := testClaim("mojang", manifest.Node{
		Kind: manifest.KindEnum, Semantic: "Mode", TypeID: "Mode.json#", Primitive: manifest.Primitive("u8").Primitive,
		Variants: []manifest.Variant{{Value: 0, Name: "Ready", Encode: manifest.Void()}, {Value: 1, Name: "Done", Encode: manifest.Void()}},
	})

	m, err := Reconcile(target, []claims.Result{
		{Pin: sourcePin("endstone"), Target: target, Claims: []claims.Claim{primitiveClaim}},
		{Pin: sourcePin("mojang"), Target: target, Claims: []claims.Claim{enumClaim}},
	}, nil)
	if err != nil {
		t.Fatalf("Reconcile: %v", err)
	}
	field := m.Packets[0].Fields[0]
	if field.Encode.Kind != manifest.KindEnum || len(field.Provenance.Pins) != 2 {
		t.Fatalf("field = %#v, want semantic enum with both source pins", field)
	}
}

func TestReconcileIgnoresNestedSourceNamesAndProvenanceForWireComparison(t *testing.T) {
	target := manifest.Target{MinecraftVersion: "fixture", ProtocolVersion: 2168}
	leftNode := manifest.Node{Kind: manifest.KindStruct, Semantic: "Left", TypeID: "Left", Fields: []manifest.Field{{Ordinal: 0, Name: "First", Encode: manifest.Primitive("u16le"), Symmetry: manifest.Symmetric, Provenance: manifest.Provenance{Pins: []string{"endstone"}}}}}
	rightNode := manifest.Node{Kind: manifest.KindStruct, Semantic: "Right", TypeID: "Right.json#", Fields: []manifest.Field{{Ordinal: 0, Name: "Other", Encode: manifest.Primitive("u16le"), Symmetry: manifest.Symmetric, Provenance: manifest.Provenance{Pins: []string{"mojang"}}}}}

	m, err := Reconcile(target, []claims.Result{
		{Pin: sourcePin("endstone"), Target: target, Claims: []claims.Claim{testClaim("endstone", leftNode)}},
		{Pin: sourcePin("mojang"), Target: target, Claims: []claims.Claim{testClaim("mojang", rightNode)}},
	}, nil)
	if err != nil {
		t.Fatalf("Reconcile: %v", err)
	}
	if got := m.Packets[0].Fields[0].Provenance.Pins; len(got) != 2 {
		t.Fatalf("pins = %v, want both sources", got)
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
