// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.44/go/protocol"

	"github.com/go-gl/mathgl/mgl32"
)

// SpawnParticleEffect is sent by the server to spawn a particle effect client-side. Unlike other
// packets that result in the appearing of particles, this packet can show particles that are not
// hardcoded in the client. They can be added and changed through behaviour packs to implement
// custom particles.
type SpawnParticleEffect struct {
	DimensionID uint8
	ActorID     int64
	// Position is the position that the particle should be spawned at. If the position is too far away
	// from the player, it will not show up. If EntityUniqueID is not -1, the position will be relative
	// to the position of the entity.
	Position   mgl32.Vec3
	EffectName string
	// MoLangVariables is an encoded JSON map of MoLang variables that may be applicable to the particle
	// spawn. This can just be left empty in most cases.
	MoLangVariables protocol.Optional[string]
}

// Marshal reads or writes SpawnParticleEffect using its canonical wire layout.
func (x *SpawnParticleEffect) Marshal(io protocol.IO) {
	io.Uint8(&x.DimensionID)
	protocol.Minimum(io, &x.DimensionID, 0)
	protocol.Maximum(io, &x.DimensionID, 255)
	io.ActorUniqueID(&x.ActorID)
	io.Vec3(&x.Position)
	io.String(&x.EffectName)
	protocol.OptionalFunc(io, &x.MoLangVariables, io.String)
}

// ID returns the protocol ID for SpawnParticleEffect.
func (*SpawnParticleEffect) ID() uint32 { return IDSpawnParticleEffect }
