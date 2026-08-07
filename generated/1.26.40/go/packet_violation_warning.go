// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type PacketViolationWarning struct {
	ViolationType     PacketViolationType
	ViolationSeverity PacketViolationSeverity
	ViolationPacketId int32
	ViolationContext  string
}

func (p *PacketViolationWarning) Encode(w Encoder) error {
	if err := w.Write("PacketViolationWarningPacket.Violation Type", Shape{Kind: "enum", Semantic: "PacketViolationType", TypeID: "enums/PacketViolationType", PrimitiveCode: "zigzag_i32", Variants: []ShapeVariant{{Value: -1, Name: "Unknown", Shape: Shape{Kind: "void"}}, {Value: 0, Name: "PacketMalformed", Shape: Shape{Kind: "void"}}}}, p.ViolationType); err != nil {
		return err
	}
	if err := w.Write("PacketViolationWarningPacket.Violation Severity", Shape{Kind: "enum", Semantic: "PacketViolationSeverity", TypeID: "enums/PacketViolationSeverity", PrimitiveCode: "zigzag_i32", Variants: []ShapeVariant{{Value: -1, Name: "Unknown", Shape: Shape{Kind: "void"}}, {Value: 0, Name: "Warning", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "FinalWarning", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "TerminatingConnection", Shape: Shape{Kind: "void"}}}}, p.ViolationSeverity); err != nil {
		return err
	}
	if err := w.Write("PacketViolationWarningPacket.Violation PacketId", Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}, p.ViolationPacketId); err != nil {
		return err
	}
	if err := w.Write("PacketViolationWarningPacket.Violation Context", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, p.ViolationContext); err != nil {
		return err
	}
	return nil
}

func DecodePacketViolationWarning(r Decoder) (PacketViolationWarning, error) {
	var p PacketViolationWarning
	{
		raw, err := r.Read("PacketViolationWarningPacket.Violation Type", Shape{Kind: "enum", Semantic: "PacketViolationType", TypeID: "enums/PacketViolationType", PrimitiveCode: "zigzag_i32", Variants: []ShapeVariant{{Value: -1, Name: "Unknown", Shape: Shape{Kind: "void"}}, {Value: 0, Name: "PacketMalformed", Shape: Shape{Kind: "void"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(PacketViolationType)
		if !ok {
			return p, fmt.Errorf("field PacketViolationWarningPacket.Violation Type has unexpected decoded type %T", raw)
		}
		p.ViolationType = value
	}
	{
		raw, err := r.Read("PacketViolationWarningPacket.Violation Severity", Shape{Kind: "enum", Semantic: "PacketViolationSeverity", TypeID: "enums/PacketViolationSeverity", PrimitiveCode: "zigzag_i32", Variants: []ShapeVariant{{Value: -1, Name: "Unknown", Shape: Shape{Kind: "void"}}, {Value: 0, Name: "Warning", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "FinalWarning", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "TerminatingConnection", Shape: Shape{Kind: "void"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(PacketViolationSeverity)
		if !ok {
			return p, fmt.Errorf("field PacketViolationWarningPacket.Violation Severity has unexpected decoded type %T", raw)
		}
		p.ViolationSeverity = value
	}
	{
		raw, err := r.Read("PacketViolationWarningPacket.Violation PacketId", Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(int32)
		if !ok {
			return p, fmt.Errorf("field PacketViolationWarningPacket.Violation PacketId has unexpected decoded type %T", raw)
		}
		p.ViolationPacketId = value
	}
	{
		raw, err := r.Read("PacketViolationWarningPacket.Violation Context", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(string)
		if !ok {
			return p, fmt.Errorf("field PacketViolationWarningPacket.Violation Context has unexpected decoded type %T", raw)
		}
		p.ViolationContext = value
	}
	return p, nil
}
