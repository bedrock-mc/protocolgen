// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type SubChunkPos struct {
	SubchunkPositionX int32
	SubchunkPositionY int32
	SubchunkPositionZ int32
}

// Marshal reads or writes SubChunkPos using its canonical wire layout.
func (x *SubChunkPos) Marshal(io IO) {
	io.Int32(&x.SubchunkPositionX)
	io.Int32(&x.SubchunkPositionY)
	io.Int32(&x.SubchunkPositionZ)
}
