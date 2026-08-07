// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type ModalFormRequest struct {
	FormID     uint32
	FormUIJSON string
}

func (p *ModalFormRequest) Encode(w Encoder) error {
	if err := w.Write("ModalFormRequestPacket.Form ID", Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, p.FormID); err != nil {
		return err
	}
	if err := w.Write("ModalFormRequestPacket.Form UI JSON", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, p.FormUIJSON); err != nil {
		return err
	}
	return nil
}

func DecodeModalFormRequest(r Decoder) (ModalFormRequest, error) {
	var p ModalFormRequest
	{
		raw, err := r.Read("ModalFormRequestPacket.Form ID", Shape{Kind: "primitive", PrimitiveCode: "var_u32"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(uint32)
		if !ok {
			return p, fmt.Errorf("field ModalFormRequestPacket.Form ID has unexpected decoded type %T", raw)
		}
		p.FormID = value
	}
	{
		raw, err := r.Read("ModalFormRequestPacket.Form UI JSON", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(string)
		if !ok {
			return p, fmt.Errorf("field ModalFormRequestPacket.Form UI JSON has unexpected decoded type %T", raw)
		}
		p.FormUIJSON = value
	}
	return p, nil
}
