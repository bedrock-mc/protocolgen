// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type BlockPickRequest struct {
	Position BlockPos
	WithData bool
	MaxSlots uint8
}

// Marshal reads or writes BlockPickRequest using its canonical wire layout.
func (x *BlockPickRequest) Marshal(io IO) {
	x.Position.Marshal(io)
	io.Bool(&x.WithData)
	io.Uint8(&x.MaxSlots)
}
