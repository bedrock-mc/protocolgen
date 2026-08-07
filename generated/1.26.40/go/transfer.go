// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type Transfer struct {
	ServerAddress           string
	ServerPort              uint16
	ReloadWorld             bool
	GatheringsConfiguration *ServerConfigurationGatheringsConfigurationJoinInfo
}

func (p *Transfer) Encode(w Encoder) error {
	if err := w.Write("TransferPacket.Server Address", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, p.ServerAddress); err != nil {
		return err
	}
	if err := w.Write("TransferPacket.Server Port", Shape{Kind: "primitive", PrimitiveCode: "u16le"}, p.ServerPort); err != nil {
		return err
	}
	if err := w.Write("TransferPacket.Reload World", Shape{Kind: "primitive", PrimitiveCode: "bool"}, p.ReloadWorld); err != nil {
		return err
	}
	if err := w.Write("TransferPacket.Gatherings Configuration", Shape{Kind: "optional", Value: &Shape{Kind: "struct", Semantic: "ServerConfiguration::GatheringsConfigurationJoinInfo", TypeID: "ServerConfiguration::GatheringsConfigurationJoinInfo", Fields: []ShapeField{{Ordinal: 0, Name: "experienceId", Shape: Shape{Kind: "primitive", PrimitiveCode: "uuid"}}, {Ordinal: 1, Name: "experienceName", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 2, Name: "worldId", Shape: Shape{Kind: "optional", Value: &Shape{Kind: "primitive", PrimitiveCode: "uuid"}}}, {Ordinal: 3, Name: "worldName", Shape: Shape{Kind: "optional", Value: &Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}, {Ordinal: 4, Name: "creatorId", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 5, Name: "targetId", Shape: Shape{Kind: "optional", Value: &Shape{Kind: "primitive", PrimitiveCode: "uuid"}}}, {Ordinal: 6, Name: "scenarioId", Shape: Shape{Kind: "optional", Value: &Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}, {Ordinal: 7, Name: "serverId", Shape: Shape{Kind: "optional", Value: &Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}}}}, p.GatheringsConfiguration); err != nil {
		return err
	}
	return nil
}

func DecodeTransfer(r Decoder) (Transfer, error) {
	var p Transfer
	{
		raw, err := r.Read("TransferPacket.Server Address", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(string)
		if !ok {
			return p, fmt.Errorf("field TransferPacket.Server Address has unexpected decoded type %T", raw)
		}
		p.ServerAddress = value
	}
	{
		raw, err := r.Read("TransferPacket.Server Port", Shape{Kind: "primitive", PrimitiveCode: "u16le"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(uint16)
		if !ok {
			return p, fmt.Errorf("field TransferPacket.Server Port has unexpected decoded type %T", raw)
		}
		p.ServerPort = value
	}
	{
		raw, err := r.Read("TransferPacket.Reload World", Shape{Kind: "primitive", PrimitiveCode: "bool"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(bool)
		if !ok {
			return p, fmt.Errorf("field TransferPacket.Reload World has unexpected decoded type %T", raw)
		}
		p.ReloadWorld = value
	}
	{
		raw, err := r.Read("TransferPacket.Gatherings Configuration", Shape{Kind: "optional", Value: &Shape{Kind: "struct", Semantic: "ServerConfiguration::GatheringsConfigurationJoinInfo", TypeID: "ServerConfiguration::GatheringsConfigurationJoinInfo", Fields: []ShapeField{{Ordinal: 0, Name: "experienceId", Shape: Shape{Kind: "primitive", PrimitiveCode: "uuid"}}, {Ordinal: 1, Name: "experienceName", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 2, Name: "worldId", Shape: Shape{Kind: "optional", Value: &Shape{Kind: "primitive", PrimitiveCode: "uuid"}}}, {Ordinal: 3, Name: "worldName", Shape: Shape{Kind: "optional", Value: &Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}, {Ordinal: 4, Name: "creatorId", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 5, Name: "targetId", Shape: Shape{Kind: "optional", Value: &Shape{Kind: "primitive", PrimitiveCode: "uuid"}}}, {Ordinal: 6, Name: "scenarioId", Shape: Shape{Kind: "optional", Value: &Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}, {Ordinal: 7, Name: "serverId", Shape: Shape{Kind: "optional", Value: &Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(*ServerConfigurationGatheringsConfigurationJoinInfo)
		if !ok {
			return p, fmt.Errorf("field TransferPacket.Gatherings Configuration has unexpected decoded type %T", raw)
		}
		p.GatheringsConfiguration = value
	}
	return p, nil
}
