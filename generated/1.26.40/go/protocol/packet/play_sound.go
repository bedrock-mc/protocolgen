// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type PlaySound struct {
	Name              string
	Position          protocol.BlockPos
	Volume            float32
	Pitch             float32
	LoopCount         int32
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
