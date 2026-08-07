// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type PlayerVideoCapture struct {
	Action PlayerVideoCaptureAction
}

func (p *PlayerVideoCapture) Encode(w Encoder) error {
	if err := w.Write("PlayerVideoCapturePacket.Action", Shape{Kind: "union", Control: &Shape{Kind: "primitive", PrimitiveCode: "u8"}, Variants: []ShapeVariant{{Value: 0, Name: "PlayerVideoCapturePacketPayload::StopVideoCapture", Shape: Shape{Kind: "struct", Semantic: "PlayerVideoCapturePacketPayload::StopVideoCapture", TypeID: "PlayerVideoCapturePacketPayload::StopVideoCapture"}}, {Value: 1, Name: "PlayerVideoCapturePacketPayload::StartVideoCapture", Shape: Shape{Kind: "struct", Semantic: "PlayerVideoCapturePacketPayload::StartVideoCapture", TypeID: "PlayerVideoCapturePacketPayload::StartVideoCapture", Fields: []ShapeField{{Ordinal: 0, Name: "FrameRate", Shape: Shape{Kind: "primitive", PrimitiveCode: "u32le"}}, {Ordinal: 1, Name: "FilePrefix", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}}}}}, p.Action); err != nil {
		return err
	}
	return nil
}

func DecodePlayerVideoCapture(r Decoder) (PlayerVideoCapture, error) {
	var p PlayerVideoCapture
	{
		raw, err := r.Read("PlayerVideoCapturePacket.Action", Shape{Kind: "union", Control: &Shape{Kind: "primitive", PrimitiveCode: "u8"}, Variants: []ShapeVariant{{Value: 0, Name: "PlayerVideoCapturePacketPayload::StopVideoCapture", Shape: Shape{Kind: "struct", Semantic: "PlayerVideoCapturePacketPayload::StopVideoCapture", TypeID: "PlayerVideoCapturePacketPayload::StopVideoCapture"}}, {Value: 1, Name: "PlayerVideoCapturePacketPayload::StartVideoCapture", Shape: Shape{Kind: "struct", Semantic: "PlayerVideoCapturePacketPayload::StartVideoCapture", TypeID: "PlayerVideoCapturePacketPayload::StartVideoCapture", Fields: []ShapeField{{Ordinal: 0, Name: "FrameRate", Shape: Shape{Kind: "primitive", PrimitiveCode: "u32le"}}, {Ordinal: 1, Name: "FilePrefix", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(PlayerVideoCaptureAction)
		if !ok {
			return p, fmt.Errorf("field PlayerVideoCapturePacket.Action has unexpected decoded type %T", raw)
		}
		p.Action = value
	}
	return p, nil
}
