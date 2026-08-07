// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.40/go/protocol"

	"github.com/go-gl/mathgl/mgl32"
)

type LevelSoundEvent struct {
	SoundEvent      string
	Position        mgl32.Vec3
	Data            int32
	ActorIdentifier string
	IsBaby          bool
	IsGlobal        bool
	ActorUniqueID   int64
	FireAtPosition  protocol.Optional[mgl32.Vec3]
}

// Marshal reads or writes LevelSoundEvent using its canonical wire layout.
func (x *LevelSoundEvent) Marshal(io protocol.IO) {
	io.String(&x.SoundEvent)
	io.Vec3(&x.Position)
	io.Varint32(&x.Data)
	io.String(&x.ActorIdentifier)
	io.Bool(&x.IsBaby)
	io.Bool(&x.IsGlobal)
	io.Int64(&x.ActorUniqueID)
	protocol.OptionalFunc(io, &x.FireAtPosition, io.Vec3)
}

// ID returns the protocol ID for LevelSoundEvent.
func (*LevelSoundEvent) ID() uint32 { return IDLevelSoundEvent }
