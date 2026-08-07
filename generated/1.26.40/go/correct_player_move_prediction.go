// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "github.com/go-gl/mathgl/mgl32"

type CorrectPlayerMovePrediction struct {
	PredictionType         RewindType
	Pos                    mgl32.Vec3
	PosDelta               mgl32.Vec3
	Rotation               mgl32.Vec2
	VehicleAngularVelocity Optional[float32]
	OnGround               bool
	Tick                   PlayerInputTick
}

// Marshal reads or writes CorrectPlayerMovePrediction using its canonical wire layout.
func (x *CorrectPlayerMovePrediction) Marshal(io IO) {
	enumValue1 := uint8(x.PredictionType)
	io.Uint8(&enumValue1)
	x.PredictionType = RewindType(enumValue1)
	switch int64(enumValue1) {
	case 0, 1:
	default:
		io.InvalidValue(enumValue1, "unknown enum value")
	}
	io.Vec3(&x.Pos)
	io.Vec3(&x.PosDelta)
	io.Vec2(&x.Rotation)
	io.Bool(&x.VehicleAngularVelocity.set)
	if x.VehicleAngularVelocity.set {
		io.Float32(&x.VehicleAngularVelocity.val)
	} else if io.Reading() {
		var zero float32
		x.VehicleAngularVelocity.val = zero
	}
	io.Bool(&x.OnGround)
	x.Tick.Marshal(io)
}
