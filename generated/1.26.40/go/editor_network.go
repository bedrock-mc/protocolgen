// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type EditorNetwork struct {
	RouteToManager bool
	RawVariantName string
	RawVariantData string
}

func (p *EditorNetwork) Encode(w Encoder) error {
	if err := w.Write("EditorNetworkPacket.Route To Manager", Shape{Kind: "primitive", PrimitiveCode: "bool"}, p.RouteToManager); err != nil {
		return err
	}
	if err := w.Write("EditorNetworkPacket.Raw Variant Name", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, p.RawVariantName); err != nil {
		return err
	}
	if err := w.Write("EditorNetworkPacket.Raw Variant Data", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, p.RawVariantData); err != nil {
		return err
	}
	return nil
}

func DecodeEditorNetwork(r Decoder) (EditorNetwork, error) {
	var p EditorNetwork
	{
		raw, err := r.Read("EditorNetworkPacket.Route To Manager", Shape{Kind: "primitive", PrimitiveCode: "bool"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(bool)
		if !ok {
			return p, fmt.Errorf("field EditorNetworkPacket.Route To Manager has unexpected decoded type %T", raw)
		}
		p.RouteToManager = value
	}
	{
		raw, err := r.Read("EditorNetworkPacket.Raw Variant Name", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(string)
		if !ok {
			return p, fmt.Errorf("field EditorNetworkPacket.Raw Variant Name has unexpected decoded type %T", raw)
		}
		p.RawVariantName = value
	}
	{
		raw, err := r.Read("EditorNetworkPacket.Raw Variant Data", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(string)
		if !ok {
			return p, fmt.Errorf("field EditorNetworkPacket.Raw Variant Data has unexpected decoded type %T", raw)
		}
		p.RawVariantData = value
	}
	return p, nil
}
