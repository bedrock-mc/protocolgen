// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type NetworkStackLatency struct {
	CreationTime uint64
	IsFromServer bool
}

// Marshal reads or writes NetworkStackLatency using its canonical wire layout.
func (x *NetworkStackLatency) Marshal(io protocol.IO) {
	io.Uint64(&x.CreationTime)
	io.Bool(&x.IsFromServer)
}
