// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.40/go/protocol"

	"github.com/go-gl/mathgl/mgl32"
)

type AddPainting struct {
	TargetActorID   int64
	TargetRuntimeID uint64
	Position        mgl32.Vec3
	Direction       int32
	Motif           string
}

// Marshal reads or writes AddPainting using its canonical wire layout.
func (x *AddPainting) Marshal(io protocol.IO) {
	io.ActorUniqueID(&x.TargetActorID)
	io.ActorRuntimeID(&x.TargetRuntimeID)
	io.Vec3(&x.Position)
	io.Varint32(&x.Direction)
	io.String(&x.Motif)
}

// ID returns the protocol ID for AddPainting.
func (*AddPainting) ID() uint32 { return IDAddPainting }
