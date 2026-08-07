// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type Login struct {
	ClientNetworkVersion int32
	ConnectionRequest    string
}

func (p *Login) Encode(w Encoder) error {
	if err := w.Write("LoginPacket.Client Network Version", Shape{Kind: "primitive", PrimitiveCode: "i32be"}, p.ClientNetworkVersion); err != nil {
		return err
	}
	if err := w.Write("LoginPacket.Connection Request", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, p.ConnectionRequest); err != nil {
		return err
	}
	return nil
}

func DecodeLogin(r Decoder) (Login, error) {
	var p Login
	{
		raw, err := r.Read("LoginPacket.Client Network Version", Shape{Kind: "primitive", PrimitiveCode: "i32be"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(int32)
		if !ok {
			return p, fmt.Errorf("field LoginPacket.Client Network Version has unexpected decoded type %T", raw)
		}
		p.ClientNetworkVersion = value
	}
	{
		raw, err := r.Read("LoginPacket.Connection Request", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(string)
		if !ok {
			return p, fmt.Errorf("field LoginPacket.Connection Request has unexpected decoded type %T", raw)
		}
		p.ConnectionRequest = value
	}
	return p, nil
}
