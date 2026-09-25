// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type SoundDataEvent interface {
	Marshaler
	tagSoundDataEvent() uint32
}

// MarshalSoundDataEvent reads or writes the SoundDataEvent union using its canonical wire layout.
func MarshalSoundDataEvent(io IO, x *SoundDataEvent) {
	Union(io, x, io.Varuint32, SoundDataEvent.tagSoundDataEvent, func(tag uint32) SoundDataEvent {
		switch tag {
		case 0:
			return new(SoundDataEventStop)
		case 1:
			return new(SoundDataEventSetVolume)
		case 2:
			return new(SoundDataEventSetPitch)
		case 3:
			return new(SoundDataEventFade)
		case 4:
			return new(SoundDataEventSeekTo)
		case 5:
			return new(SoundDataEventPause)
		case 6:
			return new(SoundDataEventResume)
		}
		return nil
	})
}

type SoundDataEventFade struct {
	Duration     float32
	TargetVolume float32
}

func (*SoundDataEventFade) tagSoundDataEvent() uint32 { return 3 }

// Marshal reads or writes SoundDataEventFade using its canonical wire layout.
func (x *SoundDataEventFade) Marshal(io IO) {
	io.Float32(&x.Duration)
	io.Float32(&x.TargetVolume)
}

type SoundDataEventPause struct {
}

func (*SoundDataEventPause) tagSoundDataEvent() uint32 { return 5 }

// Marshal reads or writes SoundDataEventPause using its canonical wire layout.
func (x *SoundDataEventPause) Marshal(io IO) {
}

type SoundDataEventResume struct {
}

func (*SoundDataEventResume) tagSoundDataEvent() uint32 { return 6 }

// Marshal reads or writes SoundDataEventResume using its canonical wire layout.
func (x *SoundDataEventResume) Marshal(io IO) {
}

type SoundDataEventSeekTo struct {
	Seconds float32
}

func (*SoundDataEventSeekTo) tagSoundDataEvent() uint32 { return 4 }

// Marshal reads or writes SoundDataEventSeekTo using its canonical wire layout.
func (x *SoundDataEventSeekTo) Marshal(io IO) {
	io.Float32(&x.Seconds)
}

type SoundDataEventSetPitch struct {
	Pitch float32
}

func (*SoundDataEventSetPitch) tagSoundDataEvent() uint32 { return 2 }

// Marshal reads or writes SoundDataEventSetPitch using its canonical wire layout.
func (x *SoundDataEventSetPitch) Marshal(io IO) {
	io.Float32(&x.Pitch)
}

type SoundDataEventSetVolume struct {
	Volume float32
}

func (*SoundDataEventSetVolume) tagSoundDataEvent() uint32 { return 1 }

// Marshal reads or writes SoundDataEventSetVolume using its canonical wire layout.
func (x *SoundDataEventSetVolume) Marshal(io IO) {
	io.Float32(&x.Volume)
}

type SoundDataEventStop struct {
}

func (*SoundDataEventStop) tagSoundDataEvent() uint32 { return 0 }

// Marshal reads or writes SoundDataEventStop using its canonical wire layout.
func (x *SoundDataEventStop) Marshal(io IO) {
}
