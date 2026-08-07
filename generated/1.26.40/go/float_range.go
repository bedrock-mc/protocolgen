// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type FloatRange struct {
	Min float32
	Max float32
}

// Marshal reads or writes FloatRange using its canonical wire layout.
func (x *FloatRange) Marshal(io IO) {
	io.Float32(&x.Min)
	io.Float32(&x.Max)
}
