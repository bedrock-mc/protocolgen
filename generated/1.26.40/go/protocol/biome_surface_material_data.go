// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type BiomeSurfaceMaterialData struct {
	TopBlock        uint32
	MidBlock        uint32
	SeaFloorBlock   uint32
	FoundationBlock uint32
	SeaBlock        uint32
	SeaFloorDepth   int32
}

// Marshal reads or writes BiomeSurfaceMaterialData using its canonical wire layout.
func (x *BiomeSurfaceMaterialData) Marshal(io IO) {
	io.Uint32(&x.TopBlock)
	io.Uint32(&x.MidBlock)
	io.Uint32(&x.SeaFloorBlock)
	io.Uint32(&x.FoundationBlock)
	io.Uint32(&x.SeaBlock)
	io.Int32(&x.SeaFloorDepth)
}
