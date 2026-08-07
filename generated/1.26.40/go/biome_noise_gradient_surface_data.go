// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type BiomeNoiseGradientSurfaceData struct {
	NonReplaceableBlocks []uint32
	GradientBlocks       []SerializedNoiseBlockSpecifier
	Noise                NoiseDescriptor
}

// Marshal reads or writes BiomeNoiseGradientSurfaceData using its canonical wire layout.
func (x *BiomeNoiseGradientSurfaceData) Marshal(io IO) {
	FuncSlice(io, &x.NonReplaceableBlocks, io.Varuint32, io.Uint32)
	FuncSlice(io, &x.GradientBlocks, io.Varuint32, func(value *SerializedNoiseBlockSpecifier) {
		item := *value
		item.Marshal(io)
		*value = item
	})
	x.Noise.Marshal(io)
}
