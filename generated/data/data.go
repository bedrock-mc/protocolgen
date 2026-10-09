// Package data reports the target release, source versions and catalog counts.
package data

// Counts reports the number of records in each generated dataset.
type Counts struct {
	BlockStates uint32 `json:"block_states"`
	Biomes      uint32 `json:"biomes"`
	VoxelShapes uint32 `json:"voxel_shapes"`
	Entities    uint32 `json:"entities"`
	Foods       uint32 `json:"foods"`
}
