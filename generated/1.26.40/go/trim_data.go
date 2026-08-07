// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type TrimData struct {
	TrimPatternList  []TrimPattern
	TrimMaterialList []TrimMaterial
}

// Marshal reads or writes TrimData using its canonical wire layout.
func (x *TrimData) Marshal(io IO) {
	FuncSlice(io, &x.TrimPatternList, io.Varuint32, func(value *TrimPattern) {
		value.Marshal(io)
	})
	FuncSlice(io, &x.TrimMaterialList, io.Varuint32, func(value *TrimMaterial) {
		value.Marshal(io)
	})
}
