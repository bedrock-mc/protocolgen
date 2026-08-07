// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type HurtArmor struct {
	Cause      int32
	Damage     int32
	ArmorSlots uint64
}

func (p *HurtArmor) Encode(w Encoder) error {
	if err := w.Write("HurtArmorPacket.Cause", Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}, p.Cause); err != nil {
		return err
	}
	if err := w.Write("HurtArmorPacket.Damage", Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}, p.Damage); err != nil {
		return err
	}
	if err := w.Write("HurtArmorPacket.Armor Slots", Shape{Kind: "primitive", PrimitiveCode: "var_u64"}, p.ArmorSlots); err != nil {
		return err
	}
	return nil
}

func DecodeHurtArmor(r Decoder) (HurtArmor, error) {
	var p HurtArmor
	{
		raw, err := r.Read("HurtArmorPacket.Cause", Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(int32)
		if !ok {
			return p, fmt.Errorf("field HurtArmorPacket.Cause has unexpected decoded type %T", raw)
		}
		p.Cause = value
	}
	{
		raw, err := r.Read("HurtArmorPacket.Damage", Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(int32)
		if !ok {
			return p, fmt.Errorf("field HurtArmorPacket.Damage has unexpected decoded type %T", raw)
		}
		p.Damage = value
	}
	{
		raw, err := r.Read("HurtArmorPacket.Armor Slots", Shape{Kind: "primitive", PrimitiveCode: "var_u64"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(uint64)
		if !ok {
			return p, fmt.Errorf("field HurtArmorPacket.Armor Slots has unexpected decoded type %T", raw)
		}
		p.ArmorSlots = value
	}
	return p, nil
}
