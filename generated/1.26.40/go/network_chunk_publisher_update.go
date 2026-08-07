// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type NetworkChunkPublisherUpdate struct {
	NewPositionForView    BlockPos
	NewRadiusForView      uint32
	ServerBuiltChunksList []ChunkPos
}

// Marshal reads or writes NetworkChunkPublisherUpdate using its canonical wire layout.
func (x *NetworkChunkPublisherUpdate) Marshal(io IO) {
	x.NewPositionForView.Marshal(io)
	io.Varuint32(&x.NewRadiusForView)
	FuncSlice(io, &x.ServerBuiltChunksList, io.Uint32, func(value *ChunkPos) {
		value.Marshal(io)
	})
}
