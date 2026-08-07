// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type LecternUpdate struct {
	NewPageToShow             uint8
	TotalPages                uint8
	PositionOfLecternToUpdate BlockPos
}

func (p *LecternUpdate) Encode(w Encoder) error {
	if err := w.Write("LecternUpdatePacket.New page to show", Shape{Kind: "primitive", PrimitiveCode: "u8"}, p.NewPageToShow); err != nil {
		return err
	}
	if err := w.Write("LecternUpdatePacket.Total Pages", Shape{Kind: "primitive", PrimitiveCode: "u8"}, p.TotalPages); err != nil {
		return err
	}
	if err := w.Write("LecternUpdatePacket.Position of Lectern to update", Shape{Kind: "struct", Semantic: "BlockPos", TypeID: "BlockPos", Fields: []ShapeField{{Ordinal: 0, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 1, Name: "Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 2, Name: "Z", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}}}, p.PositionOfLecternToUpdate); err != nil {
		return err
	}
	return nil
}

func DecodeLecternUpdate(r Decoder) (LecternUpdate, error) {
	var p LecternUpdate
	{
		raw, err := r.Read("LecternUpdatePacket.New page to show", Shape{Kind: "primitive", PrimitiveCode: "u8"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(uint8)
		if !ok {
			return p, fmt.Errorf("field LecternUpdatePacket.New page to show has unexpected decoded type %T", raw)
		}
		p.NewPageToShow = value
	}
	{
		raw, err := r.Read("LecternUpdatePacket.Total Pages", Shape{Kind: "primitive", PrimitiveCode: "u8"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(uint8)
		if !ok {
			return p, fmt.Errorf("field LecternUpdatePacket.Total Pages has unexpected decoded type %T", raw)
		}
		p.TotalPages = value
	}
	{
		raw, err := r.Read("LecternUpdatePacket.Position of Lectern to update", Shape{Kind: "struct", Semantic: "BlockPos", TypeID: "BlockPos", Fields: []ShapeField{{Ordinal: 0, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 1, Name: "Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 2, Name: "Z", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(BlockPos)
		if !ok {
			return p, fmt.Errorf("field LecternUpdatePacket.Position of Lectern to update has unexpected decoded type %T", raw)
		}
		p.PositionOfLecternToUpdate = value
	}
	return p, nil
}
