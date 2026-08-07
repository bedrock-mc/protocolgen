// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type CodeBuilder struct {
	URL                   string
	ShouldOpenCodeBuilder bool
}

func (p *CodeBuilder) Encode(w Encoder) error {
	if err := w.Write("CodeBuilderPacket.URL", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, p.URL); err != nil {
		return err
	}
	if err := w.Write("CodeBuilderPacket.Should open code builder", Shape{Kind: "primitive", PrimitiveCode: "bool"}, p.ShouldOpenCodeBuilder); err != nil {
		return err
	}
	return nil
}

func DecodeCodeBuilder(r Decoder) (CodeBuilder, error) {
	var p CodeBuilder
	{
		raw, err := r.Read("CodeBuilderPacket.URL", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(string)
		if !ok {
			return p, fmt.Errorf("field CodeBuilderPacket.URL has unexpected decoded type %T", raw)
		}
		p.URL = value
	}
	{
		raw, err := r.Read("CodeBuilderPacket.Should open code builder", Shape{Kind: "primitive", PrimitiveCode: "bool"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(bool)
		if !ok {
			return p, fmt.Errorf("field CodeBuilderPacket.Should open code builder has unexpected decoded type %T", raw)
		}
		p.ShouldOpenCodeBuilder = value
	}
	return p, nil
}
