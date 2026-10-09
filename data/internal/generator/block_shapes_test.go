package generator

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

func TestApplyBlockShapesPreservesOtherProperties(t *testing.T) {
	states := []rawBlockState{{
		Name: "minecraft:stone", Hash: 100, Hardness: 3, Friction: 0.4,
		CollisionShape: json.RawMessage(`[]`), OutlineShape: json.RawMessage(`[]`),
		VisualShape: json.RawMessage(`[0,0,0,1,1,1]`), TintMethod: "None",
	}}
	want := states[0]
	want.CollisionShape = json.RawMessage(`[[0.25,0,0.25,0.75,1,0.75]]`)
	want.OutlineShape = json.RawMessage(`[]`)
	want.TintMethod = "Grass"
	path := writeBlockShapes(t, `[{
		"name":"minecraft:stone","blockStateHash":100,
		"collisionShape":[[0.25,0,0.25,0.75,1,0.75]],
		"shape":[0,0,0,1,0,1],"tintMethod":"Grass"
	}]`)
	if err := applyBlockShapes(states, path); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(states[0], want) {
		t.Fatalf("overlaid state = %+v, want %+v", states[0], want)
	}
	boxes, available, err := decodeShape(states[0].OutlineShape)
	if err != nil || !available || len(boxes) != 0 {
		t.Fatalf("zero-volume outline = %v, available %t, error %v", boxes, available, err)
	}
}

func TestApplyBlockShapesRejectsIncompleteOrInvalidSources(t *testing.T) {
	const valid = `{"name":"minecraft:stone","blockStateHash":100,"collisionShape":[],"shape":[0,0,0,1,1,1]}`
	for _, test := range []struct {
		name, source, want string
	}{
		{"empty", `[]`, "coverage"},
		{"extra", `[` + valid + `,` + valid + `]`, "coverage"},
		{"missing_hash", `[` + strings.Replace(valid, `:100`, `:99`, 1) + `]`, "missing"},
		{"wrong_name", `[` + strings.Replace(valid, "minecraft:stone", "minecraft:air", 1) + `]`, "names"},
		{"missing_collision", `[{"name":"minecraft:stone","blockStateHash":100,"shape":[]}]`, "collision shape"},
		{"missing_outline", `[{"name":"minecraft:stone","blockStateHash":100,"collisionShape":[]}]`, "outline shape"},
		{"reversed_collision", `[` + strings.Replace(valid, `"collisionShape":[]`, `"collisionShape":[[1,0,0,0,1,1]]`, 1) + `]`, "collision shape"},
		{"malformed_outline", `[` + strings.Replace(valid, `[0,0,0,1,1,1]`, `[0,0,1]`, 1) + `]`, "outline shape"},
	} {
		t.Run(test.name, func(t *testing.T) {
			states := []rawBlockState{{Name: "minecraft:stone", Hash: 100}}
			err := applyBlockShapes(states, writeBlockShapes(t, test.source))
			if err == nil || !strings.Contains(err.Error(), test.want) {
				t.Fatalf("error = %v, want %q", err, test.want)
			}
		})
	}
	t.Run("duplicate_hash", func(t *testing.T) {
		states := []rawBlockState{{Name: "minecraft:stone", Hash: 100}, {Name: "minecraft:air", Hash: 50}}
		err := applyBlockShapes(states, writeBlockShapes(t, `[`+valid+`,`+valid+`]`))
		if err == nil || !strings.Contains(err.Error(), "duplicate") {
			t.Fatalf("error = %v, want duplicate hash rejection", err)
		}
	})
}

func TestGenerateUsesBlockShapeOverlayOnlyForBlocks(t *testing.T) {
	cfg := Config{CloudburstDir: "testdata/cloudburst", BDSDir: "testdata/bds"}
	original, _, err := Generate(cfg)
	if err != nil {
		t.Fatal(err)
	}
	cfg.BlockShapesPath = writeBlockShapes(t, `[
		{"name":"minecraft:air","blockStateHash":50,"collisionShape":[],"shape":[0,0,0,0,0,0],"tintMethod":"None"},
		{"name":"minecraft:stone","blockStateHash":100,"collisionShape":[[0.25,0,0.25,0.75,1,0.75]],"shape":[0,0,0,1,1,1],"tintMethod":"Grass"}
	]`)
	overlaid, _, err := Generate(cfg)
	if err != nil {
		t.Fatal(err)
	}
	assertContains(t, overlaid, "block/index_generated.go", `TintMethod: "Grass"`)
	assertContains(t, overlaid, "block/index_generated.go", `{0.25, 0, 0.25, 0.75, 1, 0.75}`)
	for name, content := range original {
		if strings.HasPrefix(name, "block/") || name == "semantic_sources.json" || name == "version_generated.go" {
			continue
		}
		if !bytes.Equal(content, overlaid[name]) {
			t.Errorf("block shape overlay changed %s", name)
		}
	}
}

// writeBlockShapes writes a synthetic source for a geometry overlay test.
func writeBlockShapes(t *testing.T, content string) string {
	t.Helper()
	path := filepath.Join(t.TempDir(), "block_states.json")
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}
	return path
}
