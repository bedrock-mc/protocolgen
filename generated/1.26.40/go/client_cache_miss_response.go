// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type ClientCacheMissResponse struct {
	MissingBlobs []MissingBlobData
}

func (p *ClientCacheMissResponse) Encode(w Encoder) error {
	if err := w.Write("ClientCacheMissResponsePacket.Missing Blobs", Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "struct", Semantic: "MissingBlobData", TypeID: "MissingBlobData", Fields: []ShapeField{{Ordinal: 0, Name: "Blob Id", Shape: Shape{Kind: "primitive", PrimitiveCode: "u64le"}}, {Ordinal: 1, Name: "Blob Data", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}}}, p.MissingBlobs); err != nil {
		return err
	}
	return nil
}

func DecodeClientCacheMissResponse(r Decoder) (ClientCacheMissResponse, error) {
	var p ClientCacheMissResponse
	{
		raw, err := r.Read("ClientCacheMissResponsePacket.Missing Blobs", Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "struct", Semantic: "MissingBlobData", TypeID: "MissingBlobData", Fields: []ShapeField{{Ordinal: 0, Name: "Blob Id", Shape: Shape{Kind: "primitive", PrimitiveCode: "u64le"}}, {Ordinal: 1, Name: "Blob Data", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.([]MissingBlobData)
		if !ok {
			return p, fmt.Errorf("field ClientCacheMissResponsePacket.Missing Blobs has unexpected decoded type %T", raw)
		}
		p.MissingBlobs = value
	}
	return p, nil
}
