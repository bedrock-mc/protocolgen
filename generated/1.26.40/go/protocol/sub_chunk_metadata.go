// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type SubChunkMetadata struct {
	BlobID uint64
}

// Marshal reads or writes SubChunkMetadata using its canonical wire layout.
func (x *SubChunkMetadata) Marshal(io IO) {
	io.Uint64(&x.BlobID)
}
