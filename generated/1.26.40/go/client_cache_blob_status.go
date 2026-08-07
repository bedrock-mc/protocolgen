// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type ClientCacheBlobStatus struct {
	MissingIds []uint64
	FoundIds   []uint64
}

func (p *ClientCacheBlobStatus) Encode(w Encoder) error {
	if err := w.Write("ClientCacheBlobStatusPacket.Missing Ids", Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "primitive", PrimitiveCode: "u64le"}}, p.MissingIds); err != nil {
		return err
	}
	if err := w.Write("ClientCacheBlobStatusPacket.Found Ids", Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "primitive", PrimitiveCode: "u64le"}}, p.FoundIds); err != nil {
		return err
	}
	return nil
}

func DecodeClientCacheBlobStatus(r Decoder) (ClientCacheBlobStatus, error) {
	var p ClientCacheBlobStatus
	{
		raw, err := r.Read("ClientCacheBlobStatusPacket.Missing Ids", Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "primitive", PrimitiveCode: "u64le"}})
		if err != nil {
			return p, err
		}
		value, ok := raw.([]uint64)
		if !ok {
			return p, fmt.Errorf("field ClientCacheBlobStatusPacket.Missing Ids has unexpected decoded type %T", raw)
		}
		p.MissingIds = value
	}
	{
		raw, err := r.Read("ClientCacheBlobStatusPacket.Found Ids", Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "primitive", PrimitiveCode: "u64le"}})
		if err != nil {
			return p, err
		}
		value, ok := raw.([]uint64)
		if !ok {
			return p, fmt.Errorf("field ClientCacheBlobStatusPacket.Found Ids has unexpected decoded type %T", raw)
		}
		p.FoundIds = value
	}
	return p, nil
}
