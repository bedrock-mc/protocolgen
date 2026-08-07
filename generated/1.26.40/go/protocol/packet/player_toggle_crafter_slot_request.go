// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type PlayerToggleCrafterSlotRequest struct {
	PosX       int32
	PosY       int32
	PosZ       int32
	SlotIndex  uint8
	IsDisabled bool
}

// Marshal reads or writes PlayerToggleCrafterSlotRequest using its canonical wire layout.
func (x *PlayerToggleCrafterSlotRequest) Marshal(io protocol.IO) {
	io.Int32(&x.PosX)
	io.Int32(&x.PosY)
	io.Int32(&x.PosZ)
	io.Uint8(&x.SlotIndex)
	io.Bool(&x.IsDisabled)
}

// ID returns the protocol ID for PlayerToggleCrafterSlotRequest.
func (*PlayerToggleCrafterSlotRequest) ID() uint32 { return IDPlayerToggleCrafterSlotRequest }
