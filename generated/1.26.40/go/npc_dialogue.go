// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type NpcDialogue struct {
	NpcIdRawId            uint64
	NpcDialogueActionType NpcDialogueNpcDialogueActionType
	Dialogue              string
	SceneName             string
	NpcName               string
	ActionJSON            string
}

func (p *NpcDialogue) Encode(w Encoder) error {
	if err := w.Write("NpcDialoguePacket.Npc Id Raw Id", Shape{Kind: "primitive", PrimitiveCode: "u64le"}, p.NpcIdRawId); err != nil {
		return err
	}
	if err := w.Write("NpcDialoguePacket.Npc Dialogue Action Type", Shape{Kind: "enum", Semantic: "NpcDialoguePacketPayload::NpcDialogueActionType", TypeID: "enums/NpcDialoguePacketPayload::NpcDialogueActionType", PrimitiveCode: "zigzag_i32", Variants: []ShapeVariant{{Value: 0, Name: "Open", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "Close", Shape: Shape{Kind: "void"}}}}, p.NpcDialogueActionType); err != nil {
		return err
	}
	if err := w.Write("NpcDialoguePacket.Dialogue", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, p.Dialogue); err != nil {
		return err
	}
	if err := w.Write("NpcDialoguePacket.Scene Name", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, p.SceneName); err != nil {
		return err
	}
	if err := w.Write("NpcDialoguePacket.Npc Name", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, p.NpcName); err != nil {
		return err
	}
	if err := w.Write("NpcDialoguePacket.Action JSON", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, p.ActionJSON); err != nil {
		return err
	}
	return nil
}

func DecodeNpcDialogue(r Decoder) (NpcDialogue, error) {
	var p NpcDialogue
	{
		raw, err := r.Read("NpcDialoguePacket.Npc Id Raw Id", Shape{Kind: "primitive", PrimitiveCode: "u64le"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(uint64)
		if !ok {
			return p, fmt.Errorf("field NpcDialoguePacket.Npc Id Raw Id has unexpected decoded type %T", raw)
		}
		p.NpcIdRawId = value
	}
	{
		raw, err := r.Read("NpcDialoguePacket.Npc Dialogue Action Type", Shape{Kind: "enum", Semantic: "NpcDialoguePacketPayload::NpcDialogueActionType", TypeID: "enums/NpcDialoguePacketPayload::NpcDialogueActionType", PrimitiveCode: "zigzag_i32", Variants: []ShapeVariant{{Value: 0, Name: "Open", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "Close", Shape: Shape{Kind: "void"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(NpcDialogueNpcDialogueActionType)
		if !ok {
			return p, fmt.Errorf("field NpcDialoguePacket.Npc Dialogue Action Type has unexpected decoded type %T", raw)
		}
		p.NpcDialogueActionType = value
	}
	{
		raw, err := r.Read("NpcDialoguePacket.Dialogue", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(string)
		if !ok {
			return p, fmt.Errorf("field NpcDialoguePacket.Dialogue has unexpected decoded type %T", raw)
		}
		p.Dialogue = value
	}
	{
		raw, err := r.Read("NpcDialoguePacket.Scene Name", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(string)
		if !ok {
			return p, fmt.Errorf("field NpcDialoguePacket.Scene Name has unexpected decoded type %T", raw)
		}
		p.SceneName = value
	}
	{
		raw, err := r.Read("NpcDialoguePacket.Npc Name", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(string)
		if !ok {
			return p, fmt.Errorf("field NpcDialoguePacket.Npc Name has unexpected decoded type %T", raw)
		}
		p.NpcName = value
	}
	{
		raw, err := r.Read("NpcDialoguePacket.Action JSON", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(string)
		if !ok {
			return p, fmt.Errorf("field NpcDialoguePacket.Action JSON has unexpected decoded type %T", raw)
		}
		p.ActionJSON = value
	}
	return p, nil
}
