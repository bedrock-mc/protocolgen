// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import (
	"fmt"

	"github.com/go-gl/mathgl/mgl32"
)

type SpawnExperienceOrb struct {
	Position mgl32.Vec3
	XPValue  int32
}

func (p *SpawnExperienceOrb) Encode(w Encoder) error {
	if err := w.Write("SpawnExperienceOrbPacket.Position", Shape{Kind: "struct", Semantic: "Vec3", TypeID: "Vec3", Fields: []ShapeField{{Ordinal: 0, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 1, Name: "Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 2, Name: "Z", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}}}, p.Position); err != nil {
		return err
	}
	if err := w.Write("SpawnExperienceOrbPacket.XP Value", Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}, p.XPValue); err != nil {
		return err
	}
	return nil
}

func DecodeSpawnExperienceOrb(r Decoder) (SpawnExperienceOrb, error) {
	var p SpawnExperienceOrb
	{
		raw, err := r.Read("SpawnExperienceOrbPacket.Position", Shape{Kind: "struct", Semantic: "Vec3", TypeID: "Vec3", Fields: []ShapeField{{Ordinal: 0, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 1, Name: "Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 2, Name: "Z", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(mgl32.Vec3)
		if !ok {
			return p, fmt.Errorf("field SpawnExperienceOrbPacket.Position has unexpected decoded type %T", raw)
		}
		p.Position = value
	}
	{
		raw, err := r.Read("SpawnExperienceOrbPacket.XP Value", Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(int32)
		if !ok {
			return p, fmt.Errorf("field SpawnExperienceOrbPacket.XP Value has unexpected decoded type %T", raw)
		}
		p.XPValue = value
	}
	return p, nil
}
