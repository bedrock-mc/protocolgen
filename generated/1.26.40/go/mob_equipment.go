// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type MobEquipment struct {
	TargetRuntimeID ActorRuntimeID
	Item            CerealizerNetworkItemStackDescriptorSerializedData
	Slot            uint8
	SelectedSlot    uint8
	ContainerID     uint8
}

func (p *MobEquipment) Encode(w Encoder) error {
	if err := w.Write("MobEquipmentPacket.Target Runtime ID", Shape{Kind: "struct", Semantic: "ActorRuntimeID", TypeID: "ActorRuntimeID", Fields: []ShapeField{{Ordinal: 0, Name: "Actor Runtime ID", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u64"}}}}, p.TargetRuntimeID); err != nil {
		return err
	}
	if err := w.Write("MobEquipmentPacket.Item", Shape{Kind: "struct", Semantic: "cerealizer<NetworkItemStackDescriptor>::SerializedData", TypeID: "cerealizer<NetworkItemStackDescriptor>::SerializedData", Fields: []ShapeField{{Ordinal: 0, Name: "Id", Shape: Shape{Kind: "primitive", PrimitiveCode: "i16le"}}, {Ordinal: 1, Name: "Stack size", Shape: Shape{Kind: "primitive", PrimitiveCode: "u16le"}}, {Ordinal: 2, Name: "Aux value", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, {Ordinal: 3, Name: "Net Id Variant", Shape: Shape{Kind: "optional", Value: &Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}}, {Ordinal: 4, Name: "Block Runtime Id", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, {Ordinal: 5, Name: "User Data Buffer", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}}, p.Item); err != nil {
		return err
	}
	if err := w.Write("MobEquipmentPacket.Slot", Shape{Kind: "primitive", PrimitiveCode: "u8"}, p.Slot); err != nil {
		return err
	}
	if err := w.Write("MobEquipmentPacket.Selected Slot", Shape{Kind: "primitive", PrimitiveCode: "u8"}, p.SelectedSlot); err != nil {
		return err
	}
	if err := w.Write("MobEquipmentPacket.Container ID", Shape{Kind: "primitive", PrimitiveCode: "u8"}, p.ContainerID); err != nil {
		return err
	}
	return nil
}

func DecodeMobEquipment(r Decoder) (MobEquipment, error) {
	var p MobEquipment
	{
		raw, err := r.Read("MobEquipmentPacket.Target Runtime ID", Shape{Kind: "struct", Semantic: "ActorRuntimeID", TypeID: "ActorRuntimeID", Fields: []ShapeField{{Ordinal: 0, Name: "Actor Runtime ID", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u64"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(ActorRuntimeID)
		if !ok {
			return p, fmt.Errorf("field MobEquipmentPacket.Target Runtime ID has unexpected decoded type %T", raw)
		}
		p.TargetRuntimeID = value
	}
	{
		raw, err := r.Read("MobEquipmentPacket.Item", Shape{Kind: "struct", Semantic: "cerealizer<NetworkItemStackDescriptor>::SerializedData", TypeID: "cerealizer<NetworkItemStackDescriptor>::SerializedData", Fields: []ShapeField{{Ordinal: 0, Name: "Id", Shape: Shape{Kind: "primitive", PrimitiveCode: "i16le"}}, {Ordinal: 1, Name: "Stack size", Shape: Shape{Kind: "primitive", PrimitiveCode: "u16le"}}, {Ordinal: 2, Name: "Aux value", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, {Ordinal: 3, Name: "Net Id Variant", Shape: Shape{Kind: "optional", Value: &Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}}, {Ordinal: 4, Name: "Block Runtime Id", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, {Ordinal: 5, Name: "User Data Buffer", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(CerealizerNetworkItemStackDescriptorSerializedData)
		if !ok {
			return p, fmt.Errorf("field MobEquipmentPacket.Item has unexpected decoded type %T", raw)
		}
		p.Item = value
	}
	{
		raw, err := r.Read("MobEquipmentPacket.Slot", Shape{Kind: "primitive", PrimitiveCode: "u8"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(uint8)
		if !ok {
			return p, fmt.Errorf("field MobEquipmentPacket.Slot has unexpected decoded type %T", raw)
		}
		p.Slot = value
	}
	{
		raw, err := r.Read("MobEquipmentPacket.Selected Slot", Shape{Kind: "primitive", PrimitiveCode: "u8"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(uint8)
		if !ok {
			return p, fmt.Errorf("field MobEquipmentPacket.Selected Slot has unexpected decoded type %T", raw)
		}
		p.SelectedSlot = value
	}
	{
		raw, err := r.Read("MobEquipmentPacket.Container ID", Shape{Kind: "primitive", PrimitiveCode: "u8"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(uint8)
		if !ok {
			return p, fmt.Errorf("field MobEquipmentPacket.Container ID has unexpected decoded type %T", raw)
		}
		p.ContainerID = value
	}
	return p, nil
}
