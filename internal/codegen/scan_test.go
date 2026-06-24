package codegen

import "testing"

func TestPacketIDConstsInFile(t *testing.T) {
	const src = `package packet

import "github.com/sandertv/gophertunnel/minecraft/protocol"

type Animate struct{ A uint8 }

func (*Animate) ID() uint32 { return IDAnimate }
func (pk *Animate) Marshal(io protocol.IO) { io.Uint8(&pk.A) }

// PartyInfo is a sub-type with Marshal but no ID() — must be ignored.
type PartyInfo struct{ R uint8 }

func (x *PartyInfo) Marshal(io protocol.IO) { io.Uint8(&x.R) }`

	got := packetIDConstsInFile([]byte(src))
	if len(got) != 1 || got[0] != "IDAnimate" {
		t.Errorf("packetIDConstsInFile = %v, want [IDAnimate]", got)
	}
}
