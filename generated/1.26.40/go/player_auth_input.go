// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "github.com/go-gl/mathgl/mgl32"

type PlayerAuthInput struct {
	PlayerRotation         mgl32.Vec2
	Position               mgl32.Vec3
	MoveVector             mgl32.Vec2
	PlayerHeadRotation     float32
	InputData              Optional[[]PlayerAuthInputInputData]
	InputMode              InputMode
	PlayMode               ClientPlayMode
	NewInteractionModel    NewInteractionModel
	InteractRotation       mgl32.Vec2
	ClientTick             uint64
	PosDelta               mgl32.Vec3
	ItemUseTransaction     Optional[PackedItemUseLegacyInventoryTransaction]
	ItemStackRequest       Optional[ItemStackRequestCerealRequestData]
	PlayerBlockActions     Optional[[]PlayerBlockActionData]
	VehicleRotation        Optional[mgl32.Vec2]
	ClientPredictedVehicle Optional[int64]
	AnalogMoveVector       mgl32.Vec2
	CameraOrientation      mgl32.Vec3
	RawMoveVector          mgl32.Vec2
}

// Marshal reads or writes PlayerAuthInput using its canonical wire layout.
func (x *PlayerAuthInput) Marshal(io IO) {
	io.Vec2(&x.PlayerRotation)
	io.Vec3(&x.Position)
	io.Vec2(&x.MoveVector)
	io.Float32(&x.PlayerHeadRotation)
	OptionalFunc(io, &x.InputData, func(value *[]PlayerAuthInputInputData) {
		FuncSlice(io, value, io.Varuint32, func(value *PlayerAuthInputInputData) {
			IntegerFunc(value, io.Varint32)
		})
	})
	IntegerFunc(&x.InputMode, io.Varuint32)
	IntegerFunc(&x.PlayMode, io.Varuint32)
	IntegerFunc(&x.NewInteractionModel, io.Varint32)
	io.Vec2(&x.InteractRotation)
	io.PlayerInputTick(&x.ClientTick)
	io.Vec3(&x.PosDelta)
	DoubleOptionalFunc(io, &x.ItemUseTransaction, func(value *PackedItemUseLegacyInventoryTransaction) {
		value.Marshal(io)
	})
	DoubleOptionalFunc(io, &x.ItemStackRequest, func(value *ItemStackRequestCerealRequestData) {
		value.Marshal(io)
	})
	DoubleOptionalFunc(io, &x.PlayerBlockActions, func(value *[]PlayerBlockActionData) {
		FuncSlice(io, value, io.Varuint32, func(value *PlayerBlockActionData) {
			value.Marshal(io)
		})
	})
	DoubleOptionalFunc(io, &x.VehicleRotation, io.Vec2)
	DoubleOptionalFunc(io, &x.ClientPredictedVehicle, io.ActorUniqueID)
	io.Vec2(&x.AnalogMoveVector)
	io.Vec3(&x.CameraOrientation)
	io.Vec2(&x.RawMoveVector)
}
