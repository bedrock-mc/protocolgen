// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.40/go/protocol"

	"github.com/go-gl/mathgl/mgl32"
)

type MotionPredictionHints struct {
	MRuntimeId uint64
	MMotion    mgl32.Vec3
	MOnGround  bool
}

// Marshal reads or writes MotionPredictionHints using its canonical wire layout.
func (x *MotionPredictionHints) Marshal(io protocol.IO) {
	io.ActorRuntimeID(&x.MRuntimeId)
	io.Vec3(&x.MMotion)
	io.Bool(&x.MOnGround)
}

// ID returns the protocol ID for MotionPredictionHints.
func (*MotionPredictionHints) ID() uint32 { return IDMotionPredictionHints }
