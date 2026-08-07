// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type AddActor struct {
	TargetActorID     ActorUniqueID
	TargetRuntimeID   ActorRuntimeID
	ActorType         string
	Position          Vec3
	Velocity          Vec3
	Rotation          Vec2
	YHeadRotation     float32
	YBodyRotation     float32
	AttributesList    []SyncedAttribute
	ActorData         SynchedActorDataCopyableDataList
	SynchedProperties PropertySyncData
	ActorLinks        []ActorLink
}

func (p *AddActor) Encode(w Encoder) error {
	if err := w.Write("AddActorPacket.Target Actor ID", Shape{Kind: "struct", Semantic: "ActorUniqueID", TypeID: "ActorUniqueID", Fields: []ShapeField{{Ordinal: 0, Name: "Actor Unique ID", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i64"}}}}, p.TargetActorID); err != nil {
		return err
	}
	if err := w.Write("AddActorPacket.Target Runtime ID", Shape{Kind: "struct", Semantic: "ActorRuntimeID", TypeID: "ActorRuntimeID", Fields: []ShapeField{{Ordinal: 0, Name: "Actor Runtime ID", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u64"}}}}, p.TargetRuntimeID); err != nil {
		return err
	}
	if err := w.Write("AddActorPacket.Actor Type", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, p.ActorType); err != nil {
		return err
	}
	if err := w.Write("AddActorPacket.Position", Shape{Kind: "struct", Semantic: "Vec3", TypeID: "Vec3", Fields: []ShapeField{{Ordinal: 0, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 1, Name: "Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 2, Name: "Z", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}}}, p.Position); err != nil {
		return err
	}
	if err := w.Write("AddActorPacket.Velocity", Shape{Kind: "struct", Semantic: "Vec3", TypeID: "Vec3", Fields: []ShapeField{{Ordinal: 0, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 1, Name: "Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 2, Name: "Z", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}}}, p.Velocity); err != nil {
		return err
	}
	if err := w.Write("AddActorPacket.Rotation", Shape{Kind: "struct", Semantic: "Vec2", TypeID: "Vec2", Fields: []ShapeField{{Ordinal: 0, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 1, Name: "Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}}}, p.Rotation); err != nil {
		return err
	}
	if err := w.Write("AddActorPacket.Y Head Rotation", Shape{Kind: "primitive", PrimitiveCode: "f32le"}, p.YHeadRotation); err != nil {
		return err
	}
	if err := w.Write("AddActorPacket.Y Body Rotation", Shape{Kind: "primitive", PrimitiveCode: "f32le"}, p.YBodyRotation); err != nil {
		return err
	}
	if err := w.Write("AddActorPacket.Attributes List", Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "struct", Semantic: "SyncedAttribute", TypeID: "SyncedAttribute", Fields: []ShapeField{{Ordinal: 0, Name: "Attribute Name", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 1, Name: "Min Value", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 2, Name: "Current Value", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 3, Name: "Max Value", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}}}}, p.AttributesList); err != nil {
		return err
	}
	if err := w.Write("AddActorPacket.Actor Data", Shape{Kind: "struct", Semantic: "SynchedActorData::CopyableDataList", TypeID: "SynchedActorData::CopyableDataList", Fields: []ShapeField{{Ordinal: 0, Name: "Data", Shape: Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "struct", Semantic: "DataItemEntry", TypeID: "DataItemEntry", Fields: []ShapeField{{Ordinal: 0, Name: "ID", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, {Ordinal: 1, Name: "Payload", Shape: Shape{Kind: "union", Control: &Shape{Kind: "primitive", PrimitiveCode: "u8"}, Variants: []ShapeVariant{{Value: 0, Name: "DataItemBytePayload", Shape: Shape{Kind: "struct", Semantic: "DataItemBytePayload", TypeID: "DataItemBytePayload", Fields: []ShapeField{{Ordinal: 0, Name: "Type", Shape: Shape{Kind: "enum", Semantic: "DataItemType", TypeID: "enums/DataItemType", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "Byte", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "Short", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "Int", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "Float", Shape: Shape{Kind: "void"}}, {Value: 4, Name: "String", Shape: Shape{Kind: "void"}}, {Value: 5, Name: "CompoundTag", Shape: Shape{Kind: "void"}}, {Value: 6, Name: "Pos", Shape: Shape{Kind: "void"}}, {Value: 7, Name: "Int64", Shape: Shape{Kind: "void"}}, {Value: 8, Name: "Vec3", Shape: Shape{Kind: "void"}}}}}, {Ordinal: 1, Name: "Value", Shape: Shape{Kind: "primitive", PrimitiveCode: "i8"}}}}}, {Value: 1, Name: "DataItemShortPayload", Shape: Shape{Kind: "struct", Semantic: "DataItemShortPayload", TypeID: "DataItemShortPayload", Fields: []ShapeField{{Ordinal: 0, Name: "Type", Shape: Shape{Kind: "enum", Semantic: "DataItemType", TypeID: "enums/DataItemType", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "Byte", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "Short", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "Int", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "Float", Shape: Shape{Kind: "void"}}, {Value: 4, Name: "String", Shape: Shape{Kind: "void"}}, {Value: 5, Name: "CompoundTag", Shape: Shape{Kind: "void"}}, {Value: 6, Name: "Pos", Shape: Shape{Kind: "void"}}, {Value: 7, Name: "Int64", Shape: Shape{Kind: "void"}}, {Value: 8, Name: "Vec3", Shape: Shape{Kind: "void"}}}}}, {Ordinal: 1, Name: "Value", Shape: Shape{Kind: "primitive", PrimitiveCode: "i16le"}}}}}, {Value: 2, Name: "DataItemIntPayload", Shape: Shape{Kind: "struct", Semantic: "DataItemIntPayload", TypeID: "DataItemIntPayload", Fields: []ShapeField{{Ordinal: 0, Name: "Type", Shape: Shape{Kind: "enum", Semantic: "DataItemType", TypeID: "enums/DataItemType", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "Byte", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "Short", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "Int", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "Float", Shape: Shape{Kind: "void"}}, {Value: 4, Name: "String", Shape: Shape{Kind: "void"}}, {Value: 5, Name: "CompoundTag", Shape: Shape{Kind: "void"}}, {Value: 6, Name: "Pos", Shape: Shape{Kind: "void"}}, {Value: 7, Name: "Int64", Shape: Shape{Kind: "void"}}, {Value: 8, Name: "Vec3", Shape: Shape{Kind: "void"}}}}}, {Ordinal: 1, Name: "Value", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}}}}, {Value: 3, Name: "DataItemFloatPayload", Shape: Shape{Kind: "struct", Semantic: "DataItemFloatPayload", TypeID: "DataItemFloatPayload", Fields: []ShapeField{{Ordinal: 0, Name: "Type", Shape: Shape{Kind: "enum", Semantic: "DataItemType", TypeID: "enums/DataItemType", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "Byte", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "Short", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "Int", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "Float", Shape: Shape{Kind: "void"}}, {Value: 4, Name: "String", Shape: Shape{Kind: "void"}}, {Value: 5, Name: "CompoundTag", Shape: Shape{Kind: "void"}}, {Value: 6, Name: "Pos", Shape: Shape{Kind: "void"}}, {Value: 7, Name: "Int64", Shape: Shape{Kind: "void"}}, {Value: 8, Name: "Vec3", Shape: Shape{Kind: "void"}}}}}, {Ordinal: 1, Name: "Value", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}}}}, {Value: 4, Name: "DataItemStringPayload", Shape: Shape{Kind: "struct", Semantic: "DataItemStringPayload", TypeID: "DataItemStringPayload", Fields: []ShapeField{{Ordinal: 0, Name: "Type", Shape: Shape{Kind: "enum", Semantic: "DataItemType", TypeID: "enums/DataItemType", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "Byte", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "Short", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "Int", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "Float", Shape: Shape{Kind: "void"}}, {Value: 4, Name: "String", Shape: Shape{Kind: "void"}}, {Value: 5, Name: "CompoundTag", Shape: Shape{Kind: "void"}}, {Value: 6, Name: "Pos", Shape: Shape{Kind: "void"}}, {Value: 7, Name: "Int64", Shape: Shape{Kind: "void"}}, {Value: 8, Name: "Vec3", Shape: Shape{Kind: "void"}}}}}, {Ordinal: 1, Name: "Value", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}}}, {Value: 5, Name: "DataItemCompoundTagPayload", Shape: Shape{Kind: "struct", Semantic: "DataItemCompoundTagPayload", TypeID: "DataItemCompoundTagPayload", Fields: []ShapeField{{Ordinal: 0, Name: "Type", Shape: Shape{Kind: "enum", Semantic: "DataItemType", TypeID: "enums/DataItemType", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "Byte", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "Short", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "Int", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "Float", Shape: Shape{Kind: "void"}}, {Value: 4, Name: "String", Shape: Shape{Kind: "void"}}, {Value: 5, Name: "CompoundTag", Shape: Shape{Kind: "void"}}, {Value: 6, Name: "Pos", Shape: Shape{Kind: "void"}}, {Value: 7, Name: "Int64", Shape: Shape{Kind: "void"}}, {Value: 8, Name: "Vec3", Shape: Shape{Kind: "void"}}}}}, {Ordinal: 1, Name: "Value", Shape: Shape{Kind: "primitive", PrimitiveCode: "nbt_le"}}}}}, {Value: 6, Name: "DataItemPosPayload", Shape: Shape{Kind: "struct", Semantic: "DataItemPosPayload", TypeID: "DataItemPosPayload", Fields: []ShapeField{{Ordinal: 0, Name: "Type", Shape: Shape{Kind: "enum", Semantic: "DataItemType", TypeID: "enums/DataItemType", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "Byte", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "Short", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "Int", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "Float", Shape: Shape{Kind: "void"}}, {Value: 4, Name: "String", Shape: Shape{Kind: "void"}}, {Value: 5, Name: "CompoundTag", Shape: Shape{Kind: "void"}}, {Value: 6, Name: "Pos", Shape: Shape{Kind: "void"}}, {Value: 7, Name: "Int64", Shape: Shape{Kind: "void"}}, {Value: 8, Name: "Vec3", Shape: Shape{Kind: "void"}}}}}, {Ordinal: 1, Name: "Value", Shape: Shape{Kind: "struct", Semantic: "BlockPos", TypeID: "BlockPos", Fields: []ShapeField{{Ordinal: 0, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 1, Name: "Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 2, Name: "Z", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}}}}}}}, {Value: 7, Name: "DataItemInt64Payload", Shape: Shape{Kind: "struct", Semantic: "DataItemInt64Payload", TypeID: "DataItemInt64Payload", Fields: []ShapeField{{Ordinal: 0, Name: "Type", Shape: Shape{Kind: "enum", Semantic: "DataItemType", TypeID: "enums/DataItemType", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "Byte", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "Short", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "Int", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "Float", Shape: Shape{Kind: "void"}}, {Value: 4, Name: "String", Shape: Shape{Kind: "void"}}, {Value: 5, Name: "CompoundTag", Shape: Shape{Kind: "void"}}, {Value: 6, Name: "Pos", Shape: Shape{Kind: "void"}}, {Value: 7, Name: "Int64", Shape: Shape{Kind: "void"}}, {Value: 8, Name: "Vec3", Shape: Shape{Kind: "void"}}}}}, {Ordinal: 1, Name: "Value", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i64"}}}}}, {Value: 8, Name: "DataItemVec3Payload", Shape: Shape{Kind: "struct", Semantic: "DataItemVec3Payload", TypeID: "DataItemVec3Payload", Fields: []ShapeField{{Ordinal: 0, Name: "Type", Shape: Shape{Kind: "enum", Semantic: "DataItemType", TypeID: "enums/DataItemType", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "Byte", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "Short", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "Int", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "Float", Shape: Shape{Kind: "void"}}, {Value: 4, Name: "String", Shape: Shape{Kind: "void"}}, {Value: 5, Name: "CompoundTag", Shape: Shape{Kind: "void"}}, {Value: 6, Name: "Pos", Shape: Shape{Kind: "void"}}, {Value: 7, Name: "Int64", Shape: Shape{Kind: "void"}}, {Value: 8, Name: "Vec3", Shape: Shape{Kind: "void"}}}}}, {Ordinal: 1, Name: "Value", Shape: Shape{Kind: "struct", Semantic: "Vec3", TypeID: "Vec3", Fields: []ShapeField{{Ordinal: 0, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 1, Name: "Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 2, Name: "Z", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}}}}}}}}}}}}}}}}, p.ActorData); err != nil {
		return err
	}
	if err := w.Write("AddActorPacket.Synched Properties", Shape{Kind: "struct", Semantic: "PropertySyncData", TypeID: "PropertySyncData", Fields: []ShapeField{{Ordinal: 0, Name: "Int Entries List", Shape: Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "struct", Semantic: "PropertySyncData::PropertySyncIntEntry", TypeID: "PropertySyncData::PropertySyncIntEntry", Fields: []ShapeField{{Ordinal: 0, Name: "Property Index", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, {Ordinal: 1, Name: "Data", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}}}}}, {Ordinal: 1, Name: "Float Entries List", Shape: Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "struct", Semantic: "PropertySyncData::PropertySyncFloatEntry", TypeID: "PropertySyncData::PropertySyncFloatEntry", Fields: []ShapeField{{Ordinal: 0, Name: "Property Index", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, {Ordinal: 1, Name: "Data", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}}}}}}}, p.SynchedProperties); err != nil {
		return err
	}
	if err := w.Write("AddActorPacket.Actor Links", Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "struct", Semantic: "ActorLink", TypeID: "ActorLink", Fields: []ShapeField{{Ordinal: 0, Name: "Target A", Shape: Shape{Kind: "struct", Semantic: "ActorUniqueID", TypeID: "ActorUniqueID", Fields: []ShapeField{{Ordinal: 0, Name: "Actor Unique ID", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i64"}}}}}, {Ordinal: 1, Name: "Target B", Shape: Shape{Kind: "struct", Semantic: "ActorUniqueID", TypeID: "ActorUniqueID", Fields: []ShapeField{{Ordinal: 0, Name: "Actor Unique ID", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i64"}}}}}, {Ordinal: 2, Name: "Type", Shape: Shape{Kind: "enum", Semantic: "ActorLinkType", TypeID: "enums/ActorLinkType", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "None", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "Riding", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "Passenger", Shape: Shape{Kind: "void"}}}}}, {Ordinal: 3, Name: "Immediate", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}, {Ordinal: 4, Name: "Passenger Initiated", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}, {Ordinal: 5, Name: "Vehicle Angular Velocity", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}}}}, p.ActorLinks); err != nil {
		return err
	}
	return nil
}

func DecodeAddActor(r Decoder) (AddActor, error) {
	var p AddActor
	{
		raw, err := r.Read("AddActorPacket.Target Actor ID", Shape{Kind: "struct", Semantic: "ActorUniqueID", TypeID: "ActorUniqueID", Fields: []ShapeField{{Ordinal: 0, Name: "Actor Unique ID", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i64"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(ActorUniqueID)
		if !ok {
			return p, fmt.Errorf("field AddActorPacket.Target Actor ID has unexpected decoded type %T", raw)
		}
		p.TargetActorID = value
	}
	{
		raw, err := r.Read("AddActorPacket.Target Runtime ID", Shape{Kind: "struct", Semantic: "ActorRuntimeID", TypeID: "ActorRuntimeID", Fields: []ShapeField{{Ordinal: 0, Name: "Actor Runtime ID", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u64"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(ActorRuntimeID)
		if !ok {
			return p, fmt.Errorf("field AddActorPacket.Target Runtime ID has unexpected decoded type %T", raw)
		}
		p.TargetRuntimeID = value
	}
	{
		raw, err := r.Read("AddActorPacket.Actor Type", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(string)
		if !ok {
			return p, fmt.Errorf("field AddActorPacket.Actor Type has unexpected decoded type %T", raw)
		}
		p.ActorType = value
	}
	{
		raw, err := r.Read("AddActorPacket.Position", Shape{Kind: "struct", Semantic: "Vec3", TypeID: "Vec3", Fields: []ShapeField{{Ordinal: 0, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 1, Name: "Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 2, Name: "Z", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(Vec3)
		if !ok {
			return p, fmt.Errorf("field AddActorPacket.Position has unexpected decoded type %T", raw)
		}
		p.Position = value
	}
	{
		raw, err := r.Read("AddActorPacket.Velocity", Shape{Kind: "struct", Semantic: "Vec3", TypeID: "Vec3", Fields: []ShapeField{{Ordinal: 0, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 1, Name: "Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 2, Name: "Z", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(Vec3)
		if !ok {
			return p, fmt.Errorf("field AddActorPacket.Velocity has unexpected decoded type %T", raw)
		}
		p.Velocity = value
	}
	{
		raw, err := r.Read("AddActorPacket.Rotation", Shape{Kind: "struct", Semantic: "Vec2", TypeID: "Vec2", Fields: []ShapeField{{Ordinal: 0, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 1, Name: "Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(Vec2)
		if !ok {
			return p, fmt.Errorf("field AddActorPacket.Rotation has unexpected decoded type %T", raw)
		}
		p.Rotation = value
	}
	{
		raw, err := r.Read("AddActorPacket.Y Head Rotation", Shape{Kind: "primitive", PrimitiveCode: "f32le"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(float32)
		if !ok {
			return p, fmt.Errorf("field AddActorPacket.Y Head Rotation has unexpected decoded type %T", raw)
		}
		p.YHeadRotation = value
	}
	{
		raw, err := r.Read("AddActorPacket.Y Body Rotation", Shape{Kind: "primitive", PrimitiveCode: "f32le"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(float32)
		if !ok {
			return p, fmt.Errorf("field AddActorPacket.Y Body Rotation has unexpected decoded type %T", raw)
		}
		p.YBodyRotation = value
	}
	{
		raw, err := r.Read("AddActorPacket.Attributes List", Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "struct", Semantic: "SyncedAttribute", TypeID: "SyncedAttribute", Fields: []ShapeField{{Ordinal: 0, Name: "Attribute Name", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 1, Name: "Min Value", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 2, Name: "Current Value", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 3, Name: "Max Value", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.([]SyncedAttribute)
		if !ok {
			return p, fmt.Errorf("field AddActorPacket.Attributes List has unexpected decoded type %T", raw)
		}
		p.AttributesList = value
	}
	{
		raw, err := r.Read("AddActorPacket.Actor Data", Shape{Kind: "struct", Semantic: "SynchedActorData::CopyableDataList", TypeID: "SynchedActorData::CopyableDataList", Fields: []ShapeField{{Ordinal: 0, Name: "Data", Shape: Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "struct", Semantic: "DataItemEntry", TypeID: "DataItemEntry", Fields: []ShapeField{{Ordinal: 0, Name: "ID", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, {Ordinal: 1, Name: "Payload", Shape: Shape{Kind: "union", Control: &Shape{Kind: "primitive", PrimitiveCode: "u8"}, Variants: []ShapeVariant{{Value: 0, Name: "DataItemBytePayload", Shape: Shape{Kind: "struct", Semantic: "DataItemBytePayload", TypeID: "DataItemBytePayload", Fields: []ShapeField{{Ordinal: 0, Name: "Type", Shape: Shape{Kind: "enum", Semantic: "DataItemType", TypeID: "enums/DataItemType", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "Byte", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "Short", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "Int", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "Float", Shape: Shape{Kind: "void"}}, {Value: 4, Name: "String", Shape: Shape{Kind: "void"}}, {Value: 5, Name: "CompoundTag", Shape: Shape{Kind: "void"}}, {Value: 6, Name: "Pos", Shape: Shape{Kind: "void"}}, {Value: 7, Name: "Int64", Shape: Shape{Kind: "void"}}, {Value: 8, Name: "Vec3", Shape: Shape{Kind: "void"}}}}}, {Ordinal: 1, Name: "Value", Shape: Shape{Kind: "primitive", PrimitiveCode: "i8"}}}}}, {Value: 1, Name: "DataItemShortPayload", Shape: Shape{Kind: "struct", Semantic: "DataItemShortPayload", TypeID: "DataItemShortPayload", Fields: []ShapeField{{Ordinal: 0, Name: "Type", Shape: Shape{Kind: "enum", Semantic: "DataItemType", TypeID: "enums/DataItemType", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "Byte", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "Short", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "Int", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "Float", Shape: Shape{Kind: "void"}}, {Value: 4, Name: "String", Shape: Shape{Kind: "void"}}, {Value: 5, Name: "CompoundTag", Shape: Shape{Kind: "void"}}, {Value: 6, Name: "Pos", Shape: Shape{Kind: "void"}}, {Value: 7, Name: "Int64", Shape: Shape{Kind: "void"}}, {Value: 8, Name: "Vec3", Shape: Shape{Kind: "void"}}}}}, {Ordinal: 1, Name: "Value", Shape: Shape{Kind: "primitive", PrimitiveCode: "i16le"}}}}}, {Value: 2, Name: "DataItemIntPayload", Shape: Shape{Kind: "struct", Semantic: "DataItemIntPayload", TypeID: "DataItemIntPayload", Fields: []ShapeField{{Ordinal: 0, Name: "Type", Shape: Shape{Kind: "enum", Semantic: "DataItemType", TypeID: "enums/DataItemType", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "Byte", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "Short", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "Int", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "Float", Shape: Shape{Kind: "void"}}, {Value: 4, Name: "String", Shape: Shape{Kind: "void"}}, {Value: 5, Name: "CompoundTag", Shape: Shape{Kind: "void"}}, {Value: 6, Name: "Pos", Shape: Shape{Kind: "void"}}, {Value: 7, Name: "Int64", Shape: Shape{Kind: "void"}}, {Value: 8, Name: "Vec3", Shape: Shape{Kind: "void"}}}}}, {Ordinal: 1, Name: "Value", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}}}}, {Value: 3, Name: "DataItemFloatPayload", Shape: Shape{Kind: "struct", Semantic: "DataItemFloatPayload", TypeID: "DataItemFloatPayload", Fields: []ShapeField{{Ordinal: 0, Name: "Type", Shape: Shape{Kind: "enum", Semantic: "DataItemType", TypeID: "enums/DataItemType", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "Byte", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "Short", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "Int", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "Float", Shape: Shape{Kind: "void"}}, {Value: 4, Name: "String", Shape: Shape{Kind: "void"}}, {Value: 5, Name: "CompoundTag", Shape: Shape{Kind: "void"}}, {Value: 6, Name: "Pos", Shape: Shape{Kind: "void"}}, {Value: 7, Name: "Int64", Shape: Shape{Kind: "void"}}, {Value: 8, Name: "Vec3", Shape: Shape{Kind: "void"}}}}}, {Ordinal: 1, Name: "Value", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}}}}, {Value: 4, Name: "DataItemStringPayload", Shape: Shape{Kind: "struct", Semantic: "DataItemStringPayload", TypeID: "DataItemStringPayload", Fields: []ShapeField{{Ordinal: 0, Name: "Type", Shape: Shape{Kind: "enum", Semantic: "DataItemType", TypeID: "enums/DataItemType", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "Byte", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "Short", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "Int", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "Float", Shape: Shape{Kind: "void"}}, {Value: 4, Name: "String", Shape: Shape{Kind: "void"}}, {Value: 5, Name: "CompoundTag", Shape: Shape{Kind: "void"}}, {Value: 6, Name: "Pos", Shape: Shape{Kind: "void"}}, {Value: 7, Name: "Int64", Shape: Shape{Kind: "void"}}, {Value: 8, Name: "Vec3", Shape: Shape{Kind: "void"}}}}}, {Ordinal: 1, Name: "Value", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}}}, {Value: 5, Name: "DataItemCompoundTagPayload", Shape: Shape{Kind: "struct", Semantic: "DataItemCompoundTagPayload", TypeID: "DataItemCompoundTagPayload", Fields: []ShapeField{{Ordinal: 0, Name: "Type", Shape: Shape{Kind: "enum", Semantic: "DataItemType", TypeID: "enums/DataItemType", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "Byte", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "Short", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "Int", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "Float", Shape: Shape{Kind: "void"}}, {Value: 4, Name: "String", Shape: Shape{Kind: "void"}}, {Value: 5, Name: "CompoundTag", Shape: Shape{Kind: "void"}}, {Value: 6, Name: "Pos", Shape: Shape{Kind: "void"}}, {Value: 7, Name: "Int64", Shape: Shape{Kind: "void"}}, {Value: 8, Name: "Vec3", Shape: Shape{Kind: "void"}}}}}, {Ordinal: 1, Name: "Value", Shape: Shape{Kind: "primitive", PrimitiveCode: "nbt_le"}}}}}, {Value: 6, Name: "DataItemPosPayload", Shape: Shape{Kind: "struct", Semantic: "DataItemPosPayload", TypeID: "DataItemPosPayload", Fields: []ShapeField{{Ordinal: 0, Name: "Type", Shape: Shape{Kind: "enum", Semantic: "DataItemType", TypeID: "enums/DataItemType", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "Byte", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "Short", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "Int", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "Float", Shape: Shape{Kind: "void"}}, {Value: 4, Name: "String", Shape: Shape{Kind: "void"}}, {Value: 5, Name: "CompoundTag", Shape: Shape{Kind: "void"}}, {Value: 6, Name: "Pos", Shape: Shape{Kind: "void"}}, {Value: 7, Name: "Int64", Shape: Shape{Kind: "void"}}, {Value: 8, Name: "Vec3", Shape: Shape{Kind: "void"}}}}}, {Ordinal: 1, Name: "Value", Shape: Shape{Kind: "struct", Semantic: "BlockPos", TypeID: "BlockPos", Fields: []ShapeField{{Ordinal: 0, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 1, Name: "Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 2, Name: "Z", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}}}}}}}, {Value: 7, Name: "DataItemInt64Payload", Shape: Shape{Kind: "struct", Semantic: "DataItemInt64Payload", TypeID: "DataItemInt64Payload", Fields: []ShapeField{{Ordinal: 0, Name: "Type", Shape: Shape{Kind: "enum", Semantic: "DataItemType", TypeID: "enums/DataItemType", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "Byte", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "Short", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "Int", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "Float", Shape: Shape{Kind: "void"}}, {Value: 4, Name: "String", Shape: Shape{Kind: "void"}}, {Value: 5, Name: "CompoundTag", Shape: Shape{Kind: "void"}}, {Value: 6, Name: "Pos", Shape: Shape{Kind: "void"}}, {Value: 7, Name: "Int64", Shape: Shape{Kind: "void"}}, {Value: 8, Name: "Vec3", Shape: Shape{Kind: "void"}}}}}, {Ordinal: 1, Name: "Value", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i64"}}}}}, {Value: 8, Name: "DataItemVec3Payload", Shape: Shape{Kind: "struct", Semantic: "DataItemVec3Payload", TypeID: "DataItemVec3Payload", Fields: []ShapeField{{Ordinal: 0, Name: "Type", Shape: Shape{Kind: "enum", Semantic: "DataItemType", TypeID: "enums/DataItemType", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "Byte", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "Short", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "Int", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "Float", Shape: Shape{Kind: "void"}}, {Value: 4, Name: "String", Shape: Shape{Kind: "void"}}, {Value: 5, Name: "CompoundTag", Shape: Shape{Kind: "void"}}, {Value: 6, Name: "Pos", Shape: Shape{Kind: "void"}}, {Value: 7, Name: "Int64", Shape: Shape{Kind: "void"}}, {Value: 8, Name: "Vec3", Shape: Shape{Kind: "void"}}}}}, {Ordinal: 1, Name: "Value", Shape: Shape{Kind: "struct", Semantic: "Vec3", TypeID: "Vec3", Fields: []ShapeField{{Ordinal: 0, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 1, Name: "Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 2, Name: "Z", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}}}}}}}}}}}}}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(SynchedActorDataCopyableDataList)
		if !ok {
			return p, fmt.Errorf("field AddActorPacket.Actor Data has unexpected decoded type %T", raw)
		}
		p.ActorData = value
	}
	{
		raw, err := r.Read("AddActorPacket.Synched Properties", Shape{Kind: "struct", Semantic: "PropertySyncData", TypeID: "PropertySyncData", Fields: []ShapeField{{Ordinal: 0, Name: "Int Entries List", Shape: Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "struct", Semantic: "PropertySyncData::PropertySyncIntEntry", TypeID: "PropertySyncData::PropertySyncIntEntry", Fields: []ShapeField{{Ordinal: 0, Name: "Property Index", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, {Ordinal: 1, Name: "Data", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}}}}}, {Ordinal: 1, Name: "Float Entries List", Shape: Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "struct", Semantic: "PropertySyncData::PropertySyncFloatEntry", TypeID: "PropertySyncData::PropertySyncFloatEntry", Fields: []ShapeField{{Ordinal: 0, Name: "Property Index", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, {Ordinal: 1, Name: "Data", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}}}}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(PropertySyncData)
		if !ok {
			return p, fmt.Errorf("field AddActorPacket.Synched Properties has unexpected decoded type %T", raw)
		}
		p.SynchedProperties = value
	}
	{
		raw, err := r.Read("AddActorPacket.Actor Links", Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "struct", Semantic: "ActorLink", TypeID: "ActorLink", Fields: []ShapeField{{Ordinal: 0, Name: "Target A", Shape: Shape{Kind: "struct", Semantic: "ActorUniqueID", TypeID: "ActorUniqueID", Fields: []ShapeField{{Ordinal: 0, Name: "Actor Unique ID", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i64"}}}}}, {Ordinal: 1, Name: "Target B", Shape: Shape{Kind: "struct", Semantic: "ActorUniqueID", TypeID: "ActorUniqueID", Fields: []ShapeField{{Ordinal: 0, Name: "Actor Unique ID", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i64"}}}}}, {Ordinal: 2, Name: "Type", Shape: Shape{Kind: "enum", Semantic: "ActorLinkType", TypeID: "enums/ActorLinkType", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "None", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "Riding", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "Passenger", Shape: Shape{Kind: "void"}}}}}, {Ordinal: 3, Name: "Immediate", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}, {Ordinal: 4, Name: "Passenger Initiated", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}, {Ordinal: 5, Name: "Vehicle Angular Velocity", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.([]ActorLink)
		if !ok {
			return p, fmt.Errorf("field AddActorPacket.Actor Links has unexpected decoded type %T", raw)
		}
		p.ActorLinks = value
	}
	return p, nil
}
