// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type RequestNetworkSettings struct {
	ClientNetworkVersion int32
}

func (p *RequestNetworkSettings) Encode(w Encoder) error {
	if err := w.Write("RequestNetworkSettingsPacket.ClientNetworkVersion", Shape{Kind: "primitive", PrimitiveCode: "i32be"}, p.ClientNetworkVersion); err != nil {
		return err
	}
	return nil
}

func DecodeRequestNetworkSettings(r Decoder) (RequestNetworkSettings, error) {
	var p RequestNetworkSettings
	{
		raw, err := r.Read("RequestNetworkSettingsPacket.ClientNetworkVersion", Shape{Kind: "primitive", PrimitiveCode: "i32be"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(int32)
		if !ok {
			return p, fmt.Errorf("field RequestNetworkSettingsPacket.ClientNetworkVersion has unexpected decoded type %T", raw)
		}
		p.ClientNetworkVersion = value
	}
	return p, nil
}
