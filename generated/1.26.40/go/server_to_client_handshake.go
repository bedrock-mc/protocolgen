// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type ServerToClientHandshake struct {
	HandshakeWebToken string
}

func (p *ServerToClientHandshake) Encode(w Encoder) error {
	if err := w.Write("ServerToClientHandshakePacket.Handshake WebToken", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, p.HandshakeWebToken); err != nil {
		return err
	}
	return nil
}

func DecodeServerToClientHandshake(r Decoder) (ServerToClientHandshake, error) {
	var p ServerToClientHandshake
	{
		raw, err := r.Read("ServerToClientHandshakePacket.Handshake WebToken", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(string)
		if !ok {
			return p, fmt.Errorf("field ServerToClientHandshakePacket.Handshake WebToken has unexpected decoded type %T", raw)
		}
		p.HandshakeWebToken = value
	}
	return p, nil
}
