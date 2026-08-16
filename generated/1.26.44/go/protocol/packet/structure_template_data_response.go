// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.44/go/protocol"

// StructureTemplateDataResponse is sent by the server to send data of a structure to the client in
// response to a StructureTemplateDataRequest packet.
type StructureTemplateDataResponse struct {
	// StructureName is the name of the structure that was requested. This is the name used to export
	// the structure to a file.
	StructureName string
	StructureSNBT []byte
	// ResponseType specifies the response type of the packet. This depends on the RequestType field
	// sent in the StructureTemplateDataRequest packet and is one of the constants above.
	ResponseType protocol.StructureTemplateResponseType
}

// Marshal reads or writes StructureTemplateDataResponse using its canonical wire layout.
func (x *StructureTemplateDataResponse) Marshal(io protocol.IO) {
	io.String(&x.StructureName)
	io.NBT(&x.StructureSNBT, protocol.NBTNetwork)
	protocol.IntegerFunc(&x.ResponseType, io.Uint8)
}

// ID returns the protocol ID for StructureTemplateDataResponse.
func (*StructureTemplateDataResponse) ID() uint32 { return IDStructureTemplateDataResponse }
