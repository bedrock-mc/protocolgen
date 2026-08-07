// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type PrimitiveShapes struct {
	ArrayOfPrimitiveShapesCanBeAMixOfNewUpdatedOrRemoved []PrimitiveShapeData
}

// Marshal reads or writes PrimitiveShapes using its canonical wire layout.
func (x *PrimitiveShapes) Marshal(io IO) {
	FuncSlice(io, &x.ArrayOfPrimitiveShapesCanBeAMixOfNewUpdatedOrRemoved, io.Varuint32, func(value *PrimitiveShapeData) {
		value.Marshal(io)
	})
}
