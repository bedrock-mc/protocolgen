package main

import (
	"go/format"
	"strings"
	"testing"
)

func TestEmitPacketProducesValidGo(t *testing.T) {
	pk, _ := parseGenPacket([]byte(demoDocJSON))
	src, todos, err := emitPacket(pk, "IDDemo")
	if err != nil {
		t.Fatalf("emitPacket: %v", err)
	}
	// Output must be syntactically valid Go.
	if _, err := format.Source([]byte(src)); err != nil {
		t.Fatalf("emitted source is not valid Go: %v\n%s", err, src)
	}
	for _, want := range []string{
		"package packet",
		"type Demo struct {",
		"func (*Demo) ID() uint32 {",
		"return IDDemo",
		"func (pk *Demo) Marshal(io protocol.IO) {",
		"io.Uint8(&pk.Action)",
		"io.Varuint64(&pk.Target)",
		"io.Vec3(&pk.Pos)",
		"github.com/sandertv/gophertunnel/minecraft/protocol",
		"github.com/go-gl/mathgl/mgl32",
	} {
		if !strings.Contains(src, want) {
			t.Errorf("emitted source missing %q\n---\n%s", want, src)
		}
	}
	// Body (oneOf) and Attrs (array of unknown composite) require hand-finishing.
	if len(todos) == 0 {
		t.Errorf("expected TODOs for oneOf/composite fields, got none")
	}
}

func TestEmitPacketMarshalOrderMatchesOrdinal(t *testing.T) {
	pk, _ := parseGenPacket([]byte(demoDocJSON))
	src, _, _ := emitPacket(pk, "IDDemo")
	// Action (ord 0) must marshal before Target (ord 1) before Pos (ord 2).
	ai := strings.Index(src, "&pk.Action")
	ti := strings.Index(src, "&pk.Target")
	pi := strings.Index(src, "&pk.Pos")
	if !(ai < ti && ti < pi) {
		t.Errorf("marshal order wrong: Action@%d Target@%d Pos@%d", ai, ti, pi)
	}
}
