// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type Emote struct {
	ActorRuntimeId   ActorRuntimeID
	EmoteId          string
	EmoteLengthTicks uint32
	Xuid             string
	PlatformId       string
	Flags            uint8
}

func (p *Emote) Encode(w Encoder) error {
	if err := w.Write("EmotePacket.Actor Runtime Id", Shape{Kind: "struct", Semantic: "ActorRuntimeID", TypeID: "ActorRuntimeID", Fields: []ShapeField{{Ordinal: 0, Name: "Actor Runtime ID", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u64"}}}}, p.ActorRuntimeId); err != nil {
		return err
	}
	if err := w.Write("EmotePacket.Emote Id", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, p.EmoteId); err != nil {
		return err
	}
	if err := w.Write("EmotePacket.Emote Length Ticks", Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, p.EmoteLengthTicks); err != nil {
		return err
	}
	if err := w.Write("EmotePacket.Xuid", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, p.Xuid); err != nil {
		return err
	}
	if err := w.Write("EmotePacket.PlatformId", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, p.PlatformId); err != nil {
		return err
	}
	if err := w.Write("EmotePacket.Flags", Shape{Kind: "primitive", PrimitiveCode: "u8"}, p.Flags); err != nil {
		return err
	}
	return nil
}

func DecodeEmote(r Decoder) (Emote, error) {
	var p Emote
	{
		raw, err := r.Read("EmotePacket.Actor Runtime Id", Shape{Kind: "struct", Semantic: "ActorRuntimeID", TypeID: "ActorRuntimeID", Fields: []ShapeField{{Ordinal: 0, Name: "Actor Runtime ID", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u64"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(ActorRuntimeID)
		if !ok {
			return p, fmt.Errorf("field EmotePacket.Actor Runtime Id has unexpected decoded type %T", raw)
		}
		p.ActorRuntimeId = value
	}
	{
		raw, err := r.Read("EmotePacket.Emote Id", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(string)
		if !ok {
			return p, fmt.Errorf("field EmotePacket.Emote Id has unexpected decoded type %T", raw)
		}
		p.EmoteId = value
	}
	{
		raw, err := r.Read("EmotePacket.Emote Length Ticks", Shape{Kind: "primitive", PrimitiveCode: "var_u32"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(uint32)
		if !ok {
			return p, fmt.Errorf("field EmotePacket.Emote Length Ticks has unexpected decoded type %T", raw)
		}
		p.EmoteLengthTicks = value
	}
	{
		raw, err := r.Read("EmotePacket.Xuid", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(string)
		if !ok {
			return p, fmt.Errorf("field EmotePacket.Xuid has unexpected decoded type %T", raw)
		}
		p.Xuid = value
	}
	{
		raw, err := r.Read("EmotePacket.PlatformId", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(string)
		if !ok {
			return p, fmt.Errorf("field EmotePacket.PlatformId has unexpected decoded type %T", raw)
		}
		p.PlatformId = value
	}
	{
		raw, err := r.Read("EmotePacket.Flags", Shape{Kind: "primitive", PrimitiveCode: "u8"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(uint8)
		if !ok {
			return p, fmt.Errorf("field EmotePacket.Flags has unexpected decoded type %T", raw)
		}
		p.Flags = value
	}
	return p, nil
}
