// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type BiomeSurfaceBuilderData struct {
	SurfaceMaterials           Optional[BiomeSurfaceMaterialData]
	HasDefaultOverworldSurface bool
	HasSwampSurface            bool
	HasFrozenOceanSurface      bool
	HasTheEndSurface           bool
	MesaSurface                Optional[BiomeMesaSurfaceData]
	CappedSurface              Optional[BiomeCappedSurfaceData]
	NoiseGradientSurface       Optional[BiomeNoiseGradientSurfaceData]
}

// Marshal reads or writes BiomeSurfaceBuilderData using its canonical wire layout.
func (x *BiomeSurfaceBuilderData) Marshal(io IO) {
	OptionalFunc(io, &x.SurfaceMaterials, func(value *BiomeSurfaceMaterialData) {
		item := *value
		item.Marshal(io)
		*value = item
	})
	io.Bool(&x.HasDefaultOverworldSurface)
	io.Bool(&x.HasSwampSurface)
	io.Bool(&x.HasFrozenOceanSurface)
	io.Bool(&x.HasTheEndSurface)
	OptionalFunc(io, &x.MesaSurface, func(value *BiomeMesaSurfaceData) {
		item := *value
		item.Marshal(io)
		*value = item
	})
	OptionalFunc(io, &x.CappedSurface, func(value *BiomeCappedSurfaceData) {
		item := *value
		item.Marshal(io)
		*value = item
	})
	OptionalFunc(io, &x.NoiseGradientSurface, func(value *BiomeNoiseGradientSurfaceData) {
		item := *value
		item.Marshal(io)
		*value = item
	})
}
