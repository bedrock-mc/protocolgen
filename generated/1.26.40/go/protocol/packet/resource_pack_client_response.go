// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

// ResourcePackClientResponse is sent by the client in response to resource packets sent by the
// server. It is used to let the server know what action needs to be taken for the client to have
// all resource packs ready and set.
type ResourcePackClientResponse struct {
	// Response is the response type of the response. It is one of the constants found above.
	Response protocol.ResourcePackClientResponseData
}

// Marshal reads or writes ResourcePackClientResponse using its canonical wire layout.
func (x *ResourcePackClientResponse) Marshal(io protocol.IO) {
	protocol.MarshalResourcePackClientResponseData(io, &x.Response)
}

// ID returns the protocol ID for ResourcePackClientResponse.
func (*ResourcePackClientResponse) ID() uint32 { return IDResourcePackClientResponse }
