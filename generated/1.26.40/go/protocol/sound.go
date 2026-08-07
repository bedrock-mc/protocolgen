// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type SoundDataEvent interface {
	isSoundDataEvent()
}

// MarshalSoundDataEvent reads or writes the SoundDataEvent union using its canonical wire layout.
func MarshalSoundDataEvent(io IO, x *SoundDataEvent) {
	UnionFunc(io,
		func() {
			var tag uint32
			io.Varuint32(&tag)
			switch int64(tag) {
			case 0:
				value := new(SoundDataEventStop)
				value.Marshal(io)
				*x = value
			case 1:
				value := new(SoundDataEventSetVolume)
				value.Marshal(io)
				*x = value
			case 2:
				value := new(SoundDataEventSetPitch)
				value.Marshal(io)
				*x = value
			case 3:
				value := new(SoundDataEventFade)
				value.Marshal(io)
				*x = value
			case 4:
				value := new(SoundDataEventSeekTo)
				value.Marshal(io)
				*x = value
			case 5:
				value := new(SoundDataEventPause)
				value.Marshal(io)
				*x = value
			case 6:
				value := new(SoundDataEventResume)
				value.Marshal(io)
				*x = value
			default:
				io.InvalidValue(tag, "unknown union tag")
			}
		},
		func() {
			switch value := (*x).(type) {
			case *SoundDataEventStop:
				tag := uint32(0)
				io.Varuint32(&tag)
				value.Marshal(io)
			case *SoundDataEventSetVolume:
				tag := uint32(1)
				io.Varuint32(&tag)
				value.Marshal(io)
			case *SoundDataEventSetPitch:
				tag := uint32(2)
				io.Varuint32(&tag)
				value.Marshal(io)
			case *SoundDataEventFade:
				tag := uint32(3)
				io.Varuint32(&tag)
				value.Marshal(io)
			case *SoundDataEventSeekTo:
				tag := uint32(4)
				io.Varuint32(&tag)
				value.Marshal(io)
			case *SoundDataEventPause:
				tag := uint32(5)
				io.Varuint32(&tag)
				value.Marshal(io)
			case *SoundDataEventResume:
				tag := uint32(6)
				io.Varuint32(&tag)
				value.Marshal(io)
			default:
				io.InvalidValue(*x, "unknown union value")
			}
		},
	)
}

type SoundDataEventFade struct {
	Duration     float32
	TargetVolume float32
}

func (*SoundDataEventFade) isSoundDataEvent() {}

// Marshal reads or writes SoundDataEventFade using its canonical wire layout.
func (x *SoundDataEventFade) Marshal(io IO) {
	io.Float32(&x.Duration)
	io.Float32(&x.TargetVolume)
}

type SoundDataEventPause struct {
}

func (*SoundDataEventPause) isSoundDataEvent() {}

// Marshal reads or writes SoundDataEventPause using its canonical wire layout.
func (x *SoundDataEventPause) Marshal(io IO) {
}

type SoundDataEventResume struct {
}

func (*SoundDataEventResume) isSoundDataEvent() {}

// Marshal reads or writes SoundDataEventResume using its canonical wire layout.
func (x *SoundDataEventResume) Marshal(io IO) {
}

type SoundDataEventSeekTo struct {
	Seconds float32
}

func (*SoundDataEventSeekTo) isSoundDataEvent() {}

// Marshal reads or writes SoundDataEventSeekTo using its canonical wire layout.
func (x *SoundDataEventSeekTo) Marshal(io IO) {
	io.Float32(&x.Seconds)
}

type SoundDataEventSetPitch struct {
	Pitch float32
}

func (*SoundDataEventSetPitch) isSoundDataEvent() {}

// Marshal reads or writes SoundDataEventSetPitch using its canonical wire layout.
func (x *SoundDataEventSetPitch) Marshal(io IO) {
	io.Float32(&x.Pitch)
}

type SoundDataEventSetVolume struct {
	Volume float32
}

func (*SoundDataEventSetVolume) isSoundDataEvent() {}

// Marshal reads or writes SoundDataEventSetVolume using its canonical wire layout.
func (x *SoundDataEventSetVolume) Marshal(io IO) {
	io.Float32(&x.Volume)
}

type SoundDataEventStop struct {
}

func (*SoundDataEventStop) isSoundDataEvent() {}

// Marshal reads or writes SoundDataEventStop using its canonical wire layout.
func (x *SoundDataEventStop) Marshal(io IO) {
}
