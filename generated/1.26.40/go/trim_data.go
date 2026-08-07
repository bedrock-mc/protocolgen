// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type TrimData struct {
	TrimPatternList  []TrimPattern
	TrimMaterialList []TrimMaterial
}

// Marshal reads or writes TrimData using its canonical wire layout.
func (x *TrimData) Marshal(io IO) {
	if !io.Reading() && uint64(len(x.TrimPatternList)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.TrimPatternList), "collection length overflows uint32")
		return
	}
	count1 := uint32(len(x.TrimPatternList))
	io.Varuint32(&count1)
	if io.Reading() {
		if uint64(count1) > uint64(^uint(0)>>1) {
			io.InvalidValue(count1, "collection length overflows int")
			return
		}
		x.TrimPatternList = make([]TrimPattern, int(count1))
	}
	for index2 := range x.TrimPatternList {
		x.TrimPatternList[index2].Marshal(io)
	}
	if !io.Reading() && uint64(len(x.TrimMaterialList)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.TrimMaterialList), "collection length overflows uint32")
		return
	}
	count3 := uint32(len(x.TrimMaterialList))
	io.Varuint32(&count3)
	if io.Reading() {
		if uint64(count3) > uint64(^uint(0)>>1) {
			io.InvalidValue(count3, "collection length overflows int")
			return
		}
		x.TrimMaterialList = make([]TrimMaterial, int(count3))
	}
	for index4 := range x.TrimMaterialList {
		x.TrimMaterialList[index4].Marshal(io)
	}
}
