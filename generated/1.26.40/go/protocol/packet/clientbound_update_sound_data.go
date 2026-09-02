// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

// ClientboundUpdateSoundData is sent by the server to update a sound that is currently playing,
// identified by the handle that the server sent in the PlaySound packet that started it. Each
// optional field is a Cereal union slot that may hold any SoundDataUpdate variant; its name does
// not constrain the variant on the wire.
type ClientboundUpdateSoundData struct {
	// ServerSoundHandle is the server-side handle of the sound to update.
	ServerSoundHandle protocol.ServerSoundHandle
	Stop              protocol.SoundDataEvent
	SetVolume         protocol.SoundDataEvent
	SetPitch          protocol.SoundDataEvent
	Fade              protocol.SoundDataEvent
	SeekTo            protocol.SoundDataEvent
	Pause             protocol.SoundDataEvent
	Resume            protocol.SoundDataEvent
}

// Marshal reads or writes ClientboundUpdateSoundData using its canonical wire layout.
func (x *ClientboundUpdateSoundData) Marshal(io protocol.IO) {
	x.ServerSoundHandle.Marshal(io)
	protocol.MarshalSoundDataEvent(io, &x.Stop)
	protocol.MarshalSoundDataEvent(io, &x.SetVolume)
	protocol.MarshalSoundDataEvent(io, &x.SetPitch)
	protocol.MarshalSoundDataEvent(io, &x.Fade)
	protocol.MarshalSoundDataEvent(io, &x.SeekTo)
	protocol.MarshalSoundDataEvent(io, &x.Pause)
	protocol.MarshalSoundDataEvent(io, &x.Resume)
}

// ID returns the protocol ID for ClientboundUpdateSoundData.
func (*ClientboundUpdateSoundData) ID() uint32 { return IDClientboundUpdateSoundData }
