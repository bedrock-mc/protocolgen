// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "github.com/go-gl/mathgl/mgl32"

type ActorEvent struct {
	TargetRuntimeID ActorRuntimeID
	EventID         ActorEventType
	Data            int32
	FireAtPosition  Optional[mgl32.Vec3]
}
