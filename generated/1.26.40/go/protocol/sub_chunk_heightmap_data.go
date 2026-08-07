// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type SubChunkHeightmapData struct {
	HeightMapType           SubChunkHeightMapDataType
	SubchunkHeightMap       Optional[[16][16]int8]
	RenderHeightMapType     SubChunkHeightMapDataType
	SubchunkRenderHeightMap Optional[[16][16]int8]
}

// Marshal reads or writes SubChunkHeightmapData using its canonical wire layout.
func (x *SubChunkHeightmapData) Marshal(io IO) {
	IntegerFunc(&x.HeightMapType, io.Uint8)
	OptionalFunc(io, &x.SubchunkHeightMap, func(value *[16][16]int8) {
		for index1 := range *value {
			item2 := &(*value)[index1]
			for index3 := range *item2 {
				item4 := &(*item2)[index3]
				io.Int8(item4)
			}
		}
	})
	IntegerFunc(&x.RenderHeightMapType, io.Uint8)
	OptionalFunc(io, &x.SubchunkRenderHeightMap, func(value *[16][16]int8) {
		for index5 := range *value {
			item6 := &(*value)[index5]
			for index7 := range *item6 {
				item8 := &(*item6)[index7]
				io.Int8(item8)
			}
		}
	})
}
