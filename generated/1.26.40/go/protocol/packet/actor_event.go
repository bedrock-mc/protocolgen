// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.40/go/protocol"

	"github.com/go-gl/mathgl/mgl32"
)

// ActorEvent is sent by the server when a particular event happens that has to do with an entity.
// Some of these events are entity-specific, for example a wolf shaking itself dry, but others are
// used for each entity, such as dying.
type ActorEvent struct {
	TargetRuntimeID uint64
	EventID         protocol.ActorEventType
	Data            int32
	// FireAtPosition is the position in the same world at which the event should fire. If this is not
	// present, the position entity will be used instead.
	FireAtPosition protocol.Optional[mgl32.Vec3]
}

// Marshal reads or writes ActorEvent using its canonical wire layout.
func (x *ActorEvent) Marshal(io protocol.IO) {
	io.ActorRuntimeID(&x.TargetRuntimeID)
	protocol.IntegerFunc(&x.EventID, io.Uint8)
	io.Varint32(&x.Data)
	protocol.OptionalFunc(io, &x.FireAtPosition, io.Vec3)
}

// ID returns the protocol ID for ActorEvent.
func (*ActorEvent) ID() uint32 { return IDActorEvent }
