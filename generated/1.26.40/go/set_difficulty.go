// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type SetDifficulty struct {
	Difficulty uint32
}

func (p *SetDifficulty) Encode(w Encoder) error {
	if err := w.Write("SetDifficultyPacket.Difficulty", Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, p.Difficulty); err != nil {
		return err
	}
	return nil
}

func DecodeSetDifficulty(r Decoder) (SetDifficulty, error) {
	var p SetDifficulty
	{
		raw, err := r.Read("SetDifficultyPacket.Difficulty", Shape{Kind: "primitive", PrimitiveCode: "var_u32"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(uint32)
		if !ok {
			return p, fmt.Errorf("field SetDifficultyPacket.Difficulty has unexpected decoded type %T", raw)
		}
		p.Difficulty = value
	}
	return p, nil
}
