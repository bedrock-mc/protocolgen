// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type InventorySource struct {
	SourceType  InventorySourceType
	ContainerID Optional[int8]
	BitFlags    Optional[InventorySourceInventorySourceFlags]
}

// Marshal reads or writes InventorySource using its canonical wire layout.
func (x *InventorySource) Marshal(io IO) {
	IntegerFunc(&x.SourceType, io.Varuint32)
	DoubleOptionalFunc(io, &x.ContainerID, io.Int8)
	DoubleOptionalFunc(io, &x.BitFlags, func(value *InventorySourceInventorySourceFlags) {
		IntegerFunc(value, io.Varuint32)
	})
}
