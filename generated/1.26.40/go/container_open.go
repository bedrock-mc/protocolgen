// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type ContainerOpen struct {
	ContainerId   uint8
	ContainerType uint8
	Position      BlockPos
	TargetActorID ActorUniqueID
}

func (p *ContainerOpen) Encode(w Encoder) error {
	if err := w.Write("ContainerOpenPacket.Container Id", Shape{Kind: "primitive", PrimitiveCode: "u8"}, p.ContainerId); err != nil {
		return err
	}
	if err := w.Write("ContainerOpenPacket.Container Type", Shape{Kind: "primitive", PrimitiveCode: "u8"}, p.ContainerType); err != nil {
		return err
	}
	if err := w.Write("ContainerOpenPacket.Position", Shape{Kind: "struct", Semantic: "BlockPos", TypeID: "BlockPos", Fields: []ShapeField{{Ordinal: 0, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 1, Name: "Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 2, Name: "Z", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}}}, p.Position); err != nil {
		return err
	}
	if err := w.Write("ContainerOpenPacket.Target Actor ID", Shape{Kind: "struct", Semantic: "ActorUniqueID", TypeID: "ActorUniqueID", Fields: []ShapeField{{Ordinal: 0, Name: "Actor Unique ID", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i64"}}}}, p.TargetActorID); err != nil {
		return err
	}
	return nil
}

func DecodeContainerOpen(r Decoder) (ContainerOpen, error) {
	var p ContainerOpen
	{
		raw, err := r.Read("ContainerOpenPacket.Container Id", Shape{Kind: "primitive", PrimitiveCode: "u8"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(uint8)
		if !ok {
			return p, fmt.Errorf("field ContainerOpenPacket.Container Id has unexpected decoded type %T", raw)
		}
		p.ContainerId = value
	}
	{
		raw, err := r.Read("ContainerOpenPacket.Container Type", Shape{Kind: "primitive", PrimitiveCode: "u8"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(uint8)
		if !ok {
			return p, fmt.Errorf("field ContainerOpenPacket.Container Type has unexpected decoded type %T", raw)
		}
		p.ContainerType = value
	}
	{
		raw, err := r.Read("ContainerOpenPacket.Position", Shape{Kind: "struct", Semantic: "BlockPos", TypeID: "BlockPos", Fields: []ShapeField{{Ordinal: 0, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 1, Name: "Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 2, Name: "Z", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(BlockPos)
		if !ok {
			return p, fmt.Errorf("field ContainerOpenPacket.Position has unexpected decoded type %T", raw)
		}
		p.Position = value
	}
	{
		raw, err := r.Read("ContainerOpenPacket.Target Actor ID", Shape{Kind: "struct", Semantic: "ActorUniqueID", TypeID: "ActorUniqueID", Fields: []ShapeField{{Ordinal: 0, Name: "Actor Unique ID", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i64"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(ActorUniqueID)
		if !ok {
			return p, fmt.Errorf("field ContainerOpenPacket.Target Actor ID has unexpected decoded type %T", raw)
		}
		p.TargetActorID = value
	}
	return p, nil
}
