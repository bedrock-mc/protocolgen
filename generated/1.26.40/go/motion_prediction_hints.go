// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "github.com/go-gl/mathgl/mgl32"

type MotionPredictionHints struct {
	MRuntimeId ActorRuntimeID
	MMotion    mgl32.Vec3
	MOnGround  bool
}
