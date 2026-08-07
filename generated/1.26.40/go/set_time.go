// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type SetTime struct {
	Time int32
}

func (p *SetTime) Encode(w Encoder) error {
	if err := w.Write("SetTimePacket.Time", Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}, p.Time); err != nil {
		return err
	}
	return nil
}

func DecodeSetTime(r Decoder) (SetTime, error) {
	var p SetTime
	{
		raw, err := r.Read("SetTimePacket.Time", Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(int32)
		if !ok {
			return p, fmt.Errorf("field SetTimePacket.Time has unexpected decoded type %T", raw)
		}
		p.Time = value
	}
	return p, nil
}
