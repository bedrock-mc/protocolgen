// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type PlayerHotbar struct {
	SelectedSlot     uint32
	ContainerID      uint8
	ShouldSelectSlot bool
}

// Marshal reads or writes PlayerHotbar using its canonical wire layout.
func (x *PlayerHotbar) Marshal(io protocol.IO) {
	io.Varuint32(&x.SelectedSlot)
	io.Uint8(&x.ContainerID)
	io.Bool(&x.ShouldSelectSlot)
}
