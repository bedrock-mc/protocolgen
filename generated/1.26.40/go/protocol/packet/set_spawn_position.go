// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.40/go/protocol"
)

// SetSpawnPosition is sent by the server to update the spawn position of a player, for example when sleeping
// in a bed.
type SetSpawnPosition struct {
	SpawnPositionType protocol.SpawnPositionType
	BlockPosition     protocol.BlockPos
	DimensionType     protocol.DimensionType
	SpawnBlockPos     protocol.BlockPos
}

// ID returns the protocol ID for SetSpawnPosition.
func (*SetSpawnPosition) ID() uint32 { return IDSetSpawnPosition }

// Marshal reads or writes SetSpawnPosition using its canonical wire layout.
func (pk *SetSpawnPosition) Marshal(io protocol.IO) {
	pk.SpawnPositionType.Marshal(io)
	pk.BlockPosition.Marshal(io)
	pk.DimensionType.Marshal(io)
	pk.SpawnBlockPos.Marshal(io)
}
