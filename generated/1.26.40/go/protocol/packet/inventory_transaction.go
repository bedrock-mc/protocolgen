// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type InventoryTransaction struct {
	LegacyRequestID    protocol.ItemStackLegacyRequestID
	LegacySetItemSlots protocol.Optional[[]protocol.LegacySetSlot]
	Transaction        protocol.Optional[protocol.InventoryTransactionValue]
}

// Marshal reads or writes InventoryTransaction using its canonical wire layout.
func (x *InventoryTransaction) Marshal(io protocol.IO) {
	x.LegacyRequestID.Marshal(io)
	protocol.OptionalFunc(io, &x.LegacySetItemSlots, func(value *[]protocol.LegacySetSlot) {
		protocol.Slice(io, value)
	})
	protocol.OptionalFunc(io, &x.Transaction, func(value *protocol.InventoryTransactionValue) {
		protocol.MarshalInventoryTransactionValue(io, value)
	})
}

// ID returns the protocol ID for InventoryTransaction.
func (*InventoryTransaction) ID() uint32 { return IDInventoryTransaction }
