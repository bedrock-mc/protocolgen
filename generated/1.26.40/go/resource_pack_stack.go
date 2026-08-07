// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type ResourcePackStack struct {
	TexturePackRequired bool
	TexturePackList     []PackInstanceId
	BaseGameVersion     string
	Experiments         Experiments
	IncludeEditorPacks  bool
}

func (p *ResourcePackStack) Encode(w Encoder) error {
	if err := w.Write("ResourcePackStackPacket.Texture Pack Required", Shape{Kind: "primitive", PrimitiveCode: "bool"}, p.TexturePackRequired); err != nil {
		return err
	}
	if err := w.Write("ResourcePackStackPacket.Texture Pack List", Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "struct", Semantic: "PackInstanceId", TypeID: "PackInstanceId", Fields: []ShapeField{{Ordinal: 0, Name: "Pack ID", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 1, Name: "Version", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 2, Name: "Sub Pack Name", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}}}, p.TexturePackList); err != nil {
		return err
	}
	if err := w.Write("ResourcePackStackPacket.Base Game Version", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, p.BaseGameVersion); err != nil {
		return err
	}
	if err := w.Write("ResourcePackStackPacket.Experiments", Shape{Kind: "struct", Semantic: "Experiments", TypeID: "Experiments", Fields: []ShapeField{{Ordinal: 0, Name: "Toggles", Shape: Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "u32le"}, Element: &Shape{Kind: "struct", Semantic: "cerealizer_ExperimentsAnon::ExperimentToggle", TypeID: "cerealizer_ExperimentsAnon::ExperimentToggle", Fields: []ShapeField{{Ordinal: 0, Name: "Name", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 1, Name: "Enabled", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}}}}}, {Ordinal: 1, Name: "ExperimentsEverToggled", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}}}, p.Experiments); err != nil {
		return err
	}
	if err := w.Write("ResourcePackStackPacket.Include Editor Packs", Shape{Kind: "primitive", PrimitiveCode: "bool"}, p.IncludeEditorPacks); err != nil {
		return err
	}
	return nil
}

func DecodeResourcePackStack(r Decoder) (ResourcePackStack, error) {
	var p ResourcePackStack
	{
		raw, err := r.Read("ResourcePackStackPacket.Texture Pack Required", Shape{Kind: "primitive", PrimitiveCode: "bool"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(bool)
		if !ok {
			return p, fmt.Errorf("field ResourcePackStackPacket.Texture Pack Required has unexpected decoded type %T", raw)
		}
		p.TexturePackRequired = value
	}
	{
		raw, err := r.Read("ResourcePackStackPacket.Texture Pack List", Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "struct", Semantic: "PackInstanceId", TypeID: "PackInstanceId", Fields: []ShapeField{{Ordinal: 0, Name: "Pack ID", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 1, Name: "Version", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 2, Name: "Sub Pack Name", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.([]PackInstanceId)
		if !ok {
			return p, fmt.Errorf("field ResourcePackStackPacket.Texture Pack List has unexpected decoded type %T", raw)
		}
		p.TexturePackList = value
	}
	{
		raw, err := r.Read("ResourcePackStackPacket.Base Game Version", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(string)
		if !ok {
			return p, fmt.Errorf("field ResourcePackStackPacket.Base Game Version has unexpected decoded type %T", raw)
		}
		p.BaseGameVersion = value
	}
	{
		raw, err := r.Read("ResourcePackStackPacket.Experiments", Shape{Kind: "struct", Semantic: "Experiments", TypeID: "Experiments", Fields: []ShapeField{{Ordinal: 0, Name: "Toggles", Shape: Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "u32le"}, Element: &Shape{Kind: "struct", Semantic: "cerealizer_ExperimentsAnon::ExperimentToggle", TypeID: "cerealizer_ExperimentsAnon::ExperimentToggle", Fields: []ShapeField{{Ordinal: 0, Name: "Name", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 1, Name: "Enabled", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}}}}}, {Ordinal: 1, Name: "ExperimentsEverToggled", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(Experiments)
		if !ok {
			return p, fmt.Errorf("field ResourcePackStackPacket.Experiments has unexpected decoded type %T", raw)
		}
		p.Experiments = value
	}
	{
		raw, err := r.Read("ResourcePackStackPacket.Include Editor Packs", Shape{Kind: "primitive", PrimitiveCode: "bool"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(bool)
		if !ok {
			return p, fmt.Errorf("field ResourcePackStackPacket.Include Editor Packs has unexpected decoded type %T", raw)
		}
		p.IncludeEditorPacks = value
	}
	return p, nil
}
