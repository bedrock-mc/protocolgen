// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type ClientboundDataDrivenUIShowScreen struct {
	ScreenId       string
	FormId         uint32
	DataInstanceId *uint32
}

func (p *ClientboundDataDrivenUIShowScreen) Encode(w Encoder) error {
	if err := w.Write("ClientboundDataDrivenUIShowScreenPacket.ScreenId", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, p.ScreenId); err != nil {
		return err
	}
	if err := w.Write("ClientboundDataDrivenUIShowScreenPacket.FormId", Shape{Kind: "primitive", PrimitiveCode: "u32le"}, p.FormId); err != nil {
		return err
	}
	if err := w.Write("ClientboundDataDrivenUIShowScreenPacket.DataInstanceId", Shape{Kind: "optional", Value: &Shape{Kind: "primitive", PrimitiveCode: "u32le"}}, p.DataInstanceId); err != nil {
		return err
	}
	return nil
}

func DecodeClientboundDataDrivenUIShowScreen(r Decoder) (ClientboundDataDrivenUIShowScreen, error) {
	var p ClientboundDataDrivenUIShowScreen
	{
		raw, err := r.Read("ClientboundDataDrivenUIShowScreenPacket.ScreenId", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(string)
		if !ok {
			return p, fmt.Errorf("field ClientboundDataDrivenUIShowScreenPacket.ScreenId has unexpected decoded type %T", raw)
		}
		p.ScreenId = value
	}
	{
		raw, err := r.Read("ClientboundDataDrivenUIShowScreenPacket.FormId", Shape{Kind: "primitive", PrimitiveCode: "u32le"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(uint32)
		if !ok {
			return p, fmt.Errorf("field ClientboundDataDrivenUIShowScreenPacket.FormId has unexpected decoded type %T", raw)
		}
		p.FormId = value
	}
	{
		raw, err := r.Read("ClientboundDataDrivenUIShowScreenPacket.DataInstanceId", Shape{Kind: "optional", Value: &Shape{Kind: "primitive", PrimitiveCode: "u32le"}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(*uint32)
		if !ok {
			return p, fmt.Errorf("field ClientboundDataDrivenUIShowScreenPacket.DataInstanceId has unexpected decoded type %T", raw)
		}
		p.DataInstanceId = value
	}
	return p, nil
}
