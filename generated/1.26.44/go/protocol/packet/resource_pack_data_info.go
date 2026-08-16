// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.44/go/protocol"

// ResourcePackDataInfo is sent by the server to the client to inform the client about the data
// contained in one of the resource packs that are about to be sent.
type ResourcePackDataInfo struct {
	ResourceName   string
	ChunkSize      uint32
	NumberOfChunks uint32
	FileSize       uint64
	FileHash       []byte
	IsPremiumPack  bool
	// PackType is the type of the resource pack. It is one of the resource pack types that may be found
	// in the constants above.
	PackType uint8
}

// Marshal reads or writes ResourcePackDataInfo using its canonical wire layout.
func (x *ResourcePackDataInfo) Marshal(io protocol.IO) {
	io.String(&x.ResourceName)
	io.Uint32(&x.ChunkSize)
	protocol.Minimum(io, &x.ChunkSize, 0)
	io.Uint32(&x.NumberOfChunks)
	protocol.Minimum(io, &x.NumberOfChunks, 0)
	io.Uint64(&x.FileSize)
	protocol.Minimum(io, &x.FileSize, 0)
	io.Bytes(&x.FileHash)
	io.Bool(&x.IsPremiumPack)
	io.Uint8(&x.PackType)
	protocol.Minimum(io, &x.PackType, 0)
	protocol.Maximum(io, &x.PackType, 255)
}

// ID returns the protocol ID for ResourcePackDataInfo.
func (*ResourcePackDataInfo) ID() uint32 { return IDResourcePackDataInfo }
