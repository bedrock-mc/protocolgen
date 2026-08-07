// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type AddPlayer struct {
	UUID              [16]byte
	PlayerName        string
	TargetRuntimeID   ActorRuntimeID
	PlatformChatId    string
	Position          Vec3
	Velocity          Vec3
	Rotation          Vec2
	YHeadRotation     float32
	CarriedItem       CerealizerNetworkItemStackDescriptorSerializedData
	PlayerGameType    GameType
	EntityData        SynchedActorDataCopyableDataList
	SynchedProperties PropertySyncData
	AbilitiesData     SerializedAbilitiesData
	ActorLinks        []ActorLink
	DeviceId          string
	BuildPlatform     BuildPlatform
}

func (p *AddPlayer) Encode(w Encoder) error {
	if err := w.Write("AddPlayerPacket.UUID", Shape{Kind: "primitive", Semantic: "mce::UUID", TypeID: "mce__UUID.json#", PrimitiveCode: "uuid"}, p.UUID); err != nil {
		return err
	}
	if err := w.Write("AddPlayerPacket.Player Name", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, p.PlayerName); err != nil {
		return err
	}
	if err := w.Write("AddPlayerPacket.Target Runtime ID", Shape{Kind: "struct", Semantic: "ActorRuntimeID", TypeID: "ActorRuntimeID", Fields: []ShapeField{{Ordinal: 0, Name: "Actor Runtime ID", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u64"}}}}, p.TargetRuntimeID); err != nil {
		return err
	}
	if err := w.Write("AddPlayerPacket.Platform Chat Id", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, p.PlatformChatId); err != nil {
		return err
	}
	if err := w.Write("AddPlayerPacket.Position", Shape{Kind: "struct", Semantic: "Vec3", TypeID: "Vec3", Fields: []ShapeField{{Ordinal: 0, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 1, Name: "Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 2, Name: "Z", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}}}, p.Position); err != nil {
		return err
	}
	if err := w.Write("AddPlayerPacket.Velocity", Shape{Kind: "struct", Semantic: "Vec3", TypeID: "Vec3", Fields: []ShapeField{{Ordinal: 0, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 1, Name: "Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 2, Name: "Z", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}}}, p.Velocity); err != nil {
		return err
	}
	if err := w.Write("AddPlayerPacket.Rotation", Shape{Kind: "struct", Semantic: "Vec2", TypeID: "Vec2", Fields: []ShapeField{{Ordinal: 0, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 1, Name: "Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}}}, p.Rotation); err != nil {
		return err
	}
	if err := w.Write("AddPlayerPacket.Y-Head Rotation", Shape{Kind: "primitive", PrimitiveCode: "f32le"}, p.YHeadRotation); err != nil {
		return err
	}
	if err := w.Write("AddPlayerPacket.Carried Item", Shape{Kind: "struct", Semantic: "cerealizer<NetworkItemStackDescriptor>::SerializedData", TypeID: "cerealizer<NetworkItemStackDescriptor>::SerializedData", Fields: []ShapeField{{Ordinal: 0, Name: "Id", Shape: Shape{Kind: "primitive", PrimitiveCode: "i16le"}}, {Ordinal: 1, Name: "Stack size", Shape: Shape{Kind: "primitive", PrimitiveCode: "u16le"}}, {Ordinal: 2, Name: "Aux value", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, {Ordinal: 3, Name: "Net Id Variant", Shape: Shape{Kind: "optional", Value: &Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}}, {Ordinal: 4, Name: "Block Runtime Id", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, {Ordinal: 5, Name: "User Data Buffer", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}}, p.CarriedItem); err != nil {
		return err
	}
	if err := w.Write("AddPlayerPacket.Player Game Type", Shape{Kind: "enum", Semantic: "GameType", TypeID: "enums/GameType", PrimitiveCode: "zigzag_i32", Variants: []ShapeVariant{{Value: -1, Name: "Undefined", Shape: Shape{Kind: "void"}}, {Value: 0, Name: "Survival", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "Creative", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "Adventure", Shape: Shape{Kind: "void"}}, {Value: 5, Name: "Default", Shape: Shape{Kind: "void"}}, {Value: 6, Name: "Spectator", Shape: Shape{Kind: "void"}}}}, p.PlayerGameType); err != nil {
		return err
	}
	if err := w.Write("AddPlayerPacket.Entity Data", Shape{Kind: "struct", Semantic: "SynchedActorData::CopyableDataList", TypeID: "SynchedActorData::CopyableDataList", Fields: []ShapeField{{Ordinal: 0, Name: "Data", Shape: Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "struct", Semantic: "DataItemEntry", TypeID: "DataItemEntry", Fields: []ShapeField{{Ordinal: 0, Name: "ID", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, {Ordinal: 1, Name: "Payload", Shape: Shape{Kind: "union", Control: &Shape{Kind: "primitive", PrimitiveCode: "u8"}, Variants: []ShapeVariant{{Value: 0, Name: "DataItemBytePayload", Shape: Shape{Kind: "struct", Semantic: "DataItemBytePayload", TypeID: "DataItemBytePayload", Fields: []ShapeField{{Ordinal: 0, Name: "Type", Shape: Shape{Kind: "enum", Semantic: "DataItemType", TypeID: "enums/DataItemType", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "Byte", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "Short", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "Int", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "Float", Shape: Shape{Kind: "void"}}, {Value: 4, Name: "String", Shape: Shape{Kind: "void"}}, {Value: 5, Name: "CompoundTag", Shape: Shape{Kind: "void"}}, {Value: 6, Name: "Pos", Shape: Shape{Kind: "void"}}, {Value: 7, Name: "Int64", Shape: Shape{Kind: "void"}}, {Value: 8, Name: "Vec3", Shape: Shape{Kind: "void"}}}}}, {Ordinal: 1, Name: "Value", Shape: Shape{Kind: "primitive", PrimitiveCode: "i8"}}}}}, {Value: 1, Name: "DataItemShortPayload", Shape: Shape{Kind: "struct", Semantic: "DataItemShortPayload", TypeID: "DataItemShortPayload", Fields: []ShapeField{{Ordinal: 0, Name: "Type", Shape: Shape{Kind: "enum", Semantic: "DataItemType", TypeID: "enums/DataItemType", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "Byte", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "Short", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "Int", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "Float", Shape: Shape{Kind: "void"}}, {Value: 4, Name: "String", Shape: Shape{Kind: "void"}}, {Value: 5, Name: "CompoundTag", Shape: Shape{Kind: "void"}}, {Value: 6, Name: "Pos", Shape: Shape{Kind: "void"}}, {Value: 7, Name: "Int64", Shape: Shape{Kind: "void"}}, {Value: 8, Name: "Vec3", Shape: Shape{Kind: "void"}}}}}, {Ordinal: 1, Name: "Value", Shape: Shape{Kind: "primitive", PrimitiveCode: "i16le"}}}}}, {Value: 2, Name: "DataItemIntPayload", Shape: Shape{Kind: "struct", Semantic: "DataItemIntPayload", TypeID: "DataItemIntPayload", Fields: []ShapeField{{Ordinal: 0, Name: "Type", Shape: Shape{Kind: "enum", Semantic: "DataItemType", TypeID: "enums/DataItemType", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "Byte", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "Short", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "Int", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "Float", Shape: Shape{Kind: "void"}}, {Value: 4, Name: "String", Shape: Shape{Kind: "void"}}, {Value: 5, Name: "CompoundTag", Shape: Shape{Kind: "void"}}, {Value: 6, Name: "Pos", Shape: Shape{Kind: "void"}}, {Value: 7, Name: "Int64", Shape: Shape{Kind: "void"}}, {Value: 8, Name: "Vec3", Shape: Shape{Kind: "void"}}}}}, {Ordinal: 1, Name: "Value", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}}}}, {Value: 3, Name: "DataItemFloatPayload", Shape: Shape{Kind: "struct", Semantic: "DataItemFloatPayload", TypeID: "DataItemFloatPayload", Fields: []ShapeField{{Ordinal: 0, Name: "Type", Shape: Shape{Kind: "enum", Semantic: "DataItemType", TypeID: "enums/DataItemType", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "Byte", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "Short", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "Int", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "Float", Shape: Shape{Kind: "void"}}, {Value: 4, Name: "String", Shape: Shape{Kind: "void"}}, {Value: 5, Name: "CompoundTag", Shape: Shape{Kind: "void"}}, {Value: 6, Name: "Pos", Shape: Shape{Kind: "void"}}, {Value: 7, Name: "Int64", Shape: Shape{Kind: "void"}}, {Value: 8, Name: "Vec3", Shape: Shape{Kind: "void"}}}}}, {Ordinal: 1, Name: "Value", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}}}}, {Value: 4, Name: "DataItemStringPayload", Shape: Shape{Kind: "struct", Semantic: "DataItemStringPayload", TypeID: "DataItemStringPayload", Fields: []ShapeField{{Ordinal: 0, Name: "Type", Shape: Shape{Kind: "enum", Semantic: "DataItemType", TypeID: "enums/DataItemType", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "Byte", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "Short", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "Int", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "Float", Shape: Shape{Kind: "void"}}, {Value: 4, Name: "String", Shape: Shape{Kind: "void"}}, {Value: 5, Name: "CompoundTag", Shape: Shape{Kind: "void"}}, {Value: 6, Name: "Pos", Shape: Shape{Kind: "void"}}, {Value: 7, Name: "Int64", Shape: Shape{Kind: "void"}}, {Value: 8, Name: "Vec3", Shape: Shape{Kind: "void"}}}}}, {Ordinal: 1, Name: "Value", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}}}, {Value: 5, Name: "DataItemCompoundTagPayload", Shape: Shape{Kind: "struct", Semantic: "DataItemCompoundTagPayload", TypeID: "DataItemCompoundTagPayload", Fields: []ShapeField{{Ordinal: 0, Name: "Type", Shape: Shape{Kind: "enum", Semantic: "DataItemType", TypeID: "enums/DataItemType", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "Byte", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "Short", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "Int", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "Float", Shape: Shape{Kind: "void"}}, {Value: 4, Name: "String", Shape: Shape{Kind: "void"}}, {Value: 5, Name: "CompoundTag", Shape: Shape{Kind: "void"}}, {Value: 6, Name: "Pos", Shape: Shape{Kind: "void"}}, {Value: 7, Name: "Int64", Shape: Shape{Kind: "void"}}, {Value: 8, Name: "Vec3", Shape: Shape{Kind: "void"}}}}}, {Ordinal: 1, Name: "Value", Shape: Shape{Kind: "primitive", PrimitiveCode: "nbt_le"}}}}}, {Value: 6, Name: "DataItemPosPayload", Shape: Shape{Kind: "struct", Semantic: "DataItemPosPayload", TypeID: "DataItemPosPayload", Fields: []ShapeField{{Ordinal: 0, Name: "Type", Shape: Shape{Kind: "enum", Semantic: "DataItemType", TypeID: "enums/DataItemType", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "Byte", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "Short", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "Int", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "Float", Shape: Shape{Kind: "void"}}, {Value: 4, Name: "String", Shape: Shape{Kind: "void"}}, {Value: 5, Name: "CompoundTag", Shape: Shape{Kind: "void"}}, {Value: 6, Name: "Pos", Shape: Shape{Kind: "void"}}, {Value: 7, Name: "Int64", Shape: Shape{Kind: "void"}}, {Value: 8, Name: "Vec3", Shape: Shape{Kind: "void"}}}}}, {Ordinal: 1, Name: "Value", Shape: Shape{Kind: "struct", Semantic: "BlockPos", TypeID: "BlockPos", Fields: []ShapeField{{Ordinal: 0, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 1, Name: "Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 2, Name: "Z", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}}}}}}}, {Value: 7, Name: "DataItemInt64Payload", Shape: Shape{Kind: "struct", Semantic: "DataItemInt64Payload", TypeID: "DataItemInt64Payload", Fields: []ShapeField{{Ordinal: 0, Name: "Type", Shape: Shape{Kind: "enum", Semantic: "DataItemType", TypeID: "enums/DataItemType", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "Byte", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "Short", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "Int", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "Float", Shape: Shape{Kind: "void"}}, {Value: 4, Name: "String", Shape: Shape{Kind: "void"}}, {Value: 5, Name: "CompoundTag", Shape: Shape{Kind: "void"}}, {Value: 6, Name: "Pos", Shape: Shape{Kind: "void"}}, {Value: 7, Name: "Int64", Shape: Shape{Kind: "void"}}, {Value: 8, Name: "Vec3", Shape: Shape{Kind: "void"}}}}}, {Ordinal: 1, Name: "Value", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i64"}}}}}, {Value: 8, Name: "DataItemVec3Payload", Shape: Shape{Kind: "struct", Semantic: "DataItemVec3Payload", TypeID: "DataItemVec3Payload", Fields: []ShapeField{{Ordinal: 0, Name: "Type", Shape: Shape{Kind: "enum", Semantic: "DataItemType", TypeID: "enums/DataItemType", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "Byte", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "Short", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "Int", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "Float", Shape: Shape{Kind: "void"}}, {Value: 4, Name: "String", Shape: Shape{Kind: "void"}}, {Value: 5, Name: "CompoundTag", Shape: Shape{Kind: "void"}}, {Value: 6, Name: "Pos", Shape: Shape{Kind: "void"}}, {Value: 7, Name: "Int64", Shape: Shape{Kind: "void"}}, {Value: 8, Name: "Vec3", Shape: Shape{Kind: "void"}}}}}, {Ordinal: 1, Name: "Value", Shape: Shape{Kind: "struct", Semantic: "Vec3", TypeID: "Vec3", Fields: []ShapeField{{Ordinal: 0, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 1, Name: "Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 2, Name: "Z", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}}}}}}}}}}}}}}}}, p.EntityData); err != nil {
		return err
	}
	if err := w.Write("AddPlayerPacket.Synched Properties", Shape{Kind: "struct", Semantic: "PropertySyncData", TypeID: "PropertySyncData", Fields: []ShapeField{{Ordinal: 0, Name: "Int Entries List", Shape: Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "struct", Semantic: "PropertySyncData::PropertySyncIntEntry", TypeID: "PropertySyncData::PropertySyncIntEntry", Fields: []ShapeField{{Ordinal: 0, Name: "Property Index", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, {Ordinal: 1, Name: "Data", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}}}}}, {Ordinal: 1, Name: "Float Entries List", Shape: Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "struct", Semantic: "PropertySyncData::PropertySyncFloatEntry", TypeID: "PropertySyncData::PropertySyncFloatEntry", Fields: []ShapeField{{Ordinal: 0, Name: "Property Index", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, {Ordinal: 1, Name: "Data", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}}}}}}}, p.SynchedProperties); err != nil {
		return err
	}
	if err := w.Write("AddPlayerPacket.Abilities Data", Shape{Kind: "struct", Semantic: "SerializedAbilitiesData", TypeID: "SerializedAbilitiesData", Fields: []ShapeField{{Ordinal: 0, Name: "Target Player Raw Id", Shape: Shape{Kind: "primitive", PrimitiveCode: "i64le"}}, {Ordinal: 1, Name: "Player Permissions", Shape: Shape{Kind: "enum", Semantic: "PlayerPermissionLevel", TypeID: "enums/PlayerPermissionLevel", PrimitiveCode: "i8", Variants: []ShapeVariant{{Value: 0, Name: "Visitor", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "Member", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "Operator", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "Custom", Shape: Shape{Kind: "void"}}}}}, {Ordinal: 2, Name: "Command Permissions", Shape: Shape{Kind: "enum", Semantic: "CommandPermissionLevel", TypeID: "enums/CommandPermissionLevel", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "Any", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "GameDirectors", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "Admin", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "Host", Shape: Shape{Kind: "void"}}, {Value: 4, Name: "Owner", Shape: Shape{Kind: "void"}}, {Value: 5, Name: "Internal", Shape: Shape{Kind: "void"}}}}}, {Ordinal: 3, Name: "Layers", Shape: Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "struct", Semantic: "SerializedAbilitiesData::SerializedLayer", TypeID: "SerializedAbilitiesData::SerializedLayer", Fields: []ShapeField{{Ordinal: 0, Name: "SerializedLayer", Shape: Shape{Kind: "primitive", PrimitiveCode: "u16le"}}, {Ordinal: 1, Name: "AbilitiesSet", Shape: Shape{Kind: "primitive", PrimitiveCode: "u32le"}}, {Ordinal: 2, Name: "AbilityValues", Shape: Shape{Kind: "primitive", PrimitiveCode: "u32le"}}, {Ordinal: 3, Name: "FlySpeed", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 4, Name: "VerticalFlySpeed", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 5, Name: "WalkSpeed", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}}}}}}}, p.AbilitiesData); err != nil {
		return err
	}
	if err := w.Write("AddPlayerPacket.Actor Links", Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "struct", Semantic: "ActorLink", TypeID: "ActorLink", Fields: []ShapeField{{Ordinal: 0, Name: "Target A", Shape: Shape{Kind: "struct", Semantic: "ActorUniqueID", TypeID: "ActorUniqueID", Fields: []ShapeField{{Ordinal: 0, Name: "Actor Unique ID", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i64"}}}}}, {Ordinal: 1, Name: "Target B", Shape: Shape{Kind: "struct", Semantic: "ActorUniqueID", TypeID: "ActorUniqueID", Fields: []ShapeField{{Ordinal: 0, Name: "Actor Unique ID", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i64"}}}}}, {Ordinal: 2, Name: "Type", Shape: Shape{Kind: "enum", Semantic: "ActorLinkType", TypeID: "enums/ActorLinkType", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "None", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "Riding", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "Passenger", Shape: Shape{Kind: "void"}}}}}, {Ordinal: 3, Name: "Immediate", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}, {Ordinal: 4, Name: "Passenger Initiated", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}, {Ordinal: 5, Name: "Vehicle Angular Velocity", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}}}}, p.ActorLinks); err != nil {
		return err
	}
	if err := w.Write("AddPlayerPacket.Device Id", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, p.DeviceId); err != nil {
		return err
	}
	if err := w.Write("AddPlayerPacket.Build Platform", Shape{Kind: "enum", Semantic: "BuildPlatform", TypeID: "enums/BuildPlatform", PrimitiveCode: "i32le", Variants: []ShapeVariant{{Value: -1, Name: "Unknown", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "Google", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "iOS", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "OSX", Shape: Shape{Kind: "void"}}, {Value: 4, Name: "Amazon", Shape: Shape{Kind: "void"}}, {Value: 5, Name: "GearVR", Shape: Shape{Kind: "void"}}, {Value: 7, Name: "UWP", Shape: Shape{Kind: "void"}}, {Value: 8, Name: "Win32", Shape: Shape{Kind: "void"}}, {Value: 9, Name: "Dedicated", Shape: Shape{Kind: "void"}}, {Value: 10, Name: "tvOS", Shape: Shape{Kind: "void"}}, {Value: 11, Name: "Sony", Shape: Shape{Kind: "void"}}, {Value: 12, Name: "Nx", Shape: Shape{Kind: "void"}}, {Value: 13, Name: "Xbox", Shape: Shape{Kind: "void"}}, {Value: 14, Name: "WindowsPhone", Shape: Shape{Kind: "void"}}, {Value: 15, Name: "Linux", Shape: Shape{Kind: "void"}}}}, p.BuildPlatform); err != nil {
		return err
	}
	return nil
}

func DecodeAddPlayer(r Decoder) (AddPlayer, error) {
	var p AddPlayer
	{
		raw, err := r.Read("AddPlayerPacket.UUID", Shape{Kind: "primitive", Semantic: "mce::UUID", TypeID: "mce__UUID.json#", PrimitiveCode: "uuid"})
		if err != nil {
			return p, err
		}
		value, ok := raw.([16]byte)
		if !ok {
			return p, fmt.Errorf("field AddPlayerPacket.UUID has unexpected decoded type %T", raw)
		}
		p.UUID = value
	}
	{
		raw, err := r.Read("AddPlayerPacket.Player Name", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(string)
		if !ok {
			return p, fmt.Errorf("field AddPlayerPacket.Player Name has unexpected decoded type %T", raw)
		}
		p.PlayerName = value
	}
	{
		raw, err := r.Read("AddPlayerPacket.Target Runtime ID", Shape{Kind: "struct", Semantic: "ActorRuntimeID", TypeID: "ActorRuntimeID", Fields: []ShapeField{{Ordinal: 0, Name: "Actor Runtime ID", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u64"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(ActorRuntimeID)
		if !ok {
			return p, fmt.Errorf("field AddPlayerPacket.Target Runtime ID has unexpected decoded type %T", raw)
		}
		p.TargetRuntimeID = value
	}
	{
		raw, err := r.Read("AddPlayerPacket.Platform Chat Id", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(string)
		if !ok {
			return p, fmt.Errorf("field AddPlayerPacket.Platform Chat Id has unexpected decoded type %T", raw)
		}
		p.PlatformChatId = value
	}
	{
		raw, err := r.Read("AddPlayerPacket.Position", Shape{Kind: "struct", Semantic: "Vec3", TypeID: "Vec3", Fields: []ShapeField{{Ordinal: 0, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 1, Name: "Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 2, Name: "Z", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(Vec3)
		if !ok {
			return p, fmt.Errorf("field AddPlayerPacket.Position has unexpected decoded type %T", raw)
		}
		p.Position = value
	}
	{
		raw, err := r.Read("AddPlayerPacket.Velocity", Shape{Kind: "struct", Semantic: "Vec3", TypeID: "Vec3", Fields: []ShapeField{{Ordinal: 0, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 1, Name: "Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 2, Name: "Z", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(Vec3)
		if !ok {
			return p, fmt.Errorf("field AddPlayerPacket.Velocity has unexpected decoded type %T", raw)
		}
		p.Velocity = value
	}
	{
		raw, err := r.Read("AddPlayerPacket.Rotation", Shape{Kind: "struct", Semantic: "Vec2", TypeID: "Vec2", Fields: []ShapeField{{Ordinal: 0, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 1, Name: "Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(Vec2)
		if !ok {
			return p, fmt.Errorf("field AddPlayerPacket.Rotation has unexpected decoded type %T", raw)
		}
		p.Rotation = value
	}
	{
		raw, err := r.Read("AddPlayerPacket.Y-Head Rotation", Shape{Kind: "primitive", PrimitiveCode: "f32le"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(float32)
		if !ok {
			return p, fmt.Errorf("field AddPlayerPacket.Y-Head Rotation has unexpected decoded type %T", raw)
		}
		p.YHeadRotation = value
	}
	{
		raw, err := r.Read("AddPlayerPacket.Carried Item", Shape{Kind: "struct", Semantic: "cerealizer<NetworkItemStackDescriptor>::SerializedData", TypeID: "cerealizer<NetworkItemStackDescriptor>::SerializedData", Fields: []ShapeField{{Ordinal: 0, Name: "Id", Shape: Shape{Kind: "primitive", PrimitiveCode: "i16le"}}, {Ordinal: 1, Name: "Stack size", Shape: Shape{Kind: "primitive", PrimitiveCode: "u16le"}}, {Ordinal: 2, Name: "Aux value", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, {Ordinal: 3, Name: "Net Id Variant", Shape: Shape{Kind: "optional", Value: &Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}}, {Ordinal: 4, Name: "Block Runtime Id", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, {Ordinal: 5, Name: "User Data Buffer", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(CerealizerNetworkItemStackDescriptorSerializedData)
		if !ok {
			return p, fmt.Errorf("field AddPlayerPacket.Carried Item has unexpected decoded type %T", raw)
		}
		p.CarriedItem = value
	}
	{
		raw, err := r.Read("AddPlayerPacket.Player Game Type", Shape{Kind: "enum", Semantic: "GameType", TypeID: "enums/GameType", PrimitiveCode: "zigzag_i32", Variants: []ShapeVariant{{Value: -1, Name: "Undefined", Shape: Shape{Kind: "void"}}, {Value: 0, Name: "Survival", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "Creative", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "Adventure", Shape: Shape{Kind: "void"}}, {Value: 5, Name: "Default", Shape: Shape{Kind: "void"}}, {Value: 6, Name: "Spectator", Shape: Shape{Kind: "void"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(GameType)
		if !ok {
			return p, fmt.Errorf("field AddPlayerPacket.Player Game Type has unexpected decoded type %T", raw)
		}
		p.PlayerGameType = value
	}
	{
		raw, err := r.Read("AddPlayerPacket.Entity Data", Shape{Kind: "struct", Semantic: "SynchedActorData::CopyableDataList", TypeID: "SynchedActorData::CopyableDataList", Fields: []ShapeField{{Ordinal: 0, Name: "Data", Shape: Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "struct", Semantic: "DataItemEntry", TypeID: "DataItemEntry", Fields: []ShapeField{{Ordinal: 0, Name: "ID", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, {Ordinal: 1, Name: "Payload", Shape: Shape{Kind: "union", Control: &Shape{Kind: "primitive", PrimitiveCode: "u8"}, Variants: []ShapeVariant{{Value: 0, Name: "DataItemBytePayload", Shape: Shape{Kind: "struct", Semantic: "DataItemBytePayload", TypeID: "DataItemBytePayload", Fields: []ShapeField{{Ordinal: 0, Name: "Type", Shape: Shape{Kind: "enum", Semantic: "DataItemType", TypeID: "enums/DataItemType", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "Byte", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "Short", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "Int", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "Float", Shape: Shape{Kind: "void"}}, {Value: 4, Name: "String", Shape: Shape{Kind: "void"}}, {Value: 5, Name: "CompoundTag", Shape: Shape{Kind: "void"}}, {Value: 6, Name: "Pos", Shape: Shape{Kind: "void"}}, {Value: 7, Name: "Int64", Shape: Shape{Kind: "void"}}, {Value: 8, Name: "Vec3", Shape: Shape{Kind: "void"}}}}}, {Ordinal: 1, Name: "Value", Shape: Shape{Kind: "primitive", PrimitiveCode: "i8"}}}}}, {Value: 1, Name: "DataItemShortPayload", Shape: Shape{Kind: "struct", Semantic: "DataItemShortPayload", TypeID: "DataItemShortPayload", Fields: []ShapeField{{Ordinal: 0, Name: "Type", Shape: Shape{Kind: "enum", Semantic: "DataItemType", TypeID: "enums/DataItemType", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "Byte", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "Short", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "Int", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "Float", Shape: Shape{Kind: "void"}}, {Value: 4, Name: "String", Shape: Shape{Kind: "void"}}, {Value: 5, Name: "CompoundTag", Shape: Shape{Kind: "void"}}, {Value: 6, Name: "Pos", Shape: Shape{Kind: "void"}}, {Value: 7, Name: "Int64", Shape: Shape{Kind: "void"}}, {Value: 8, Name: "Vec3", Shape: Shape{Kind: "void"}}}}}, {Ordinal: 1, Name: "Value", Shape: Shape{Kind: "primitive", PrimitiveCode: "i16le"}}}}}, {Value: 2, Name: "DataItemIntPayload", Shape: Shape{Kind: "struct", Semantic: "DataItemIntPayload", TypeID: "DataItemIntPayload", Fields: []ShapeField{{Ordinal: 0, Name: "Type", Shape: Shape{Kind: "enum", Semantic: "DataItemType", TypeID: "enums/DataItemType", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "Byte", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "Short", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "Int", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "Float", Shape: Shape{Kind: "void"}}, {Value: 4, Name: "String", Shape: Shape{Kind: "void"}}, {Value: 5, Name: "CompoundTag", Shape: Shape{Kind: "void"}}, {Value: 6, Name: "Pos", Shape: Shape{Kind: "void"}}, {Value: 7, Name: "Int64", Shape: Shape{Kind: "void"}}, {Value: 8, Name: "Vec3", Shape: Shape{Kind: "void"}}}}}, {Ordinal: 1, Name: "Value", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}}}}, {Value: 3, Name: "DataItemFloatPayload", Shape: Shape{Kind: "struct", Semantic: "DataItemFloatPayload", TypeID: "DataItemFloatPayload", Fields: []ShapeField{{Ordinal: 0, Name: "Type", Shape: Shape{Kind: "enum", Semantic: "DataItemType", TypeID: "enums/DataItemType", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "Byte", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "Short", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "Int", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "Float", Shape: Shape{Kind: "void"}}, {Value: 4, Name: "String", Shape: Shape{Kind: "void"}}, {Value: 5, Name: "CompoundTag", Shape: Shape{Kind: "void"}}, {Value: 6, Name: "Pos", Shape: Shape{Kind: "void"}}, {Value: 7, Name: "Int64", Shape: Shape{Kind: "void"}}, {Value: 8, Name: "Vec3", Shape: Shape{Kind: "void"}}}}}, {Ordinal: 1, Name: "Value", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}}}}, {Value: 4, Name: "DataItemStringPayload", Shape: Shape{Kind: "struct", Semantic: "DataItemStringPayload", TypeID: "DataItemStringPayload", Fields: []ShapeField{{Ordinal: 0, Name: "Type", Shape: Shape{Kind: "enum", Semantic: "DataItemType", TypeID: "enums/DataItemType", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "Byte", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "Short", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "Int", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "Float", Shape: Shape{Kind: "void"}}, {Value: 4, Name: "String", Shape: Shape{Kind: "void"}}, {Value: 5, Name: "CompoundTag", Shape: Shape{Kind: "void"}}, {Value: 6, Name: "Pos", Shape: Shape{Kind: "void"}}, {Value: 7, Name: "Int64", Shape: Shape{Kind: "void"}}, {Value: 8, Name: "Vec3", Shape: Shape{Kind: "void"}}}}}, {Ordinal: 1, Name: "Value", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}}}, {Value: 5, Name: "DataItemCompoundTagPayload", Shape: Shape{Kind: "struct", Semantic: "DataItemCompoundTagPayload", TypeID: "DataItemCompoundTagPayload", Fields: []ShapeField{{Ordinal: 0, Name: "Type", Shape: Shape{Kind: "enum", Semantic: "DataItemType", TypeID: "enums/DataItemType", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "Byte", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "Short", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "Int", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "Float", Shape: Shape{Kind: "void"}}, {Value: 4, Name: "String", Shape: Shape{Kind: "void"}}, {Value: 5, Name: "CompoundTag", Shape: Shape{Kind: "void"}}, {Value: 6, Name: "Pos", Shape: Shape{Kind: "void"}}, {Value: 7, Name: "Int64", Shape: Shape{Kind: "void"}}, {Value: 8, Name: "Vec3", Shape: Shape{Kind: "void"}}}}}, {Ordinal: 1, Name: "Value", Shape: Shape{Kind: "primitive", PrimitiveCode: "nbt_le"}}}}}, {Value: 6, Name: "DataItemPosPayload", Shape: Shape{Kind: "struct", Semantic: "DataItemPosPayload", TypeID: "DataItemPosPayload", Fields: []ShapeField{{Ordinal: 0, Name: "Type", Shape: Shape{Kind: "enum", Semantic: "DataItemType", TypeID: "enums/DataItemType", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "Byte", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "Short", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "Int", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "Float", Shape: Shape{Kind: "void"}}, {Value: 4, Name: "String", Shape: Shape{Kind: "void"}}, {Value: 5, Name: "CompoundTag", Shape: Shape{Kind: "void"}}, {Value: 6, Name: "Pos", Shape: Shape{Kind: "void"}}, {Value: 7, Name: "Int64", Shape: Shape{Kind: "void"}}, {Value: 8, Name: "Vec3", Shape: Shape{Kind: "void"}}}}}, {Ordinal: 1, Name: "Value", Shape: Shape{Kind: "struct", Semantic: "BlockPos", TypeID: "BlockPos", Fields: []ShapeField{{Ordinal: 0, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 1, Name: "Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 2, Name: "Z", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}}}}}}}, {Value: 7, Name: "DataItemInt64Payload", Shape: Shape{Kind: "struct", Semantic: "DataItemInt64Payload", TypeID: "DataItemInt64Payload", Fields: []ShapeField{{Ordinal: 0, Name: "Type", Shape: Shape{Kind: "enum", Semantic: "DataItemType", TypeID: "enums/DataItemType", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "Byte", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "Short", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "Int", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "Float", Shape: Shape{Kind: "void"}}, {Value: 4, Name: "String", Shape: Shape{Kind: "void"}}, {Value: 5, Name: "CompoundTag", Shape: Shape{Kind: "void"}}, {Value: 6, Name: "Pos", Shape: Shape{Kind: "void"}}, {Value: 7, Name: "Int64", Shape: Shape{Kind: "void"}}, {Value: 8, Name: "Vec3", Shape: Shape{Kind: "void"}}}}}, {Ordinal: 1, Name: "Value", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i64"}}}}}, {Value: 8, Name: "DataItemVec3Payload", Shape: Shape{Kind: "struct", Semantic: "DataItemVec3Payload", TypeID: "DataItemVec3Payload", Fields: []ShapeField{{Ordinal: 0, Name: "Type", Shape: Shape{Kind: "enum", Semantic: "DataItemType", TypeID: "enums/DataItemType", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "Byte", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "Short", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "Int", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "Float", Shape: Shape{Kind: "void"}}, {Value: 4, Name: "String", Shape: Shape{Kind: "void"}}, {Value: 5, Name: "CompoundTag", Shape: Shape{Kind: "void"}}, {Value: 6, Name: "Pos", Shape: Shape{Kind: "void"}}, {Value: 7, Name: "Int64", Shape: Shape{Kind: "void"}}, {Value: 8, Name: "Vec3", Shape: Shape{Kind: "void"}}}}}, {Ordinal: 1, Name: "Value", Shape: Shape{Kind: "struct", Semantic: "Vec3", TypeID: "Vec3", Fields: []ShapeField{{Ordinal: 0, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 1, Name: "Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 2, Name: "Z", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}}}}}}}}}}}}}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(SynchedActorDataCopyableDataList)
		if !ok {
			return p, fmt.Errorf("field AddPlayerPacket.Entity Data has unexpected decoded type %T", raw)
		}
		p.EntityData = value
	}
	{
		raw, err := r.Read("AddPlayerPacket.Synched Properties", Shape{Kind: "struct", Semantic: "PropertySyncData", TypeID: "PropertySyncData", Fields: []ShapeField{{Ordinal: 0, Name: "Int Entries List", Shape: Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "struct", Semantic: "PropertySyncData::PropertySyncIntEntry", TypeID: "PropertySyncData::PropertySyncIntEntry", Fields: []ShapeField{{Ordinal: 0, Name: "Property Index", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, {Ordinal: 1, Name: "Data", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}}}}}, {Ordinal: 1, Name: "Float Entries List", Shape: Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "struct", Semantic: "PropertySyncData::PropertySyncFloatEntry", TypeID: "PropertySyncData::PropertySyncFloatEntry", Fields: []ShapeField{{Ordinal: 0, Name: "Property Index", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, {Ordinal: 1, Name: "Data", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}}}}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(PropertySyncData)
		if !ok {
			return p, fmt.Errorf("field AddPlayerPacket.Synched Properties has unexpected decoded type %T", raw)
		}
		p.SynchedProperties = value
	}
	{
		raw, err := r.Read("AddPlayerPacket.Abilities Data", Shape{Kind: "struct", Semantic: "SerializedAbilitiesData", TypeID: "SerializedAbilitiesData", Fields: []ShapeField{{Ordinal: 0, Name: "Target Player Raw Id", Shape: Shape{Kind: "primitive", PrimitiveCode: "i64le"}}, {Ordinal: 1, Name: "Player Permissions", Shape: Shape{Kind: "enum", Semantic: "PlayerPermissionLevel", TypeID: "enums/PlayerPermissionLevel", PrimitiveCode: "i8", Variants: []ShapeVariant{{Value: 0, Name: "Visitor", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "Member", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "Operator", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "Custom", Shape: Shape{Kind: "void"}}}}}, {Ordinal: 2, Name: "Command Permissions", Shape: Shape{Kind: "enum", Semantic: "CommandPermissionLevel", TypeID: "enums/CommandPermissionLevel", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "Any", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "GameDirectors", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "Admin", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "Host", Shape: Shape{Kind: "void"}}, {Value: 4, Name: "Owner", Shape: Shape{Kind: "void"}}, {Value: 5, Name: "Internal", Shape: Shape{Kind: "void"}}}}}, {Ordinal: 3, Name: "Layers", Shape: Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "struct", Semantic: "SerializedAbilitiesData::SerializedLayer", TypeID: "SerializedAbilitiesData::SerializedLayer", Fields: []ShapeField{{Ordinal: 0, Name: "SerializedLayer", Shape: Shape{Kind: "primitive", PrimitiveCode: "u16le"}}, {Ordinal: 1, Name: "AbilitiesSet", Shape: Shape{Kind: "primitive", PrimitiveCode: "u32le"}}, {Ordinal: 2, Name: "AbilityValues", Shape: Shape{Kind: "primitive", PrimitiveCode: "u32le"}}, {Ordinal: 3, Name: "FlySpeed", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 4, Name: "VerticalFlySpeed", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 5, Name: "WalkSpeed", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}}}}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(SerializedAbilitiesData)
		if !ok {
			return p, fmt.Errorf("field AddPlayerPacket.Abilities Data has unexpected decoded type %T", raw)
		}
		p.AbilitiesData = value
	}
	{
		raw, err := r.Read("AddPlayerPacket.Actor Links", Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "struct", Semantic: "ActorLink", TypeID: "ActorLink", Fields: []ShapeField{{Ordinal: 0, Name: "Target A", Shape: Shape{Kind: "struct", Semantic: "ActorUniqueID", TypeID: "ActorUniqueID", Fields: []ShapeField{{Ordinal: 0, Name: "Actor Unique ID", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i64"}}}}}, {Ordinal: 1, Name: "Target B", Shape: Shape{Kind: "struct", Semantic: "ActorUniqueID", TypeID: "ActorUniqueID", Fields: []ShapeField{{Ordinal: 0, Name: "Actor Unique ID", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i64"}}}}}, {Ordinal: 2, Name: "Type", Shape: Shape{Kind: "enum", Semantic: "ActorLinkType", TypeID: "enums/ActorLinkType", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "None", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "Riding", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "Passenger", Shape: Shape{Kind: "void"}}}}}, {Ordinal: 3, Name: "Immediate", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}, {Ordinal: 4, Name: "Passenger Initiated", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}, {Ordinal: 5, Name: "Vehicle Angular Velocity", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.([]ActorLink)
		if !ok {
			return p, fmt.Errorf("field AddPlayerPacket.Actor Links has unexpected decoded type %T", raw)
		}
		p.ActorLinks = value
	}
	{
		raw, err := r.Read("AddPlayerPacket.Device Id", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(string)
		if !ok {
			return p, fmt.Errorf("field AddPlayerPacket.Device Id has unexpected decoded type %T", raw)
		}
		p.DeviceId = value
	}
	{
		raw, err := r.Read("AddPlayerPacket.Build Platform", Shape{Kind: "enum", Semantic: "BuildPlatform", TypeID: "enums/BuildPlatform", PrimitiveCode: "i32le", Variants: []ShapeVariant{{Value: -1, Name: "Unknown", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "Google", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "iOS", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "OSX", Shape: Shape{Kind: "void"}}, {Value: 4, Name: "Amazon", Shape: Shape{Kind: "void"}}, {Value: 5, Name: "GearVR", Shape: Shape{Kind: "void"}}, {Value: 7, Name: "UWP", Shape: Shape{Kind: "void"}}, {Value: 8, Name: "Win32", Shape: Shape{Kind: "void"}}, {Value: 9, Name: "Dedicated", Shape: Shape{Kind: "void"}}, {Value: 10, Name: "tvOS", Shape: Shape{Kind: "void"}}, {Value: 11, Name: "Sony", Shape: Shape{Kind: "void"}}, {Value: 12, Name: "Nx", Shape: Shape{Kind: "void"}}, {Value: 13, Name: "Xbox", Shape: Shape{Kind: "void"}}, {Value: 14, Name: "WindowsPhone", Shape: Shape{Kind: "void"}}, {Value: 15, Name: "Linux", Shape: Shape{Kind: "void"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(BuildPlatform)
		if !ok {
			return p, fmt.Errorf("field AddPlayerPacket.Build Platform has unexpected decoded type %T", raw)
		}
		p.BuildPlatform = value
	}
	return p, nil
}
