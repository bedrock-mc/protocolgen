// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type CurrentStructureFeature struct {
	CurrentStructureFeature string
}

func (p *CurrentStructureFeature) Encode(w Encoder) error {
	if err := w.Write("CurrentStructureFeaturePacket.Current Structure Feature", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, p.CurrentStructureFeature); err != nil {
		return err
	}
	return nil
}

func DecodeCurrentStructureFeature(r Decoder) (CurrentStructureFeature, error) {
	var p CurrentStructureFeature
	{
		raw, err := r.Read("CurrentStructureFeaturePacket.Current Structure Feature", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(string)
		if !ok {
			return p, fmt.Errorf("field CurrentStructureFeaturePacket.Current Structure Feature has unexpected decoded type %T", raw)
		}
		p.CurrentStructureFeature = value
	}
	return p, nil
}
