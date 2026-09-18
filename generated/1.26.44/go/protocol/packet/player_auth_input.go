// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.44/go/protocol"

	"github.com/go-gl/mathgl/mgl32"
)

// PlayerAuthInput is sent by the client to allow for server authoritative movement. It is used to synchronise
// the player input with the position server-side. The client sends this packet when the
// ServerAuthoritativeMovementMode field in the StartGame packet is set to true, instead of the MovePlayer
// packet. The client will send this packet once every tick.
type PlayerAuthInput struct {
	PlayerRotation mgl32.Vec2
	// Position holds the position that the player reports it has.
	Position mgl32.Vec3
	// MoveVector is a Vec2 that specifies the direction in which the player moved, as a combination of X/Z values
	// which are created using the WASD/controller stick state.
	MoveVector mgl32.Vec2
	// Pitch and Yaw hold the rotation that the player reports it has.
	PlayerHeadRotation float32
	// InputData is the set of input flags that together specify the way the player moved last tick. It holds the
	// flags above.
	InputData protocol.Optional[[]protocol.InputData]
	// InputMode specifies the way that the client inputs data to the screen. It is one of the constants that may
	// be found above.
	InputMode protocol.InputMode
	// PlayMode specifies the way that the player is playing. The values it holds, which are rather random, may be
	// found above.
	PlayMode protocol.ClientPlayMode
	// InteractionModel is a constant representing the interaction model the player is using. It is one of the
	// constants that may be found above.
	NewInteractionModel protocol.NewInteractionModel
	InteractRotation    mgl32.Vec2
	ClientTick          uint64
	// Delta was the delta between the old and the new position. There isn't any practical use for this field as
	// it can be calculated by the server itself.
	PosDelta mgl32.Vec3
	// ItemInteractionData is the transaction data if the InputData includes an item interaction.
	ItemUseTransaction protocol.Optional[protocol.PackedItemUseLegacyInventoryTransaction]
	// ItemStackRequest is sent by the client to change an item in their inventory.
	ItemStackRequest protocol.Optional[protocol.ItemStackRequestData]
	// BlockActions is a slice of block actions that the client has interacted with.
	PlayerBlockActions protocol.Optional[[]protocol.PlayerBlockActionData]
	// VehicleRotation is the rotation of the vehicle that the player is in, if any.
	VehicleRotation protocol.Optional[mgl32.Vec2]
	// ClientPredictedVehicle is the unique ID of the vehicle that the client predicts the player to be in.
	ClientPredictedVehicle protocol.Optional[int64]
	// AnalogueMoveVector is a Vec2 that specifies the direction in which the player moved, as a combination of
	// X/Z values which are created using an analogue input.
	AnalogMoveVector mgl32.Vec2
	// CameraOrientation is the vector that represents the camera's forward direction which can be used to
	// transform movement to be camera relative.
	CameraOrientation mgl32.Vec3
	// RawMoveVector is the value of MoveVector before it is affected by input permissions, sneaking/fly speeds
	// and isn't normalised for analogue inputs.
	RawMoveVector mgl32.Vec2
}

// ID ...
func (*PlayerAuthInput) ID() uint32 {
	return IDPlayerAuthInput
}

func (pk *PlayerAuthInput) Marshal(io protocol.IO) {
	io.Vec2(&pk.PlayerRotation)
	io.Vec3(&pk.Position)
	io.Vec2(&pk.MoveVector)
	io.Float32(&pk.PlayerHeadRotation)
	protocol.OptionalFunc(io, &pk.InputData, func(value *[]protocol.InputData) {
		protocol.Slice(io, value)
	})
	pk.InputMode.Marshal(io)
	pk.PlayMode.Marshal(io)
	pk.NewInteractionModel.Marshal(io)
	io.Vec2(&pk.InteractRotation)
	io.PlayerInputTick(&pk.ClientTick)
	io.Vec3(&pk.PosDelta)
	protocol.DoubleOptionalFunc(io, &pk.ItemUseTransaction, func(value *protocol.PackedItemUseLegacyInventoryTransaction) {
		value.Marshal(io)
	})
	protocol.DoubleOptionalFunc(io, &pk.ItemStackRequest, func(value *protocol.ItemStackRequestData) {
		value.Marshal(io)
	})
	protocol.DoubleOptionalFunc(io, &pk.PlayerBlockActions, func(value *[]protocol.PlayerBlockActionData) {
		protocol.SliceLimits(io, value, 0, 100)
	})
	protocol.DoubleOptionalFunc(io, &pk.VehicleRotation, io.Vec2)
	protocol.DoubleOptionalFunc(io, &pk.ClientPredictedVehicle, io.ActorUniqueID)
	io.Vec2(&pk.AnalogMoveVector)
	io.Vec3(&pk.CameraOrientation)
	io.Vec2(&pk.RawMoveVector)
}
