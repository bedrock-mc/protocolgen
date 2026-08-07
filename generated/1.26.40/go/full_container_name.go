// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type FullContainerName struct {
	ContainerName ContainerEnumName
	DynamicID     Optional[uint32]
}

// Marshal reads or writes FullContainerName using its canonical wire layout.
func (x *FullContainerName) Marshal(io IO) {
	IntegerFunc(&x.ContainerName, io.Uint8)
	OptionalFunc(io, &x.DynamicID, io.Uint32)
}
