// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type GameTestResults struct {
	Succeeded bool
	Error     string
	TestName  string
}

func (p *GameTestResults) Encode(w Encoder) error {
	if err := w.Write("GameTestResultsPacket.Succeeded", Shape{Kind: "primitive", PrimitiveCode: "bool"}, p.Succeeded); err != nil {
		return err
	}
	if err := w.Write("GameTestResultsPacket.Error", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, p.Error); err != nil {
		return err
	}
	if err := w.Write("GameTestResultsPacket.TestName", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, p.TestName); err != nil {
		return err
	}
	return nil
}

func DecodeGameTestResults(r Decoder) (GameTestResults, error) {
	var p GameTestResults
	{
		raw, err := r.Read("GameTestResultsPacket.Succeeded", Shape{Kind: "primitive", PrimitiveCode: "bool"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(bool)
		if !ok {
			return p, fmt.Errorf("field GameTestResultsPacket.Succeeded has unexpected decoded type %T", raw)
		}
		p.Succeeded = value
	}
	{
		raw, err := r.Read("GameTestResultsPacket.Error", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(string)
		if !ok {
			return p, fmt.Errorf("field GameTestResultsPacket.Error has unexpected decoded type %T", raw)
		}
		p.Error = value
	}
	{
		raw, err := r.Read("GameTestResultsPacket.TestName", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(string)
		if !ok {
			return p, fmt.Errorf("field GameTestResultsPacket.TestName has unexpected decoded type %T", raw)
		}
		p.TestName = value
	}
	return p, nil
}
