// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

import "github.com/go-gl/mathgl/mgl32"

type ItemReleaseInventoryTransaction struct {
	Actions      InventoryTransactionData
	ActionType   ItemReleaseInventoryTransactionActionType
	Slot         int32
	Item         CerealizerNetworkItemStackDescriptorSerializedData
	FromPosition mgl32.Vec3
}

func (ItemReleaseInventoryTransaction) isInventoryTransactionTransactionValue() {}

// Marshal reads or writes ItemReleaseInventoryTransaction using its canonical wire layout.
func (x *ItemReleaseInventoryTransaction) Marshal(io IO) {
	x.Actions.Marshal(io)
	IntegerFunc(&x.ActionType, io.Varint32)
	io.Varint32(&x.Slot)
	x.Item.Marshal(io)
	io.Vec3(&x.FromPosition)
}
