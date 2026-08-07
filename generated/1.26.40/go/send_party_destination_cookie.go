// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type SendPartyDestinationCookie struct {
	Cookie          string
	Intent          string
	DestinationName string
}

func (p *SendPartyDestinationCookie) Encode(w Encoder) error {
	if err := w.Write("SendPartyDestinationCookiePacket.cookie", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, p.Cookie); err != nil {
		return err
	}
	if err := w.Write("SendPartyDestinationCookiePacket.intent", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, p.Intent); err != nil {
		return err
	}
	if err := w.Write("SendPartyDestinationCookiePacket.destination_name", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, p.DestinationName); err != nil {
		return err
	}
	return nil
}

func DecodeSendPartyDestinationCookie(r Decoder) (SendPartyDestinationCookie, error) {
	var p SendPartyDestinationCookie
	{
		raw, err := r.Read("SendPartyDestinationCookiePacket.cookie", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(string)
		if !ok {
			return p, fmt.Errorf("field SendPartyDestinationCookiePacket.cookie has unexpected decoded type %T", raw)
		}
		p.Cookie = value
	}
	{
		raw, err := r.Read("SendPartyDestinationCookiePacket.intent", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(string)
		if !ok {
			return p, fmt.Errorf("field SendPartyDestinationCookiePacket.intent has unexpected decoded type %T", raw)
		}
		p.Intent = value
	}
	{
		raw, err := r.Read("SendPartyDestinationCookiePacket.destination_name", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(string)
		if !ok {
			return p, fmt.Errorf("field SendPartyDestinationCookiePacket.destination_name has unexpected decoded type %T", raw)
		}
		p.DestinationName = value
	}
	return p, nil
}
