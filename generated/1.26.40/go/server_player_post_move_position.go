// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type ServerPlayerPostMovePosition struct {
	Pos Vec3
}

func (p *ServerPlayerPostMovePosition) Encode(w Encoder) error {
	if err := w.Write("ServerPlayerPostMovePositionPacket.Pos", Shape{Kind: "struct", Semantic: "Vec3", TypeID: "Vec3", Fields: []ShapeField{{Ordinal: 0, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 1, Name: "Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 2, Name: "Z", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}}}, p.Pos); err != nil {
		return err
	}
	return nil
}

func DecodeServerPlayerPostMovePosition(r Decoder) (ServerPlayerPostMovePosition, error) {
	var p ServerPlayerPostMovePosition
	{
		raw, err := r.Read("ServerPlayerPostMovePositionPacket.Pos", Shape{Kind: "struct", Semantic: "Vec3", TypeID: "Vec3", Fields: []ShapeField{{Ordinal: 0, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 1, Name: "Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 2, Name: "Z", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(Vec3)
		if !ok {
			return p, fmt.Errorf("field ServerPlayerPostMovePositionPacket.Pos has unexpected decoded type %T", raw)
		}
		p.Pos = value
	}
	return p, nil
}
