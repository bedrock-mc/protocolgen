// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type SettingsCommand struct {
	Command        string
	SuppressOutput bool
}

func (p *SettingsCommand) Encode(w Encoder) error {
	if err := w.Write("SettingsCommandPacket.Command", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, p.Command); err != nil {
		return err
	}
	if err := w.Write("SettingsCommandPacket.Suppress Output?", Shape{Kind: "primitive", PrimitiveCode: "bool"}, p.SuppressOutput); err != nil {
		return err
	}
	return nil
}

func DecodeSettingsCommand(r Decoder) (SettingsCommand, error) {
	var p SettingsCommand
	{
		raw, err := r.Read("SettingsCommandPacket.Command", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(string)
		if !ok {
			return p, fmt.Errorf("field SettingsCommandPacket.Command has unexpected decoded type %T", raw)
		}
		p.Command = value
	}
	{
		raw, err := r.Read("SettingsCommandPacket.Suppress Output?", Shape{Kind: "primitive", PrimitiveCode: "bool"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(bool)
		if !ok {
			return p, fmt.Errorf("field SettingsCommandPacket.Suppress Output? has unexpected decoded type %T", raw)
		}
		p.SuppressOutput = value
	}
	return p, nil
}
