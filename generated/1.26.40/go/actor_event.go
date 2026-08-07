// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "github.com/go-gl/mathgl/mgl32"

type ActorEvent struct {
	TargetRuntimeID uint64
	EventID         ActorEventType
	Data            int32
	FireAtPosition  Optional[mgl32.Vec3]
}

// Marshal reads or writes ActorEvent using its canonical wire layout.
func (x *ActorEvent) Marshal(io IO) {
	io.ActorRuntimeID(&x.TargetRuntimeID)
	IntegerFunc(&x.EventID, io.Uint8)
	io.Varint32(&x.Data)
	OptionalFunc(io, &x.FireAtPosition, io.Vec3)
}
