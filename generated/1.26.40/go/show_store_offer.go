// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "github.com/google/uuid"

type ShowStoreOffer struct {
	OfferId      uuid.UUID
	RedirectType ShowStoreOfferRedirectType
}

// Marshal reads or writes ShowStoreOffer using its canonical wire layout.
func (x *ShowStoreOffer) Marshal(io IO) {
	io.UUID(&x.OfferId)
	IntegerFunc(&x.RedirectType, io.Uint8)
}
