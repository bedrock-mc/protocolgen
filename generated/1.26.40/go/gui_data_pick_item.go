// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type GuiDataPickItem struct {
	ItemName       string
	ItemEffectName string
	Slot           int32
}

func (p *GuiDataPickItem) Encode(w Encoder) error {
	if err := w.Write("GuiDataPickItemPacket.Item Name", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, p.ItemName); err != nil {
		return err
	}
	if err := w.Write("GuiDataPickItemPacket.Item Effect Name", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, p.ItemEffectName); err != nil {
		return err
	}
	if err := w.Write("GuiDataPickItemPacket.Slot", Shape{Kind: "primitive", PrimitiveCode: "i32le"}, p.Slot); err != nil {
		return err
	}
	return nil
}

func DecodeGuiDataPickItem(r Decoder) (GuiDataPickItem, error) {
	var p GuiDataPickItem
	{
		raw, err := r.Read("GuiDataPickItemPacket.Item Name", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(string)
		if !ok {
			return p, fmt.Errorf("field GuiDataPickItemPacket.Item Name has unexpected decoded type %T", raw)
		}
		p.ItemName = value
	}
	{
		raw, err := r.Read("GuiDataPickItemPacket.Item Effect Name", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(string)
		if !ok {
			return p, fmt.Errorf("field GuiDataPickItemPacket.Item Effect Name has unexpected decoded type %T", raw)
		}
		p.ItemEffectName = value
	}
	{
		raw, err := r.Read("GuiDataPickItemPacket.Slot", Shape{Kind: "primitive", PrimitiveCode: "i32le"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(int32)
		if !ok {
			return p, fmt.Errorf("field GuiDataPickItemPacket.Slot has unexpected decoded type %T", raw)
		}
		p.Slot = value
	}
	return p, nil
}
