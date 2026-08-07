// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type ServerStoreInfo struct {
	ClientStoreEntryPointConfiguration *ServerConfigurationClientStoreEntryPointConfiguration
}

func (p *ServerStoreInfo) Encode(w Encoder) error {
	if err := w.Write("ServerStoreInfoPacket.client_store_entry_point_configuration", Shape{Kind: "optional", Value: &Shape{Kind: "struct", Semantic: "ServerConfiguration::ClientStoreEntryPointConfiguration", TypeID: "ServerConfiguration::ClientStoreEntryPointConfiguration", Fields: []ShapeField{{Ordinal: 0, Name: "storeId", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 1, Name: "storeName", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}}}, p.ClientStoreEntryPointConfiguration); err != nil {
		return err
	}
	return nil
}

func DecodeServerStoreInfo(r Decoder) (ServerStoreInfo, error) {
	var p ServerStoreInfo
	{
		raw, err := r.Read("ServerStoreInfoPacket.client_store_entry_point_configuration", Shape{Kind: "optional", Value: &Shape{Kind: "struct", Semantic: "ServerConfiguration::ClientStoreEntryPointConfiguration", TypeID: "ServerConfiguration::ClientStoreEntryPointConfiguration", Fields: []ShapeField{{Ordinal: 0, Name: "storeId", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 1, Name: "storeName", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(*ServerConfigurationClientStoreEntryPointConfiguration)
		if !ok {
			return p, fmt.Errorf("field ServerStoreInfoPacket.client_store_entry_point_configuration has unexpected decoded type %T", raw)
		}
		p.ClientStoreEntryPointConfiguration = value
	}
	return p, nil
}
