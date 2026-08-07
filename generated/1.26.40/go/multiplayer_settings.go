// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type MultiplayerSettings struct {
	PacketType MultiplayerSettingsPacketType
}

func (p *MultiplayerSettings) Encode(w Encoder) error {
	if err := w.Write("MultiplayerSettingsPacket.PacketType", Shape{Kind: "enum", Semantic: "MultiplayerSettingsPacketType", TypeID: "enums/MultiplayerSettingsPacketType", PrimitiveCode: "zigzag_i32", Variants: []ShapeVariant{{Value: 0, Name: "EnableMultiplayer", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "DisableMultiplayer", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "RefreshJoincode", Shape: Shape{Kind: "void"}}}}, p.PacketType); err != nil {
		return err
	}
	return nil
}

func DecodeMultiplayerSettings(r Decoder) (MultiplayerSettings, error) {
	var p MultiplayerSettings
	{
		raw, err := r.Read("MultiplayerSettingsPacket.PacketType", Shape{Kind: "enum", Semantic: "MultiplayerSettingsPacketType", TypeID: "enums/MultiplayerSettingsPacketType", PrimitiveCode: "zigzag_i32", Variants: []ShapeVariant{{Value: 0, Name: "EnableMultiplayer", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "DisableMultiplayer", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "RefreshJoincode", Shape: Shape{Kind: "void"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(MultiplayerSettingsPacketType)
		if !ok {
			return p, fmt.Errorf("field MultiplayerSettingsPacket.PacketType has unexpected decoded type %T", raw)
		}
		p.PacketType = value
	}
	return p, nil
}
