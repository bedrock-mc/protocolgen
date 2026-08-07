// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type PartyChanged struct {
	PartyInfo *PlayerPartyInfo
}

func (p *PartyChanged) Encode(w Encoder) error {
	if err := w.Write("PartyChangedPacket.party_info", Shape{Kind: "optional", Value: &Shape{Kind: "struct", Semantic: "PlayerPartyInfo", TypeID: "PlayerPartyInfo", Fields: []ShapeField{{Ordinal: 0, Name: "party_id", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 1, Name: "is_party_leader", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}}}}, p.PartyInfo); err != nil {
		return err
	}
	return nil
}

func DecodePartyChanged(r Decoder) (PartyChanged, error) {
	var p PartyChanged
	{
		raw, err := r.Read("PartyChangedPacket.party_info", Shape{Kind: "optional", Value: &Shape{Kind: "struct", Semantic: "PlayerPartyInfo", TypeID: "PlayerPartyInfo", Fields: []ShapeField{{Ordinal: 0, Name: "party_id", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 1, Name: "is_party_leader", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(*PlayerPartyInfo)
		if !ok {
			return p, fmt.Errorf("field PartyChangedPacket.party_info has unexpected decoded type %T", raw)
		}
		p.PartyInfo = value
	}
	return p, nil
}
