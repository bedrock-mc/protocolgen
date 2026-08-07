// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type Text struct {
	Localize        bool
	Body            TextBody
	SenderSXUID     string
	PlatformId      string
	FilteredMessage *string
}

func (p *Text) Encode(w Encoder) error {
	if err := w.Write("TextPacket.Localize?", Shape{Kind: "primitive", PrimitiveCode: "bool"}, p.Localize); err != nil {
		return err
	}
	if err := w.Write("TextPacket.Body", Shape{Kind: "union", Control: &Shape{Kind: "primitive", PrimitiveCode: "u8"}, Variants: []ShapeVariant{{Value: 0, Name: "raw", Shape: Shape{Kind: "struct", Semantic: "TextPacketPayload::MessageOnly", TypeID: "TextPacketPayload::MessageOnly", Fields: []ShapeField{{Ordinal: 1, Name: "Message", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}}}, {Value: 1, Name: "chat", Shape: Shape{Kind: "struct", Semantic: "TextPacketPayload::AuthorAndMessage", TypeID: "TextPacketPayload::AuthorAndMessage", Fields: []ShapeField{{Ordinal: 1, Name: "Player Name", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 2, Name: "Message", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}}}, {Value: 2, Name: "translate", Shape: Shape{Kind: "struct", Semantic: "TextPacketPayload::MessageAndParams", TypeID: "TextPacketPayload::MessageAndParams", Fields: []ShapeField{{Ordinal: 1, Name: "Message", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 2, Name: "Parameter List", Shape: Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}}}}, {Value: 3, Name: "popup", Shape: Shape{Kind: "struct", Semantic: "TextPacketPayload::MessageAndParams", TypeID: "TextPacketPayload::MessageAndParams", Fields: []ShapeField{{Ordinal: 1, Name: "Message", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 2, Name: "Parameter List", Shape: Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}}}}, {Value: 4, Name: "jukeboxPopup", Shape: Shape{Kind: "struct", Semantic: "TextPacketPayload::MessageAndParams", TypeID: "TextPacketPayload::MessageAndParams", Fields: []ShapeField{{Ordinal: 1, Name: "Message", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 2, Name: "Parameter List", Shape: Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}}}}, {Value: 5, Name: "tip", Shape: Shape{Kind: "struct", Semantic: "TextPacketPayload::MessageOnly", TypeID: "TextPacketPayload::MessageOnly", Fields: []ShapeField{{Ordinal: 1, Name: "Message", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}}}, {Value: 6, Name: "systemMessage", Shape: Shape{Kind: "struct", Semantic: "TextPacketPayload::MessageOnly", TypeID: "TextPacketPayload::MessageOnly", Fields: []ShapeField{{Ordinal: 1, Name: "Message", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}}}, {Value: 7, Name: "whisper", Shape: Shape{Kind: "struct", Semantic: "TextPacketPayload::AuthorAndMessage", TypeID: "TextPacketPayload::AuthorAndMessage", Fields: []ShapeField{{Ordinal: 1, Name: "Player Name", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 2, Name: "Message", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}}}, {Value: 8, Name: "announcement", Shape: Shape{Kind: "struct", Semantic: "TextPacketPayload::AuthorAndMessage", TypeID: "TextPacketPayload::AuthorAndMessage", Fields: []ShapeField{{Ordinal: 1, Name: "Player Name", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 2, Name: "Message", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}}}, {Value: 9, Name: "textObjectWhisper", Shape: Shape{Kind: "struct", Semantic: "TextPacketPayload::MessageOnly", TypeID: "TextPacketPayload::MessageOnly", Fields: []ShapeField{{Ordinal: 1, Name: "Message", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}}}, {Value: 10, Name: "textObject", Shape: Shape{Kind: "struct", Semantic: "TextPacketPayload::MessageOnly", TypeID: "TextPacketPayload::MessageOnly", Fields: []ShapeField{{Ordinal: 1, Name: "Message", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}}}, {Value: 11, Name: "textObjectAnnouncement", Shape: Shape{Kind: "struct", Semantic: "TextPacketPayload::MessageOnly", TypeID: "TextPacketPayload::MessageOnly", Fields: []ShapeField{{Ordinal: 1, Name: "Message", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}}}}}, p.Body); err != nil {
		return err
	}
	if err := w.Write("TextPacket.Sender's XUID", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, p.SenderSXUID); err != nil {
		return err
	}
	if err := w.Write("TextPacket.Platform Id", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, p.PlatformId); err != nil {
		return err
	}
	if err := w.Write("TextPacket.Filtered Message", Shape{Kind: "optional", Value: &Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, p.FilteredMessage); err != nil {
		return err
	}
	return nil
}

func DecodeText(r Decoder) (Text, error) {
	var p Text
	{
		raw, err := r.Read("TextPacket.Localize?", Shape{Kind: "primitive", PrimitiveCode: "bool"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(bool)
		if !ok {
			return p, fmt.Errorf("field TextPacket.Localize? has unexpected decoded type %T", raw)
		}
		p.Localize = value
	}
	{
		raw, err := r.Read("TextPacket.Body", Shape{Kind: "union", Control: &Shape{Kind: "primitive", PrimitiveCode: "u8"}, Variants: []ShapeVariant{{Value: 0, Name: "raw", Shape: Shape{Kind: "struct", Semantic: "TextPacketPayload::MessageOnly", TypeID: "TextPacketPayload::MessageOnly", Fields: []ShapeField{{Ordinal: 1, Name: "Message", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}}}, {Value: 1, Name: "chat", Shape: Shape{Kind: "struct", Semantic: "TextPacketPayload::AuthorAndMessage", TypeID: "TextPacketPayload::AuthorAndMessage", Fields: []ShapeField{{Ordinal: 1, Name: "Player Name", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 2, Name: "Message", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}}}, {Value: 2, Name: "translate", Shape: Shape{Kind: "struct", Semantic: "TextPacketPayload::MessageAndParams", TypeID: "TextPacketPayload::MessageAndParams", Fields: []ShapeField{{Ordinal: 1, Name: "Message", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 2, Name: "Parameter List", Shape: Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}}}}, {Value: 3, Name: "popup", Shape: Shape{Kind: "struct", Semantic: "TextPacketPayload::MessageAndParams", TypeID: "TextPacketPayload::MessageAndParams", Fields: []ShapeField{{Ordinal: 1, Name: "Message", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 2, Name: "Parameter List", Shape: Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}}}}, {Value: 4, Name: "jukeboxPopup", Shape: Shape{Kind: "struct", Semantic: "TextPacketPayload::MessageAndParams", TypeID: "TextPacketPayload::MessageAndParams", Fields: []ShapeField{{Ordinal: 1, Name: "Message", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 2, Name: "Parameter List", Shape: Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}}}}, {Value: 5, Name: "tip", Shape: Shape{Kind: "struct", Semantic: "TextPacketPayload::MessageOnly", TypeID: "TextPacketPayload::MessageOnly", Fields: []ShapeField{{Ordinal: 1, Name: "Message", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}}}, {Value: 6, Name: "systemMessage", Shape: Shape{Kind: "struct", Semantic: "TextPacketPayload::MessageOnly", TypeID: "TextPacketPayload::MessageOnly", Fields: []ShapeField{{Ordinal: 1, Name: "Message", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}}}, {Value: 7, Name: "whisper", Shape: Shape{Kind: "struct", Semantic: "TextPacketPayload::AuthorAndMessage", TypeID: "TextPacketPayload::AuthorAndMessage", Fields: []ShapeField{{Ordinal: 1, Name: "Player Name", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 2, Name: "Message", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}}}, {Value: 8, Name: "announcement", Shape: Shape{Kind: "struct", Semantic: "TextPacketPayload::AuthorAndMessage", TypeID: "TextPacketPayload::AuthorAndMessage", Fields: []ShapeField{{Ordinal: 1, Name: "Player Name", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 2, Name: "Message", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}}}, {Value: 9, Name: "textObjectWhisper", Shape: Shape{Kind: "struct", Semantic: "TextPacketPayload::MessageOnly", TypeID: "TextPacketPayload::MessageOnly", Fields: []ShapeField{{Ordinal: 1, Name: "Message", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}}}, {Value: 10, Name: "textObject", Shape: Shape{Kind: "struct", Semantic: "TextPacketPayload::MessageOnly", TypeID: "TextPacketPayload::MessageOnly", Fields: []ShapeField{{Ordinal: 1, Name: "Message", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}}}, {Value: 11, Name: "textObjectAnnouncement", Shape: Shape{Kind: "struct", Semantic: "TextPacketPayload::MessageOnly", TypeID: "TextPacketPayload::MessageOnly", Fields: []ShapeField{{Ordinal: 1, Name: "Message", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(TextBody)
		if !ok {
			return p, fmt.Errorf("field TextPacket.Body has unexpected decoded type %T", raw)
		}
		p.Body = value
	}
	{
		raw, err := r.Read("TextPacket.Sender's XUID", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(string)
		if !ok {
			return p, fmt.Errorf("field TextPacket.Sender's XUID has unexpected decoded type %T", raw)
		}
		p.SenderSXUID = value
	}
	{
		raw, err := r.Read("TextPacket.Platform Id", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(string)
		if !ok {
			return p, fmt.Errorf("field TextPacket.Platform Id has unexpected decoded type %T", raw)
		}
		p.PlatformId = value
	}
	{
		raw, err := r.Read("TextPacket.Filtered Message", Shape{Kind: "optional", Value: &Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(*string)
		if !ok {
			return p, fmt.Errorf("field TextPacket.Filtered Message has unexpected decoded type %T", raw)
		}
		p.FilteredMessage = value
	}
	return p, nil
}
