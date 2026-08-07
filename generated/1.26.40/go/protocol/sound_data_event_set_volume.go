// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type SoundDataEventSetVolume struct {
	Volume float32
}

func (*SoundDataEventSetVolume) isSoundDataEvent() {}

// Marshal reads or writes SoundDataEventSetVolume using its canonical wire layout.
func (x *SoundDataEventSetVolume) Marshal(io IO) {
	io.Float32(&x.Volume)
}
