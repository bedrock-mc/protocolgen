// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type SetScoreboardIdentity struct {
	ScoreboardIdentityPacketType ScoreboardIdentityPacketType
	ScoreboardIdentityInfo       []ScoreboardIdentityPacketInfo
}

func (p *SetScoreboardIdentity) Encode(w Encoder) error {
	if err := w.Write("SetScoreboardIdentityPacket.Scoreboard Identity Packet Type", Shape{Kind: "enum", Semantic: "ScoreboardIdentityPacketType", TypeID: "enums/ScoreboardIdentityPacketType", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "Update", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "Remove", Shape: Shape{Kind: "void"}}}}, p.ScoreboardIdentityPacketType); err != nil {
		return err
	}
	if err := w.Write("SetScoreboardIdentityPacket.Scoreboard Identity Info", Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "struct", Semantic: "ScoreboardIdentityPacketInfo", TypeID: "ScoreboardIdentityPacketInfo", Fields: []ShapeField{{Ordinal: 0, Name: "Scoreboard Id", Shape: Shape{Kind: "struct", Semantic: "ScoreboardId", TypeID: "ScoreboardId", Fields: []ShapeField{{Ordinal: 0, Name: "Scoreboard Id", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i64"}}}}}, {Ordinal: 1, Name: "Player Unique Id", Shape: Shape{Kind: "optional", Value: &Shape{Kind: "primitive", PrimitiveCode: "zigzag_i64"}}}}}}, p.ScoreboardIdentityInfo); err != nil {
		return err
	}
	return nil
}

func DecodeSetScoreboardIdentity(r Decoder) (SetScoreboardIdentity, error) {
	var p SetScoreboardIdentity
	{
		raw, err := r.Read("SetScoreboardIdentityPacket.Scoreboard Identity Packet Type", Shape{Kind: "enum", Semantic: "ScoreboardIdentityPacketType", TypeID: "enums/ScoreboardIdentityPacketType", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "Update", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "Remove", Shape: Shape{Kind: "void"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(ScoreboardIdentityPacketType)
		if !ok {
			return p, fmt.Errorf("field SetScoreboardIdentityPacket.Scoreboard Identity Packet Type has unexpected decoded type %T", raw)
		}
		p.ScoreboardIdentityPacketType = value
	}
	{
		raw, err := r.Read("SetScoreboardIdentityPacket.Scoreboard Identity Info", Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "struct", Semantic: "ScoreboardIdentityPacketInfo", TypeID: "ScoreboardIdentityPacketInfo", Fields: []ShapeField{{Ordinal: 0, Name: "Scoreboard Id", Shape: Shape{Kind: "struct", Semantic: "ScoreboardId", TypeID: "ScoreboardId", Fields: []ShapeField{{Ordinal: 0, Name: "Scoreboard Id", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i64"}}}}}, {Ordinal: 1, Name: "Player Unique Id", Shape: Shape{Kind: "optional", Value: &Shape{Kind: "primitive", PrimitiveCode: "zigzag_i64"}}}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.([]ScoreboardIdentityPacketInfo)
		if !ok {
			return p, fmt.Errorf("field SetScoreboardIdentityPacket.Scoreboard Identity Info has unexpected decoded type %T", raw)
		}
		p.ScoreboardIdentityInfo = value
	}
	return p, nil
}
