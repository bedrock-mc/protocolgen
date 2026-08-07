// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type ClientboundUpdateSoundData struct {
	ServerSoundHandle ServerSoundHandle
	Stop              Optional[SoundDataEvent]
	SetVolume         Optional[SoundDataEvent]
	SetPitch          Optional[SoundDataEvent]
	Fade              Optional[SoundDataEvent]
	SeekTo            Optional[SoundDataEvent]
	Pause             Optional[SoundDataEvent]
	Resume            Optional[SoundDataEvent]
}

// Marshal reads or writes ClientboundUpdateSoundData using its canonical wire layout.
func (x *ClientboundUpdateSoundData) Marshal(io IO) {
	x.ServerSoundHandle.Marshal(io)
	OptionalFunc(io, &x.Stop, func(value *SoundDataEvent) {
		marshalSoundDataEvent(io, value)
	})
	OptionalFunc(io, &x.SetVolume, func(value *SoundDataEvent) {
		marshalSoundDataEvent(io, value)
	})
	OptionalFunc(io, &x.SetPitch, func(value *SoundDataEvent) {
		marshalSoundDataEvent(io, value)
	})
	OptionalFunc(io, &x.Fade, func(value *SoundDataEvent) {
		marshalSoundDataEvent(io, value)
	})
	OptionalFunc(io, &x.SeekTo, func(value *SoundDataEvent) {
		marshalSoundDataEvent(io, value)
	})
	OptionalFunc(io, &x.Pause, func(value *SoundDataEvent) {
		marshalSoundDataEvent(io, value)
	})
	OptionalFunc(io, &x.Resume, func(value *SoundDataEvent) {
		marshalSoundDataEvent(io, value)
	})
}
