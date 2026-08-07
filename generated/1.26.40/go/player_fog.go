// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type PlayerFog struct {
	FogStack []string
}

func (p *PlayerFog) Encode(w Encoder) error {
	if err := w.Write("PlayerFogPacket.Fog Stack", Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, p.FogStack); err != nil {
		return err
	}
	return nil
}

func DecodePlayerFog(r Decoder) (PlayerFog, error) {
	var p PlayerFog
	{
		raw, err := r.Read("PlayerFogPacket.Fog Stack", Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.([]string)
		if !ok {
			return p, fmt.Errorf("field PlayerFogPacket.Fog Stack has unexpected decoded type %T", raw)
		}
		p.FogStack = value
	}
	return p, nil
}
