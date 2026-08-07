// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type BiomeSurfaceMaterialAdjustmentData struct {
	Adjustments []BiomeElementData
}

// Marshal reads or writes BiomeSurfaceMaterialAdjustmentData using its canonical wire layout.
func (x *BiomeSurfaceMaterialAdjustmentData) Marshal(io IO) {
	FuncSlice(io, &x.Adjustments, io.Varuint32, func(value *BiomeElementData) {
		item := *value
		item.Marshal(io)
		*value = item
	})
}
