//go:build protocolgen_12640

package vanilladata

import (
	"testing"

	genprotocol "protocolgen/generated/1.26.40/go/protocol"
	genpacket "protocolgen/generated/1.26.40/go/protocol/packet"
)

func testGeneratedTarget() Target {
	return Target{MinecraftVersion: genprotocol.GAME_VERSION, ProtocolVersion: genprotocol.PROTOCOL_VERSION}
}

func testGeneratedVersion() string { return genprotocol.GAME_VERSION }

func testItemRegistryPayload(t *testing.T, emptyNBT, componentNBT []byte) []byte {
	t.Helper()
	payload, err := genpacket.Encode(&genpacket.ItemRegistry{ItemData: []genprotocol.ItemData{
		{ItemName: "minecraft:stick", ItemID: 280, ItemVersion: genprotocol.ItemVersionNone, ItemComponentData: emptyNBT},
		{ItemName: "minecraft:test", ItemID: -7, IsComponentBased: true, ItemVersion: genprotocol.ItemVersionDataDriven, ItemComponentData: componentNBT},
	}})
	if err != nil {
		t.Fatal(err)
	}
	return payload
}

func testBiomeDefinitionListPayload(t *testing.T) []byte {
	t.Helper()
	payload, err := genpacket.Encode(&genpacket.BiomeDefinitionList{
		MapOfBiomeNamesToData: []genprotocol.OrderedEntry[uint16, genprotocol.BiomeDefinitionData]{
			{Key: 0, Value: genprotocol.BiomeDefinitionData{
				ID: 42, Temperature: 0.8, Downfall: 0.4, FoliageSnow: 0.25, Depth: 0.1, Scale: 0.2,
				MapWaterColorARGB: -1525384027, Rain: true,
				Tags: genprotocol.Option(genprotocol.BiomeTagsData{Tags: []uint16{1}}),
			}},
		},
		StringList: genprotocol.BiomeStringList{Strings: []string{"minecraft:test_biome", "overworld"}},
	})
	if err != nil {
		t.Fatal(err)
	}
	return payload
}
