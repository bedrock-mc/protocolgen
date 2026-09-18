// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.44/go/protocol"
)

// RefreshEntitlements is sent by the client to the server to refresh the entitlements of the player.
type RefreshEntitlements struct {
}

// ID returns the protocol ID for RefreshEntitlements.
func (*RefreshEntitlements) ID() uint32 { return IDRefreshEntitlements }

// Marshal reads or writes RefreshEntitlements using its canonical wire layout.
func (pk *RefreshEntitlements) Marshal(io protocol.IO) {
}
