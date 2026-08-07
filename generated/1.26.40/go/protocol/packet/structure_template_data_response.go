// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type StructureTemplateDataResponse struct {
	StructureName string
	StructureSNBT []byte
	ResponseType  protocol.StructureTemplateResponseType
}

// Marshal reads or writes StructureTemplateDataResponse using its canonical wire layout.
func (x *StructureTemplateDataResponse) Marshal(io protocol.IO) {
	io.String(&x.StructureName)
	io.NBT(&x.StructureSNBT, protocol.NBTNetwork)
	protocol.IntegerFunc(&x.ResponseType, io.Uint8)
}

// ID returns the protocol ID for StructureTemplateDataResponse.
func (*StructureTemplateDataResponse) ID() uint32 { return IDStructureTemplateDataResponse }
