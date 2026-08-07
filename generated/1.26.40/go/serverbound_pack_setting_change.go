// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type ServerboundPackSettingChange struct {
	PackId           [16]byte
	PackSettingName  string
	PackSettingValue ServerboundPackSettingChangePackSettingValue
}

func (p *ServerboundPackSettingChange) Encode(w Encoder) error {
	if err := w.Write("ServerboundPackSettingChangePacket.PackId", Shape{Kind: "primitive", Semantic: "mce::UUID", TypeID: "mce__UUID.json#", PrimitiveCode: "uuid"}, p.PackId); err != nil {
		return err
	}
	if err := w.Write("ServerboundPackSettingChangePacket.PackSettingName", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, p.PackSettingName); err != nil {
		return err
	}
	if err := w.Write("ServerboundPackSettingChangePacket.PackSettingValue", Shape{Kind: "union", Control: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Variants: []ShapeVariant{{Value: 0, Name: "float", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Value: 1, Name: "bool", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}, {Value: 2, Name: "string", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}}, p.PackSettingValue); err != nil {
		return err
	}
	return nil
}

func DecodeServerboundPackSettingChange(r Decoder) (ServerboundPackSettingChange, error) {
	var p ServerboundPackSettingChange
	{
		raw, err := r.Read("ServerboundPackSettingChangePacket.PackId", Shape{Kind: "primitive", Semantic: "mce::UUID", TypeID: "mce__UUID.json#", PrimitiveCode: "uuid"})
		if err != nil {
			return p, err
		}
		value, ok := raw.([16]byte)
		if !ok {
			return p, fmt.Errorf("field ServerboundPackSettingChangePacket.PackId has unexpected decoded type %T", raw)
		}
		p.PackId = value
	}
	{
		raw, err := r.Read("ServerboundPackSettingChangePacket.PackSettingName", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(string)
		if !ok {
			return p, fmt.Errorf("field ServerboundPackSettingChangePacket.PackSettingName has unexpected decoded type %T", raw)
		}
		p.PackSettingName = value
	}
	{
		raw, err := r.Read("ServerboundPackSettingChangePacket.PackSettingValue", Shape{Kind: "union", Control: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Variants: []ShapeVariant{{Value: 0, Name: "float", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Value: 1, Name: "bool", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}, {Value: 2, Name: "string", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(ServerboundPackSettingChangePackSettingValue)
		if !ok {
			return p, fmt.Errorf("field ServerboundPackSettingChangePacket.PackSettingValue has unexpected decoded type %T", raw)
		}
		p.PackSettingValue = value
	}
	return p, nil
}
