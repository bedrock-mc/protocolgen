// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type AvailableActorIdentifiers struct {
	IdentifierList []byte
}

func (p *AvailableActorIdentifiers) Encode(w Encoder) error {
	if err := w.Write("AvailableActorIdentifiersPacket.Identifier List", Shape{Kind: "primitive", PrimitiveCode: "nbt_le"}, p.IdentifierList); err != nil {
		return err
	}
	return nil
}

func DecodeAvailableActorIdentifiers(r Decoder) (AvailableActorIdentifiers, error) {
	var p AvailableActorIdentifiers
	{
		raw, err := r.Read("AvailableActorIdentifiersPacket.Identifier List", Shape{Kind: "primitive", PrimitiveCode: "nbt_le"})
		if err != nil {
			return p, err
		}
		value, ok := raw.([]byte)
		if !ok {
			return p, fmt.Errorf("field AvailableActorIdentifiersPacket.Identifier List has unexpected decoded type %T", raw)
		}
		p.IdentifierList = value
	}
	return p, nil
}
