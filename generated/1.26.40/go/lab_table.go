// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type LabTable struct {
	Type     LabTableType
	Position BlockPos
	Reaction LabTableReactionType
}

// Marshal reads or writes LabTable using its canonical wire layout.
func (x *LabTable) Marshal(io IO) {
	IntegerFunc(&x.Type, io.Uint8)
	x.Position.Marshal(io)
	IntegerFunc(&x.Reaction, io.Uint8)
}
