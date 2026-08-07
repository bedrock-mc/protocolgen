// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type SetSpawnPosition struct {
	SpawnPositionType protocol.SpawnPositionType
	BlockPosition     protocol.BlockPos
	DimensionType     protocol.DimensionType
	SpawnBlockPos     protocol.BlockPos
}

// Marshal reads or writes SetSpawnPosition using its canonical wire layout.
func (x *SetSpawnPosition) Marshal(io protocol.IO) {
	protocol.IntegerFunc(&x.SpawnPositionType, io.Varint32)
	x.BlockPosition.Marshal(io)
	x.DimensionType.Marshal(io)
	x.SpawnBlockPos.Marshal(io)
}

// ID returns the protocol ID for SetSpawnPosition.
func (*SetSpawnPosition) ID() uint32 { return IDSetSpawnPosition }
