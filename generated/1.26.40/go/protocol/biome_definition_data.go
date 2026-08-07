// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type BiomeDefinitionData struct {
	Id                uint16
	Temperature       float32
	Downfall          float32
	FoliageSnow       float32
	Depth             float32
	Scale             float32
	MapWaterColorARGB int32
	Rain              bool
	Tags              Optional[BiomeTagsData]
	ChunkGenData      Optional[BiomeDefinitionChunkGenData]
}

// Marshal reads or writes BiomeDefinitionData using its canonical wire layout.
func (x *BiomeDefinitionData) Marshal(io IO) {
	io.Uint16(&x.Id)
	io.Float32(&x.Temperature)
	io.Float32(&x.Downfall)
	io.Float32(&x.FoliageSnow)
	io.Float32(&x.Depth)
	io.Float32(&x.Scale)
	io.Int32(&x.MapWaterColorARGB)
	io.Bool(&x.Rain)
	OptionalFunc(io, &x.Tags, func(value *BiomeTagsData) {
		value.Marshal(io)
	})
	OptionalFunc(io, &x.ChunkGenData, func(value *BiomeDefinitionChunkGenData) {
		value.Marshal(io)
	})
}
