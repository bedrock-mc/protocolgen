// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type VoxelShapes struct {
	Shapes           []VoxelShapesSerializableVoxelShape
	NameMap          []OrderedEntry[string, VoxelShapesRegistryHandle]
	CustomShapeCount uint16
}

// Marshal reads or writes VoxelShapes using its canonical wire layout.
func (x *VoxelShapes) Marshal(io IO) {
	FuncSlice(io, &x.Shapes, io.Varuint32, func(value *VoxelShapesSerializableVoxelShape) {
		item := *value
		item.Marshal(io)
		*value = item
	})
	OrderedMap(io, &x.NameMap, io.Varuint32, io.String, func(value *VoxelShapesRegistryHandle) {
		item := *value
		item.Marshal(io)
		*value = item
	})
	io.Uint16(&x.CustomShapeCount)
}
