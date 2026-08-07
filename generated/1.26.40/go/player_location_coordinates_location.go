// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "github.com/go-gl/mathgl/mgl32"

type PlayerLocationCoordinatesLocation struct {
	PacketType PlayerLocationType
	Position   mgl32.Vec3
}

func (PlayerLocationCoordinatesLocation) isPlayerLocationLocation() {}

// Marshal reads or writes PlayerLocationCoordinatesLocation using its canonical wire layout.
func (x *PlayerLocationCoordinatesLocation) Marshal(io IO) {
	IntegerFunc(&x.PacketType, io.Varint32)
	io.Vec3(&x.Position)
}
