// Package biome provides generated Bedrock biome properties.
package biome

import "sort"

// Biome contains static climate and presentation properties for a biome.
type Biome struct {
	Name              string
	ID                int32
	HasID             bool
	Temperature       float32
	Downfall          float32
	Depth             float32
	Scale             float32
	RedSporeDensity   float32
	BlueSporeDensity  float32
	AshDensity        float32
	WhiteAshDensity   float32
	FoliageSnow       float32
	MapWaterColorRGBA uint32
	Rain              bool
	Tags              []string
}

// All returns all biome specifications in name order. The returned slice and
// nested tag slices must be treated as immutable.
func All() []Biome {
	return all
}

// Lookup finds a biome specification by its namespaced identifier.
func Lookup(name string) (Biome, bool) {
	index := sort.Search(len(all), func(i int) bool { return all[i].Name >= name })
	if index == len(all) || all[index].Name != name {
		return Biome{}, false
	}
	return all[index], true
}

var byID = func() map[int32]Biome {
	lookup := make(map[int32]Biome, len(all))
	for _, biome := range all {
		if biome.HasID {
			lookup[biome.ID] = biome
		}
	}
	return lookup
}()

// LookupID finds a biome by its numeric ID in the target retail release.
func LookupID(id int32) (Biome, bool) {
	biome, ok := byID[id]
	return biome, ok
}
