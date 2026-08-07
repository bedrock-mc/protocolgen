// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type ClientboundTextureShift struct {
	ActionID             ClientboundTextureShiftAction
	CollectionName       string
	FromStep             string
	ToStep               string
	AllSteps             []string
	CurrentLengthInTicks uint64
	TotalLengthInTicks   uint64
	Enabled              bool
}

func (p *ClientboundTextureShift) Encode(w Encoder) error {
	if err := w.Write("ClientboundTextureShiftPacket.Action ID", Shape{Kind: "enum", Semantic: "ClientboundTextureShiftPacketPayload::Action", TypeID: "enums/ClientboundTextureShiftPacketPayload::Action", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "Invalid", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "Initialize", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "Start", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "SetEnabled", Shape: Shape{Kind: "void"}}, {Value: 4, Name: "Sync", Shape: Shape{Kind: "void"}}}}, p.ActionID); err != nil {
		return err
	}
	if err := w.Write("ClientboundTextureShiftPacket.Collection Name", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, p.CollectionName); err != nil {
		return err
	}
	if err := w.Write("ClientboundTextureShiftPacket.From Step", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, p.FromStep); err != nil {
		return err
	}
	if err := w.Write("ClientboundTextureShiftPacket.To Step", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, p.ToStep); err != nil {
		return err
	}
	if err := w.Write("ClientboundTextureShiftPacket.All Steps", Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, p.AllSteps); err != nil {
		return err
	}
	if err := w.Write("ClientboundTextureShiftPacket.Current Length In Ticks", Shape{Kind: "primitive", PrimitiveCode: "var_u64"}, p.CurrentLengthInTicks); err != nil {
		return err
	}
	if err := w.Write("ClientboundTextureShiftPacket.Total Length In Ticks", Shape{Kind: "primitive", PrimitiveCode: "var_u64"}, p.TotalLengthInTicks); err != nil {
		return err
	}
	if err := w.Write("ClientboundTextureShiftPacket.Enabled", Shape{Kind: "primitive", PrimitiveCode: "bool"}, p.Enabled); err != nil {
		return err
	}
	return nil
}

func DecodeClientboundTextureShift(r Decoder) (ClientboundTextureShift, error) {
	var p ClientboundTextureShift
	{
		raw, err := r.Read("ClientboundTextureShiftPacket.Action ID", Shape{Kind: "enum", Semantic: "ClientboundTextureShiftPacketPayload::Action", TypeID: "enums/ClientboundTextureShiftPacketPayload::Action", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "Invalid", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "Initialize", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "Start", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "SetEnabled", Shape: Shape{Kind: "void"}}, {Value: 4, Name: "Sync", Shape: Shape{Kind: "void"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(ClientboundTextureShiftAction)
		if !ok {
			return p, fmt.Errorf("field ClientboundTextureShiftPacket.Action ID has unexpected decoded type %T", raw)
		}
		p.ActionID = value
	}
	{
		raw, err := r.Read("ClientboundTextureShiftPacket.Collection Name", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(string)
		if !ok {
			return p, fmt.Errorf("field ClientboundTextureShiftPacket.Collection Name has unexpected decoded type %T", raw)
		}
		p.CollectionName = value
	}
	{
		raw, err := r.Read("ClientboundTextureShiftPacket.From Step", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(string)
		if !ok {
			return p, fmt.Errorf("field ClientboundTextureShiftPacket.From Step has unexpected decoded type %T", raw)
		}
		p.FromStep = value
	}
	{
		raw, err := r.Read("ClientboundTextureShiftPacket.To Step", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(string)
		if !ok {
			return p, fmt.Errorf("field ClientboundTextureShiftPacket.To Step has unexpected decoded type %T", raw)
		}
		p.ToStep = value
	}
	{
		raw, err := r.Read("ClientboundTextureShiftPacket.All Steps", Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.([]string)
		if !ok {
			return p, fmt.Errorf("field ClientboundTextureShiftPacket.All Steps has unexpected decoded type %T", raw)
		}
		p.AllSteps = value
	}
	{
		raw, err := r.Read("ClientboundTextureShiftPacket.Current Length In Ticks", Shape{Kind: "primitive", PrimitiveCode: "var_u64"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(uint64)
		if !ok {
			return p, fmt.Errorf("field ClientboundTextureShiftPacket.Current Length In Ticks has unexpected decoded type %T", raw)
		}
		p.CurrentLengthInTicks = value
	}
	{
		raw, err := r.Read("ClientboundTextureShiftPacket.Total Length In Ticks", Shape{Kind: "primitive", PrimitiveCode: "var_u64"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(uint64)
		if !ok {
			return p, fmt.Errorf("field ClientboundTextureShiftPacket.Total Length In Ticks has unexpected decoded type %T", raw)
		}
		p.TotalLengthInTicks = value
	}
	{
		raw, err := r.Read("ClientboundTextureShiftPacket.Enabled", Shape{Kind: "primitive", PrimitiveCode: "bool"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(bool)
		if !ok {
			return p, fmt.Errorf("field ClientboundTextureShiftPacket.Enabled has unexpected decoded type %T", raw)
		}
		p.Enabled = value
	}
	return p, nil
}
