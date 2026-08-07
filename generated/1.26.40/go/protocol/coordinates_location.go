// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

import "github.com/go-gl/mathgl/mgl32"

type CoordinatesLocation struct {
	PacketType PlayerLocationType
	Position   mgl32.Vec3
}

func (*CoordinatesLocation) isPlayerLocationData() {}

// Marshal reads or writes CoordinatesLocation using its canonical wire layout.
func (x *CoordinatesLocation) Marshal(io IO) {
	IntegerFunc(&x.PacketType, io.Varint32)
	io.Vec3(&x.Position)
}
