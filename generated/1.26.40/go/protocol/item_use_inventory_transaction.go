// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

import "github.com/go-gl/mathgl/mgl32"

type ItemUseInventoryTransaction struct {
	Actions                  InventoryTransactionData
	ActionType               ItemUseInventoryTransactionActionType
	TriggerType              ItemUseInventoryTransactionTriggerType
	Position                 BlockPos
	Face                     uint8
	Slot                     int32
	Item                     CerealizerNetworkItemStackDescriptorSerializedData
	FromPosition             mgl32.Vec3
	ClickPosition            mgl32.Vec3
	TargetBlockId            uint32
	ClientInteractPrediction ItemUseInventoryTransactionPredictedResult
	ClientCooldownState      ItemUseInventoryTransactionClientCooldownState
}

func (ItemUseInventoryTransaction) isInventoryTransactionTransactionValue() {}

// Marshal reads or writes ItemUseInventoryTransaction using its canonical wire layout.
func (x *ItemUseInventoryTransaction) Marshal(io IO) {
	x.Actions.Marshal(io)
	IntegerFunc(&x.ActionType, io.Varint32)
	IntegerFunc(&x.TriggerType, io.Uint8)
	x.Position.Marshal(io)
	io.Uint8(&x.Face)
	io.Varint32(&x.Slot)
	x.Item.Marshal(io)
	io.Vec3(&x.FromPosition)
	io.Vec3(&x.ClickPosition)
	io.Varuint32(&x.TargetBlockId)
	IntegerFunc(&x.ClientInteractPrediction, io.Uint8)
	IntegerFunc(&x.ClientCooldownState, io.Uint8)
}
