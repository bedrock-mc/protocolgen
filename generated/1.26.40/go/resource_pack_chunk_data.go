// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type ResourcePackChunkData struct {
	ResourceName string
	ChunkID      uint32
	ByteOffset   uint64
	ChunkData    string
}

func (p *ResourcePackChunkData) Encode(w Encoder) error {
	if err := w.Write("ResourcePackChunkDataPacket.Resource Name", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, p.ResourceName); err != nil {
		return err
	}
	if err := w.Write("ResourcePackChunkDataPacket.Chunk ID", Shape{Kind: "primitive", PrimitiveCode: "u32le"}, p.ChunkID); err != nil {
		return err
	}
	if err := w.Write("ResourcePackChunkDataPacket.Byte Offset", Shape{Kind: "primitive", PrimitiveCode: "u64le"}, p.ByteOffset); err != nil {
		return err
	}
	if err := w.Write("ResourcePackChunkDataPacket.Chunk Data", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, p.ChunkData); err != nil {
		return err
	}
	return nil
}

func DecodeResourcePackChunkData(r Decoder) (ResourcePackChunkData, error) {
	var p ResourcePackChunkData
	{
		raw, err := r.Read("ResourcePackChunkDataPacket.Resource Name", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(string)
		if !ok {
			return p, fmt.Errorf("field ResourcePackChunkDataPacket.Resource Name has unexpected decoded type %T", raw)
		}
		p.ResourceName = value
	}
	{
		raw, err := r.Read("ResourcePackChunkDataPacket.Chunk ID", Shape{Kind: "primitive", PrimitiveCode: "u32le"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(uint32)
		if !ok {
			return p, fmt.Errorf("field ResourcePackChunkDataPacket.Chunk ID has unexpected decoded type %T", raw)
		}
		p.ChunkID = value
	}
	{
		raw, err := r.Read("ResourcePackChunkDataPacket.Byte Offset", Shape{Kind: "primitive", PrimitiveCode: "u64le"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(uint64)
		if !ok {
			return p, fmt.Errorf("field ResourcePackChunkDataPacket.Byte Offset has unexpected decoded type %T", raw)
		}
		p.ByteOffset = value
	}
	{
		raw, err := r.Read("ResourcePackChunkDataPacket.Chunk Data", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(string)
		if !ok {
			return p, fmt.Errorf("field ResourcePackChunkDataPacket.Chunk Data has unexpected decoded type %T", raw)
		}
		p.ChunkData = value
	}
	return p, nil
}
