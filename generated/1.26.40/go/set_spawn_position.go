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
	enumValue1 := int32(x.SpawnPositionType)
	io.Varint32(&enumValue1)
	x.SpawnPositionType = SpawnPositionType(enumValue1)
	switch int64(enumValue1) {
	case 0, 1:
	default:
		io.InvalidValue(enumValue1, "unknown enum value")
	}
	x.BlockPosition.Marshal(io)
	x.DimensionType.Marshal(io)
	x.SpawnBlockPos.Marshal(io)
}
