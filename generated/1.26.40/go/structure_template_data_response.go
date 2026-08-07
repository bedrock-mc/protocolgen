// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type StructureTemplateDataResponse struct {
	StructureName string
	StructureSNBT []byte
	ResponseType  StructureTemplateResponseType
}

// Marshal reads or writes StructureTemplateDataResponse using its canonical wire layout.
func (x *StructureTemplateDataResponse) Marshal(io IO) {
	io.String(&x.StructureName)
	io.NBT(&x.StructureSNBT)
	IntegerFunc(&x.ResponseType, io.Uint8)
}
