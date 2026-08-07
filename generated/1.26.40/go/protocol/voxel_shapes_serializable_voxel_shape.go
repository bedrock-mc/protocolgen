// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

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
