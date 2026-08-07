// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type BiomeReplacementData struct {
	ReplacementBiome    uint16
	Dimension           uint16
	TargetBiomes        []uint16
	Amount              float32
	NoiseFrequencyScale float32
	ReplacementIndex    uint32
}

// Marshal reads or writes BiomeReplacementData using its canonical wire layout.
func (x *BiomeReplacementData) Marshal(io IO) {
	io.Uint16(&x.ReplacementBiome)
	io.Uint16(&x.Dimension)
	FuncSlice(io, &x.TargetBiomes, io.Varuint32, io.Uint16)
	io.Float32(&x.Amount)
	io.Float32(&x.NoiseFrequencyScale)
	io.Uint32(&x.ReplacementIndex)
}
