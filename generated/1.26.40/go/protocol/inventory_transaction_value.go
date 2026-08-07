// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type InventoryTransactionValue interface {
	isInventoryTransactionValue()
}

// MarshalInventoryTransactionValue reads or writes the InventoryTransactionValue union using its canonical wire layout.
func MarshalInventoryTransactionValue(io IO, x *InventoryTransactionValue) {
	UnionFunc(io,
		func() {
			var tag uint32
			io.Varuint32(&tag)
			switch int64(tag) {
			case 0:
				value := new(NormalTransactionData)
				value.Marshal(io)
				*x = value
			case 1:
				value := new(InventoryMismatchData)
				value.Marshal(io)
				*x = value
			case 2:
				value := new(ItemUseInventoryTransaction)
				value.Marshal(io)
				*x = value
			case 3:
				value := new(ItemUseOnActorInventoryTransaction)
				value.Marshal(io)
				*x = value
			case 4:
				value := new(ItemReleaseInventoryTransaction)
				value.Marshal(io)
				*x = value
			default:
				io.InvalidValue(tag, "unknown union tag")
			}
		},
		func() {
			switch value := (*x).(type) {
			case *NormalTransactionData:
				tag := uint32(0)
				io.Varuint32(&tag)
				value.Marshal(io)
			case *InventoryMismatchData:
				tag := uint32(1)
				io.Varuint32(&tag)
				value.Marshal(io)
			case *ItemUseInventoryTransaction:
				tag := uint32(2)
				io.Varuint32(&tag)
				value.Marshal(io)
			case *ItemUseOnActorInventoryTransaction:
				tag := uint32(3)
				io.Varuint32(&tag)
				value.Marshal(io)
			case *ItemReleaseInventoryTransaction:
				tag := uint32(4)
				io.Varuint32(&tag)
				value.Marshal(io)
			default:
				io.InvalidValue(*x, "unknown union value")
			}
		},
	)
}
