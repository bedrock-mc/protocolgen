// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type StructureBlockUpdate struct {
	BlockPosition BlockPos
	StructureData StructureEditorData
	Trigger       bool
	IsWaterlogged bool
}

// Marshal reads or writes StructureBlockUpdate using its canonical wire layout.
func (x *StructureBlockUpdate) Marshal(io IO) {
	x.BlockPosition.Marshal(io)
	x.StructureData.Marshal(io)
	io.Bool(&x.Trigger)
	io.Bool(&x.IsWaterlogged)
}
