package codegen

import "testing"

const demoDocJSON = `{
	"title": "DemoPacket",
	"type": "object",
	"definitions": {
		"100": {"title": "RuntimeID", "type": "object", "properties": {
			"v": {"type": "integer", "x-underlying-type": "uint64", "x-serialization-options": ["Compression"], "x-ordinal-index": 0}
		}},
		"200": {"title": "AttrData", "type": "object", "properties": {
			"a": {"type": "number", "x-underlying-type": "float", "x-ordinal-index": 0},
			"b": {"type": "number", "x-underlying-type": "float", "x-ordinal-index": 1}
		}},
		"300": {"title": "Vec3", "type": "object", "properties": {
			"X": {"type": "number", "x-underlying-type": "float", "x-ordinal-index": 0},
			"Y": {"type": "number", "x-underlying-type": "float", "x-ordinal-index": 1},
			"Z": {"type": "number", "x-underlying-type": "float", "x-ordinal-index": 2}
		}}
	},
	"properties": {
		"Action": {"type": "string", "enum": ["A", "B"], "x-underlying-type": "uint8", "x-serialization-options": ["Enum-as-Value"], "x-ordinal-index": 0},
		"Target": {"$ref": "#/definitions/100", "x-ordinal-index": 1},
		"Pos": {"$ref": "#/definitions/300", "x-ordinal-index": 2},
		"Attrs": {"type": "array", "items": {"$ref": "#/definitions/200"}, "x-ordinal-index": 3},
		"Slots": {"type": "array", "items": {"type": "integer", "x-underlying-type": "uint8"}, "x-ordinal-index": 4},
		"Body": {"oneOf": [{"type": "object"}], "x-ordinal-index": 5}
	},
	"required": ["Action", "Target", "Pos", "Attrs", "Slots"],
	"$metaProperties": {"[cereal:packet]": 44}
}`

func TestParseGenPacketBasics(t *testing.T) {
	pk, err := parseGenPacket([]byte(demoDocJSON))
	if err != nil {
		t.Fatalf("parseGenPacket: %v", err)
	}
	if pk.TypeName != "Demo" {
		t.Errorf("TypeName = %q, want Demo", pk.TypeName)
	}
	if pk.ID != 44 {
		t.Errorf("ID = %d, want 44", pk.ID)
	}
	wantOrder := []string{"Action", "Target", "Pos", "Attrs", "Slots", "Body"}
	if len(pk.Fields) != len(wantOrder) {
		t.Fatalf("len(Fields) = %d, want %d", len(pk.Fields), len(wantOrder))
	}
	for i, w := range wantOrder {
		if pk.Fields[i].DocName != w {
			t.Errorf("Fields[%d].DocName = %q, want %q", i, pk.Fields[i].DocName, w)
		}
	}
}

func TestParseGenPacketFieldShapes(t *testing.T) {
	pk, _ := parseGenPacket([]byte(demoDocJSON))
	byName := map[string]genField{}
	for _, f := range pk.Fields {
		byName[f.DocName] = f
	}

	if a := byName["Action"]; !a.Enum || !a.Required || a.Under != "uint8" {
		t.Errorf("Action = %+v, want enum+required+uint8", a)
	}
	if tgt := byName["Target"]; tgt.Composite || tgt.Under != "uint64" || !hasOption(tgt.Options, "Compression") || tgt.RefTitle != "RuntimeID" {
		t.Errorf("Target = %+v, want flattened uint64+Compression ref RuntimeID", tgt)
	}
	if p := byName["Pos"]; !p.Composite || p.RefTitle != "Vec3" {
		t.Errorf("Pos = %+v, want composite Vec3", p)
	}
	if a := byName["Attrs"]; !a.IsArray || a.Elem == nil || !a.Elem.Composite || a.Elem.RefTitle != "AttrData" {
		t.Errorf("Attrs = %+v, want array of composite AttrData", a)
	}
	if s := byName["Slots"]; !s.IsArray || s.Elem == nil || s.Elem.Under != "uint8" {
		t.Errorf("Slots = %+v, want array of uint8", s)
	}
	if b := byName["Body"]; !b.OneOf {
		t.Errorf("Body = %+v, want oneOf", b)
	}
	if byName["Body"].Required {
		t.Errorf("Body should be optional (not in required)")
	}
}
