// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type SoundDataEventSeekTo struct {
	Seconds float32
}

func (*SoundDataEventSeekTo) isSoundDataEvent() {}

// Marshal reads or writes SoundDataEventSeekTo using its canonical wire layout.
func (x *SoundDataEventSeekTo) Marshal(io IO) {
	io.Float32(&x.Seconds)
}
