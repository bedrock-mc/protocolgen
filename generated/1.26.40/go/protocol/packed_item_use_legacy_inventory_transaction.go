// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type PackedItemUseLegacyInventoryTransaction struct {
	LegacyRequestID    ItemStackLegacyRequestID
	LegacySetItemSlots Optional[[]LegacySetSlot]
	ItemUseTransaction Optional[ItemUseInventoryTransaction]
}

// Marshal reads or writes PackedItemUseLegacyInventoryTransaction using its canonical wire layout.
func (x *PackedItemUseLegacyInventoryTransaction) Marshal(io IO) {
	x.LegacyRequestID.Marshal(io)
	OptionalFunc(io, &x.LegacySetItemSlots, func(value *[]LegacySetSlot) {
		Slice(io, value)
	})
	OptionalFunc(io, &x.ItemUseTransaction, func(value *ItemUseInventoryTransaction) {
		value.Marshal(io)
	})
}
