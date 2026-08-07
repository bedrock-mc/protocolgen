// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "github.com/go-gl/mathgl/mgl32"

type Interact struct {
	Action          InteractAction
	TargetRuntimeID ActorRuntimeID
	Position        Optional[mgl32.Vec3]
}

// Marshal reads or writes Interact using its canonical wire layout.
func (x *Interact) Marshal(io IO) {
	enumValue1 := uint8(x.Action)
	io.Uint8(&enumValue1)
	x.Action = InteractAction(enumValue1)
	switch int64(enumValue1) {
	case 0, 3, 4, 5, 6:
	default:
		io.InvalidValue(enumValue1, "unknown enum value")
	}
	x.TargetRuntimeID.Marshal(io)
	io.Bool(&x.Position.set)
	if x.Position.set {
		io.Vec3(&x.Position.val)
	} else if io.Reading() {
		var zero mgl32.Vec3
		x.Position.val = zero
	}
}
