// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "github.com/go-gl/mathgl/mgl32"

type Respawn struct {
	Position        mgl32.Vec3
	State           PlayerRespawnState
	PlayerRuntimeId ActorRuntimeID
}

// Marshal reads or writes Respawn using its canonical wire layout.
func (x *Respawn) Marshal(io IO) {
	io.Vec3(&x.Position)
	enumValue1 := uint8(x.State)
	io.Uint8(&enumValue1)
	x.State = PlayerRespawnState(enumValue1)
	switch int64(enumValue1) {
	case 0, 1, 2:
	default:
		io.InvalidValue(enumValue1, "unknown enum value")
	}
	x.PlayerRuntimeId.Marshal(io)
}
