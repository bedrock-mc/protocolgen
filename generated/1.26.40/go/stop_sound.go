// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type StopSound struct {
	SoundName       string
	StopAllSounds   bool
	StopMusicLegacy bool
}

func (p *StopSound) Encode(w Encoder) error {
	if err := w.Write("StopSoundPacket.Sound Name", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, p.SoundName); err != nil {
		return err
	}
	if err := w.Write("StopSoundPacket.Stop All Sounds?", Shape{Kind: "primitive", PrimitiveCode: "bool"}, p.StopAllSounds); err != nil {
		return err
	}
	if err := w.Write("StopSoundPacket.Stop Music (Legacy)", Shape{Kind: "primitive", PrimitiveCode: "bool"}, p.StopMusicLegacy); err != nil {
		return err
	}
	return nil
}

func DecodeStopSound(r Decoder) (StopSound, error) {
	var p StopSound
	{
		raw, err := r.Read("StopSoundPacket.Sound Name", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(string)
		if !ok {
			return p, fmt.Errorf("field StopSoundPacket.Sound Name has unexpected decoded type %T", raw)
		}
		p.SoundName = value
	}
	{
		raw, err := r.Read("StopSoundPacket.Stop All Sounds?", Shape{Kind: "primitive", PrimitiveCode: "bool"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(bool)
		if !ok {
			return p, fmt.Errorf("field StopSoundPacket.Stop All Sounds? has unexpected decoded type %T", raw)
		}
		p.StopAllSounds = value
	}
	{
		raw, err := r.Read("StopSoundPacket.Stop Music (Legacy)", Shape{Kind: "primitive", PrimitiveCode: "bool"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(bool)
		if !ok {
			return p, fmt.Errorf("field StopSoundPacket.Stop Music (Legacy) has unexpected decoded type %T", raw)
		}
		p.StopMusicLegacy = value
	}
	return p, nil
}
