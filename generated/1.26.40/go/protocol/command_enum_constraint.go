// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type CommandEnumConstraint struct {
	EnumValueSymbol   uint32
	EnumSymbol        uint32
	ConstraintIndices []uint8
}

// Marshal reads or writes CommandEnumConstraint using its canonical wire layout.
func (x *CommandEnumConstraint) Marshal(io IO) {
	io.Uint32(&x.EnumValueSymbol)
	io.Uint32(&x.EnumSymbol)
	FuncSlice(io, &x.ConstraintIndices, io.Varuint32, io.Uint8)
}
