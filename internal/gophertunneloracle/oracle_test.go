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

func TestFixedArrayGroupingNormalizesOnlyWireEquivalentScalarLayout(t *testing.T) {
	nested, reasons := manifestNodeAtoms("Nested", manifest.FixedArray(16, manifest.FixedArray(16, manifest.Primitive("i8"))))
	flat, flatReasons := sourceOperationAtoms(sourceOperation{Kind: "fixed_array", Length: 256, Element: []sourceOperation{{Kind: "primitive", Code: "i8"}}})
	if len(reasons) != 0 || len(flatReasons) != 0 || !atomsEqual(normalizeFixedArrayGrouping(nested), normalizeFixedArrayGrouping(flat)) {
		t.Fatalf("nested and flat fixed arrays did not normalize: nested=%#v flat=%#v", nested, flat)
	}
	short, _ := sourceOperationAtoms(sourceOperation{Kind: "fixed_array", Length: 255, Element: []sourceOperation{{Kind: "primitive", Code: "i8"}}})
	if atomsEqual(normalizeFixedArrayGrouping(nested), normalizeFixedArrayGrouping(short)) {
		t.Fatal("different scalar counts were normalized as equivalent")
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

func TestExtractBuildsFinitePathsForSwitchAndConditional(t *testing.T) {
	root := t.TempDir()
	packetDir := filepath.Join(root, "minecraft", "protocol", "packet")
	if err := os.MkdirAll(packetDir, 0o755); err != nil {
		t.Fatal(err)
	}
	idSource := `package packet
const (
	IDFixture = iota + 1
	ModeA = 1
	ModeB = 2
)
`
	packetSource := `package packet
import "github.com/sandertv/gophertunnel/minecraft/protocol"
type Fixture struct { Mode uint8; Flag bool; Value uint16 }
func (*Fixture) ID() uint32 { return IDFixture }
func (pk *Fixture) Marshal(io protocol.IO) {
	io.Uint8(&pk.Mode)
	switch pk.Mode {
	case ModeA:
		io.Uint16(&pk.Value)
	case ModeB:
		io.Bool(&pk.Flag)
	}
	if pk.Flag {
		io.Uint16(&pk.Value)
	}
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
	if len(result.Packets) != 1 {
		t.Fatalf("packets = %#v", result.Packets)
	}
	packet := result.Packets[0]
	if len(packet.Paths) != 4 {
		t.Fatalf("paths = %#v, want switch x if path expansion", packet.Paths)
	}
	for _, path := range packet.Paths {
		if len(path.Constraints) == 0 || len(path.Operations) == 0 {
			t.Fatalf("path lost control-flow metadata: %#v", path)
		}
	}
}

func TestExtractFollowsStaticallyResolvableLocalWireCall(t *testing.T) {
	root := t.TempDir()
	packetDir := filepath.Join(root, "minecraft", "protocol", "packet")
	if err := os.MkdirAll(packetDir, 0o755); err != nil {
		t.Fatal(err)
	}
	idSource := `package packet
const IDFixture = 1
`
	packetSource := `package packet
import "github.com/sandertv/gophertunnel/minecraft/protocol"
type Fixture struct { Value uint16 }
func (*Fixture) ID() uint32 { return IDFixture }
func encodeValue(io protocol.IO, value *uint16) { io.Uint16(value) }
func (pk *Fixture) Marshal(io protocol.IO) { encodeValue(io, &pk.Value) }
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
	if len(result.Packets) != 1 || len(result.Packets[0].Operations) != 1 {
		t.Fatalf("packets = %#v", result.Packets)
	}
	operation := result.Packets[0].Operations[0]
	if operation.Kind != "primitive" || operation.Code != "u16le" {
		t.Fatalf("operation = %#v, want resolved local call", operation)
	}
}

func TestExtractTraversesNestedFunctionLiteral(t *testing.T) {
	root := t.TempDir()
	packetDir := filepath.Join(root, "minecraft", "protocol", "packet")
	if err := os.MkdirAll(packetDir, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(packetDir, "fixture.go"), []byte(`package packet
import "github.com/sandertv/gophertunnel/minecraft/protocol"
const IDFixture = 1
type Fixture struct { Value uint16 }
func (*Fixture) ID() uint32 { return IDFixture }
func (pk *Fixture) Marshal(io protocol.IO) {
	encode := func() { io.Uint16(&pk.Value) }
	encode()
}
`), 0o644); err != nil {
		t.Fatal(err)
	}
	result, err := Extract(root)
	if err != nil {
		t.Fatal(err)
	}
	if len(result.Packets) != 1 || len(result.Packets[0].Operations) != 1 || result.Packets[0].Operations[0].Code != "u16le" {
		t.Fatalf("operations = %#v, want nested function wire operation", result.Packets)
	}
}

func TestComparePreservesExactUnionVariantValues(t *testing.T) {
	union := manifest.Union(
		manifest.Primitive("var_u32"),
		manifest.Variant{Value: 0, Name: "Zero", Encode: manifest.Void()},
		manifest.Variant{Value: 1, Name: "One", Encode: manifest.Void()},
	)
	m := fixtureManifest(union)
	operations := []sourceOperation{
		{Kind: "primitive", Code: "var_u32"},
		{Kind: "switch", CompareTo: "Field", Variants: []sourceVariant{
			{Values: []int64{0}},
			{Values: []int64{2}},
		}},
	}
	source := extraction{Packets: []sourcePacket{{ID: 1, Name: "Fixture", Operations: operations, Paths: expandSourcePaths(operations)}}}
	report := Compare(m, source, fixtureLock(), emptyAccepted(), "manifest.json")
	if report.Counts.Divergence != 1 || report.Counts.Agreement != 0 {
		t.Fatalf("counts = %#v, packets = %#v", report.Counts, report.Packets)
	}
	if len(report.Packets[0].Paths) == 0 || report.Packets[0].Paths[0].ManifestConstraint == "" {
		t.Fatalf("variant path evidence missing: %#v", report.Packets[0])
	}
}

func TestSymbolicComparisonAvoidsOptionalCartesianProduct(t *testing.T) {
	fields := make([]manifest.Node, 6)
	operations := make([]sourceOperation, 0, len(fields))
	for index := range fields {
		fields[index] = manifest.Optional(manifest.Union(
			manifest.Primitive("var_u32"),
			manifest.Variant{Value: 0, Name: "Zero", Encode: manifest.Primitive("u8")},
			manifest.Variant{Value: 1, Name: "One", Encode: manifest.Primitive("u8")},
		))
		operations = append(operations, sourceOperation{Kind: "optional", Field: "Field", Presence: "bool", Value: []sourceOperation{
			{Kind: "primitive", Code: "var_u32"},
			{Kind: "union", Control: "", Variants: []sourceVariant{
				{Value: 0, Ops: []sourceOperation{{Kind: "primitive", Code: "u8"}}},
				{Value: 1, Ops: []sourceOperation{{Kind: "primitive", Code: "u8"}}},
			}},
		}})
	}
	m := fixtureManifest(fields...)
	source := extraction{Packets: []sourcePacket{{ID: 1, Name: "Fixture", Operations: operations, Paths: expandSourcePaths(operations)}}}
	if len(source.Packets[0].Paths) != 1 {
		t.Fatalf("fixture should hit the bounded path product: %d paths", len(source.Packets[0].Paths))
	}
	report := Compare(m, source, fixtureLock(), emptyAccepted(), "manifest.json")
	if report.Counts.Agreement != 1 || report.Counts.Unresolved != 0 {
		t.Fatalf("counts = %#v, packets = %#v", report.Counts, report.Packets)
	}
}

func TestExternalLengthIsCoalescedWithTheFollowingArray(t *testing.T) {
	atoms, reasons := sourcePathAtoms(sourcePath{Operations: []sourceOperation{
		{Kind: "primitive", Code: "u32le"},
		{Kind: "array", Prefix: "u32le", ConsumesPrefix: true, Element: []sourceOperation{{Kind: "primitive", Code: "u8"}}},
	}})
	if len(reasons) != 0 || len(atoms) != 1 || atoms[0].Token != "LEN:FIXED32LE" {
		t.Fatalf("external length was not coalesced: atoms=%#v reasons=%#v", atoms, reasons)
	}
}

func TestReviewedHelperRequiresPinnedRevision(t *testing.T) {
	root := t.TempDir()
	protocolDir := filepath.Join(root, "minecraft", "protocol")
	packetDir := filepath.Join(protocolDir, "packet")
	if err := os.MkdirAll(packetDir, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(protocolDir, "writer.go"), []byte(`package protocol
type IO interface{}
type Color struct{}
type Writer struct{}
func (w *Writer) Float32(*float32) {}
func (w *Writer) RGB(x *Color) { var value float32; w.Float32(&value); w.Float32(&value); w.Float32(&value) }
`), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(packetDir, "fixture.go"), []byte(`package packet
import "example/protocol"
const IDFixture = 1
type Fixture struct { Value uint32 }
func (*Fixture) ID() uint32 { return IDFixture }
func (pk *Fixture) Marshal(io protocol.IO) { io.RGB(nil) }
`), 0o644); err != nil {
		t.Fatal(err)
	}
	unpinned, err := ExtractAtRevision(root, "1111111111111111111111111111111111111111")
	if err != nil {
		t.Fatal(err)
	}
	if len(unpinned.Packets) != 1 || len(unpinned.Packets[0].Operations) != 1 || unpinned.Packets[0].Operations[0].Kind != "unresolved" {
		t.Fatalf("unpinned helper was admitted: %#v", unpinned.Packets)
	}
	pinned, err := ExtractAtRevision(root, reviewedHelperRevision)
	if err != nil {
		t.Fatal(err)
	}
	if len(pinned.Packets) != 1 || len(pinned.Packets[0].Operations) != 3 || pinned.Packets[0].Operations[0].Code != "f32le" {
		t.Fatalf("pinned helper was not expanded: %#v", pinned.Packets)
	}
}

func emptyAccepted() AcceptedFile {
	return AcceptedFile{SchemaVersion: AcceptedSchemaVersion, MinecraftVersion: "1.26.40", ProtocolVersion: 2168}
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

func TestAcceptedEvidenceMustCiteTheLockedOracle(t *testing.T) {
	lock := fixtureLock()
	lock.Gophertunnel.Repo = "https://github.com/example/gophertunnel.git"
	entry := AcceptedDivergence{ID: 1, Name: "P", Reason: "reviewed", WhatWouldSettleIt: "capture"}
	for _, locator := range []string{
		"https://github.com/other/gophertunnel/blob/" + lock.Gophertunnel.Commit + "/minecraft/protocol/packet/p.go",
		"https://github.com/example/gophertunnel/blob/1111111111111111111111111111111111111111/minecraft/protocol/packet/p.go",
	} {
		entry.Evidence = []Evidence{{Locator: locator, Summary: "pinned marshal"}}
		file := AcceptedFile{SchemaVersion: AcceptedSchemaVersion, MinecraftVersion: "1.26.40", ProtocolVersion: 2168, Divergences: []AcceptedDivergence{entry}}
		if err := checkAcceptedEvidence(lock, file); err == nil {
			t.Errorf("locator %q outside the locked oracle was accepted", locator)
		}
	}

	entry.Evidence = []Evidence{
		{Locator: "https://github.com/example/gophertunnel/blob/" + lock.Gophertunnel.Commit + "/minecraft/protocol/packet/p.go", Summary: "pinned marshal"},
		{Locator: "https://github.com/CloudburstMC/Protocol/blob/fbbeee7/Serializer.java", Summary: "independent serializer"},
	}
	file := AcceptedFile{SchemaVersion: AcceptedSchemaVersion, MinecraftVersion: "1.26.40", ProtocolVersion: 2168, Divergences: []AcceptedDivergence{entry}}
	if err := checkAcceptedEvidence(lock, file); err != nil {
		t.Fatalf("locked-oracle and independent evidence were rejected: %v", err)
	}
}

// TestCheckedInOracleBaselineMatchesTheCanonicalManifest keeps the committed
// lock and reviewed baseline from drifting away from the manifest without a
// gophertunnel checkout being available.
func TestCheckedInOracleBaselineMatchesTheCanonicalManifest(t *testing.T) {
	canonical, err := manifest.Load(filepath.Join("..", "..", "generated", "1.26.40", "manifest.json"))
	if err != nil {
		t.Fatal(err)
	}
	lock, err := LoadLock(filepath.Join("..", "..", "tools", "gophertunnel-oracle", "lock.json"))
	if err != nil {
		t.Fatal(err)
	}
	accepted, err := LoadAccepted(filepath.Join("..", "..", "tools", "gophertunnel-oracle", "accepted-divergences.json"))
	if err != nil {
		t.Fatal(err)
	}
	if lock.MinecraftVersion != canonical.Target.MinecraftVersion || lock.ProtocolVersion != canonical.Target.ProtocolVersion {
		t.Fatalf("lock targets %s/%d, manifest targets %s/%d", lock.MinecraftVersion, lock.ProtocolVersion, canonical.Target.MinecraftVersion, canonical.Target.ProtocolVersion)
	}
	if accepted.MinecraftVersion != canonical.Target.MinecraftVersion || accepted.ProtocolVersion != canonical.Target.ProtocolVersion {
		t.Fatalf("baseline targets %s/%d, manifest targets %s/%d", accepted.MinecraftVersion, accepted.ProtocolVersion, canonical.Target.MinecraftVersion, canonical.Target.ProtocolVersion)
	}
	if err := checkAcceptedEvidence(lock, accepted); err != nil {
		t.Fatal(err)
	}
	names := make(map[uint32]string, len(canonical.Packets))
	for _, packet := range canonical.Packets {
		names[packet.ID] = packet.Name
	}
	for _, entry := range accepted.Divergences {
		name, ok := names[entry.ID]
		if !ok {
			t.Errorf("accepted divergence %d (%s) has no manifest packet", entry.ID, entry.Name)
			continue
		}
		if name != entry.Name {
			t.Errorf("accepted divergence %d names %q, manifest packet is %q", entry.ID, entry.Name, name)
		}
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
