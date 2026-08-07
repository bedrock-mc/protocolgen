// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type ServerSettingsResponse struct {
	FormID     uint32
	FormUIJSON string
}

func (p *ServerSettingsResponse) Encode(w Encoder) error {
	if err := w.Write("ServerSettingsResponsePacket.Form ID", Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, p.FormID); err != nil {
		return err
	}
	if err := w.Write("ServerSettingsResponsePacket.Form UI JSON", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, p.FormUIJSON); err != nil {
		return err
	}
	return nil
}

func DecodeServerSettingsResponse(r Decoder) (ServerSettingsResponse, error) {
	var p ServerSettingsResponse
	{
		raw, err := r.Read("ServerSettingsResponsePacket.Form ID", Shape{Kind: "primitive", PrimitiveCode: "var_u32"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(uint32)
		if !ok {
			return p, fmt.Errorf("field ServerSettingsResponsePacket.Form ID has unexpected decoded type %T", raw)
		}
		p.FormID = value
	}
	{
		raw, err := r.Read("ServerSettingsResponsePacket.Form UI JSON", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(string)
		if !ok {
			return p, fmt.Errorf("field ServerSettingsResponsePacket.Form UI JSON has unexpected decoded type %T", raw)
		}
		p.FormUIJSON = value
	}
	return p, nil
}
