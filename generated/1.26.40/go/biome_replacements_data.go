// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type BiomeReplacementsData struct {
	BiomeReplacements []BiomeReplacementData
}

// Marshal reads or writes BiomeReplacementsData using its canonical wire layout.
func (x *BiomeReplacementsData) Marshal(io IO) {
	FuncSlice(io, &x.BiomeReplacements, io.Varuint32, func(value *BiomeReplacementData) {
		item := *value
		item.Marshal(io)
		*value = item
	})
}
