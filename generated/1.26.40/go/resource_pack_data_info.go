// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type ResourcePackDataInfo struct {
	ResourceName   string
	ChunkSize      uint32
	NumberOfChunks uint32
	FileSize       uint64
	FileHash       string
	IsPremiumPack  bool
	PackType       uint8
}

// Marshal reads or writes ResourcePackDataInfo using its canonical wire layout.
func (x *ResourcePackDataInfo) Marshal(io IO) {
	io.String(&x.ResourceName)
	io.Uint32(&x.ChunkSize)
	io.Uint32(&x.NumberOfChunks)
	io.Uint64(&x.FileSize)
	io.String(&x.FileHash)
	io.Bool(&x.IsPremiumPack)
	io.Uint8(&x.PackType)
}
