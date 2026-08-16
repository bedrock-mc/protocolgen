//go:build protocolgen_12640

package vanilladata

import (
	"testing"

	"github.com/google/uuid"
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

func testResourcePacksPayload(t *testing.T) (info, stack, dataInfo []byte) {
	t.Helper()
	packUUID := uuid.MustParse("d34cfa4b-2ad1-453d-a0db-668b429a3ea0")
	info, err := genpacket.Encode(&genpacket.ResourcePacksInfo{
		ResourcePackRequired:       false,
		HasAddonPacks:              true,
		HasScripts:                 true,
		ForceDisableVibrantVisuals: true,
		WorldTemplateIDAndVersion: genprotocol.PackIDVersion{
			PackUUID:    uuid.MustParse("11111111-2222-3333-4444-555555555555"),
			PackVersion: genprotocol.SemVersion{Version: "2.3.4"},
		},
		ResourcePacks: []genprotocol.PackInfoData{{
			PackIDVersion:       genprotocol.PackIDVersionData{PackUUID: packUUID, PackVersion: genprotocol.SemVersionData{Version: "1.26.40"}},
			PackSize:            123456,
			ContentKey:          "content-key",
			SubpackName:         "subpack",
			ContentIdentity:     genprotocol.ContentIdentity{Identity: "content-identity"},
			HasScripts:          true,
			IsAddonPack:         true,
			IsRayTracingCapable: true,
			CDNURL:              "https://example.invalid/pack",
		}},
	})
	if err != nil {
		t.Fatal(err)
	}
	stack, err = genpacket.Encode(&genpacket.ResourcePackStack{
		TexturePackRequired: false,
		TexturePackList: []genprotocol.PackInstanceID{
			{PackID: packUUID.String(), Version: "1.26.40", SubPackName: "subpack"},
			{PackID: "b41c2785-c512-4a49-af56-3a87afd47c57", Version: "1.21.30"},
		},
		BaseGameVersion: "1.26.40",
		Experiments: genprotocol.Experiments{
			Toggles:                []genprotocol.ExperimentToggle{{Name: "cameras", Enabled: true}},
			ExperimentsEverToggled: true,
		},
		IncludeEditorPacks: true,
	})
	if err != nil {
		t.Fatal(err)
	}
	dataInfo, err = genpacket.Encode(&genpacket.ResourcePackDataInfo{
		ResourceName:   packUUID.String() + "_1.26.40",
		ChunkSize:      524288,
		NumberOfChunks: 1,
		FileSize:       123456,
		FileHash:       []byte{0xde, 0xad, 0xbe, 0xef},
		IsPremiumPack:  false,
		PackType:       6,
	})
	if err != nil {
		t.Fatal(err)
	}
	return info, stack, dataInfo
}
