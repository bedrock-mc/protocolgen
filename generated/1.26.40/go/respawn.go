// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type Respawn struct {
	Position        Vec3
	State           PlayerRespawnState
	PlayerRuntimeId ActorRuntimeID
}

func (p *Respawn) Encode(w Encoder) error {
	if err := w.Write("RespawnPacket.Position", Shape{Kind: "struct", Semantic: "Vec3", TypeID: "Vec3", Fields: []ShapeField{{Ordinal: 0, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 1, Name: "Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 2, Name: "Z", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}}}, p.Position); err != nil {
		return err
	}
	if err := w.Write("RespawnPacket.State", Shape{Kind: "enum", Semantic: "PlayerRespawnState", TypeID: "enums/PlayerRespawnState", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "SearchingForSpawn", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "ReadyToSpawn", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "ClientReadyToSpawn", Shape: Shape{Kind: "void"}}}}, p.State); err != nil {
		return err
	}
	if err := w.Write("RespawnPacket.Player Runtime Id", Shape{Kind: "struct", Semantic: "ActorRuntimeID", TypeID: "ActorRuntimeID", Fields: []ShapeField{{Ordinal: 0, Name: "Actor Runtime ID", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u64"}}}}, p.PlayerRuntimeId); err != nil {
		return err
	}
	return nil
}

func DecodeRespawn(r Decoder) (Respawn, error) {
	var p Respawn
	{
		raw, err := r.Read("RespawnPacket.Position", Shape{Kind: "struct", Semantic: "Vec3", TypeID: "Vec3", Fields: []ShapeField{{Ordinal: 0, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 1, Name: "Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 2, Name: "Z", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(Vec3)
		if !ok {
			return p, fmt.Errorf("field RespawnPacket.Position has unexpected decoded type %T", raw)
		}
		p.Position = value
	}
	{
		raw, err := r.Read("RespawnPacket.State", Shape{Kind: "enum", Semantic: "PlayerRespawnState", TypeID: "enums/PlayerRespawnState", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "SearchingForSpawn", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "ReadyToSpawn", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "ClientReadyToSpawn", Shape: Shape{Kind: "void"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(PlayerRespawnState)
		if !ok {
			return p, fmt.Errorf("field RespawnPacket.State has unexpected decoded type %T", raw)
		}
		p.State = value
	}
	{
		raw, err := r.Read("RespawnPacket.Player Runtime Id", Shape{Kind: "struct", Semantic: "ActorRuntimeID", TypeID: "ActorRuntimeID", Fields: []ShapeField{{Ordinal: 0, Name: "Actor Runtime ID", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u64"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(ActorRuntimeID)
		if !ok {
			return p, fmt.Errorf("field RespawnPacket.Player Runtime Id has unexpected decoded type %T", raw)
		}
		p.PlayerRuntimeId = value
	}
	return p, nil
}
