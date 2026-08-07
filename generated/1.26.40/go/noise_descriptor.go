// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type NoiseDescriptor struct {
	Name        string
	FirstOctave int32
	Amplitudes  []float32
}

// Marshal reads or writes NoiseDescriptor using its canonical wire layout.
func (x *NoiseDescriptor) Marshal(io IO) {
	io.String(&x.Name)
	io.Int32(&x.FirstOctave)
	FuncSlice(io, &x.Amplitudes, io.Varuint32, io.Float32)
}
