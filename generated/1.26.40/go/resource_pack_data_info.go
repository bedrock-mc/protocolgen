// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type ResourcePackDataInfo struct {
	ResourceName   string
	ChunkSize      uint32
	NumberOfChunks uint32
	FileSize       uint64
	FileHash       string
	IsPremiumPack  bool
	PackType       uint8
}

func (p *ResourcePackDataInfo) Encode(w Encoder) error {
	if err := w.Write("ResourcePackDataInfoPacket.Resource Name", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, p.ResourceName); err != nil {
		return err
	}
	if err := w.Write("ResourcePackDataInfoPacket.Chunk Size", Shape{Kind: "primitive", PrimitiveCode: "u32le"}, p.ChunkSize); err != nil {
		return err
	}
	if err := w.Write("ResourcePackDataInfoPacket.Number of Chunks", Shape{Kind: "primitive", PrimitiveCode: "u32le"}, p.NumberOfChunks); err != nil {
		return err
	}
	if err := w.Write("ResourcePackDataInfoPacket.File Size", Shape{Kind: "primitive", PrimitiveCode: "u64le"}, p.FileSize); err != nil {
		return err
	}
	if err := w.Write("ResourcePackDataInfoPacket.File Hash", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, p.FileHash); err != nil {
		return err
	}
	if err := w.Write("ResourcePackDataInfoPacket.Is Premium Pack", Shape{Kind: "primitive", PrimitiveCode: "bool"}, p.IsPremiumPack); err != nil {
		return err
	}
	if err := w.Write("ResourcePackDataInfoPacket.Pack Type", Shape{Kind: "primitive", PrimitiveCode: "u8"}, p.PackType); err != nil {
		return err
	}
	return nil
}

func DecodeResourcePackDataInfo(r Decoder) (ResourcePackDataInfo, error) {
	var p ResourcePackDataInfo
	{
		raw, err := r.Read("ResourcePackDataInfoPacket.Resource Name", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(string)
		if !ok {
			return p, fmt.Errorf("field ResourcePackDataInfoPacket.Resource Name has unexpected decoded type %T", raw)
		}
		p.ResourceName = value
	}
	{
		raw, err := r.Read("ResourcePackDataInfoPacket.Chunk Size", Shape{Kind: "primitive", PrimitiveCode: "u32le"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(uint32)
		if !ok {
			return p, fmt.Errorf("field ResourcePackDataInfoPacket.Chunk Size has unexpected decoded type %T", raw)
		}
		p.ChunkSize = value
	}
	{
		raw, err := r.Read("ResourcePackDataInfoPacket.Number of Chunks", Shape{Kind: "primitive", PrimitiveCode: "u32le"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(uint32)
		if !ok {
			return p, fmt.Errorf("field ResourcePackDataInfoPacket.Number of Chunks has unexpected decoded type %T", raw)
		}
		p.NumberOfChunks = value
	}
	{
		raw, err := r.Read("ResourcePackDataInfoPacket.File Size", Shape{Kind: "primitive", PrimitiveCode: "u64le"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(uint64)
		if !ok {
			return p, fmt.Errorf("field ResourcePackDataInfoPacket.File Size has unexpected decoded type %T", raw)
		}
		p.FileSize = value
	}
	{
		raw, err := r.Read("ResourcePackDataInfoPacket.File Hash", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(string)
		if !ok {
			return p, fmt.Errorf("field ResourcePackDataInfoPacket.File Hash has unexpected decoded type %T", raw)
		}
		p.FileHash = value
	}
	{
		raw, err := r.Read("ResourcePackDataInfoPacket.Is Premium Pack", Shape{Kind: "primitive", PrimitiveCode: "bool"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(bool)
		if !ok {
			return p, fmt.Errorf("field ResourcePackDataInfoPacket.Is Premium Pack has unexpected decoded type %T", raw)
		}
		p.IsPremiumPack = value
	}
	{
		raw, err := r.Read("ResourcePackDataInfoPacket.Pack Type", Shape{Kind: "primitive", PrimitiveCode: "u8"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(uint8)
		if !ok {
			return p, fmt.Errorf("field ResourcePackDataInfoPacket.Pack Type has unexpected decoded type %T", raw)
		}
		p.PackType = value
	}
	return p, nil
}
