// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

// PlaySound is sent by the server to play a sound to the client. Some of the sounds may only be
// started using this packet and must be stopped using the StopSound packet.
type PlaySound struct {
	// Name is the name of the sound to play.
	Name string
	// Position is the position at which the sound was played. Some sounds do not depend on a position,
	// which will then ignore it, but most of them will play with the direction based on the position
	// compared to the player's position.
	Position protocol.BlockPos
	// Volume is the relative volume of the sound to play. It will be less loud for the player if it is
	// farther away from the position of the sound.
	Volume float32
	// Pitch is the pitch of the sound to play. Some sounds completely ignore this field, whereas others
	// use it to specify the pitch as the field is intended.
	Pitch float32
	// LoopCount is the number of times to loop the sound before stopping. -1 means no looping at all.
	LoopCount int32
	// ServerSoundHandle is an optional sound handle ID. It is currently unknown what this is for, and
	// is not required to be set by servers.
	ServerSoundHandle protocol.Optional[protocol.ServerSoundHandle]
}

// Marshal reads or writes PlaySound using its canonical wire layout.
func (x *PlaySound) Marshal(io protocol.IO) {
	io.String(&x.Name)
	x.Position.Marshal(io)
	io.Float32(&x.Volume)
	io.Float32(&x.Pitch)
	io.Varint32(&x.LoopCount)
	protocol.OptionalFunc(io, &x.ServerSoundHandle, func(value *protocol.ServerSoundHandle) {
		value.Marshal(io)
	})
}

// ID returns the protocol ID for PlaySound.
func (*PlaySound) ID() uint32 { return IDPlaySound }
