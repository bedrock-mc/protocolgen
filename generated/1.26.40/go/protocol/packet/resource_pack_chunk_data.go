// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

// ResourcePackChunkData is sent to the client so that the client can download the resource pack.
// Each packet holds a chunk of the compressed resource pack, of which the size is defined in the
// ResourcePackDataInfo packet sent before.
type ResourcePackChunkData struct {
	// ResourceName is the unique ID of the resource pack that the chunk of data is taken out of.
	ResourceName string
	// ChunkID is the current chunk index of the chunk. It is a number that starts at 0 and is
	// incremented for each resource pack data chunk sent to the client.
	ChunkID uint32
	// ByteOffset is the current progress in bytes or offset in the data that the resource pack data
	// chunk is taken from.
	ByteOffset uint64
	// RawPayload is a byte slice containing a chunk of data from the resource pack. It must be of the
	// same size or less than the DataChunkSize set in the ResourcePackDataInfo packet.
	ChunkData []byte
}

// Marshal reads or writes ResourcePackChunkData using its canonical wire layout.
func (x *ResourcePackChunkData) Marshal(io protocol.IO) {
	io.String(&x.ResourceName)
	io.Uint32(&x.ChunkID)
	protocol.Minimum(io, &x.ChunkID, 0)
	io.Uint64(&x.ByteOffset)
	protocol.Minimum(io, &x.ByteOffset, 0)
	io.Bytes(&x.ChunkData)
}

// ID returns the protocol ID for ResourcePackChunkData.
func (*ResourcePackChunkData) ID() uint32 { return IDResourcePackChunkData }
