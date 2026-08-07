// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type ChunkPos struct {
	X int32
	Z int32
}

// Marshal reads or writes ChunkPos using its canonical wire layout.
func (x *ChunkPos) Marshal(io IO) {
	io.Varint32(&x.X)
	io.Varint32(&x.Z)
}
