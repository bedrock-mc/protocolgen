// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

// ClientCacheBlobStatus is part of the blob cache protocol. It is sent by the client to let the
// server know what blobs it needs and which blobs it already has, in an ACK type system.
type ClientCacheBlobStatus struct {
	// MissHashes is a list of blob hashes that the client does not have a blob available for. The
	// server should send the blobs matching these hashes as soon as possible.
	MissingIds []uint64
	// HitHashes is a list of blob hashes that the client has a blob available for. The blobs hashes
	// here mean that the client already has them: The server does not need to send the blobs anymore.
	FoundIds []uint64
}

// Marshal reads or writes ClientCacheBlobStatus using its canonical wire layout.
func (x *ClientCacheBlobStatus) Marshal(io protocol.IO) {
	protocol.FuncSlice(io, &x.MissingIds, io.Varuint32, io.Uint64)
	protocol.FuncSlice(io, &x.FoundIds, io.Varuint32, io.Uint64)
}

// ID returns the protocol ID for ClientCacheBlobStatus.
func (*ClientCacheBlobStatus) ID() uint32 { return IDClientCacheBlobStatus }
