// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type CommandOutput struct {
	OriginData CommandOriginData
	Output     CommandOutputData
}

func (p *CommandOutput) Encode(w Encoder) error {
	if err := w.Write("CommandOutputPacket.Origin Data", Shape{Kind: "struct", Semantic: "CommandOriginData", TypeID: "CommandOriginData.json#", Fields: []ShapeField{{Ordinal: 0, Name: "Type", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 1, Name: "UUID", Shape: Shape{Kind: "primitive", Semantic: "mce::UUID", TypeID: "mce__UUID.json#", PrimitiveCode: "uuid"}}, {Ordinal: 2, Name: "RequestId", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 3, Name: "PlayerId", Shape: Shape{Kind: "primitive", PrimitiveCode: "i64le"}}}}, p.OriginData); err != nil {
		return err
	}
	if err := w.Write("CommandOutputPacket.Output", Shape{Kind: "struct", Semantic: "CommandOutput", TypeID: "CommandOutput", Fields: []ShapeField{{Ordinal: 0, Name: "Output Type", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 1, Name: "Success Count", Shape: Shape{Kind: "primitive", PrimitiveCode: "u32le"}}, {Ordinal: 2, Name: "Output Messages", Shape: Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "struct", Semantic: "CommandOutputMessage", TypeID: "CommandOutputMessage", Fields: []ShapeField{{Ordinal: 0, Name: "Message ID", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 1, Name: "Successful?", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}, {Ordinal: 2, Name: "Parameters", Shape: Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}}}}}, {Ordinal: 3, Name: "Data Set", Shape: Shape{Kind: "optional", Value: &Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}}}, p.Output); err != nil {
		return err
	}
	return nil
}

func DecodeCommandOutput(r Decoder) (CommandOutput, error) {
	var p CommandOutput
	{
		raw, err := r.Read("CommandOutputPacket.Origin Data", Shape{Kind: "struct", Semantic: "CommandOriginData", TypeID: "CommandOriginData.json#", Fields: []ShapeField{{Ordinal: 0, Name: "Type", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 1, Name: "UUID", Shape: Shape{Kind: "primitive", Semantic: "mce::UUID", TypeID: "mce__UUID.json#", PrimitiveCode: "uuid"}}, {Ordinal: 2, Name: "RequestId", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 3, Name: "PlayerId", Shape: Shape{Kind: "primitive", PrimitiveCode: "i64le"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(CommandOriginData)
		if !ok {
			return p, fmt.Errorf("field CommandOutputPacket.Origin Data has unexpected decoded type %T", raw)
		}
		p.OriginData = value
	}
	{
		raw, err := r.Read("CommandOutputPacket.Output", Shape{Kind: "struct", Semantic: "CommandOutput", TypeID: "CommandOutput", Fields: []ShapeField{{Ordinal: 0, Name: "Output Type", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 1, Name: "Success Count", Shape: Shape{Kind: "primitive", PrimitiveCode: "u32le"}}, {Ordinal: 2, Name: "Output Messages", Shape: Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "struct", Semantic: "CommandOutputMessage", TypeID: "CommandOutputMessage", Fields: []ShapeField{{Ordinal: 0, Name: "Message ID", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 1, Name: "Successful?", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}, {Ordinal: 2, Name: "Parameters", Shape: Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}}}}}, {Ordinal: 3, Name: "Data Set", Shape: Shape{Kind: "optional", Value: &Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(CommandOutputData)
		if !ok {
			return p, fmt.Errorf("field CommandOutputPacket.Output has unexpected decoded type %T", raw)
		}
		p.Output = value
	}
	return p, nil
}
