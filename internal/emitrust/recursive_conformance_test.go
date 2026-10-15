package emitrust

import (
	"testing"

	"protocolgen/internal/manifest"
)

// recursiveFixture uses non-default tags and a fixed-width list length so the
// conformance fixtures must follow the schema rather than one protocol's names.
func recursiveFixture() manifest.Manifest {
	const identity = "fixture::Value"
	value := manifest.Union(manifest.Primitive("i16le"),
		manifest.Variant{Value: 10, Name: "Leaf", Encode: manifest.Void()},
		manifest.Variant{Value: 20, Name: "Sequence", Encode: manifest.Array(manifest.Primitive("u32le"), manifest.Recursive(identity))},
		manifest.Variant{Value: 30, Name: "Dictionary", Encode: manifest.Map(manifest.Primitive("var_u32"), manifest.String(manifest.Primitive("var_u32")), manifest.Recursive(identity))},
	)
	value.TypeID = identity
	return manifest.Manifest{
		SchemaVersion: 2,
		Target:        manifest.Target{MinecraftVersion: "fixture", ProtocolVersion: 1},
		Sources:       []manifest.SourcePin{{ID: "fixture", Kind: "synthetic", Revision: "1", Digest: "fixture:recursive", MinecraftVersion: "fixture", ProtocolVersion: 1}},
		Packets: []manifest.Packet{{ID: 1, Name: "FixturePacket", Direction: manifest.DirectionClientbound, Fields: []manifest.Field{
			{Ordinal: 0, Name: "Value", Encode: value, Symmetry: manifest.Symmetric, Provenance: manifest.Provenance{Pins: []string{"fixture"}}},
		}}},
	}
}

func TestRecursiveConformanceFollowsSchema(t *testing.T) {
	for _, contract := range []int{0, CodecContractVersion} {
		g, _, err := prepareWithOptions(recursiveFixture(), Options{CodecContract: contract})
		if err != nil {
			t.Fatal(err)
		}
		var fixtures []recursiveValueFixture
		for _, item := range g.sortedDefinitions() {
			if fixture, ok := recursiveValueShape(item); ok {
				fixtures = append(fixtures, fixture)
			}
		}
		if len(fixtures) != 1 {
			t.Fatalf("contract %d: got %d recursive fixtures, want 1", contract, len(fixtures))
		}
		got := fixtures[0]
		if got.control != "I16LE" || got.listPrefix != "U32LE" || got.mapPrefix != "VarUInt" || got.keyPrefix != "VarUInt" || got.leaf != 10 || got.list != 20 || got.dictionary != 30 || got.unknown != 0 {
			t.Fatalf("contract %d: fixture lost schema shape: %+v", contract, got)
		}
		files, err := GenerateFilesWithOptions(recursiveFixture(), Options{CodecContract: contract})
		if err != nil {
			t.Fatal(err)
		}
		if files["tests/recursive_values.rs"] == "" {
			t.Fatalf("contract %d: no public recursive-value conformance tests", contract)
		}
	}
}

func TestRecursiveConformanceSkipsIncompatibleSingletonConstraints(t *testing.T) {
	m := recursiveFixture()
	value := &m.Packets[0].Fields[0].Encode
	value.Variants[1].Encode.Constraints = &manifest.Constraints{MinItems: pointerTo(uint64(2))}
	g, _, err := prepareWithOptions(m, Options{})
	if err != nil {
		t.Fatal(err)
	}
	for _, item := range g.sortedDefinitions() {
		if _, ok := recursiveValueShape(item); ok {
			t.Fatal("conformance singleton would violate the list's schema minimum")
		}
	}
}

func TestRecursiveConformanceSkipsSaturatedTagDomain(t *testing.T) {
	m := recursiveFixture()
	value := &m.Packets[0].Fields[0].Encode
	// Shape selection consumes prepared definitions; names do not participate.
	item := definition{Kind: manifest.KindUnion, TypeID: value.TypeID, ControlCode: "u8"}
	for _, variant := range value.Variants {
		item.Union = append(item.Union, rustVariant{Discriminant: variant.Value, Node: variant.Encode})
	}
	for tag := int64(0); tag < 256; tag++ {
		if tag != 10 && tag != 20 && tag != 30 {
			item.Union = append(item.Union, rustVariant{Discriminant: tag, Node: manifest.Void()})
		}
	}
	if _, ok := recursiveValueShape(item); ok {
		t.Fatal("conformance fixture would truncate its unknown tag into a valid tag")
	}
}
