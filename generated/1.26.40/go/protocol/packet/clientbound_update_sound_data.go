// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type ClientboundUpdateSoundData struct {
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
