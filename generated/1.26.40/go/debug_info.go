// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type DebugInfo struct {
	ActorId ActorUniqueID
	Data    string
}

func (p *DebugInfo) Encode(w Encoder) error {
	if err := w.Write("DebugInfoPacket.Actor Id", Shape{Kind: "struct", Semantic: "ActorUniqueID", TypeID: "ActorUniqueID", Fields: []ShapeField{{Ordinal: 0, Name: "Actor Unique ID", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i64"}}}}, p.ActorId); err != nil {
		return err
	}
	if err := w.Write("DebugInfoPacket.Data", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, p.Data); err != nil {
		return err
	}
	return nil
}

func DecodeDebugInfo(r Decoder) (DebugInfo, error) {
	var p DebugInfo
	{
		raw, err := r.Read("DebugInfoPacket.Actor Id", Shape{Kind: "struct", Semantic: "ActorUniqueID", TypeID: "ActorUniqueID", Fields: []ShapeField{{Ordinal: 0, Name: "Actor Unique ID", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i64"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(ActorUniqueID)
		if !ok {
			return p, fmt.Errorf("field DebugInfoPacket.Actor Id has unexpected decoded type %T", raw)
		}
		p.ActorId = value
	}
	{
		raw, err := r.Read("DebugInfoPacket.Data", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(string)
		if !ok {
			return p, fmt.Errorf("field DebugInfoPacket.Data has unexpected decoded type %T", raw)
		}
		p.Data = value
	}
	return p, nil
}
