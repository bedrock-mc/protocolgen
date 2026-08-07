// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type SoundDataEventSetPitch struct {
	Pitch float32
}

func (SoundDataEventSetPitch) isSoundDataEvent() {}

// Marshal reads or writes SoundDataEventSetPitch using its canonical wire layout.
func (x *SoundDataEventSetPitch) Marshal(io IO) {
	io.Float32(&x.Pitch)
}
