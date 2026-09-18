// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.50/go/protocol"
)

// SetHud is sent by the server to set the visibility of individual HUD elements on the client.
type SetHud struct {
	HudElement []protocol.HudElement
	HudVisible protocol.HudVisibility
}

// ID ...
func (*SetHud) ID() uint32 {
	return IDSetHud
}

func (pk *SetHud) Marshal(io protocol.IO) {
	protocol.Slice(io, &pk.HudElement)
	pk.HudVisible.Marshal(io)
}
