// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type BiomeDefinitionList struct {
	MapOfBiomeNamesToData []OrderedEntry[uint16, BiomeDefinitionData]
	StringList            BiomeStringList
}

// Marshal reads or writes BiomeDefinitionList using its canonical wire layout.
func (x *BiomeDefinitionList) Marshal(io IO) {
	OrderedMap(io, &x.MapOfBiomeNamesToData, io.Varuint32, io.Uint16, func(value *BiomeDefinitionData) {
		value.Marshal(io)
	})
	x.StringList.Marshal(io)
}
