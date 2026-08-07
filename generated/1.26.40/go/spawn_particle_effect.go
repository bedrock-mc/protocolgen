// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import (
	"fmt"

	"github.com/go-gl/mathgl/mgl32"
)

type SpawnParticleEffect struct {
	DimensionId     uint8
	ActorId         ActorUniqueID
	Position        mgl32.Vec3
	EffectName      string
	MolangVariables *string
}

func (p *SpawnParticleEffect) Encode(w Encoder) error {
	if err := w.Write("SpawnParticleEffectPacket.Dimension Id", Shape{Kind: "primitive", PrimitiveCode: "u8"}, p.DimensionId); err != nil {
		return err
	}
	if err := w.Write("SpawnParticleEffectPacket.Actor Id", Shape{Kind: "struct", Semantic: "ActorUniqueID", TypeID: "ActorUniqueID", Fields: []ShapeField{{Ordinal: 0, Name: "Actor Unique ID", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i64"}}}}, p.ActorId); err != nil {
		return err
	}
	if err := w.Write("SpawnParticleEffectPacket.Position", Shape{Kind: "struct", Semantic: "Vec3", TypeID: "Vec3", Fields: []ShapeField{{Ordinal: 0, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 1, Name: "Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 2, Name: "Z", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}}}, p.Position); err != nil {
		return err
	}
	if err := w.Write("SpawnParticleEffectPacket.Effect Name", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, p.EffectName); err != nil {
		return err
	}
	if err := w.Write("SpawnParticleEffectPacket.Molang Variables", Shape{Kind: "optional", Value: &Shape{Kind: "string", Semantic: "Json::Value", TypeID: "MolangVariableMap.json#", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, p.MolangVariables); err != nil {
		return err
	}
	return nil
}

func DecodeSpawnParticleEffect(r Decoder) (SpawnParticleEffect, error) {
	var p SpawnParticleEffect
	{
		raw, err := r.Read("SpawnParticleEffectPacket.Dimension Id", Shape{Kind: "primitive", PrimitiveCode: "u8"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(uint8)
		if !ok {
			return p, fmt.Errorf("field SpawnParticleEffectPacket.Dimension Id has unexpected decoded type %T", raw)
		}
		p.DimensionId = value
	}
	{
		raw, err := r.Read("SpawnParticleEffectPacket.Actor Id", Shape{Kind: "struct", Semantic: "ActorUniqueID", TypeID: "ActorUniqueID", Fields: []ShapeField{{Ordinal: 0, Name: "Actor Unique ID", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i64"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(ActorUniqueID)
		if !ok {
			return p, fmt.Errorf("field SpawnParticleEffectPacket.Actor Id has unexpected decoded type %T", raw)
		}
		p.ActorId = value
	}
	{
		raw, err := r.Read("SpawnParticleEffectPacket.Position", Shape{Kind: "struct", Semantic: "Vec3", TypeID: "Vec3", Fields: []ShapeField{{Ordinal: 0, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 1, Name: "Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 2, Name: "Z", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(mgl32.Vec3)
		if !ok {
			return p, fmt.Errorf("field SpawnParticleEffectPacket.Position has unexpected decoded type %T", raw)
		}
		p.Position = value
	}
	{
		raw, err := r.Read("SpawnParticleEffectPacket.Effect Name", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(string)
		if !ok {
			return p, fmt.Errorf("field SpawnParticleEffectPacket.Effect Name has unexpected decoded type %T", raw)
		}
		p.EffectName = value
	}
	{
		raw, err := r.Read("SpawnParticleEffectPacket.Molang Variables", Shape{Kind: "optional", Value: &Shape{Kind: "string", Semantic: "Json::Value", TypeID: "MolangVariableMap.json#", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(*string)
		if !ok {
			return p, fmt.Errorf("field SpawnParticleEffectPacket.Molang Variables has unexpected decoded type %T", raw)
		}
		p.MolangVariables = value
	}
	return p, nil
}
