// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type CreativeContent struct {
	Groups  []CreativeGroupInfo
	Entries []CreativeItemEntry
}

func (p *CreativeContent) Encode(w Encoder) error {
	if err := w.Write("CreativeContentPacket.Groups", Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "struct", Semantic: "CreativeGroupInfoPayload", TypeID: "CreativeGroupInfoPayload", Fields: []ShapeField{{Ordinal: 0, Name: "Creative Category", Shape: Shape{Kind: "enum", Semantic: "SharedTypes::CreativeItemCategory", TypeID: "enums/SharedTypes::CreativeItemCategory", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 1, Name: "Construction", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "Nature", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "Equipment", Shape: Shape{Kind: "void"}}, {Value: 4, Name: "Items", Shape: Shape{Kind: "void"}}, {Value: 5, Name: "ItemCommandOnly", Shape: Shape{Kind: "void"}}}}}, {Ordinal: 1, Name: "Name", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 2, Name: "Group Icon Item", Shape: Shape{Kind: "struct", Semantic: "cerealizer<NetworkItemInstanceDescriptor>::SerializedData", TypeID: "cerealizer<NetworkItemInstanceDescriptor>::SerializedData", Fields: []ShapeField{{Ordinal: 0, Name: "Id", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 1, Name: "Stack size", Shape: Shape{Kind: "primitive", PrimitiveCode: "u16le"}}, {Ordinal: 2, Name: "Aux value", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, {Ordinal: 3, Name: "Block Runtime Id", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 4, Name: "User Data Buffer", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}}}}}}, p.Groups); err != nil {
		return err
	}
	if err := w.Write("CreativeContentPacket.Entries", Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "struct", Semantic: "CreativeItemEntryPayload", TypeID: "CreativeItemEntryPayload", Fields: []ShapeField{{Ordinal: 0, Name: "Creative Net Id", Shape: Shape{Kind: "struct", Semantic: "TypedServerNetId<struct CreativeItemNetIdTag>", TypeID: "TypedServerNetId<struct CreativeItemNetIdTag>", Fields: []ShapeField{{Ordinal: 0, Name: "ID", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}}, {Ordinal: 1, Name: "Item Instance", Shape: Shape{Kind: "struct", Semantic: "cerealizer<NetworkItemInstanceDescriptor>::SerializedData", TypeID: "cerealizer<NetworkItemInstanceDescriptor>::SerializedData", Fields: []ShapeField{{Ordinal: 0, Name: "Id", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 1, Name: "Stack size", Shape: Shape{Kind: "primitive", PrimitiveCode: "u16le"}}, {Ordinal: 2, Name: "Aux value", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, {Ordinal: 3, Name: "Block Runtime Id", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 4, Name: "User Data Buffer", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}}}, {Ordinal: 2, Name: "Group Index", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}}, p.Entries); err != nil {
		return err
	}
	return nil
}

func DecodeCreativeContent(r Decoder) (CreativeContent, error) {
	var p CreativeContent
	{
		raw, err := r.Read("CreativeContentPacket.Groups", Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "struct", Semantic: "CreativeGroupInfoPayload", TypeID: "CreativeGroupInfoPayload", Fields: []ShapeField{{Ordinal: 0, Name: "Creative Category", Shape: Shape{Kind: "enum", Semantic: "SharedTypes::CreativeItemCategory", TypeID: "enums/SharedTypes::CreativeItemCategory", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 1, Name: "Construction", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "Nature", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "Equipment", Shape: Shape{Kind: "void"}}, {Value: 4, Name: "Items", Shape: Shape{Kind: "void"}}, {Value: 5, Name: "ItemCommandOnly", Shape: Shape{Kind: "void"}}}}}, {Ordinal: 1, Name: "Name", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 2, Name: "Group Icon Item", Shape: Shape{Kind: "struct", Semantic: "cerealizer<NetworkItemInstanceDescriptor>::SerializedData", TypeID: "cerealizer<NetworkItemInstanceDescriptor>::SerializedData", Fields: []ShapeField{{Ordinal: 0, Name: "Id", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 1, Name: "Stack size", Shape: Shape{Kind: "primitive", PrimitiveCode: "u16le"}}, {Ordinal: 2, Name: "Aux value", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, {Ordinal: 3, Name: "Block Runtime Id", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 4, Name: "User Data Buffer", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}}}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.([]CreativeGroupInfo)
		if !ok {
			return p, fmt.Errorf("field CreativeContentPacket.Groups has unexpected decoded type %T", raw)
		}
		p.Groups = value
	}
	{
		raw, err := r.Read("CreativeContentPacket.Entries", Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "struct", Semantic: "CreativeItemEntryPayload", TypeID: "CreativeItemEntryPayload", Fields: []ShapeField{{Ordinal: 0, Name: "Creative Net Id", Shape: Shape{Kind: "struct", Semantic: "TypedServerNetId<struct CreativeItemNetIdTag>", TypeID: "TypedServerNetId<struct CreativeItemNetIdTag>", Fields: []ShapeField{{Ordinal: 0, Name: "ID", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}}, {Ordinal: 1, Name: "Item Instance", Shape: Shape{Kind: "struct", Semantic: "cerealizer<NetworkItemInstanceDescriptor>::SerializedData", TypeID: "cerealizer<NetworkItemInstanceDescriptor>::SerializedData", Fields: []ShapeField{{Ordinal: 0, Name: "Id", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 1, Name: "Stack size", Shape: Shape{Kind: "primitive", PrimitiveCode: "u16le"}}, {Ordinal: 2, Name: "Aux value", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, {Ordinal: 3, Name: "Block Runtime Id", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 4, Name: "User Data Buffer", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}}}, {Ordinal: 2, Name: "Group Index", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.([]CreativeItemEntry)
		if !ok {
			return p, fmt.Errorf("field CreativeContentPacket.Entries has unexpected decoded type %T", raw)
		}
		p.Entries = value
	}
	return p, nil
}
