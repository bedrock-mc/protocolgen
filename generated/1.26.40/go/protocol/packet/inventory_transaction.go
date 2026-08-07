// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type InventoryTransaction struct {
	LegacyRequestID    protocol.TypedClientNetIdStructItemStackLegacyRequestIdTagInt32T0
	LegacySetItemSlots protocol.Optional[[]protocol.LegacySetSlot]
	Transaction        protocol.Optional[protocol.InventoryTransactionTransactionValue]
}

// Marshal reads or writes InventoryTransaction using its canonical wire layout.
func (x *InventoryTransaction) Marshal(io protocol.IO) {
	x.LegacyRequestID.Marshal(io)
	protocol.OptionalFunc(io, &x.LegacySetItemSlots, func(value *[]protocol.LegacySetSlot) {
		protocol.FuncSlice(io, value, io.Varuint32, func(value *protocol.LegacySetSlot) {
			value.Marshal(io)
		})
	})
	protocol.OptionalFunc(io, &x.Transaction, func(value *protocol.InventoryTransactionTransactionValue) {
		protocol.MarshalInventoryTransactionTransactionValue(io, value)
	})
}
