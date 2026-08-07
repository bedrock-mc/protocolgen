// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type ResourcePackClientResponse struct {
	Response protocol.ResourcePackClientResponseResponse
}

// Marshal reads or writes ResourcePackClientResponse using its canonical wire layout.
func (x *ResourcePackClientResponse) Marshal(io protocol.IO) {
	protocol.MarshalResourcePackClientResponseResponse(io, &x.Response)
}
