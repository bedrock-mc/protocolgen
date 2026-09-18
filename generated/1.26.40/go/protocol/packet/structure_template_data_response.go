// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.40/go/protocol"
)

// StructureTemplateDataResponse is sent by the server to send data of a structure to the client in response
// to a StructureTemplateDataRequest packet.
type StructureTemplateDataResponse struct {
	// StructureName is the name of the structure that was requested. This is the name used to export the
	// structure to a file.
	StructureName string
	StructureSNBT []byte
	// ResponseType specifies the response type of the packet. This depends on the RequestType field sent in the
	// StructureTemplateDataRequest packet and is one of the constants above.
	ResponseType protocol.StructureTemplateResponseType
}

// ID returns the protocol ID for StructureTemplateDataResponse.
func (*StructureTemplateDataResponse) ID() uint32 { return IDStructureTemplateDataResponse }

// Marshal reads or writes StructureTemplateDataResponse using its canonical wire layout.
func (pk *StructureTemplateDataResponse) Marshal(io protocol.IO) {
	io.String(&pk.StructureName)
	io.NBT(&pk.StructureSNBT, protocol.NBTNetwork)
	pk.ResponseType.Marshal(io)
}
