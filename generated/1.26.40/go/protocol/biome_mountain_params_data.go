// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type BiomeMountainParamsData struct {
	SteepBlock      uint32
	NorthSlopes     bool
	SouthSlopes     bool
	WestSlopes      bool
	EastSlopes      bool
	TopSlideEnabled bool
}

// Marshal reads or writes BiomeMountainParamsData using its canonical wire layout.
func (x *BiomeMountainParamsData) Marshal(io IO) {
	io.Uint32(&x.SteepBlock)
	io.Bool(&x.NorthSlopes)
	io.Bool(&x.SouthSlopes)
	io.Bool(&x.WestSlopes)
	io.Bool(&x.EastSlopes)
	io.Bool(&x.TopSlideEnabled)
}
