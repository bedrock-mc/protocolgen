// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.40/go/protocol"

	"github.com/go-gl/mathgl/mgl32"
)

type Interact struct {
	Action          protocol.InteractAction
	TargetRuntimeID uint64
	Position        protocol.Optional[mgl32.Vec3]
}

// Marshal reads or writes Interact using its canonical wire layout.
func (x *Interact) Marshal(io protocol.IO) {
	protocol.IntegerFunc(&x.Action, io.Uint8)
	io.ActorRuntimeID(&x.TargetRuntimeID)
	protocol.OptionalFunc(io, &x.Position, io.Vec3)
}

// ID returns the protocol ID for Interact.
func (*Interact) ID() uint32 { return IDInteract }
