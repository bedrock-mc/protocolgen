// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type FeatureRegistry struct {
	FeaturesDataList []FeatureRegistryFeatureBinaryJsonFormat
}

func (p *FeatureRegistry) Encode(w Encoder) error {
	if err := w.Write("FeatureRegistryPacket.FeaturesDataList", Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "struct", Semantic: "FeatureRegistry::FeatureBinaryJsonFormat", TypeID: "FeatureRegistry::FeatureBinaryJsonFormat", Fields: []ShapeField{{Ordinal: 0, Name: "Feature Name", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 1, Name: "Binary Json Output", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}}}, p.FeaturesDataList); err != nil {
		return err
	}
	return nil
}

func DecodeFeatureRegistry(r Decoder) (FeatureRegistry, error) {
	var p FeatureRegistry
	{
		raw, err := r.Read("FeatureRegistryPacket.FeaturesDataList", Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "struct", Semantic: "FeatureRegistry::FeatureBinaryJsonFormat", TypeID: "FeatureRegistry::FeatureBinaryJsonFormat", Fields: []ShapeField{{Ordinal: 0, Name: "Feature Name", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 1, Name: "Binary Json Output", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.([]FeatureRegistryFeatureBinaryJsonFormat)
		if !ok {
			return p, fmt.Errorf("field FeatureRegistryPacket.FeaturesDataList has unexpected decoded type %T", raw)
		}
		p.FeaturesDataList = value
	}
	return p, nil
}
