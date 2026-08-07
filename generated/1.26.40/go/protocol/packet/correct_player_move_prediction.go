// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.40/go/protocol"

	"github.com/go-gl/mathgl/mgl32"
)

// CorrectPlayerMovePrediction is sent by the server if and only if
// StartGame.ServerAuthoritativeMovementMode is set to AuthoritativeMovementModeServerWithRewind.
// The packet is used to correct movement at a specific point in time.
type CorrectPlayerMovePrediction struct {
	// PredictionType is the type of prediction that was corrected. It is one of the constants above.
	PredictionType protocol.RewindType
	Pos            mgl32.Vec3
	PosDelta       mgl32.Vec3
	// Rotation is the rotation of the player at the tick written in the field below.
	Rotation mgl32.Vec2
	// VehicleAngularVelocity is the angular velocity of the vehicle that the rider is riding.
	VehicleAngularVelocity protocol.Optional[float32]
	// OnGround specifies if the player was on the ground at the time of the tick below.
	OnGround bool
	// Tick is the tick of the movement which was corrected by this packet.
	Tick uint64
}

// Marshal reads or writes CorrectPlayerMovePrediction using its canonical wire layout.
func (x *CorrectPlayerMovePrediction) Marshal(io protocol.IO) {
	protocol.IntegerFunc(&x.PredictionType, io.Uint8)
	io.Vec3(&x.Pos)
	io.Vec3(&x.PosDelta)
	io.Vec2(&x.Rotation)
	protocol.OptionalFunc(io, &x.VehicleAngularVelocity, io.Float32)
	io.Bool(&x.OnGround)
	io.PlayerInputTick(&x.Tick)
}

// ID returns the protocol ID for CorrectPlayerMovePrediction.
func (*CorrectPlayerMovePrediction) ID() uint32 { return IDCorrectPlayerMovePrediction }
