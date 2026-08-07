// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type UpdateEquip struct {
	ContainerId    uint8
	Type           uint8
	Size           int32
	EntityUniqueId ActorUniqueID
	Data           []byte
}

func (p *UpdateEquip) Encode(w Encoder) error {
	if err := w.Write("UpdateEquipPacket.Container Id", Shape{Kind: "primitive", PrimitiveCode: "u8"}, p.ContainerId); err != nil {
		return err
	}
	if err := w.Write("UpdateEquipPacket.Type", Shape{Kind: "primitive", PrimitiveCode: "u8"}, p.Type); err != nil {
		return err
	}
	if err := w.Write("UpdateEquipPacket.Size", Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}, p.Size); err != nil {
		return err
	}
	if err := w.Write("UpdateEquipPacket.Entity Unique Id", Shape{Kind: "struct", Semantic: "ActorUniqueID", TypeID: "ActorUniqueID", Fields: []ShapeField{{Ordinal: 0, Name: "Actor Unique ID", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i64"}}}}, p.EntityUniqueId); err != nil {
		return err
	}
	if err := w.Write("UpdateEquipPacket.Data", Shape{Kind: "primitive", PrimitiveCode: "nbt_le"}, p.Data); err != nil {
		return err
	}
	return nil
}

func DecodeUpdateEquip(r Decoder) (UpdateEquip, error) {
	var p UpdateEquip
	{
		raw, err := r.Read("UpdateEquipPacket.Container Id", Shape{Kind: "primitive", PrimitiveCode: "u8"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(uint8)
		if !ok {
			return p, fmt.Errorf("field UpdateEquipPacket.Container Id has unexpected decoded type %T", raw)
		}
		p.ContainerId = value
	}
	{
		raw, err := r.Read("UpdateEquipPacket.Type", Shape{Kind: "primitive", PrimitiveCode: "u8"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(uint8)
		if !ok {
			return p, fmt.Errorf("field UpdateEquipPacket.Type has unexpected decoded type %T", raw)
		}
		p.Type = value
	}
	{
		raw, err := r.Read("UpdateEquipPacket.Size", Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(int32)
		if !ok {
			return p, fmt.Errorf("field UpdateEquipPacket.Size has unexpected decoded type %T", raw)
		}
		p.Size = value
	}
	{
		raw, err := r.Read("UpdateEquipPacket.Entity Unique Id", Shape{Kind: "struct", Semantic: "ActorUniqueID", TypeID: "ActorUniqueID", Fields: []ShapeField{{Ordinal: 0, Name: "Actor Unique ID", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i64"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(ActorUniqueID)
		if !ok {
			return p, fmt.Errorf("field UpdateEquipPacket.Entity Unique Id has unexpected decoded type %T", raw)
		}
		p.EntityUniqueId = value
	}
	{
		raw, err := r.Read("UpdateEquipPacket.Data", Shape{Kind: "primitive", PrimitiveCode: "nbt_le"})
		if err != nil {
			return p, err
		}
		value, ok := raw.([]byte)
		if !ok {
			return p, fmt.Errorf("field UpdateEquipPacket.Data has unexpected decoded type %T", raw)
		}
		p.Data = value
	}
	return p, nil
}
