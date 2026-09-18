// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.44/go/protocol"
)

// ClientboundCloseForm is sent by the server to clear the entire form stack of the client. This means that
// all forms that are currently open will be closed. This does not affect inventories and other containers.
type ClientboundCloseForm struct {
}

// ID returns the protocol ID for ClientboundCloseForm.
func (*ClientboundCloseForm) ID() uint32 { return IDClientboundCloseForm }

// Marshal reads or writes ClientboundCloseForm using its canonical wire layout.
func (pk *ClientboundCloseForm) Marshal(io protocol.IO) {
}
