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
	FuncSlice(io, &x.HillsTransformations, io.Varuint32, func(value *BiomeWeightedData) {
		value.Marshal(io)
	})
	FuncSlice(io, &x.MutateTransformations, io.Varuint32, func(value *BiomeWeightedData) {
		value.Marshal(io)
	})
	FuncSlice(io, &x.RiverTransformations, io.Varuint32, func(value *BiomeWeightedData) {
		value.Marshal(io)
	})
	FuncSlice(io, &x.ShoreTransformations, io.Varuint32, func(value *BiomeWeightedData) {
		value.Marshal(io)
	})
	FuncSlice(io, &x.PreHillsEdge, io.Varuint32, func(value *BiomeConditionalTransformationData) {
		value.Marshal(io)
	})
	FuncSlice(io, &x.PostShoreEdge, io.Varuint32, func(value *BiomeConditionalTransformationData) {
		value.Marshal(io)
	})
	FuncSlice(io, &x.Climate, io.Varuint32, func(value *BiomeWeightedTemperatureData) {
		value.Marshal(io)
	})
}
