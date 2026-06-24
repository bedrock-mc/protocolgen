package main

import "testing"

// synthetic source mirroring the structure of gophertunnel's animate.go.
const animateSrc = `package packet

import "github.com/sandertv/gophertunnel/minecraft/protocol"

type Animate struct {
	ActionType      uint8
	EntityRuntimeID uint64
	Data            float32
	SwingSource     uint8
}

func (*Animate) ID() uint32 { return IDAnimate }

func (pk *Animate) Marshal(io protocol.IO) {
	var swingSource protocol.Optional[string]
	if pk.SwingSource != 0 {
		swingSource = protocol.Option(swingSourceToString(pk.SwingSource))
	}
	io.Uint8(&pk.ActionType)
	io.Varuint64(&pk.EntityRuntimeID)
	io.Float32(&pk.Data)
	protocol.OptionalFunc(io, &swingSource, io.String)
	if val, ok := swingSource.Value(); ok {
		swingSourceFromString(io, &pk.SwingSource, val)
	}
}`

func TestAnalyzePacketsIgnoresMarshalerSubtypesWithoutID(t *testing.T) {
	// PartyInfo is a sub-type (a protocol.Marshaler) with a Marshal method but no
	// ID() method, so it must not be treated as a packet.
	const src = `package packet

import "github.com/sandertv/gophertunnel/minecraft/protocol"

type PartyChanged struct {
	Info protocol.Optional[PartyInfo]
}

func (*PartyChanged) ID() uint32 { return IDPartyChanged }

func (pk *PartyChanged) Marshal(io protocol.IO) {
	protocol.OptionalMarshaler(io, &pk.Info)
}

type PartyInfo struct {
	Role uint8
}

func (x *PartyInfo) Marshal(io protocol.IO) {
	io.Uint8(&x.Role)
}`
	pks, err := analyzePackets([]byte(src))
	if err != nil {
		t.Fatalf("analyzePackets: %v", err)
	}
	if len(pks) != 1 {
		t.Fatalf("len(packets) = %d, want 1 (PartyInfo sub-type must be excluded); got %+v", len(pks), pks)
	}
	if pks[0].TypeName != "PartyChanged" {
		t.Errorf("TypeName = %q, want PartyChanged", pks[0].TypeName)
	}
}

func TestAnalyzePacketsExtractsTypeAndIDConst(t *testing.T) {
	pks, err := analyzePackets([]byte(animateSrc))
	if err != nil {
		t.Fatalf("analyzePackets: %v", err)
	}
	if len(pks) != 1 {
		t.Fatalf("len(packets) = %d, want 1", len(pks))
	}
	if pks[0].TypeName != "Animate" {
		t.Errorf("TypeName = %q, want Animate", pks[0].TypeName)
	}
	if pks[0].IDConst != "IDAnimate" {
		t.Errorf("IDConst = %q, want IDAnimate", pks[0].IDConst)
	}
}

func TestAnalyzePacketsExtractsOrderedWireOps(t *testing.T) {
	pks, err := analyzePackets([]byte(animateSrc))
	if err != nil {
		t.Fatalf("analyzePackets: %v", err)
	}
	ops := pks[0].Ops
	type want struct {
		method  string
		kind    string
		field   string
		guarded bool
	}
	wants := []want{
		{"Uint8", "io", "ActionType", false},
		{"Varuint64", "io", "EntityRuntimeID", false},
		{"Float32", "io", "Data", false},
		{"OptionalFunc", "protocol", "", false},
		{"swingSourceFromString", "custom", "SwingSource", true},
	}
	if len(ops) != len(wants) {
		t.Fatalf("len(ops) = %d, want %d; got %+v", len(ops), len(wants), ops)
	}
	for i, w := range wants {
		if ops[i].Method != w.method || ops[i].Kind != w.kind || ops[i].Field != w.field || ops[i].Guarded != w.guarded {
			t.Errorf("ops[%d] = {Method:%q Kind:%q Field:%q Guarded:%v}, want {%q %q %q %v}",
				i, ops[i].Method, ops[i].Kind, ops[i].Field, ops[i].Guarded, w.method, w.kind, w.field, w.guarded)
		}
	}
}

func TestAnalyzePacketsCapturesInnerIOMethod(t *testing.T) {
	pks, err := analyzePackets([]byte(animateSrc))
	if err != nil {
		t.Fatalf("analyzePackets: %v", err)
	}
	// protocol.OptionalFunc(io, &swingSource, io.String) -> inner io method is "String".
	var optional WireOp
	for _, op := range pks[0].Ops {
		if op.Method == "OptionalFunc" {
			optional = op
		}
	}
	if optional.Inner != "String" {
		t.Errorf("OptionalFunc inner = %q, want String", optional.Inner)
	}
}

func TestAnalyzePacketsSkipsSetupCallsNotTouchingIO(t *testing.T) {
	pks, err := analyzePackets([]byte(animateSrc))
	if err != nil {
		t.Fatalf("analyzePackets: %v", err)
	}
	// protocol.Option(...) and swingSourceToString(...) do not reference io and must not be ops.
	for _, op := range pks[0].Ops {
		if op.Method == "Option" || op.Method == "swingSourceToString" {
			t.Errorf("non-io setup call %q was wrongly recorded as a wire op", op.Method)
		}
	}
}

func TestParsePacketIDsHandlesIotaOffsetAndBlanks(t *testing.T) {
	const src = `package packet

const (
	IDLogin = iota + 1
	IDPlayStatus
	IDServerToClientHandshake
	_
	IDDisconnect
)`
	ids, err := parsePacketIDs([]byte(src))
	if err != nil {
		t.Fatalf("parsePacketIDs: %v", err)
	}
	want := map[string]int{
		"IDLogin":                   1,
		"IDPlayStatus":              2,
		"IDServerToClientHandshake": 3,
		// blank consumes 4
		"IDDisconnect": 5,
	}
	for name, val := range want {
		if ids[name] != val {
			t.Errorf("ids[%q] = %d, want %d", name, ids[name], val)
		}
	}
	if _, ok := ids["_"]; ok {
		t.Errorf("blank identifier should not be recorded")
	}
}
