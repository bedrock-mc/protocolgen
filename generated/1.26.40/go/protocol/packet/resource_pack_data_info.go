// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

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
func (x *ResourcePackDataInfo) Marshal(io protocol.IO) {
	io.String(&x.ResourceName)
	io.Uint32(&x.ChunkSize)
	io.Uint32(&x.NumberOfChunks)
	io.Uint64(&x.FileSize)
	io.String(&x.FileHash)
	io.Bool(&x.IsPremiumPack)
	io.Uint8(&x.PackType)
}

// ID returns the protocol ID for ResourcePackDataInfo.
func (*ResourcePackDataInfo) ID() uint32 { return IDResourcePackDataInfo }
