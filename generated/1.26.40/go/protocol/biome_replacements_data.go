// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type BiomeReplacementsData struct {
	BiomeReplacements []BiomeReplacementData
}

// Marshal reads or writes BiomeReplacementsData using its canonical wire layout.
func (x *BiomeReplacementsData) Marshal(io IO) {
	Slice(io, &x.BiomeReplacements)
}
