// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type StackRequestSlotInfo struct {
	FullContainerName FullContainerName
	Slot              uint8
	NetIDVariant      int32
}

// Marshal reads or writes StackRequestSlotInfo using its canonical wire layout.
func (x *StackRequestSlotInfo) Marshal(io IO) {
	x.FullContainerName.Marshal(io)
	io.Uint8(&x.Slot)
	io.Int32(&x.NetIDVariant)
}
