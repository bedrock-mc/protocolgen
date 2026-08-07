// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type BiomeMultinoiseGenRulesData struct {
	Temperature float32
	Humidity    float32
	Altitude    float32
	Weirdness   float32
	Weight      float32
}

// Marshal reads or writes BiomeMultinoiseGenRulesData using its canonical wire layout.
func (x *BiomeMultinoiseGenRulesData) Marshal(io IO) {
	io.Float32(&x.Temperature)
	io.Float32(&x.Humidity)
	io.Float32(&x.Altitude)
	io.Float32(&x.Weirdness)
	io.Float32(&x.Weight)
}
