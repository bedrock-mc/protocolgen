// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type ShowCredits struct {
	PlayerRuntimeID uint64
	CreditsState    int32
}

// Marshal reads or writes ShowCredits using its canonical wire layout.
func (x *ShowCredits) Marshal(io protocol.IO) {
	io.ActorRuntimeID(&x.PlayerRuntimeID)
	io.Varint32(&x.CreditsState)
}

// ID returns the protocol ID for ShowCredits.
func (*ShowCredits) ID() uint32 { return IDShowCredits }
