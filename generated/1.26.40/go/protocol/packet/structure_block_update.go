// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.40/go/protocol"
)

// StructureBlockUpdate is sent by the client when it updates a structure block using the in-game UI. The data
// it contains depends on the type of structure block that it is. In Minecraft Bedrock Edition v1.11, there is
// only the Export structure block type, but in v1.13 the ones present in Java Edition will, according to the
// wiki, be added too.
type StructureBlockUpdate struct {
	BlockPosition protocol.BlockPos
	StructureData protocol.StructureEditorData
	Trigger       bool
	IsWaterlogged bool
}

// ID ...
func (*StructureBlockUpdate) ID() uint32 {
	return IDStructureBlockUpdate
}

func (pk *StructureBlockUpdate) Marshal(io protocol.IO) {
	pk.BlockPosition.Marshal(io)
	pk.StructureData.Marshal(io)
	io.Bool(&pk.Trigger)
	io.Bool(&pk.IsWaterlogged)
}
