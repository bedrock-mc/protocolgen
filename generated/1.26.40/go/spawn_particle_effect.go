// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "github.com/go-gl/mathgl/mgl32"

type SpawnParticleEffect struct {
	DimensionId     uint8
	ActorId         int64
	Position        mgl32.Vec3
	EffectName      string
	MolangVariables Optional[string]
}

// Marshal reads or writes SpawnParticleEffect using its canonical wire layout.
func (x *SpawnParticleEffect) Marshal(io IO) {
	io.Uint8(&x.DimensionId)
	io.ActorUniqueID(&x.ActorId)
	io.Vec3(&x.Position)
	io.String(&x.EffectName)
	OptionalFunc(io, &x.MolangVariables, io.String)
}
