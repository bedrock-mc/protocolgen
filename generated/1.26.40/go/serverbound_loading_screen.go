// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type ServerboundLoadingScreen struct {
	LoadingScreenPacketType ServerboundLoadingScreenPacketType
	LoadingScreenId         *uint32
}

func (p *ServerboundLoadingScreen) Encode(w Encoder) error {
	if err := w.Write("ServerboundLoadingScreenPacket.Loading Screen Packet Type", Shape{Kind: "enum", Semantic: "ServerboundLoadingScreenPacketType", TypeID: "enums/ServerboundLoadingScreenPacketType", PrimitiveCode: "zigzag_i32", Variants: []ShapeVariant{{Value: 1, Name: "StartLoadingScreen", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "EndLoadingScreen", Shape: Shape{Kind: "void"}}}}, p.LoadingScreenPacketType); err != nil {
		return err
	}
	if err := w.Write("ServerboundLoadingScreenPacket.Loading Screen Id", Shape{Kind: "optional", Value: &Shape{Kind: "primitive", PrimitiveCode: "u32le"}}, p.LoadingScreenId); err != nil {
		return err
	}
	return nil
}

func DecodeServerboundLoadingScreen(r Decoder) (ServerboundLoadingScreen, error) {
	var p ServerboundLoadingScreen
	{
		raw, err := r.Read("ServerboundLoadingScreenPacket.Loading Screen Packet Type", Shape{Kind: "enum", Semantic: "ServerboundLoadingScreenPacketType", TypeID: "enums/ServerboundLoadingScreenPacketType", PrimitiveCode: "zigzag_i32", Variants: []ShapeVariant{{Value: 1, Name: "StartLoadingScreen", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "EndLoadingScreen", Shape: Shape{Kind: "void"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(ServerboundLoadingScreenPacketType)
		if !ok {
			return p, fmt.Errorf("field ServerboundLoadingScreenPacket.Loading Screen Packet Type has unexpected decoded type %T", raw)
		}
		p.LoadingScreenPacketType = value
	}
	{
		raw, err := r.Read("ServerboundLoadingScreenPacket.Loading Screen Id", Shape{Kind: "optional", Value: &Shape{Kind: "primitive", PrimitiveCode: "u32le"}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(*uint32)
		if !ok {
			return p, fmt.Errorf("field ServerboundLoadingScreenPacket.Loading Screen Id has unexpected decoded type %T", raw)
		}
		p.LoadingScreenId = value
	}
	return p, nil
}
