// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.40/go/protocol"

	"github.com/go-gl/mathgl/mgl32"
)

type Respawn struct {
	Position        mgl32.Vec3
	State           protocol.PlayerRespawnState
	PlayerRuntimeId uint64
}

// Marshal reads or writes Respawn using its canonical wire layout.
func (x *Respawn) Marshal(io protocol.IO) {
	io.Vec3(&x.Position)
	protocol.IntegerFunc(&x.State, io.Uint8)
	io.ActorRuntimeID(&x.PlayerRuntimeId)
}
