package generator

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math"
	"sort"
)

// voxelShapeNames rejects duplicate identifiers instead of silently replacing them.
type voxelShapeNames map[string]int

// UnmarshalJSON reads each named grid reference exactly once.
func (names *voxelShapeNames) UnmarshalJSON(data []byte) error {
	decoder := json.NewDecoder(bytes.NewReader(data))
	token, err := decoder.Token()
	if err != nil {
		return err
	}
	if token != json.Delim('{') {
		return fmt.Errorf("voxel shape names must be an object")
	}
	values := make(voxelShapeNames)
	for decoder.More() {
		token, err := decoder.Token()
		if err != nil {
			return err
		}
		name := token.(string)
		if name == "" {
			return fmt.Errorf("voxel shape name is empty")
		}
		if _, exists := values[name]; exists {
			return fmt.Errorf("duplicate voxel shape %s", name)
		}
		var index *int
		if err := decoder.Decode(&index); err != nil {
			return err
		}
		if index == nil {
			return fmt.Errorf("%s has no shape index", name)
		}
		values[name] = *index
	}
	if _, err := decoder.Token(); err != nil {
		return err
	}
	*names = values
	return nil
}

// rawVoxelGrid stores Cloudburst's block-local coordinates and packed occupancy.
type rawVoxelGrid struct {
	Cells *struct {
		XSize   *int                 `json:"xSize"`
		YSize   *int                 `json:"ySize"`
		ZSize   *int                 `json:"zSize"`
		Storage voxelNumbers[uint16] `json:"storage"`
	} `json:"cells"`
	XCoordinates voxelNumbers[float32] `json:"xCoordinates"`
	YCoordinates voxelNumbers[float32] `json:"yCoordinates"`
	ZCoordinates voxelNumbers[float32] `json:"zCoordinates"`
}

// voxelNumbers keeps JSON null values from silently becoming zero coordinates or bits.
type voxelNumbers[T float32 | uint16] []T

// UnmarshalJSON requires an explicit number at every array position.
func (numbers *voxelNumbers[T]) UnmarshalJSON(data []byte) error {
	var values []*T
	if err := json.Unmarshal(data, &values); err != nil {
		return err
	}
	if values == nil {
		return fmt.Errorf("numeric array is null")
	}
	result := make(voxelNumbers[T], len(values))
	for index, value := range values {
		if value == nil {
			return fmt.Errorf("numeric array element %d is null", index)
		}
		result[index] = *value
	}
	*numbers = result
	return nil
}

type generatedVoxelShape struct {
	Name  string
	Boxes []generatedBox
}

// generateVoxelShapes publishes named Cloudburst grids as deterministic box unions.
func generateVoxelShapes(path string) (map[string][]byte, int, error) {
	var raw struct {
		Names  voxelShapeNames `json:"names"`
		Shapes []rawVoxelGrid  `json:"shapes"`
	}
	if err := readJSON(path, &raw, false); err != nil {
		return nil, 0, err
	}
	if len(raw.Names) == 0 {
		return nil, 0, fmt.Errorf("no named voxel shapes in %s", path)
	}
	grids := make([][]generatedBox, len(raw.Shapes))
	for index, grid := range raw.Shapes {
		boxes, err := voxelGridBoxes(grid)
		if err != nil {
			return nil, 0, fmt.Errorf("voxel grid %d: %w", index, err)
		}
		grids[index] = boxes
	}
	values := make([]generatedVoxelShape, 0, len(raw.Names))
	for name, index := range raw.Names {
		if index < 0 || index >= len(grids) {
			return nil, 0, fmt.Errorf("%s has invalid shape index %d", name, index)
		}
		values = append(values, generatedVoxelShape{Name: name, Boxes: grids[index]})
	}
	sort.Slice(values, func(i, j int) bool { return values[i].Name < values[j].Name })

	names := make([]string, len(values))
	for index, value := range values {
		names[index] = value.Name
	}
	identifiers, err := exportedIdentifiers(names)
	if err != nil {
		return nil, 0, err
	}
	fileNames, err := generatedFileNames(names)
	if err != nil {
		return nil, 0, err
	}

	files := make(map[string][]byte, len(values)+1)
	for _, value := range values {
		var output bytes.Buffer
		fmt.Fprintf(&output, "// %s is the generated definition for %s.\nvar %s = Shape{\nName: %q,\nBoxes: []Box{\n",
			identifiers[value.Name], value.Name, identifiers[value.Name], value.Name)
		for _, box := range value.Boxes {
			fmt.Fprintf(&output, "{%s, %s, %s, %s, %s, %s},\n",
				goFloat(box[0]), goFloat(box[1]), goFloat(box[2]),
				goFloat(box[3]), goFloat(box[4]), goFloat(box[5]))
		}
		output.WriteString("},\n}\n\n")
		fileName := "voxelshape/" + fileNames[value.Name]
		formatted, err := formattedGo(fileName, generatedSource("voxelshape", output.String()))
		if err != nil {
			return nil, 0, err
		}
		files[fileName] = formatted
	}
	var output bytes.Buffer
	output.WriteString("var all = []Shape{\n")
	for _, value := range values {
		fmt.Fprintf(&output, "%s,\n", identifiers[value.Name])
	}
	output.WriteString("}\n")
	const indexName = "voxelshape/index_generated.go"
	formatted, err := formattedGo(indexName, generatedSource("voxelshape", output.String()))
	if err != nil {
		return nil, 0, err
	}
	files[indexName] = formatted
	return files, len(values), nil
}

// voxelGridBoxes validates a packed grid and greedily merges adjacent filled cells.
// The fixed z, y, x expansion order makes the non-overlapping boxes reproducible.
func voxelGridBoxes(grid rawVoxelGrid) ([]generatedBox, error) {
	if grid.Cells == nil {
		return nil, fmt.Errorf("missing cells")
	}
	var dimensions [3]int
	for axis, size := range [3]*int{grid.Cells.XSize, grid.Cells.YSize, grid.Cells.ZSize} {
		if size == nil {
			return nil, fmt.Errorf("axis %d size is missing", axis)
		}
		dimensions[axis] = *size
	}
	coordinates := [3][]float32{grid.XCoordinates, grid.YCoordinates, grid.ZCoordinates}
	count := 1
	for axis, size := range dimensions {
		if size < 0 || len(coordinates[axis]) == 0 || size != len(coordinates[axis])-1 {
			return nil, fmt.Errorf("axis %d size does not match coordinates", axis)
		}
		for index, value := range coordinates[axis] {
			if math.IsNaN(float64(value)) || math.IsInf(float64(value), 0) {
				return nil, fmt.Errorf("axis %d has non-finite coordinates", axis)
			}
			if index > 0 && value <= coordinates[axis][index-1] {
				return nil, fmt.Errorf("axis %d coordinates are not strictly increasing", axis)
			}
		}
		if size > 0 && count > math.MaxInt/size {
			return nil, fmt.Errorf("cell count overflows")
		}
		count *= size
	}
	storage := grid.Cells.Storage
	if storage == nil {
		return nil, fmt.Errorf("storage is missing")
	}
	expectedBytes := count / 8
	if count%8 != 0 {
		expectedBytes++
	}
	if len(storage) != expectedBytes {
		return nil, fmt.Errorf("storage has %d bytes, want %d", len(storage), expectedBytes)
	}
	for _, value := range storage {
		if value > 255 {
			return nil, fmt.Errorf("storage value %d is not a byte", value)
		}
	}
	if count%8 != 0 && storage[len(storage)-1]>>(count%8) != 0 {
		return nil, fmt.Errorf("storage has set padding bits")
	}
	occupied := make([]bool, count)
	for index := range occupied {
		occupied[index] = storage[index/8]&(1<<(index%8)) != 0
	}
	boxes := make([]generatedBox, 0)
	for x := 0; x < dimensions[0]; x++ {
		for y := 0; y < dimensions[1]; y++ {
			for z := 0; z < dimensions[2]; z++ {
				start := [3]int{x, y, z}
				end := [3]int{x + 1, y + 1, z + 1}
				if !voxelRegionFilled(occupied, dimensions, start, end) {
					continue
				}
				for axis := 2; axis >= 0; axis-- {
					for end[axis] < dimensions[axis] {
						end[axis]++
						if !voxelRegionFilled(occupied, dimensions, start, end) {
							end[axis]--
							break
						}
					}
				}
				boxes = append(boxes, generatedBox{
					coordinates[0][x], coordinates[1][y], coordinates[2][z],
					coordinates[0][end[0]], coordinates[1][end[1]], coordinates[2][end[2]],
				})
				for a := x; a < end[0]; a++ {
					for b := y; b < end[1]; b++ {
						for c := z; c < end[2]; c++ {
							occupied[(a*dimensions[1]+b)*dimensions[2]+c] = false
						}
					}
				}
			}
		}
	}
	return boxes, nil
}

// voxelRegionFilled checks a half-open cell region using Cloudburst's z-fast order.
func voxelRegionFilled(occupied []bool, dimensions, start, end [3]int) bool {
	for x := start[0]; x < end[0]; x++ {
		for y := start[1]; y < end[1]; y++ {
			for z := start[2]; z < end[2]; z++ {
				if !occupied[(x*dimensions[1]+y)*dimensions[2]+z] {
					return false
				}
			}
		}
	}
	return true
}
