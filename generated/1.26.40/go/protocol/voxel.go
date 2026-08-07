// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type VoxelShapesRegistryHandle struct {
	Value uint16
}

// Marshal reads or writes VoxelShapesRegistryHandle using its canonical wire layout.
func (x *VoxelShapesRegistryHandle) Marshal(io IO) {
	io.Uint16(&x.Value)
}

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

type VoxelShapesSerializableVoxelShape struct {
	Cells        VoxelShapesSerializableCells
	XCoordinates []float32
	YCoordinates []float32
	ZCoordinates []float32
}

// Marshal reads or writes VoxelShapesSerializableVoxelShape using its canonical wire layout.
func (x *VoxelShapesSerializableVoxelShape) Marshal(io IO) {
	x.Cells.Marshal(io)
	FuncSlice(io, &x.XCoordinates, io.Varuint32, io.Float32)
	FuncSlice(io, &x.YCoordinates, io.Varuint32, io.Float32)
	FuncSlice(io, &x.ZCoordinates, io.Varuint32, io.Float32)
}
