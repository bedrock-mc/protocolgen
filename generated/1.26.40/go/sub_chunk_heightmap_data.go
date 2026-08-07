// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

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
		item := *value
		for index1 := range item {
			for index2 := range item[index1] {
				io.Int8(&item[index1][index2])
			}
		}
		*value = item
	})
	IntegerFunc(&x.RenderHeightMapType, io.Uint8)
	OptionalFunc(io, &x.SubchunkRenderHeightMap, func(value *[16][16]int8) {
		item := *value
		for index3 := range item {
			for index4 := range item[index3] {
				io.Int8(&item[index3][index4])
			}
		}
		*value = item
	})
}
