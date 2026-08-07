// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type PlayerUpdateEntityOverrides struct {
	TargetID      ActorUniqueID
	PropertyIndex uint32
	Update        PlayerUpdateEntityOverridesUpdate
}

func (p *PlayerUpdateEntityOverrides) Encode(w Encoder) error {
	if err := w.Write("PlayerUpdateEntityOverridesPacket.Target ID", Shape{Kind: "struct", Semantic: "ActorUniqueID", TypeID: "ActorUniqueID", Fields: []ShapeField{{Ordinal: 0, Name: "Actor Unique ID", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i64"}}}}, p.TargetID); err != nil {
		return err
	}
	if err := w.Write("PlayerUpdateEntityOverridesPacket.Property Index", Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, p.PropertyIndex); err != nil {
		return err
	}
	if err := w.Write("PlayerUpdateEntityOverridesPacket.Update", Shape{Kind: "union", Control: &Shape{Kind: "primitive", PrimitiveCode: "u8"}, Variants: []ShapeVariant{{Value: 0, Name: "PlayerUpdateEntityOverridesPacketPayload::ClearOverride", Shape: Shape{Kind: "struct", Semantic: "PlayerUpdateEntityOverridesPacketPayload::ClearOverride", TypeID: "PlayerUpdateEntityOverridesPacketPayload::ClearOverride", Fields: []ShapeField{{Ordinal: 0, Name: "Type", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}}}, {Value: 1, Name: "PlayerUpdateEntityOverridesPacketPayload::RemoveOverride", Shape: Shape{Kind: "struct", Semantic: "PlayerUpdateEntityOverridesPacketPayload::RemoveOverride", TypeID: "PlayerUpdateEntityOverridesPacketPayload::RemoveOverride", Fields: []ShapeField{{Ordinal: 0, Name: "Type", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}}}, {Value: 2, Name: "PlayerUpdateEntityOverridesPacketPayload::IntOverride", Shape: Shape{Kind: "struct", Semantic: "PlayerUpdateEntityOverridesPacketPayload::IntOverride", TypeID: "PlayerUpdateEntityOverridesPacketPayload::IntOverride", Fields: []ShapeField{{Ordinal: 0, Name: "Type", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 1, Name: "Value", Shape: Shape{Kind: "primitive", PrimitiveCode: "i32le"}}}}}, {Value: 3, Name: "PlayerUpdateEntityOverridesPacketPayload::FloatOverride", Shape: Shape{Kind: "struct", Semantic: "PlayerUpdateEntityOverridesPacketPayload::FloatOverride", TypeID: "PlayerUpdateEntityOverridesPacketPayload::FloatOverride", Fields: []ShapeField{{Ordinal: 0, Name: "Type", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 1, Name: "Value", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}}}}}}, p.Update); err != nil {
		return err
	}
	return nil
}

func DecodePlayerUpdateEntityOverrides(r Decoder) (PlayerUpdateEntityOverrides, error) {
	var p PlayerUpdateEntityOverrides
	{
		raw, err := r.Read("PlayerUpdateEntityOverridesPacket.Target ID", Shape{Kind: "struct", Semantic: "ActorUniqueID", TypeID: "ActorUniqueID", Fields: []ShapeField{{Ordinal: 0, Name: "Actor Unique ID", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i64"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(ActorUniqueID)
		if !ok {
			return p, fmt.Errorf("field PlayerUpdateEntityOverridesPacket.Target ID has unexpected decoded type %T", raw)
		}
		p.TargetID = value
	}
	{
		raw, err := r.Read("PlayerUpdateEntityOverridesPacket.Property Index", Shape{Kind: "primitive", PrimitiveCode: "var_u32"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(uint32)
		if !ok {
			return p, fmt.Errorf("field PlayerUpdateEntityOverridesPacket.Property Index has unexpected decoded type %T", raw)
		}
		p.PropertyIndex = value
	}
	{
		raw, err := r.Read("PlayerUpdateEntityOverridesPacket.Update", Shape{Kind: "union", Control: &Shape{Kind: "primitive", PrimitiveCode: "u8"}, Variants: []ShapeVariant{{Value: 0, Name: "PlayerUpdateEntityOverridesPacketPayload::ClearOverride", Shape: Shape{Kind: "struct", Semantic: "PlayerUpdateEntityOverridesPacketPayload::ClearOverride", TypeID: "PlayerUpdateEntityOverridesPacketPayload::ClearOverride", Fields: []ShapeField{{Ordinal: 0, Name: "Type", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}}}, {Value: 1, Name: "PlayerUpdateEntityOverridesPacketPayload::RemoveOverride", Shape: Shape{Kind: "struct", Semantic: "PlayerUpdateEntityOverridesPacketPayload::RemoveOverride", TypeID: "PlayerUpdateEntityOverridesPacketPayload::RemoveOverride", Fields: []ShapeField{{Ordinal: 0, Name: "Type", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}}}, {Value: 2, Name: "PlayerUpdateEntityOverridesPacketPayload::IntOverride", Shape: Shape{Kind: "struct", Semantic: "PlayerUpdateEntityOverridesPacketPayload::IntOverride", TypeID: "PlayerUpdateEntityOverridesPacketPayload::IntOverride", Fields: []ShapeField{{Ordinal: 0, Name: "Type", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 1, Name: "Value", Shape: Shape{Kind: "primitive", PrimitiveCode: "i32le"}}}}}, {Value: 3, Name: "PlayerUpdateEntityOverridesPacketPayload::FloatOverride", Shape: Shape{Kind: "struct", Semantic: "PlayerUpdateEntityOverridesPacketPayload::FloatOverride", TypeID: "PlayerUpdateEntityOverridesPacketPayload::FloatOverride", Fields: []ShapeField{{Ordinal: 0, Name: "Type", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 1, Name: "Value", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}}}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(PlayerUpdateEntityOverridesUpdate)
		if !ok {
			return p, fmt.Errorf("field PlayerUpdateEntityOverridesPacket.Update has unexpected decoded type %T", raw)
		}
		p.Update = value
	}
	return p, nil
}
