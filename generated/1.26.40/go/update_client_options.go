// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type UpdateClientOptions struct {
	GraphicsModeChange    *GraphicsMode
	FilterProfanityChange *bool
}

func (p *UpdateClientOptions) Encode(w Encoder) error {
	if err := w.Write("UpdateClientOptionsPacket.Graphics Mode Change", Shape{Kind: "optional", Value: &Shape{Kind: "enum", Semantic: "GraphicsMode", TypeID: "enums/GraphicsMode", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "Simple", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "Fancy", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "Advanced", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "RayTraced", Shape: Shape{Kind: "void"}}}}}, p.GraphicsModeChange); err != nil {
		return err
	}
	if err := w.Write("UpdateClientOptionsPacket.Filter Profanity Change", Shape{Kind: "optional", Value: &Shape{Kind: "primitive", PrimitiveCode: "bool"}}, p.FilterProfanityChange); err != nil {
		return err
	}
	return nil
}

func DecodeUpdateClientOptions(r Decoder) (UpdateClientOptions, error) {
	var p UpdateClientOptions
	{
		raw, err := r.Read("UpdateClientOptionsPacket.Graphics Mode Change", Shape{Kind: "optional", Value: &Shape{Kind: "enum", Semantic: "GraphicsMode", TypeID: "enums/GraphicsMode", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "Simple", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "Fancy", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "Advanced", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "RayTraced", Shape: Shape{Kind: "void"}}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(*GraphicsMode)
		if !ok {
			return p, fmt.Errorf("field UpdateClientOptionsPacket.Graphics Mode Change has unexpected decoded type %T", raw)
		}
		p.GraphicsModeChange = value
	}
	{
		raw, err := r.Read("UpdateClientOptionsPacket.Filter Profanity Change", Shape{Kind: "optional", Value: &Shape{Kind: "primitive", PrimitiveCode: "bool"}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(*bool)
		if !ok {
			return p, fmt.Errorf("field UpdateClientOptionsPacket.Filter Profanity Change has unexpected decoded type %T", raw)
		}
		p.FilterProfanityChange = value
	}
	return p, nil
}
