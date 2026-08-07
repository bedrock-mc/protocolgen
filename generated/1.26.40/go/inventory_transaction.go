// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type InventoryTransaction struct {
	LegacyRequestID    TypedClientNetIdStructItemStackLegacyRequestIdTagInt32T0
	LegacySetItemSlots Optional[[]LegacySetSlot]
	Transaction        Optional[InventoryTransactionTransactionValue]
}

// Marshal reads or writes InventoryTransaction using its canonical wire layout.
func (x *InventoryTransaction) Marshal(io IO) {
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
	OptionalFunc(io, &x.Transaction, func(value *InventoryTransactionTransactionValue) {
		item := *value
		marshalInventoryTransactionTransactionValue(io, &item)
		*value = item
	})
}
