// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type OpenSign struct {
	Pos         BlockPos
	IsFrontSide bool
}

func (p *OpenSign) Encode(w Encoder) error {
	if err := w.Write("OpenSignPacket.Pos", Shape{Kind: "struct", Semantic: "BlockPos", TypeID: "BlockPos", Fields: []ShapeField{{Ordinal: 0, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 1, Name: "Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 2, Name: "Z", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}}}, p.Pos); err != nil {
		return err
	}
	if err := w.Write("OpenSignPacket.Is Front Side", Shape{Kind: "primitive", PrimitiveCode: "bool"}, p.IsFrontSide); err != nil {
		return err
	}
	return nil
}

func DecodeOpenSign(r Decoder) (OpenSign, error) {
	var p OpenSign
	{
		raw, err := r.Read("OpenSignPacket.Pos", Shape{Kind: "struct", Semantic: "BlockPos", TypeID: "BlockPos", Fields: []ShapeField{{Ordinal: 0, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 1, Name: "Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 2, Name: "Z", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(BlockPos)
		if !ok {
			return p, fmt.Errorf("field OpenSignPacket.Pos has unexpected decoded type %T", raw)
		}
		p.Pos = value
	}
	{
		raw, err := r.Read("OpenSignPacket.Is Front Side", Shape{Kind: "primitive", PrimitiveCode: "bool"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(bool)
		if !ok {
			return p, fmt.Errorf("field OpenSignPacket.Is Front Side has unexpected decoded type %T", raw)
		}
		p.IsFrontSide = value
	}
	return p, nil
}
