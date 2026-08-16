package vanilladata

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/sandertv/gophertunnel/minecraft/nbt"
	genprotocol "protocolgen/generated/1.26.44/go/protocol"
	genpacket "protocolgen/generated/1.26.44/go/protocol/packet"
)

func TestBuildDerivedArtifactsUsesGeneratedProtocolPackets(t *testing.T) {
	emptyNBT, err := nbt.MarshalEncoding(map[string]any{}, nbt.NetworkLittleEndian)
	if err != nil {
		t.Fatal(err)
	}
	componentNBT, err := nbt.MarshalEncoding(map[string]any{"minecraft:test": int32(1)}, nbt.NetworkLittleEndian)
	if err != nil {
		t.Fatal(err)
	}
	items, err := genpacket.Encode(&genpacket.ItemRegistry{ItemData: []genprotocol.ItemData{
		{ItemName: "minecraft:stick", ItemID: 280, ItemVersion: genprotocol.ItemVersionNone, ItemComponentData: emptyNBT},
		{ItemName: "minecraft:test", ItemID: -7, IsComponentBased: true, ItemVersion: genprotocol.ItemVersionDataDriven, ItemComponentData: componentNBT},
	}})
	if err != nil {
		t.Fatal(err)
	}
	biomes, err := genpacket.Encode(&genpacket.BiomeDefinitionList{
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

	files, err := BuildDerivedArtifacts(map[string][]byte{
		"ItemRegistryPacket":        items,
		"BiomeDefinitionListPacket": biomes,
	})
	if err != nil {
		t.Fatalf("BuildDerivedArtifacts: %v", err)
	}
	var required map[string]struct {
		RuntimeID      int16  `json:"runtime_id"`
		ComponentBased bool   `json:"component_based"`
		Version        int32  `json:"version"`
		ComponentNBT   string `json:"component_nbt,omitempty"`
	}
	if err := json.Unmarshal(files["required_item_list.json"], &required); err != nil {
		t.Fatal(err)
	}
	if got := required["minecraft:test"]; got.RuntimeID != -7 || !got.ComponentBased || got.Version != 1 || got.ComponentNBT == "" {
		t.Fatalf("required item entry = %#v", got)
	}

	var definitions map[string]struct {
		ID          uint16   `json:"id"`
		FoliageSnow float32  `json:"foliageSnow"`
		Tags        []string `json:"tags"`
	}
	if err := json.Unmarshal(files["biome_definitions.json"], &definitions); err != nil {
		t.Fatal(err)
	}
	if got := definitions["minecraft:test_biome"]; got.ID != 42 || got.FoliageSnow != 0.25 || !reflect.DeepEqual(got.Tags, []string{"overworld"}) {
		t.Fatalf("biome definition = %#v", got)
	}
}

func TestBuildDerivedArtifactsRejectsPayloadThatGeneratedCodecCannotDecode(t *testing.T) {
	_, err := BuildDerivedArtifacts(map[string][]byte{"ItemRegistryPacket": {0xff}})
	if err == nil {
		t.Fatal("BuildDerivedArtifacts accepted malformed generated-protocol payload")
	}
}

func TestValidateGeneratedTargetRejectsAnotherGeneratedVersion(t *testing.T) {
	if err := ValidateGeneratedTarget(Target{MinecraftVersion: "1.26.44", ProtocolVersion: 2168}); err != nil {
		t.Fatalf("ValidateGeneratedTarget exact match: %v", err)
	}
	if err := ValidateGeneratedTarget(Target{MinecraftVersion: "1.26.50", ProtocolVersion: 2200}); err == nil {
		t.Fatal("ValidateGeneratedTarget accepted another generated version")
	}
}

func TestCheckedInCaptureReproducesDerivedArtifacts(t *testing.T) {
	root := filepath.Join("..", "..", "..")
	capture := filepath.Join(root, "generated", "1.26.44", "vanilla-data")
	payloads := make(map[string][]byte)
	for _, spec := range DefaultPacketSpecs() {
		data, err := os.ReadFile(filepath.Join(capture, spec.File))
		if os.IsNotExist(err) && spec.Optional {
			continue
		}
		if err != nil {
			t.Fatal(err)
		}
		payloads[spec.Name] = data
	}
	files, err := BuildDerivedArtifacts(payloads)
	if err != nil {
		t.Fatalf("BuildDerivedArtifacts checked-in capture: %v", err)
	}
	for name, got := range files {
		want, err := os.ReadFile(filepath.Join(capture, filepath.FromSlash(name)))
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("checked-in %s is not reproducible", name)
		}
	}
}
