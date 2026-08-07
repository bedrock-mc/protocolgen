// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type ServerboundDataDrivenScreenClosed struct {
	FormId      uint32
	CloseReason string
}

func (p *ServerboundDataDrivenScreenClosed) Encode(w Encoder) error {
	if err := w.Write("ServerboundDataDrivenScreenClosedPacket.FormId", Shape{Kind: "primitive", PrimitiveCode: "u32le"}, p.FormId); err != nil {
		return err
	}
	if err := w.Write("ServerboundDataDrivenScreenClosedPacket.CloseReason", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, p.CloseReason); err != nil {
		return err
	}
	return nil
}

func DecodeServerboundDataDrivenScreenClosed(r Decoder) (ServerboundDataDrivenScreenClosed, error) {
	var p ServerboundDataDrivenScreenClosed
	{
		raw, err := r.Read("ServerboundDataDrivenScreenClosedPacket.FormId", Shape{Kind: "primitive", PrimitiveCode: "u32le"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(uint32)
		if !ok {
			return p, fmt.Errorf("field ServerboundDataDrivenScreenClosedPacket.FormId has unexpected decoded type %T", raw)
		}
		p.FormId = value
	}
	{
		raw, err := r.Read("ServerboundDataDrivenScreenClosedPacket.CloseReason", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(string)
		if !ok {
			return p, fmt.Errorf("field ServerboundDataDrivenScreenClosedPacket.CloseReason has unexpected decoded type %T", raw)
		}
		p.CloseReason = value
	}
	return p, nil
}
