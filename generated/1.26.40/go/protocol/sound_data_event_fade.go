// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

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
