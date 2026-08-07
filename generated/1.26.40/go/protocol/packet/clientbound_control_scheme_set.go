// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type ClientboundControlSchemeSet struct {
	ControlScheme protocol.ControlSchemeScheme
}

// Marshal reads or writes ClientboundControlSchemeSet using its canonical wire layout.
func (x *ClientboundControlSchemeSet) Marshal(io protocol.IO) {
	protocol.IntegerFunc(&x.ControlScheme, io.Uint8)
}

// ID returns the protocol ID for ClientboundControlSchemeSet.
func (*ClientboundControlSchemeSet) ID() uint32 { return IDClientboundControlSchemeSet }
