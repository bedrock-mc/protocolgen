// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.44/go/protocol"

type PlayerHotbar struct {
	SelectedSlot     uint32
	ContainerID      uint8
	ShouldSelectSlot bool
}

// Marshal reads or writes PlayerHotbar using its canonical wire layout.
func (x *PlayerHotbar) Marshal(io protocol.IO) {
	io.Varuint32(&x.SelectedSlot)
	protocol.Minimum(io, &x.SelectedSlot, 0)
	io.Uint8(&x.ContainerID)
	protocol.Minimum(io, &x.ContainerID, 0)
	protocol.Maximum(io, &x.ContainerID, 255)
	io.Bool(&x.ShouldSelectSlot)
}

// ID returns the protocol ID for PlayerHotbar.
func (*PlayerHotbar) ID() uint32 { return IDPlayerHotbar }
