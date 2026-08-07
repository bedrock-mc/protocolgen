// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type VoxelShapesSerializableCells struct {
	XSize   uint8
	YSize   uint8
	ZSize   uint8
	Storage []uint8
}

// Marshal reads or writes VoxelShapesSerializableCells using its canonical wire layout.
func (x *VoxelShapesSerializableCells) Marshal(io IO) {
	io.Uint8(&x.XSize)
	io.Uint8(&x.YSize)
	io.Uint8(&x.ZSize)
	FuncSlice(io, &x.Storage, io.Varuint32, io.Uint8)
}
