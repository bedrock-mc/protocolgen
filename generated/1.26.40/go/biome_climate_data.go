// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type BiomeClimateData struct {
	Temperature         float32
	Downfall            float32
	SnowAccumulationMin float32
	SnowAccumulationMax float32
}

// Marshal reads or writes BiomeClimateData using its canonical wire layout.
func (x *BiomeClimateData) Marshal(io IO) {
	io.Float32(&x.Temperature)
	io.Float32(&x.Downfall)
	io.Float32(&x.SnowAccumulationMin)
	io.Float32(&x.SnowAccumulationMax)
}
