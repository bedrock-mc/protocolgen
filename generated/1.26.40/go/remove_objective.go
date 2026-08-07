// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type RemoveObjective struct {
	ObjectiveName string
}

func (p *RemoveObjective) Encode(w Encoder) error {
	if err := w.Write("RemoveObjectivePacket.Objective Name", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, p.ObjectiveName); err != nil {
		return err
	}
	return nil
}

func DecodeRemoveObjective(r Decoder) (RemoveObjective, error) {
	var p RemoveObjective
	{
		raw, err := r.Read("RemoveObjectivePacket.Objective Name", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(string)
		if !ok {
			return p, fmt.Errorf("field RemoveObjectivePacket.Objective Name has unexpected decoded type %T", raw)
		}
		p.ObjectiveName = value
	}
	return p, nil
}
