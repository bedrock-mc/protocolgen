// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type UpdateAttributes struct {
	TargetRuntimeID ActorRuntimeID
	AttributeList   []AttributeData
	Tick            PlayerInputTick
}

func (p *UpdateAttributes) Encode(w Encoder) error {
	if err := w.Write("UpdateAttributesPacket.Target Runtime ID", Shape{Kind: "struct", Semantic: "ActorRuntimeID", TypeID: "ActorRuntimeID", Fields: []ShapeField{{Ordinal: 0, Name: "Actor Runtime ID", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u64"}}}}, p.TargetRuntimeID); err != nil {
		return err
	}
	if err := w.Write("UpdateAttributesPacket.Attribute List", Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "struct", Semantic: "AttributeData", TypeID: "AttributeData", Fields: []ShapeField{{Ordinal: 0, Name: "Min Value", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 1, Name: "Max Value", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 2, Name: "Current Value", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 3, Name: "Default Min Value", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 4, Name: "Default Max Value", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 5, Name: "Default Value", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 6, Name: "Name", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 7, Name: "Modifiers", Shape: Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "struct", Semantic: "AttributeModifier", TypeID: "AttributeModifier", Fields: []ShapeField{{Ordinal: 0, Name: "Id", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 1, Name: "Name", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 2, Name: "Amount", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 3, Name: "Operation", Shape: Shape{Kind: "primitive", PrimitiveCode: "i32le"}}, {Ordinal: 4, Name: "Operand", Shape: Shape{Kind: "primitive", PrimitiveCode: "i32le"}}, {Ordinal: 5, Name: "Is Serializable?", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}}}}}}}}, p.AttributeList); err != nil {
		return err
	}
	if err := w.Write("UpdateAttributesPacket.Tick", Shape{Kind: "struct", Semantic: "PlayerInputTick", TypeID: "PlayerInputTick", Fields: []ShapeField{{Ordinal: 0, Name: "Input tick", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u64"}}}}, p.Tick); err != nil {
		return err
	}
	return nil
}

func DecodeUpdateAttributes(r Decoder) (UpdateAttributes, error) {
	var p UpdateAttributes
	{
		raw, err := r.Read("UpdateAttributesPacket.Target Runtime ID", Shape{Kind: "struct", Semantic: "ActorRuntimeID", TypeID: "ActorRuntimeID", Fields: []ShapeField{{Ordinal: 0, Name: "Actor Runtime ID", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u64"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(ActorRuntimeID)
		if !ok {
			return p, fmt.Errorf("field UpdateAttributesPacket.Target Runtime ID has unexpected decoded type %T", raw)
		}
		p.TargetRuntimeID = value
	}
	{
		raw, err := r.Read("UpdateAttributesPacket.Attribute List", Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "struct", Semantic: "AttributeData", TypeID: "AttributeData", Fields: []ShapeField{{Ordinal: 0, Name: "Min Value", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 1, Name: "Max Value", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 2, Name: "Current Value", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 3, Name: "Default Min Value", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 4, Name: "Default Max Value", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 5, Name: "Default Value", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 6, Name: "Name", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 7, Name: "Modifiers", Shape: Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "struct", Semantic: "AttributeModifier", TypeID: "AttributeModifier", Fields: []ShapeField{{Ordinal: 0, Name: "Id", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 1, Name: "Name", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 2, Name: "Amount", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 3, Name: "Operation", Shape: Shape{Kind: "primitive", PrimitiveCode: "i32le"}}, {Ordinal: 4, Name: "Operand", Shape: Shape{Kind: "primitive", PrimitiveCode: "i32le"}}, {Ordinal: 5, Name: "Is Serializable?", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}}}}}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.([]AttributeData)
		if !ok {
			return p, fmt.Errorf("field UpdateAttributesPacket.Attribute List has unexpected decoded type %T", raw)
		}
		p.AttributeList = value
	}
	{
		raw, err := r.Read("UpdateAttributesPacket.Tick", Shape{Kind: "struct", Semantic: "PlayerInputTick", TypeID: "PlayerInputTick", Fields: []ShapeField{{Ordinal: 0, Name: "Input tick", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u64"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(PlayerInputTick)
		if !ok {
			return p, fmt.Errorf("field UpdateAttributesPacket.Tick has unexpected decoded type %T", raw)
		}
		p.Tick = value
	}
	return p, nil
}
