// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type ContainerSetData struct {
	ContainerID uint8
	ID          int32
	Value       int32
}

func (p *ContainerSetData) Encode(w Encoder) error {
	if err := w.Write("ContainerSetDataPacket.Container ID", Shape{Kind: "primitive", PrimitiveCode: "u8"}, p.ContainerID); err != nil {
		return err
	}
	if err := w.Write("ContainerSetDataPacket.ID", Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}, p.ID); err != nil {
		return err
	}
	if err := w.Write("ContainerSetDataPacket.Value", Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}, p.Value); err != nil {
		return err
	}
	return nil
}

func DecodeContainerSetData(r Decoder) (ContainerSetData, error) {
	var p ContainerSetData
	{
		raw, err := r.Read("ContainerSetDataPacket.Container ID", Shape{Kind: "primitive", PrimitiveCode: "u8"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(uint8)
		if !ok {
			return p, fmt.Errorf("field ContainerSetDataPacket.Container ID has unexpected decoded type %T", raw)
		}
		p.ContainerID = value
	}
	{
		raw, err := r.Read("ContainerSetDataPacket.ID", Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(int32)
		if !ok {
			return p, fmt.Errorf("field ContainerSetDataPacket.ID has unexpected decoded type %T", raw)
		}
		p.ID = value
	}
	{
		raw, err := r.Read("ContainerSetDataPacket.Value", Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(int32)
		if !ok {
			return p, fmt.Errorf("field ContainerSetDataPacket.Value has unexpected decoded type %T", raw)
		}
		p.Value = value
	}
	return p, nil
}
