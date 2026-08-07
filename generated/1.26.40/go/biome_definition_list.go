// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type BiomeDefinitionList struct {
	MapOfBiomeNamesToData []OrderedEntry[uint16, BiomeDefinitionData]
	StringList            BiomeStringList
}

// Marshal reads or writes BiomeDefinitionList using its canonical wire layout.
func (x *BiomeDefinitionList) Marshal(io IO) {
	if !io.Reading() && uint64(len(x.MapOfBiomeNamesToData)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.MapOfBiomeNamesToData), "map length overflows uint32")
		return
	}
	count1 := uint32(len(x.MapOfBiomeNamesToData))
	io.Varuint32(&count1)
	if io.Reading() {
		if uint64(count1) > uint64(^uint(0)>>1) {
			io.InvalidValue(count1, "map length overflows int")
			return
		}
		x.MapOfBiomeNamesToData = make([]OrderedEntry[uint16, BiomeDefinitionData], int(count1))
	}
	for index2 := range x.MapOfBiomeNamesToData {
		io.Uint16(&x.MapOfBiomeNamesToData[index2].Key)
		x.MapOfBiomeNamesToData[index2].Value.Marshal(io)
	}
	x.StringList.Marshal(io)
}
