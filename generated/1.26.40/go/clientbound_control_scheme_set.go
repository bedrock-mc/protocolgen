// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type ClientboundControlSchemeSet struct {
	ControlScheme ControlSchemeScheme
}

func (p *ClientboundControlSchemeSet) Encode(w Encoder) error {
	if err := w.Write("ClientboundControlSchemeSetPacket.Control Scheme", Shape{Kind: "enum", Semantic: "ControlScheme::Scheme", TypeID: "enums/ControlScheme::Scheme", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "locked_player_relative_strafe", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "camera_relative", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "camera_relative_strafe", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "player_relative", Shape: Shape{Kind: "void"}}, {Value: 4, Name: "player_relative_strafe", Shape: Shape{Kind: "void"}}}}, p.ControlScheme); err != nil {
		return err
	}
	return nil
}

func DecodeClientboundControlSchemeSet(r Decoder) (ClientboundControlSchemeSet, error) {
	var p ClientboundControlSchemeSet
	{
		raw, err := r.Read("ClientboundControlSchemeSetPacket.Control Scheme", Shape{Kind: "enum", Semantic: "ControlScheme::Scheme", TypeID: "enums/ControlScheme::Scheme", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "locked_player_relative_strafe", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "camera_relative", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "camera_relative_strafe", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "player_relative", Shape: Shape{Kind: "void"}}, {Value: 4, Name: "player_relative_strafe", Shape: Shape{Kind: "void"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(ControlSchemeScheme)
		if !ok {
			return p, fmt.Errorf("field ClientboundControlSchemeSetPacket.Control Scheme has unexpected decoded type %T", raw)
		}
		p.ControlScheme = value
	}
	return p, nil
}
