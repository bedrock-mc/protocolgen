// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.40/go/protocol"

	"github.com/go-gl/mathgl/mgl32"
)

type PlayerAuthInput struct {
	PlayerRotation         mgl32.Vec2
	Position               mgl32.Vec3
	MoveVector             mgl32.Vec2
	PlayerHeadRotation     float32
	InputData              protocol.Optional[[]protocol.PlayerAuthInputInputData]
	InputMode              protocol.InputMode
	PlayMode               protocol.ClientPlayMode
	NewInteractionModel    protocol.NewInteractionModel
	InteractRotation       mgl32.Vec2
	ClientTick             uint64
	PosDelta               mgl32.Vec3
	ItemUseTransaction     protocol.Optional[protocol.PackedItemUseLegacyInventoryTransaction]
	ItemStackRequest       protocol.Optional[protocol.ItemStackRequestCerealRequestData]
	PlayerBlockActions     protocol.Optional[[]protocol.PlayerBlockActionData]
	VehicleRotation        protocol.Optional[mgl32.Vec2]
	ClientPredictedVehicle protocol.Optional[int64]
	AnalogMoveVector       mgl32.Vec2
	CameraOrientation      mgl32.Vec3
	RawMoveVector          mgl32.Vec2
}

// Marshal reads or writes PlayerAuthInput using its canonical wire layout.
func (x *PlayerAuthInput) Marshal(io protocol.IO) {
	io.Vec2(&x.PlayerRotation)
	io.Vec3(&x.Position)
	io.Vec2(&x.MoveVector)
	io.Float32(&x.PlayerHeadRotation)
	protocol.OptionalFunc(io, &x.InputData, func(value *[]protocol.PlayerAuthInputInputData) {
		protocol.FuncSlice(io, value, io.Varuint32, func(value *protocol.PlayerAuthInputInputData) {
			protocol.IntegerFunc(value, io.Varint32)
		})
	})
	protocol.IntegerFunc(&x.InputMode, io.Varuint32)
	protocol.IntegerFunc(&x.PlayMode, io.Varuint32)
	protocol.IntegerFunc(&x.NewInteractionModel, io.Varint32)
	io.Vec2(&x.InteractRotation)
	io.PlayerInputTick(&x.ClientTick)
	io.Vec3(&x.PosDelta)
	protocol.DoubleOptionalFunc(io, &x.ItemUseTransaction, func(value *protocol.PackedItemUseLegacyInventoryTransaction) {
		value.Marshal(io)
	})
	protocol.DoubleOptionalFunc(io, &x.ItemStackRequest, func(value *protocol.ItemStackRequestCerealRequestData) {
		value.Marshal(io)
	})
	protocol.DoubleOptionalFunc(io, &x.PlayerBlockActions, func(value *[]protocol.PlayerBlockActionData) {
		protocol.Slice(io, value)
	})
	protocol.DoubleOptionalFunc(io, &x.VehicleRotation, io.Vec2)
	protocol.DoubleOptionalFunc(io, &x.ClientPredictedVehicle, io.ActorUniqueID)
	io.Vec2(&x.AnalogMoveVector)
	io.Vec3(&x.CameraOrientation)
	io.Vec2(&x.RawMoveVector)
}

// ID returns the protocol ID for PlayerAuthInput.
func (*PlayerAuthInput) ID() uint32 { return IDPlayerAuthInput }
