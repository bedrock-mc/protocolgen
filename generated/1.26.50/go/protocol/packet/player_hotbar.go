// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.50/go/protocol"
)

type PlayerHotbar struct {
	SelectedSlot     uint32
	ContainerID      uint8
	ShouldSelectSlot bool
}

// ID ...
func (*PlayerHotbar) ID() uint32 {
	return IDPlayerHotbar
}

func (pk *PlayerHotbar) Marshal(io protocol.IO) {
	io.Varuint32(&pk.SelectedSlot)
	io.Uint8(&pk.ContainerID)
	io.Bool(&pk.ShouldSelectSlot)
}
