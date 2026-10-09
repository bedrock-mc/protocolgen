package data_test

import (
	"testing"

	"github.com/bedrock-mc/protocolgen/data"
	"github.com/bedrock-mc/protocolgen/data/biome"
	"github.com/bedrock-mc/protocolgen/data/block"
	"github.com/bedrock-mc/protocolgen/data/entity"
	"github.com/bedrock-mc/protocolgen/data/item"
	"github.com/bedrock-mc/protocolgen/data/voxelshape"
)

// TestGeneratedCountsAndIndexes verifies cross-package data references.
func TestGeneratedCountsAndIndexes(t *testing.T) {
	counts := data.GeneratedCounts
	if got := uint32(block.StateCount()); got != counts.BlockStates {
		t.Fatalf("block state count: got %d, want %d", got, counts.BlockStates)
	}
	if got := uint32(len(biome.All())); got != counts.Biomes {
		t.Fatalf("biome count: got %d, want %d", got, counts.Biomes)
	}
	if got := uint32(len(voxelshape.All())); got != counts.VoxelShapes {
		t.Fatalf("voxel shape count: got %d, want %d", got, counts.VoxelShapes)
	}
	if got := uint32(len(entity.All())); got != counts.Entities {
		t.Fatalf("entity count: got %d, want %d", got, counts.Entities)
	}
	if got := uint32(len(item.Foods())); got != counts.Foods {
		t.Fatalf("food count: got %d, want %d", got, counts.Foods)
	}

	for index := range block.StateCount() {
		state, ok := block.StateAt(index)
		if !ok {
			t.Fatalf("state %d is not addressable", index)
		}
		resolved, ok := block.StateByHash(state.Hash)
		if !ok || resolved != state {
			t.Fatalf("state %d hash %d does not round-trip", index, state.Hash)
		}
		if _, ok := block.BlockAt(state.Block); !ok {
			t.Fatalf("state %d references missing block %d", index, state.Block)
		}
		properties, ok := block.PropertiesAt(state.Properties)
		if !ok {
			t.Fatalf("state %d references missing properties %d", index, state.Properties)
		}
		for _, shape := range []uint16{
			properties.CollisionShape,
			properties.OutlineShape,
			properties.VisualShape,
			properties.UIShape,
			properties.LiquidClipShape,
		} {
			_, available := block.Shape(shape)
			if available != (shape != 0) {
				t.Fatalf("state %d has invalid shape reference %d", index, shape)
			}
		}
	}
}

func TestRepresentativeLookups(t *testing.T) {
	if value, ok := biome.Lookup("minecraft:plains"); !ok || value.Temperature != 0.8 {
		t.Fatalf("unexpected plains biome: %+v, found=%t", value, ok)
	}
	if value, ok := voxelshape.Lookup("minecraft:anvil"); !ok || len(value.Boxes) != 7 {
		t.Fatalf("unexpected anvil shape: %+v, found=%t", value, ok)
	}
	if value, ok := entity.Lookup("minecraft:arrow"); !ok ||
		!value.ProjectileGravity.Present || value.ProjectileGravity.Value != 0.05 {
		t.Fatalf("unexpected arrow entity: %+v, found=%t", value, ok)
	}
	if value, ok := item.LookupFood("minecraft:apple"); !ok || value.Nutrition != 4 {
		t.Fatalf("unexpected apple food: %+v, found=%t", value, ok)
	}
}
