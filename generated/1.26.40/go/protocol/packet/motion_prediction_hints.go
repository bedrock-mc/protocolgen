// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.40/go/protocol"

	"github.com/go-gl/mathgl/mgl32"
)

// MotionPredictionHints is sent by the server to the client. There is a predictive movement component for
// entities. This packet fills the "history" of that component and entity movement is computed based on the
// points. Vanilla sends this packet instead of the SetActorMotion packet when 'spatial optimisations' are
// enabled.
type MotionPredictionHints struct {
	MRuntimeID uint64
	MMotion    mgl32.Vec3
	MOnGround  bool
}

// ID ...
func (*MotionPredictionHints) ID() uint32 {
	return IDMotionPredictionHints
}

func (pk *MotionPredictionHints) Marshal(io protocol.IO) {
	io.ActorRuntimeID(&pk.MRuntimeID)
	io.Vec3(&pk.MMotion)
	io.Bool(&pk.MOnGround)
}
