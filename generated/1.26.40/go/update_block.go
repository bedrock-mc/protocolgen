// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type UpdateBlock struct {
	BlockPosition  BlockPos
	BlockRuntimeID uint32
	Flags          uint32
	Layer          uint32
}

// Marshal reads or writes UpdateBlock using its canonical wire layout.
func (x *UpdateBlock) Marshal(io IO) {
	x.BlockPosition.Marshal(io)
	io.Varuint32(&x.BlockRuntimeID)
	io.Varuint32(&x.Flags)
	io.Varuint32(&x.Layer)
}
