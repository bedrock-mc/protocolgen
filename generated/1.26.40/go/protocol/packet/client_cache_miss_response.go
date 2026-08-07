// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type ClientCacheMissResponse struct {
	MissingBlobs []protocol.MissingBlobData
}

// Marshal reads or writes ClientCacheMissResponse using its canonical wire layout.
func (x *ClientCacheMissResponse) Marshal(io protocol.IO) {
	protocol.FuncSlice(io, &x.MissingBlobs, io.Varuint32, func(value *protocol.MissingBlobData) {
		value.Marshal(io)
	})
}
