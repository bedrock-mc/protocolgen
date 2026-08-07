// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type UpdateClientInputLocks struct {
	InputLockComponentData uint32
}

func (p *UpdateClientInputLocks) Encode(w Encoder) error {
	if err := w.Write("UpdateClientInputLocksPacket.Input Lock ComponentData", Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, p.InputLockComponentData); err != nil {
		return err
	}
	return nil
}

func DecodeUpdateClientInputLocks(r Decoder) (UpdateClientInputLocks, error) {
	var p UpdateClientInputLocks
	{
		raw, err := r.Read("UpdateClientInputLocksPacket.Input Lock ComponentData", Shape{Kind: "primitive", PrimitiveCode: "var_u32"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(uint32)
		if !ok {
			return p, fmt.Errorf("field UpdateClientInputLocksPacket.Input Lock ComponentData has unexpected decoded type %T", raw)
		}
		p.InputLockComponentData = value
	}
	return p, nil
}
