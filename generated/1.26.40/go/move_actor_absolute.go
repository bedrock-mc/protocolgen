// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type MoveActorAbsolute struct {
	MoveData MoveActorAbsoluteData
}

func (p *MoveActorAbsolute) Encode(w Encoder) error {
	if err := w.Write("MoveActorAbsolutePacket.Move Data", Shape{Kind: "struct", Semantic: "MoveActorAbsoluteData", TypeID: "MoveActorAbsoluteData", Fields: []ShapeField{{Ordinal: 0, Name: "ActorRuntimeID", Shape: Shape{Kind: "struct", Semantic: "ActorRuntimeID", TypeID: "ActorRuntimeID", Fields: []ShapeField{{Ordinal: 0, Name: "Actor Runtime ID", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u64"}}}}}, {Ordinal: 1, Name: "Header", Shape: Shape{Kind: "primitive", PrimitiveCode: "u8"}}, {Ordinal: 2, Name: "Position", Shape: Shape{Kind: "struct", Semantic: "Vec3", TypeID: "Vec3", Fields: []ShapeField{{Ordinal: 0, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 1, Name: "Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 2, Name: "Z", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}}}}, {Ordinal: 3, Name: "Rotation X", Shape: Shape{Kind: "primitive", PrimitiveCode: "u8"}}, {Ordinal: 4, Name: "Rotation Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "u8"}}, {Ordinal: 5, Name: "Rotation Y Head", Shape: Shape{Kind: "primitive", PrimitiveCode: "u8"}}}}, p.MoveData); err != nil {
		return err
	}
	return nil
}

func DecodeMoveActorAbsolute(r Decoder) (MoveActorAbsolute, error) {
	var p MoveActorAbsolute
	{
		raw, err := r.Read("MoveActorAbsolutePacket.Move Data", Shape{Kind: "struct", Semantic: "MoveActorAbsoluteData", TypeID: "MoveActorAbsoluteData", Fields: []ShapeField{{Ordinal: 0, Name: "ActorRuntimeID", Shape: Shape{Kind: "struct", Semantic: "ActorRuntimeID", TypeID: "ActorRuntimeID", Fields: []ShapeField{{Ordinal: 0, Name: "Actor Runtime ID", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u64"}}}}}, {Ordinal: 1, Name: "Header", Shape: Shape{Kind: "primitive", PrimitiveCode: "u8"}}, {Ordinal: 2, Name: "Position", Shape: Shape{Kind: "struct", Semantic: "Vec3", TypeID: "Vec3", Fields: []ShapeField{{Ordinal: 0, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 1, Name: "Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 2, Name: "Z", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}}}}, {Ordinal: 3, Name: "Rotation X", Shape: Shape{Kind: "primitive", PrimitiveCode: "u8"}}, {Ordinal: 4, Name: "Rotation Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "u8"}}, {Ordinal: 5, Name: "Rotation Y Head", Shape: Shape{Kind: "primitive", PrimitiveCode: "u8"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(MoveActorAbsoluteData)
		if !ok {
			return p, fmt.Errorf("field MoveActorAbsolutePacket.Move Data has unexpected decoded type %T", raw)
		}
		p.MoveData = value
	}
	return p, nil
}
