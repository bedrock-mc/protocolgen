// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

// StructureTemplateDataRequest is sent by the client to request data of a structure.
type StructureTemplateDataRequest struct {
	// StructureName is the name of the structure that was set in the structure block's UI. This is the
	// name used to export the structure to a file.
	StructureName string
	// StructurePosition is the position of the structure block that has its template data requested.
	StructurePosition protocol.BlockPos
	// StructureSettings is a struct of settings that should be used for exporting the structure. These
	// settings are identical to the last sent in the StructureBlockUpdate packet by the client.
	StructureSettings protocol.StructureSettings
	// RequestedOperation specifies the type of template data request that the player sent. It is one of
	// the constants found above.
	RequestedOperation protocol.StructureTemplateRequestOperation
}

// Marshal reads or writes StructureTemplateDataRequest using its canonical wire layout.
func (x *StructureTemplateDataRequest) Marshal(io protocol.IO) {
	io.String(&x.StructureName)
	x.StructurePosition.Marshal(io)
	x.StructureSettings.Marshal(io)
	protocol.IntegerFunc(&x.RequestedOperation, io.Uint8)
}

// ID returns the protocol ID for StructureTemplateDataRequest.
func (*StructureTemplateDataRequest) ID() uint32 { return IDStructureTemplateDataRequest }
