package main

import (
	"strings"
	"testing"
)

func TestResolveScalars(t *testing.T) {
	cases := []struct {
		f       genField
		goType  string
		marshal string
	}{
		{genField{GoName: "Action", Under: "uint8", JSONType: "integer", Required: true}, "uint8", "io.Uint8(&pk.Action)"},
		{genField{GoName: "RID", Under: "uint64", Options: []string{"Compression"}, JSONType: "integer", Required: true}, "uint64", "io.Varuint64(&pk.RID)"},
		{genField{GoName: "Dim", Under: "int32", Options: []string{"Compression"}, JSONType: "integer", Required: true}, "int32", "io.Varint32(&pk.Dim)"},
		{genField{GoName: "Data", Under: "float", JSONType: "number", Required: true}, "float32", "io.Float32(&pk.Data)"},
		{genField{GoName: "Flag", Under: "boolean", JSONType: "boolean", Required: true}, "bool", "io.Bool(&pk.Flag)"},
		{genField{GoName: "Name", Under: "string", JSONType: "string", Required: true}, "string", "io.String(&pk.Name)"},
	}
	for _, c := range cases {
		s := resolve(c.f)
		if s.GoType != c.goType || s.Marshal != c.marshal {
			t.Errorf("resolve(%+v) = {GoType:%q Marshal:%q}, want {%q %q}", c.f, s.GoType, s.Marshal, c.goType, c.marshal)
		}
		if s.Todo != "" {
			t.Errorf("resolve(%+v) unexpected Todo %q", c.f, s.Todo)
		}
	}
}

func TestResolveStringWithoutUnderlyingType(t *testing.T) {
	// Plain string fields in the docs carry only "type":"string", no x-underlying-type.
	s := resolve(genField{GoName: "LevelID", JSONType: "string", Required: true})
	if s.GoType != "string" || s.Marshal != "io.String(&pk.LevelID)" {
		t.Errorf("string fallback = {%q %q}, want {string io.String(&pk.LevelID)}", s.GoType, s.Marshal)
	}
	if s.Todo != "" {
		t.Errorf("plain string should be clean, got TODO %q", s.Todo)
	}
}

func TestResolveNamespacedKnownComposite(t *testing.T) {
	// "mce::UUID" should resolve to the known uuid.UUID composite.
	s := resolve(genField{GoName: "WorldTemplateID", Composite: true, RefTitle: "mce::UUID"})
	if s.GoType != "uuid.UUID" || s.Marshal != "io.UUID(&pk.WorldTemplateID)" {
		t.Errorf("mce::UUID = {%q %q}, want {uuid.UUID io.UUID(&pk.WorldTemplateID)}", s.GoType, s.Marshal)
	}
}

func TestResolveMessyCompositeYieldsValidIdentifier(t *testing.T) {
	s := resolve(genField{GoName: "Cfg", Composite: true, RefTitle: "Json::Value"})
	if s.GoType != "Value" {
		t.Errorf("GoType = %q, want Value (sanitized)", s.GoType)
	}
	if s.Todo == "" {
		t.Errorf("unknown composite should still carry a TODO")
	}
}

func TestResolveItemStackComposite(t *testing.T) {
	// NetworkItemStackDescriptor is an io-method composite -> io.ItemInstance.
	s := resolve(genField{GoName: "HeldItem", Composite: true, RefTitle: "NetworkItemStackDescriptor"})
	if s.GoType != "protocol.ItemInstance" || s.Marshal != "io.ItemInstance(&pk.HeldItem)" {
		t.Errorf("item single = {%q %q}, want {protocol.ItemInstance io.ItemInstance(&pk.HeldItem)}", s.GoType, s.Marshal)
	}
	if s.Todo != "" {
		t.Errorf("mapped item should have no TODO, got %q", s.Todo)
	}
	a := resolve(genField{GoName: "Content", IsArray: true, Elem: &genField{Composite: true, RefTitle: "NetworkItemStackDescriptor"}})
	if a.GoType != "[]protocol.ItemInstance" || a.Marshal != "protocol.FuncSlice(io, &pk.Content, io.ItemInstance)" {
		t.Errorf("item array = {%q %q}", a.GoType, a.Marshal)
	}
}

func TestResolveCompositeMappingKeysAreCleaned(t *testing.T) {
	// A doc title with spaces must resolve via its cleaned key.
	s := resolve(genField{GoName: "Update", Composite: true, RefTitle: "Data Store Update"})
	if s.GoType != "protocol.DataStoreUpdate" {
		t.Errorf("GoType = %q, want protocol.DataStoreUpdate (cleaned-key lookup)", s.GoType)
	}
	if s.Todo != "" {
		t.Errorf("mapped composite should have no TODO, got %q", s.Todo)
	}
}

func TestResolveMarshalerCompositeSingle(t *testing.T) {
	s := resolve(genField{GoName: "Blob", Composite: true, RefTitle: "MissingBlobData"})
	if s.GoType != "protocol.CacheBlob" {
		t.Errorf("GoType = %q, want protocol.CacheBlob", s.GoType)
	}
	if s.Marshal != "protocol.Single(io, &pk.Blob)" {
		t.Errorf("Marshal = %q, want protocol.Single(io, &pk.Blob)", s.Marshal)
	}
	if s.Todo != "" {
		t.Errorf("mapped composite should have no TODO, got %q", s.Todo)
	}
}

func TestResolveMarshalerCompositeArray(t *testing.T) {
	s := resolve(genField{GoName: "MissingBlobs", IsArray: true, Elem: &genField{Composite: true, RefTitle: "MissingBlobData"}})
	if s.GoType != "[]protocol.CacheBlob" {
		t.Errorf("GoType = %q, want []protocol.CacheBlob", s.GoType)
	}
	if s.Marshal != "protocol.Slice(io, &pk.MissingBlobs)" {
		t.Errorf("Marshal = %q, want protocol.Slice(io, &pk.MissingBlobs)", s.Marshal)
	}
	if s.Todo != "" {
		t.Errorf("mapped composite array should have no TODO, got %q", s.Todo)
	}
}

func TestResolveKnownComposites(t *testing.T) {
	v3 := resolve(genField{GoName: "Position", Composite: true, RefTitle: "Vec3"})
	if v3.GoType != "mgl32.Vec3" || v3.Marshal != "io.Vec3(&pk.Position)" {
		t.Errorf("Vec3 = {%q %q}, want {mgl32.Vec3 io.Vec3(&pk.Position)}", v3.GoType, v3.Marshal)
	}
	if !contains(v3.Imports, "github.com/go-gl/mathgl/mgl32") {
		t.Errorf("Vec3 should import mgl32, got %v", v3.Imports)
	}
	bp := resolve(genField{GoName: "Pos", Composite: true, RefTitle: "BlockPos"})
	if bp.GoType != "protocol.BlockPos" || bp.Marshal != "io.BlockPos(&pk.Pos)" {
		t.Errorf("BlockPos = {%q %q}", bp.GoType, bp.Marshal)
	}
}

func TestResolveArrayOfScalar(t *testing.T) {
	f := genField{GoName: "Slots", IsArray: true, Elem: &genField{Under: "uint8", JSONType: "integer"}}
	s := resolve(f)
	if s.GoType != "[]uint8" {
		t.Errorf("GoType = %q, want []uint8", s.GoType)
	}
	if s.Marshal != "protocol.FuncSlice(io, &pk.Slots, io.Uint8)" {
		t.Errorf("Marshal = %q, want protocol.FuncSlice(io, &pk.Slots, io.Uint8)", s.Marshal)
	}
}

func TestResolveArrayOfComposite(t *testing.T) {
	f := genField{GoName: "Entries", IsArray: true, Elem: &genField{Composite: true, RefTitle: "TotallyUnknownComposite"}}
	s := resolve(f)
	if s.GoType != "[]TotallyUnknownComposite" {
		t.Errorf("GoType = %q, want []TotallyUnknownComposite", s.GoType)
	}
	if !strings.Contains(s.Marshal, "protocol.Slice(io, &pk.Entries)") {
		t.Errorf("Marshal = %q, want protocol.Slice(...)", s.Marshal)
	}
	if s.Todo == "" {
		t.Errorf("array of unknown composite should carry a TODO")
	}
}

func TestResolveUnknownCompositeNeedsTodo(t *testing.T) {
	s := resolve(genField{GoName: "Name", Composite: true, RefTitle: "AttributeName"})
	if s.Todo == "" {
		t.Errorf("unknown composite should carry a TODO; got shape %+v", s)
	}
	if !strings.Contains(s.Marshal, "protocol.Single(io, &pk.Name)") {
		t.Errorf("Marshal = %q, want protocol.Single(...)", s.Marshal)
	}
}

func TestResolveOneOfNeedsTodo(t *testing.T) {
	s := resolve(genField{GoName: "Body", OneOf: true})
	if s.Todo == "" {
		t.Errorf("oneOf union should carry a TODO")
	}
}

func contains(ss []string, s string) bool {
	for _, x := range ss {
		if x == s {
			return true
		}
	}
	return false
}
