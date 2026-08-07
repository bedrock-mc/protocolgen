// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "github.com/go-gl/mathgl/mgl32"

type WorldPosition struct {
	Position      mgl32.Vec3
	DimensionType DimensionType
}

// Marshal reads or writes WorldPosition using its canonical wire layout.
func (x *WorldPosition) Marshal(io IO) {
	io.Vec3(&x.Position)
	x.DimensionType.Marshal(io)
}
