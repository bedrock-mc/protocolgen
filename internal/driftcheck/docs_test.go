package driftcheck

import "testing"

// synthetic doc modeled on Mojang's JSON-Schema format (NOT vendored Mojang content).
const exampleDocJSON = `{
	"$schema": "http://json-schema.org/draft-07/schema#",
	"title": "ExamplePacket",
	"type": "object",
	"definitions": {
		"111": {
			"title": "RuntimeID",
			"type": "object",
			"properties": {
				"Runtime ID": {"type": "integer", "x-underlying-type": "uint64", "x-serialization-options": ["Compression"], "x-ordinal-index": 0}
			},
			"required": ["Runtime ID"]
		},
		"222": {
			"title": "BlockPos",
			"type": "object",
			"properties": {
				"X": {"type": "integer", "x-underlying-type": "int32", "x-serialization-options": ["Compression"], "x-ordinal-index": 0},
				"Y": {"type": "integer", "x-underlying-type": "int32", "x-serialization-options": ["Compression"], "x-ordinal-index": 1},
				"Z": {"type": "integer", "x-underlying-type": "int32", "x-serialization-options": ["Compression"], "x-ordinal-index": 2}
			}
		}
	},
	"properties": {
		"Action": {"type": "string", "enum": ["A", "B"], "x-underlying-type": "uint8", "x-serialization-options": ["Enum-as-Value"], "x-ordinal-index": 0},
		"Data": {"type": "number", "x-underlying-type": "float", "x-ordinal-index": 3},
		"Target": {"$ref": "#/definitions/111", "x-ordinal-index": 1},
		"Position": {"$ref": "#/definitions/222", "x-ordinal-index": 2}
	},
	"required": ["Action", "Data", "Target"],
	"$metaProperties": {"[cereal:packet]": 44}
}`

func TestParseDocPacketExtractsNameAndID(t *testing.T) {
	pk, err := parseDocPacket([]byte(exampleDocJSON))
	if err != nil {
		t.Fatalf("parseDocPacket: %v", err)
	}
	if pk.Name != "ExamplePacket" {
		t.Errorf("Name = %q, want ExamplePacket", pk.Name)
	}
	if pk.ID != 44 {
		t.Errorf("ID = %d, want 44", pk.ID)
	}
}

func TestParseDocPacketSortsFieldsByOrdinal(t *testing.T) {
	pk, err := parseDocPacket([]byte(exampleDocJSON))
	if err != nil {
		t.Fatalf("parseDocPacket: %v", err)
	}
	want := []string{"Action", "Target", "Position", "Data"}
	if len(pk.Fields) != len(want) {
		t.Fatalf("len(Fields) = %d, want %d", len(pk.Fields), len(want))
	}
	for i, name := range want {
		if pk.Fields[i].Name != name {
			t.Errorf("Fields[%d].Name = %q, want %q", i, pk.Fields[i].Name, name)
		}
	}
}

func TestParseDocPacketFlattensSingleScalarRef(t *testing.T) {
	pk, err := parseDocPacket([]byte(exampleDocJSON))
	if err != nil {
		t.Fatalf("parseDocPacket: %v", err)
	}
	target := pk.Fields[1] // "Target", a $ref to single-scalar RuntimeID
	if target.UnderlyingType != "uint64" {
		t.Errorf("Target.UnderlyingType = %q, want uint64 (flattened from ref)", target.UnderlyingType)
	}
	if !target.HasOption("Compression") {
		t.Errorf("Target should carry the Compression option from the referenced scalar")
	}
	if target.RefTitle != "RuntimeID" {
		t.Errorf("Target.RefTitle = %q, want RuntimeID", target.RefTitle)
	}
	if target.Composite {
		t.Errorf("single-scalar ref should not be marked Composite")
	}
}

func TestParseDocPacketRecordsEnum(t *testing.T) {
	pk, err := parseDocPacket([]byte(exampleDocJSON))
	if err != nil {
		t.Fatalf("parseDocPacket: %v", err)
	}
	if !pk.Fields[0].Enum { // "Action" has an enum
		t.Errorf("Action.Enum = false, want true")
	}
	if pk.Fields[3].Enum { // "Data" is a plain float
		t.Errorf("Data.Enum = true, want false")
	}
}

func TestParseDocPacketMarksMultiFieldRefAsComposite(t *testing.T) {
	pk, err := parseDocPacket([]byte(exampleDocJSON))
	if err != nil {
		t.Fatalf("parseDocPacket: %v", err)
	}
	pos := pk.Fields[2] // "Position", a $ref to multi-field BlockPos
	if !pos.Composite {
		t.Errorf("multi-field ref should be marked Composite")
	}
	if pos.RefTitle != "BlockPos" {
		t.Errorf("Position.RefTitle = %q, want BlockPos", pos.RefTitle)
	}
}
