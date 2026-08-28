// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.50/go/protocol"

// ClientboundUpdateSoundData is sent by the server to update a sound that is currently playing,
// identified by the handle that the server sent in the PlaySound packet that started it. Each
// optional field is a Cereal union slot that may hold any SoundDataUpdate variant; its name does
// not constrain the variant on the wire.
type ClientboundUpdateSoundData struct {
	// ServerSoundHandle is the server-side handle of the sound to update.
	ServerSoundHandle protocol.ServerSoundHandle
	Stop              protocol.Optional[protocol.SoundDataEvent]
	SetVolume         protocol.Optional[protocol.SoundDataEvent]
	SetPitch          protocol.Optional[protocol.SoundDataEvent]
	Fade              protocol.Optional[protocol.SoundDataEvent]
	SeekTo            protocol.Optional[protocol.SoundDataEvent]
	Pause             protocol.Optional[protocol.SoundDataEvent]
	Resume            protocol.Optional[protocol.SoundDataEvent]
}

// Marshal reads or writes ClientboundUpdateSoundData using its canonical wire layout.
func (x *ClientboundUpdateSoundData) Marshal(io protocol.IO) {
	x.ServerSoundHandle.Marshal(io)
	protocol.OptionalFunc(io, &x.Stop, func(value *protocol.SoundDataEvent) {
		protocol.MarshalSoundDataEvent(io, value)
	})
	protocol.OptionalFunc(io, &x.SetVolume, func(value *protocol.SoundDataEvent) {
		protocol.MarshalSoundDataEvent(io, value)
	})
	protocol.OptionalFunc(io, &x.SetPitch, func(value *protocol.SoundDataEvent) {
		protocol.MarshalSoundDataEvent(io, value)
	})
	protocol.OptionalFunc(io, &x.Fade, func(value *protocol.SoundDataEvent) {
		protocol.MarshalSoundDataEvent(io, value)
	})
	protocol.OptionalFunc(io, &x.SeekTo, func(value *protocol.SoundDataEvent) {
		protocol.MarshalSoundDataEvent(io, value)
	})
	protocol.OptionalFunc(io, &x.Pause, func(value *protocol.SoundDataEvent) {
		protocol.MarshalSoundDataEvent(io, value)
	})
	protocol.OptionalFunc(io, &x.Resume, func(value *protocol.SoundDataEvent) {
		protocol.MarshalSoundDataEvent(io, value)
	})
}

// ID returns the protocol ID for ClientboundUpdateSoundData.
func (*ClientboundUpdateSoundData) ID() uint32 { return IDClientboundUpdateSoundData }
