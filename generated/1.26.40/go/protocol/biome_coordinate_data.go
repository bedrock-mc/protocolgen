// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type BiomeCoordinateData struct {
	MinValueType int32
	MinValue     uint16
	MaxValueType int32
	MaxValue     uint16
	GridOffset   uint32
	GridStepSize uint32
	Distribution RandomDistributionType
}

// Marshal reads or writes BiomeCoordinateData using its canonical wire layout.
func (x *BiomeCoordinateData) Marshal(io IO) {
	io.Varint32(&x.MinValueType)
	io.Uint16(&x.MinValue)
	io.Varint32(&x.MaxValueType)
	io.Uint16(&x.MaxValue)
	io.Uint32(&x.GridOffset)
	io.Uint32(&x.GridStepSize)
	IntegerFunc(&x.Distribution, io.Varint32)
}
