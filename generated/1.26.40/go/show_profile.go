// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type ShowProfile struct {
	PlayerXUID string
}

func (p *ShowProfile) Encode(w Encoder) error {
	if err := w.Write("ShowProfilePacket.Player XUID", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, p.PlayerXUID); err != nil {
		return err
	}
	return nil
}

func DecodeShowProfile(r Decoder) (ShowProfile, error) {
	var p ShowProfile
	{
		raw, err := r.Read("ShowProfilePacket.Player XUID", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(string)
		if !ok {
			return p, fmt.Errorf("field ShowProfilePacket.Player XUID has unexpected decoded type %T", raw)
		}
		p.PlayerXUID = value
	}
	return p, nil
}
