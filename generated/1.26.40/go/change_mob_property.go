// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type ChangeMobProperty struct {
	ActorId              ActorUniqueID
	PropertyName         string
	BoolComponentValue   bool
	StringComponentValue string
	IntComponentValue    int32
	FloatComponentValue  float32
}

func (p *ChangeMobProperty) Encode(w Encoder) error {
	if err := w.Write("ChangeMobPropertyPacket.Actor Id", Shape{Kind: "struct", Semantic: "ActorUniqueID", TypeID: "ActorUniqueID", Fields: []ShapeField{{Ordinal: 0, Name: "Actor Unique ID", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i64"}}}}, p.ActorId); err != nil {
		return err
	}
	if err := w.Write("ChangeMobPropertyPacket.Property Name", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, p.PropertyName); err != nil {
		return err
	}
	if err := w.Write("ChangeMobPropertyPacket.BoolComponent Value", Shape{Kind: "primitive", PrimitiveCode: "bool"}, p.BoolComponentValue); err != nil {
		return err
	}
	if err := w.Write("ChangeMobPropertyPacket.StringComponent Value", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, p.StringComponentValue); err != nil {
		return err
	}
	if err := w.Write("ChangeMobPropertyPacket.IntComponent Value", Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}, p.IntComponentValue); err != nil {
		return err
	}
	if err := w.Write("ChangeMobPropertyPacket.FloatComponent Value", Shape{Kind: "primitive", PrimitiveCode: "f32le"}, p.FloatComponentValue); err != nil {
		return err
	}
	return nil
}

func DecodeChangeMobProperty(r Decoder) (ChangeMobProperty, error) {
	var p ChangeMobProperty
	{
		raw, err := r.Read("ChangeMobPropertyPacket.Actor Id", Shape{Kind: "struct", Semantic: "ActorUniqueID", TypeID: "ActorUniqueID", Fields: []ShapeField{{Ordinal: 0, Name: "Actor Unique ID", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i64"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(ActorUniqueID)
		if !ok {
			return p, fmt.Errorf("field ChangeMobPropertyPacket.Actor Id has unexpected decoded type %T", raw)
		}
		p.ActorId = value
	}
	{
		raw, err := r.Read("ChangeMobPropertyPacket.Property Name", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(string)
		if !ok {
			return p, fmt.Errorf("field ChangeMobPropertyPacket.Property Name has unexpected decoded type %T", raw)
		}
		p.PropertyName = value
	}
	{
		raw, err := r.Read("ChangeMobPropertyPacket.BoolComponent Value", Shape{Kind: "primitive", PrimitiveCode: "bool"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(bool)
		if !ok {
			return p, fmt.Errorf("field ChangeMobPropertyPacket.BoolComponent Value has unexpected decoded type %T", raw)
		}
		p.BoolComponentValue = value
	}
	{
		raw, err := r.Read("ChangeMobPropertyPacket.StringComponent Value", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(string)
		if !ok {
			return p, fmt.Errorf("field ChangeMobPropertyPacket.StringComponent Value has unexpected decoded type %T", raw)
		}
		p.StringComponentValue = value
	}
	{
		raw, err := r.Read("ChangeMobPropertyPacket.IntComponent Value", Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(int32)
		if !ok {
			return p, fmt.Errorf("field ChangeMobPropertyPacket.IntComponent Value has unexpected decoded type %T", raw)
		}
		p.IntComponentValue = value
	}
	{
		raw, err := r.Read("ChangeMobPropertyPacket.FloatComponent Value", Shape{Kind: "primitive", PrimitiveCode: "f32le"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(float32)
		if !ok {
			return p, fmt.Errorf("field ChangeMobPropertyPacket.FloatComponent Value has unexpected decoded type %T", raw)
		}
		p.FloatComponentValue = value
	}
	return p, nil
}
