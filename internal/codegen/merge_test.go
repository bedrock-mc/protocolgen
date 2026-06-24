package codegen

import (
	"go/format"
	"strings"
	"testing"
)

const existingAnimate = `package packet

import "github.com/sandertv/gophertunnel/minecraft/protocol"

const (
	AnimateActionSwingArm = iota + 1
	AnimateActionStopSleep
)

// Animate is sent by the server to send a player animation from one player to all
// viewers of that player.
type Animate struct {
	// ActionType is the ID of the animation action to execute.
	ActionType uint8
	// EntityRuntimeID is the runtime ID of the player.
	EntityRuntimeID uint64
	// Data ...
	Data float32
}

// ID ...
func (*Animate) ID() uint32 {
	return IDAnimate
}

func (pk *Animate) Marshal(io protocol.IO) {
	io.Uint8(&pk.ActionType)
	io.Varuint64(&pk.EntityRuntimeID)
	io.Float32(&pk.Data)
}

func swingHelper() string { return "x" }
`

func demoAnimatePacket() genPacket {
	return genPacket{TypeName: "Animate", ID: 44, Fields: []genField{
		{GoName: "Action", Under: "uint8", JSONType: "integer", Required: true},
		{GoName: "EntityRuntimeID", Under: "uint64", Options: []string{"Compression"}, JSONType: "integer", Required: true},
		{GoName: "Data", Under: "float", JSONType: "number", Required: true},
	}}
}

func TestMergePreservesCommentsConstsHelpers(t *testing.T) {
	merged, _, err := mergeIntoExisting([]byte(existingAnimate), demoAnimatePacket(), false)
	if err != nil {
		t.Fatalf("mergeIntoExisting: %v", err)
	}
	if _, err := format.Source([]byte(merged)); err != nil {
		t.Fatalf("merged source not valid Go: %v\n%s", err, merged)
	}
	for _, want := range []string{
		"// Animate is sent by the server",         // type doc comment preserved
		"AnimateActionSwingArm = iota + 1",         // const block preserved
		"func swingHelper() string",                // helper preserved
		"// ActionType is the ID of the animation", // field comment preserved (by position)
	} {
		if !strings.Contains(merged, want) {
			t.Errorf("merged output missing preserved %q\n---\n%s", want, merged)
		}
	}
}

func TestMergeReplacesFieldsAndMarshal(t *testing.T) {
	merged, _, err := mergeIntoExisting([]byte(existingAnimate), demoAnimatePacket(), false)
	if err != nil {
		t.Fatalf("mergeIntoExisting: %v", err)
	}
	// Field 0 renamed ActionType -> Action; Marshal must reference the new name.
	if !strings.Contains(merged, "Action uint8") {
		t.Errorf("merged output missing regenerated field 'Action uint8'\n%s", merged)
	}
	if strings.Contains(merged, "io.Uint8(&pk.ActionType)") {
		t.Errorf("old Marshal op still present; body was not replaced\n%s", merged)
	}
	if !strings.Contains(merged, "io.Uint8(&pk.Action)") {
		t.Errorf("merged output missing regenerated Marshal op\n%s", merged)
	}
}

func TestMergePreservesWireMatchedNames(t *testing.T) {
	// The docs rename both fields but the wire is unchanged; with preserveMatching
	// the gopher names (and ops) must be kept verbatim — no rename churn.
	const existing = `package packet

import "github.com/sandertv/gophertunnel/minecraft/protocol"

type Foo struct {
	// EntityRuntimeID is the runtime id.
	EntityRuntimeID uint64
	Health          int32
}

func (*Foo) ID() uint32 { return IDFoo }

func (pk *Foo) Marshal(io protocol.IO) {
	io.Varuint64(&pk.EntityRuntimeID)
	io.Varint32(&pk.Health)
}`
	pk := genPacket{TypeName: "Foo", ID: 1, Fields: []genField{
		{GoName: "TargetActorRuntimeID", Under: "uint64", Options: []string{"Compression"}, Required: true},
		{GoName: "HealthValue", Under: "int32", Options: []string{"Compression"}, Required: true},
	}}
	merged, _, err := mergeIntoExisting([]byte(existing), pk, true)
	if err != nil {
		t.Fatalf("mergeIntoExisting: %v", err)
	}
	if !strings.Contains(merged, "EntityRuntimeID uint64") || !strings.Contains(merged, "io.Varuint64(&pk.EntityRuntimeID)") {
		t.Errorf("gopher field EntityRuntimeID should be preserved\n%s", merged)
	}
	if strings.Contains(merged, "TargetActorRuntimeID") {
		t.Errorf("doc name should NOT appear for a wire-matched field\n%s", merged)
	}
}

func TestMergeWireChangedFieldUsesGenerated(t *testing.T) {
	// The wire encoding genuinely changed (uint16 -> varint32), so the generated
	// field must win and the change must be visible.
	const existing = `package packet

import "github.com/sandertv/gophertunnel/minecraft/protocol"

type Foo struct {
	Count uint16
}

func (*Foo) ID() uint32 { return IDFoo }

func (pk *Foo) Marshal(io protocol.IO) {
	io.Uint16(&pk.Count)
}`
	pk := genPacket{TypeName: "Foo", ID: 1, Fields: []genField{
		{GoName: "Count", Under: "int32", Options: []string{"Compression"}, Required: true},
	}}
	merged, _, err := mergeIntoExisting([]byte(existing), pk, true)
	if err != nil {
		t.Fatalf("mergeIntoExisting: %v", err)
	}
	if !strings.Contains(merged, "io.Varint32(&pk.Count)") {
		t.Errorf("a real wire change must use the generated op\n%s", merged)
	}
}

func TestMergePreservesByteSliceOverString(t *testing.T) {
	// The docs type raw byte blobs as "string"; gophertunnel uses []byte. Since
	// io.String and io.ByteSlice are wire-identical, preserve the existing []byte.
	const existing = `package packet

import "github.com/sandertv/gophertunnel/minecraft/protocol"

type Foo struct {
	// Payload is the raw data.
	Payload []byte
}

func (*Foo) ID() uint32 { return IDFoo }

func (pk *Foo) Marshal(io protocol.IO) {
	io.ByteSlice(&pk.Payload)
}`
	pk := genPacket{TypeName: "Foo", ID: 1, Fields: []genField{
		{GoName: "Payload", JSONType: "string", Required: true}, // doc models the blob as string
	}}
	merged, _, err := mergeIntoExisting([]byte(existing), pk, false)
	if err != nil {
		t.Fatalf("mergeIntoExisting: %v", err)
	}
	if !strings.Contains(merged, "Payload []byte") {
		t.Errorf("merged should keep []byte, not downgrade to string\n%s", merged)
	}
	if !strings.Contains(merged, "io.ByteSlice(&pk.Payload)") {
		t.Errorf("merged should keep io.ByteSlice\n%s", merged)
	}
	if strings.Contains(merged, "io.String") {
		t.Errorf("merged should not use io.String for a []byte field\n%s", merged)
	}
}

func TestMergeAddsNeededImport(t *testing.T) {
	// A packet whose generated fields need mgl32 must gain that import.
	pk := genPacket{TypeName: "Animate", ID: 44, Fields: []genField{
		{GoName: "Pos", Composite: true, RefTitle: "Vec3"},
	}}
	merged, _, err := mergeIntoExisting([]byte(existingAnimate), pk, false)
	if err != nil {
		t.Fatalf("mergeIntoExisting: %v", err)
	}
	if !strings.Contains(merged, "github.com/go-gl/mathgl/mgl32") {
		t.Errorf("merged output missing added mgl32 import\n%s", merged)
	}
}
