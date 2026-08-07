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
	io.Bool(&x.LegacySetItemSlots.set)
	if x.LegacySetItemSlots.set {
		if !io.Reading() && uint64(len(x.LegacySetItemSlots.val)) > uint64(^uint32(0)) {
			io.InvalidValue(len(x.LegacySetItemSlots.val), "collection length overflows uint32")
			return
		}
		count1 := uint32(len(x.LegacySetItemSlots.val))
		io.Varuint32(&count1)
		if io.Reading() {
			if uint64(count1) > uint64(^uint(0)>>1) {
				io.InvalidValue(count1, "collection length overflows int")
				return
			}
			x.LegacySetItemSlots.val = make([]LegacySetSlot, int(count1))
		}
		for index2 := range x.LegacySetItemSlots.val {
			x.LegacySetItemSlots.val[index2].Marshal(io)
		}
	} else if io.Reading() {
		var zero []LegacySetSlot
		x.LegacySetItemSlots.val = zero
	}
	io.Bool(&x.Transaction.set)
	if x.Transaction.set {
		marshalInventoryTransactionTransactionValue(io, &x.Transaction.val)
	} else if io.Reading() {
		var zero InventoryTransactionTransactionValue
		x.Transaction.val = zero
	}
}
