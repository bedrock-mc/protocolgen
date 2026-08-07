// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type StructureTemplateDataRequest struct {
	StructureName      string
	StructurePosition  protocol.BlockPos
	StructureSettings  protocol.StructureSettings
	RequestedOperation protocol.StructureTemplateRequestOperation
}

// Marshal reads or writes StructureTemplateDataRequest using its canonical wire layout.
func (x *StructureTemplateDataRequest) Marshal(io protocol.IO) {
	io.String(&x.StructureName)
	x.StructurePosition.Marshal(io)
	x.StructureSettings.Marshal(io)
	protocol.IntegerFunc(&x.RequestedOperation, io.Uint8)
}
