// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "github.com/go-gl/mathgl/mgl32"

type Interact struct {
	Action          InteractAction
	TargetRuntimeID uint64
	Position        Optional[mgl32.Vec3]
}

// Marshal reads or writes Interact using its canonical wire layout.
func (x *Interact) Marshal(io IO) {
	IntegerFunc(&x.Action, io.Uint8)
	io.ActorRuntimeID(&x.TargetRuntimeID)
	OptionalFunc(io, &x.Position, io.Vec3)
}
