// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type LabTable struct {
	Type     protocol.LabTableType
	Position protocol.BlockPos
	Reaction protocol.LabTableReactionType
}

// Marshal reads or writes LabTable using its canonical wire layout.
func (x *LabTable) Marshal(io protocol.IO) {
	protocol.IntegerFunc(&x.Type, io.Uint8)
	x.Position.Marshal(io)
	protocol.IntegerFunc(&x.Reaction, io.Uint8)
}
