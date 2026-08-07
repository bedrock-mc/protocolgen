// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type ResourcePackChunkData struct {
	ResourceName string
	ChunkID      uint32
	ByteOffset   uint64
	ChunkData    []byte
}

// Marshal reads or writes ResourcePackChunkData using its canonical wire layout.
func (x *ResourcePackChunkData) Marshal(io protocol.IO) {
	io.String(&x.ResourceName)
	io.Uint32(&x.ChunkID)
	io.Uint64(&x.ByteOffset)
	io.Bytes(&x.ChunkData)
}
