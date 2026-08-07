// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type GameTestRequest struct {
	MaxTestsPerBatch int32
	RepeatCount      int32
	Rotation         Rotation
	StopOnFailure    bool
	TestPos          BlockPos
	TestsPerRow      int32
	TestName         string
}

func (p *GameTestRequest) Encode(w Encoder) error {
	if err := w.Write("GameTestRequestPacket.MaxTestsPerBatch", Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}, p.MaxTestsPerBatch); err != nil {
		return err
	}
	if err := w.Write("GameTestRequestPacket.RepeatCount", Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}, p.RepeatCount); err != nil {
		return err
	}
	if err := w.Write("GameTestRequestPacket.Rotation", Shape{Kind: "enum", Semantic: "Rotation", TypeID: "enums/Rotation", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "None", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "Rotate90", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "Rotate180", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "Rotate270", Shape: Shape{Kind: "void"}}}}, p.Rotation); err != nil {
		return err
	}
	if err := w.Write("GameTestRequestPacket.StopOnFailure", Shape{Kind: "primitive", PrimitiveCode: "bool"}, p.StopOnFailure); err != nil {
		return err
	}
	if err := w.Write("GameTestRequestPacket.TestPos", Shape{Kind: "struct", Semantic: "BlockPos", TypeID: "BlockPos", Fields: []ShapeField{{Ordinal: 0, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 1, Name: "Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 2, Name: "Z", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}}}, p.TestPos); err != nil {
		return err
	}
	if err := w.Write("GameTestRequestPacket.TestsPerRow", Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}, p.TestsPerRow); err != nil {
		return err
	}
	if err := w.Write("GameTestRequestPacket.TestName", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, p.TestName); err != nil {
		return err
	}
	return nil
}

func DecodeGameTestRequest(r Decoder) (GameTestRequest, error) {
	var p GameTestRequest
	{
		raw, err := r.Read("GameTestRequestPacket.MaxTestsPerBatch", Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(int32)
		if !ok {
			return p, fmt.Errorf("field GameTestRequestPacket.MaxTestsPerBatch has unexpected decoded type %T", raw)
		}
		p.MaxTestsPerBatch = value
	}
	{
		raw, err := r.Read("GameTestRequestPacket.RepeatCount", Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(int32)
		if !ok {
			return p, fmt.Errorf("field GameTestRequestPacket.RepeatCount has unexpected decoded type %T", raw)
		}
		p.RepeatCount = value
	}
	{
		raw, err := r.Read("GameTestRequestPacket.Rotation", Shape{Kind: "enum", Semantic: "Rotation", TypeID: "enums/Rotation", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "None", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "Rotate90", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "Rotate180", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "Rotate270", Shape: Shape{Kind: "void"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(Rotation)
		if !ok {
			return p, fmt.Errorf("field GameTestRequestPacket.Rotation has unexpected decoded type %T", raw)
		}
		p.Rotation = value
	}
	{
		raw, err := r.Read("GameTestRequestPacket.StopOnFailure", Shape{Kind: "primitive", PrimitiveCode: "bool"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(bool)
		if !ok {
			return p, fmt.Errorf("field GameTestRequestPacket.StopOnFailure has unexpected decoded type %T", raw)
		}
		p.StopOnFailure = value
	}
	{
		raw, err := r.Read("GameTestRequestPacket.TestPos", Shape{Kind: "struct", Semantic: "BlockPos", TypeID: "BlockPos", Fields: []ShapeField{{Ordinal: 0, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 1, Name: "Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 2, Name: "Z", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(BlockPos)
		if !ok {
			return p, fmt.Errorf("field GameTestRequestPacket.TestPos has unexpected decoded type %T", raw)
		}
		p.TestPos = value
	}
	{
		raw, err := r.Read("GameTestRequestPacket.TestsPerRow", Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(int32)
		if !ok {
			return p, fmt.Errorf("field GameTestRequestPacket.TestsPerRow has unexpected decoded type %T", raw)
		}
		p.TestsPerRow = value
	}
	{
		raw, err := r.Read("GameTestRequestPacket.TestName", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(string)
		if !ok {
			return p, fmt.Errorf("field GameTestRequestPacket.TestName has unexpected decoded type %T", raw)
		}
		p.TestName = value
	}
	return p, nil
}
