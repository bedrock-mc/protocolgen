// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.40/go/protocol"

	"github.com/go-gl/mathgl/mgl32"
)

type SpawnExperienceOrb struct {
	Position mgl32.Vec3
	XPValue  int32
}

// Marshal reads or writes SpawnExperienceOrb using its canonical wire layout.
func (x *SpawnExperienceOrb) Marshal(io protocol.IO) {
	io.Vec3(&x.Position)
	io.Varint32(&x.XPValue)
}

// ID returns the protocol ID for SpawnExperienceOrb.
func (*SpawnExperienceOrb) ID() uint32 { return IDSpawnExperienceOrb }
