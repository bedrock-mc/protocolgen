// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type ClientCameraAimAssist struct {
	CameraPresetId string
	Action         ClientCameraAimAssistPacketAction
	AllowAimAssist bool
}

func (p *ClientCameraAimAssist) Encode(w Encoder) error {
	if err := w.Write("ClientCameraAimAssistPacket.Camera Preset Id", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, p.CameraPresetId); err != nil {
		return err
	}
	if err := w.Write("ClientCameraAimAssistPacket.Action", Shape{Kind: "enum", Semantic: "ClientCameraAimAssistPacketAction", TypeID: "enums/ClientCameraAimAssistPacketAction", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "SetFromCameraPreset", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "Clear", Shape: Shape{Kind: "void"}}}}, p.Action); err != nil {
		return err
	}
	if err := w.Write("ClientCameraAimAssistPacket.Allow aim assist", Shape{Kind: "primitive", PrimitiveCode: "bool"}, p.AllowAimAssist); err != nil {
		return err
	}
	return nil
}

func DecodeClientCameraAimAssist(r Decoder) (ClientCameraAimAssist, error) {
	var p ClientCameraAimAssist
	{
		raw, err := r.Read("ClientCameraAimAssistPacket.Camera Preset Id", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(string)
		if !ok {
			return p, fmt.Errorf("field ClientCameraAimAssistPacket.Camera Preset Id has unexpected decoded type %T", raw)
		}
		p.CameraPresetId = value
	}
	{
		raw, err := r.Read("ClientCameraAimAssistPacket.Action", Shape{Kind: "enum", Semantic: "ClientCameraAimAssistPacketAction", TypeID: "enums/ClientCameraAimAssistPacketAction", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "SetFromCameraPreset", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "Clear", Shape: Shape{Kind: "void"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(ClientCameraAimAssistPacketAction)
		if !ok {
			return p, fmt.Errorf("field ClientCameraAimAssistPacket.Action has unexpected decoded type %T", raw)
		}
		p.Action = value
	}
	{
		raw, err := r.Read("ClientCameraAimAssistPacket.Allow aim assist", Shape{Kind: "primitive", PrimitiveCode: "bool"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(bool)
		if !ok {
			return p, fmt.Errorf("field ClientCameraAimAssistPacket.Allow aim assist has unexpected decoded type %T", raw)
		}
		p.AllowAimAssist = value
	}
	return p, nil
}
