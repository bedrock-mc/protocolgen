// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type EASAttributeLayerSettings struct {
	Priority          int32
	Weight            float32
	Enabled           bool
	TransitionsPaused bool
}

// Marshal reads or writes EASAttributeLayerSettings using its canonical wire layout.
func (x *EASAttributeLayerSettings) Marshal(io IO) {
	io.Int32(&x.Priority)
	io.Float32(&x.Weight)
	io.Bool(&x.Enabled)
	io.Bool(&x.TransitionsPaused)
}
