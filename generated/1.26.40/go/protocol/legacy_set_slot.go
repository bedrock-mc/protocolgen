// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type LegacySetSlot struct {
	ContainerEnum ContainerEnumName
	Slots         []uint8
}

// Marshal reads or writes LegacySetSlot using its canonical wire layout.
func (x *LegacySetSlot) Marshal(io IO) {
	IntegerFunc(&x.ContainerEnum, io.Uint8)
	FuncSlice(io, &x.Slots, io.Varuint32, io.Uint8)
}
