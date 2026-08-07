// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type ClientCacheMissResponse struct {
	MissingBlobs []protocol.MissingBlobData
}

// Marshal reads or writes ClientCacheMissResponse using its canonical wire layout.
func (x *ClientCacheMissResponse) Marshal(io protocol.IO) {
	protocol.Slice(io, &x.MissingBlobs)
}

// ID returns the protocol ID for ClientCacheMissResponse.
func (*ClientCacheMissResponse) ID() uint32 { return IDClientCacheMissResponse }
