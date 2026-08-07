// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type PlayerArmorDamage struct {
	ArmorSlotAndDamagePairs []ArmorSlotAndDamagePair
}

func (p *PlayerArmorDamage) Encode(w Encoder) error {
	if err := w.Write("PlayerArmorDamagePacket.Armor Slot and Damage Pairs", Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "struct", Semantic: "ArmorSlotAndDamagePair", TypeID: "ArmorSlotAndDamagePair", Fields: []ShapeField{{Ordinal: 0, Name: "Armor Slot", Shape: Shape{Kind: "enum", Semantic: "SharedTypes::Legacy::ArmorSlot", TypeID: "enums/SharedTypes::Legacy::ArmorSlot", PrimitiveCode: "zigzag_i32", Variants: []ShapeVariant{{Value: 0, Name: "Head", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "Torso", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "Legs", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "Feet", Shape: Shape{Kind: "void"}}, {Value: 4, Name: "Body", Shape: Shape{Kind: "void"}}}}}, {Ordinal: 1, Name: "Damage", Shape: Shape{Kind: "primitive", PrimitiveCode: "i16le"}}}}}, p.ArmorSlotAndDamagePairs); err != nil {
		return err
	}
	return nil
}

func DecodePlayerArmorDamage(r Decoder) (PlayerArmorDamage, error) {
	var p PlayerArmorDamage
	{
		raw, err := r.Read("PlayerArmorDamagePacket.Armor Slot and Damage Pairs", Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "struct", Semantic: "ArmorSlotAndDamagePair", TypeID: "ArmorSlotAndDamagePair", Fields: []ShapeField{{Ordinal: 0, Name: "Armor Slot", Shape: Shape{Kind: "enum", Semantic: "SharedTypes::Legacy::ArmorSlot", TypeID: "enums/SharedTypes::Legacy::ArmorSlot", PrimitiveCode: "zigzag_i32", Variants: []ShapeVariant{{Value: 0, Name: "Head", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "Torso", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "Legs", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "Feet", Shape: Shape{Kind: "void"}}, {Value: 4, Name: "Body", Shape: Shape{Kind: "void"}}}}}, {Ordinal: 1, Name: "Damage", Shape: Shape{Kind: "primitive", PrimitiveCode: "i16le"}}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.([]ArmorSlotAndDamagePair)
		if !ok {
			return p, fmt.Errorf("field PlayerArmorDamagePacket.Armor Slot and Damage Pairs has unexpected decoded type %T", raw)
		}
		p.ArmorSlotAndDamagePairs = value
	}
	return p, nil
}
