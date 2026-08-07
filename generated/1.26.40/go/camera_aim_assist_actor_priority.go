// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type CameraAimAssistActorPriority struct {
	CameraAimAssistActorPriorityList []CameraAimAssistActorPriorityPriorityData
}

func (p *CameraAimAssistActorPriority) Encode(w Encoder) error {
	if err := w.Write("CameraAimAssistActorPriorityPacket.Camera Aim-Assist Actor Priority List", Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "struct", Semantic: "CameraAimAssistActorPriority::PriorityData", TypeID: "CameraAimAssistActorPriority::PriorityData", Fields: []ShapeField{{Ordinal: 0, Name: "Preset Index", Shape: Shape{Kind: "primitive", PrimitiveCode: "i32le"}}, {Ordinal: 1, Name: "Category Index", Shape: Shape{Kind: "primitive", PrimitiveCode: "i32le"}}, {Ordinal: 2, Name: "Actor Index", Shape: Shape{Kind: "primitive", PrimitiveCode: "i32le"}}, {Ordinal: 3, Name: "Priority Value", Shape: Shape{Kind: "primitive", PrimitiveCode: "i32le"}}}}}, p.CameraAimAssistActorPriorityList); err != nil {
		return err
	}
	return nil
}

func DecodeCameraAimAssistActorPriority(r Decoder) (CameraAimAssistActorPriority, error) {
	var p CameraAimAssistActorPriority
	{
		raw, err := r.Read("CameraAimAssistActorPriorityPacket.Camera Aim-Assist Actor Priority List", Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "struct", Semantic: "CameraAimAssistActorPriority::PriorityData", TypeID: "CameraAimAssistActorPriority::PriorityData", Fields: []ShapeField{{Ordinal: 0, Name: "Preset Index", Shape: Shape{Kind: "primitive", PrimitiveCode: "i32le"}}, {Ordinal: 1, Name: "Category Index", Shape: Shape{Kind: "primitive", PrimitiveCode: "i32le"}}, {Ordinal: 2, Name: "Actor Index", Shape: Shape{Kind: "primitive", PrimitiveCode: "i32le"}}, {Ordinal: 3, Name: "Priority Value", Shape: Shape{Kind: "primitive", PrimitiveCode: "i32le"}}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.([]CameraAimAssistActorPriorityPriorityData)
		if !ok {
			return p, fmt.Errorf("field CameraAimAssistActorPriorityPacket.Camera Aim-Assist Actor Priority List has unexpected decoded type %T", raw)
		}
		p.CameraAimAssistActorPriorityList = value
	}
	return p, nil
}
