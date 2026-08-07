// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type ResourcePacksInfo struct {
	ResourcePackRequired       bool
	HasAddonPacks              bool
	HasScripts                 bool
	ForceDisableVibrantVisuals bool
	WorldTemplateIdAndVersion  PackIdVersion
	ResourcePacks              []PackInfoData
}

func (p *ResourcePacksInfo) Encode(w Encoder) error {
	if err := w.Write("ResourcePacksInfoPacket.Resource Pack Required", Shape{Kind: "primitive", PrimitiveCode: "bool"}, p.ResourcePackRequired); err != nil {
		return err
	}
	if err := w.Write("ResourcePacksInfoPacket.Has Addon Packs", Shape{Kind: "primitive", PrimitiveCode: "bool"}, p.HasAddonPacks); err != nil {
		return err
	}
	if err := w.Write("ResourcePacksInfoPacket.Has Scripts", Shape{Kind: "primitive", PrimitiveCode: "bool"}, p.HasScripts); err != nil {
		return err
	}
	if err := w.Write("ResourcePacksInfoPacket.Force Disable Vibrant Visuals", Shape{Kind: "primitive", PrimitiveCode: "bool"}, p.ForceDisableVibrantVisuals); err != nil {
		return err
	}
	if err := w.Write("ResourcePacksInfoPacket.World Template Id And Version", Shape{Kind: "struct", Semantic: "PackIdVersion", TypeID: "PackIdVersion.json#", Fields: []ShapeField{{Ordinal: 0, Name: "Pack UUID", Shape: Shape{Kind: "primitive", Semantic: "mce::UUID", TypeID: "mce__UUID.json#", PrimitiveCode: "uuid"}}, {Ordinal: 1, Name: "Pack Version", Shape: Shape{Kind: "struct", Semantic: "SemVersion", TypeID: "SemVersion.json#", Fields: []ShapeField{{Ordinal: 0, Name: "Version", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}}}}}, p.WorldTemplateIdAndVersion); err != nil {
		return err
	}
	if err := w.Write("ResourcePacksInfoPacket.Resource Packs", Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "struct", Semantic: "PackInfoData", TypeID: "PackInfoData", Fields: []ShapeField{{Ordinal: 0, Name: "Pack Id Version", Shape: Shape{Kind: "struct", Semantic: "PackIdVersion", TypeID: "PackIdVersion", Fields: []ShapeField{{Ordinal: 0, Name: "Pack UUID", Shape: Shape{Kind: "primitive", PrimitiveCode: "uuid"}}, {Ordinal: 1, Name: "Pack Version", Shape: Shape{Kind: "struct", Semantic: "SemVersion", TypeID: "SemVersion", Fields: []ShapeField{{Ordinal: 0, Name: "Version", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}}}}}}, {Ordinal: 1, Name: "Pack Size", Shape: Shape{Kind: "primitive", PrimitiveCode: "u64le"}}, {Ordinal: 2, Name: "Content Key", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 3, Name: "Subpack Name", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 4, Name: "Content Identity", Shape: Shape{Kind: "struct", Semantic: "ContentIdentity", TypeID: "ContentIdentity", Fields: []ShapeField{{Ordinal: 0, Name: "Identity", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}}}, {Ordinal: 5, Name: "Has Scripts", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}, {Ordinal: 6, Name: "Is Addon Pack", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}, {Ordinal: 7, Name: "Is Ray Tracing Capable", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}, {Ordinal: 8, Name: "CDN URL", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}}}, p.ResourcePacks); err != nil {
		return err
	}
	return nil
}

func DecodeResourcePacksInfo(r Decoder) (ResourcePacksInfo, error) {
	var p ResourcePacksInfo
	{
		raw, err := r.Read("ResourcePacksInfoPacket.Resource Pack Required", Shape{Kind: "primitive", PrimitiveCode: "bool"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(bool)
		if !ok {
			return p, fmt.Errorf("field ResourcePacksInfoPacket.Resource Pack Required has unexpected decoded type %T", raw)
		}
		p.ResourcePackRequired = value
	}
	{
		raw, err := r.Read("ResourcePacksInfoPacket.Has Addon Packs", Shape{Kind: "primitive", PrimitiveCode: "bool"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(bool)
		if !ok {
			return p, fmt.Errorf("field ResourcePacksInfoPacket.Has Addon Packs has unexpected decoded type %T", raw)
		}
		p.HasAddonPacks = value
	}
	{
		raw, err := r.Read("ResourcePacksInfoPacket.Has Scripts", Shape{Kind: "primitive", PrimitiveCode: "bool"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(bool)
		if !ok {
			return p, fmt.Errorf("field ResourcePacksInfoPacket.Has Scripts has unexpected decoded type %T", raw)
		}
		p.HasScripts = value
	}
	{
		raw, err := r.Read("ResourcePacksInfoPacket.Force Disable Vibrant Visuals", Shape{Kind: "primitive", PrimitiveCode: "bool"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(bool)
		if !ok {
			return p, fmt.Errorf("field ResourcePacksInfoPacket.Force Disable Vibrant Visuals has unexpected decoded type %T", raw)
		}
		p.ForceDisableVibrantVisuals = value
	}
	{
		raw, err := r.Read("ResourcePacksInfoPacket.World Template Id And Version", Shape{Kind: "struct", Semantic: "PackIdVersion", TypeID: "PackIdVersion.json#", Fields: []ShapeField{{Ordinal: 0, Name: "Pack UUID", Shape: Shape{Kind: "primitive", Semantic: "mce::UUID", TypeID: "mce__UUID.json#", PrimitiveCode: "uuid"}}, {Ordinal: 1, Name: "Pack Version", Shape: Shape{Kind: "struct", Semantic: "SemVersion", TypeID: "SemVersion.json#", Fields: []ShapeField{{Ordinal: 0, Name: "Version", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(PackIdVersion)
		if !ok {
			return p, fmt.Errorf("field ResourcePacksInfoPacket.World Template Id And Version has unexpected decoded type %T", raw)
		}
		p.WorldTemplateIdAndVersion = value
	}
	{
		raw, err := r.Read("ResourcePacksInfoPacket.Resource Packs", Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "struct", Semantic: "PackInfoData", TypeID: "PackInfoData", Fields: []ShapeField{{Ordinal: 0, Name: "Pack Id Version", Shape: Shape{Kind: "struct", Semantic: "PackIdVersion", TypeID: "PackIdVersion", Fields: []ShapeField{{Ordinal: 0, Name: "Pack UUID", Shape: Shape{Kind: "primitive", PrimitiveCode: "uuid"}}, {Ordinal: 1, Name: "Pack Version", Shape: Shape{Kind: "struct", Semantic: "SemVersion", TypeID: "SemVersion", Fields: []ShapeField{{Ordinal: 0, Name: "Version", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}}}}}}, {Ordinal: 1, Name: "Pack Size", Shape: Shape{Kind: "primitive", PrimitiveCode: "u64le"}}, {Ordinal: 2, Name: "Content Key", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 3, Name: "Subpack Name", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 4, Name: "Content Identity", Shape: Shape{Kind: "struct", Semantic: "ContentIdentity", TypeID: "ContentIdentity", Fields: []ShapeField{{Ordinal: 0, Name: "Identity", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}}}, {Ordinal: 5, Name: "Has Scripts", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}, {Ordinal: 6, Name: "Is Addon Pack", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}, {Ordinal: 7, Name: "Is Ray Tracing Capable", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}, {Ordinal: 8, Name: "CDN URL", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.([]PackInfoData)
		if !ok {
			return p, fmt.Errorf("field ResourcePacksInfoPacket.Resource Packs has unexpected decoded type %T", raw)
		}
		p.ResourcePacks = value
	}
	return p, nil
}
