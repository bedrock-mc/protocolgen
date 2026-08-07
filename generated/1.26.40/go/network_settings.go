// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type NetworkSettings struct {
	CompressionThreshold    uint16
	CompressionAlgorithm    PacketCompressionAlgorithm
	ClientThrottleEnabled   bool
	ClientThrottleThreshold uint8
	ClientThrottleScalar    float32
}

func (p *NetworkSettings) Encode(w Encoder) error {
	if err := w.Write("NetworkSettingsPacket.Compression Threshold", Shape{Kind: "primitive", PrimitiveCode: "u16le"}, p.CompressionThreshold); err != nil {
		return err
	}
	if err := w.Write("NetworkSettingsPacket.CompressionAlgorithm", Shape{Kind: "enum", Semantic: "PacketCompressionAlgorithm", TypeID: "enums/PacketCompressionAlgorithm", PrimitiveCode: "u16le", Variants: []ShapeVariant{{Value: 0, Name: "ZLib", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "Snappy", Shape: Shape{Kind: "void"}}, {Value: 65535, Name: "None", Shape: Shape{Kind: "void"}}}}, p.CompressionAlgorithm); err != nil {
		return err
	}
	if err := w.Write("NetworkSettingsPacket.Client Throttle Enabled", Shape{Kind: "primitive", PrimitiveCode: "bool"}, p.ClientThrottleEnabled); err != nil {
		return err
	}
	if err := w.Write("NetworkSettingsPacket.Client Throttle Threshold", Shape{Kind: "primitive", PrimitiveCode: "u8"}, p.ClientThrottleThreshold); err != nil {
		return err
	}
	if err := w.Write("NetworkSettingsPacket.Client Throttle Scalar", Shape{Kind: "primitive", PrimitiveCode: "f32le"}, p.ClientThrottleScalar); err != nil {
		return err
	}
	return nil
}

func DecodeNetworkSettings(r Decoder) (NetworkSettings, error) {
	var p NetworkSettings
	{
		raw, err := r.Read("NetworkSettingsPacket.Compression Threshold", Shape{Kind: "primitive", PrimitiveCode: "u16le"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(uint16)
		if !ok {
			return p, fmt.Errorf("field NetworkSettingsPacket.Compression Threshold has unexpected decoded type %T", raw)
		}
		p.CompressionThreshold = value
	}
	{
		raw, err := r.Read("NetworkSettingsPacket.CompressionAlgorithm", Shape{Kind: "enum", Semantic: "PacketCompressionAlgorithm", TypeID: "enums/PacketCompressionAlgorithm", PrimitiveCode: "u16le", Variants: []ShapeVariant{{Value: 0, Name: "ZLib", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "Snappy", Shape: Shape{Kind: "void"}}, {Value: 65535, Name: "None", Shape: Shape{Kind: "void"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(PacketCompressionAlgorithm)
		if !ok {
			return p, fmt.Errorf("field NetworkSettingsPacket.CompressionAlgorithm has unexpected decoded type %T", raw)
		}
		p.CompressionAlgorithm = value
	}
	{
		raw, err := r.Read("NetworkSettingsPacket.Client Throttle Enabled", Shape{Kind: "primitive", PrimitiveCode: "bool"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(bool)
		if !ok {
			return p, fmt.Errorf("field NetworkSettingsPacket.Client Throttle Enabled has unexpected decoded type %T", raw)
		}
		p.ClientThrottleEnabled = value
	}
	{
		raw, err := r.Read("NetworkSettingsPacket.Client Throttle Threshold", Shape{Kind: "primitive", PrimitiveCode: "u8"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(uint8)
		if !ok {
			return p, fmt.Errorf("field NetworkSettingsPacket.Client Throttle Threshold has unexpected decoded type %T", raw)
		}
		p.ClientThrottleThreshold = value
	}
	{
		raw, err := r.Read("NetworkSettingsPacket.Client Throttle Scalar", Shape{Kind: "primitive", PrimitiveCode: "f32le"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(float32)
		if !ok {
			return p, fmt.Errorf("field NetworkSettingsPacket.Client Throttle Scalar has unexpected decoded type %T", raw)
		}
		p.ClientThrottleScalar = value
	}
	return p, nil
}
