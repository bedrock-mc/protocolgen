// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type RequestChunkRadius struct {
	ChunkRadius    int32
	MaxChunkRadius uint8
}

// Marshal reads or writes RequestChunkRadius using its canonical wire layout.
func (x *RequestChunkRadius) Marshal(io protocol.IO) {
	io.Varint32(&x.ChunkRadius)
	io.Uint8(&x.MaxChunkRadius)
}
