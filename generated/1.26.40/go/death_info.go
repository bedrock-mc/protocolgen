// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type DeathInfo struct {
	DeathCauseAttackName  string
	DeathCauseMessageList []string
}

func (p *DeathInfo) Encode(w Encoder) error {
	if err := w.Write("DeathInfoPacket.Death Cause Attack Name", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, p.DeathCauseAttackName); err != nil {
		return err
	}
	if err := w.Write("DeathInfoPacket.Death Cause Message List", Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, p.DeathCauseMessageList); err != nil {
		return err
	}
	return nil
}

func DecodeDeathInfo(r Decoder) (DeathInfo, error) {
	var p DeathInfo
	{
		raw, err := r.Read("DeathInfoPacket.Death Cause Attack Name", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(string)
		if !ok {
			return p, fmt.Errorf("field DeathInfoPacket.Death Cause Attack Name has unexpected decoded type %T", raw)
		}
		p.DeathCauseAttackName = value
	}
	{
		raw, err := r.Read("DeathInfoPacket.Death Cause Message List", Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.([]string)
		if !ok {
			return p, fmt.Errorf("field DeathInfoPacket.Death Cause Message List has unexpected decoded type %T", raw)
		}
		p.DeathCauseMessageList = value
	}
	return p, nil
}
