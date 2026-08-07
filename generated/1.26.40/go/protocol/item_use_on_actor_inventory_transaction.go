// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

import "github.com/go-gl/mathgl/mgl32"

type ItemUseOnActorInventoryTransaction struct {
	Actions      InventoryTransactionData
	RuntimeId    uint64
	ActionType   ItemUseOnActorInventoryTransactionActionType
	Slot         int32
	Item         CerealizerNetworkItemStackDescriptorSerializedData
	FromPosition mgl32.Vec3
	HitPosition  mgl32.Vec3
}

func (ItemUseOnActorInventoryTransaction) isInventoryTransactionTransactionValue() {}

// Marshal reads or writes ItemUseOnActorInventoryTransaction using its canonical wire layout.
func (x *ItemUseOnActorInventoryTransaction) Marshal(io IO) {
	x.Actions.Marshal(io)
	io.ActorRuntimeID(&x.RuntimeId)
	IntegerFunc(&x.ActionType, io.Varint32)
	io.Varint32(&x.Slot)
	x.Item.Marshal(io)
	io.Vec3(&x.FromPosition)
	io.Vec3(&x.HitPosition)
}
