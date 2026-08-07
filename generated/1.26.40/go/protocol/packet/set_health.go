// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

// SetHealth is sent by the server. It sets the health of the player it is sent to. The SetHealth
// packet should no longer be used. Instead, the health attribute should be used so that the health
// and maximum health may be changed directly.
type SetHealth struct {
	// Health is the new health of the player.
	Health int32
}

// Marshal reads or writes SetHealth using its canonical wire layout.
func (x *SetHealth) Marshal(io protocol.IO) {
	io.Varint32(&x.Health)
}

// ID returns the protocol ID for SetHealth.
func (*SetHealth) ID() uint32 { return IDSetHealth }
