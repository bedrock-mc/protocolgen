// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type GraphicsOverrideParameter struct {
	ParameterKeyframeValues []OrderedEntry[float32, Vec3]
	FloatValue              *float32
	Vec3Value               *Vec3
	BiomeIdentifier         string
	PlayerIdentifier        *string
	IdentifierForParameter  GraphicsOverrideParameterType
	ResetParameter          bool
}

func (p *GraphicsOverrideParameter) Encode(w Encoder) error {
	if err := w.Write("GraphicsOverrideParameterPacket.Parameter Keyframe Values", Shape{Kind: "map", Representation: "ordered_entries", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Value: &Shape{Kind: "struct", Semantic: "Vec3", TypeID: "Vec3", Fields: []ShapeField{{Ordinal: 0, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 1, Name: "Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 2, Name: "Z", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}}}, Key: &Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, p.ParameterKeyframeValues); err != nil {
		return err
	}
	if err := w.Write("GraphicsOverrideParameterPacket.Float Value", Shape{Kind: "optional", Value: &Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, p.FloatValue); err != nil {
		return err
	}
	if err := w.Write("GraphicsOverrideParameterPacket.Vec3 Value", Shape{Kind: "optional", Value: &Shape{Kind: "struct", Semantic: "Vec3", TypeID: "Vec3", Fields: []ShapeField{{Ordinal: 0, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 1, Name: "Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 2, Name: "Z", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}}}}, p.Vec3Value); err != nil {
		return err
	}
	if err := w.Write("GraphicsOverrideParameterPacket.Biome Identifier", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, p.BiomeIdentifier); err != nil {
		return err
	}
	if err := w.Write("GraphicsOverrideParameterPacket.Player Identifier", Shape{Kind: "optional", Value: &Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, p.PlayerIdentifier); err != nil {
		return err
	}
	if err := w.Write("GraphicsOverrideParameterPacket.Identifier for Parameter", Shape{Kind: "enum", Semantic: "GraphicsOverrideParameterType", TypeID: "enums/GraphicsOverrideParameterType", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "SkyZenithColor", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "SkyHorizonColor", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "HorizonBlendMin", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "HorizonBlendMax", Shape: Shape{Kind: "void"}}, {Value: 4, Name: "HorizonBlendStart", Shape: Shape{Kind: "void"}}, {Value: 5, Name: "HorizonBlendMieStart", Shape: Shape{Kind: "void"}}, {Value: 6, Name: "RayleighStrength", Shape: Shape{Kind: "void"}}, {Value: 7, Name: "SunMieStrength", Shape: Shape{Kind: "void"}}, {Value: 8, Name: "MoonMieStrength", Shape: Shape{Kind: "void"}}, {Value: 9, Name: "SunGlareShape", Shape: Shape{Kind: "void"}}, {Value: 10, Name: "Chlorophyll", Shape: Shape{Kind: "void"}}, {Value: 11, Name: "CDOM", Shape: Shape{Kind: "void"}}, {Value: 12, Name: "SuspendedSediment", Shape: Shape{Kind: "void"}}, {Value: 13, Name: "WavesDepth", Shape: Shape{Kind: "void"}}, {Value: 14, Name: "WavesFrequency", Shape: Shape{Kind: "void"}}, {Value: 15, Name: "WavesFrequencyScaling", Shape: Shape{Kind: "void"}}, {Value: 16, Name: "WavesSpeed", Shape: Shape{Kind: "void"}}, {Value: 17, Name: "WavesSpeedScaling", Shape: Shape{Kind: "void"}}, {Value: 18, Name: "WavesShape", Shape: Shape{Kind: "void"}}, {Value: 19, Name: "WavesOctaves", Shape: Shape{Kind: "void"}}, {Value: 20, Name: "WavesMix", Shape: Shape{Kind: "void"}}, {Value: 21, Name: "WavesPull", Shape: Shape{Kind: "void"}}, {Value: 22, Name: "WavesDirectionIncrement", Shape: Shape{Kind: "void"}}, {Value: 23, Name: "MidtonesContrast", Shape: Shape{Kind: "void"}}, {Value: 24, Name: "HighlightsContrast", Shape: Shape{Kind: "void"}}, {Value: 25, Name: "ShadowsContrast", Shape: Shape{Kind: "void"}}, {Value: 26, Name: "HighlightsGain", Shape: Shape{Kind: "void"}}, {Value: 27, Name: "HighlightsGamma", Shape: Shape{Kind: "void"}}, {Value: 28, Name: "HighlightsOffset", Shape: Shape{Kind: "void"}}, {Value: 29, Name: "HighlightsSaturation", Shape: Shape{Kind: "void"}}, {Value: 30, Name: "MidtonesGain", Shape: Shape{Kind: "void"}}, {Value: 31, Name: "MidtonesGamma", Shape: Shape{Kind: "void"}}, {Value: 32, Name: "MidtonesOffset", Shape: Shape{Kind: "void"}}, {Value: 33, Name: "MidtonesSaturation", Shape: Shape{Kind: "void"}}, {Value: 34, Name: "ShadowsGain", Shape: Shape{Kind: "void"}}, {Value: 35, Name: "ShadowsGamma", Shape: Shape{Kind: "void"}}, {Value: 36, Name: "ShadowsOffset", Shape: Shape{Kind: "void"}}, {Value: 37, Name: "ShadowsSaturation", Shape: Shape{Kind: "void"}}, {Value: 38, Name: "HighlightsMin", Shape: Shape{Kind: "void"}}, {Value: 39, Name: "ShadowsMax", Shape: Shape{Kind: "void"}}, {Value: 40, Name: "Temperature", Shape: Shape{Kind: "void"}}, {Value: 41, Name: "SunColor", Shape: Shape{Kind: "void"}}, {Value: 42, Name: "SunIlluminance", Shape: Shape{Kind: "void"}}, {Value: 43, Name: "MoonColor", Shape: Shape{Kind: "void"}}, {Value: 44, Name: "MoonIlluminance", Shape: Shape{Kind: "void"}}, {Value: 45, Name: "FlashColor", Shape: Shape{Kind: "void"}}, {Value: 46, Name: "FlashIlluminance", Shape: Shape{Kind: "void"}}, {Value: 47, Name: "AmbientColor", Shape: Shape{Kind: "void"}}, {Value: 48, Name: "AmbientIlluminance", Shape: Shape{Kind: "void"}}, {Value: 49, Name: "EmissiveDesaturation", Shape: Shape{Kind: "void"}}, {Value: 50, Name: "SkyIntensity", Shape: Shape{Kind: "void"}}, {Value: 51, Name: "OrbitalOffsetDegrees", Shape: Shape{Kind: "void"}}}}, p.IdentifierForParameter); err != nil {
		return err
	}
	if err := w.Write("GraphicsOverrideParameterPacket.Reset Parameter", Shape{Kind: "primitive", PrimitiveCode: "bool"}, p.ResetParameter); err != nil {
		return err
	}
	return nil
}

func DecodeGraphicsOverrideParameter(r Decoder) (GraphicsOverrideParameter, error) {
	var p GraphicsOverrideParameter
	{
		raw, err := r.Read("GraphicsOverrideParameterPacket.Parameter Keyframe Values", Shape{Kind: "map", Representation: "ordered_entries", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Value: &Shape{Kind: "struct", Semantic: "Vec3", TypeID: "Vec3", Fields: []ShapeField{{Ordinal: 0, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 1, Name: "Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 2, Name: "Z", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}}}, Key: &Shape{Kind: "primitive", PrimitiveCode: "f32le"}})
		if err != nil {
			return p, err
		}
		value, ok := raw.([]OrderedEntry[float32, Vec3])
		if !ok {
			return p, fmt.Errorf("field GraphicsOverrideParameterPacket.Parameter Keyframe Values has unexpected decoded type %T", raw)
		}
		p.ParameterKeyframeValues = value
	}
	{
		raw, err := r.Read("GraphicsOverrideParameterPacket.Float Value", Shape{Kind: "optional", Value: &Shape{Kind: "primitive", PrimitiveCode: "f32le"}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(*float32)
		if !ok {
			return p, fmt.Errorf("field GraphicsOverrideParameterPacket.Float Value has unexpected decoded type %T", raw)
		}
		p.FloatValue = value
	}
	{
		raw, err := r.Read("GraphicsOverrideParameterPacket.Vec3 Value", Shape{Kind: "optional", Value: &Shape{Kind: "struct", Semantic: "Vec3", TypeID: "Vec3", Fields: []ShapeField{{Ordinal: 0, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 1, Name: "Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 2, Name: "Z", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(*Vec3)
		if !ok {
			return p, fmt.Errorf("field GraphicsOverrideParameterPacket.Vec3 Value has unexpected decoded type %T", raw)
		}
		p.Vec3Value = value
	}
	{
		raw, err := r.Read("GraphicsOverrideParameterPacket.Biome Identifier", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(string)
		if !ok {
			return p, fmt.Errorf("field GraphicsOverrideParameterPacket.Biome Identifier has unexpected decoded type %T", raw)
		}
		p.BiomeIdentifier = value
	}
	{
		raw, err := r.Read("GraphicsOverrideParameterPacket.Player Identifier", Shape{Kind: "optional", Value: &Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(*string)
		if !ok {
			return p, fmt.Errorf("field GraphicsOverrideParameterPacket.Player Identifier has unexpected decoded type %T", raw)
		}
		p.PlayerIdentifier = value
	}
	{
		raw, err := r.Read("GraphicsOverrideParameterPacket.Identifier for Parameter", Shape{Kind: "enum", Semantic: "GraphicsOverrideParameterType", TypeID: "enums/GraphicsOverrideParameterType", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "SkyZenithColor", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "SkyHorizonColor", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "HorizonBlendMin", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "HorizonBlendMax", Shape: Shape{Kind: "void"}}, {Value: 4, Name: "HorizonBlendStart", Shape: Shape{Kind: "void"}}, {Value: 5, Name: "HorizonBlendMieStart", Shape: Shape{Kind: "void"}}, {Value: 6, Name: "RayleighStrength", Shape: Shape{Kind: "void"}}, {Value: 7, Name: "SunMieStrength", Shape: Shape{Kind: "void"}}, {Value: 8, Name: "MoonMieStrength", Shape: Shape{Kind: "void"}}, {Value: 9, Name: "SunGlareShape", Shape: Shape{Kind: "void"}}, {Value: 10, Name: "Chlorophyll", Shape: Shape{Kind: "void"}}, {Value: 11, Name: "CDOM", Shape: Shape{Kind: "void"}}, {Value: 12, Name: "SuspendedSediment", Shape: Shape{Kind: "void"}}, {Value: 13, Name: "WavesDepth", Shape: Shape{Kind: "void"}}, {Value: 14, Name: "WavesFrequency", Shape: Shape{Kind: "void"}}, {Value: 15, Name: "WavesFrequencyScaling", Shape: Shape{Kind: "void"}}, {Value: 16, Name: "WavesSpeed", Shape: Shape{Kind: "void"}}, {Value: 17, Name: "WavesSpeedScaling", Shape: Shape{Kind: "void"}}, {Value: 18, Name: "WavesShape", Shape: Shape{Kind: "void"}}, {Value: 19, Name: "WavesOctaves", Shape: Shape{Kind: "void"}}, {Value: 20, Name: "WavesMix", Shape: Shape{Kind: "void"}}, {Value: 21, Name: "WavesPull", Shape: Shape{Kind: "void"}}, {Value: 22, Name: "WavesDirectionIncrement", Shape: Shape{Kind: "void"}}, {Value: 23, Name: "MidtonesContrast", Shape: Shape{Kind: "void"}}, {Value: 24, Name: "HighlightsContrast", Shape: Shape{Kind: "void"}}, {Value: 25, Name: "ShadowsContrast", Shape: Shape{Kind: "void"}}, {Value: 26, Name: "HighlightsGain", Shape: Shape{Kind: "void"}}, {Value: 27, Name: "HighlightsGamma", Shape: Shape{Kind: "void"}}, {Value: 28, Name: "HighlightsOffset", Shape: Shape{Kind: "void"}}, {Value: 29, Name: "HighlightsSaturation", Shape: Shape{Kind: "void"}}, {Value: 30, Name: "MidtonesGain", Shape: Shape{Kind: "void"}}, {Value: 31, Name: "MidtonesGamma", Shape: Shape{Kind: "void"}}, {Value: 32, Name: "MidtonesOffset", Shape: Shape{Kind: "void"}}, {Value: 33, Name: "MidtonesSaturation", Shape: Shape{Kind: "void"}}, {Value: 34, Name: "ShadowsGain", Shape: Shape{Kind: "void"}}, {Value: 35, Name: "ShadowsGamma", Shape: Shape{Kind: "void"}}, {Value: 36, Name: "ShadowsOffset", Shape: Shape{Kind: "void"}}, {Value: 37, Name: "ShadowsSaturation", Shape: Shape{Kind: "void"}}, {Value: 38, Name: "HighlightsMin", Shape: Shape{Kind: "void"}}, {Value: 39, Name: "ShadowsMax", Shape: Shape{Kind: "void"}}, {Value: 40, Name: "Temperature", Shape: Shape{Kind: "void"}}, {Value: 41, Name: "SunColor", Shape: Shape{Kind: "void"}}, {Value: 42, Name: "SunIlluminance", Shape: Shape{Kind: "void"}}, {Value: 43, Name: "MoonColor", Shape: Shape{Kind: "void"}}, {Value: 44, Name: "MoonIlluminance", Shape: Shape{Kind: "void"}}, {Value: 45, Name: "FlashColor", Shape: Shape{Kind: "void"}}, {Value: 46, Name: "FlashIlluminance", Shape: Shape{Kind: "void"}}, {Value: 47, Name: "AmbientColor", Shape: Shape{Kind: "void"}}, {Value: 48, Name: "AmbientIlluminance", Shape: Shape{Kind: "void"}}, {Value: 49, Name: "EmissiveDesaturation", Shape: Shape{Kind: "void"}}, {Value: 50, Name: "SkyIntensity", Shape: Shape{Kind: "void"}}, {Value: 51, Name: "OrbitalOffsetDegrees", Shape: Shape{Kind: "void"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(GraphicsOverrideParameterType)
		if !ok {
			return p, fmt.Errorf("field GraphicsOverrideParameterPacket.Identifier for Parameter has unexpected decoded type %T", raw)
		}
		p.IdentifierForParameter = value
	}
	{
		raw, err := r.Read("GraphicsOverrideParameterPacket.Reset Parameter", Shape{Kind: "primitive", PrimitiveCode: "bool"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(bool)
		if !ok {
			return p, fmt.Errorf("field GraphicsOverrideParameterPacket.Reset Parameter has unexpected decoded type %T", raw)
		}
		p.ResetParameter = value
	}
	return p, nil
}
