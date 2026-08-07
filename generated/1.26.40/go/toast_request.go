// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type ToastRequest struct {
	Title   string
	Content string
}

func (p *ToastRequest) Encode(w Encoder) error {
	if err := w.Write("ToastRequestPacket.Title", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, p.Title); err != nil {
		return err
	}
	if err := w.Write("ToastRequestPacket.Content", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, p.Content); err != nil {
		return err
	}
	return nil
}

func DecodeToastRequest(r Decoder) (ToastRequest, error) {
	var p ToastRequest
	{
		raw, err := r.Read("ToastRequestPacket.Title", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(string)
		if !ok {
			return p, fmt.Errorf("field ToastRequestPacket.Title has unexpected decoded type %T", raw)
		}
		p.Title = value
	}
	{
		raw, err := r.Read("ToastRequestPacket.Content", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(string)
		if !ok {
			return p, fmt.Errorf("field ToastRequestPacket.Content has unexpected decoded type %T", raw)
		}
		p.Content = value
	}
	return p, nil
}
