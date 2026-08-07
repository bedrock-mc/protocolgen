// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type HeightmapData struct {
	HeightMapType           HeightMapDataType
	SubchunkHeightMap       Optional[[16][16]int8]
	RenderHeightMapType     HeightMapDataType
	SubchunkRenderHeightMap Optional[[16][16]int8]
}

// Marshal reads or writes HeightmapData using its canonical wire layout.
func (x *HeightmapData) Marshal(io IO) {
	IntegerFunc(&x.HeightMapType, io.Uint8)
	OptionalFunc(io, &x.SubchunkHeightMap, func(value *[16][16]int8) {
		for index1 := range *value {
			for index2 := range *&(*value)[index1] {
				io.Int8(&(*&(*value)[index1])[index2])
			}
		}
	})
	IntegerFunc(&x.RenderHeightMapType, io.Uint8)
	OptionalFunc(io, &x.SubchunkRenderHeightMap, func(value *[16][16]int8) {
		for index3 := range *value {
			for index4 := range *&(*value)[index3] {
				io.Int8(&(*&(*value)[index3])[index4])
			}
		}
	})
}
