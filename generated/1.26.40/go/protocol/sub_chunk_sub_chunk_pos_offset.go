// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type SubChunkSubChunkPosOffset struct {
	SubchunkOffsetX int8
	SubchunkOffsetY int8
	SubchunkOffsetZ int8
}

// Marshal reads or writes SubChunkSubChunkPosOffset using its canonical wire layout.
func (x *SubChunkSubChunkPosOffset) Marshal(io IO) {
	io.Int8(&x.SubchunkOffsetX)
	io.Int8(&x.SubchunkOffsetY)
	io.Int8(&x.SubchunkOffsetZ)
}
