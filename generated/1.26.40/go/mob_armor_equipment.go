// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type MobArmorEquipment struct {
	TargetRuntimeID ActorRuntimeID
	Head            CerealizerNetworkItemStackDescriptorSerializedData
	Torso           CerealizerNetworkItemStackDescriptorSerializedData
	Legs            CerealizerNetworkItemStackDescriptorSerializedData
	Feet            CerealizerNetworkItemStackDescriptorSerializedData
	Body            CerealizerNetworkItemStackDescriptorSerializedData
}

func (p *MobArmorEquipment) Encode(w Encoder) error {
	if err := w.Write("MobArmorEquipmentPacket.Target Runtime ID", Shape{Kind: "struct", Semantic: "ActorRuntimeID", TypeID: "ActorRuntimeID", Fields: []ShapeField{{Ordinal: 0, Name: "Actor Runtime ID", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u64"}}}}, p.TargetRuntimeID); err != nil {
		return err
	}
	if err := w.Write("MobArmorEquipmentPacket.Head", Shape{Kind: "struct", Semantic: "cerealizer<NetworkItemStackDescriptor>::SerializedData", TypeID: "cerealizer<NetworkItemStackDescriptor>::SerializedData", Fields: []ShapeField{{Ordinal: 0, Name: "Id", Shape: Shape{Kind: "primitive", PrimitiveCode: "i16le"}}, {Ordinal: 1, Name: "Stack size", Shape: Shape{Kind: "primitive", PrimitiveCode: "u16le"}}, {Ordinal: 2, Name: "Aux value", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, {Ordinal: 3, Name: "Net Id Variant", Shape: Shape{Kind: "optional", Value: &Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}}, {Ordinal: 4, Name: "Block Runtime Id", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, {Ordinal: 5, Name: "User Data Buffer", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}}, p.Head); err != nil {
		return err
	}
	if err := w.Write("MobArmorEquipmentPacket.Torso", Shape{Kind: "struct", Semantic: "cerealizer<NetworkItemStackDescriptor>::SerializedData", TypeID: "cerealizer<NetworkItemStackDescriptor>::SerializedData", Fields: []ShapeField{{Ordinal: 0, Name: "Id", Shape: Shape{Kind: "primitive", PrimitiveCode: "i16le"}}, {Ordinal: 1, Name: "Stack size", Shape: Shape{Kind: "primitive", PrimitiveCode: "u16le"}}, {Ordinal: 2, Name: "Aux value", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, {Ordinal: 3, Name: "Net Id Variant", Shape: Shape{Kind: "optional", Value: &Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}}, {Ordinal: 4, Name: "Block Runtime Id", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, {Ordinal: 5, Name: "User Data Buffer", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}}, p.Torso); err != nil {
		return err
	}
	if err := w.Write("MobArmorEquipmentPacket.Legs", Shape{Kind: "struct", Semantic: "cerealizer<NetworkItemStackDescriptor>::SerializedData", TypeID: "cerealizer<NetworkItemStackDescriptor>::SerializedData", Fields: []ShapeField{{Ordinal: 0, Name: "Id", Shape: Shape{Kind: "primitive", PrimitiveCode: "i16le"}}, {Ordinal: 1, Name: "Stack size", Shape: Shape{Kind: "primitive", PrimitiveCode: "u16le"}}, {Ordinal: 2, Name: "Aux value", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, {Ordinal: 3, Name: "Net Id Variant", Shape: Shape{Kind: "optional", Value: &Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}}, {Ordinal: 4, Name: "Block Runtime Id", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, {Ordinal: 5, Name: "User Data Buffer", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}}, p.Legs); err != nil {
		return err
	}
	if err := w.Write("MobArmorEquipmentPacket.Feet", Shape{Kind: "struct", Semantic: "cerealizer<NetworkItemStackDescriptor>::SerializedData", TypeID: "cerealizer<NetworkItemStackDescriptor>::SerializedData", Fields: []ShapeField{{Ordinal: 0, Name: "Id", Shape: Shape{Kind: "primitive", PrimitiveCode: "i16le"}}, {Ordinal: 1, Name: "Stack size", Shape: Shape{Kind: "primitive", PrimitiveCode: "u16le"}}, {Ordinal: 2, Name: "Aux value", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, {Ordinal: 3, Name: "Net Id Variant", Shape: Shape{Kind: "optional", Value: &Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}}, {Ordinal: 4, Name: "Block Runtime Id", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, {Ordinal: 5, Name: "User Data Buffer", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}}, p.Feet); err != nil {
		return err
	}
	if err := w.Write("MobArmorEquipmentPacket.Body", Shape{Kind: "struct", Semantic: "cerealizer<NetworkItemStackDescriptor>::SerializedData", TypeID: "cerealizer<NetworkItemStackDescriptor>::SerializedData", Fields: []ShapeField{{Ordinal: 0, Name: "Id", Shape: Shape{Kind: "primitive", PrimitiveCode: "i16le"}}, {Ordinal: 1, Name: "Stack size", Shape: Shape{Kind: "primitive", PrimitiveCode: "u16le"}}, {Ordinal: 2, Name: "Aux value", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, {Ordinal: 3, Name: "Net Id Variant", Shape: Shape{Kind: "optional", Value: &Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}}, {Ordinal: 4, Name: "Block Runtime Id", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, {Ordinal: 5, Name: "User Data Buffer", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}}, p.Body); err != nil {
		return err
	}
	return nil
}

func DecodeMobArmorEquipment(r Decoder) (MobArmorEquipment, error) {
	var p MobArmorEquipment
	{
		raw, err := r.Read("MobArmorEquipmentPacket.Target Runtime ID", Shape{Kind: "struct", Semantic: "ActorRuntimeID", TypeID: "ActorRuntimeID", Fields: []ShapeField{{Ordinal: 0, Name: "Actor Runtime ID", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u64"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(ActorRuntimeID)
		if !ok {
			return p, fmt.Errorf("field MobArmorEquipmentPacket.Target Runtime ID has unexpected decoded type %T", raw)
		}
		p.TargetRuntimeID = value
	}
	{
		raw, err := r.Read("MobArmorEquipmentPacket.Head", Shape{Kind: "struct", Semantic: "cerealizer<NetworkItemStackDescriptor>::SerializedData", TypeID: "cerealizer<NetworkItemStackDescriptor>::SerializedData", Fields: []ShapeField{{Ordinal: 0, Name: "Id", Shape: Shape{Kind: "primitive", PrimitiveCode: "i16le"}}, {Ordinal: 1, Name: "Stack size", Shape: Shape{Kind: "primitive", PrimitiveCode: "u16le"}}, {Ordinal: 2, Name: "Aux value", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, {Ordinal: 3, Name: "Net Id Variant", Shape: Shape{Kind: "optional", Value: &Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}}, {Ordinal: 4, Name: "Block Runtime Id", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, {Ordinal: 5, Name: "User Data Buffer", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(CerealizerNetworkItemStackDescriptorSerializedData)
		if !ok {
			return p, fmt.Errorf("field MobArmorEquipmentPacket.Head has unexpected decoded type %T", raw)
		}
		p.Head = value
	}
	{
		raw, err := r.Read("MobArmorEquipmentPacket.Torso", Shape{Kind: "struct", Semantic: "cerealizer<NetworkItemStackDescriptor>::SerializedData", TypeID: "cerealizer<NetworkItemStackDescriptor>::SerializedData", Fields: []ShapeField{{Ordinal: 0, Name: "Id", Shape: Shape{Kind: "primitive", PrimitiveCode: "i16le"}}, {Ordinal: 1, Name: "Stack size", Shape: Shape{Kind: "primitive", PrimitiveCode: "u16le"}}, {Ordinal: 2, Name: "Aux value", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, {Ordinal: 3, Name: "Net Id Variant", Shape: Shape{Kind: "optional", Value: &Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}}, {Ordinal: 4, Name: "Block Runtime Id", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, {Ordinal: 5, Name: "User Data Buffer", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(CerealizerNetworkItemStackDescriptorSerializedData)
		if !ok {
			return p, fmt.Errorf("field MobArmorEquipmentPacket.Torso has unexpected decoded type %T", raw)
		}
		p.Torso = value
	}
	{
		raw, err := r.Read("MobArmorEquipmentPacket.Legs", Shape{Kind: "struct", Semantic: "cerealizer<NetworkItemStackDescriptor>::SerializedData", TypeID: "cerealizer<NetworkItemStackDescriptor>::SerializedData", Fields: []ShapeField{{Ordinal: 0, Name: "Id", Shape: Shape{Kind: "primitive", PrimitiveCode: "i16le"}}, {Ordinal: 1, Name: "Stack size", Shape: Shape{Kind: "primitive", PrimitiveCode: "u16le"}}, {Ordinal: 2, Name: "Aux value", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, {Ordinal: 3, Name: "Net Id Variant", Shape: Shape{Kind: "optional", Value: &Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}}, {Ordinal: 4, Name: "Block Runtime Id", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, {Ordinal: 5, Name: "User Data Buffer", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(CerealizerNetworkItemStackDescriptorSerializedData)
		if !ok {
			return p, fmt.Errorf("field MobArmorEquipmentPacket.Legs has unexpected decoded type %T", raw)
		}
		p.Legs = value
	}
	{
		raw, err := r.Read("MobArmorEquipmentPacket.Feet", Shape{Kind: "struct", Semantic: "cerealizer<NetworkItemStackDescriptor>::SerializedData", TypeID: "cerealizer<NetworkItemStackDescriptor>::SerializedData", Fields: []ShapeField{{Ordinal: 0, Name: "Id", Shape: Shape{Kind: "primitive", PrimitiveCode: "i16le"}}, {Ordinal: 1, Name: "Stack size", Shape: Shape{Kind: "primitive", PrimitiveCode: "u16le"}}, {Ordinal: 2, Name: "Aux value", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, {Ordinal: 3, Name: "Net Id Variant", Shape: Shape{Kind: "optional", Value: &Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}}, {Ordinal: 4, Name: "Block Runtime Id", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, {Ordinal: 5, Name: "User Data Buffer", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(CerealizerNetworkItemStackDescriptorSerializedData)
		if !ok {
			return p, fmt.Errorf("field MobArmorEquipmentPacket.Feet has unexpected decoded type %T", raw)
		}
		p.Feet = value
	}
	{
		raw, err := r.Read("MobArmorEquipmentPacket.Body", Shape{Kind: "struct", Semantic: "cerealizer<NetworkItemStackDescriptor>::SerializedData", TypeID: "cerealizer<NetworkItemStackDescriptor>::SerializedData", Fields: []ShapeField{{Ordinal: 0, Name: "Id", Shape: Shape{Kind: "primitive", PrimitiveCode: "i16le"}}, {Ordinal: 1, Name: "Stack size", Shape: Shape{Kind: "primitive", PrimitiveCode: "u16le"}}, {Ordinal: 2, Name: "Aux value", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, {Ordinal: 3, Name: "Net Id Variant", Shape: Shape{Kind: "optional", Value: &Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}}, {Ordinal: 4, Name: "Block Runtime Id", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, {Ordinal: 5, Name: "User Data Buffer", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(CerealizerNetworkItemStackDescriptorSerializedData)
		if !ok {
			return p, fmt.Errorf("field MobArmorEquipmentPacket.Body has unexpected decoded type %T", raw)
		}
		p.Body = value
	}
	return p, nil
}
