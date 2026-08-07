// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type StructureBlockUpdate struct {
	BlockPosition protocol.BlockPos
	StructureData protocol.StructureEditorData
	Trigger       bool
	IsWaterlogged bool
}

// Marshal reads or writes StructureBlockUpdate using its canonical wire layout.
func (x *StructureBlockUpdate) Marshal(io protocol.IO) {
	x.BlockPosition.Marshal(io)
	x.StructureData.Marshal(io)
	io.Bool(&x.Trigger)
	io.Bool(&x.IsWaterlogged)
}

// ID returns the protocol ID for StructureBlockUpdate.
func (*StructureBlockUpdate) ID() uint32 { return IDStructureBlockUpdate }
