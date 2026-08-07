// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type ClientCacheStatus struct {
	IsCacheSupported bool
}

// Marshal reads or writes ClientCacheStatus using its canonical wire layout.
func (x *ClientCacheStatus) Marshal(io protocol.IO) {
	io.Bool(&x.IsCacheSupported)
}

// ID returns the protocol ID for ClientCacheStatus.
func (*ClientCacheStatus) ID() uint32 { return IDClientCacheStatus }
