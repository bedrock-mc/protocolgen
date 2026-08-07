// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type BiomeDefinitionList struct {
	MapOfBiomeNamesToData []protocol.OrderedEntry[uint16, protocol.BiomeDefinitionData]
	StringList            protocol.BiomeStringList
}

// Marshal reads or writes BiomeDefinitionList using its canonical wire layout.
func (x *BiomeDefinitionList) Marshal(io protocol.IO) {
	protocol.OrderedMap(io, &x.MapOfBiomeNamesToData, io.Varuint32, io.Uint16, func(value *protocol.BiomeDefinitionData) {
		value.Marshal(io)
	})
	x.StringList.Marshal(io)
}
