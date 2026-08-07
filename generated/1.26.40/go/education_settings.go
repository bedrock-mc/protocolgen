// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type EducationSettings struct {
	EducationLevelSettings EducationLevelSettings
}

func (p *EducationSettings) Encode(w Encoder) error {
	if err := w.Write("EducationSettingsPacket.Education Level Settings", Shape{Kind: "struct", Semantic: "EducationLevelSettings", TypeID: "EducationLevelSettings", Fields: []ShapeField{{Ordinal: 0, Name: "Code Builder Default URI", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 1, Name: "Code Builder Title", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 2, Name: "Can resize Code Builder", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}, {Ordinal: 3, Name: "Disable legacy title bar", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}, {Ordinal: 4, Name: "Post Process Filter", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 5, Name: "Screenshot Border Resource Path", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 6, Name: "Agent Capabilities", Shape: Shape{Kind: "optional", Value: &Shape{Kind: "struct", Semantic: "AgentCapabilities", TypeID: "AgentCapabilities", Fields: []ShapeField{{Ordinal: 0, Name: "Can Modify Blocks", Shape: Shape{Kind: "optional", Value: &Shape{Kind: "primitive", PrimitiveCode: "bool"}}}}}}}, {Ordinal: 7, Name: "Local Settings", Shape: Shape{Kind: "struct", Semantic: "EducationLocalLevelSettings", TypeID: "EducationLocalLevelSettings", Fields: []ShapeField{{Ordinal: 0, Name: "Code Builder Override Uri", Shape: Shape{Kind: "optional", Value: &Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}}}}, {Ordinal: 8, Name: "(Deprecated) Always False", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}, {Ordinal: 9, Name: "External Link Settings", Shape: Shape{Kind: "optional", Value: &Shape{Kind: "struct", Semantic: "ExternalLinkSettings", TypeID: "ExternalLinkSettings", Fields: []ShapeField{{Ordinal: 0, Name: "URL", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 1, Name: "Display Name", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}}}}}}, p.EducationLevelSettings); err != nil {
		return err
	}
	return nil
}

func DecodeEducationSettings(r Decoder) (EducationSettings, error) {
	var p EducationSettings
	{
		raw, err := r.Read("EducationSettingsPacket.Education Level Settings", Shape{Kind: "struct", Semantic: "EducationLevelSettings", TypeID: "EducationLevelSettings", Fields: []ShapeField{{Ordinal: 0, Name: "Code Builder Default URI", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 1, Name: "Code Builder Title", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 2, Name: "Can resize Code Builder", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}, {Ordinal: 3, Name: "Disable legacy title bar", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}, {Ordinal: 4, Name: "Post Process Filter", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 5, Name: "Screenshot Border Resource Path", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 6, Name: "Agent Capabilities", Shape: Shape{Kind: "optional", Value: &Shape{Kind: "struct", Semantic: "AgentCapabilities", TypeID: "AgentCapabilities", Fields: []ShapeField{{Ordinal: 0, Name: "Can Modify Blocks", Shape: Shape{Kind: "optional", Value: &Shape{Kind: "primitive", PrimitiveCode: "bool"}}}}}}}, {Ordinal: 7, Name: "Local Settings", Shape: Shape{Kind: "struct", Semantic: "EducationLocalLevelSettings", TypeID: "EducationLocalLevelSettings", Fields: []ShapeField{{Ordinal: 0, Name: "Code Builder Override Uri", Shape: Shape{Kind: "optional", Value: &Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}}}}, {Ordinal: 8, Name: "(Deprecated) Always False", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}, {Ordinal: 9, Name: "External Link Settings", Shape: Shape{Kind: "optional", Value: &Shape{Kind: "struct", Semantic: "ExternalLinkSettings", TypeID: "ExternalLinkSettings", Fields: []ShapeField{{Ordinal: 0, Name: "URL", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 1, Name: "Display Name", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}}}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(EducationLevelSettings)
		if !ok {
			return p, fmt.Errorf("field EducationSettingsPacket.Education Level Settings has unexpected decoded type %T", raw)
		}
		p.EducationLevelSettings = value
	}
	return p, nil
}
