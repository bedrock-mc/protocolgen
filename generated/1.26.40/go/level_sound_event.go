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
	FireAtPosition  Optional[mgl32.Vec3]
}

// Marshal reads or writes LevelSoundEvent using its canonical wire layout.
func (x *LevelSoundEvent) Marshal(io IO) {
	io.String(&x.SoundEvent)
	io.Vec3(&x.Position)
	io.Varint32(&x.Data)
	io.String(&x.ActorIdentifier)
	io.Bool(&x.IsBaby)
	io.Bool(&x.IsGlobal)
	io.Int64(&x.ActorUniqueId)
	io.Bool(&x.FireAtPosition.set)
	if x.FireAtPosition.set {
		io.Vec3(&x.FireAtPosition.val)
	} else if io.Reading() {
		var zero mgl32.Vec3
		x.FireAtPosition.val = zero
	}
}
