// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.50/go/protocol"
)

// SetTitle is sent by the server to make a title, subtitle or action bar shown to a player. It has several
// fields that allow setting the duration of the titles.
type SetTitle struct {
	TitleType   protocol.TitleType
	TitleText   string
	FadeInTime  int32
	StayTime    int32
	FadeOutTime int32
	// Xuid is the XBOX Live user ID of the player, which will remain consistent as long as the player is logged
	// in with the XBOX Live account. It is empty if the user is not logged into its XBL account.
	Xuid string
	// PlatformOnlineID is either a uint64 or an empty string.
	PlatformOnlineID     string
	FilteredTitleMessage string
}

// ID returns the protocol ID for SetTitle.
func (*SetTitle) ID() uint32 { return IDSetTitle }

// Marshal reads or writes SetTitle using its canonical wire layout.
func (pk *SetTitle) Marshal(io protocol.IO) {
	pk.TitleType.Marshal(io)
	io.String(&pk.TitleText)
	io.Varint32(&pk.FadeInTime)
	io.Varint32(&pk.StayTime)
	io.Varint32(&pk.FadeOutTime)
	io.String(&pk.Xuid)
	io.String(&pk.PlatformOnlineID)
	io.String(&pk.FilteredTitleMessage)
}
