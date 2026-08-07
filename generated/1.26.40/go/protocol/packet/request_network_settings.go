// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type RequestNetworkSettings struct {
	ClientNetworkVersion int32
}

// Marshal reads or writes RequestNetworkSettings using its canonical wire layout.
func (x *RequestNetworkSettings) Marshal(io protocol.IO) {
	io.BEInt32(&x.ClientNetworkVersion)
}
