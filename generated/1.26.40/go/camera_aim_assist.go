// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type CameraAimAssist struct {
	PresetId        string
	ViewAngle       Vec2
	Distance        float32
	TargetMode      CameraAimAssistTargetModeType
	Action          CameraAimAssistAction
	ShowDebugRender bool
}

func (p *CameraAimAssist) Encode(w Encoder) error {
	if err := w.Write("CameraAimAssistPacket.Preset Id", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, p.PresetId); err != nil {
		return err
	}
	if err := w.Write("CameraAimAssistPacket.View Angle", Shape{Kind: "struct", Semantic: "Vec2", TypeID: "Vec2", Fields: []ShapeField{{Ordinal: 0, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 1, Name: "Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}}}, p.ViewAngle); err != nil {
		return err
	}
	if err := w.Write("CameraAimAssistPacket.Distance", Shape{Kind: "primitive", PrimitiveCode: "f32le"}, p.Distance); err != nil {
		return err
	}
	if err := w.Write("CameraAimAssistPacket.Target Mode", Shape{Kind: "enum", Semantic: "CameraAimAssistPacketPayload::TargetMode", TypeID: "enums/CameraAimAssistPacketPayload::TargetMode", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "Angle", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "Distance", Shape: Shape{Kind: "void"}}}}, p.TargetMode); err != nil {
		return err
	}
	if err := w.Write("CameraAimAssistPacket.Action", Shape{Kind: "enum", Semantic: "CameraAimAssistPacketPayload::Action", TypeID: "enums/CameraAimAssistPacketPayload::Action", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "Set", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "Clear", Shape: Shape{Kind: "void"}}}}, p.Action); err != nil {
		return err
	}
	if err := w.Write("CameraAimAssistPacket.Show Debug Render", Shape{Kind: "primitive", PrimitiveCode: "bool"}, p.ShowDebugRender); err != nil {
		return err
	}
	return nil
}

func DecodeCameraAimAssist(r Decoder) (CameraAimAssist, error) {
	var p CameraAimAssist
	{
		raw, err := r.Read("CameraAimAssistPacket.Preset Id", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(string)
		if !ok {
			return p, fmt.Errorf("field CameraAimAssistPacket.Preset Id has unexpected decoded type %T", raw)
		}
		p.PresetId = value
	}
	{
		raw, err := r.Read("CameraAimAssistPacket.View Angle", Shape{Kind: "struct", Semantic: "Vec2", TypeID: "Vec2", Fields: []ShapeField{{Ordinal: 0, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 1, Name: "Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(Vec2)
		if !ok {
			return p, fmt.Errorf("field CameraAimAssistPacket.View Angle has unexpected decoded type %T", raw)
		}
		p.ViewAngle = value
	}
	{
		raw, err := r.Read("CameraAimAssistPacket.Distance", Shape{Kind: "primitive", PrimitiveCode: "f32le"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(float32)
		if !ok {
			return p, fmt.Errorf("field CameraAimAssistPacket.Distance has unexpected decoded type %T", raw)
		}
		p.Distance = value
	}
	{
		raw, err := r.Read("CameraAimAssistPacket.Target Mode", Shape{Kind: "enum", Semantic: "CameraAimAssistPacketPayload::TargetMode", TypeID: "enums/CameraAimAssistPacketPayload::TargetMode", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "Angle", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "Distance", Shape: Shape{Kind: "void"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(CameraAimAssistTargetModeType)
		if !ok {
			return p, fmt.Errorf("field CameraAimAssistPacket.Target Mode has unexpected decoded type %T", raw)
		}
		p.TargetMode = value
	}
	{
		raw, err := r.Read("CameraAimAssistPacket.Action", Shape{Kind: "enum", Semantic: "CameraAimAssistPacketPayload::Action", TypeID: "enums/CameraAimAssistPacketPayload::Action", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "Set", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "Clear", Shape: Shape{Kind: "void"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(CameraAimAssistAction)
		if !ok {
			return p, fmt.Errorf("field CameraAimAssistPacket.Action has unexpected decoded type %T", raw)
		}
		p.Action = value
	}
	{
		raw, err := r.Read("CameraAimAssistPacket.Show Debug Render", Shape{Kind: "primitive", PrimitiveCode: "bool"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(bool)
		if !ok {
			return p, fmt.Errorf("field CameraAimAssistPacket.Show Debug Render has unexpected decoded type %T", raw)
		}
		p.ShowDebugRender = value
	}
	return p, nil
}
