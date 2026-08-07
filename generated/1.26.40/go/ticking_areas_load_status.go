// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type TickingAreasLoadStatus struct {
	WaitingForPreload bool
}

func (p *TickingAreasLoadStatus) Encode(w Encoder) error {
	if err := w.Write("TickingAreasLoadStatusPacket.Waiting For Preload", Shape{Kind: "primitive", PrimitiveCode: "bool"}, p.WaitingForPreload); err != nil {
		return err
	}
	return nil
}

func DecodeTickingAreasLoadStatus(r Decoder) (TickingAreasLoadStatus, error) {
	var p TickingAreasLoadStatus
	{
		raw, err := r.Read("TickingAreasLoadStatusPacket.Waiting For Preload", Shape{Kind: "primitive", PrimitiveCode: "bool"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(bool)
		if !ok {
			return p, fmt.Errorf("field TickingAreasLoadStatusPacket.Waiting For Preload has unexpected decoded type %T", raw)
		}
		p.WaitingForPreload = value
	}
	return p, nil
}
