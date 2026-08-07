// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "github.com/go-gl/mathgl/mgl32"

type Respawn struct {
	Position        mgl32.Vec3
	State           PlayerRespawnState
	PlayerRuntimeId uint64
}

// Marshal reads or writes Respawn using its canonical wire layout.
func (x *Respawn) Marshal(io IO) {
	io.Vec3(&x.Position)
	IntegerFunc(&x.State, io.Uint8)
	io.ActorRuntimeID(&x.PlayerRuntimeId)
}
