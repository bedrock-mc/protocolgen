// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type SetCommandsEnabled struct {
	CommandsEnabled bool
}

func (p *SetCommandsEnabled) Encode(w Encoder) error {
	if err := w.Write("SetCommandsEnabledPacket.Commands Enabled", Shape{Kind: "primitive", PrimitiveCode: "bool"}, p.CommandsEnabled); err != nil {
		return err
	}
	return nil
}

func DecodeSetCommandsEnabled(r Decoder) (SetCommandsEnabled, error) {
	var p SetCommandsEnabled
	{
		raw, err := r.Read("SetCommandsEnabledPacket.Commands Enabled", Shape{Kind: "primitive", PrimitiveCode: "bool"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(bool)
		if !ok {
			return p, fmt.Errorf("field SetCommandsEnabledPacket.Commands Enabled has unexpected decoded type %T", raw)
		}
		p.CommandsEnabled = value
	}
	return p, nil
}
