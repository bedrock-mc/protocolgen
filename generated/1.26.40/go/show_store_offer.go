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
	enumValue1 := uint8(x.RedirectType)
	io.Uint8(&enumValue1)
	x.RedirectType = ShowStoreOfferRedirectType(enumValue1)
	switch int64(enumValue1) {
	case 0, 1, 2:
	default:
		io.InvalidValue(enumValue1, "unknown enum value")
	}
}
