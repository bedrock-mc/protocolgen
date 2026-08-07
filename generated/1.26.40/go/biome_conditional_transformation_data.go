// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type BiomeConditionalTransformationData struct {
	TransformsInto      []BiomeWeightedData
	ConditionJson       uint16
	MinPassingNeighbors uint32
}

// Marshal reads or writes BiomeConditionalTransformationData using its canonical wire layout.
func (x *BiomeConditionalTransformationData) Marshal(io IO) {
	FuncSlice(io, &x.TransformsInto, io.Varuint32, func(value *BiomeWeightedData) {
		item := *value
		item.Marshal(io)
		*value = item
	})
	io.Uint16(&x.ConditionJson)
	io.Uint32(&x.MinPassingNeighbors)
}
