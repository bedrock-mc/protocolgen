// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type OnScreenTextureAnimation struct {
	EffectId uint32
}

func (p *OnScreenTextureAnimation) Encode(w Encoder) error {
	if err := w.Write("OnScreenTextureAnimationPacket.Effect Id", Shape{Kind: "primitive", PrimitiveCode: "u32le"}, p.EffectId); err != nil {
		return err
	}
	return nil
}

func DecodeOnScreenTextureAnimation(r Decoder) (OnScreenTextureAnimation, error) {
	var p OnScreenTextureAnimation
	{
		raw, err := r.Read("OnScreenTextureAnimationPacket.Effect Id", Shape{Kind: "primitive", PrimitiveCode: "u32le"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(uint32)
		if !ok {
			return p, fmt.Errorf("field OnScreenTextureAnimationPacket.Effect Id has unexpected decoded type %T", raw)
		}
		p.EffectId = value
	}
	return p, nil
}
