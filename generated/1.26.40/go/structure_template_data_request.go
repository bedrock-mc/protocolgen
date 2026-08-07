// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type StructureTemplateDataRequest struct {
	StructureName      string
	StructurePosition  BlockPos
	StructureSettings  StructureSettings
	RequestedOperation StructureTemplateRequestOperation
}

// Marshal reads or writes StructureTemplateDataRequest using its canonical wire layout.
func (x *StructureTemplateDataRequest) Marshal(io IO) {
	io.String(&x.StructureName)
	x.StructurePosition.Marshal(io)
	x.StructureSettings.Marshal(io)
	IntegerFunc(&x.RequestedOperation, io.Uint8)
}
