// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.50/go/protocol"
)

type ClientboundControlSchemeSet struct {
	ControlScheme protocol.ControlScheme
}

// ID returns the protocol ID for ClientboundControlSchemeSet.
func (*ClientboundControlSchemeSet) ID() uint32 { return IDClientboundControlSchemeSet }

// Marshal reads or writes ClientboundControlSchemeSet using its canonical wire layout.
func (pk *ClientboundControlSchemeSet) Marshal(io protocol.IO) {
	pk.ControlScheme.Marshal(io)
}
