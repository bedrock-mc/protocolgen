package packet

import "github.com/sandertv/gophertunnel/minecraft/protocol"

// Dynamic mimics gophertunnel's Unknown packet: it has an ID() method but the id
// is a runtime value, not a static IDXxx constant, so it cannot be compared by id.
type Dynamic struct {
	PacketID uint32
}

func (pk *Dynamic) ID() uint32 { return pk.PacketID }

func (pk *Dynamic) Marshal(io protocol.IO) {
	io.Varuint32(&pk.PacketID)
}
