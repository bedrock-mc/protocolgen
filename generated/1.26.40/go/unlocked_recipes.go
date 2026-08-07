// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type UnlockedRecipes struct {
	PacketType          UnlockedRecipesPacketType
	UnlockedRecipesList []string
}

func (p *UnlockedRecipes) Encode(w Encoder) error {
	if err := w.Write("UnlockedRecipesPacket.Packet Type", Shape{Kind: "enum", Semantic: "UnlockedRecipesPacketPayload::PacketType", TypeID: "enums/UnlockedRecipesPacketPayload::PacketType", PrimitiveCode: "u32le", Variants: []ShapeVariant{{Value: 0, Name: "Empty", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "InitiallyUnlockedRecipes", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "NewlyUnlockedRecipes", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "RemoveUnlockedRecipes", Shape: Shape{Kind: "void"}}, {Value: 4, Name: "RemoveAllUnlockedRecipes", Shape: Shape{Kind: "void"}}}}, p.PacketType); err != nil {
		return err
	}
	if err := w.Write("UnlockedRecipesPacket.Unlocked Recipes List", Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, p.UnlockedRecipesList); err != nil {
		return err
	}
	return nil
}

func DecodeUnlockedRecipes(r Decoder) (UnlockedRecipes, error) {
	var p UnlockedRecipes
	{
		raw, err := r.Read("UnlockedRecipesPacket.Packet Type", Shape{Kind: "enum", Semantic: "UnlockedRecipesPacketPayload::PacketType", TypeID: "enums/UnlockedRecipesPacketPayload::PacketType", PrimitiveCode: "u32le", Variants: []ShapeVariant{{Value: 0, Name: "Empty", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "InitiallyUnlockedRecipes", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "NewlyUnlockedRecipes", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "RemoveUnlockedRecipes", Shape: Shape{Kind: "void"}}, {Value: 4, Name: "RemoveAllUnlockedRecipes", Shape: Shape{Kind: "void"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(UnlockedRecipesPacketType)
		if !ok {
			return p, fmt.Errorf("field UnlockedRecipesPacket.Packet Type has unexpected decoded type %T", raw)
		}
		p.PacketType = value
	}
	{
		raw, err := r.Read("UnlockedRecipesPacket.Unlocked Recipes List", Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.([]string)
		if !ok {
			return p, fmt.Errorf("field UnlockedRecipesPacket.Unlocked Recipes List has unexpected decoded type %T", raw)
		}
		p.UnlockedRecipesList = value
	}
	return p, nil
}
