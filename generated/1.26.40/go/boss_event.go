// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type BossEvent struct {
	TargetActorID ActorUniqueID
	PlayerID      ActorUniqueID
	EventType     BossEventUpdateType
	Name          string
	FilteredName  string
	HealthPercent float32
	Color         BossBarColor
	Overlay       BossBarOverlay
}

func (p *BossEvent) Encode(w Encoder) error {
	if err := w.Write("BossEventPacket.Target Actor ID", Shape{Kind: "struct", Semantic: "ActorUniqueID", TypeID: "ActorUniqueID", Fields: []ShapeField{{Ordinal: 0, Name: "Actor Unique ID", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i64"}}}}, p.TargetActorID); err != nil {
		return err
	}
	if err := w.Write("BossEventPacket.Player ID", Shape{Kind: "struct", Semantic: "ActorUniqueID", TypeID: "ActorUniqueID", Fields: []ShapeField{{Ordinal: 0, Name: "Actor Unique ID", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i64"}}}}, p.PlayerID); err != nil {
		return err
	}
	if err := w.Write("BossEventPacket.Event Type", Shape{Kind: "enum", Semantic: "BossEventUpdateType", TypeID: "enums/BossEventUpdateType", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "Add", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "PlayerAdded", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "Remove", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "PlayerRemoved", Shape: Shape{Kind: "void"}}, {Value: 4, Name: "Update_Percent", Shape: Shape{Kind: "void"}}, {Value: 5, Name: "Update_Name", Shape: Shape{Kind: "void"}}, {Value: 6, Name: "Update_Properties", Shape: Shape{Kind: "void"}}, {Value: 7, Name: "Update_Style", Shape: Shape{Kind: "void"}}, {Value: 8, Name: "Query", Shape: Shape{Kind: "void"}}}}, p.EventType); err != nil {
		return err
	}
	if err := w.Write("BossEventPacket.Name", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, p.Name); err != nil {
		return err
	}
	if err := w.Write("BossEventPacket.FilteredName", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, p.FilteredName); err != nil {
		return err
	}
	if err := w.Write("BossEventPacket.Health Percent", Shape{Kind: "primitive", PrimitiveCode: "f32le"}, p.HealthPercent); err != nil {
		return err
	}
	if err := w.Write("BossEventPacket.Color", Shape{Kind: "enum", Semantic: "BossBarColor", TypeID: "enums/BossBarColor", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "PINK", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "BLUE", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "RED", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "GREEN", Shape: Shape{Kind: "void"}}, {Value: 4, Name: "YELLOW", Shape: Shape{Kind: "void"}}, {Value: 5, Name: "PURPLE", Shape: Shape{Kind: "void"}}, {Value: 6, Name: "REBECCA_PURPLE", Shape: Shape{Kind: "void"}}, {Value: 7, Name: "WHITE", Shape: Shape{Kind: "void"}}}}, p.Color); err != nil {
		return err
	}
	if err := w.Write("BossEventPacket.Overlay", Shape{Kind: "enum", Semantic: "BossBarOverlay", TypeID: "enums/BossBarOverlay", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "PROGRESS", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "NOTCHED_6", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "NOTCHED_10", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "NOTCHED_12", Shape: Shape{Kind: "void"}}, {Value: 4, Name: "NOTCHED_20", Shape: Shape{Kind: "void"}}}}, p.Overlay); err != nil {
		return err
	}
	return nil
}

func DecodeBossEvent(r Decoder) (BossEvent, error) {
	var p BossEvent
	{
		raw, err := r.Read("BossEventPacket.Target Actor ID", Shape{Kind: "struct", Semantic: "ActorUniqueID", TypeID: "ActorUniqueID", Fields: []ShapeField{{Ordinal: 0, Name: "Actor Unique ID", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i64"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(ActorUniqueID)
		if !ok {
			return p, fmt.Errorf("field BossEventPacket.Target Actor ID has unexpected decoded type %T", raw)
		}
		p.TargetActorID = value
	}
	{
		raw, err := r.Read("BossEventPacket.Player ID", Shape{Kind: "struct", Semantic: "ActorUniqueID", TypeID: "ActorUniqueID", Fields: []ShapeField{{Ordinal: 0, Name: "Actor Unique ID", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i64"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(ActorUniqueID)
		if !ok {
			return p, fmt.Errorf("field BossEventPacket.Player ID has unexpected decoded type %T", raw)
		}
		p.PlayerID = value
	}
	{
		raw, err := r.Read("BossEventPacket.Event Type", Shape{Kind: "enum", Semantic: "BossEventUpdateType", TypeID: "enums/BossEventUpdateType", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "Add", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "PlayerAdded", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "Remove", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "PlayerRemoved", Shape: Shape{Kind: "void"}}, {Value: 4, Name: "Update_Percent", Shape: Shape{Kind: "void"}}, {Value: 5, Name: "Update_Name", Shape: Shape{Kind: "void"}}, {Value: 6, Name: "Update_Properties", Shape: Shape{Kind: "void"}}, {Value: 7, Name: "Update_Style", Shape: Shape{Kind: "void"}}, {Value: 8, Name: "Query", Shape: Shape{Kind: "void"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(BossEventUpdateType)
		if !ok {
			return p, fmt.Errorf("field BossEventPacket.Event Type has unexpected decoded type %T", raw)
		}
		p.EventType = value
	}
	{
		raw, err := r.Read("BossEventPacket.Name", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(string)
		if !ok {
			return p, fmt.Errorf("field BossEventPacket.Name has unexpected decoded type %T", raw)
		}
		p.Name = value
	}
	{
		raw, err := r.Read("BossEventPacket.FilteredName", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(string)
		if !ok {
			return p, fmt.Errorf("field BossEventPacket.FilteredName has unexpected decoded type %T", raw)
		}
		p.FilteredName = value
	}
	{
		raw, err := r.Read("BossEventPacket.Health Percent", Shape{Kind: "primitive", PrimitiveCode: "f32le"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(float32)
		if !ok {
			return p, fmt.Errorf("field BossEventPacket.Health Percent has unexpected decoded type %T", raw)
		}
		p.HealthPercent = value
	}
	{
		raw, err := r.Read("BossEventPacket.Color", Shape{Kind: "enum", Semantic: "BossBarColor", TypeID: "enums/BossBarColor", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "PINK", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "BLUE", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "RED", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "GREEN", Shape: Shape{Kind: "void"}}, {Value: 4, Name: "YELLOW", Shape: Shape{Kind: "void"}}, {Value: 5, Name: "PURPLE", Shape: Shape{Kind: "void"}}, {Value: 6, Name: "REBECCA_PURPLE", Shape: Shape{Kind: "void"}}, {Value: 7, Name: "WHITE", Shape: Shape{Kind: "void"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(BossBarColor)
		if !ok {
			return p, fmt.Errorf("field BossEventPacket.Color has unexpected decoded type %T", raw)
		}
		p.Color = value
	}
	{
		raw, err := r.Read("BossEventPacket.Overlay", Shape{Kind: "enum", Semantic: "BossBarOverlay", TypeID: "enums/BossBarOverlay", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "PROGRESS", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "NOTCHED_6", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "NOTCHED_10", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "NOTCHED_12", Shape: Shape{Kind: "void"}}, {Value: 4, Name: "NOTCHED_20", Shape: Shape{Kind: "void"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(BossBarOverlay)
		if !ok {
			return p, fmt.Errorf("field BossEventPacket.Overlay has unexpected decoded type %T", raw)
		}
		p.Overlay = value
	}
	return p, nil
}
