// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type LabTable struct {
	Type     LabTableType
	Position BlockPos
	Reaction LabTableReactionType
}

// Marshal reads or writes LabTable using its canonical wire layout.
func (x *LabTable) Marshal(io IO) {
	enumValue1 := uint8(x.Type)
	io.Uint8(&enumValue1)
	x.Type = LabTableType(enumValue1)
	switch int64(enumValue1) {
	case 0, 1, 2:
	default:
		io.InvalidValue(enumValue1, "unknown enum value")
	}
	x.Position.Marshal(io)
	enumValue2 := uint8(x.Reaction)
	io.Uint8(&enumValue2)
	x.Reaction = LabTableReactionType(enumValue2)
	switch int64(enumValue2) {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12:
	default:
		io.InvalidValue(enumValue2, "unknown enum value")
	}
}
