// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type ResourcePackChunkRequest struct {
	ResourceName string
	Chunk        int32
}

// Marshal reads or writes ResourcePackChunkRequest using its canonical wire layout.
func (x *ResourcePackChunkRequest) Marshal(io protocol.IO) {
	io.String(&x.ResourceName)
	io.Int32(&x.Chunk)
}

// ID returns the protocol ID for ResourcePackChunkRequest.
func (*ResourcePackChunkRequest) ID() uint32 { return IDResourcePackChunkRequest }
