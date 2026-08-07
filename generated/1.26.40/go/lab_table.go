// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type LabTable struct {
	Type     LabTableType
	Position BlockPos
	Reaction LabTableReactionType
}

func (p *LabTable) Encode(w Encoder) error {
	if err := w.Write("LabTablePacket.Type", Shape{Kind: "enum", Semantic: "LabTablePacketPayload::Type", TypeID: "enums/LabTablePacketPayload::Type", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "StartCombine", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "StartReaction", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "Reset", Shape: Shape{Kind: "void"}}}}, p.Type); err != nil {
		return err
	}
	if err := w.Write("LabTablePacket.Position", Shape{Kind: "struct", Semantic: "BlockPos", TypeID: "BlockPos", Fields: []ShapeField{{Ordinal: 0, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 1, Name: "Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 2, Name: "Z", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}}}, p.Position); err != nil {
		return err
	}
	if err := w.Write("LabTablePacket.Reaction", Shape{Kind: "enum", Semantic: "LabTableReactionType", TypeID: "enums/LabTableReactionType", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "None", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "IceBomb", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "Bleach", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "ElephantToothpaste", Shape: Shape{Kind: "void"}}, {Value: 4, Name: "Fertilizer", Shape: Shape{Kind: "void"}}, {Value: 5, Name: "HeatBlock", Shape: Shape{Kind: "void"}}, {Value: 6, Name: "MagnesiumSalts", Shape: Shape{Kind: "void"}}, {Value: 7, Name: "MiscFire", Shape: Shape{Kind: "void"}}, {Value: 8, Name: "MiscExplosion", Shape: Shape{Kind: "void"}}, {Value: 9, Name: "MiscLava", Shape: Shape{Kind: "void"}}, {Value: 10, Name: "MiscMystical", Shape: Shape{Kind: "void"}}, {Value: 11, Name: "MiscSmoke", Shape: Shape{Kind: "void"}}, {Value: 12, Name: "MiscLargeSmoke", Shape: Shape{Kind: "void"}}}}, p.Reaction); err != nil {
		return err
	}
	return nil
}

func DecodeLabTable(r Decoder) (LabTable, error) {
	var p LabTable
	{
		raw, err := r.Read("LabTablePacket.Type", Shape{Kind: "enum", Semantic: "LabTablePacketPayload::Type", TypeID: "enums/LabTablePacketPayload::Type", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "StartCombine", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "StartReaction", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "Reset", Shape: Shape{Kind: "void"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(LabTableType)
		if !ok {
			return p, fmt.Errorf("field LabTablePacket.Type has unexpected decoded type %T", raw)
		}
		p.Type = value
	}
	{
		raw, err := r.Read("LabTablePacket.Position", Shape{Kind: "struct", Semantic: "BlockPos", TypeID: "BlockPos", Fields: []ShapeField{{Ordinal: 0, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 1, Name: "Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 2, Name: "Z", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(BlockPos)
		if !ok {
			return p, fmt.Errorf("field LabTablePacket.Position has unexpected decoded type %T", raw)
		}
		p.Position = value
	}
	{
		raw, err := r.Read("LabTablePacket.Reaction", Shape{Kind: "enum", Semantic: "LabTableReactionType", TypeID: "enums/LabTableReactionType", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "None", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "IceBomb", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "Bleach", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "ElephantToothpaste", Shape: Shape{Kind: "void"}}, {Value: 4, Name: "Fertilizer", Shape: Shape{Kind: "void"}}, {Value: 5, Name: "HeatBlock", Shape: Shape{Kind: "void"}}, {Value: 6, Name: "MagnesiumSalts", Shape: Shape{Kind: "void"}}, {Value: 7, Name: "MiscFire", Shape: Shape{Kind: "void"}}, {Value: 8, Name: "MiscExplosion", Shape: Shape{Kind: "void"}}, {Value: 9, Name: "MiscLava", Shape: Shape{Kind: "void"}}, {Value: 10, Name: "MiscMystical", Shape: Shape{Kind: "void"}}, {Value: 11, Name: "MiscSmoke", Shape: Shape{Kind: "void"}}, {Value: 12, Name: "MiscLargeSmoke", Shape: Shape{Kind: "void"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(LabTableReactionType)
		if !ok {
			return p, fmt.Errorf("field LabTablePacket.Reaction has unexpected decoded type %T", raw)
		}
		p.Reaction = value
	}
	return p, nil
}
