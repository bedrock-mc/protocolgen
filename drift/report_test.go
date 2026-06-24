package main

import (
	"strings"
	"testing"
)

func TestRenderShowsMismatchDetail(t *testing.T) {
	doc := animateDoc()
	src := animateSource()
	src.Ops[1].Method = "Uint64" // mismatch on the Target field
	rep := Compare([]DocPacket{doc}, []SourcePacket{src})

	out := render(rep, false)
	if !strings.Contains(out, "AnimatePacket") {
		t.Errorf("output missing packet name:\n%s", out)
	}
	if !strings.Contains(out, "Target") {
		t.Errorf("output missing mismatched field name:\n%s", out)
	}
	if !strings.Contains(out, "mismatch") {
		t.Errorf("output missing mismatch marker:\n%s", out)
	}
	if !strings.Contains(out, "Uint64") {
		t.Errorf("output missing the offending source op:\n%s", out)
	}
}

func TestRenderListsNewAndRemovedPackets(t *testing.T) {
	rep := Compare(
		[]DocPacket{{Name: "NewPacket", ID: 200}},
		[]SourcePacket{{TypeName: "Legacy", IDConst: "IDLegacy", ID: 201}},
	)
	out := render(rep, false)
	if !strings.Contains(out, "NewPacket") {
		t.Errorf("output missing new packet:\n%s", out)
	}
	if !strings.Contains(out, "Legacy") {
		t.Errorf("output missing removed packet:\n%s", out)
	}
}

func TestRenderCleanReportSaysNoDrift(t *testing.T) {
	doc := DocPacket{Name: "P", ID: 5, Fields: []DocField{
		{Name: "A", Ordinal: 0, UnderlyingType: "uint8", JSONType: "integer"},
	}}
	src := SourcePacket{TypeName: "P", ID: 5, Ops: []WireOp{{Kind: "io", Method: "Uint8", Field: "A"}}}
	out := render(Compare([]DocPacket{doc}, []SourcePacket{src}), false)
	if !strings.Contains(strings.ToLower(out), "no drift") {
		t.Errorf("clean report should say no drift:\n%s", out)
	}
}
