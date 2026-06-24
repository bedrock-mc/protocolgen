package packet

import "github.com/sandertv/gophertunnel/minecraft/protocol"

type Alpha struct {
	A uint8
	B uint32
}

func (*Alpha) ID() uint32 { return IDAlpha }

func (pk *Alpha) Marshal(io protocol.IO) {
	io.Uint8(&pk.A)
	io.Varuint32(&pk.B)
}
