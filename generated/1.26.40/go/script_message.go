// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type ScriptMessage struct {
	MessageId    string
	MessageValue string
}

func (p *ScriptMessage) Encode(w Encoder) error {
	if err := w.Write("ScriptMessagePacket.Message Id", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, p.MessageId); err != nil {
		return err
	}
	if err := w.Write("ScriptMessagePacket.Message Value", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, p.MessageValue); err != nil {
		return err
	}
	return nil
}

func DecodeScriptMessage(r Decoder) (ScriptMessage, error) {
	var p ScriptMessage
	{
		raw, err := r.Read("ScriptMessagePacket.Message Id", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(string)
		if !ok {
			return p, fmt.Errorf("field ScriptMessagePacket.Message Id has unexpected decoded type %T", raw)
		}
		p.MessageId = value
	}
	{
		raw, err := r.Read("ScriptMessagePacket.Message Value", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(string)
		if !ok {
			return p, fmt.Errorf("field ScriptMessagePacket.Message Value has unexpected decoded type %T", raw)
		}
		p.MessageValue = value
	}
	return p, nil
}
