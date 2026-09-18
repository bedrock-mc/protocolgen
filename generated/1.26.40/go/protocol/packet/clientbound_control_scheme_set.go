// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.40/go/protocol"
)

type ClientboundControlSchemeSet struct {
	ControlScheme protocol.ControlScheme
}

// ID ...
func (*ClientboundControlSchemeSet) ID() uint32 {
	return IDClientboundControlSchemeSet
}

func (pk *ClientboundControlSchemeSet) Marshal(io protocol.IO) {
	pk.ControlScheme.Marshal(io)
}
