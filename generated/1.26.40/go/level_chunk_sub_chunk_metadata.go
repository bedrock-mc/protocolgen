// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type LevelChunkSubChunkMetadata struct {
	BlobId uint64
}

// Marshal reads or writes LevelChunkSubChunkMetadata using its canonical wire layout.
func (x *LevelChunkSubChunkMetadata) Marshal(io IO) {
	io.Uint64(&x.BlobId)
}
