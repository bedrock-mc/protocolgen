// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type SubClientLogin struct {
	SubClientConnectionRequest string
}

func (p *SubClientLogin) Encode(w Encoder) error {
	if err := w.Write("SubClientLoginPacket.Sub Client Connection Request", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, p.SubClientConnectionRequest); err != nil {
		return err
	}
	return nil
}

func DecodeSubClientLogin(r Decoder) (SubClientLogin, error) {
	var p SubClientLogin
	{
		raw, err := r.Read("SubClientLoginPacket.Sub Client Connection Request", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(string)
		if !ok {
			return p, fmt.Errorf("field SubClientLoginPacket.Sub Client Connection Request has unexpected decoded type %T", raw)
		}
		p.SubClientConnectionRequest = value
	}
	return p, nil
}
