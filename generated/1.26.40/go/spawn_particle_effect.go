// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "github.com/go-gl/mathgl/mgl32"

type SpawnParticleEffect struct {
	DimensionId     uint8
	ActorId         ActorUniqueID
	Position        mgl32.Vec3
	EffectName      string
	MolangVariables Optional[string]
}

// Marshal reads or writes SpawnParticleEffect using its canonical wire layout.
func (x *SpawnParticleEffect) Marshal(io IO) {
	io.Uint8(&x.DimensionId)
	x.ActorId.Marshal(io)
	io.Vec3(&x.Position)
	io.String(&x.EffectName)
	io.Bool(&x.MolangVariables.set)
	if x.MolangVariables.set {
		io.String(&x.MolangVariables.val)
	} else if io.Reading() {
		var zero string
		x.MolangVariables.val = zero
	}
}
