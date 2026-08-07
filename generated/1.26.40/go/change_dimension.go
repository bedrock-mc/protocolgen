// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import (
	"fmt"

	"github.com/go-gl/mathgl/mgl32"
)

type ChangeDimension struct {
	DimensionID     DimensionType
	Position        mgl32.Vec3
	Respawn         bool
	LoadingScreenId *uint32
}

func (p *ChangeDimension) Encode(w Encoder) error {
	if err := w.Write("ChangeDimensionPacket.Dimension ID", Shape{Kind: "struct", Semantic: "DimensionType", TypeID: "DimensionType", Fields: []ShapeField{{Ordinal: 0, Name: "value", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}}}, p.DimensionID); err != nil {
		return err
	}
	if err := w.Write("ChangeDimensionPacket.Position", Shape{Kind: "struct", Semantic: "Vec3", TypeID: "Vec3", Fields: []ShapeField{{Ordinal: 0, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 1, Name: "Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 2, Name: "Z", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}}}, p.Position); err != nil {
		return err
	}
	if err := w.Write("ChangeDimensionPacket.Respawn", Shape{Kind: "primitive", PrimitiveCode: "bool"}, p.Respawn); err != nil {
		return err
	}
	if err := w.Write("ChangeDimensionPacket.Loading Screen Id", Shape{Kind: "optional", Value: &Shape{Kind: "primitive", PrimitiveCode: "u32le"}}, p.LoadingScreenId); err != nil {
		return err
	}
	return nil
}

func DecodeChangeDimension(r Decoder) (ChangeDimension, error) {
	var p ChangeDimension
	{
		raw, err := r.Read("ChangeDimensionPacket.Dimension ID", Shape{Kind: "struct", Semantic: "DimensionType", TypeID: "DimensionType", Fields: []ShapeField{{Ordinal: 0, Name: "value", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(DimensionType)
		if !ok {
			return p, fmt.Errorf("field ChangeDimensionPacket.Dimension ID has unexpected decoded type %T", raw)
		}
		p.DimensionID = value
	}
	{
		raw, err := r.Read("ChangeDimensionPacket.Position", Shape{Kind: "struct", Semantic: "Vec3", TypeID: "Vec3", Fields: []ShapeField{{Ordinal: 0, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 1, Name: "Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 2, Name: "Z", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(mgl32.Vec3)
		if !ok {
			return p, fmt.Errorf("field ChangeDimensionPacket.Position has unexpected decoded type %T", raw)
		}
		p.Position = value
	}
	{
		raw, err := r.Read("ChangeDimensionPacket.Respawn", Shape{Kind: "primitive", PrimitiveCode: "bool"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(bool)
		if !ok {
			return p, fmt.Errorf("field ChangeDimensionPacket.Respawn has unexpected decoded type %T", raw)
		}
		p.Respawn = value
	}
	{
		raw, err := r.Read("ChangeDimensionPacket.Loading Screen Id", Shape{Kind: "optional", Value: &Shape{Kind: "primitive", PrimitiveCode: "u32le"}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(*uint32)
		if !ok {
			return p, fmt.Errorf("field ChangeDimensionPacket.Loading Screen Id has unexpected decoded type %T", raw)
		}
		p.LoadingScreenId = value
	}
	return p, nil
}
