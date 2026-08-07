// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type ResourcePackChunkRequest struct {
	ResourceName string
	Chunk        int32
}

func (p *ResourcePackChunkRequest) Encode(w Encoder) error {
	if err := w.Write("ResourcePackChunkRequestPacket.Resource Name", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, p.ResourceName); err != nil {
		return err
	}
	if err := w.Write("ResourcePackChunkRequestPacket.Chunk", Shape{Kind: "primitive", PrimitiveCode: "i32le"}, p.Chunk); err != nil {
		return err
	}
	return nil
}

func DecodeResourcePackChunkRequest(r Decoder) (ResourcePackChunkRequest, error) {
	var p ResourcePackChunkRequest
	{
		raw, err := r.Read("ResourcePackChunkRequestPacket.Resource Name", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(string)
		if !ok {
			return p, fmt.Errorf("field ResourcePackChunkRequestPacket.Resource Name has unexpected decoded type %T", raw)
		}
		p.ResourceName = value
	}
	{
		raw, err := r.Read("ResourcePackChunkRequestPacket.Chunk", Shape{Kind: "primitive", PrimitiveCode: "i32le"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(int32)
		if !ok {
			return p, fmt.Errorf("field ResourcePackChunkRequestPacket.Chunk has unexpected decoded type %T", raw)
		}
		p.Chunk = value
	}
	return p, nil
}
