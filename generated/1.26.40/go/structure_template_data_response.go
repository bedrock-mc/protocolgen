// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type StructureTemplateDataResponse struct {
	StructureName string
	StructureSNBT []byte
	ResponseType  StructureTemplateResponseType
}

func (p *StructureTemplateDataResponse) Encode(w Encoder) error {
	if err := w.Write("StructureTemplateDataResponsePacket.Structure Name", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, p.StructureName); err != nil {
		return err
	}
	if err := w.Write("StructureTemplateDataResponsePacket.Structure's NBT", Shape{Kind: "primitive", PrimitiveCode: "nbt_le"}, p.StructureSNBT); err != nil {
		return err
	}
	if err := w.Write("StructureTemplateDataResponsePacket.Response Type", Shape{Kind: "enum", Semantic: "StructureTemplateResponseType", TypeID: "enums/StructureTemplateResponseType", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "None", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "Export", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "Query", Shape: Shape{Kind: "void"}}}}, p.ResponseType); err != nil {
		return err
	}
	return nil
}

func DecodeStructureTemplateDataResponse(r Decoder) (StructureTemplateDataResponse, error) {
	var p StructureTemplateDataResponse
	{
		raw, err := r.Read("StructureTemplateDataResponsePacket.Structure Name", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(string)
		if !ok {
			return p, fmt.Errorf("field StructureTemplateDataResponsePacket.Structure Name has unexpected decoded type %T", raw)
		}
		p.StructureName = value
	}
	{
		raw, err := r.Read("StructureTemplateDataResponsePacket.Structure's NBT", Shape{Kind: "primitive", PrimitiveCode: "nbt_le"})
		if err != nil {
			return p, err
		}
		value, ok := raw.([]byte)
		if !ok {
			return p, fmt.Errorf("field StructureTemplateDataResponsePacket.Structure's NBT has unexpected decoded type %T", raw)
		}
		p.StructureSNBT = value
	}
	{
		raw, err := r.Read("StructureTemplateDataResponsePacket.Response Type", Shape{Kind: "enum", Semantic: "StructureTemplateResponseType", TypeID: "enums/StructureTemplateResponseType", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "None", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "Export", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "Query", Shape: Shape{Kind: "void"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(StructureTemplateResponseType)
		if !ok {
			return p, fmt.Errorf("field StructureTemplateDataResponsePacket.Response Type has unexpected decoded type %T", raw)
		}
		p.ResponseType = value
	}
	return p, nil
}
