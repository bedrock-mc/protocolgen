// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type ServerPresenceInfo struct {
	PresenceConfiguration *ServerConfigurationPresenceConfiguration
}

func (p *ServerPresenceInfo) Encode(w Encoder) error {
	if err := w.Write("ServerPresenceInfoPacket.presence_configuration", Shape{Kind: "optional", Value: &Shape{Kind: "struct", Semantic: "ServerConfiguration::PresenceConfiguration", TypeID: "ServerConfiguration::PresenceConfiguration", Fields: []ShapeField{{Ordinal: 0, Name: "richPresenceId", Shape: Shape{Kind: "optional", Value: &Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}}}}, p.PresenceConfiguration); err != nil {
		return err
	}
	return nil
}

func DecodeServerPresenceInfo(r Decoder) (ServerPresenceInfo, error) {
	var p ServerPresenceInfo
	{
		raw, err := r.Read("ServerPresenceInfoPacket.presence_configuration", Shape{Kind: "optional", Value: &Shape{Kind: "struct", Semantic: "ServerConfiguration::PresenceConfiguration", TypeID: "ServerConfiguration::PresenceConfiguration", Fields: []ShapeField{{Ordinal: 0, Name: "richPresenceId", Shape: Shape{Kind: "optional", Value: &Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(*ServerConfigurationPresenceConfiguration)
		if !ok {
			return p, fmt.Errorf("field ServerPresenceInfoPacket.presence_configuration has unexpected decoded type %T", raw)
		}
		p.PresenceConfiguration = value
	}
	return p, nil
}
