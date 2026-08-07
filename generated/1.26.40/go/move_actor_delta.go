// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type MoveActorDelta struct {
	MoveData MoveActorDeltaData
}

func (p *MoveActorDelta) Encode(w Encoder) error {
	if err := w.Write("MoveActorDeltaPacket.Move Data", Shape{Kind: "struct", Semantic: "MoveActorDeltaData", TypeID: "MoveActorDeltaData", Fields: []ShapeField{{Ordinal: 0, Name: "Actor Runtime ID", Shape: Shape{Kind: "struct", Semantic: "ActorRuntimeID", TypeID: "ActorRuntimeID", Fields: []ShapeField{{Ordinal: 0, Name: "Actor Runtime ID", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u64"}}}}}, {Ordinal: 1, Name: "New Position X", Shape: Shape{Kind: "optional", Value: &Shape{Kind: "primitive", PrimitiveCode: "f32le"}}}, {Ordinal: 2, Name: "New Position Y", Shape: Shape{Kind: "optional", Value: &Shape{Kind: "primitive", PrimitiveCode: "f32le"}}}, {Ordinal: 3, Name: "New Position Z", Shape: Shape{Kind: "optional", Value: &Shape{Kind: "primitive", PrimitiveCode: "f32le"}}}, {Ordinal: 4, Name: "Rotation X", Shape: Shape{Kind: "optional", Value: &Shape{Kind: "primitive", PrimitiveCode: "i8"}}}, {Ordinal: 5, Name: "Rotation Y", Shape: Shape{Kind: "optional", Value: &Shape{Kind: "primitive", PrimitiveCode: "i8"}}}, {Ordinal: 6, Name: "Rotation Y Head", Shape: Shape{Kind: "optional", Value: &Shape{Kind: "primitive", PrimitiveCode: "i8"}}}, {Ordinal: 7, Name: "Is On Ground", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}, {Ordinal: 8, Name: "Force Move", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}, {Ordinal: 9, Name: "Force Move Local Entity", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}, {Ordinal: 10, Name: "Force Completion", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}}}, p.MoveData); err != nil {
		return err
	}
	return nil
}

func DecodeMoveActorDelta(r Decoder) (MoveActorDelta, error) {
	var p MoveActorDelta
	{
		raw, err := r.Read("MoveActorDeltaPacket.Move Data", Shape{Kind: "struct", Semantic: "MoveActorDeltaData", TypeID: "MoveActorDeltaData", Fields: []ShapeField{{Ordinal: 0, Name: "Actor Runtime ID", Shape: Shape{Kind: "struct", Semantic: "ActorRuntimeID", TypeID: "ActorRuntimeID", Fields: []ShapeField{{Ordinal: 0, Name: "Actor Runtime ID", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u64"}}}}}, {Ordinal: 1, Name: "New Position X", Shape: Shape{Kind: "optional", Value: &Shape{Kind: "primitive", PrimitiveCode: "f32le"}}}, {Ordinal: 2, Name: "New Position Y", Shape: Shape{Kind: "optional", Value: &Shape{Kind: "primitive", PrimitiveCode: "f32le"}}}, {Ordinal: 3, Name: "New Position Z", Shape: Shape{Kind: "optional", Value: &Shape{Kind: "primitive", PrimitiveCode: "f32le"}}}, {Ordinal: 4, Name: "Rotation X", Shape: Shape{Kind: "optional", Value: &Shape{Kind: "primitive", PrimitiveCode: "i8"}}}, {Ordinal: 5, Name: "Rotation Y", Shape: Shape{Kind: "optional", Value: &Shape{Kind: "primitive", PrimitiveCode: "i8"}}}, {Ordinal: 6, Name: "Rotation Y Head", Shape: Shape{Kind: "optional", Value: &Shape{Kind: "primitive", PrimitiveCode: "i8"}}}, {Ordinal: 7, Name: "Is On Ground", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}, {Ordinal: 8, Name: "Force Move", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}, {Ordinal: 9, Name: "Force Move Local Entity", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}, {Ordinal: 10, Name: "Force Completion", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(MoveActorDeltaData)
		if !ok {
			return p, fmt.Errorf("field MoveActorDeltaPacket.Move Data has unexpected decoded type %T", raw)
		}
		p.MoveData = value
	}
	return p, nil
}
