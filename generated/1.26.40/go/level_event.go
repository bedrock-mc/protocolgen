// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import (
	"fmt"

	"github.com/go-gl/mathgl/mgl32"
)

type LevelEvent struct {
	EventId  int32
	Position mgl32.Vec3
	Data     int32
}

func (p *LevelEvent) Encode(w Encoder) error {
	if err := w.Write("LevelEventPacket.Event Id", Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}, p.EventId); err != nil {
		return err
	}
	if err := w.Write("LevelEventPacket.Position", Shape{Kind: "struct", Semantic: "Vec3", TypeID: "Vec3", Fields: []ShapeField{{Ordinal: 0, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 1, Name: "Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 2, Name: "Z", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}}}, p.Position); err != nil {
		return err
	}
	if err := w.Write("LevelEventPacket.Data", Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}, p.Data); err != nil {
		return err
	}
	return nil
}

func DecodeLevelEvent(r Decoder) (LevelEvent, error) {
	var p LevelEvent
	{
		raw, err := r.Read("LevelEventPacket.Event Id", Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(int32)
		if !ok {
			return p, fmt.Errorf("field LevelEventPacket.Event Id has unexpected decoded type %T", raw)
		}
		p.EventId = value
	}
	{
		raw, err := r.Read("LevelEventPacket.Position", Shape{Kind: "struct", Semantic: "Vec3", TypeID: "Vec3", Fields: []ShapeField{{Ordinal: 0, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 1, Name: "Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 2, Name: "Z", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(mgl32.Vec3)
		if !ok {
			return p, fmt.Errorf("field LevelEventPacket.Position has unexpected decoded type %T", raw)
		}
		p.Position = value
	}
	{
		raw, err := r.Read("LevelEventPacket.Data", Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(int32)
		if !ok {
			return p, fmt.Errorf("field LevelEventPacket.Data has unexpected decoded type %T", raw)
		}
		p.Data = value
	}
	return p, nil
}
