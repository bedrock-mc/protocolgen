// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type BiomeMesaSurfaceData struct {
	ClayMaterial     uint32
	HardClayMaterial uint32
	BrycePillars     bool
	HasForest        bool
}

// Marshal reads or writes BiomeMesaSurfaceData using its canonical wire layout.
func (x *BiomeMesaSurfaceData) Marshal(io IO) {
	io.Uint32(&x.ClayMaterial)
	io.Uint32(&x.HardClayMaterial)
	io.Bool(&x.BrycePillars)
	io.Bool(&x.HasForest)
}
