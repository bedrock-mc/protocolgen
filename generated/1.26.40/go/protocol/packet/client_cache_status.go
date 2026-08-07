// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

// ClientCacheStatus is sent by the client to the server at the start of the game. It is sent to let
// the server know if it supports the client-side blob cache. Clients such as Nintendo Switch do not
// support the cache, and attempting to use it anyway will fail.
type ClientCacheStatus struct {
	// Enabled specifies if the blob cache is enabled. If false, the server should not attempt to use
	// the blob cache. If true, it may do so, but it may also choose not to use it.
	IsCacheSupported bool
}

// Marshal reads or writes ClientCacheStatus using its canonical wire layout.
func (x *ClientCacheStatus) Marshal(io protocol.IO) {
	io.Bool(&x.IsCacheSupported)
}

// ID returns the protocol ID for ClientCacheStatus.
func (*ClientCacheStatus) ID() uint32 { return IDClientCacheStatus }
