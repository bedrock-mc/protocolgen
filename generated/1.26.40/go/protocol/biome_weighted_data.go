// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type BiomeWeightedData struct {
	BiomeIdentifier uint16
	Weight          uint32
}

// Marshal reads or writes BiomeWeightedData using its canonical wire layout.
func (x *BiomeWeightedData) Marshal(io IO) {
	io.Uint16(&x.BiomeIdentifier)
	io.Uint32(&x.Weight)
}
