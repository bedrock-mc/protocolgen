// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.44/go/protocol"

// ResourcePackChunkRequest is sent by the client to request a chunk of data from a particular
// resource pack, that it has obtained information about in a ResourcePackDataInfo packet.
type ResourcePackChunkRequest struct {
	// ResourceName is the unique ID of the resource pack that the chunk of data is requested from.
	ResourceName string
	// Chunk is the requested chunk index of the chunk. It is a number that starts at 0 and is
	// incremented for each resource pack data chunk requested.
	Chunk int32
}

// Marshal reads or writes ResourcePackChunkRequest using its canonical wire layout.
func (x *ResourcePackChunkRequest) Marshal(io protocol.IO) {
	io.String(&x.ResourceName)
	protocol.Pattern(io, &x.ResourceName, "A string in the format of <uuid>_<semver>, where <uuid> is a valid UUID and <semver> is a valid semantic version")
	io.Int32(&x.Chunk)
	protocol.Minimum(io, &x.Chunk, 0)
}

// ID returns the protocol ID for ResourcePackChunkRequest.
func (*ResourcePackChunkRequest) ID() uint32 { return IDResourcePackChunkRequest }
