// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

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
		item := *value
		item.Marshal(io)
		*value = item
	})
	FuncSlice(io, &x.MutateTransformations, io.Varuint32, func(value *BiomeWeightedData) {
		item := *value
		item.Marshal(io)
		*value = item
	})
	FuncSlice(io, &x.RiverTransformations, io.Varuint32, func(value *BiomeWeightedData) {
		item := *value
		item.Marshal(io)
		*value = item
	})
	FuncSlice(io, &x.ShoreTransformations, io.Varuint32, func(value *BiomeWeightedData) {
		item := *value
		item.Marshal(io)
		*value = item
	})
	FuncSlice(io, &x.PreHillsEdge, io.Varuint32, func(value *BiomeConditionalTransformationData) {
		item := *value
		item.Marshal(io)
		*value = item
	})
	FuncSlice(io, &x.PostShoreEdge, io.Varuint32, func(value *BiomeConditionalTransformationData) {
		item := *value
		item.Marshal(io)
		*value = item
	})
	FuncSlice(io, &x.Climate, io.Varuint32, func(value *BiomeWeightedTemperatureData) {
		item := *value
		item.Marshal(io)
		*value = item
	})
}
