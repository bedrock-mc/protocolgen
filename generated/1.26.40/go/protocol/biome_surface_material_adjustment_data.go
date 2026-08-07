// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type BiomeSurfaceMaterialAdjustmentData struct {
	Adjustments []BiomeElementData
}

// Marshal reads or writes BiomeSurfaceMaterialAdjustmentData using its canonical wire layout.
func (x *BiomeSurfaceMaterialAdjustmentData) Marshal(io IO) {
	Slice(io, &x.Adjustments)
}
