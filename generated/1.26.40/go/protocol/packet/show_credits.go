// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

// ShowCredits is sent by the server to show the Minecraft credits screen to the client. It is
// typically sent when the player beats the ender dragon and leaves the End.
type ShowCredits struct {
	// PlayerRuntimeID is the entity runtime ID of the player to show the credits to. It's not clear why
	// this field is actually here in the first place.
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
