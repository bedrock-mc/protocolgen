// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type BiomeConsolidatedFeaturesData struct {
	Features []BiomeConsolidatedFeatureData
}

// Marshal reads or writes BiomeConsolidatedFeaturesData using its canonical wire layout.
func (x *BiomeConsolidatedFeaturesData) Marshal(io IO) {
	Slice(io, &x.Features)
}
