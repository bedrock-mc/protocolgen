// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.40/go/protocol"
)

// ResourcePackDataInfo is sent by the server to the client to inform the client about the data contained in
// one of the resource packs that are about to be sent.
type ResourcePackDataInfo struct {
	ResourceName   string
	ChunkSize      uint32
	NumberOfChunks uint32
	FileSize       uint64
	FileHash       []byte
	IsPremiumPack  bool
	// PackType is the type of the resource pack. It is one of the resource pack types that may be found in the
	// constants above.
	PackType uint8
}

// ID ...
func (*ResourcePackDataInfo) ID() uint32 {
	return IDResourcePackDataInfo
}

func (pk *ResourcePackDataInfo) Marshal(io protocol.IO) {
	io.String(&pk.ResourceName)
	io.Uint32(&pk.ChunkSize)
	io.Uint32(&pk.NumberOfChunks)
	io.Uint64(&pk.FileSize)
	io.ByteSlice(&pk.FileHash)
	io.Bool(&pk.IsPremiumPack)
	io.Uint8(&pk.PackType)
}
