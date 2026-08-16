// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.44/go/protocol"

type ClientboundCloseForm struct {
}

// Marshal reads or writes ClientboundCloseForm using its canonical wire layout.
func (x *ClientboundCloseForm) Marshal(io protocol.IO) {
}

// ID returns the protocol ID for ClientboundCloseForm.
func (*ClientboundCloseForm) ID() uint32 { return IDClientboundCloseForm }
