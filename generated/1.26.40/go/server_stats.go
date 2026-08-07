// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type ServerStats struct {
	ServerTime  float32
	NetworkTime float32
}

func (p *ServerStats) Encode(w Encoder) error {
	if err := w.Write("ServerStatsPacket.ServerTime", Shape{Kind: "primitive", PrimitiveCode: "f32le"}, p.ServerTime); err != nil {
		return err
	}
	if err := w.Write("ServerStatsPacket.NetworkTime", Shape{Kind: "primitive", PrimitiveCode: "f32le"}, p.NetworkTime); err != nil {
		return err
	}
	return nil
}

func DecodeServerStats(r Decoder) (ServerStats, error) {
	var p ServerStats
	{
		raw, err := r.Read("ServerStatsPacket.ServerTime", Shape{Kind: "primitive", PrimitiveCode: "f32le"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(float32)
		if !ok {
			return p, fmt.Errorf("field ServerStatsPacket.ServerTime has unexpected decoded type %T", raw)
		}
		p.ServerTime = value
	}
	{
		raw, err := r.Read("ServerStatsPacket.NetworkTime", Shape{Kind: "primitive", PrimitiveCode: "f32le"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(float32)
		if !ok {
			return p, fmt.Errorf("field ServerStatsPacket.NetworkTime has unexpected decoded type %T", raw)
		}
		p.NetworkTime = value
	}
	return p, nil
}
