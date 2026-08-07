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
	enumValue1 := uint8(x.RequestedOperation)
	io.Uint8(&enumValue1)
	x.RequestedOperation = StructureTemplateRequestOperation(enumValue1)
	switch int64(enumValue1) {
	case 0, 1, 2, 3:
	default:
		io.InvalidValue(enumValue1, "unknown enum value")
	}
}
