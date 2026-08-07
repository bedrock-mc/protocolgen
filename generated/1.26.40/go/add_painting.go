// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "github.com/go-gl/mathgl/mgl32"

type AddPainting struct {
	TargetActorID   ActorUniqueID
	TargetRuntimeID ActorRuntimeID
	Position        mgl32.Vec3
	Direction       int32
	Motif           string
}

// Marshal reads or writes AddPainting using its canonical wire layout.
func (x *AddPainting) Marshal(io IO) {
	x.TargetActorID.Marshal(io)
	x.TargetRuntimeID.Marshal(io)
	io.Vec3(&x.Position)
	io.Varint32(&x.Direction)
	io.String(&x.Motif)
}
