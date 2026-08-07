// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type ClientCacheStatus struct {
	IsCacheSupported bool
}

func (p *ClientCacheStatus) Encode(w Encoder) error {
	if err := w.Write("ClientCacheStatusPacket.Is cache supported?", Shape{Kind: "primitive", PrimitiveCode: "bool"}, p.IsCacheSupported); err != nil {
		return err
	}
	return nil
}

func DecodeClientCacheStatus(r Decoder) (ClientCacheStatus, error) {
	var p ClientCacheStatus
	{
		raw, err := r.Read("ClientCacheStatusPacket.Is cache supported?", Shape{Kind: "primitive", PrimitiveCode: "bool"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(bool)
		if !ok {
			return p, fmt.Errorf("field ClientCacheStatusPacket.Is cache supported? has unexpected decoded type %T", raw)
		}
		p.IsCacheSupported = value
	}
	return p, nil
}
