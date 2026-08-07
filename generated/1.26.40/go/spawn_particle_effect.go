// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "github.com/go-gl/mathgl/mgl32"

type SpawnParticleEffect struct {
	DimensionId     uint8
	ActorId         ActorUniqueID
	Position        mgl32.Vec3
	EffectName      string
	MolangVariables *string
}
