// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

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
		value.Marshal(io)
	})
	OptionalFunc(io, &x.ConsolidatedFeatures, func(value *BiomeConsolidatedFeaturesData) {
		value.Marshal(io)
	})
	OptionalFunc(io, &x.MountainParams, func(value *BiomeMountainParamsData) {
		value.Marshal(io)
	})
	OptionalFunc(io, &x.SurfaceMaterialAdjustments, func(value *BiomeSurfaceMaterialAdjustmentData) {
		value.Marshal(io)
	})
	OptionalFunc(io, &x.OverworldGenRules, func(value *BiomeOverworldGenRulesData) {
		value.Marshal(io)
	})
	OptionalFunc(io, &x.MultinoiseGenRules, func(value *BiomeMultinoiseGenRulesData) {
		value.Marshal(io)
	})
	OptionalFunc(io, &x.LegacyWorldGenRules, func(value *BiomeLegacyWorldGenRulesData) {
		value.Marshal(io)
	})
	OptionalFunc(io, &x.ReplacementBiomes, func(value *BiomeReplacementsData) {
		value.Marshal(io)
	})
	OptionalFunc(io, &x.VillageType, func(value *VillageType) {
		IntegerFunc(value, io.Uint8)
	})
	OptionalFunc(io, &x.SurfaceBuilderData, func(value *BiomeSurfaceBuilderData) {
		value.Marshal(io)
	})
	OptionalFunc(io, &x.SubsurfaceBuilderData, func(value *BiomeSurfaceBuilderData) {
		value.Marshal(io)
	})
}
