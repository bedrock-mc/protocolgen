// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type PlayerToggleCrafterSlotRequest struct {
	PosX       int32
	PosY       int32
	PosZ       int32
	SlotIndex  uint8
	IsDisabled bool
}

// Marshal reads or writes PlayerToggleCrafterSlotRequest using its canonical wire layout.
func (x *PlayerToggleCrafterSlotRequest) Marshal(io IO) {
	io.Int32(&x.PosX)
	io.Int32(&x.PosY)
	io.Int32(&x.PosZ)
	io.Uint8(&x.SlotIndex)
	io.Bool(&x.IsDisabled)
}
