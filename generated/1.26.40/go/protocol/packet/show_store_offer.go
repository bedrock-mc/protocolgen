// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.40/go/protocol"

	"github.com/google/uuid"
)

// ShowStoreOffer is sent by the server to show a Marketplace store offer to a player. It opens a
// window client-side that displays the item. The ShowStoreOffer packet only works on the partnered
// servers: Servers that are not partnered will not have a store buttons show up in the in-game
// pause menu and will, as a result, not be able to open store offers on the client side. Sending
// the packet does therefore not work when using a proxy that is not connected to with the domain of
// one of the partnered servers.
type ShowStoreOffer struct {
	// OfferID is a UUID that identifies the offer for which a window should be opened.
	OfferID      uuid.UUID
	RedirectType protocol.ShowStoreOfferRedirectType
}

// Marshal reads or writes ShowStoreOffer using its canonical wire layout.
func (x *ShowStoreOffer) Marshal(io protocol.IO) {
	io.UUID(&x.OfferID)
	protocol.IntegerFunc(&x.RedirectType, io.Uint8)
}

// ID returns the protocol ID for ShowStoreOffer.
func (*ShowStoreOffer) ID() uint32 { return IDShowStoreOffer }
