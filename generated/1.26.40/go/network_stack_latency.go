// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type NetworkStackLatency struct {
	CreationTime uint64
	IsFromServer bool
}

func (p *NetworkStackLatency) Encode(w Encoder) error {
	if err := w.Write("NetworkStackLatencyPacket.Creation Time", Shape{Kind: "primitive", PrimitiveCode: "u64le"}, p.CreationTime); err != nil {
		return err
	}
	if err := w.Write("NetworkStackLatencyPacket.Is From Server", Shape{Kind: "primitive", PrimitiveCode: "bool"}, p.IsFromServer); err != nil {
		return err
	}
	return nil
}

func DecodeNetworkStackLatency(r Decoder) (NetworkStackLatency, error) {
	var p NetworkStackLatency
	{
		raw, err := r.Read("NetworkStackLatencyPacket.Creation Time", Shape{Kind: "primitive", PrimitiveCode: "u64le"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(uint64)
		if !ok {
			return p, fmt.Errorf("field NetworkStackLatencyPacket.Creation Time has unexpected decoded type %T", raw)
		}
		p.CreationTime = value
	}
	{
		raw, err := r.Read("NetworkStackLatencyPacket.Is From Server", Shape{Kind: "primitive", PrimitiveCode: "bool"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(bool)
		if !ok {
			return p, fmt.Errorf("field NetworkStackLatencyPacket.Is From Server has unexpected decoded type %T", raw)
		}
		p.IsFromServer = value
	}
	return p, nil
}
