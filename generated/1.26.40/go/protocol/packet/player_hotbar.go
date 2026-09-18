// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.40/go/protocol"
)

type PlayerHotbar struct {
	SelectedSlot     uint32
	ContainerID      uint8
	ShouldSelectSlot bool
}

// ID returns the protocol ID for PlayerHotbar.
func (*PlayerHotbar) ID() uint32 { return IDPlayerHotbar }

// Marshal reads or writes PlayerHotbar using its canonical wire layout.
func (pk *PlayerHotbar) Marshal(io protocol.IO) {
	io.Varuint32(&pk.SelectedSlot)
	io.Uint8(&pk.ContainerID)
	io.Bool(&pk.ShouldSelectSlot)
}
