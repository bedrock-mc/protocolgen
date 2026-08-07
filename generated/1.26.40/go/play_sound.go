// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type PlaySound struct {
	Name              string
	Position          BlockPos
	Volume            float32
	Pitch             float32
	LoopCount         int32
	ServerSoundHandle *ServerSoundHandle
}

func (p *PlaySound) Encode(w Encoder) error {
	if err := w.Write("PlaySoundPacket.Name", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, p.Name); err != nil {
		return err
	}
	if err := w.Write("PlaySoundPacket.Position", Shape{Kind: "struct", Semantic: "BlockPos", TypeID: "BlockPos", Fields: []ShapeField{{Ordinal: 0, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 1, Name: "Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 2, Name: "Z", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}}}, p.Position); err != nil {
		return err
	}
	if err := w.Write("PlaySoundPacket.Volume", Shape{Kind: "primitive", PrimitiveCode: "f32le"}, p.Volume); err != nil {
		return err
	}
	if err := w.Write("PlaySoundPacket.Pitch", Shape{Kind: "primitive", PrimitiveCode: "f32le"}, p.Pitch); err != nil {
		return err
	}
	if err := w.Write("PlaySoundPacket.Loop Count", Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}, p.LoopCount); err != nil {
		return err
	}
	if err := w.Write("PlaySoundPacket.Server Sound Handle", Shape{Kind: "optional", Value: &Shape{Kind: "struct", Semantic: "ServerSoundHandle", TypeID: "ServerSoundHandle", Fields: []ShapeField{{Ordinal: 0, Name: "Server Sound Handle", Shape: Shape{Kind: "primitive", PrimitiveCode: "u64le"}}}}}, p.ServerSoundHandle); err != nil {
		return err
	}
	return nil
}

func DecodePlaySound(r Decoder) (PlaySound, error) {
	var p PlaySound
	{
		raw, err := r.Read("PlaySoundPacket.Name", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(string)
		if !ok {
			return p, fmt.Errorf("field PlaySoundPacket.Name has unexpected decoded type %T", raw)
		}
		p.Name = value
	}
	{
		raw, err := r.Read("PlaySoundPacket.Position", Shape{Kind: "struct", Semantic: "BlockPos", TypeID: "BlockPos", Fields: []ShapeField{{Ordinal: 0, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 1, Name: "Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 2, Name: "Z", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(BlockPos)
		if !ok {
			return p, fmt.Errorf("field PlaySoundPacket.Position has unexpected decoded type %T", raw)
		}
		p.Position = value
	}
	{
		raw, err := r.Read("PlaySoundPacket.Volume", Shape{Kind: "primitive", PrimitiveCode: "f32le"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(float32)
		if !ok {
			return p, fmt.Errorf("field PlaySoundPacket.Volume has unexpected decoded type %T", raw)
		}
		p.Volume = value
	}
	{
		raw, err := r.Read("PlaySoundPacket.Pitch", Shape{Kind: "primitive", PrimitiveCode: "f32le"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(float32)
		if !ok {
			return p, fmt.Errorf("field PlaySoundPacket.Pitch has unexpected decoded type %T", raw)
		}
		p.Pitch = value
	}
	{
		raw, err := r.Read("PlaySoundPacket.Loop Count", Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(int32)
		if !ok {
			return p, fmt.Errorf("field PlaySoundPacket.Loop Count has unexpected decoded type %T", raw)
		}
		p.LoopCount = value
	}
	{
		raw, err := r.Read("PlaySoundPacket.Server Sound Handle", Shape{Kind: "optional", Value: &Shape{Kind: "struct", Semantic: "ServerSoundHandle", TypeID: "ServerSoundHandle", Fields: []ShapeField{{Ordinal: 0, Name: "Server Sound Handle", Shape: Shape{Kind: "primitive", PrimitiveCode: "u64le"}}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(*ServerSoundHandle)
		if !ok {
			return p, fmt.Errorf("field PlaySoundPacket.Server Sound Handle has unexpected decoded type %T", raw)
		}
		p.ServerSoundHandle = value
	}
	return p, nil
}
