// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type EduUriResource struct {
	EduSharedURIResource EduSharedUriResource
}

func (p *EduUriResource) Encode(w Encoder) error {
	if err := w.Write("EduUriResourcePacket.Edu Shared URI Resource", Shape{Kind: "struct", Semantic: "EduSharedUriResource", TypeID: "EduSharedUriResource", Fields: []ShapeField{{Ordinal: 0, Name: "Button Name", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 1, Name: "Link Uri", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}}, p.EduSharedURIResource); err != nil {
		return err
	}
	return nil
}

func DecodeEduUriResource(r Decoder) (EduUriResource, error) {
	var p EduUriResource
	{
		raw, err := r.Read("EduUriResourcePacket.Edu Shared URI Resource", Shape{Kind: "struct", Semantic: "EduSharedUriResource", TypeID: "EduSharedUriResource", Fields: []ShapeField{{Ordinal: 0, Name: "Button Name", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 1, Name: "Link Uri", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(EduSharedUriResource)
		if !ok {
			return p, fmt.Errorf("field EduUriResourcePacket.Edu Shared URI Resource has unexpected decoded type %T", raw)
		}
		p.EduSharedURIResource = value
	}
	return p, nil
}
