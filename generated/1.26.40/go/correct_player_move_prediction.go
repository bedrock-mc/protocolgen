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
