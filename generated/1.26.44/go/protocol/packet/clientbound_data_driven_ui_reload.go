// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.44/go/protocol"
)

// ClientboundDataDrivenUIReload is sent by the server to reload the data-driven UI on the client.
type ClientboundDataDrivenUIReload struct {
}

// ID returns the protocol ID for ClientboundDataDrivenUIReload.
func (*ClientboundDataDrivenUIReload) ID() uint32 { return IDClientboundDataDrivenUIReload }

// Marshal reads or writes ClientboundDataDrivenUIReload using its canonical wire layout.
func (pk *ClientboundDataDrivenUIReload) Marshal(io protocol.IO) {
}
