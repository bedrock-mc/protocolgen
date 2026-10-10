package generator

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestLockedBiomeIDsMatchRetailCatalog(t *testing.T) {
	data, err := os.ReadFile(filepath.Join("..", "..", "source", "inputs", "1.26.50", "biome_ids.json"))
	if err != nil {
		t.Fatal(err)
	}
	ids, err := DecodeBiomeIDs(data, "1.26.50", 2193)
	if err != nil {
		t.Fatal(err)
	}
	if len(ids) != 89 || ids["minecraft:ocean"] != 0 || ids["minecraft:plains"] != 1 || ids["minecraft:dappled_forest"] != 195 {
		t.Fatalf("unexpected retail biome IDs: count=%d ocean=%d plains=%d dappled_forest=%d", len(ids), ids["minecraft:ocean"], ids["minecraft:plains"], ids["minecraft:dappled_forest"])
	}
}

func TestBiomeIDsRejectWrongReleaseAndDuplicates(t *testing.T) {
	base := `{"schema_version":1,"minecraft_version":"1.26.50","protocol_version":2193,"source":"fixture","source_sha256":"1111111111111111111111111111111111111111111111111111111111111111","biomes":[{"name":"minecraft:plains","id":1},{"name":"minecraft:ocean","id":0}]}`
	for _, test := range []struct {
		name string
		data string
		want string
	}{
		{"wrong protocol", strings.Replace(base, `"protocol_version":2193`, `"protocol_version":2187`, 1), "do not match"},
		{"duplicate ID", strings.Replace(base, `"minecraft:ocean","id":0`, `"minecraft:ocean","id":1`, 1), "duplicate biome ID"},
		{"duplicate name", strings.Replace(base, `"minecraft:ocean","id":0`, `"minecraft:plains","id":0`, 1), "duplicate biome name"},
		{"missing ID", strings.Replace(base, `"minecraft:ocean","id":0`, `"minecraft:ocean"`, 1), "has no numeric ID"},
		{"null ID", strings.Replace(base, `"minecraft:ocean","id":0`, `"minecraft:ocean","id":null`, 1), "has no numeric ID"},
	} {
		t.Run(test.name, func(t *testing.T) {
			if _, err := DecodeBiomeIDs([]byte(test.data), "1.26.50", 2193); err == nil || !strings.Contains(err.Error(), test.want) {
				t.Fatalf("DecodeBiomeIDs error = %v, want %q", err, test.want)
			}
		})
	}
}

func TestBiomeIDsMustCoverCloudburstNames(t *testing.T) {
	path := filepath.Join("testdata", "cloudburst", "stripped_biome_definitions.json")
	for _, test := range []struct {
		name string
		ids  map[string]int32
		want string
	}{
		{"missing", map[string]int32{}, "cover 0 names"},
		{"wrong name", map[string]int32{"minecraft:ocean": 0}, `"minecraft:plains" has no locked numeric ID`},
	} {
		t.Run(test.name, func(t *testing.T) {
			if _, _, err := generateBiomes(path, test.ids); err == nil || !strings.Contains(err.Error(), test.want) {
				t.Fatalf("generateBiomes error = %v, want %q", err, test.want)
			}
		})
	}
}
