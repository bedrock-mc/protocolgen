// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

// RefreshEntitlements is sent by the client to the server to refresh the entitlements of the
// player.
type RefreshEntitlements struct {
}

// Marshal reads or writes RefreshEntitlements using its canonical wire layout.
func (x *RefreshEntitlements) Marshal(io protocol.IO) {
}

// ID returns the protocol ID for RefreshEntitlements.
func (*RefreshEntitlements) ID() uint32 { return IDRefreshEntitlements }
