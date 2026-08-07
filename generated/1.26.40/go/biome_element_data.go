// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type BiomeElementData struct {
	NoiseFreqScale    float32
	NoiseLowerBound   float32
	NoiseUpperBound   float32
	HeightMinType     int32
	HeightMin         uint16
	HeightMaxType     int32
	HeightMax         uint16
	AdjustedMaterials BiomeSurfaceMaterialData
}

// Marshal reads or writes BiomeElementData using its canonical wire layout.
func (x *BiomeElementData) Marshal(io IO) {
	io.Float32(&x.NoiseFreqScale)
	io.Float32(&x.NoiseLowerBound)
	io.Float32(&x.NoiseUpperBound)
	io.Varint32(&x.HeightMinType)
	io.Uint16(&x.HeightMin)
	io.Varint32(&x.HeightMaxType)
	io.Uint16(&x.HeightMax)
	x.AdjustedMaterials.Marshal(io)
}
