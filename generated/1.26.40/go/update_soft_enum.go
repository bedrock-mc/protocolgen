// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type UpdateSoftEnum struct {
	EnumName   string
	Values     []string
	UpdateType SoftEnumUpdateType
}

func (p *UpdateSoftEnum) Encode(w Encoder) error {
	if err := w.Write("UpdateSoftEnumPacket.Enum Name", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, p.EnumName); err != nil {
		return err
	}
	if err := w.Write("UpdateSoftEnumPacket.Values", Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, p.Values); err != nil {
		return err
	}
	if err := w.Write("UpdateSoftEnumPacket.Update Type", Shape{Kind: "enum", Semantic: "SoftEnumUpdateType", TypeID: "enums/SoftEnumUpdateType", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "Add", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "Remove", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "Replace", Shape: Shape{Kind: "void"}}}}, p.UpdateType); err != nil {
		return err
	}
	return nil
}

func DecodeUpdateSoftEnum(r Decoder) (UpdateSoftEnum, error) {
	var p UpdateSoftEnum
	{
		raw, err := r.Read("UpdateSoftEnumPacket.Enum Name", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(string)
		if !ok {
			return p, fmt.Errorf("field UpdateSoftEnumPacket.Enum Name has unexpected decoded type %T", raw)
		}
		p.EnumName = value
	}
	{
		raw, err := r.Read("UpdateSoftEnumPacket.Values", Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.([]string)
		if !ok {
			return p, fmt.Errorf("field UpdateSoftEnumPacket.Values has unexpected decoded type %T", raw)
		}
		p.Values = value
	}
	{
		raw, err := r.Read("UpdateSoftEnumPacket.Update Type", Shape{Kind: "enum", Semantic: "SoftEnumUpdateType", TypeID: "enums/SoftEnumUpdateType", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "Add", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "Remove", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "Replace", Shape: Shape{Kind: "void"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(SoftEnumUpdateType)
		if !ok {
			return p, fmt.Errorf("field UpdateSoftEnumPacket.Update Type has unexpected decoded type %T", raw)
		}
		p.UpdateType = value
	}
	return p, nil
}
