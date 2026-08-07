// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type PlayerLocation struct {
	TargetActorID ActorUniqueID
	Location      PlayerLocationLocation
}

func (p *PlayerLocation) Encode(w Encoder) error {
	if err := w.Write("PlayerLocationPacket.Target Actor ID", Shape{Kind: "struct", Semantic: "ActorUniqueID", TypeID: "ActorUniqueID", Fields: []ShapeField{{Ordinal: 0, Name: "Actor Unique ID", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i64"}}}}, p.TargetActorID); err != nil {
		return err
	}
	if err := w.Write("PlayerLocationPacket.Location", Shape{Kind: "union", Control: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Variants: []ShapeVariant{{Value: 0, Name: "PLAYER_LOCATION_COORDINATES", Shape: Shape{Kind: "struct", Semantic: "PlayerLocationPacketPayload::CoordinatesLocation", TypeID: "PlayerLocationPacketPayload::CoordinatesLocation", Fields: []ShapeField{{Ordinal: 0, Name: "Packet Type", Shape: Shape{Kind: "enum", Semantic: "PlayerLocationPacketPayload::Type", TypeID: "enums/PlayerLocationPacketPayload::Type", PrimitiveCode: "zigzag_i32", Variants: []ShapeVariant{{Value: 0, Name: "PLAYER_LOCATION_COORDINATES", Shape: Shape{Kind: "void"}}}}}, {Ordinal: 1, Name: "Position", Shape: Shape{Kind: "struct", Semantic: "Vec3", TypeID: "Vec3", Fields: []ShapeField{{Ordinal: 0, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 1, Name: "Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 2, Name: "Z", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}}}}}}}, {Value: 1, Name: "PLAYER_LOCATION_HIDE", Shape: Shape{Kind: "struct", Semantic: "PlayerLocationPacketPayload::HiddenLocation", TypeID: "PlayerLocationPacketPayload::HiddenLocation", Fields: []ShapeField{{Ordinal: 0, Name: "Packet Type", Shape: Shape{Kind: "enum", Semantic: "PlayerLocationPacketPayload::Type", TypeID: "enums/PlayerLocationPacketPayload::Type", PrimitiveCode: "zigzag_i32", Variants: []ShapeVariant{{Value: 1, Name: "PLAYER_LOCATION_HIDE", Shape: Shape{Kind: "void"}}}}}}}}}}, p.Location); err != nil {
		return err
	}
	return nil
}

func DecodePlayerLocation(r Decoder) (PlayerLocation, error) {
	var p PlayerLocation
	{
		raw, err := r.Read("PlayerLocationPacket.Target Actor ID", Shape{Kind: "struct", Semantic: "ActorUniqueID", TypeID: "ActorUniqueID", Fields: []ShapeField{{Ordinal: 0, Name: "Actor Unique ID", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i64"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(ActorUniqueID)
		if !ok {
			return p, fmt.Errorf("field PlayerLocationPacket.Target Actor ID has unexpected decoded type %T", raw)
		}
		p.TargetActorID = value
	}
	{
		raw, err := r.Read("PlayerLocationPacket.Location", Shape{Kind: "union", Control: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Variants: []ShapeVariant{{Value: 0, Name: "PLAYER_LOCATION_COORDINATES", Shape: Shape{Kind: "struct", Semantic: "PlayerLocationPacketPayload::CoordinatesLocation", TypeID: "PlayerLocationPacketPayload::CoordinatesLocation", Fields: []ShapeField{{Ordinal: 0, Name: "Packet Type", Shape: Shape{Kind: "enum", Semantic: "PlayerLocationPacketPayload::Type", TypeID: "enums/PlayerLocationPacketPayload::Type", PrimitiveCode: "zigzag_i32", Variants: []ShapeVariant{{Value: 0, Name: "PLAYER_LOCATION_COORDINATES", Shape: Shape{Kind: "void"}}}}}, {Ordinal: 1, Name: "Position", Shape: Shape{Kind: "struct", Semantic: "Vec3", TypeID: "Vec3", Fields: []ShapeField{{Ordinal: 0, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 1, Name: "Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 2, Name: "Z", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}}}}}}}, {Value: 1, Name: "PLAYER_LOCATION_HIDE", Shape: Shape{Kind: "struct", Semantic: "PlayerLocationPacketPayload::HiddenLocation", TypeID: "PlayerLocationPacketPayload::HiddenLocation", Fields: []ShapeField{{Ordinal: 0, Name: "Packet Type", Shape: Shape{Kind: "enum", Semantic: "PlayerLocationPacketPayload::Type", TypeID: "enums/PlayerLocationPacketPayload::Type", PrimitiveCode: "zigzag_i32", Variants: []ShapeVariant{{Value: 1, Name: "PLAYER_LOCATION_HIDE", Shape: Shape{Kind: "void"}}}}}}}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(PlayerLocationLocation)
		if !ok {
			return p, fmt.Errorf("field PlayerLocationPacket.Location has unexpected decoded type %T", raw)
		}
		p.Location = value
	}
	return p, nil
}
