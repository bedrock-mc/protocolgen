// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type PrimitiveShapes struct {
	ArrayOfPrimitiveShapesCanBeAMixOfNewUpdatedOrRemoved []PrimitiveShapeData
}

// Marshal reads or writes PrimitiveShapes using its canonical wire layout.
func (x *PrimitiveShapes) Marshal(io IO) {
	if !io.Reading() && uint64(len(x.ArrayOfPrimitiveShapesCanBeAMixOfNewUpdatedOrRemoved)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.ArrayOfPrimitiveShapesCanBeAMixOfNewUpdatedOrRemoved), "collection length overflows uint32")
		return
	}
	count1 := uint32(len(x.ArrayOfPrimitiveShapesCanBeAMixOfNewUpdatedOrRemoved))
	io.Varuint32(&count1)
	if io.Reading() {
		if uint64(count1) > uint64(^uint(0)>>1) {
			io.InvalidValue(count1, "collection length overflows int")
			return
		}
		x.ArrayOfPrimitiveShapesCanBeAMixOfNewUpdatedOrRemoved = make([]PrimitiveShapeData, int(count1))
	}
	for index2 := range x.ArrayOfPrimitiveShapesCanBeAMixOfNewUpdatedOrRemoved {
		x.ArrayOfPrimitiveShapesCanBeAMixOfNewUpdatedOrRemoved[index2].Marshal(io)
	}
}
