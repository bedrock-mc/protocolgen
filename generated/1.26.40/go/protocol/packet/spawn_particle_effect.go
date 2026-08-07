// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.40/go/protocol"

	"github.com/go-gl/mathgl/mgl32"
)

type SpawnParticleEffect struct {
	DimensionID     uint8
	ActorID         int64
	Position        mgl32.Vec3
	EffectName      string
	MoLangVariables protocol.Optional[string]
}

// Marshal reads or writes SpawnParticleEffect using its canonical wire layout.
func (x *SpawnParticleEffect) Marshal(io protocol.IO) {
	io.Uint8(&x.DimensionID)
	io.ActorUniqueID(&x.ActorID)
	io.Vec3(&x.Position)
	io.String(&x.EffectName)
	protocol.OptionalFunc(io, &x.MoLangVariables, io.String)
}

// ID returns the protocol ID for SpawnParticleEffect.
func (*SpawnParticleEffect) ID() uint32 { return IDSpawnParticleEffect }
