// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type CameraShake struct {
	Intensity   float32
	Seconds     float32
	ShakeType   CameraShakeType
	ShakeAction CameraShakeAction
}

func (p *CameraShake) Encode(w Encoder) error {
	if err := w.Write("CameraShakePacket.Intensity", Shape{Kind: "primitive", PrimitiveCode: "f32le"}, p.Intensity); err != nil {
		return err
	}
	if err := w.Write("CameraShakePacket.Seconds", Shape{Kind: "primitive", PrimitiveCode: "f32le"}, p.Seconds); err != nil {
		return err
	}
	if err := w.Write("CameraShakePacket.Shake Type", Shape{Kind: "enum", Semantic: "CameraShakeType", TypeID: "enums/CameraShakeType", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "Positional", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "Rotational", Shape: Shape{Kind: "void"}}}}, p.ShakeType); err != nil {
		return err
	}
	if err := w.Write("CameraShakePacket.Shake Action", Shape{Kind: "enum", Semantic: "CameraShakeAction", TypeID: "enums/CameraShakeAction", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "Add", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "Stop", Shape: Shape{Kind: "void"}}}}, p.ShakeAction); err != nil {
		return err
	}
	return nil
}

func DecodeCameraShake(r Decoder) (CameraShake, error) {
	var p CameraShake
	{
		raw, err := r.Read("CameraShakePacket.Intensity", Shape{Kind: "primitive", PrimitiveCode: "f32le"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(float32)
		if !ok {
			return p, fmt.Errorf("field CameraShakePacket.Intensity has unexpected decoded type %T", raw)
		}
		p.Intensity = value
	}
	{
		raw, err := r.Read("CameraShakePacket.Seconds", Shape{Kind: "primitive", PrimitiveCode: "f32le"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(float32)
		if !ok {
			return p, fmt.Errorf("field CameraShakePacket.Seconds has unexpected decoded type %T", raw)
		}
		p.Seconds = value
	}
	{
		raw, err := r.Read("CameraShakePacket.Shake Type", Shape{Kind: "enum", Semantic: "CameraShakeType", TypeID: "enums/CameraShakeType", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "Positional", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "Rotational", Shape: Shape{Kind: "void"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(CameraShakeType)
		if !ok {
			return p, fmt.Errorf("field CameraShakePacket.Shake Type has unexpected decoded type %T", raw)
		}
		p.ShakeType = value
	}
	{
		raw, err := r.Read("CameraShakePacket.Shake Action", Shape{Kind: "enum", Semantic: "CameraShakeAction", TypeID: "enums/CameraShakeAction", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "Add", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "Stop", Shape: Shape{Kind: "void"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(CameraShakeAction)
		if !ok {
			return p, fmt.Errorf("field CameraShakePacket.Shake Action has unexpected decoded type %T", raw)
		}
		p.ShakeAction = value
	}
	return p, nil
}
