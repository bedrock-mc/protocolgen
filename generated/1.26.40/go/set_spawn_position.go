// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type SetSpawnPosition struct {
	SpawnPositionType SpawnPositionType
	BlockPosition     BlockPos
	DimensionType     DimensionType
	SpawnBlockPos     BlockPos
}

// Marshal reads or writes SetSpawnPosition using its canonical wire layout.
func (x *SetSpawnPosition) Marshal(io IO) {
	IntegerFunc(&x.SpawnPositionType, io.Varint32)
	x.BlockPosition.Marshal(io)
	x.DimensionType.Marshal(io)
	x.SpawnBlockPos.Marshal(io)
}
