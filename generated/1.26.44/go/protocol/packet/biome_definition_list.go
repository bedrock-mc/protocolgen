// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.44/go/protocol"

// BiomeDefinitionList is sent by the server to let the client know all biomes that are available
// and implemented on the server side. When enabled, it also includes information for the client to
// accurately recreate the server-side generation in vanilla worlds/servers for increased
// performance.
type BiomeDefinitionList struct {
	MapOfBiomeNamesToData []protocol.OrderedEntry[uint16, protocol.BiomeDefinitionData]
	// StringList is a makeshift dictionary implementation Mojang created to try and reduce the size of
	// the overall packet. It is a list of common strings that are used in the biome definitions, such
	// as biome names, float values or query expressions.
	StringList protocol.BiomeStringList
}

// Marshal reads or writes BiomeDefinitionList using its canonical wire layout.
func (x *BiomeDefinitionList) Marshal(io protocol.IO) {
	protocol.OrderedMap(io, &x.MapOfBiomeNamesToData, io.Varuint32, io.Uint16, func(value *protocol.BiomeDefinitionData) {
		value.Marshal(io)
	})
	x.StringList.Marshal(io)
}

// ID returns the protocol ID for BiomeDefinitionList.
func (*BiomeDefinitionList) ID() uint32 { return IDBiomeDefinitionList }
