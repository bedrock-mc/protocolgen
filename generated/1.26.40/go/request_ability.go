// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type RequestAbility struct {
	Ability   int32
	ValueType RequestAbilityType
	Bool      bool
	Float     float32
}

func (p *RequestAbility) Encode(w Encoder) error {
	if err := w.Write("RequestAbilityPacket.Ability", Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}, p.Ability); err != nil {
		return err
	}
	if err := w.Write("RequestAbilityPacket.Value Type", Shape{Kind: "enum", Semantic: "RequestAbilityPacketPayload::Type", TypeID: "enums/RequestAbilityPacketPayload::Type", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "Unset", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "Bool", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "Float", Shape: Shape{Kind: "void"}}}}, p.ValueType); err != nil {
		return err
	}
	if err := w.Write("RequestAbilityPacket.Bool", Shape{Kind: "primitive", PrimitiveCode: "bool"}, p.Bool); err != nil {
		return err
	}
	if err := w.Write("RequestAbilityPacket.Float", Shape{Kind: "primitive", PrimitiveCode: "f32le"}, p.Float); err != nil {
		return err
	}
	return nil
}

func DecodeRequestAbility(r Decoder) (RequestAbility, error) {
	var p RequestAbility
	{
		raw, err := r.Read("RequestAbilityPacket.Ability", Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(int32)
		if !ok {
			return p, fmt.Errorf("field RequestAbilityPacket.Ability has unexpected decoded type %T", raw)
		}
		p.Ability = value
	}
	{
		raw, err := r.Read("RequestAbilityPacket.Value Type", Shape{Kind: "enum", Semantic: "RequestAbilityPacketPayload::Type", TypeID: "enums/RequestAbilityPacketPayload::Type", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "Unset", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "Bool", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "Float", Shape: Shape{Kind: "void"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(RequestAbilityType)
		if !ok {
			return p, fmt.Errorf("field RequestAbilityPacket.Value Type has unexpected decoded type %T", raw)
		}
		p.ValueType = value
	}
	{
		raw, err := r.Read("RequestAbilityPacket.Bool", Shape{Kind: "primitive", PrimitiveCode: "bool"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(bool)
		if !ok {
			return p, fmt.Errorf("field RequestAbilityPacket.Bool has unexpected decoded type %T", raw)
		}
		p.Bool = value
	}
	{
		raw, err := r.Read("RequestAbilityPacket.Float", Shape{Kind: "primitive", PrimitiveCode: "f32le"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(float32)
		if !ok {
			return p, fmt.Errorf("field RequestAbilityPacket.Float has unexpected decoded type %T", raw)
		}
		p.Float = value
	}
	return p, nil
}
