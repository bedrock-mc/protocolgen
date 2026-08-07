// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type SerializedNoiseBlockSpecifier struct {
	Noise     string
	Threshold float32
	Range     FloatRange
	Block     uint32
}

// Marshal reads or writes SerializedNoiseBlockSpecifier using its canonical wire layout.
func (x *SerializedNoiseBlockSpecifier) Marshal(io IO) {
	io.String(&x.Noise)
	io.Float32(&x.Threshold)
	x.Range.Marshal(io)
	io.Uint32(&x.Block)
}
