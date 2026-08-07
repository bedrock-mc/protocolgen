// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type ClientCacheBlobStatus struct {
	MissingIds []uint64
	FoundIds   []uint64
}

// Marshal reads or writes ClientCacheBlobStatus using its canonical wire layout.
func (x *ClientCacheBlobStatus) Marshal(io protocol.IO) {
	protocol.FuncSlice(io, &x.MissingIds, io.Varuint32, io.Uint64)
	protocol.FuncSlice(io, &x.FoundIds, io.Varuint32, io.Uint64)
}

// ID returns the protocol ID for ClientCacheBlobStatus.
func (*ClientCacheBlobStatus) ID() uint32 { return IDClientCacheBlobStatus }
