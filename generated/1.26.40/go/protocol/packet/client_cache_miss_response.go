// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

// ClientCacheMissResponse is part of the blob cache protocol. It is sent by the server in response
// to a ClientCacheBlobStatus packet and contains the blob data of all blobs that the client
// acknowledged not to have yet.
type ClientCacheMissResponse struct {
	// MissingBlobs is a list of all blobs that the client sent misses for in the ClientCacheBlobStatus.
	// These blobs hold the data of the blobs with the hashes they are matched with.
	MissingBlobs []protocol.MissingBlobData
}

// Marshal reads or writes ClientCacheMissResponse using its canonical wire layout.
func (x *ClientCacheMissResponse) Marshal(io protocol.IO) {
	protocol.SliceLimits(io, &x.MissingBlobs, 0, 4095)
}

// ID returns the protocol ID for ClientCacheMissResponse.
func (*ClientCacheMissResponse) ID() uint32 { return IDClientCacheMissResponse }
