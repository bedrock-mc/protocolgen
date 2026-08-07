// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type BiomeOverworldGenRulesData struct {
	HillsTransformations  []BiomeWeightedData
	MutateTransformations []BiomeWeightedData
	RiverTransformations  []BiomeWeightedData
	ShoreTransformations  []BiomeWeightedData
	PreHillsEdge          []BiomeConditionalTransformationData
	PostShoreEdge         []BiomeConditionalTransformationData
	Climate               []BiomeWeightedTemperatureData
}

// Marshal reads or writes BiomeOverworldGenRulesData using its canonical wire layout.
func (x *BiomeOverworldGenRulesData) Marshal(io IO) {
	Slice(io, &x.HillsTransformations)
	Slice(io, &x.MutateTransformations)
	Slice(io, &x.RiverTransformations)
	Slice(io, &x.ShoreTransformations)
	Slice(io, &x.PreHillsEdge)
	Slice(io, &x.PostShoreEdge)
	Slice(io, &x.Climate)
}
