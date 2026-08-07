// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type PartyDestinationCookieResponse struct {
	Cookie   string
	Accepted bool
}

func (p *PartyDestinationCookieResponse) Encode(w Encoder) error {
	if err := w.Write("PartyDestinationCookieResponsePacket.cookie", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, p.Cookie); err != nil {
		return err
	}
	if err := w.Write("PartyDestinationCookieResponsePacket.accepted", Shape{Kind: "primitive", PrimitiveCode: "bool"}, p.Accepted); err != nil {
		return err
	}
	return nil
}

func DecodePartyDestinationCookieResponse(r Decoder) (PartyDestinationCookieResponse, error) {
	var p PartyDestinationCookieResponse
	{
		raw, err := r.Read("PartyDestinationCookieResponsePacket.cookie", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(string)
		if !ok {
			return p, fmt.Errorf("field PartyDestinationCookieResponsePacket.cookie has unexpected decoded type %T", raw)
		}
		p.Cookie = value
	}
	{
		raw, err := r.Read("PartyDestinationCookieResponsePacket.accepted", Shape{Kind: "primitive", PrimitiveCode: "bool"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(bool)
		if !ok {
			return p, fmt.Errorf("field PartyDestinationCookieResponsePacket.accepted has unexpected decoded type %T", raw)
		}
		p.Accepted = value
	}
	return p, nil
}
