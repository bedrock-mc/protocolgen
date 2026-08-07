// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type BiomeConsolidatedFeatureData struct {
	Scatter               BiomeScatterParamData
	Feature               uint16
	Identifier            uint16
	Pass                  uint16
	CanUseInternalFeature bool
}

// Marshal reads or writes BiomeConsolidatedFeatureData using its canonical wire layout.
func (x *BiomeConsolidatedFeatureData) Marshal(io IO) {
	x.Scatter.Marshal(io)
	io.Uint16(&x.Feature)
	io.Uint16(&x.Identifier)
	io.Uint16(&x.Pass)
	io.Bool(&x.CanUseInternalFeature)
}
