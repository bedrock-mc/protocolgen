// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type ShowCredits struct {
	PlayerRuntimeID ActorRuntimeID
	CreditsState    int32
}

func (p *ShowCredits) Encode(w Encoder) error {
	if err := w.Write("ShowCreditsPacket.Player Runtime ID", Shape{Kind: "struct", Semantic: "ActorRuntimeID", TypeID: "ActorRuntimeID", Fields: []ShapeField{{Ordinal: 0, Name: "Actor Runtime ID", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u64"}}}}, p.PlayerRuntimeID); err != nil {
		return err
	}
	if err := w.Write("ShowCreditsPacket.Credits State", Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}, p.CreditsState); err != nil {
		return err
	}
	return nil
}

func DecodeShowCredits(r Decoder) (ShowCredits, error) {
	var p ShowCredits
	{
		raw, err := r.Read("ShowCreditsPacket.Player Runtime ID", Shape{Kind: "struct", Semantic: "ActorRuntimeID", TypeID: "ActorRuntimeID", Fields: []ShapeField{{Ordinal: 0, Name: "Actor Runtime ID", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u64"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(ActorRuntimeID)
		if !ok {
			return p, fmt.Errorf("field ShowCreditsPacket.Player Runtime ID has unexpected decoded type %T", raw)
		}
		p.PlayerRuntimeID = value
	}
	{
		raw, err := r.Read("ShowCreditsPacket.Credits State", Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(int32)
		if !ok {
			return p, fmt.Errorf("field ShowCreditsPacket.Credits State has unexpected decoded type %T", raw)
		}
		p.CreditsState = value
	}
	return p, nil
}
