// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type UpdateAdventureSettings struct {
	AdventureSettings AdventureSettings
}

func (p *UpdateAdventureSettings) Encode(w Encoder) error {
	if err := w.Write("UpdateAdventureSettingsPacket.Adventure Settings", Shape{Kind: "struct", Semantic: "AdventureSettings", TypeID: "AdventureSettings", Fields: []ShapeField{{Ordinal: 0, Name: "no PvM", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}, {Ordinal: 1, Name: "no MvP", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}, {Ordinal: 2, Name: "Immutable World", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}, {Ordinal: 3, Name: "Show Name Tags", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}, {Ordinal: 4, Name: "Auto Jump", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}}}, p.AdventureSettings); err != nil {
		return err
	}
	return nil
}

func DecodeUpdateAdventureSettings(r Decoder) (UpdateAdventureSettings, error) {
	var p UpdateAdventureSettings
	{
		raw, err := r.Read("UpdateAdventureSettingsPacket.Adventure Settings", Shape{Kind: "struct", Semantic: "AdventureSettings", TypeID: "AdventureSettings", Fields: []ShapeField{{Ordinal: 0, Name: "no PvM", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}, {Ordinal: 1, Name: "no MvP", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}, {Ordinal: 2, Name: "Immutable World", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}, {Ordinal: 3, Name: "Show Name Tags", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}, {Ordinal: 4, Name: "Auto Jump", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(AdventureSettings)
		if !ok {
			return p, fmt.Errorf("field UpdateAdventureSettingsPacket.Adventure Settings has unexpected decoded type %T", raw)
		}
		p.AdventureSettings = value
	}
	return p, nil
}
