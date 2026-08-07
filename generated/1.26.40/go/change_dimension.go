// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "github.com/go-gl/mathgl/mgl32"

type ChangeDimension struct {
	DimensionID     DimensionType
	Position        mgl32.Vec3
	Respawn         bool
	LoadingScreenId Optional[uint32]
}

// Marshal reads or writes ChangeDimension using its canonical wire layout.
func (x *ChangeDimension) Marshal(io IO) {
	x.DimensionID.Marshal(io)
	io.Vec3(&x.Position)
	io.Bool(&x.Respawn)
	io.Bool(&x.LoadingScreenId.set)
	if x.LoadingScreenId.set {
		io.Uint32(&x.LoadingScreenId.val)
	} else if io.Reading() {
		var zero uint32
		x.LoadingScreenId.val = zero
	}
}
