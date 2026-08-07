// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type AgentActionEvent struct {
	RequestId string
	Action    AgentActionType
	Response  string
}

func (p *AgentActionEvent) Encode(w Encoder) error {
	if err := w.Write("AgentActionEventPacket.Request Id", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, p.RequestId); err != nil {
		return err
	}
	if err := w.Write("AgentActionEventPacket.Action", Shape{Kind: "enum", Semantic: "AgentActionType", TypeID: "enums/AgentActionType", PrimitiveCode: "i32le", Variants: []ShapeVariant{{Value: 1, Name: "Attack", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "Collect", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "Destroy", Shape: Shape{Kind: "void"}}, {Value: 4, Name: "DetectRedstone", Shape: Shape{Kind: "void"}}, {Value: 5, Name: "DetectObstacle", Shape: Shape{Kind: "void"}}, {Value: 6, Name: "Drop", Shape: Shape{Kind: "void"}}, {Value: 7, Name: "DropAll", Shape: Shape{Kind: "void"}}, {Value: 8, Name: "Inspect", Shape: Shape{Kind: "void"}}, {Value: 9, Name: "InspectData", Shape: Shape{Kind: "void"}}, {Value: 10, Name: "InspectItemCount", Shape: Shape{Kind: "void"}}, {Value: 11, Name: "InspectItemDetail", Shape: Shape{Kind: "void"}}, {Value: 12, Name: "InspectItemSpace", Shape: Shape{Kind: "void"}}, {Value: 13, Name: "Interact", Shape: Shape{Kind: "void"}}, {Value: 14, Name: "Move", Shape: Shape{Kind: "void"}}, {Value: 15, Name: "PlaceBlock", Shape: Shape{Kind: "void"}}, {Value: 16, Name: "Till", Shape: Shape{Kind: "void"}}, {Value: 17, Name: "TransferItemTo", Shape: Shape{Kind: "void"}}, {Value: 18, Name: "Turn", Shape: Shape{Kind: "void"}}}}, p.Action); err != nil {
		return err
	}
	if err := w.Write("AgentActionEventPacket.Response", Shape{Kind: "string", Semantic: "Json::Value", TypeID: "Json__Value.json#", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, p.Response); err != nil {
		return err
	}
	return nil
}

func DecodeAgentActionEvent(r Decoder) (AgentActionEvent, error) {
	var p AgentActionEvent
	{
		raw, err := r.Read("AgentActionEventPacket.Request Id", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(string)
		if !ok {
			return p, fmt.Errorf("field AgentActionEventPacket.Request Id has unexpected decoded type %T", raw)
		}
		p.RequestId = value
	}
	{
		raw, err := r.Read("AgentActionEventPacket.Action", Shape{Kind: "enum", Semantic: "AgentActionType", TypeID: "enums/AgentActionType", PrimitiveCode: "i32le", Variants: []ShapeVariant{{Value: 1, Name: "Attack", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "Collect", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "Destroy", Shape: Shape{Kind: "void"}}, {Value: 4, Name: "DetectRedstone", Shape: Shape{Kind: "void"}}, {Value: 5, Name: "DetectObstacle", Shape: Shape{Kind: "void"}}, {Value: 6, Name: "Drop", Shape: Shape{Kind: "void"}}, {Value: 7, Name: "DropAll", Shape: Shape{Kind: "void"}}, {Value: 8, Name: "Inspect", Shape: Shape{Kind: "void"}}, {Value: 9, Name: "InspectData", Shape: Shape{Kind: "void"}}, {Value: 10, Name: "InspectItemCount", Shape: Shape{Kind: "void"}}, {Value: 11, Name: "InspectItemDetail", Shape: Shape{Kind: "void"}}, {Value: 12, Name: "InspectItemSpace", Shape: Shape{Kind: "void"}}, {Value: 13, Name: "Interact", Shape: Shape{Kind: "void"}}, {Value: 14, Name: "Move", Shape: Shape{Kind: "void"}}, {Value: 15, Name: "PlaceBlock", Shape: Shape{Kind: "void"}}, {Value: 16, Name: "Till", Shape: Shape{Kind: "void"}}, {Value: 17, Name: "TransferItemTo", Shape: Shape{Kind: "void"}}, {Value: 18, Name: "Turn", Shape: Shape{Kind: "void"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(AgentActionType)
		if !ok {
			return p, fmt.Errorf("field AgentActionEventPacket.Action has unexpected decoded type %T", raw)
		}
		p.Action = value
	}
	{
		raw, err := r.Read("AgentActionEventPacket.Response", Shape{Kind: "string", Semantic: "Json::Value", TypeID: "Json__Value.json#", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(string)
		if !ok {
			return p, fmt.Errorf("field AgentActionEventPacket.Response has unexpected decoded type %T", raw)
		}
		p.Response = value
	}
	return p, nil
}
