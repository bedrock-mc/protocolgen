package main

import (
	"go/format"
	"strings"
	"testing"
)

func TestRawResolvesCompositeToLocalType(t *testing.T) {
	s := resolveRaw(genField{GoName: "Pos", Composite: true, RefTitle: "Vec3"})
	if s.GoType != "Vec3" || s.Marshal != "pk.Pos.Marshal(io)" {
		t.Errorf("composite = {%q %q}, want {Vec3 pk.Pos.Marshal(io)}", s.GoType, s.Marshal)
	}
	a := resolveRaw(genField{GoName: "Items", IsArray: true, Elem: &genField{Composite: true, RefTitle: "ItemData"}})
	if a.GoType != "[]ItemData" || a.Marshal != "Slice(io, &pk.Items)" {
		t.Errorf("composite array = {%q %q}", a.GoType, a.Marshal)
	}
	sc := resolveRaw(genField{GoName: "A", Under: "uint64", Options: []string{"Compression"}, JSONType: "integer"})
	if sc.GoType != "uint64" || sc.Marshal != "io.Varuint64(&pk.A)" {
		t.Errorf("scalar = {%q %q}", sc.GoType, sc.Marshal)
	}
}

func TestRawEmitPacketIsValidGo(t *testing.T) {
	pk := genPacket{Title: "FooPacket", TypeName: "Foo", ID: 42, Fields: []genField{
		{GoName: "A", Under: "uint8", JSONType: "integer", Required: true},
		{GoName: "Pos", Composite: true, RefTitle: "Vec3"},
	}}
	src, _, err := rawFormat{pkg: "x"}.EmitPacket(pk, "")
	if err != nil {
		t.Fatalf("EmitPacket: %v", err)
	}
	if _, e := format.Source([]byte(src)); e != nil {
		t.Fatalf("raw output not valid Go: %v\n%s", e, src)
	}
	for _, want := range []string{"package x", "type FooPacket struct", "func (*FooPacket) ID() uint32 { return 42 }", "pk.Pos.Marshal(io)"} {
		if !strings.Contains(src, want) {
			t.Errorf("raw output missing %q\n%s", want, src)
		}
	}
}

func TestRawAvoidsReservedFieldNames(t *testing.T) {
	pk := genPacket{Title: "FooPacket", ID: 1, Fields: []genField{
		{GoName: "ID", Under: "uint32", JSONType: "integer", Required: true},
	}}
	src, _, _ := rawFormat{pkg: "x"}.EmitPacket(pk, "")
	if _, e := format.Source([]byte(src)); e != nil {
		t.Fatalf("field named ID must be renamed to avoid the ID() method; got invalid Go: %v\n%s", e, src)
	}
	if !strings.Contains(src, "IDField") {
		t.Errorf("field ID should be renamed to IDField\n%s", src)
	}
}
