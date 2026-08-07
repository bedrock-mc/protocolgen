// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type CommandRequest struct {
	Command    string
	Origin     CommandOriginData
	IsInternal bool
	Version    string
}

func (p *CommandRequest) Encode(w Encoder) error {
	if err := w.Write("CommandRequestPacket.Command", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, p.Command); err != nil {
		return err
	}
	if err := w.Write("CommandRequestPacket.Origin", Shape{Kind: "struct", Semantic: "CommandOriginData", TypeID: "CommandOriginData.json#", Fields: []ShapeField{{Ordinal: 0, Name: "Type", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 1, Name: "UUID", Shape: Shape{Kind: "primitive", Semantic: "mce::UUID", TypeID: "mce__UUID.json#", PrimitiveCode: "uuid"}}, {Ordinal: 2, Name: "RequestId", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 3, Name: "PlayerId", Shape: Shape{Kind: "primitive", PrimitiveCode: "i64le"}}}}, p.Origin); err != nil {
		return err
	}
	if err := w.Write("CommandRequestPacket.IsInternal", Shape{Kind: "primitive", PrimitiveCode: "bool"}, p.IsInternal); err != nil {
		return err
	}
	if err := w.Write("CommandRequestPacket.Version", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, p.Version); err != nil {
		return err
	}
	return nil
}

func DecodeCommandRequest(r Decoder) (CommandRequest, error) {
	var p CommandRequest
	{
		raw, err := r.Read("CommandRequestPacket.Command", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(string)
		if !ok {
			return p, fmt.Errorf("field CommandRequestPacket.Command has unexpected decoded type %T", raw)
		}
		p.Command = value
	}
	{
		raw, err := r.Read("CommandRequestPacket.Origin", Shape{Kind: "struct", Semantic: "CommandOriginData", TypeID: "CommandOriginData.json#", Fields: []ShapeField{{Ordinal: 0, Name: "Type", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 1, Name: "UUID", Shape: Shape{Kind: "primitive", Semantic: "mce::UUID", TypeID: "mce__UUID.json#", PrimitiveCode: "uuid"}}, {Ordinal: 2, Name: "RequestId", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 3, Name: "PlayerId", Shape: Shape{Kind: "primitive", PrimitiveCode: "i64le"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(CommandOriginData)
		if !ok {
			return p, fmt.Errorf("field CommandRequestPacket.Origin has unexpected decoded type %T", raw)
		}
		p.Origin = value
	}
	{
		raw, err := r.Read("CommandRequestPacket.IsInternal", Shape{Kind: "primitive", PrimitiveCode: "bool"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(bool)
		if !ok {
			return p, fmt.Errorf("field CommandRequestPacket.IsInternal has unexpected decoded type %T", raw)
		}
		p.IsInternal = value
	}
	{
		raw, err := r.Read("CommandRequestPacket.Version", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(string)
		if !ok {
			return p, fmt.Errorf("field CommandRequestPacket.Version has unexpected decoded type %T", raw)
		}
		p.Version = value
	}
	return p, nil
}
