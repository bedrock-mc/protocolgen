// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type Interact struct {
	Action          InteractAction
	TargetRuntimeID ActorRuntimeID
	Position        *Vec3
}

func (p *Interact) Encode(w Encoder) error {
	if err := w.Write("InteractPacket.Action", Shape{Kind: "enum", Semantic: "InteractPacketPayload::Action", TypeID: "enums/InteractPacketPayload::Action", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "Invalid", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "StopRiding", Shape: Shape{Kind: "void"}}, {Value: 4, Name: "InteractUpdate", Shape: Shape{Kind: "void"}}, {Value: 5, Name: "NpcOpen", Shape: Shape{Kind: "void"}}, {Value: 6, Name: "OpenInventory", Shape: Shape{Kind: "void"}}}}, p.Action); err != nil {
		return err
	}
	if err := w.Write("InteractPacket.Target Runtime ID", Shape{Kind: "struct", Semantic: "ActorRuntimeID", TypeID: "ActorRuntimeID", Fields: []ShapeField{{Ordinal: 0, Name: "Actor Runtime ID", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u64"}}}}, p.TargetRuntimeID); err != nil {
		return err
	}
	if err := w.Write("InteractPacket.Position", Shape{Kind: "optional", Value: &Shape{Kind: "struct", Semantic: "Vec3", TypeID: "Vec3", Fields: []ShapeField{{Ordinal: 0, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 1, Name: "Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 2, Name: "Z", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}}}}, p.Position); err != nil {
		return err
	}
	return nil
}

func DecodeInteract(r Decoder) (Interact, error) {
	var p Interact
	{
		raw, err := r.Read("InteractPacket.Action", Shape{Kind: "enum", Semantic: "InteractPacketPayload::Action", TypeID: "enums/InteractPacketPayload::Action", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "Invalid", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "StopRiding", Shape: Shape{Kind: "void"}}, {Value: 4, Name: "InteractUpdate", Shape: Shape{Kind: "void"}}, {Value: 5, Name: "NpcOpen", Shape: Shape{Kind: "void"}}, {Value: 6, Name: "OpenInventory", Shape: Shape{Kind: "void"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(InteractAction)
		if !ok {
			return p, fmt.Errorf("field InteractPacket.Action has unexpected decoded type %T", raw)
		}
		p.Action = value
	}
	{
		raw, err := r.Read("InteractPacket.Target Runtime ID", Shape{Kind: "struct", Semantic: "ActorRuntimeID", TypeID: "ActorRuntimeID", Fields: []ShapeField{{Ordinal: 0, Name: "Actor Runtime ID", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u64"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(ActorRuntimeID)
		if !ok {
			return p, fmt.Errorf("field InteractPacket.Target Runtime ID has unexpected decoded type %T", raw)
		}
		p.TargetRuntimeID = value
	}
	{
		raw, err := r.Read("InteractPacket.Position", Shape{Kind: "optional", Value: &Shape{Kind: "struct", Semantic: "Vec3", TypeID: "Vec3", Fields: []ShapeField{{Ordinal: 0, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 1, Name: "Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 2, Name: "Z", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(*Vec3)
		if !ok {
			return p, fmt.Errorf("field InteractPacket.Position has unexpected decoded type %T", raw)
		}
		p.Position = value
	}
	return p, nil
}
