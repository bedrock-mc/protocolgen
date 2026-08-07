package gophertunneloracle

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"protocolgen/internal/manifest"
)

func TestNormalizationAppliesOnlyDocumentedByteEquivalences(t *testing.T) {
	m := fixtureManifest(
		manifest.Primitive("i16le"),
		manifest.String(manifest.Primitive("var_u32")),
		manifest.Bytes(manifest.Primitive("var_u32")),
		manifest.Array(manifest.Primitive("var_u32"), manifest.Primitive("u8")),
		manifest.Primitive("uuid"),
	)
	source := extraction{Packets: []sourcePacket{{ID: 1, Name: "Fixture", Operations: []sourceOperation{
		{Kind: "primitive", Code: "u16le"},
		{Kind: "bytes", Prefix: "var_u32"},
		{Kind: "string", Prefix: "var_u32"},
		{Kind: "array", Prefix: "var_u32", Element: []sourceOperation{{Kind: "primitive", Code: "u8"}}},
		{Kind: "uuid"},
	}}}}
	lock := fixtureLock()
	report := Compare(m, source, lock, AcceptedFile{SchemaVersion: AcceptedSchemaVersion, MinecraftVersion: "1.26.40", ProtocolVersion: 2168}, "manifest.json")
	if report.Counts.Agreement != 1 || report.Counts.Divergence != 0 {
		t.Fatalf("counts = %#v, want one agreement: %#v", report.Counts, report.Packets)
	}
}

func TestUUIDNormalizesOnlyItsExactSixteenByteShape(t *testing.T) {
	manifestAtoms, manifestReasons := manifestNodeAtoms("UUID", manifest.FixedArray(16, manifest.Primitive("u8")))
	sourceAtoms, sourceReasons := sourceOperationAtoms(sourceOperation{Kind: "uuid", Field: "UUID"})
	if len(manifestReasons) != 0 || len(sourceReasons) != 0 || !atomsEqual(manifestAtoms, sourceAtoms) {
		t.Fatalf("UUID shape did not normalize: manifest=%#v/%#v source=%#v/%#v", manifestAtoms, manifestReasons, sourceAtoms, sourceReasons)
	}
	shortAtoms, _ := sourceOperationAtoms(sourceOperation{Kind: "fixed_array", Length: 15, Element: []sourceOperation{{Kind: "primitive", Code: "u8"}}})
	if atomsEqual(manifestAtoms, shortAtoms) {
		t.Fatal("UUID normalization collapsed a non-16-byte fixed array")
	}
}

func TestNormalizationPreservesWireShapeDistinctions(t *testing.T) {
	for name, pair := range map[string]struct {
		want sourceOperation
		got  sourceOperation
	}{
		"endianness": {
			want: sourceOperation{Kind: "primitive", Code: "u32le"},
			got:  sourceOperation{Kind: "primitive", Code: "u32be"},
		},
		"varint-family": {
			want: sourceOperation{Kind: "primitive", Code: "var_u32"},
			got:  sourceOperation{Kind: "primitive", Code: "zigzag_i32"},
		},
		"option-presence": {
			want: sourceOperation{Kind: "optional", Presence: "bool", Value: []sourceOperation{{Kind: "primitive", Code: "u8"}}},
			got:  sourceOperation{Kind: "optional", Presence: "u8", Value: []sourceOperation{{Kind: "primitive", Code: "u8"}}},
		},
		"array-prefix": {
			want: sourceOperation{Kind: "array", Prefix: "var_u32", Element: []sourceOperation{{Kind: "primitive", Code: "u16le"}}},
			got:  sourceOperation{Kind: "array", Prefix: "u8", Element: []sourceOperation{{Kind: "primitive", Code: "u16le"}}},
		},
		"fixed-array-length": {
			want: sourceOperation{Kind: "fixed_array", Length: 2, Element: []sourceOperation{{Kind: "primitive", Code: "u8"}}},
			got:  sourceOperation{Kind: "fixed_array", Length: 3, Element: []sourceOperation{{Kind: "primitive", Code: "u8"}}},
		},
		"float-versus-integer": {
			want: sourceOperation{Kind: "primitive", Code: "f32le"},
			got:  sourceOperation{Kind: "primitive", Code: "i32le"},
		},
	} {
		wantAtoms, wantReasons := sourceOperationAtoms(pair.want)
		gotAtoms, gotReasons := sourceOperationAtoms(pair.got)
		if len(wantReasons) != 0 || len(gotReasons) != 0 || atomsEqual(wantAtoms, gotAtoms) {
			t.Errorf("%s was collapsed: want=%#v got=%#v", name, wantAtoms, gotAtoms)
		}
	}

	unionWant := sourceOperation{Kind: "union", Control: "var_u32", Variants: []sourceVariant{{Value: 0}, {Value: 1}}}
	unionGot := sourceOperation{Kind: "union", Control: "var_u32", Variants: []sourceVariant{{Value: 0}, {Value: 2}}}
	wantAtoms, _ := sourceOperationAtoms(unionWant)
	gotAtoms, _ := sourceOperationAtoms(unionGot)
	if atomsEqual(wantAtoms, gotAtoms) {
		t.Fatal("union discriminant was collapsed")
	}
}

func TestRuntimeBranchesBecomeUnresolved(t *testing.T) {
	m := fixtureManifest(manifest.Primitive("u8"))
	source := extraction{Packets: []sourcePacket{{ID: 1, Name: "Fixture", Operations: []sourceOperation{
		{Kind: "unresolved", Reason: "runtime conditional branch", Site: "fixture.go:10"},
	}}}}
	report := Compare(m, source, fixtureLock(), AcceptedFile{SchemaVersion: AcceptedSchemaVersion, MinecraftVersion: "1.26.40", ProtocolVersion: 2168}, "manifest.json")
	if report.Counts.Unresolved != 1 || report.Counts.Agreement != 0 {
		t.Fatalf("counts = %#v", report.Counts)
	}
	if !strings.Contains(report.Packets[0].Reasons[0], "runtime conditional") {
		t.Fatalf("reasons = %#v", report.Packets[0].Reasons)
	}
}

func TestExtractRecursesThroughSliceCallbackWithoutTypeChecking(t *testing.T) {
	root := t.TempDir()
	packetDir := filepath.Join(root, "minecraft", "protocol", "packet")
	if err := os.MkdirAll(packetDir, 0o755); err != nil {
		t.Fatal(err)
	}
	idSource := `package packet
const ( IDFixture = iota + 1 )
`
	packetSource := `package packet
import "github.com/sandertv/gophertunnel/minecraft/protocol"
type Fixture struct { Values []uint8; Value uint16 }
func (*Fixture) ID() uint32 { return IDFixture }
func (pk *Fixture) Marshal(io protocol.IO) {
    protocol.FuncSlice(io, &pk.Values, io.Uint8)
    io.Uint16(&pk.Value)
}
`
	if err := os.WriteFile(filepath.Join(packetDir, "id.go"), []byte(idSource), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(packetDir, "fixture.go"), []byte(packetSource), 0o644); err != nil {
		t.Fatal(err)
	}
	result, err := Extract(root)
	if err != nil {
		t.Fatal(err)
	}
	if len(result.Packets) != 1 || result.Packets[0].ID != 1 {
		t.Fatalf("packets = %#v", result.Packets)
	}
	if len(result.Packets[0].Operations) != 2 || result.Packets[0].Operations[0].Kind != "array" || result.Packets[0].Operations[1].Code != "u16le" {
		t.Fatalf("operations = %#v", result.Packets[0].Operations)
	}
}

func TestLoadAcceptedRequiresEvidence(t *testing.T) {
	path := filepath.Join(t.TempDir(), "accepted.json")
	data := `{"schema_version":1,"minecraft_version":"1.26.40","protocol_version":2168,"divergences":[{"id":1,"name":"P","reason":"reviewed","what_would_settle_it":"capture","evidence":[]}]}`
	if err := os.WriteFile(path, []byte(data), 0o644); err != nil {
		t.Fatal(err)
	}
	if _, err := LoadAccepted(path); err == nil {
		t.Fatal("LoadAccepted accepted an entry without evidence")
	}
}

func fixtureManifest(nodes ...manifest.Node) manifest.Manifest {
	fields := make([]manifest.Field, len(nodes))
	for index, node := range nodes {
		fields[index] = manifest.Field{Ordinal: index, Name: "Field", Encode: node, Symmetry: manifest.Symmetric, Provenance: manifest.Provenance{Pins: []string{"fixture"}}}
	}
	return manifest.Manifest{
		SchemaVersion: 2,
		Target:        manifest.Target{MinecraftVersion: "1.26.40", ProtocolVersion: 2168},
		Sources:       []manifest.SourcePin{{ID: "fixture", Kind: "synthetic", Revision: "fixture", Digest: "fixture:oracle", MinecraftVersion: "1.26.40", ProtocolVersion: 2168}},
		Packets:       []manifest.Packet{{ID: 1, Name: "Fixture", Direction: manifest.DirectionUnknown, Fields: fields}},
	}
}

func fixtureLock() Lock {
	var lock Lock
	lock.SchemaVersion = LockSchemaVersion
	lock.MinecraftVersion = "1.26.40"
	lock.ProtocolVersion = 2168
	lock.Gophertunnel.Repo = "fixture"
	lock.Gophertunnel.Commit = "0123456789012345678901234567890123456789"
	return lock
}
