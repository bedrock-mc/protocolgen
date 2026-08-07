// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type UpdateAbilities struct {
	Data SerializedAbilitiesData
}

func (p *UpdateAbilities) Encode(w Encoder) error {
	if err := w.Write("UpdateAbilitiesPacket.Data", Shape{Kind: "struct", Semantic: "SerializedAbilitiesData", TypeID: "SerializedAbilitiesData", Fields: []ShapeField{{Ordinal: 0, Name: "Target Player Raw Id", Shape: Shape{Kind: "primitive", PrimitiveCode: "i64le"}}, {Ordinal: 1, Name: "Player Permissions", Shape: Shape{Kind: "enum", Semantic: "PlayerPermissionLevel", TypeID: "enums/PlayerPermissionLevel", PrimitiveCode: "i8", Variants: []ShapeVariant{{Value: 0, Name: "Visitor", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "Member", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "Operator", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "Custom", Shape: Shape{Kind: "void"}}}}}, {Ordinal: 2, Name: "Command Permissions", Shape: Shape{Kind: "enum", Semantic: "CommandPermissionLevel", TypeID: "enums/CommandPermissionLevel", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "Any", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "GameDirectors", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "Admin", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "Host", Shape: Shape{Kind: "void"}}, {Value: 4, Name: "Owner", Shape: Shape{Kind: "void"}}, {Value: 5, Name: "Internal", Shape: Shape{Kind: "void"}}}}}, {Ordinal: 3, Name: "Layers", Shape: Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "struct", Semantic: "SerializedAbilitiesData::SerializedLayer", TypeID: "SerializedAbilitiesData::SerializedLayer", Fields: []ShapeField{{Ordinal: 0, Name: "SerializedLayer", Shape: Shape{Kind: "primitive", PrimitiveCode: "u16le"}}, {Ordinal: 1, Name: "AbilitiesSet", Shape: Shape{Kind: "primitive", PrimitiveCode: "u32le"}}, {Ordinal: 2, Name: "AbilityValues", Shape: Shape{Kind: "primitive", PrimitiveCode: "u32le"}}, {Ordinal: 3, Name: "FlySpeed", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 4, Name: "VerticalFlySpeed", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 5, Name: "WalkSpeed", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}}}}}}}, p.Data); err != nil {
		return err
	}
	return nil
}

func DecodeUpdateAbilities(r Decoder) (UpdateAbilities, error) {
	var p UpdateAbilities
	{
		raw, err := r.Read("UpdateAbilitiesPacket.Data", Shape{Kind: "struct", Semantic: "SerializedAbilitiesData", TypeID: "SerializedAbilitiesData", Fields: []ShapeField{{Ordinal: 0, Name: "Target Player Raw Id", Shape: Shape{Kind: "primitive", PrimitiveCode: "i64le"}}, {Ordinal: 1, Name: "Player Permissions", Shape: Shape{Kind: "enum", Semantic: "PlayerPermissionLevel", TypeID: "enums/PlayerPermissionLevel", PrimitiveCode: "i8", Variants: []ShapeVariant{{Value: 0, Name: "Visitor", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "Member", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "Operator", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "Custom", Shape: Shape{Kind: "void"}}}}}, {Ordinal: 2, Name: "Command Permissions", Shape: Shape{Kind: "enum", Semantic: "CommandPermissionLevel", TypeID: "enums/CommandPermissionLevel", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "Any", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "GameDirectors", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "Admin", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "Host", Shape: Shape{Kind: "void"}}, {Value: 4, Name: "Owner", Shape: Shape{Kind: "void"}}, {Value: 5, Name: "Internal", Shape: Shape{Kind: "void"}}}}}, {Ordinal: 3, Name: "Layers", Shape: Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "struct", Semantic: "SerializedAbilitiesData::SerializedLayer", TypeID: "SerializedAbilitiesData::SerializedLayer", Fields: []ShapeField{{Ordinal: 0, Name: "SerializedLayer", Shape: Shape{Kind: "primitive", PrimitiveCode: "u16le"}}, {Ordinal: 1, Name: "AbilitiesSet", Shape: Shape{Kind: "primitive", PrimitiveCode: "u32le"}}, {Ordinal: 2, Name: "AbilityValues", Shape: Shape{Kind: "primitive", PrimitiveCode: "u32le"}}, {Ordinal: 3, Name: "FlySpeed", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 4, Name: "VerticalFlySpeed", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 5, Name: "WalkSpeed", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}}}}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(SerializedAbilitiesData)
		if !ok {
			return p, fmt.Errorf("field UpdateAbilitiesPacket.Data has unexpected decoded type %T", raw)
		}
		p.Data = value
	}
	return p, nil
}
