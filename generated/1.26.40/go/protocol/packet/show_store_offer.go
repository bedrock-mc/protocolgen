// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.40/go/protocol"

	"github.com/google/uuid"
)

type ShowStoreOffer struct {
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
