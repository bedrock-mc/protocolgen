// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type ResourcePackClientResponse struct {
	Response ResourcePackClientResponseResponse
}

func (p *ResourcePackClientResponse) Encode(w Encoder) error {
	if err := w.Write("ResourcePackClientResponsePacket.Response", Shape{Kind: "union", Control: &Shape{Kind: "primitive", PrimitiveCode: "i8"}, Variants: []ShapeVariant{{Value: 1, Name: "ResourcePackClientResponsePacketPayload::Cancel", Shape: Shape{Kind: "struct", Semantic: "ResourcePackClientResponsePacketPayload::Cancel", TypeID: "ResourcePackClientResponsePacketPayload::Cancel", Fields: []ShapeField{{Ordinal: 0, Name: "Response Type", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}}}, {Value: 2, Name: "ResourcePackClientResponsePacketPayload::Downloading", Shape: Shape{Kind: "struct", Semantic: "ResourcePackClientResponsePacketPayload::Downloading", TypeID: "ResourcePackClientResponsePacketPayload::Downloading", Fields: []ShapeField{{Ordinal: 0, Name: "Response Type", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 1, Name: "Downloading Packs", Shape: Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}}}}, {Value: 3, Name: "ResourcePackClientResponsePacketPayload::DownloadingFinished", Shape: Shape{Kind: "struct", Semantic: "ResourcePackClientResponsePacketPayload::DownloadingFinished", TypeID: "ResourcePackClientResponsePacketPayload::DownloadingFinished", Fields: []ShapeField{{Ordinal: 0, Name: "Response Type", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}}}, {Value: 4, Name: "ResourcePackClientResponsePacketPayload::ResourcePackStackFinished", Shape: Shape{Kind: "struct", Semantic: "ResourcePackClientResponsePacketPayload::ResourcePackStackFinished", TypeID: "ResourcePackClientResponsePacketPayload::ResourcePackStackFinished", Fields: []ShapeField{{Ordinal: 0, Name: "Response Type", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}}}}}, p.Response); err != nil {
		return err
	}
	return nil
}

func DecodeResourcePackClientResponse(r Decoder) (ResourcePackClientResponse, error) {
	var p ResourcePackClientResponse
	{
		raw, err := r.Read("ResourcePackClientResponsePacket.Response", Shape{Kind: "union", Control: &Shape{Kind: "primitive", PrimitiveCode: "i8"}, Variants: []ShapeVariant{{Value: 1, Name: "ResourcePackClientResponsePacketPayload::Cancel", Shape: Shape{Kind: "struct", Semantic: "ResourcePackClientResponsePacketPayload::Cancel", TypeID: "ResourcePackClientResponsePacketPayload::Cancel", Fields: []ShapeField{{Ordinal: 0, Name: "Response Type", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}}}, {Value: 2, Name: "ResourcePackClientResponsePacketPayload::Downloading", Shape: Shape{Kind: "struct", Semantic: "ResourcePackClientResponsePacketPayload::Downloading", TypeID: "ResourcePackClientResponsePacketPayload::Downloading", Fields: []ShapeField{{Ordinal: 0, Name: "Response Type", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 1, Name: "Downloading Packs", Shape: Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}}}}, {Value: 3, Name: "ResourcePackClientResponsePacketPayload::DownloadingFinished", Shape: Shape{Kind: "struct", Semantic: "ResourcePackClientResponsePacketPayload::DownloadingFinished", TypeID: "ResourcePackClientResponsePacketPayload::DownloadingFinished", Fields: []ShapeField{{Ordinal: 0, Name: "Response Type", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}}}, {Value: 4, Name: "ResourcePackClientResponsePacketPayload::ResourcePackStackFinished", Shape: Shape{Kind: "struct", Semantic: "ResourcePackClientResponsePacketPayload::ResourcePackStackFinished", TypeID: "ResourcePackClientResponsePacketPayload::ResourcePackStackFinished", Fields: []ShapeField{{Ordinal: 0, Name: "Response Type", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(ResourcePackClientResponseResponse)
		if !ok {
			return p, fmt.Errorf("field ResourcePackClientResponsePacket.Response has unexpected decoded type %T", raw)
		}
		p.Response = value
	}
	return p, nil
}
