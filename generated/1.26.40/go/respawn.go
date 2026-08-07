// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "github.com/go-gl/mathgl/mgl32"

type Respawn struct {
	Position        mgl32.Vec3
	State           PlayerRespawnState
	PlayerRuntimeId ActorRuntimeID
}
