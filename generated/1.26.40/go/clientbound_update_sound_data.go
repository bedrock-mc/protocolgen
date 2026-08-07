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
	io.Bool(&x.Stop.set)
	if x.Stop.set {
		marshalSoundDataEvent(io, &x.Stop.val)
	} else if io.Reading() {
		var zero SoundDataEvent
		x.Stop.val = zero
	}
	io.Bool(&x.SetVolume.set)
	if x.SetVolume.set {
		marshalSoundDataEvent(io, &x.SetVolume.val)
	} else if io.Reading() {
		var zero SoundDataEvent
		x.SetVolume.val = zero
	}
	io.Bool(&x.SetPitch.set)
	if x.SetPitch.set {
		marshalSoundDataEvent(io, &x.SetPitch.val)
	} else if io.Reading() {
		var zero SoundDataEvent
		x.SetPitch.val = zero
	}
	io.Bool(&x.Fade.set)
	if x.Fade.set {
		marshalSoundDataEvent(io, &x.Fade.val)
	} else if io.Reading() {
		var zero SoundDataEvent
		x.Fade.val = zero
	}
	io.Bool(&x.SeekTo.set)
	if x.SeekTo.set {
		marshalSoundDataEvent(io, &x.SeekTo.val)
	} else if io.Reading() {
		var zero SoundDataEvent
		x.SeekTo.val = zero
	}
	io.Bool(&x.Pause.set)
	if x.Pause.set {
		marshalSoundDataEvent(io, &x.Pause.val)
	} else if io.Reading() {
		var zero SoundDataEvent
		x.Pause.val = zero
	}
	io.Bool(&x.Resume.set)
	if x.Resume.set {
		marshalSoundDataEvent(io, &x.Resume.val)
	} else if io.Reading() {
		var zero SoundDataEvent
		x.Resume.val = zero
	}
}
