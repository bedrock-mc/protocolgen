// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type SimpleEvent struct {
	Type SimpleEventSubtype
}

func (p *SimpleEvent) Encode(w Encoder) error {
	if err := w.Write("SimpleEventPacket.Type", Shape{Kind: "enum", Semantic: "SimpleEventPacketPayload::Subtype", TypeID: "enums/SimpleEventPacketPayload::Subtype", PrimitiveCode: "u16le", Variants: []ShapeVariant{{Value: 0, Name: "UninitializedSubtype", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "EnableCommands", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "DisableCommands", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "UnlockWorldTemplateSettings", Shape: Shape{Kind: "void"}}}}, p.Type); err != nil {
		return err
	}
	return nil
}

func DecodeSimpleEvent(r Decoder) (SimpleEvent, error) {
	var p SimpleEvent
	{
		raw, err := r.Read("SimpleEventPacket.Type", Shape{Kind: "enum", Semantic: "SimpleEventPacketPayload::Subtype", TypeID: "enums/SimpleEventPacketPayload::Subtype", PrimitiveCode: "u16le", Variants: []ShapeVariant{{Value: 0, Name: "UninitializedSubtype", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "EnableCommands", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "DisableCommands", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "UnlockWorldTemplateSettings", Shape: Shape{Kind: "void"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(SimpleEventSubtype)
		if !ok {
			return p, fmt.Errorf("field SimpleEventPacket.Type has unexpected decoded type %T", raw)
		}
		p.Type = value
	}
	return p, nil
}
