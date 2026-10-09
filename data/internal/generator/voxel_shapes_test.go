package generator

import (
	"encoding/json"
	"math"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

// TestVoxelGridBoxesPreservesOccupancy checks axis order, disconnected regions,
// bits across byte and 64-bit boundaries, and compact merging of solid regions.
func TestVoxelGridBoxesPreservesOccupancy(t *testing.T) {
	for _, test := range []struct {
		name       string
		dimensions [3]int
		filled     []int
		all        bool
		wantBoxes  int
	}{
		{name: "empty", dimensions: [3]int{}, wantBoxes: 0},
		{name: "disconnected", dimensions: [3]int{2, 3, 13}, filled: []int{0, 1, 13, 14, 65, 76, 77}, wantBoxes: 3},
		{name: "solid", dimensions: [3]int{2, 3, 13}, all: true, wantBoxes: 1},
	} {
		t.Run(test.name, func(t *testing.T) {
			count := test.dimensions[0] * test.dimensions[1] * test.dimensions[2]
			filled := make([]bool, count)
			for _, index := range test.filled {
				filled[index] = true
			}
			if test.all {
				for index := range filled {
					filled[index] = true
				}
			}
			grid := voxelTestGrid(t, test.dimensions, filled)
			boxes, err := voxelGridBoxes(grid)
			if err != nil {
				t.Fatal(err)
			}
			if len(boxes) != test.wantBoxes {
				t.Fatalf("got %d boxes, want %d: %v", len(boxes), test.wantBoxes, boxes)
			}
			coordinates := [3][]float32{grid.XCoordinates, grid.YCoordinates, grid.ZCoordinates}
			for x := 0; x < test.dimensions[0]; x++ {
				for y := 0; y < test.dimensions[1]; y++ {
					for z := 0; z < test.dimensions[2]; z++ {
						position := [3]int{x, y, z}
						var midpoint [3]float32
						for axis, index := range position {
							midpoint[axis] = (coordinates[axis][index] + coordinates[axis][index+1]) / 2
						}
						coverage := 0
						for _, box := range boxes {
							if midpoint[0] > box[0] && midpoint[0] < box[3] && midpoint[1] > box[1] && midpoint[1] < box[4] && midpoint[2] > box[2] && midpoint[2] < box[5] {
								coverage++
							}
						}
						want := 0
						if filled[(x*test.dimensions[1]+y)*test.dimensions[2]+z] {
							want = 1
						}
						if coverage != want {
							t.Fatalf("cell %v covered %d times, want %d", position, coverage, want)
						}
					}
				}
			}
			again, err := voxelGridBoxes(grid)
			if err != nil || !reflect.DeepEqual(boxes, again) {
				t.Fatalf("decomposition changed on second call: %v, %v", again, err)
			}
		})
	}
}

// TestVoxelGridRejectsMalformedData prevents malformed grids from changing shape geometry.
func TestVoxelGridRejectsMalformedData(t *testing.T) {
	for _, test := range []struct {
		name   string
		change func(*rawVoxelGrid)
		want   string
	}{
		{"missing_cells", func(g *rawVoxelGrid) { g.Cells = nil }, "missing cells"},
		{"missing_size", func(g *rawVoxelGrid) { g.Cells.XSize = nil }, "size is missing"},
		{"negative_size", func(g *rawVoxelGrid) { *g.Cells.XSize = -1 }, "size"},
		{"wrong_size", func(g *rawVoxelGrid) { *g.Cells.XSize++ }, "size"},
		{"missing_coordinates", func(g *rawVoxelGrid) { g.XCoordinates = nil }, "coordinates"},
		{"reversed_coordinates", func(g *rawVoxelGrid) { g.YCoordinates[1] = -1 }, "increasing"},
		{"duplicate_coordinates", func(g *rawVoxelGrid) { g.ZCoordinates[1] = g.ZCoordinates[0] }, "increasing"},
		{"infinite_coordinate", func(g *rawVoxelGrid) { g.XCoordinates[1] = float32(math.Inf(1)) }, "non-finite"},
		{"nan_coordinate", func(g *rawVoxelGrid) { g.XCoordinates[1] = float32(math.NaN()) }, "non-finite"},
		{"short_storage", func(g *rawVoxelGrid) { g.Cells.Storage = nil }, "storage"},
		{"long_storage", func(g *rawVoxelGrid) { g.Cells.Storage = append(g.Cells.Storage, 0) }, "storage"},
		{"large_byte", func(g *rawVoxelGrid) { g.Cells.Storage[0] = 256 }, "not a byte"},
		{"padding_bits", func(g *rawVoxelGrid) { g.Cells.Storage[0] = 128 }, "padding bits"},
	} {
		t.Run(test.name, func(t *testing.T) {
			grid := voxelTestGrid(t, [3]int{1, 1, 1}, []bool{true})
			test.change(&grid)
			if _, err := voxelGridBoxes(grid); err == nil || !strings.Contains(err.Error(), test.want) {
				t.Fatalf("error = %v, want %q", err, test.want)
			}
		})
	}
}

// TestEmptyVoxelGridRequiresFields distinguishes an empty shape from incomplete input.
func TestEmptyVoxelGridRequiresFields(t *testing.T) {
	for _, cells := range []string{
		`{}`,
		`{"xSize":null,"ySize":0,"zSize":0,"storage":[]}`,
		`{"xSize":0,"ySize":0,"zSize":0}`,
		`{"xSize":0,"ySize":0,"zSize":0,"storage":null}`,
	} {
		var grid rawVoxelGrid
		err := json.Unmarshal([]byte(`{"cells":`+cells+`,"xCoordinates":[0],"yCoordinates":[0],"zCoordinates":[0]}`), &grid)
		if err == nil {
			_, err = voxelGridBoxes(grid)
		}
		if err == nil {
			t.Fatalf("accepted incomplete empty grid: %s", cells)
		}
	}
}

// TestVoxelGridRejectsNullNumbers prevents JSON null from fabricating zero-valued geometry.
func TestVoxelGridRejectsNullNumbers(t *testing.T) {
	for _, source := range []string{
		`{"cells":{"xSize":1,"ySize":1,"zSize":1,"storage":[null]},"xCoordinates":[0,1],"yCoordinates":[0,1],"zCoordinates":[0,1]}`,
		`{"cells":{"xSize":1,"ySize":1,"zSize":1,"storage":[1]},"xCoordinates":[null,1],"yCoordinates":[0,1],"zCoordinates":[0,1]}`,
	} {
		var grid rawVoxelGrid
		if err := json.Unmarshal([]byte(source), &grid); err == nil || !strings.Contains(err.Error(), "is null") {
			t.Fatalf("error = %v, want null number rejection", err)
		}
	}
}

// TestGenerateVoxelShapesValidatesNames checks named references and stable output order.
func TestGenerateVoxelShapesValidatesNames(t *testing.T) {
	grid := voxelTestGrid(t, [3]int{1, 1, 1}, []bool{true})
	encoded, err := json.Marshal([]rawVoxelGrid{grid})
	if err != nil {
		t.Fatal(err)
	}
	for _, test := range []struct {
		name  string
		names string
		want  string
	}{
		{"duplicate", `{"minecraft:cube":0,"minecraft:cube":0}`, "duplicate"},
		{"negative", `{"minecraft:cube":-1}`, "invalid shape index"},
		{"outside", `{"minecraft:cube":1}`, "invalid shape index"},
		{"null_index", `{"minecraft:cube":null}`, "no shape index"},
		{"empty_name", `{"":0}`, "name is empty"},
		{"empty_names", `{}`, "no named voxel shapes"},
		{"null_names", `null`, "must be an object"},
	} {
		t.Run(test.name, func(t *testing.T) {
			path := voxelTestFile(t, `{"names":`+test.names+`,"shapes":`+string(encoded)+`}`)
			if _, _, err := generateVoxelShapes(path); err == nil || !strings.Contains(err.Error(), test.want) {
				t.Fatalf("error = %v, want %q", err, test.want)
			}
		})
	}
	first, count, err := generateVoxelShapes(voxelTestFile(t, `{"names":{"minecraft:z":0,"minecraft:a":0},"shapes":`+string(encoded)+`}`))
	if err != nil || count != 2 {
		t.Fatalf("generate: count=%d, error=%v", count, err)
	}
	second, _, err := generateVoxelShapes(voxelTestFile(t, `{"names":{"minecraft:a":0,"minecraft:z":0},"shapes":`+string(encoded)+`}`))
	if err != nil || !reflect.DeepEqual(first, second) {
		t.Fatalf("name ordering changed output: %v", err)
	}
}

// voxelTestGrid builds an asymmetric grid with distinct block-local coordinate steps.
func voxelTestGrid(t *testing.T, dimensions [3]int, filled []bool) rawVoxelGrid {
	t.Helper()
	storage := make([]uint16, (len(filled)+7)/8)
	for index, value := range filled {
		if value {
			storage[index/8] |= 1 << (index % 8)
		}
	}
	var coordinates [3][]float32
	for axis, size := range dimensions {
		coordinates[axis] = make([]float32, size+1)
		for index := range coordinates[axis] {
			coordinates[axis][index] = float32(index*(axis+1)) / 8
		}
	}
	encoded, err := json.Marshal(map[string]any{
		"cells":        map[string]any{"xSize": dimensions[0], "ySize": dimensions[1], "zSize": dimensions[2], "storage": storage},
		"xCoordinates": coordinates[0], "yCoordinates": coordinates[1], "zCoordinates": coordinates[2],
	})
	if err != nil {
		t.Fatal(err)
	}
	var grid rawVoxelGrid
	if err := json.Unmarshal(encoded, &grid); err != nil {
		t.Fatal(err)
	}
	return grid
}

// voxelTestFile writes one source document in an isolated test directory.
func voxelTestFile(t *testing.T, source string) string {
	t.Helper()
	path := filepath.Join(t.TempDir(), "voxel_shapes.json")
	if err := os.WriteFile(path, []byte(source), 0o600); err != nil {
		t.Fatal(err)
	}
	return path
}
