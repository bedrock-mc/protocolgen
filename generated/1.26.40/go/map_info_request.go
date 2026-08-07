// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type MapInfoRequest struct {
	MapUniqueID      ActorUniqueID
	ClientPixelsList []MapInfoRequestPacketAnonClientPixelsProxy
}

func (p *MapInfoRequest) Encode(w Encoder) error {
	if err := w.Write("MapInfoRequestPacket.Map Unique ID", Shape{Kind: "struct", Semantic: "ActorUniqueID", TypeID: "ActorUniqueID", Fields: []ShapeField{{Ordinal: 0, Name: "Actor Unique ID", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i64"}}}}, p.MapUniqueID); err != nil {
		return err
	}
	if err := w.Write("MapInfoRequestPacket.Client Pixels List", Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "u32le"}, Element: &Shape{Kind: "struct", Semantic: "MapInfoRequestPacketAnon::ClientPixelsProxy", TypeID: "MapInfoRequestPacketAnon::ClientPixelsProxy", Fields: []ShapeField{{Ordinal: 0, Name: "pixel", Shape: Shape{Kind: "primitive", PrimitiveCode: "u32le"}}, {Ordinal: 1, Name: "index", Shape: Shape{Kind: "primitive", PrimitiveCode: "u16le"}}}}}, p.ClientPixelsList); err != nil {
		return err
	}
	return nil
}

func DecodeMapInfoRequest(r Decoder) (MapInfoRequest, error) {
	var p MapInfoRequest
	{
		raw, err := r.Read("MapInfoRequestPacket.Map Unique ID", Shape{Kind: "struct", Semantic: "ActorUniqueID", TypeID: "ActorUniqueID", Fields: []ShapeField{{Ordinal: 0, Name: "Actor Unique ID", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i64"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(ActorUniqueID)
		if !ok {
			return p, fmt.Errorf("field MapInfoRequestPacket.Map Unique ID has unexpected decoded type %T", raw)
		}
		p.MapUniqueID = value
	}
	{
		raw, err := r.Read("MapInfoRequestPacket.Client Pixels List", Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "u32le"}, Element: &Shape{Kind: "struct", Semantic: "MapInfoRequestPacketAnon::ClientPixelsProxy", TypeID: "MapInfoRequestPacketAnon::ClientPixelsProxy", Fields: []ShapeField{{Ordinal: 0, Name: "pixel", Shape: Shape{Kind: "primitive", PrimitiveCode: "u32le"}}, {Ordinal: 1, Name: "index", Shape: Shape{Kind: "primitive", PrimitiveCode: "u16le"}}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.([]MapInfoRequestPacketAnonClientPixelsProxy)
		if !ok {
			return p, fmt.Errorf("field MapInfoRequestPacket.Client Pixels List has unexpected decoded type %T", raw)
		}
		p.ClientPixelsList = value
	}
	return p, nil
}
