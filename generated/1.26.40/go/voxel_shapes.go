// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type VoxelShapes struct {
	Shapes           []VoxelShapesSerializableVoxelShape
	NameMap          []OrderedEntry[string, VoxelShapesRegistryHandle]
	CustomShapeCount uint16
}

// Marshal reads or writes VoxelShapes using its canonical wire layout.
func (x *VoxelShapes) Marshal(io IO) {
	if !io.Reading() && uint64(len(x.Shapes)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.Shapes), "collection length overflows uint32")
		return
	}
	count1 := uint32(len(x.Shapes))
	io.Varuint32(&count1)
	if io.Reading() {
		if uint64(count1) > uint64(^uint(0)>>1) {
			io.InvalidValue(count1, "collection length overflows int")
			return
		}
		x.Shapes = make([]VoxelShapesSerializableVoxelShape, int(count1))
	}
	for index2 := range x.Shapes {
		x.Shapes[index2].Marshal(io)
	}
	if !io.Reading() && uint64(len(x.NameMap)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.NameMap), "map length overflows uint32")
		return
	}
	count3 := uint32(len(x.NameMap))
	io.Varuint32(&count3)
	if io.Reading() {
		if uint64(count3) > uint64(^uint(0)>>1) {
			io.InvalidValue(count3, "map length overflows int")
			return
		}
		x.NameMap = make([]OrderedEntry[string, VoxelShapesRegistryHandle], int(count3))
	}
	for index4 := range x.NameMap {
		io.String(&x.NameMap[index4].Key)
		x.NameMap[index4].Value.Marshal(io)
	}
	io.Uint16(&x.CustomShapeCount)
}
