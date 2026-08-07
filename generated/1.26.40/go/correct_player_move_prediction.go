// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type CorrectPlayerMovePrediction struct {
	PredictionType         RewindType
	Pos                    Vec3
	PosDelta               Vec3
	Rotation               Vec2
	VehicleAngularVelocity *float32
	OnGround               bool
	Tick                   PlayerInputTick
}

func (p *CorrectPlayerMovePrediction) Encode(w Encoder) error {
	if err := w.Write("CorrectPlayerMovePredictionPacket.PredictionType", Shape{Kind: "enum", Semantic: "RewindType", TypeID: "enums/RewindType", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "Player", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "Vehicle", Shape: Shape{Kind: "void"}}}}, p.PredictionType); err != nil {
		return err
	}
	if err := w.Write("CorrectPlayerMovePredictionPacket.Pos", Shape{Kind: "struct", Semantic: "Vec3", TypeID: "Vec3", Fields: []ShapeField{{Ordinal: 0, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 1, Name: "Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 2, Name: "Z", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}}}, p.Pos); err != nil {
		return err
	}
	if err := w.Write("CorrectPlayerMovePredictionPacket.Pos Delta", Shape{Kind: "struct", Semantic: "Vec3", TypeID: "Vec3", Fields: []ShapeField{{Ordinal: 0, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 1, Name: "Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 2, Name: "Z", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}}}, p.PosDelta); err != nil {
		return err
	}
	if err := w.Write("CorrectPlayerMovePredictionPacket.Rotation", Shape{Kind: "struct", Semantic: "Vec2", TypeID: "Vec2", Fields: []ShapeField{{Ordinal: 0, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 1, Name: "Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}}}, p.Rotation); err != nil {
		return err
	}
	if err := w.Write("CorrectPlayerMovePredictionPacket.VehicleAngularVelocity", Shape{Kind: "optional", Value: &Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, p.VehicleAngularVelocity); err != nil {
		return err
	}
	if err := w.Write("CorrectPlayerMovePredictionPacket.On Ground", Shape{Kind: "primitive", PrimitiveCode: "bool"}, p.OnGround); err != nil {
		return err
	}
	if err := w.Write("CorrectPlayerMovePredictionPacket.Tick", Shape{Kind: "struct", Semantic: "PlayerInputTick", TypeID: "PlayerInputTick", Fields: []ShapeField{{Ordinal: 0, Name: "Input tick", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u64"}}}}, p.Tick); err != nil {
		return err
	}
	return nil
}

func DecodeCorrectPlayerMovePrediction(r Decoder) (CorrectPlayerMovePrediction, error) {
	var p CorrectPlayerMovePrediction
	{
		raw, err := r.Read("CorrectPlayerMovePredictionPacket.PredictionType", Shape{Kind: "enum", Semantic: "RewindType", TypeID: "enums/RewindType", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "Player", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "Vehicle", Shape: Shape{Kind: "void"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(RewindType)
		if !ok {
			return p, fmt.Errorf("field CorrectPlayerMovePredictionPacket.PredictionType has unexpected decoded type %T", raw)
		}
		p.PredictionType = value
	}
	{
		raw, err := r.Read("CorrectPlayerMovePredictionPacket.Pos", Shape{Kind: "struct", Semantic: "Vec3", TypeID: "Vec3", Fields: []ShapeField{{Ordinal: 0, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 1, Name: "Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 2, Name: "Z", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(Vec3)
		if !ok {
			return p, fmt.Errorf("field CorrectPlayerMovePredictionPacket.Pos has unexpected decoded type %T", raw)
		}
		p.Pos = value
	}
	{
		raw, err := r.Read("CorrectPlayerMovePredictionPacket.Pos Delta", Shape{Kind: "struct", Semantic: "Vec3", TypeID: "Vec3", Fields: []ShapeField{{Ordinal: 0, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 1, Name: "Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 2, Name: "Z", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(Vec3)
		if !ok {
			return p, fmt.Errorf("field CorrectPlayerMovePredictionPacket.Pos Delta has unexpected decoded type %T", raw)
		}
		p.PosDelta = value
	}
	{
		raw, err := r.Read("CorrectPlayerMovePredictionPacket.Rotation", Shape{Kind: "struct", Semantic: "Vec2", TypeID: "Vec2", Fields: []ShapeField{{Ordinal: 0, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 1, Name: "Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(Vec2)
		if !ok {
			return p, fmt.Errorf("field CorrectPlayerMovePredictionPacket.Rotation has unexpected decoded type %T", raw)
		}
		p.Rotation = value
	}
	{
		raw, err := r.Read("CorrectPlayerMovePredictionPacket.VehicleAngularVelocity", Shape{Kind: "optional", Value: &Shape{Kind: "primitive", PrimitiveCode: "f32le"}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(*float32)
		if !ok {
			return p, fmt.Errorf("field CorrectPlayerMovePredictionPacket.VehicleAngularVelocity has unexpected decoded type %T", raw)
		}
		p.VehicleAngularVelocity = value
	}
	{
		raw, err := r.Read("CorrectPlayerMovePredictionPacket.On Ground", Shape{Kind: "primitive", PrimitiveCode: "bool"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(bool)
		if !ok {
			return p, fmt.Errorf("field CorrectPlayerMovePredictionPacket.On Ground has unexpected decoded type %T", raw)
		}
		p.OnGround = value
	}
	{
		raw, err := r.Read("CorrectPlayerMovePredictionPacket.Tick", Shape{Kind: "struct", Semantic: "PlayerInputTick", TypeID: "PlayerInputTick", Fields: []ShapeField{{Ordinal: 0, Name: "Input tick", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u64"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(PlayerInputTick)
		if !ok {
			return p, fmt.Errorf("field CorrectPlayerMovePredictionPacket.Tick has unexpected decoded type %T", raw)
		}
		p.Tick = value
	}
	return p, nil
}
