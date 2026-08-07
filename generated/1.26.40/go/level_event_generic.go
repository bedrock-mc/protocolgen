// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type LevelEventGeneric struct {
	EventId int32
	CTD     []byte
}

func (p *LevelEventGeneric) Encode(w Encoder) error {
	if err := w.Write("LevelEventGenericPacket.Event Id", Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}, p.EventId); err != nil {
		return err
	}
	if err := w.Write("LevelEventGenericPacket.__[[CTD]]__", Shape{Kind: "primitive", PrimitiveCode: "nbt_le"}, p.CTD); err != nil {
		return err
	}
	return nil
}

func DecodeLevelEventGeneric(r Decoder) (LevelEventGeneric, error) {
	var p LevelEventGeneric
	{
		raw, err := r.Read("LevelEventGenericPacket.Event Id", Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(int32)
		if !ok {
			return p, fmt.Errorf("field LevelEventGenericPacket.Event Id has unexpected decoded type %T", raw)
		}
		p.EventId = value
	}
	{
		raw, err := r.Read("LevelEventGenericPacket.__[[CTD]]__", Shape{Kind: "primitive", PrimitiveCode: "nbt_le"})
		if err != nil {
			return p, err
		}
		value, ok := raw.([]byte)
		if !ok {
			return p, fmt.Errorf("field LevelEventGenericPacket.__[[CTD]]__ has unexpected decoded type %T", raw)
		}
		p.CTD = value
	}
	return p, nil
}
