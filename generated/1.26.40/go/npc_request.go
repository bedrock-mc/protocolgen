// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type NpcRequest struct {
	NPCRuntimeID ActorRuntimeID
	RequestType  NpcRequestRequestType
	Actions      string
	ActionIndex  uint8
	SceneName    string
}

func (p *NpcRequest) Encode(w Encoder) error {
	if err := w.Write("NpcRequestPacket.NPC Runtime ID", Shape{Kind: "struct", Semantic: "ActorRuntimeID", TypeID: "ActorRuntimeID", Fields: []ShapeField{{Ordinal: 0, Name: "Actor Runtime ID", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u64"}}}}, p.NPCRuntimeID); err != nil {
		return err
	}
	if err := w.Write("NpcRequestPacket.Request Type", Shape{Kind: "enum", Semantic: "NpcRequestPacketPayload::RequestType", TypeID: "enums/NpcRequestPacketPayload::RequestType", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "SetActions", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "ExecuteAction", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "ExecuteClosingCommands", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "SetName", Shape: Shape{Kind: "void"}}, {Value: 4, Name: "SetSkin", Shape: Shape{Kind: "void"}}, {Value: 5, Name: "SetInteractText", Shape: Shape{Kind: "void"}}, {Value: 6, Name: "ExecuteOpeningCommands", Shape: Shape{Kind: "void"}}}}, p.RequestType); err != nil {
		return err
	}
	if err := w.Write("NpcRequestPacket.Actions", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, p.Actions); err != nil {
		return err
	}
	if err := w.Write("NpcRequestPacket.Action Index", Shape{Kind: "primitive", PrimitiveCode: "u8"}, p.ActionIndex); err != nil {
		return err
	}
	if err := w.Write("NpcRequestPacket.Scene Name", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, p.SceneName); err != nil {
		return err
	}
	return nil
}

func DecodeNpcRequest(r Decoder) (NpcRequest, error) {
	var p NpcRequest
	{
		raw, err := r.Read("NpcRequestPacket.NPC Runtime ID", Shape{Kind: "struct", Semantic: "ActorRuntimeID", TypeID: "ActorRuntimeID", Fields: []ShapeField{{Ordinal: 0, Name: "Actor Runtime ID", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u64"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(ActorRuntimeID)
		if !ok {
			return p, fmt.Errorf("field NpcRequestPacket.NPC Runtime ID has unexpected decoded type %T", raw)
		}
		p.NPCRuntimeID = value
	}
	{
		raw, err := r.Read("NpcRequestPacket.Request Type", Shape{Kind: "enum", Semantic: "NpcRequestPacketPayload::RequestType", TypeID: "enums/NpcRequestPacketPayload::RequestType", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "SetActions", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "ExecuteAction", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "ExecuteClosingCommands", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "SetName", Shape: Shape{Kind: "void"}}, {Value: 4, Name: "SetSkin", Shape: Shape{Kind: "void"}}, {Value: 5, Name: "SetInteractText", Shape: Shape{Kind: "void"}}, {Value: 6, Name: "ExecuteOpeningCommands", Shape: Shape{Kind: "void"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(NpcRequestRequestType)
		if !ok {
			return p, fmt.Errorf("field NpcRequestPacket.Request Type has unexpected decoded type %T", raw)
		}
		p.RequestType = value
	}
	{
		raw, err := r.Read("NpcRequestPacket.Actions", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(string)
		if !ok {
			return p, fmt.Errorf("field NpcRequestPacket.Actions has unexpected decoded type %T", raw)
		}
		p.Actions = value
	}
	{
		raw, err := r.Read("NpcRequestPacket.Action Index", Shape{Kind: "primitive", PrimitiveCode: "u8"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(uint8)
		if !ok {
			return p, fmt.Errorf("field NpcRequestPacket.Action Index has unexpected decoded type %T", raw)
		}
		p.ActionIndex = value
	}
	{
		raw, err := r.Read("NpcRequestPacket.Scene Name", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(string)
		if !ok {
			return p, fmt.Errorf("field NpcRequestPacket.Scene Name has unexpected decoded type %T", raw)
		}
		p.SceneName = value
	}
	return p, nil
}
