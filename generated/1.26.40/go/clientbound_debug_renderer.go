// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type ClientboundDebugRenderer struct {
	Type            string
	DebugMarkerData *ClientboundDebugRendererDebugMarkerData
}

func (p *ClientboundDebugRenderer) Encode(w Encoder) error {
	if err := w.Write("ClientboundDebugRendererPacket.Type", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, p.Type); err != nil {
		return err
	}
	if err := w.Write("ClientboundDebugRendererPacket.DebugMarkerData", Shape{Kind: "optional", Value: &Shape{Kind: "struct", Semantic: "ClientboundDebugRendererPacketPayload::DebugMarkerData", TypeID: "ClientboundDebugRendererPacketPayload::DebugMarkerData", Fields: []ShapeField{{Ordinal: 0, Name: "Text", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 1, Name: "Position", Shape: Shape{Kind: "struct", Semantic: "Vec3", TypeID: "Vec3", Fields: []ShapeField{{Ordinal: 0, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 1, Name: "Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 2, Name: "Z", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}}}}, {Ordinal: 2, Name: "Color", Shape: Shape{Kind: "struct", Semantic: "mce::Color", TypeID: "mce::Color", Fields: []ShapeField{{Ordinal: 0, Name: "Color", Shape: Shape{Kind: "primitive", PrimitiveCode: "i32le"}}}}}, {Ordinal: 3, Name: "duration", Shape: Shape{Kind: "primitive", PrimitiveCode: "u64le"}}}}}, p.DebugMarkerData); err != nil {
		return err
	}
	return nil
}

func DecodeClientboundDebugRenderer(r Decoder) (ClientboundDebugRenderer, error) {
	var p ClientboundDebugRenderer
	{
		raw, err := r.Read("ClientboundDebugRendererPacket.Type", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(string)
		if !ok {
			return p, fmt.Errorf("field ClientboundDebugRendererPacket.Type has unexpected decoded type %T", raw)
		}
		p.Type = value
	}
	{
		raw, err := r.Read("ClientboundDebugRendererPacket.DebugMarkerData", Shape{Kind: "optional", Value: &Shape{Kind: "struct", Semantic: "ClientboundDebugRendererPacketPayload::DebugMarkerData", TypeID: "ClientboundDebugRendererPacketPayload::DebugMarkerData", Fields: []ShapeField{{Ordinal: 0, Name: "Text", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 1, Name: "Position", Shape: Shape{Kind: "struct", Semantic: "Vec3", TypeID: "Vec3", Fields: []ShapeField{{Ordinal: 0, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 1, Name: "Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 2, Name: "Z", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}}}}, {Ordinal: 2, Name: "Color", Shape: Shape{Kind: "struct", Semantic: "mce::Color", TypeID: "mce::Color", Fields: []ShapeField{{Ordinal: 0, Name: "Color", Shape: Shape{Kind: "primitive", PrimitiveCode: "i32le"}}}}}, {Ordinal: 3, Name: "duration", Shape: Shape{Kind: "primitive", PrimitiveCode: "u64le"}}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(*ClientboundDebugRendererDebugMarkerData)
		if !ok {
			return p, fmt.Errorf("field ClientboundDebugRendererPacket.DebugMarkerData has unexpected decoded type %T", raw)
		}
		p.DebugMarkerData = value
	}
	return p, nil
}
