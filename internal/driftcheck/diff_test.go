package driftcheck

import "testing"

func animateDoc() DocPacket {
	return DocPacket{Name: "AnimatePacket", ID: 44, Fields: []DocField{
		{Name: "Action", Ordinal: 0, UnderlyingType: "uint8", Enum: true, Options: []string{"Enum-as-Value"}, JSONType: "string"},
		{Name: "Target", Ordinal: 1, UnderlyingType: "uint64", Options: []string{"Compression"}, RefTitle: "RuntimeID"},
		{Name: "Data", Ordinal: 2, UnderlyingType: "float", JSONType: "number"},
		{Name: "Swing Source", Ordinal: 3, UnderlyingType: "uint8", Enum: true, JSONType: "string"},
	}}
}

func animateSource() SourcePacket {
	return SourcePacket{TypeName: "Animate", IDConst: "IDAnimate", ID: 44, Ops: []WireOp{
		{Kind: "io", Method: "Uint8", Field: "ActionType"},
		{Kind: "io", Method: "Varuint64", Field: "EntityRuntimeID"},
		{Kind: "io", Method: "Float32", Field: "Data"},
		{Kind: "protocol", Method: "OptionalFunc", Inner: "String"},
		{Kind: "custom", Method: "swingSourceFromString", Field: "SwingSource", Guarded: true},
	}}
}

func TestCompareMatchesByIDAndAlignsFields(t *testing.T) {
	rep := Compare([]DocPacket{animateDoc()}, []SourcePacket{animateSource()})
	if len(rep.MissingInSource) != 0 {
		t.Errorf("MissingInSource = %d, want 0", len(rep.MissingInSource))
	}
	if len(rep.MissingInDocs) != 0 {
		t.Errorf("MissingInDocs = %d, want 0", len(rep.MissingInDocs))
	}
	if len(rep.Diffs) != 1 {
		t.Fatalf("Diffs = %d, want 1", len(rep.Diffs))
	}
	d := rep.Diffs[0]
	// Positions 0-2 encode cleanly; the Swing Source field maps to a protocol
	// helper (review); the trailing guarded custom op is a conditional read-back
	// (review, not drift).
	wantStatuses := []string{statusOK, statusOK, statusOK, statusReview, statusReview}
	if len(d.Fields) != len(wantStatuses) {
		t.Fatalf("len(Fields) = %d, want %d; got %+v", len(d.Fields), len(wantStatuses), d.Fields)
	}
	for i, want := range wantStatuses {
		if d.Fields[i].Status != want {
			t.Errorf("Fields[%d].Status = %q, want %q", i, d.Fields[i].Status, want)
		}
	}
}

func TestCompareDetectsEncodingMismatch(t *testing.T) {
	doc := animateDoc()
	src := animateSource()
	src.Ops[1].Method = "Uint64" // fixed instead of varint for a Compression field
	rep := Compare([]DocPacket{doc}, []SourcePacket{src})
	if rep.Diffs[0].Fields[1].Status != statusMismatch {
		t.Errorf("Fields[1].Status = %q, want %q", rep.Diffs[0].Fields[1].Status, statusMismatch)
	}
	if !rep.HasDrift() {
		t.Errorf("HasDrift() = false, want true on an encoding mismatch")
	}
}

func TestCompareReportsMissingAndExtraPackets(t *testing.T) {
	docOnly := DocPacket{Name: "NewPacket", ID: 200}
	srcOnly := SourcePacket{TypeName: "Legacy", IDConst: "IDLegacy", ID: 201}
	rep := Compare([]DocPacket{animateDoc(), docOnly}, []SourcePacket{animateSource(), srcOnly})
	if len(rep.MissingInSource) != 1 || rep.MissingInSource[0].ID != 200 {
		t.Errorf("MissingInSource = %+v, want [200]", rep.MissingInSource)
	}
	if len(rep.MissingInDocs) != 1 || rep.MissingInDocs[0].ID != 201 {
		t.Errorf("MissingInDocs = %+v, want [201]", rep.MissingInDocs)
	}
	if !rep.HasDrift() {
		t.Errorf("HasDrift() = false, want true when packets are added/removed")
	}
}

func TestCompareUnguardedExtraOpIsDrift(t *testing.T) {
	doc := DocPacket{Name: "P", ID: 7, Fields: []DocField{
		{Name: "A", Ordinal: 0, UnderlyingType: "uint8", JSONType: "integer"},
	}}
	src := SourcePacket{TypeName: "P", ID: 7, Ops: []WireOp{
		{Kind: "io", Method: "Uint8", Field: "A"},
		{Kind: "io", Method: "Float32", Field: "Extra"}, // unguarded extra op
	}}
	rep := Compare([]DocPacket{doc}, []SourcePacket{src})
	if rep.Diffs[0].Fields[1].Status != statusExtraSrc {
		t.Errorf("Fields[1].Status = %q, want %q", rep.Diffs[0].Fields[1].Status, statusExtraSrc)
	}
	if !rep.HasDrift() {
		t.Errorf("HasDrift() = false, want true for an unguarded extra op")
	}
}

func TestCompareGuardedPacketDowngradesAlignmentToReview(t *testing.T) {
	// A packet with conditional (guarded) ops has unreliable positional
	// alignment, so trailing overflow ops should be review, not asserted drift.
	doc := DocPacket{Name: "P", ID: 8, Fields: []DocField{
		{Name: "A", Ordinal: 0, UnderlyingType: "uint8", JSONType: "integer"},
	}}
	src := SourcePacket{TypeName: "P", ID: 8, Ops: []WireOp{
		{Kind: "io", Method: "Uint8", Field: "A"},
		{Kind: "io", Method: "String", Field: "B", Guarded: true}, // conditional field
		{Kind: "io", Method: "String", Field: "C"},                // overflow after a conditional
	}}
	rep := Compare([]DocPacket{doc}, []SourcePacket{src})
	if got := rep.Diffs[0].Fields[2].Status; got != statusReview {
		t.Errorf("Fields[2].Status = %q, want %q (overflow in a guarded packet)", got, statusReview)
	}
	if rep.HasDrift() {
		t.Errorf("HasDrift() = true; alignment overflow in a guarded packet should not be drift")
	}
}

func TestCompareGuardedPacketStillFlagsEncodingMismatch(t *testing.T) {
	// Even in a guarded packet, an outright encoding disagreement is real drift.
	doc := DocPacket{Name: "P", ID: 9, Fields: []DocField{
		{Name: "A", Ordinal: 0, UnderlyingType: "uint32", Options: []string{"Compression"}, JSONType: "integer"},
		{Name: "B", Ordinal: 1, UnderlyingType: "uint8", JSONType: "integer"},
	}}
	src := SourcePacket{TypeName: "P", ID: 9, Ops: []WireOp{
		{Kind: "io", Method: "Uint32", Field: "A"}, // mismatch: fixed vs varint
		{Kind: "io", Method: "Uint8", Field: "B", Guarded: true},
	}}
	rep := Compare([]DocPacket{doc}, []SourcePacket{src})
	if got := rep.Diffs[0].Fields[0].Status; got != statusMismatch {
		t.Errorf("Fields[0].Status = %q, want %q", got, statusMismatch)
	}
	if !rep.HasDrift() {
		t.Errorf("HasDrift() = false, want true for an encoding mismatch even in a guarded packet")
	}
}

func TestCompareExpandsCompositeIntoScalarOps(t *testing.T) {
	// The docs group rotation as a Vec2 (one field = two floats on the wire);
	// gophertunnel writes two separate Float32 ops. This must align with no
	// cascade of false mismatches on the following fields.
	doc := DocPacket{Name: "P", ID: 20, Fields: []DocField{
		{Name: "Rotation", Ordinal: 0, Composite: true, RefTitle: "Vec2"},
		{Name: "HeadYaw", Ordinal: 1, UnderlyingType: "float", JSONType: "number"},
		{Name: "Flag", Ordinal: 2, UnderlyingType: "uint8", JSONType: "integer"},
	}}
	src := SourcePacket{TypeName: "P", ID: 20, Ops: []WireOp{
		{Kind: "io", Method: "Float32", Field: "Yaw"},
		{Kind: "io", Method: "Float32", Field: "Pitch"},
		{Kind: "io", Method: "Float32", Field: "HeadYaw"},
		{Kind: "io", Method: "Uint8", Field: "Flag"},
	}}
	rep := Compare([]DocPacket{doc}, []SourcePacket{src})
	if rep.HasDrift() {
		t.Errorf("HasDrift() = true; a Vec2 should expand to two floats with no cascade; fields=%+v", rep.Diffs[0].Fields)
	}
}

func TestCompareCompositeTypedOpStillMatches(t *testing.T) {
	// When gophertunnel uses the dedicated composite op (io.Vec3), it must still
	// match the doc's composite field directly.
	doc := DocPacket{Name: "P", ID: 21, Fields: []DocField{
		{Name: "Pos", Ordinal: 0, Composite: true, RefTitle: "Vec3"},
		{Name: "Flag", Ordinal: 1, UnderlyingType: "uint8", JSONType: "integer"},
	}}
	src := SourcePacket{TypeName: "P", ID: 21, Ops: []WireOp{
		{Kind: "io", Method: "Vec3", Field: "Pos"},
		{Kind: "io", Method: "Uint8", Field: "Flag"},
	}}
	rep := Compare([]DocPacket{doc}, []SourcePacket{src})
	if rep.HasDrift() {
		t.Errorf("HasDrift() = true; io.Vec3 should match composite Vec3; fields=%+v", rep.Diffs[0].Fields)
	}
}

func TestCompareRealMismatchAfterCompositeStillDetected(t *testing.T) {
	// A genuine encoding mismatch following a composite must still be caught
	// (expansion must not mask real drift).
	doc := DocPacket{Name: "P", ID: 22, Fields: []DocField{
		{Name: "Rotation", Ordinal: 0, Composite: true, RefTitle: "Vec2"},
		{Name: "Count", Ordinal: 1, UnderlyingType: "uint32", Options: []string{"Compression"}, JSONType: "integer"},
	}}
	src := SourcePacket{TypeName: "P", ID: 22, Ops: []WireOp{
		{Kind: "io", Method: "Float32", Field: "Yaw"},
		{Kind: "io", Method: "Float32", Field: "Pitch"},
		{Kind: "io", Method: "Uint32", Field: "Count"}, // mismatch: fixed vs varint
	}}
	rep := Compare([]DocPacket{doc}, []SourcePacket{src})
	if !rep.HasDrift() {
		t.Errorf("HasDrift() = false; a real mismatch after a composite must be caught")
	}
}

func TestCompareCleanPacketHasNoDrift(t *testing.T) {
	// A packet whose ops all match and align 1:1 should report no drift.
	doc := DocPacket{Name: "P", ID: 5, Fields: []DocField{
		{Name: "A", Ordinal: 0, UnderlyingType: "uint8", JSONType: "integer"},
		{Name: "B", Ordinal: 1, UnderlyingType: "uint32", Options: []string{"Compression"}, JSONType: "integer"},
	}}
	src := SourcePacket{TypeName: "P", ID: 5, Ops: []WireOp{
		{Kind: "io", Method: "Uint8", Field: "A"},
		{Kind: "io", Method: "Varuint32", Field: "B"},
	}}
	rep := Compare([]DocPacket{doc}, []SourcePacket{src})
	if rep.HasDrift() {
		t.Errorf("HasDrift() = true, want false for a fully matching packet; diffs=%+v", rep.Diffs[0].Fields)
	}
}
