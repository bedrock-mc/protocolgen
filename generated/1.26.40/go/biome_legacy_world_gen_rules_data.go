// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type BiomeLegacyWorldGenRulesData struct {
	LegacyPreHillsEdge []BiomeConditionalTransformationData
}

// Marshal reads or writes BiomeLegacyWorldGenRulesData using its canonical wire layout.
func (x *BiomeLegacyWorldGenRulesData) Marshal(io IO) {
	FuncSlice(io, &x.LegacyPreHillsEdge, io.Varuint32, func(value *BiomeConditionalTransformationData) {
		item := *value
		item.Marshal(io)
		*value = item
	})
}
