// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "github.com/go-gl/mathgl/mgl32"

type ActorEvent struct {
	TargetRuntimeID ActorRuntimeID
	EventID         ActorEventType
	Data            int32
	FireAtPosition  Optional[mgl32.Vec3]
}

// Marshal reads or writes ActorEvent using its canonical wire layout.
func (x *ActorEvent) Marshal(io IO) {
	x.TargetRuntimeID.Marshal(io)
	enumValue1 := uint8(x.EventID)
	io.Uint8(&enumValue1)
	x.EventID = ActorEventType(enumValue1)
	switch int64(enumValue1) {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 57, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 76, 77, 78, 79, 80, 81:
	default:
		io.InvalidValue(enumValue1, "unknown enum value")
	}
	io.Varint32(&x.Data)
	io.Bool(&x.FireAtPosition.set)
	if x.FireAtPosition.set {
		io.Vec3(&x.FireAtPosition.val)
	} else if io.Reading() {
		var zero mgl32.Vec3
		x.FireAtPosition.val = zero
	}
}
