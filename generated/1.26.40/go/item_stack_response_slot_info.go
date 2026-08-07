// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type ItemStackResponseSlotInfo struct {
	RequestedSlot        uint8
	Slot                 uint8
	Amount               uint8
	ItemStackNetId       Optional[TypedServerNetIdStructItemStackNetIdTagInt32T0]
	CustomName           BedrockSafetyRedactableString
	DurabilityCorrection int32
}

// Marshal reads or writes ItemStackResponseSlotInfo using its canonical wire layout.
func (x *ItemStackResponseSlotInfo) Marshal(io IO) {
	io.Uint8(&x.RequestedSlot)
	io.Uint8(&x.Slot)
	io.Uint8(&x.Amount)
	DoubleOptionalFunc(io, &x.ItemStackNetId, func(value *TypedServerNetIdStructItemStackNetIdTagInt32T0) {
		item := *value
		item.Marshal(io)
		*value = item
	})
	x.CustomName.Marshal(io)
	io.Varint32(&x.DurabilityCorrection)
}
