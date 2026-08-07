// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "github.com/go-gl/mathgl/mgl32"

type LevelEvent struct {
	EventId  int32
	Position mgl32.Vec3
	Data     int32
}

// Marshal reads or writes LevelEvent using its canonical wire layout.
func (x *LevelEvent) Marshal(io IO) {
	io.Varint32(&x.EventId)
	io.Vec3(&x.Position)
	io.Varint32(&x.Data)
}
