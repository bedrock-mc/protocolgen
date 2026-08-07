// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "github.com/go-gl/mathgl/mgl32"

type LevelSoundEvent struct {
	SoundEvent      string
	Position        mgl32.Vec3
	Data            int32
	ActorIdentifier string
	IsBaby          bool
	IsGlobal        bool
	ActorUniqueId   int64
	FireAtPosition  *mgl32.Vec3
}
