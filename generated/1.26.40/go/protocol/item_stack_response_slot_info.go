// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type ItemStackResponseSlotInfo struct {
	RequestedSlot        uint8
	Slot                 uint8
	Amount               uint8
	ItemStackNetID       Optional[ItemStackNetID]
	CustomName           BedrockSafetyRedactableString
	DurabilityCorrection int32
}

// Marshal reads or writes ItemStackResponseSlotInfo using its canonical wire layout.
func (x *ItemStackResponseSlotInfo) Marshal(io IO) {
	io.Uint8(&x.RequestedSlot)
	io.Uint8(&x.Slot)
	io.Uint8(&x.Amount)
	DoubleOptionalFunc(io, &x.ItemStackNetID, func(value *ItemStackNetID) {
		value.Marshal(io)
	})
	x.CustomName.Marshal(io)
	io.Varint32(&x.DurabilityCorrection)
}
