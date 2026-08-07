// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type ChunkRadiusUpdated struct {
	ChunkRadius int32
}

// Marshal reads or writes ChunkRadiusUpdated using its canonical wire layout.
func (x *ChunkRadiusUpdated) Marshal(io IO) {
	io.Varint32(&x.ChunkRadius)
}
