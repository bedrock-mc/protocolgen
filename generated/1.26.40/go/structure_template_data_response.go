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
	enumValue1 := uint8(x.ResponseType)
	io.Uint8(&enumValue1)
	x.ResponseType = StructureTemplateResponseType(enumValue1)
	switch int64(enumValue1) {
	case 0, 1, 2:
	default:
		io.InvalidValue(enumValue1, "unknown enum value")
	}
}
