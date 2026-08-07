// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.40/go/protocol"

	"github.com/go-gl/mathgl/mgl32"
)

type LevelEvent struct {
	EventId  int32
	Position mgl32.Vec3
	Data     int32
}

// Marshal reads or writes LevelEvent using its canonical wire layout.
func (x *LevelEvent) Marshal(io protocol.IO) {
	io.Varint32(&x.EventId)
	io.Vec3(&x.Position)
	io.Varint32(&x.Data)
}

// ID returns the protocol ID for LevelEvent.
func (*LevelEvent) ID() uint32 { return IDLevelEvent }
