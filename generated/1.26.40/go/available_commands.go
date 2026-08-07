// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type AvailableCommands struct {
	EnumValues              []string
	ChainedSubcommandValues []string
	PostFixes               []string
	EnumData                []AvailableCommandsEnumData
	ChainedSubcommandData   []AvailableCommandsChainedSubcommandData
	Commands                []AvailableCommandsPacketCommandData
	SoftEnums               []AvailableCommandsSoftEnumData
	Constraints             []AvailableCommandsConstrainedValueData
}

func (p *AvailableCommands) Encode(w Encoder) error {
	if err := w.Write("AvailableCommandsPacket.Enum Values", Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, p.EnumValues); err != nil {
		return err
	}
	if err := w.Write("AvailableCommandsPacket.Chained Subcommand Values", Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, p.ChainedSubcommandValues); err != nil {
		return err
	}
	if err := w.Write("AvailableCommandsPacket.Post Fixes", Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, p.PostFixes); err != nil {
		return err
	}
	if err := w.Write("AvailableCommandsPacket.Enum Data", Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "struct", Semantic: "AvailableCommandsPacketPayload::EnumData", TypeID: "AvailableCommandsPacketPayload::EnumData", Fields: []ShapeField{{Ordinal: 0, Name: "Name", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 1, Name: "Values", Shape: Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "primitive", PrimitiveCode: "u32le"}}}}}}, p.EnumData); err != nil {
		return err
	}
	if err := w.Write("AvailableCommandsPacket.Chained Subcommand Data", Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "struct", Semantic: "AvailableCommandsPacketPayload::ChainedSubcommandData", TypeID: "AvailableCommandsPacketPayload::ChainedSubcommandData", Fields: []ShapeField{{Ordinal: 0, Name: "Name", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 1, Name: "SubCommand values", Shape: Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "struct", Semantic: "AvailableCommandsPacketPayload::ChainedSubcommandRelationship", TypeID: "AvailableCommandsPacketPayload::ChainedSubcommandRelationship", Fields: []ShapeField{{Ordinal: 0, Name: "SubCommand First Value", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, {Ordinal: 1, Name: "SubCommand Second Value", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}}}}}}, p.ChainedSubcommandData); err != nil {
		return err
	}
	if err := w.Write("AvailableCommandsPacket.Commands", Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "struct", Semantic: "AvailableCommandsPacketCommandData", TypeID: "AvailableCommandsPacketCommandData.json#", Fields: []ShapeField{{Ordinal: 0, Name: "Name", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 1, Name: "Description", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 2, Name: "Flags", Shape: Shape{Kind: "primitive", PrimitiveCode: "u16le"}}, {Ordinal: 3, Name: "Permission Level", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 4, Name: "Alias Enum", Shape: Shape{Kind: "primitive", PrimitiveCode: "i32le"}}, {Ordinal: 5, Name: "CommandData Chained Subcommand Indexes", Shape: Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "primitive", PrimitiveCode: "u32le"}}}, {Ordinal: 6, Name: "Overloads", Shape: Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "struct", Semantic: "AvailableCommandsPacketPayload::OverloadData", TypeID: "AvailableCommandsPacketPayload::OverloadData", Fields: []ShapeField{{Ordinal: 0, Name: "isChaining", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}, {Ordinal: 1, Name: "Parameter Data", Shape: Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "struct", Semantic: "AvailableCommandsPacketPayload::ParamData", TypeID: "AvailableCommandsPacketPayload::ParamData", Fields: []ShapeField{{Ordinal: 0, Name: "Name", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 1, Name: "Parse Symbol", Shape: Shape{Kind: "primitive", PrimitiveCode: "u32le"}}, {Ordinal: 2, Name: "Is Optional?", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}, {Ordinal: 3, Name: "Options", Shape: Shape{Kind: "primitive", PrimitiveCode: "u8"}}}}}}}}}}}}}, p.Commands); err != nil {
		return err
	}
	if err := w.Write("AvailableCommandsPacket.Soft Enums", Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "struct", Semantic: "AvailableCommandsPacketPayload::SoftEnumData", TypeID: "AvailableCommandsPacketPayload::SoftEnumData", Fields: []ShapeField{{Ordinal: 0, Name: "Enum Name", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 1, Name: "Enum Options", Shape: Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}}}}, p.SoftEnums); err != nil {
		return err
	}
	if err := w.Write("AvailableCommandsPacket.Constraints", Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "struct", Semantic: "AvailableCommandsPacketPayload::ConstrainedValueData", TypeID: "AvailableCommandsPacketPayload::ConstrainedValueData", Fields: []ShapeField{{Ordinal: 0, Name: "Enum Value Symbol", Shape: Shape{Kind: "primitive", PrimitiveCode: "u32le"}}, {Ordinal: 1, Name: "Enum Symbol", Shape: Shape{Kind: "primitive", PrimitiveCode: "u32le"}}, {Ordinal: 2, Name: "Constraint Indices", Shape: Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "primitive", PrimitiveCode: "u8"}}}}}}, p.Constraints); err != nil {
		return err
	}
	return nil
}

func DecodeAvailableCommands(r Decoder) (AvailableCommands, error) {
	var p AvailableCommands
	{
		raw, err := r.Read("AvailableCommandsPacket.Enum Values", Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.([]string)
		if !ok {
			return p, fmt.Errorf("field AvailableCommandsPacket.Enum Values has unexpected decoded type %T", raw)
		}
		p.EnumValues = value
	}
	{
		raw, err := r.Read("AvailableCommandsPacket.Chained Subcommand Values", Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.([]string)
		if !ok {
			return p, fmt.Errorf("field AvailableCommandsPacket.Chained Subcommand Values has unexpected decoded type %T", raw)
		}
		p.ChainedSubcommandValues = value
	}
	{
		raw, err := r.Read("AvailableCommandsPacket.Post Fixes", Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.([]string)
		if !ok {
			return p, fmt.Errorf("field AvailableCommandsPacket.Post Fixes has unexpected decoded type %T", raw)
		}
		p.PostFixes = value
	}
	{
		raw, err := r.Read("AvailableCommandsPacket.Enum Data", Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "struct", Semantic: "AvailableCommandsPacketPayload::EnumData", TypeID: "AvailableCommandsPacketPayload::EnumData", Fields: []ShapeField{{Ordinal: 0, Name: "Name", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 1, Name: "Values", Shape: Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "primitive", PrimitiveCode: "u32le"}}}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.([]AvailableCommandsEnumData)
		if !ok {
			return p, fmt.Errorf("field AvailableCommandsPacket.Enum Data has unexpected decoded type %T", raw)
		}
		p.EnumData = value
	}
	{
		raw, err := r.Read("AvailableCommandsPacket.Chained Subcommand Data", Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "struct", Semantic: "AvailableCommandsPacketPayload::ChainedSubcommandData", TypeID: "AvailableCommandsPacketPayload::ChainedSubcommandData", Fields: []ShapeField{{Ordinal: 0, Name: "Name", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 1, Name: "SubCommand values", Shape: Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "struct", Semantic: "AvailableCommandsPacketPayload::ChainedSubcommandRelationship", TypeID: "AvailableCommandsPacketPayload::ChainedSubcommandRelationship", Fields: []ShapeField{{Ordinal: 0, Name: "SubCommand First Value", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, {Ordinal: 1, Name: "SubCommand Second Value", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}}}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.([]AvailableCommandsChainedSubcommandData)
		if !ok {
			return p, fmt.Errorf("field AvailableCommandsPacket.Chained Subcommand Data has unexpected decoded type %T", raw)
		}
		p.ChainedSubcommandData = value
	}
	{
		raw, err := r.Read("AvailableCommandsPacket.Commands", Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "struct", Semantic: "AvailableCommandsPacketCommandData", TypeID: "AvailableCommandsPacketCommandData.json#", Fields: []ShapeField{{Ordinal: 0, Name: "Name", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 1, Name: "Description", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 2, Name: "Flags", Shape: Shape{Kind: "primitive", PrimitiveCode: "u16le"}}, {Ordinal: 3, Name: "Permission Level", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 4, Name: "Alias Enum", Shape: Shape{Kind: "primitive", PrimitiveCode: "i32le"}}, {Ordinal: 5, Name: "CommandData Chained Subcommand Indexes", Shape: Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "primitive", PrimitiveCode: "u32le"}}}, {Ordinal: 6, Name: "Overloads", Shape: Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "struct", Semantic: "AvailableCommandsPacketPayload::OverloadData", TypeID: "AvailableCommandsPacketPayload::OverloadData", Fields: []ShapeField{{Ordinal: 0, Name: "isChaining", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}, {Ordinal: 1, Name: "Parameter Data", Shape: Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "struct", Semantic: "AvailableCommandsPacketPayload::ParamData", TypeID: "AvailableCommandsPacketPayload::ParamData", Fields: []ShapeField{{Ordinal: 0, Name: "Name", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 1, Name: "Parse Symbol", Shape: Shape{Kind: "primitive", PrimitiveCode: "u32le"}}, {Ordinal: 2, Name: "Is Optional?", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}, {Ordinal: 3, Name: "Options", Shape: Shape{Kind: "primitive", PrimitiveCode: "u8"}}}}}}}}}}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.([]AvailableCommandsPacketCommandData)
		if !ok {
			return p, fmt.Errorf("field AvailableCommandsPacket.Commands has unexpected decoded type %T", raw)
		}
		p.Commands = value
	}
	{
		raw, err := r.Read("AvailableCommandsPacket.Soft Enums", Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "struct", Semantic: "AvailableCommandsPacketPayload::SoftEnumData", TypeID: "AvailableCommandsPacketPayload::SoftEnumData", Fields: []ShapeField{{Ordinal: 0, Name: "Enum Name", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 1, Name: "Enum Options", Shape: Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.([]AvailableCommandsSoftEnumData)
		if !ok {
			return p, fmt.Errorf("field AvailableCommandsPacket.Soft Enums has unexpected decoded type %T", raw)
		}
		p.SoftEnums = value
	}
	{
		raw, err := r.Read("AvailableCommandsPacket.Constraints", Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "struct", Semantic: "AvailableCommandsPacketPayload::ConstrainedValueData", TypeID: "AvailableCommandsPacketPayload::ConstrainedValueData", Fields: []ShapeField{{Ordinal: 0, Name: "Enum Value Symbol", Shape: Shape{Kind: "primitive", PrimitiveCode: "u32le"}}, {Ordinal: 1, Name: "Enum Symbol", Shape: Shape{Kind: "primitive", PrimitiveCode: "u32le"}}, {Ordinal: 2, Name: "Constraint Indices", Shape: Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "primitive", PrimitiveCode: "u8"}}}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.([]AvailableCommandsConstrainedValueData)
		if !ok {
			return p, fmt.Errorf("field AvailableCommandsPacket.Constraints has unexpected decoded type %T", raw)
		}
		p.Constraints = value
	}
	return p, nil
}
