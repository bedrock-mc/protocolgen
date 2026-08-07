// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type ItemStackRequestCerealSlotInfoData struct {
	FullContainerName FullContainerName
	Slot              uint8
	NetIdVariant      int32
}

// Marshal reads or writes ItemStackRequestCerealSlotInfoData using its canonical wire layout.
func (x *ItemStackRequestCerealSlotInfoData) Marshal(io IO) {
	x.FullContainerName.Marshal(io)
	io.Uint8(&x.Slot)
	io.Int32(&x.NetIdVariant)
}
