// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type SimulationType struct {
	SimType SimulationTypeType
}

func (p *SimulationType) Encode(w Encoder) error {
	if err := w.Write("SimulationTypePacket.Sim Type", Shape{Kind: "enum", Semantic: "SimulationType", TypeID: "enums/SimulationType", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "Game", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "Editor", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "Test", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "INVALID", Shape: Shape{Kind: "void"}}}}, p.SimType); err != nil {
		return err
	}
	return nil
}

func DecodeSimulationType(r Decoder) (SimulationType, error) {
	var p SimulationType
	{
		raw, err := r.Read("SimulationTypePacket.Sim Type", Shape{Kind: "enum", Semantic: "SimulationType", TypeID: "enums/SimulationType", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "Game", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "Editor", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "Test", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "INVALID", Shape: Shape{Kind: "void"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(SimulationTypeType)
		if !ok {
			return p, fmt.Errorf("field SimulationTypePacket.Sim Type has unexpected decoded type %T", raw)
		}
		p.SimType = value
	}
	return p, nil
}
