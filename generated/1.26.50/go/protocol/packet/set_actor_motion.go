// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.50/go/protocol"

	"github.com/go-gl/mathgl/mgl32"
)

// SetActorMotion is sent by the server to change the client-side velocity of an entity. It is
// usually used in combination with server-side movement calculation.
type SetActorMotion struct {
	TargetRuntimeID uint64
	Motion          mgl32.Vec3
	// Tick is the server tick at which the packet was sent. It is used in relation to
	// CorrectPlayerMovePrediction.
	Tick uint64
}

// Marshal reads or writes SetActorMotion using its canonical wire layout.
func (x *SetActorMotion) Marshal(io protocol.IO) {
	io.ActorRuntimeID(&x.TargetRuntimeID)
	io.Vec3(&x.Motion)
	io.PlayerInputTick(&x.Tick)
}

// ID returns the protocol ID for SetActorMotion.
func (*SetActorMotion) ID() uint32 { return IDSetActorMotion }
