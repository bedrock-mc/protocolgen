// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import (
	"fmt"

	"github.com/google/uuid"
)

type ShowStoreOffer struct {
	OfferId      uuid.UUID
	RedirectType ShowStoreOfferRedirectType
}

func (p *ShowStoreOffer) Encode(w Encoder) error {
	if err := w.Write("ShowStoreOfferPacket.Offer Id", Shape{Kind: "primitive", Semantic: "mce::UUID", TypeID: "mce__UUID.json#", PrimitiveCode: "uuid"}, p.OfferId); err != nil {
		return err
	}
	if err := w.Write("ShowStoreOfferPacket.Redirect Type", Shape{Kind: "enum", Semantic: "ShowStoreOfferRedirectType", TypeID: "enums/ShowStoreOfferRedirectType", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "MarketplaceOffer", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "DressingRoomOffer", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "ThirdPartyServerPage", Shape: Shape{Kind: "void"}}}}, p.RedirectType); err != nil {
		return err
	}
	return nil
}

func DecodeShowStoreOffer(r Decoder) (ShowStoreOffer, error) {
	var p ShowStoreOffer
	{
		raw, err := r.Read("ShowStoreOfferPacket.Offer Id", Shape{Kind: "primitive", Semantic: "mce::UUID", TypeID: "mce__UUID.json#", PrimitiveCode: "uuid"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(uuid.UUID)
		if !ok {
			return p, fmt.Errorf("field ShowStoreOfferPacket.Offer Id has unexpected decoded type %T", raw)
		}
		p.OfferId = value
	}
	{
		raw, err := r.Read("ShowStoreOfferPacket.Redirect Type", Shape{Kind: "enum", Semantic: "ShowStoreOfferRedirectType", TypeID: "enums/ShowStoreOfferRedirectType", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "MarketplaceOffer", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "DressingRoomOffer", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "ThirdPartyServerPage", Shape: Shape{Kind: "void"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(ShowStoreOfferRedirectType)
		if !ok {
			return p, fmt.Errorf("field ShowStoreOfferPacket.Redirect Type has unexpected decoded type %T", raw)
		}
		p.RedirectType = value
	}
	return p, nil
}
