// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type AutomationClientConnect struct {
	WebSocketData WebSocketPacketData
}

func (p *AutomationClientConnect) Encode(w Encoder) error {
	if err := w.Write("AutomationClientConnectPacket.Web Socket Data", Shape{Kind: "struct", Semantic: "WebSocketPacketData", TypeID: "WebSocketPacketData", Fields: []ShapeField{{Ordinal: 0, Name: "Websocket Server URI", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}}, p.WebSocketData); err != nil {
		return err
	}
	return nil
}

func DecodeAutomationClientConnect(r Decoder) (AutomationClientConnect, error) {
	var p AutomationClientConnect
	{
		raw, err := r.Read("AutomationClientConnectPacket.Web Socket Data", Shape{Kind: "struct", Semantic: "WebSocketPacketData", TypeID: "WebSocketPacketData", Fields: []ShapeField{{Ordinal: 0, Name: "Websocket Server URI", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(WebSocketPacketData)
		if !ok {
			return p, fmt.Errorf("field AutomationClientConnectPacket.Web Socket Data has unexpected decoded type %T", raw)
		}
		p.WebSocketData = value
	}
	return p, nil
}
