// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type InventoryTransactionTransactionValue interface {
	isInventoryTransactionTransactionValue()
}

// MarshalInventoryTransactionTransactionValue reads or writes the InventoryTransactionTransactionValue union using its canonical wire layout.
func MarshalInventoryTransactionTransactionValue(io IO, x *InventoryTransactionTransactionValue) {
	UnionFunc(io,
		func() {
			var tag uint32
			io.Varuint32(&tag)
			switch int64(tag) {
			case 0:
				var value NormalTransactionData
				value.Marshal(io)
				*x = value
			case 1:
				var value InventoryMismatchData
				value.Marshal(io)
				*x = value
			case 2:
				var value ItemUseInventoryTransaction
				value.Marshal(io)
				*x = value
			case 3:
				var value ItemUseOnActorInventoryTransaction
				value.Marshal(io)
				*x = value
			case 4:
				var value ItemReleaseInventoryTransaction
				value.Marshal(io)
				*x = value
			default:
				io.InvalidValue(tag, "unknown union tag")
			}
		},
		func() {
			switch value := (*x).(type) {
			case NormalTransactionData:
				tag := uint32(0)
				io.Varuint32(&tag)
				value.Marshal(io)
			case InventoryMismatchData:
				tag := uint32(1)
				io.Varuint32(&tag)
				value.Marshal(io)
			case ItemUseInventoryTransaction:
				tag := uint32(2)
				io.Varuint32(&tag)
				value.Marshal(io)
			case ItemUseOnActorInventoryTransaction:
				tag := uint32(3)
				io.Varuint32(&tag)
				value.Marshal(io)
			case ItemReleaseInventoryTransaction:
				tag := uint32(4)
				io.Varuint32(&tag)
				value.Marshal(io)
			default:
				io.InvalidValue(*x, "unknown union value")
			}
		},
	)
}
