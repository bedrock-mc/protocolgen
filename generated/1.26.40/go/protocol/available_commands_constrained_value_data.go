// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type AvailableCommandsConstrainedValueData struct {
	EnumValueSymbol   uint32
	EnumSymbol        uint32
	ConstraintIndices []uint8
}

// Marshal reads or writes AvailableCommandsConstrainedValueData using its canonical wire layout.
func (x *AvailableCommandsConstrainedValueData) Marshal(io IO) {
	io.Uint32(&x.EnumValueSymbol)
	io.Uint32(&x.EnumSymbol)
	FuncSlice(io, &x.ConstraintIndices, io.Varuint32, io.Uint8)
}
