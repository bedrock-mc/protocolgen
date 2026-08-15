// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.40/go/protocol"

	"github.com/go-gl/mathgl/mgl32"
)

// PlayerAuthInput is sent by the client to allow for server authoritative movement. It is used to
// synchronise the player input with the position server-side. The client sends this packet when the
// ServerAuthoritativeMovementMode field in the StartGame packet is set to true, instead of the
// MovePlayer packet. The client will send this packet once every tick.
type PlayerAuthInput struct {
	PlayerRotation mgl32.Vec2
	// Position holds the position that the player reports it has.
	Position mgl32.Vec3
	// MoveVector is a Vec2 that specifies the direction in which the player moved, as a combination of
	// X/Z values which are created using the WASD/controller stick state.
	MoveVector         mgl32.Vec2
	PlayerHeadRotation float32
	// InputData is the set of input flags that together specify the way the player moved last tick. It
	// holds the flags above.
	InputData protocol.Optional[[]protocol.InputData]
	// InputMode specifies the way that the client inputs data to the screen. It is one of the constants
	// that may be found above.
	InputMode protocol.InputMode
	// PlayMode specifies the way that the player is playing. The values it holds, which are rather
	// random, may be found above.
	PlayMode            protocol.ClientPlayMode
	NewInteractionModel protocol.NewInteractionModel
	InteractRotation    mgl32.Vec2
	ClientTick          uint64
	PosDelta            mgl32.Vec3
	ItemUseTransaction  protocol.Optional[protocol.PackedItemUseLegacyInventoryTransaction]
	// ItemStackRequest is sent by the client to change an item in their inventory.
	ItemStackRequest   protocol.Optional[protocol.ItemStackRequestData]
	PlayerBlockActions protocol.Optional[[]protocol.PlayerBlockActionData]
	// VehicleRotation is the rotation of the vehicle that the player is in, if any.
	VehicleRotation protocol.Optional[mgl32.Vec2]
	// ClientPredictedVehicle is the unique ID of the vehicle that the client predicts the player to be
	// in.
	ClientPredictedVehicle protocol.Optional[int64]
	AnalogMoveVector       mgl32.Vec2
	// CameraOrientation is the vector that represents the camera's forward direction which can be used
	// to transform movement to be camera relative.
	CameraOrientation mgl32.Vec3
	// RawMoveVector is the value of MoveVector before it is affected by input permissions, sneaking/fly
	// speeds and isn't normalised for analogue inputs.
	RawMoveVector mgl32.Vec2
}

// Marshal reads or writes PlayerAuthInput using its canonical wire layout.
func (x *PlayerAuthInput) Marshal(io protocol.IO) {
	io.Vec2(&x.PlayerRotation)
	io.Vec3(&x.Position)
	io.Vec2(&x.MoveVector)
	io.Float32(&x.PlayerHeadRotation)
	protocol.OptionalFunc(io, &x.InputData, func(value *[]protocol.InputData) {
		protocol.FuncSlice(io, value, io.Varuint32, func(value *protocol.InputData) {
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
	protocol.DoubleOptionalFunc(io, &x.ItemStackRequest, func(value *protocol.ItemStackRequestData) {
		value.Marshal(io)
	})
	protocol.DoubleOptionalFunc(io, &x.PlayerBlockActions, func(value *[]protocol.PlayerBlockActionData) {
		protocol.SliceLimits(io, value, 0, 100)
	})
	protocol.DoubleOptionalFunc(io, &x.VehicleRotation, io.Vec2)
	protocol.DoubleOptionalFunc(io, &x.ClientPredictedVehicle, io.ActorUniqueID)
	io.Vec2(&x.AnalogMoveVector)
	io.Vec3(&x.CameraOrientation)
	io.Vec2(&x.RawMoveVector)
}

// ID returns the protocol ID for PlayerAuthInput.
func (*PlayerAuthInput) ID() uint32 { return IDPlayerAuthInput }
