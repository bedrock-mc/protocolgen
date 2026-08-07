// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "github.com/go-gl/mathgl/mgl32"

type MotionPredictionHints struct {
	MRuntimeId uint64
	MMotion    mgl32.Vec3
	MOnGround  bool
}

// Marshal reads or writes MotionPredictionHints using its canonical wire layout.
func (x *MotionPredictionHints) Marshal(io IO) {
	io.ActorRuntimeID(&x.MRuntimeId)
	io.Vec3(&x.MMotion)
	io.Bool(&x.MOnGround)
}
