// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type BiomeDefinitionChunkGenData struct {
	Climate                    Optional[BiomeClimateData]
	ConsolidatedFeatures       Optional[BiomeConsolidatedFeaturesData]
	MountainParams             Optional[BiomeMountainParamsData]
	SurfaceMaterialAdjustments Optional[BiomeSurfaceMaterialAdjustmentData]
	OverworldGenRules          Optional[BiomeOverworldGenRulesData]
	MultinoiseGenRules         Optional[BiomeMultinoiseGenRulesData]
	LegacyWorldGenRules        Optional[BiomeLegacyWorldGenRulesData]
	ReplacementBiomes          Optional[BiomeReplacementsData]
	VillageType                Optional[VillageType]
	SurfaceBuilderData         Optional[BiomeSurfaceBuilderData]
	SubsurfaceBuilderData      Optional[BiomeSurfaceBuilderData]
}

// Marshal reads or writes BiomeDefinitionChunkGenData using its canonical wire layout.
func (x *BiomeDefinitionChunkGenData) Marshal(io IO) {
	OptionalFunc(io, &x.Climate, func(value *BiomeClimateData) {
		item := *value
		item.Marshal(io)
		*value = item
	})
	OptionalFunc(io, &x.ConsolidatedFeatures, func(value *BiomeConsolidatedFeaturesData) {
		item := *value
		item.Marshal(io)
		*value = item
	})
	OptionalFunc(io, &x.MountainParams, func(value *BiomeMountainParamsData) {
		item := *value
		item.Marshal(io)
		*value = item
	})
	OptionalFunc(io, &x.SurfaceMaterialAdjustments, func(value *BiomeSurfaceMaterialAdjustmentData) {
		item := *value
		item.Marshal(io)
		*value = item
	})
	OptionalFunc(io, &x.OverworldGenRules, func(value *BiomeOverworldGenRulesData) {
		item := *value
		item.Marshal(io)
		*value = item
	})
	OptionalFunc(io, &x.MultinoiseGenRules, func(value *BiomeMultinoiseGenRulesData) {
		item := *value
		item.Marshal(io)
		*value = item
	})
	OptionalFunc(io, &x.LegacyWorldGenRules, func(value *BiomeLegacyWorldGenRulesData) {
		item := *value
		item.Marshal(io)
		*value = item
	})
	OptionalFunc(io, &x.ReplacementBiomes, func(value *BiomeReplacementsData) {
		item := *value
		item.Marshal(io)
		*value = item
	})
	OptionalFunc(io, &x.VillageType, func(value *VillageType) {
		item := *value
		IntegerFunc(&item, io.Uint8)
		*value = item
	})
	OptionalFunc(io, &x.SurfaceBuilderData, func(value *BiomeSurfaceBuilderData) {
		item := *value
		item.Marshal(io)
		*value = item
	})
	OptionalFunc(io, &x.SubsurfaceBuilderData, func(value *BiomeSurfaceBuilderData) {
		item := *value
		item.Marshal(io)
		*value = item
	})
}
