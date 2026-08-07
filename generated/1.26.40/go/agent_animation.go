// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type AgentAnimation struct {
	AgentAnimation AgentAnimationType
	RuntimeId      ActorRuntimeID
}

func (p *AgentAnimation) Encode(w Encoder) error {
	if err := w.Write("AgentAnimationPacket.Agent Animation", Shape{Kind: "enum", Semantic: "AgentAnimation", TypeID: "enums/AgentAnimation", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "ArmSwing", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "Shrug", Shape: Shape{Kind: "void"}}}}, p.AgentAnimation); err != nil {
		return err
	}
	if err := w.Write("AgentAnimationPacket.Runtime Id", Shape{Kind: "struct", Semantic: "ActorRuntimeID", TypeID: "ActorRuntimeID", Fields: []ShapeField{{Ordinal: 0, Name: "Actor Runtime ID", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u64"}}}}, p.RuntimeId); err != nil {
		return err
	}
	return nil
}

func DecodeAgentAnimation(r Decoder) (AgentAnimation, error) {
	var p AgentAnimation
	{
		raw, err := r.Read("AgentAnimationPacket.Agent Animation", Shape{Kind: "enum", Semantic: "AgentAnimation", TypeID: "enums/AgentAnimation", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "ArmSwing", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "Shrug", Shape: Shape{Kind: "void"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(AgentAnimationType)
		if !ok {
			return p, fmt.Errorf("field AgentAnimationPacket.Agent Animation has unexpected decoded type %T", raw)
		}
		p.AgentAnimation = value
	}
	{
		raw, err := r.Read("AgentAnimationPacket.Runtime Id", Shape{Kind: "struct", Semantic: "ActorRuntimeID", TypeID: "ActorRuntimeID", Fields: []ShapeField{{Ordinal: 0, Name: "Actor Runtime ID", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u64"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(ActorRuntimeID)
		if !ok {
			return p, fmt.Errorf("field AgentAnimationPacket.Runtime Id has unexpected decoded type %T", raw)
		}
		p.RuntimeId = value
	}
	return p, nil
}
