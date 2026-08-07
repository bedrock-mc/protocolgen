// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.40/go/protocol"

	"github.com/go-gl/mathgl/mgl32"
)

type CorrectPlayerMovePrediction struct {
	PredictionType         protocol.RewindType
	Pos                    mgl32.Vec3
	PosDelta               mgl32.Vec3
	Rotation               mgl32.Vec2
	VehicleAngularVelocity protocol.Optional[float32]
	OnGround               bool
	Tick                   uint64
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
