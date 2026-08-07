// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type ClientboundDataDrivenUICloseScreen struct {
	FormId *uint32
}

func (p *ClientboundDataDrivenUICloseScreen) Encode(w Encoder) error {
	if err := w.Write("ClientboundDataDrivenUICloseScreenPacket.FormId", Shape{Kind: "optional", Value: &Shape{Kind: "primitive", PrimitiveCode: "u32le"}}, p.FormId); err != nil {
		return err
	}
	return nil
}

func DecodeClientboundDataDrivenUICloseScreen(r Decoder) (ClientboundDataDrivenUICloseScreen, error) {
	var p ClientboundDataDrivenUICloseScreen
	{
		raw, err := r.Read("ClientboundDataDrivenUICloseScreenPacket.FormId", Shape{Kind: "optional", Value: &Shape{Kind: "primitive", PrimitiveCode: "u32le"}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(*uint32)
		if !ok {
			return p, fmt.Errorf("field ClientboundDataDrivenUICloseScreenPacket.FormId has unexpected decoded type %T", raw)
		}
		p.FormId = value
	}
	return p, nil
}
