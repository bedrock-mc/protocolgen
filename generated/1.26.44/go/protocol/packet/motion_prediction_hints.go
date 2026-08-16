// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.44/go/protocol"

	"github.com/go-gl/mathgl/mgl32"
)

// MotionPredictionHints is sent by the server to the client. There is a predictive movement
// component for entities. This packet fills the "history" of that component and entity movement is
// computed based on the points. Vanilla sends this packet instead of the SetActorMotion packet when
// 'spatial optimisations' are enabled.
type MotionPredictionHints struct {
	MRuntimeID uint64
	MMotion    mgl32.Vec3
	MOnGround  bool
}

// Marshal reads or writes MotionPredictionHints using its canonical wire layout.
func (x *MotionPredictionHints) Marshal(io protocol.IO) {
	io.ActorRuntimeID(&x.MRuntimeID)
	io.Vec3(&x.MMotion)
	io.Bool(&x.MOnGround)
}

// ID returns the protocol ID for MotionPredictionHints.
func (*MotionPredictionHints) ID() uint32 { return IDMotionPredictionHints }
