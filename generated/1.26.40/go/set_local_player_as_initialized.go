// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type SetLocalPlayerAsInitialized struct {
	PlayerID ActorRuntimeID
}

func (p *SetLocalPlayerAsInitialized) Encode(w Encoder) error {
	if err := w.Write("SetLocalPlayerAsInitializedPacket.Player ID", Shape{Kind: "struct", Semantic: "ActorRuntimeID", TypeID: "ActorRuntimeID", Fields: []ShapeField{{Ordinal: 0, Name: "Actor Runtime ID", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u64"}}}}, p.PlayerID); err != nil {
		return err
	}
	return nil
}

func DecodeSetLocalPlayerAsInitialized(r Decoder) (SetLocalPlayerAsInitialized, error) {
	var p SetLocalPlayerAsInitialized
	{
		raw, err := r.Read("SetLocalPlayerAsInitializedPacket.Player ID", Shape{Kind: "struct", Semantic: "ActorRuntimeID", TypeID: "ActorRuntimeID", Fields: []ShapeField{{Ordinal: 0, Name: "Actor Runtime ID", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u64"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(ActorRuntimeID)
		if !ok {
			return p, fmt.Errorf("field SetLocalPlayerAsInitializedPacket.Player ID has unexpected decoded type %T", raw)
		}
		p.PlayerID = value
	}
	return p, nil
}
