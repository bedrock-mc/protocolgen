// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type PackedItemUseLegacyInventoryTransaction struct {
	LegacyRequestID    TypedClientNetIdStructItemStackLegacyRequestIdTagInt32T0
	LegacySetItemSlots Optional[[]LegacySetSlot]
	ItemUseTransaction Optional[ItemUseInventoryTransaction]
}

// Marshal reads or writes PackedItemUseLegacyInventoryTransaction using its canonical wire layout.
func (x *PackedItemUseLegacyInventoryTransaction) Marshal(io IO) {
	x.LegacyRequestID.Marshal(io)
	OptionalFunc(io, &x.LegacySetItemSlots, func(value *[]LegacySetSlot) {
		item := *value
		FuncSlice(io, &item, io.Varuint32, func(value *LegacySetSlot) {
			item := *value
			item.Marshal(io)
			*value = item
		})
		*value = item
	})
	OptionalFunc(io, &x.ItemUseTransaction, func(value *ItemUseInventoryTransaction) {
		item := *value
		item.Marshal(io)
		*value = item
	})
}
