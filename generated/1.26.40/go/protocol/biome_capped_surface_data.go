// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type BiomeCappedSurfaceData struct {
	FloorBlocks     []uint32
	CeilingBlocks   []uint32
	SeaBlock        Optional[uint32]
	FoundationBlock Optional[uint32]
	BeachBlock      Optional[uint32]
}

// Marshal reads or writes BiomeCappedSurfaceData using its canonical wire layout.
func (x *BiomeCappedSurfaceData) Marshal(io IO) {
	FuncSlice(io, &x.FloorBlocks, io.Varuint32, io.Uint32)
	FuncSlice(io, &x.CeilingBlocks, io.Varuint32, io.Uint32)
	OptionalFunc(io, &x.SeaBlock, io.Uint32)
	OptionalFunc(io, &x.FoundationBlock, io.Uint32)
	OptionalFunc(io, &x.BeachBlock, io.Uint32)
}
