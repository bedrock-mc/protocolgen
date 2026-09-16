// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

// SetHud is sent by the server to set the visibility of individual HUD elements on the client.
type SetHud struct {
	HudElement []protocol.HudElement
	HudVisible protocol.HudVisibility
}

// Marshal reads or writes SetHud using its canonical wire layout.
func (x *SetHud) Marshal(io protocol.IO) {
	protocol.Slice(io, &x.HudElement)
	x.HudVisible.Marshal(io)
}

// ID returns the protocol ID for SetHud.
func (*SetHud) ID() uint32 { return IDSetHud }
