// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type SetTitle struct {
	TitleType            SetTitleTitleType
	TitleText            string
	FadeInTime           int32
	StayTime             int32
	FadeOutTime          int32
	Xuid                 string
	PlatformOnlineId     string
	FilteredTitleMessage string
}

func (p *SetTitle) Encode(w Encoder) error {
	if err := w.Write("SetTitlePacket.Title Type", Shape{Kind: "enum", Semantic: "SetTitlePacketPayload::TitleType", TypeID: "enums/SetTitlePacketPayload::TitleType", PrimitiveCode: "zigzag_i32", Variants: []ShapeVariant{{Value: 0, Name: "Clear", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "Reset", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "Title", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "Subtitle", Shape: Shape{Kind: "void"}}, {Value: 4, Name: "Actionbar", Shape: Shape{Kind: "void"}}, {Value: 5, Name: "Times", Shape: Shape{Kind: "void"}}, {Value: 6, Name: "TitleTextObject", Shape: Shape{Kind: "void"}}, {Value: 7, Name: "SubtitleTextObject", Shape: Shape{Kind: "void"}}, {Value: 8, Name: "ActionbarTextObject", Shape: Shape{Kind: "void"}}}}, p.TitleType); err != nil {
		return err
	}
	if err := w.Write("SetTitlePacket.Title Text", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, p.TitleText); err != nil {
		return err
	}
	if err := w.Write("SetTitlePacket.Fade In Time", Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}, p.FadeInTime); err != nil {
		return err
	}
	if err := w.Write("SetTitlePacket.Stay Time", Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}, p.StayTime); err != nil {
		return err
	}
	if err := w.Write("SetTitlePacket.Fade Out Time", Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}, p.FadeOutTime); err != nil {
		return err
	}
	if err := w.Write("SetTitlePacket.Xuid", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, p.Xuid); err != nil {
		return err
	}
	if err := w.Write("SetTitlePacket.Platform Online Id", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, p.PlatformOnlineId); err != nil {
		return err
	}
	if err := w.Write("SetTitlePacket.Filtered Title Message", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, p.FilteredTitleMessage); err != nil {
		return err
	}
	return nil
}

func DecodeSetTitle(r Decoder) (SetTitle, error) {
	var p SetTitle
	{
		raw, err := r.Read("SetTitlePacket.Title Type", Shape{Kind: "enum", Semantic: "SetTitlePacketPayload::TitleType", TypeID: "enums/SetTitlePacketPayload::TitleType", PrimitiveCode: "zigzag_i32", Variants: []ShapeVariant{{Value: 0, Name: "Clear", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "Reset", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "Title", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "Subtitle", Shape: Shape{Kind: "void"}}, {Value: 4, Name: "Actionbar", Shape: Shape{Kind: "void"}}, {Value: 5, Name: "Times", Shape: Shape{Kind: "void"}}, {Value: 6, Name: "TitleTextObject", Shape: Shape{Kind: "void"}}, {Value: 7, Name: "SubtitleTextObject", Shape: Shape{Kind: "void"}}, {Value: 8, Name: "ActionbarTextObject", Shape: Shape{Kind: "void"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(SetTitleTitleType)
		if !ok {
			return p, fmt.Errorf("field SetTitlePacket.Title Type has unexpected decoded type %T", raw)
		}
		p.TitleType = value
	}
	{
		raw, err := r.Read("SetTitlePacket.Title Text", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(string)
		if !ok {
			return p, fmt.Errorf("field SetTitlePacket.Title Text has unexpected decoded type %T", raw)
		}
		p.TitleText = value
	}
	{
		raw, err := r.Read("SetTitlePacket.Fade In Time", Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(int32)
		if !ok {
			return p, fmt.Errorf("field SetTitlePacket.Fade In Time has unexpected decoded type %T", raw)
		}
		p.FadeInTime = value
	}
	{
		raw, err := r.Read("SetTitlePacket.Stay Time", Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(int32)
		if !ok {
			return p, fmt.Errorf("field SetTitlePacket.Stay Time has unexpected decoded type %T", raw)
		}
		p.StayTime = value
	}
	{
		raw, err := r.Read("SetTitlePacket.Fade Out Time", Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(int32)
		if !ok {
			return p, fmt.Errorf("field SetTitlePacket.Fade Out Time has unexpected decoded type %T", raw)
		}
		p.FadeOutTime = value
	}
	{
		raw, err := r.Read("SetTitlePacket.Xuid", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(string)
		if !ok {
			return p, fmt.Errorf("field SetTitlePacket.Xuid has unexpected decoded type %T", raw)
		}
		p.Xuid = value
	}
	{
		raw, err := r.Read("SetTitlePacket.Platform Online Id", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(string)
		if !ok {
			return p, fmt.Errorf("field SetTitlePacket.Platform Online Id has unexpected decoded type %T", raw)
		}
		p.PlatformOnlineId = value
	}
	{
		raw, err := r.Read("SetTitlePacket.Filtered Title Message", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(string)
		if !ok {
			return p, fmt.Errorf("field SetTitlePacket.Filtered Title Message has unexpected decoded type %T", raw)
		}
		p.FilteredTitleMessage = value
	}
	return p, nil
}
