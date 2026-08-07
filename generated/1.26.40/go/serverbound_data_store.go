// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type ServerboundDataStore struct {
	Update BedrockDDUIDataStoreUpdate
}

func (p *ServerboundDataStore) Encode(w Encoder) error {
	if err := w.Write("ServerboundDataStorePacket.Update", Shape{Kind: "struct", Semantic: "Bedrock::DDUI::DataStoreUpdate", TypeID: "Bedrock::DDUI::DataStoreUpdate", Fields: []ShapeField{{Ordinal: 0, Name: "Data Store Name", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 1, Name: "Property", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 2, Name: "Path", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 3, Name: "Data", Shape: Shape{Kind: "union", Control: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Variants: []ShapeVariant{{Value: 0, Name: "double", Shape: Shape{Kind: "primitive", PrimitiveCode: "f64le"}}, {Value: 1, Name: "bool", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}, {Value: 2, Name: "string", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}}}, {Ordinal: 4, Name: "Property Update Count", Shape: Shape{Kind: "primitive", PrimitiveCode: "u32le"}}, {Ordinal: 5, Name: "Path Update Count", Shape: Shape{Kind: "primitive", PrimitiveCode: "u32le"}}}}, p.Update); err != nil {
		return err
	}
	return nil
}

func DecodeServerboundDataStore(r Decoder) (ServerboundDataStore, error) {
	var p ServerboundDataStore
	{
		raw, err := r.Read("ServerboundDataStorePacket.Update", Shape{Kind: "struct", Semantic: "Bedrock::DDUI::DataStoreUpdate", TypeID: "Bedrock::DDUI::DataStoreUpdate", Fields: []ShapeField{{Ordinal: 0, Name: "Data Store Name", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 1, Name: "Property", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 2, Name: "Path", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 3, Name: "Data", Shape: Shape{Kind: "union", Control: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Variants: []ShapeVariant{{Value: 0, Name: "double", Shape: Shape{Kind: "primitive", PrimitiveCode: "f64le"}}, {Value: 1, Name: "bool", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}, {Value: 2, Name: "string", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}}}, {Ordinal: 4, Name: "Property Update Count", Shape: Shape{Kind: "primitive", PrimitiveCode: "u32le"}}, {Ordinal: 5, Name: "Path Update Count", Shape: Shape{Kind: "primitive", PrimitiveCode: "u32le"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(BedrockDDUIDataStoreUpdate)
		if !ok {
			return p, fmt.Errorf("field ServerboundDataStorePacket.Update has unexpected decoded type %T", raw)
		}
		p.Update = value
	}
	return p, nil
}
