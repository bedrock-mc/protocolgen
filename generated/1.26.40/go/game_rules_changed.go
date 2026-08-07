// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type GameRulesChanged struct {
	RuleData GameRulesChangedPacketData
}

func (p *GameRulesChanged) Encode(w Encoder) error {
	if err := w.Write("GameRulesChangedPacket.Rule Data", Shape{Kind: "struct", Semantic: "GameRulesChangedPacketData", TypeID: "GameRulesChangedPacketData", Fields: []ShapeField{{Ordinal: 0, Name: "Rules List", Shape: Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "struct", Semantic: "GameRule", TypeID: "GameRule", Fields: []ShapeField{{Ordinal: 0, Name: "Rule Name", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 1, Name: "Rule Can Be Modified", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}, {Ordinal: 2, Name: "Rule Value", Shape: Shape{Kind: "union", Control: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Variants: []ShapeVariant{{Value: 0, Name: "Empty0", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "bool", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}, {Value: 2, Name: "int32", Shape: Shape{Kind: "primitive", PrimitiveCode: "i32le"}}, {Value: 3, Name: "float", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}}}}}}}}}}, p.RuleData); err != nil {
		return err
	}
	return nil
}

func DecodeGameRulesChanged(r Decoder) (GameRulesChanged, error) {
	var p GameRulesChanged
	{
		raw, err := r.Read("GameRulesChangedPacket.Rule Data", Shape{Kind: "struct", Semantic: "GameRulesChangedPacketData", TypeID: "GameRulesChangedPacketData", Fields: []ShapeField{{Ordinal: 0, Name: "Rules List", Shape: Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "struct", Semantic: "GameRule", TypeID: "GameRule", Fields: []ShapeField{{Ordinal: 0, Name: "Rule Name", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 1, Name: "Rule Can Be Modified", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}, {Ordinal: 2, Name: "Rule Value", Shape: Shape{Kind: "union", Control: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Variants: []ShapeVariant{{Value: 0, Name: "Empty0", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "bool", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}, {Value: 2, Name: "int32", Shape: Shape{Kind: "primitive", PrimitiveCode: "i32le"}}, {Value: 3, Name: "float", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}}}}}}}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(GameRulesChangedPacketData)
		if !ok {
			return p, fmt.Errorf("field GameRulesChangedPacket.Rule Data has unexpected decoded type %T", raw)
		}
		p.RuleData = value
	}
	return p, nil
}
