// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type SetHealth struct {
	Health int32
}

func (p *SetHealth) Encode(w Encoder) error {
	if err := w.Write("SetHealthPacket.Health", Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}, p.Health); err != nil {
		return err
	}
	return nil
}

func DecodeSetHealth(r Decoder) (SetHealth, error) {
	var p SetHealth
	{
		raw, err := r.Read("SetHealthPacket.Health", Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(int32)
		if !ok {
			return p, fmt.Errorf("field SetHealthPacket.Health has unexpected decoded type %T", raw)
		}
		p.Health = value
	}
	return p, nil
}
