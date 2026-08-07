// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type BiomeNoiseGradientSurfaceData struct {
	NonReplaceableBlocks []uint32
	GradientBlocks       []SerializedNoiseBlockSpecifier
	Noise                NoiseDescriptor
}

// Marshal reads or writes BiomeNoiseGradientSurfaceData using its canonical wire layout.
func (x *BiomeNoiseGradientSurfaceData) Marshal(io IO) {
	FuncSlice(io, &x.NonReplaceableBlocks, io.Varuint32, io.Uint32)
	Slice(io, &x.GradientBlocks)
	x.Noise.Marshal(io)
}
