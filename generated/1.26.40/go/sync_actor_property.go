// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type SyncActorProperty struct {
	PropertyData []byte
}

func (p *SyncActorProperty) Encode(w Encoder) error {
	if err := w.Write("SyncActorPropertyPacket.Property Data", Shape{Kind: "primitive", PrimitiveCode: "nbt_le"}, p.PropertyData); err != nil {
		return err
	}
	return nil
}

func DecodeSyncActorProperty(r Decoder) (SyncActorProperty, error) {
	var p SyncActorProperty
	{
		raw, err := r.Read("SyncActorPropertyPacket.Property Data", Shape{Kind: "primitive", PrimitiveCode: "nbt_le"})
		if err != nil {
			return p, err
		}
		value, ok := raw.([]byte)
		if !ok {
			return p, fmt.Errorf("field SyncActorPropertyPacket.Property Data has unexpected decoded type %T", raw)
		}
		p.PropertyData = value
	}
	return p, nil
}
