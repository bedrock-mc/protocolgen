// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type BiomeWeightedTemperatureData struct {
	Temperature int32
	Weight      uint32
}

// Marshal reads or writes BiomeWeightedTemperatureData using its canonical wire layout.
func (x *BiomeWeightedTemperatureData) Marshal(io IO) {
	io.Varint32(&x.Temperature)
	io.Uint32(&x.Weight)
}
