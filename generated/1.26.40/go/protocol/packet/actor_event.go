// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.40/go/protocol"

	"github.com/go-gl/mathgl/mgl32"
)

type ActorEvent struct {
	TargetRuntimeID uint64
	EventID         protocol.ActorEventType
	Data            int32
	FireAtPosition  protocol.Optional[mgl32.Vec3]
}

// Marshal reads or writes ActorEvent using its canonical wire layout.
func (x *ActorEvent) Marshal(io protocol.IO) {
	io.ActorRuntimeID(&x.TargetRuntimeID)
	protocol.IntegerFunc(&x.EventID, io.Uint8)
	io.Varint32(&x.Data)
	protocol.OptionalFunc(io, &x.FireAtPosition, io.Vec3)
}
