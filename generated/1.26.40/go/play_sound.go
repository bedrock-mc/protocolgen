// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type PlaySound struct {
	Name              string
	Position          BlockPos
	Volume            float32
	Pitch             float32
	LoopCount         int32
	ServerSoundHandle Optional[ServerSoundHandle]
}

// Marshal reads or writes PlaySound using its canonical wire layout.
func (x *PlaySound) Marshal(io IO) {
	io.String(&x.Name)
	x.Position.Marshal(io)
	io.Float32(&x.Volume)
	io.Float32(&x.Pitch)
	io.Varint32(&x.LoopCount)
	io.Bool(&x.ServerSoundHandle.set)
	if x.ServerSoundHandle.set {
		x.ServerSoundHandle.val.Marshal(io)
	} else if io.Reading() {
		var zero ServerSoundHandle
		x.ServerSoundHandle.val = zero
	}
}
