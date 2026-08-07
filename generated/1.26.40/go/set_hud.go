// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type SetHud struct {
	HudElement []HudElement
	HudVisible HudVisibility
}

func (p *SetHud) Encode(w Encoder) error {
	if err := w.Write("SetHudPacket.Hud Element", Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "enum", Semantic: "HudElement", TypeID: "enums/HudElement", PrimitiveCode: "zigzag_i32", Variants: []ShapeVariant{{Value: 0, Name: "PaperDoll", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "Armor", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "ToolTips", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "TouchControls", Shape: Shape{Kind: "void"}}, {Value: 4, Name: "Crosshair", Shape: Shape{Kind: "void"}}, {Value: 5, Name: "HotBar", Shape: Shape{Kind: "void"}}, {Value: 6, Name: "Health", Shape: Shape{Kind: "void"}}, {Value: 7, Name: "ProgressBar", Shape: Shape{Kind: "void"}}, {Value: 8, Name: "Hunger", Shape: Shape{Kind: "void"}}, {Value: 9, Name: "AirBubbles", Shape: Shape{Kind: "void"}}, {Value: 10, Name: "HorseHealth", Shape: Shape{Kind: "void"}}, {Value: 11, Name: "StatusEffects", Shape: Shape{Kind: "void"}}, {Value: 12, Name: "ItemText", Shape: Shape{Kind: "void"}}}}}, p.HudElement); err != nil {
		return err
	}
	if err := w.Write("SetHudPacket.Hud Visible", Shape{Kind: "enum", Semantic: "HudVisibility", TypeID: "enums/HudVisibility", PrimitiveCode: "zigzag_i32", Variants: []ShapeVariant{{Value: 0, Name: "Hide", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "Reset", Shape: Shape{Kind: "void"}}}}, p.HudVisible); err != nil {
		return err
	}
	return nil
}

func DecodeSetHud(r Decoder) (SetHud, error) {
	var p SetHud
	{
		raw, err := r.Read("SetHudPacket.Hud Element", Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "enum", Semantic: "HudElement", TypeID: "enums/HudElement", PrimitiveCode: "zigzag_i32", Variants: []ShapeVariant{{Value: 0, Name: "PaperDoll", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "Armor", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "ToolTips", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "TouchControls", Shape: Shape{Kind: "void"}}, {Value: 4, Name: "Crosshair", Shape: Shape{Kind: "void"}}, {Value: 5, Name: "HotBar", Shape: Shape{Kind: "void"}}, {Value: 6, Name: "Health", Shape: Shape{Kind: "void"}}, {Value: 7, Name: "ProgressBar", Shape: Shape{Kind: "void"}}, {Value: 8, Name: "Hunger", Shape: Shape{Kind: "void"}}, {Value: 9, Name: "AirBubbles", Shape: Shape{Kind: "void"}}, {Value: 10, Name: "HorseHealth", Shape: Shape{Kind: "void"}}, {Value: 11, Name: "StatusEffects", Shape: Shape{Kind: "void"}}, {Value: 12, Name: "ItemText", Shape: Shape{Kind: "void"}}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.([]HudElement)
		if !ok {
			return p, fmt.Errorf("field SetHudPacket.Hud Element has unexpected decoded type %T", raw)
		}
		p.HudElement = value
	}
	{
		raw, err := r.Read("SetHudPacket.Hud Visible", Shape{Kind: "enum", Semantic: "HudVisibility", TypeID: "enums/HudVisibility", PrimitiveCode: "zigzag_i32", Variants: []ShapeVariant{{Value: 0, Name: "Hide", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "Reset", Shape: Shape{Kind: "void"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(HudVisibility)
		if !ok {
			return p, fmt.Errorf("field SetHudPacket.Hud Visible has unexpected decoded type %T", raw)
		}
		p.HudVisible = value
	}
	return p, nil
}
