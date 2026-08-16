package vanilladata

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/sandertv/gophertunnel/minecraft/nbt"
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
	items := testItemRegistryPayload(t, emptyNBT, componentNBT)
	biomes := testBiomeDefinitionListPayload(t)

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

func TestBuildDerivedArtifactsEmitsResourcePackMetadata(t *testing.T) {
	info, stack, dataInfo := testResourcePacksPayload(t)
	files, err := BuildDerivedArtifacts(map[string][]byte{
		"ResourcePacksInfoPacket":    info,
		"ResourcePackStackPacket":    stack,
		"ResourcePackDataInfoPacket": dataInfo,
	})
	if err != nil {
		t.Fatalf("BuildDerivedArtifacts resource packs: %v", err)
	}
	var document struct {
		Info struct {
			ResourcePackRequired       bool `json:"resource_pack_required"`
			HasAddonPacks              bool `json:"has_addon_packs"`
			HasScripts                 bool `json:"has_scripts"`
			ForceDisableVibrantVisuals bool `json:"force_disable_vibrant_visuals"`
			WorldTemplate              struct {
				UUID    string `json:"uuid"`
				Version string `json:"version"`
			} `json:"world_template"`
		} `json:"info"`
		Packs []struct {
			UUID                string `json:"uuid"`
			Version             string `json:"version"`
			Size                uint64 `json:"size"`
			ContentKey          string `json:"content_key"`
			SubpackName         string `json:"subpack_name"`
			ContentIdentity     string `json:"content_identity"`
			HasScripts          bool   `json:"has_scripts"`
			IsAddonPack         bool   `json:"is_addon_pack"`
			IsRayTracingCapable bool   `json:"is_ray_tracing_capable"`
			CDNURL              string `json:"cdn_url"`
		} `json:"packs"`
		Stack struct {
			TexturePackRequired bool `json:"texture_pack_required"`
			Packs               []struct {
				PackID      string `json:"pack_id"`
				Version     string `json:"version"`
				SubpackName string `json:"subpack_name"`
			} `json:"packs"`
			BaseGameVersion string `json:"base_game_version"`
			Experiments     []struct {
				Name    string `json:"name"`
				Enabled bool   `json:"enabled"`
			} `json:"experiments"`
			ExperimentsEverToggled bool `json:"experiments_ever_toggled"`
			IncludeEditorPacks     bool `json:"include_editor_packs"`
		} `json:"stack"`
		Data struct {
			ResourceName   string `json:"resource_name"`
			ChunkSize      uint32 `json:"chunk_size"`
			NumberOfChunks uint32 `json:"number_of_chunks"`
			FileSize       uint64 `json:"file_size"`
			FileHash       string `json:"file_hash"`
			PackType       uint8  `json:"pack_type"`
		} `json:"data_info"`
	}
	if err := json.Unmarshal(files["resource_packs.json"], &document); err != nil {
		t.Fatal(err)
	}
	if document.Info.ResourcePackRequired || !document.Info.HasAddonPacks || !document.Info.HasScripts || !document.Info.ForceDisableVibrantVisuals {
		t.Fatalf("resource pack info flags = %#v", document.Info)
	}
	if document.Info.WorldTemplate.UUID != "11111111-2222-3333-4444-555555555555" || document.Info.WorldTemplate.Version != "2.3.4" {
		t.Fatalf("world template = %#v", document.Info.WorldTemplate)
	}
	if len(document.Packs) != 1 || document.Packs[0].UUID != "d34cfa4b-2ad1-453d-a0db-668b429a3ea0" || document.Packs[0].Size != 123456 || document.Packs[0].ContentKey != "content-key" || !document.Packs[0].IsAddonPack {
		t.Fatalf("resource pack metadata = %#v", document.Packs)
	}
	if len(document.Stack.Packs) != 2 || document.Stack.Packs[1].PackID != "b41c2785-c512-4a49-af56-3a87afd47c57" || document.Stack.Packs[0].SubpackName != "subpack" {
		t.Fatalf("resource pack stack = %#v", document.Stack.Packs)
	}
	if document.Stack.BaseGameVersion != testGeneratedVersion() || len(document.Stack.Experiments) != 1 || document.Stack.Experiments[0].Name != "cameras" || !document.Stack.Experiments[0].Enabled || !document.Stack.ExperimentsEverToggled || !document.Stack.IncludeEditorPacks {
		t.Fatalf("resource pack stack metadata = %#v", document.Stack)
	}
	if document.Data.ResourceName == "" || document.Data.ChunkSize != 524288 || document.Data.NumberOfChunks != 1 || document.Data.FileSize != 123456 || document.Data.FileHash != "deadbeef" || document.Data.PackType != 6 {
		t.Fatalf("resource pack data info = %#v", document.Data)
	}
}

func TestBuildDerivedArtifactsRejectsResourcePackInfoWithoutStack(t *testing.T) {
	info, _, _ := testResourcePacksPayload(t)
	if _, err := BuildDerivedArtifacts(map[string][]byte{"ResourcePacksInfoPacket": info}); err == nil {
		t.Fatal("BuildDerivedArtifacts accepted ResourcePacksInfoPacket without ResourcePackStackPacket")
	}
}

func TestBuildDerivedArtifactsRejectsPayloadThatGeneratedCodecCannotDecode(t *testing.T) {
	_, err := BuildDerivedArtifacts(map[string][]byte{"ItemRegistryPacket": {0xff}})
	if err == nil {
		t.Fatal("BuildDerivedArtifacts accepted malformed generated-protocol payload")
	}
}

func TestValidateGeneratedTargetRejectsAnotherGeneratedVersion(t *testing.T) {
	if err := ValidateGeneratedTarget(testGeneratedTarget()); err != nil {
		t.Fatalf("ValidateGeneratedTarget exact match: %v", err)
	}
	if err := ValidateGeneratedTarget(Target{MinecraftVersion: "1.26.50", ProtocolVersion: 2200}); err == nil {
		t.Fatal("ValidateGeneratedTarget accepted another generated version")
	}
}

func TestCheckedInCaptureReproducesDerivedArtifacts(t *testing.T) {
	root := filepath.Join("..", "..", "..")
	capture := filepath.Join(root, "generated", testGeneratedVersion(), "vanilla-data")
	if _, err := os.Stat(capture); os.IsNotExist(err) {
		t.Skipf("checked-in capture for %s is not present yet", testGeneratedVersion())
	} else if err != nil {
		t.Fatal(err)
	}
	payloads := make(map[string][]byte)
	for _, spec := range DefaultPacketSpecs() {
		data, err := os.ReadFile(filepath.Join(capture, spec.File))
		if os.IsNotExist(err) && spec.Optional {
			continue
		}
		if os.IsNotExist(err) && (spec.Name == "ResourcePacksInfoPacket" || spec.Name == "ResourcePackStackPacket") {
			t.Skipf("checked-in capture for %s predates resource-pack packet capture", testGeneratedVersion())
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
