// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import (
	"image/color"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/google/uuid"
)

// Marshal reads or writes ActorDataBoundingBoxComponent using its canonical wire layout.
func (x *ActorDataBoundingBoxComponent) Marshal(io IO) {
	for index1 := range x.ActorDataBoundingBox {
		io.Float32(&x.ActorDataBoundingBox[index1])
	}
}

// Marshal reads or writes ActorDataFlagComponent using its canonical wire layout.
func (x *ActorDataFlagComponent) Marshal(io IO) {
	io.Bitset(x.ActorFlagBitsetData[:], 131)
}

// Marshal reads or writes ActorLink using its canonical wire layout.
func (x *ActorLink) Marshal(io IO) {
	x.TargetA.Marshal(io)
	x.TargetB.Marshal(io)
	enumValue2 := uint8(x.Type)
	io.Uint8(&enumValue2)
	x.Type = ActorLinkType(enumValue2)
	switch int64(enumValue2) {
	case 0, 1, 2:
	default:
		io.InvalidValue(enumValue2, "unknown enum value")
	}
	io.Bool(&x.Immediate)
	io.Bool(&x.PassengerInitiated)
	io.Float32(&x.VehicleAngularVelocity)
}

// Marshal reads or writes ActorRuntimeID using its canonical wire layout.
func (x *ActorRuntimeID) Marshal(io IO) {
	io.Varuint64(&x.ActorRuntimeID)
}

// Marshal reads or writes ActorUniqueID using its canonical wire layout.
func (x *ActorUniqueID) Marshal(io IO) {
	io.Varint64(&x.ActorUniqueID)
}

// Marshal reads or writes AdventureSettings using its canonical wire layout.
func (x *AdventureSettings) Marshal(io IO) {
	io.Bool(&x.NoPvM)
	io.Bool(&x.NoMvP)
	io.Bool(&x.ImmutableWorld)
	io.Bool(&x.ShowNameTags)
	io.Bool(&x.AutoJump)
}

// Marshal reads or writes AgentCapabilities using its canonical wire layout.
func (x *AgentCapabilities) Marshal(io IO) {
	io.Bool(&x.CanModifyBlocks.set)
	if x.CanModifyBlocks.set {
		io.Bool(&x.CanModifyBlocks.val)
	} else if io.Reading() {
		var zero bool
		x.CanModifyBlocks.val = zero
	}
}

// Marshal reads or writes AnimatedImageData using its canonical wire layout.
func (x *AnimatedImageData) Marshal(io IO) {
	x.SkinImage.Marshal(io)
	enumValue3 := uint32(x.AnimatedTextureType)
	io.Varuint32(&enumValue3)
	x.AnimatedTextureType = PersonaAnimatedTextureType(enumValue3)
	switch int64(enumValue3) {
	case 1, 2, 3:
	default:
		io.InvalidValue(enumValue3, "unknown enum value")
	}
	io.Float32(&x.Frames)
	enumValue4 := uint32(x.AnimationExpression)
	io.Varuint32(&enumValue4)
	x.AnimationExpression = PersonaAnimationExpression(enumValue4)
	switch int64(enumValue4) {
	case 0, 1:
	default:
		io.InvalidValue(enumValue4, "unknown enum value")
	}
}

// Marshal reads or writes ArmorSlotAndDamagePair using its canonical wire layout.
func (x *ArmorSlotAndDamagePair) Marshal(io IO) {
	enumValue5 := int32(x.ArmorSlot)
	io.Varint32(&enumValue5)
	x.ArmorSlot = LegacyArmorSlot(enumValue5)
	switch int64(enumValue5) {
	case 0, 1, 2, 3, 4:
	default:
		io.InvalidValue(enumValue5, "unknown enum value")
	}
	io.Int16(&x.Damage)
}

// Marshal reads or writes ArrowData using its canonical wire layout.
func (x *ArrowData) Marshal(io IO) {
	io.Bool(&x.ArrowEndLocation.set)
	if x.ArrowEndLocation.set {
		io.Vec3(&x.ArrowEndLocation.val)
	} else if io.Reading() {
		var zero mgl32.Vec3
		x.ArrowEndLocation.val = zero
	}
	io.Bool(&x.ArrowHeadLength.set)
	if x.ArrowHeadLength.set {
		io.Float32(&x.ArrowHeadLength.val)
	} else if io.Reading() {
		var zero float32
		x.ArrowHeadLength.val = zero
	}
	io.Bool(&x.ArrowHeadRadius.set)
	if x.ArrowHeadRadius.set {
		io.Float32(&x.ArrowHeadRadius.val)
	} else if io.Reading() {
		var zero float32
		x.ArrowHeadRadius.val = zero
	}
	io.Bool(&x.NumSegments.set)
	if x.NumSegments.set {
		io.Uint8(&x.NumSegments.val)
	} else if io.Reading() {
		var zero uint8
		x.NumSegments.val = zero
	}
}

// Marshal reads or writes AttributeData using its canonical wire layout.
func (x *AttributeData) Marshal(io IO) {
	io.Float32(&x.MinValue)
	io.Float32(&x.MaxValue)
	io.Float32(&x.CurrentValue)
	io.Float32(&x.DefaultMinValue)
	io.Float32(&x.DefaultMaxValue)
	io.Float32(&x.DefaultValue)
	io.String(&x.Name)
	if !io.Reading() && uint64(len(x.Modifiers)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.Modifiers), "collection length overflows uint32")
		return
	}
	count6 := uint32(len(x.Modifiers))
	io.Varuint32(&count6)
	if io.Reading() {
		if uint64(count6) > uint64(^uint(0)>>1) {
			io.InvalidValue(count6, "collection length overflows int")
			return
		}
		x.Modifiers = make([]AttributeModifier, int(count6))
	}
	for index7 := range x.Modifiers {
		x.Modifiers[index7].Marshal(io)
	}
}

func marshalAttributeLayerSyncPacketData(io IO, x *AttributeLayerSyncPacketData) {
	if io.Reading() {
		var tag uint32
		io.Varuint32(&tag)
		switch int64(tag) {
		case 0:
			var value AttributeLayerSyncPacketDataUpdateAttributeLayersData
			value.Marshal(io)
			*x = value
		case 1:
			var value AttributeLayerSyncPacketDataUpdateAttributeLayerSettingsData
			value.Marshal(io)
			*x = value
		case 2:
			var value AttributeLayerSyncPacketDataUpdateEnvironmentAttributesData
			value.Marshal(io)
			*x = value
		case 3:
			var value AttributeLayerSyncPacketDataRemoveEnvironmentAttributesData
			value.Marshal(io)
			*x = value
		default:
			io.InvalidValue(tag, "unknown union tag")
		}
		return
	}
	switch value := (*x).(type) {
	case AttributeLayerSyncPacketDataUpdateAttributeLayersData:
		tag := uint32(0)
		io.Varuint32(&tag)
		value.Marshal(io)
	case AttributeLayerSyncPacketDataUpdateAttributeLayerSettingsData:
		tag := uint32(1)
		io.Varuint32(&tag)
		value.Marshal(io)
	case AttributeLayerSyncPacketDataUpdateEnvironmentAttributesData:
		tag := uint32(2)
		io.Varuint32(&tag)
		value.Marshal(io)
	case AttributeLayerSyncPacketDataRemoveEnvironmentAttributesData:
		tag := uint32(3)
		io.Varuint32(&tag)
		value.Marshal(io)
	default:
		io.InvalidValue(*x, "unknown union value")
	}
}

// Marshal reads or writes AttributeLayerSyncPacketDataRemoveEnvironmentAttributesData using its canonical wire layout.
func (x *AttributeLayerSyncPacketDataRemoveEnvironmentAttributesData) Marshal(io IO) {
	io.String(&x.AttributeLayerName)
	x.AttributeLayerDimension.Marshal(io)
	if !io.Reading() && uint64(len(x.Attributes)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.Attributes), "collection length overflows uint32")
		return
	}
	count8 := uint32(len(x.Attributes))
	io.Varuint32(&count8)
	if io.Reading() {
		if uint64(count8) > uint64(^uint(0)>>1) {
			io.InvalidValue(count8, "collection length overflows int")
			return
		}
		x.Attributes = make([]string, int(count8))
	}
	for index9 := range x.Attributes {
		io.String(&x.Attributes[index9])
	}
}

// Marshal reads or writes AttributeLayerSyncPacketDataUpdateAttributeLayerSettingsData using its canonical wire layout.
func (x *AttributeLayerSyncPacketDataUpdateAttributeLayerSettingsData) Marshal(io IO) {
	io.String(&x.AttributeLayerName)
	x.AttributeLayerDimension.Marshal(io)
	x.AttributesLayerSettings.Marshal(io)
}

// Marshal reads or writes AttributeLayerSyncPacketDataUpdateAttributeLayersData using its canonical wire layout.
func (x *AttributeLayerSyncPacketDataUpdateAttributeLayersData) Marshal(io IO) {
	if !io.Reading() && uint64(len(x.AttributeLayers)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.AttributeLayers), "collection length overflows uint32")
		return
	}
	count10 := uint32(len(x.AttributeLayers))
	io.Varuint32(&count10)
	if io.Reading() {
		if uint64(count10) > uint64(^uint(0)>>1) {
			io.InvalidValue(count10, "collection length overflows int")
			return
		}
		x.AttributeLayers = make([]EASAttributeLayerData, int(count10))
	}
	for index11 := range x.AttributeLayers {
		x.AttributeLayers[index11].Marshal(io)
	}
}

// Marshal reads or writes AttributeLayerSyncPacketDataUpdateEnvironmentAttributesData using its canonical wire layout.
func (x *AttributeLayerSyncPacketDataUpdateEnvironmentAttributesData) Marshal(io IO) {
	io.String(&x.AttributeLayerName)
	x.AttributeLayerDimension.Marshal(io)
	if !io.Reading() && uint64(len(x.Attributes)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.Attributes), "collection length overflows uint32")
		return
	}
	count12 := uint32(len(x.Attributes))
	io.Varuint32(&count12)
	if io.Reading() {
		if uint64(count12) > uint64(^uint(0)>>1) {
			io.InvalidValue(count12, "collection length overflows int")
			return
		}
		x.Attributes = make([]EASEnvironmentAttributeData, int(count12))
	}
	for index13 := range x.Attributes {
		x.Attributes[index13].Marshal(io)
	}
}

// Marshal reads or writes AttributeModifier using its canonical wire layout.
func (x *AttributeModifier) Marshal(io IO) {
	io.String(&x.Id)
	io.String(&x.Name)
	io.Float32(&x.Amount)
	io.Int32(&x.Operation)
	io.Int32(&x.Operand)
	io.Bool(&x.IsSerializable)
}

// Marshal reads or writes AvailableCommandsChainedSubcommandData using its canonical wire layout.
func (x *AvailableCommandsChainedSubcommandData) Marshal(io IO) {
	io.String(&x.Name)
	if !io.Reading() && uint64(len(x.SubCommandValues)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.SubCommandValues), "collection length overflows uint32")
		return
	}
	count14 := uint32(len(x.SubCommandValues))
	io.Varuint32(&count14)
	if io.Reading() {
		if uint64(count14) > uint64(^uint(0)>>1) {
			io.InvalidValue(count14, "collection length overflows int")
			return
		}
		x.SubCommandValues = make([]AvailableCommandsChainedSubcommandRelationship, int(count14))
	}
	for index15 := range x.SubCommandValues {
		x.SubCommandValues[index15].Marshal(io)
	}
}

// Marshal reads or writes AvailableCommandsChainedSubcommandRelationship using its canonical wire layout.
func (x *AvailableCommandsChainedSubcommandRelationship) Marshal(io IO) {
	io.Varuint32(&x.SubCommandFirstValue)
	io.Varuint32(&x.SubCommandSecondValue)
}

// Marshal reads or writes AvailableCommandsConstrainedValueData using its canonical wire layout.
func (x *AvailableCommandsConstrainedValueData) Marshal(io IO) {
	io.Uint32(&x.EnumValueSymbol)
	io.Uint32(&x.EnumSymbol)
	if !io.Reading() && uint64(len(x.ConstraintIndices)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.ConstraintIndices), "collection length overflows uint32")
		return
	}
	count16 := uint32(len(x.ConstraintIndices))
	io.Varuint32(&count16)
	if io.Reading() {
		if uint64(count16) > uint64(^uint(0)>>1) {
			io.InvalidValue(count16, "collection length overflows int")
			return
		}
		x.ConstraintIndices = make([]uint8, int(count16))
	}
	for index17 := range x.ConstraintIndices {
		io.Uint8(&x.ConstraintIndices[index17])
	}
}

// Marshal reads or writes AvailableCommandsEnumData using its canonical wire layout.
func (x *AvailableCommandsEnumData) Marshal(io IO) {
	io.String(&x.Name)
	if !io.Reading() && uint64(len(x.Values)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.Values), "collection length overflows uint32")
		return
	}
	count18 := uint32(len(x.Values))
	io.Varuint32(&count18)
	if io.Reading() {
		if uint64(count18) > uint64(^uint(0)>>1) {
			io.InvalidValue(count18, "collection length overflows int")
			return
		}
		x.Values = make([]uint32, int(count18))
	}
	for index19 := range x.Values {
		io.Uint32(&x.Values[index19])
	}
}

// Marshal reads or writes AvailableCommandsOverloadData using its canonical wire layout.
func (x *AvailableCommandsOverloadData) Marshal(io IO) {
	io.Bool(&x.IsChaining)
	if !io.Reading() && uint64(len(x.ParameterData)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.ParameterData), "collection length overflows uint32")
		return
	}
	count20 := uint32(len(x.ParameterData))
	io.Varuint32(&count20)
	if io.Reading() {
		if uint64(count20) > uint64(^uint(0)>>1) {
			io.InvalidValue(count20, "collection length overflows int")
			return
		}
		x.ParameterData = make([]AvailableCommandsParamData, int(count20))
	}
	for index21 := range x.ParameterData {
		x.ParameterData[index21].Marshal(io)
	}
}

// Marshal reads or writes AvailableCommandsPacketCommandData using its canonical wire layout.
func (x *AvailableCommandsPacketCommandData) Marshal(io IO) {
	io.String(&x.Name)
	io.String(&x.Description)
	io.Uint16(&x.Flags)
	io.String(&x.PermissionLevel)
	io.Int32(&x.AliasEnum)
	if !io.Reading() && uint64(len(x.CommandDataChainedSubcommandIndexes)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.CommandDataChainedSubcommandIndexes), "collection length overflows uint32")
		return
	}
	count22 := uint32(len(x.CommandDataChainedSubcommandIndexes))
	io.Varuint32(&count22)
	if io.Reading() {
		if uint64(count22) > uint64(^uint(0)>>1) {
			io.InvalidValue(count22, "collection length overflows int")
			return
		}
		x.CommandDataChainedSubcommandIndexes = make([]uint32, int(count22))
	}
	for index23 := range x.CommandDataChainedSubcommandIndexes {
		io.Uint32(&x.CommandDataChainedSubcommandIndexes[index23])
	}
	if !io.Reading() && uint64(len(x.Overloads)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.Overloads), "collection length overflows uint32")
		return
	}
	count24 := uint32(len(x.Overloads))
	io.Varuint32(&count24)
	if io.Reading() {
		if uint64(count24) > uint64(^uint(0)>>1) {
			io.InvalidValue(count24, "collection length overflows int")
			return
		}
		x.Overloads = make([]AvailableCommandsOverloadData, int(count24))
	}
	for index25 := range x.Overloads {
		x.Overloads[index25].Marshal(io)
	}
}

// Marshal reads or writes AvailableCommandsParamData using its canonical wire layout.
func (x *AvailableCommandsParamData) Marshal(io IO) {
	io.String(&x.Name)
	io.Uint32(&x.ParseSymbol)
	io.Bool(&x.IsOptional)
	io.Uint8(&x.Options)
}

// Marshal reads or writes AvailableCommandsSoftEnumData using its canonical wire layout.
func (x *AvailableCommandsSoftEnumData) Marshal(io IO) {
	io.String(&x.EnumName)
	if !io.Reading() && uint64(len(x.EnumOptions)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.EnumOptions), "collection length overflows uint32")
		return
	}
	count26 := uint32(len(x.EnumOptions))
	io.Varuint32(&count26)
	if io.Reading() {
		if uint64(count26) > uint64(^uint(0)>>1) {
			io.InvalidValue(count26, "collection length overflows int")
			return
		}
		x.EnumOptions = make([]string, int(count26))
	}
	for index27 := range x.EnumOptions {
		io.String(&x.EnumOptions[index27])
	}
}

func marshalBedrockDDUI(io IO, x *BedrockDDUI) {
	if io.Reading() {
		var tag uint32
		io.Varuint32(&tag)
		switch int64(tag) {
		case 0:
			var value BedrockDDUIDataStoreUpdate
			value.Marshal(io)
			*x = value
		case 1:
			var value BedrockDDUIDataStoreChange
			value.Marshal(io)
			*x = value
		case 2:
			var value BedrockDDUIDataStoreRemoval
			value.Marshal(io)
			*x = value
		default:
			io.InvalidValue(tag, "unknown union tag")
		}
		return
	}
	switch value := (*x).(type) {
	case BedrockDDUIDataStoreUpdate:
		tag := uint32(0)
		io.Varuint32(&tag)
		value.Marshal(io)
	case BedrockDDUIDataStoreChange:
		tag := uint32(1)
		io.Varuint32(&tag)
		value.Marshal(io)
	case BedrockDDUIDataStoreRemoval:
		tag := uint32(2)
		io.Varuint32(&tag)
		value.Marshal(io)
	default:
		io.InvalidValue(*x, "unknown union value")
	}
}

// Marshal reads or writes BedrockDDUIDataStoreChange using its canonical wire layout.
func (x *BedrockDDUIDataStoreChange) Marshal(io IO) {
	io.String(&x.DataStoreName)
	io.String(&x.Property)
	io.Uint32(&x.UpdateCount)
	marshalCerealDynamicValue(io, &x.TheNewPropertyValue)
}

// Marshal reads or writes BedrockDDUIDataStoreRemoval using its canonical wire layout.
func (x *BedrockDDUIDataStoreRemoval) Marshal(io IO) {
	io.String(&x.DataStoreName)
}

// Marshal reads or writes BedrockDDUIDataStoreUpdate using its canonical wire layout.
func (x *BedrockDDUIDataStoreUpdate) Marshal(io IO) {
	io.String(&x.DataStoreName)
	io.String(&x.Property)
	io.String(&x.Path)
	marshalBedrockDDUIDataStoreUpdateData(io, &x.Data)
	io.Uint32(&x.PropertyUpdateCount)
	io.Uint32(&x.PathUpdateCount)
}

func marshalBedrockDDUIDataStoreUpdateData(io IO, x *BedrockDDUIDataStoreUpdateData) {
	if io.Reading() {
		var tag uint32
		io.Varuint32(&tag)
		switch int64(tag) {
		case 0:
			var value BedrockDDUIDataStoreUpdateDataDouble
			value.Marshal(io)
			*x = value
		case 1:
			var value BedrockDDUIDataStoreUpdateDataBool
			value.Marshal(io)
			*x = value
		case 2:
			var value BedrockDDUIDataStoreUpdateDataString
			value.Marshal(io)
			*x = value
		default:
			io.InvalidValue(tag, "unknown union tag")
		}
		return
	}
	switch value := (*x).(type) {
	case BedrockDDUIDataStoreUpdateDataDouble:
		tag := uint32(0)
		io.Varuint32(&tag)
		value.Marshal(io)
	case BedrockDDUIDataStoreUpdateDataBool:
		tag := uint32(1)
		io.Varuint32(&tag)
		value.Marshal(io)
	case BedrockDDUIDataStoreUpdateDataString:
		tag := uint32(2)
		io.Varuint32(&tag)
		value.Marshal(io)
	default:
		io.InvalidValue(*x, "unknown union value")
	}
}

// Marshal reads or writes BedrockDDUIDataStoreUpdateDataBool using its canonical wire layout.
func (x *BedrockDDUIDataStoreUpdateDataBool) Marshal(io IO) {
	io.Bool(&x.Value)
}

// Marshal reads or writes BedrockDDUIDataStoreUpdateDataDouble using its canonical wire layout.
func (x *BedrockDDUIDataStoreUpdateDataDouble) Marshal(io IO) {
	io.Float64(&x.Value)
}

// Marshal reads or writes BedrockDDUIDataStoreUpdateDataString using its canonical wire layout.
func (x *BedrockDDUIDataStoreUpdateDataString) Marshal(io IO) {
	io.String(&x.Value)
}

// Marshal reads or writes BedrockProfileWhiskerDiagnosticsScopeDataSummary using its canonical wire layout.
func (x *BedrockProfileWhiskerDiagnosticsScopeDataSummary) Marshal(io IO) {
	io.String(&x.Label)
	io.String(&x.Indentation)
	io.Uint64(&x.TotalHighCostNS)
	io.Uint64(&x.TotalMidCostNS)
	io.Uint64(&x.TotalLowCostNS)
}

// Marshal reads or writes BedrockSafetyRedactableString using its canonical wire layout.
func (x *BedrockSafetyRedactableString) Marshal(io IO) {
	io.String(&x.Unredacted)
	io.Bool(&x.Redacted.set)
	if x.Redacted.set {
		io.String(&x.Redacted.val)
	} else if io.Reading() {
		var zero string
		x.Redacted.val = zero
	}
}

// Marshal reads or writes BiomeCappedSurfaceData using its canonical wire layout.
func (x *BiomeCappedSurfaceData) Marshal(io IO) {
	if !io.Reading() && uint64(len(x.FloorBlocks)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.FloorBlocks), "collection length overflows uint32")
		return
	}
	count28 := uint32(len(x.FloorBlocks))
	io.Varuint32(&count28)
	if io.Reading() {
		if uint64(count28) > uint64(^uint(0)>>1) {
			io.InvalidValue(count28, "collection length overflows int")
			return
		}
		x.FloorBlocks = make([]uint32, int(count28))
	}
	for index29 := range x.FloorBlocks {
		io.Uint32(&x.FloorBlocks[index29])
	}
	if !io.Reading() && uint64(len(x.CeilingBlocks)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.CeilingBlocks), "collection length overflows uint32")
		return
	}
	count30 := uint32(len(x.CeilingBlocks))
	io.Varuint32(&count30)
	if io.Reading() {
		if uint64(count30) > uint64(^uint(0)>>1) {
			io.InvalidValue(count30, "collection length overflows int")
			return
		}
		x.CeilingBlocks = make([]uint32, int(count30))
	}
	for index31 := range x.CeilingBlocks {
		io.Uint32(&x.CeilingBlocks[index31])
	}
	io.Bool(&x.SeaBlock.set)
	if x.SeaBlock.set {
		io.Uint32(&x.SeaBlock.val)
	} else if io.Reading() {
		var zero uint32
		x.SeaBlock.val = zero
	}
	io.Bool(&x.FoundationBlock.set)
	if x.FoundationBlock.set {
		io.Uint32(&x.FoundationBlock.val)
	} else if io.Reading() {
		var zero uint32
		x.FoundationBlock.val = zero
	}
	io.Bool(&x.BeachBlock.set)
	if x.BeachBlock.set {
		io.Uint32(&x.BeachBlock.val)
	} else if io.Reading() {
		var zero uint32
		x.BeachBlock.val = zero
	}
}

// Marshal reads or writes BiomeClimateData using its canonical wire layout.
func (x *BiomeClimateData) Marshal(io IO) {
	io.Float32(&x.Temperature)
	io.Float32(&x.Downfall)
	io.Float32(&x.SnowAccumulationMin)
	io.Float32(&x.SnowAccumulationMax)
}

// Marshal reads or writes BiomeConditionalTransformationData using its canonical wire layout.
func (x *BiomeConditionalTransformationData) Marshal(io IO) {
	if !io.Reading() && uint64(len(x.TransformsInto)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.TransformsInto), "collection length overflows uint32")
		return
	}
	count32 := uint32(len(x.TransformsInto))
	io.Varuint32(&count32)
	if io.Reading() {
		if uint64(count32) > uint64(^uint(0)>>1) {
			io.InvalidValue(count32, "collection length overflows int")
			return
		}
		x.TransformsInto = make([]BiomeWeightedData, int(count32))
	}
	for index33 := range x.TransformsInto {
		x.TransformsInto[index33].Marshal(io)
	}
	io.Uint16(&x.ConditionJson)
	io.Uint32(&x.MinPassingNeighbors)
}

// Marshal reads or writes BiomeConsolidatedFeatureData using its canonical wire layout.
func (x *BiomeConsolidatedFeatureData) Marshal(io IO) {
	x.Scatter.Marshal(io)
	io.Uint16(&x.Feature)
	io.Uint16(&x.Identifier)
	io.Uint16(&x.Pass)
	io.Bool(&x.CanUseInternalFeature)
}

// Marshal reads or writes BiomeConsolidatedFeaturesData using its canonical wire layout.
func (x *BiomeConsolidatedFeaturesData) Marshal(io IO) {
	if !io.Reading() && uint64(len(x.Features)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.Features), "collection length overflows uint32")
		return
	}
	count34 := uint32(len(x.Features))
	io.Varuint32(&count34)
	if io.Reading() {
		if uint64(count34) > uint64(^uint(0)>>1) {
			io.InvalidValue(count34, "collection length overflows int")
			return
		}
		x.Features = make([]BiomeConsolidatedFeatureData, int(count34))
	}
	for index35 := range x.Features {
		x.Features[index35].Marshal(io)
	}
}

// Marshal reads or writes BiomeCoordinateData using its canonical wire layout.
func (x *BiomeCoordinateData) Marshal(io IO) {
	io.Varint32(&x.MinValueType)
	io.Uint16(&x.MinValue)
	io.Varint32(&x.MaxValueType)
	io.Uint16(&x.MaxValue)
	io.Uint32(&x.GridOffset)
	io.Uint32(&x.GridStepSize)
	enumValue36 := int32(x.Distribution)
	io.Varint32(&enumValue36)
	x.Distribution = RandomDistributionType(enumValue36)
	switch int64(enumValue36) {
	case 0, 1, 2, 3, 4, 5, 6:
	default:
		io.InvalidValue(enumValue36, "unknown enum value")
	}
}

// Marshal reads or writes BiomeDefinitionChunkGenData using its canonical wire layout.
func (x *BiomeDefinitionChunkGenData) Marshal(io IO) {
	io.Bool(&x.Climate.set)
	if x.Climate.set {
		x.Climate.val.Marshal(io)
	} else if io.Reading() {
		var zero BiomeClimateData
		x.Climate.val = zero
	}
	io.Bool(&x.ConsolidatedFeatures.set)
	if x.ConsolidatedFeatures.set {
		x.ConsolidatedFeatures.val.Marshal(io)
	} else if io.Reading() {
		var zero BiomeConsolidatedFeaturesData
		x.ConsolidatedFeatures.val = zero
	}
	io.Bool(&x.MountainParams.set)
	if x.MountainParams.set {
		x.MountainParams.val.Marshal(io)
	} else if io.Reading() {
		var zero BiomeMountainParamsData
		x.MountainParams.val = zero
	}
	io.Bool(&x.SurfaceMaterialAdjustments.set)
	if x.SurfaceMaterialAdjustments.set {
		x.SurfaceMaterialAdjustments.val.Marshal(io)
	} else if io.Reading() {
		var zero BiomeSurfaceMaterialAdjustmentData
		x.SurfaceMaterialAdjustments.val = zero
	}
	io.Bool(&x.OverworldGenRules.set)
	if x.OverworldGenRules.set {
		x.OverworldGenRules.val.Marshal(io)
	} else if io.Reading() {
		var zero BiomeOverworldGenRulesData
		x.OverworldGenRules.val = zero
	}
	io.Bool(&x.MultinoiseGenRules.set)
	if x.MultinoiseGenRules.set {
		x.MultinoiseGenRules.val.Marshal(io)
	} else if io.Reading() {
		var zero BiomeMultinoiseGenRulesData
		x.MultinoiseGenRules.val = zero
	}
	io.Bool(&x.LegacyWorldGenRules.set)
	if x.LegacyWorldGenRules.set {
		x.LegacyWorldGenRules.val.Marshal(io)
	} else if io.Reading() {
		var zero BiomeLegacyWorldGenRulesData
		x.LegacyWorldGenRules.val = zero
	}
	io.Bool(&x.ReplacementBiomes.set)
	if x.ReplacementBiomes.set {
		x.ReplacementBiomes.val.Marshal(io)
	} else if io.Reading() {
		var zero BiomeReplacementsData
		x.ReplacementBiomes.val = zero
	}
	io.Bool(&x.VillageType.set)
	if x.VillageType.set {
		enumValue37 := uint8(x.VillageType.val)
		io.Uint8(&enumValue37)
		x.VillageType.val = VillageType(enumValue37)
		switch int64(enumValue37) {
		case 0, 1, 2, 3, 4:
		default:
			io.InvalidValue(enumValue37, "unknown enum value")
		}
	} else if io.Reading() {
		var zero VillageType
		x.VillageType.val = zero
	}
	io.Bool(&x.SurfaceBuilderData.set)
	if x.SurfaceBuilderData.set {
		x.SurfaceBuilderData.val.Marshal(io)
	} else if io.Reading() {
		var zero BiomeSurfaceBuilderData
		x.SurfaceBuilderData.val = zero
	}
	io.Bool(&x.SubsurfaceBuilderData.set)
	if x.SubsurfaceBuilderData.set {
		x.SubsurfaceBuilderData.val.Marshal(io)
	} else if io.Reading() {
		var zero BiomeSurfaceBuilderData
		x.SubsurfaceBuilderData.val = zero
	}
}

// Marshal reads or writes BiomeDefinitionData using its canonical wire layout.
func (x *BiomeDefinitionData) Marshal(io IO) {
	io.Uint16(&x.Id)
	io.Float32(&x.Temperature)
	io.Float32(&x.Downfall)
	io.Float32(&x.FoliageSnow)
	io.Float32(&x.Depth)
	io.Float32(&x.Scale)
	io.Int32(&x.MapWaterColorARGB)
	io.Bool(&x.Rain)
	io.Bool(&x.Tags.set)
	if x.Tags.set {
		x.Tags.val.Marshal(io)
	} else if io.Reading() {
		var zero BiomeTagsData
		x.Tags.val = zero
	}
	io.Bool(&x.ChunkGenData.set)
	if x.ChunkGenData.set {
		x.ChunkGenData.val.Marshal(io)
	} else if io.Reading() {
		var zero BiomeDefinitionChunkGenData
		x.ChunkGenData.val = zero
	}
}

// Marshal reads or writes BiomeElementData using its canonical wire layout.
func (x *BiomeElementData) Marshal(io IO) {
	io.Float32(&x.NoiseFreqScale)
	io.Float32(&x.NoiseLowerBound)
	io.Float32(&x.NoiseUpperBound)
	io.Varint32(&x.HeightMinType)
	io.Uint16(&x.HeightMin)
	io.Varint32(&x.HeightMaxType)
	io.Uint16(&x.HeightMax)
	x.AdjustedMaterials.Marshal(io)
}

// Marshal reads or writes BiomeLegacyWorldGenRulesData using its canonical wire layout.
func (x *BiomeLegacyWorldGenRulesData) Marshal(io IO) {
	if !io.Reading() && uint64(len(x.LegacyPreHillsEdge)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.LegacyPreHillsEdge), "collection length overflows uint32")
		return
	}
	count38 := uint32(len(x.LegacyPreHillsEdge))
	io.Varuint32(&count38)
	if io.Reading() {
		if uint64(count38) > uint64(^uint(0)>>1) {
			io.InvalidValue(count38, "collection length overflows int")
			return
		}
		x.LegacyPreHillsEdge = make([]BiomeConditionalTransformationData, int(count38))
	}
	for index39 := range x.LegacyPreHillsEdge {
		x.LegacyPreHillsEdge[index39].Marshal(io)
	}
}

// Marshal reads or writes BiomeMesaSurfaceData using its canonical wire layout.
func (x *BiomeMesaSurfaceData) Marshal(io IO) {
	io.Uint32(&x.ClayMaterial)
	io.Uint32(&x.HardClayMaterial)
	io.Bool(&x.BrycePillars)
	io.Bool(&x.HasForest)
}

// Marshal reads or writes BiomeMountainParamsData using its canonical wire layout.
func (x *BiomeMountainParamsData) Marshal(io IO) {
	io.Uint32(&x.SteepBlock)
	io.Bool(&x.NorthSlopes)
	io.Bool(&x.SouthSlopes)
	io.Bool(&x.WestSlopes)
	io.Bool(&x.EastSlopes)
	io.Bool(&x.TopSlideEnabled)
}

// Marshal reads or writes BiomeMultinoiseGenRulesData using its canonical wire layout.
func (x *BiomeMultinoiseGenRulesData) Marshal(io IO) {
	io.Float32(&x.Temperature)
	io.Float32(&x.Humidity)
	io.Float32(&x.Altitude)
	io.Float32(&x.Weirdness)
	io.Float32(&x.Weight)
}

// Marshal reads or writes BiomeNoiseGradientSurfaceData using its canonical wire layout.
func (x *BiomeNoiseGradientSurfaceData) Marshal(io IO) {
	if !io.Reading() && uint64(len(x.NonReplaceableBlocks)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.NonReplaceableBlocks), "collection length overflows uint32")
		return
	}
	count40 := uint32(len(x.NonReplaceableBlocks))
	io.Varuint32(&count40)
	if io.Reading() {
		if uint64(count40) > uint64(^uint(0)>>1) {
			io.InvalidValue(count40, "collection length overflows int")
			return
		}
		x.NonReplaceableBlocks = make([]uint32, int(count40))
	}
	for index41 := range x.NonReplaceableBlocks {
		io.Uint32(&x.NonReplaceableBlocks[index41])
	}
	if !io.Reading() && uint64(len(x.GradientBlocks)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.GradientBlocks), "collection length overflows uint32")
		return
	}
	count42 := uint32(len(x.GradientBlocks))
	io.Varuint32(&count42)
	if io.Reading() {
		if uint64(count42) > uint64(^uint(0)>>1) {
			io.InvalidValue(count42, "collection length overflows int")
			return
		}
		x.GradientBlocks = make([]SerializedNoiseBlockSpecifier, int(count42))
	}
	for index43 := range x.GradientBlocks {
		x.GradientBlocks[index43].Marshal(io)
	}
	x.Noise.Marshal(io)
}

// Marshal reads or writes BiomeOverworldGenRulesData using its canonical wire layout.
func (x *BiomeOverworldGenRulesData) Marshal(io IO) {
	if !io.Reading() && uint64(len(x.HillsTransformations)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.HillsTransformations), "collection length overflows uint32")
		return
	}
	count44 := uint32(len(x.HillsTransformations))
	io.Varuint32(&count44)
	if io.Reading() {
		if uint64(count44) > uint64(^uint(0)>>1) {
			io.InvalidValue(count44, "collection length overflows int")
			return
		}
		x.HillsTransformations = make([]BiomeWeightedData, int(count44))
	}
	for index45 := range x.HillsTransformations {
		x.HillsTransformations[index45].Marshal(io)
	}
	if !io.Reading() && uint64(len(x.MutateTransformations)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.MutateTransformations), "collection length overflows uint32")
		return
	}
	count46 := uint32(len(x.MutateTransformations))
	io.Varuint32(&count46)
	if io.Reading() {
		if uint64(count46) > uint64(^uint(0)>>1) {
			io.InvalidValue(count46, "collection length overflows int")
			return
		}
		x.MutateTransformations = make([]BiomeWeightedData, int(count46))
	}
	for index47 := range x.MutateTransformations {
		x.MutateTransformations[index47].Marshal(io)
	}
	if !io.Reading() && uint64(len(x.RiverTransformations)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.RiverTransformations), "collection length overflows uint32")
		return
	}
	count48 := uint32(len(x.RiverTransformations))
	io.Varuint32(&count48)
	if io.Reading() {
		if uint64(count48) > uint64(^uint(0)>>1) {
			io.InvalidValue(count48, "collection length overflows int")
			return
		}
		x.RiverTransformations = make([]BiomeWeightedData, int(count48))
	}
	for index49 := range x.RiverTransformations {
		x.RiverTransformations[index49].Marshal(io)
	}
	if !io.Reading() && uint64(len(x.ShoreTransformations)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.ShoreTransformations), "collection length overflows uint32")
		return
	}
	count50 := uint32(len(x.ShoreTransformations))
	io.Varuint32(&count50)
	if io.Reading() {
		if uint64(count50) > uint64(^uint(0)>>1) {
			io.InvalidValue(count50, "collection length overflows int")
			return
		}
		x.ShoreTransformations = make([]BiomeWeightedData, int(count50))
	}
	for index51 := range x.ShoreTransformations {
		x.ShoreTransformations[index51].Marshal(io)
	}
	if !io.Reading() && uint64(len(x.PreHillsEdge)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.PreHillsEdge), "collection length overflows uint32")
		return
	}
	count52 := uint32(len(x.PreHillsEdge))
	io.Varuint32(&count52)
	if io.Reading() {
		if uint64(count52) > uint64(^uint(0)>>1) {
			io.InvalidValue(count52, "collection length overflows int")
			return
		}
		x.PreHillsEdge = make([]BiomeConditionalTransformationData, int(count52))
	}
	for index53 := range x.PreHillsEdge {
		x.PreHillsEdge[index53].Marshal(io)
	}
	if !io.Reading() && uint64(len(x.PostShoreEdge)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.PostShoreEdge), "collection length overflows uint32")
		return
	}
	count54 := uint32(len(x.PostShoreEdge))
	io.Varuint32(&count54)
	if io.Reading() {
		if uint64(count54) > uint64(^uint(0)>>1) {
			io.InvalidValue(count54, "collection length overflows int")
			return
		}
		x.PostShoreEdge = make([]BiomeConditionalTransformationData, int(count54))
	}
	for index55 := range x.PostShoreEdge {
		x.PostShoreEdge[index55].Marshal(io)
	}
	if !io.Reading() && uint64(len(x.Climate)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.Climate), "collection length overflows uint32")
		return
	}
	count56 := uint32(len(x.Climate))
	io.Varuint32(&count56)
	if io.Reading() {
		if uint64(count56) > uint64(^uint(0)>>1) {
			io.InvalidValue(count56, "collection length overflows int")
			return
		}
		x.Climate = make([]BiomeWeightedTemperatureData, int(count56))
	}
	for index57 := range x.Climate {
		x.Climate[index57].Marshal(io)
	}
}

// Marshal reads or writes BiomeReplacementData using its canonical wire layout.
func (x *BiomeReplacementData) Marshal(io IO) {
	io.Uint16(&x.ReplacementBiome)
	io.Uint16(&x.Dimension)
	if !io.Reading() && uint64(len(x.TargetBiomes)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.TargetBiomes), "collection length overflows uint32")
		return
	}
	count58 := uint32(len(x.TargetBiomes))
	io.Varuint32(&count58)
	if io.Reading() {
		if uint64(count58) > uint64(^uint(0)>>1) {
			io.InvalidValue(count58, "collection length overflows int")
			return
		}
		x.TargetBiomes = make([]uint16, int(count58))
	}
	for index59 := range x.TargetBiomes {
		io.Uint16(&x.TargetBiomes[index59])
	}
	io.Float32(&x.Amount)
	io.Float32(&x.NoiseFrequencyScale)
	io.Uint32(&x.ReplacementIndex)
}

// Marshal reads or writes BiomeReplacementsData using its canonical wire layout.
func (x *BiomeReplacementsData) Marshal(io IO) {
	if !io.Reading() && uint64(len(x.BiomeReplacements)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.BiomeReplacements), "collection length overflows uint32")
		return
	}
	count60 := uint32(len(x.BiomeReplacements))
	io.Varuint32(&count60)
	if io.Reading() {
		if uint64(count60) > uint64(^uint(0)>>1) {
			io.InvalidValue(count60, "collection length overflows int")
			return
		}
		x.BiomeReplacements = make([]BiomeReplacementData, int(count60))
	}
	for index61 := range x.BiomeReplacements {
		x.BiomeReplacements[index61].Marshal(io)
	}
}

// Marshal reads or writes BiomeScatterParamData using its canonical wire layout.
func (x *BiomeScatterParamData) Marshal(io IO) {
	if !io.Reading() && uint64(len(x.Coordinates)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.Coordinates), "collection length overflows uint32")
		return
	}
	count62 := uint32(len(x.Coordinates))
	io.Varuint32(&count62)
	if io.Reading() {
		if uint64(count62) > uint64(^uint(0)>>1) {
			io.InvalidValue(count62, "collection length overflows int")
			return
		}
		x.Coordinates = make([]BiomeCoordinateData, int(count62))
	}
	for index63 := range x.Coordinates {
		x.Coordinates[index63].Marshal(io)
	}
	enumValue64 := int32(x.EvalOrder)
	io.Varint32(&enumValue64)
	x.EvalOrder = CoordinateEvaluationOrder(enumValue64)
	switch int64(enumValue64) {
	case 0, 1, 2, 3, 4, 5:
	default:
		io.InvalidValue(enumValue64, "unknown enum value")
	}
	io.Varint32(&x.ChancePercentType)
	io.Uint16(&x.ChancePercent)
	io.Int32(&x.ChanceNumerator)
	io.Int32(&x.ChanceDenominator)
	io.Varint32(&x.IterationsType)
	io.Uint16(&x.Iterations)
}

// Marshal reads or writes BiomeStringList using its canonical wire layout.
func (x *BiomeStringList) Marshal(io IO) {
	if !io.Reading() && uint64(len(x.Strings)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.Strings), "collection length overflows uint32")
		return
	}
	count65 := uint32(len(x.Strings))
	io.Varuint32(&count65)
	if io.Reading() {
		if uint64(count65) > uint64(^uint(0)>>1) {
			io.InvalidValue(count65, "collection length overflows int")
			return
		}
		x.Strings = make([]string, int(count65))
	}
	for index66 := range x.Strings {
		io.String(&x.Strings[index66])
	}
}

// Marshal reads or writes BiomeSurfaceBuilderData using its canonical wire layout.
func (x *BiomeSurfaceBuilderData) Marshal(io IO) {
	io.Bool(&x.SurfaceMaterials.set)
	if x.SurfaceMaterials.set {
		x.SurfaceMaterials.val.Marshal(io)
	} else if io.Reading() {
		var zero BiomeSurfaceMaterialData
		x.SurfaceMaterials.val = zero
	}
	io.Bool(&x.HasDefaultOverworldSurface)
	io.Bool(&x.HasSwampSurface)
	io.Bool(&x.HasFrozenOceanSurface)
	io.Bool(&x.HasTheEndSurface)
	io.Bool(&x.MesaSurface.set)
	if x.MesaSurface.set {
		x.MesaSurface.val.Marshal(io)
	} else if io.Reading() {
		var zero BiomeMesaSurfaceData
		x.MesaSurface.val = zero
	}
	io.Bool(&x.CappedSurface.set)
	if x.CappedSurface.set {
		x.CappedSurface.val.Marshal(io)
	} else if io.Reading() {
		var zero BiomeCappedSurfaceData
		x.CappedSurface.val = zero
	}
	io.Bool(&x.NoiseGradientSurface.set)
	if x.NoiseGradientSurface.set {
		x.NoiseGradientSurface.val.Marshal(io)
	} else if io.Reading() {
		var zero BiomeNoiseGradientSurfaceData
		x.NoiseGradientSurface.val = zero
	}
}

// Marshal reads or writes BiomeSurfaceMaterialAdjustmentData using its canonical wire layout.
func (x *BiomeSurfaceMaterialAdjustmentData) Marshal(io IO) {
	if !io.Reading() && uint64(len(x.Adjustments)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.Adjustments), "collection length overflows uint32")
		return
	}
	count67 := uint32(len(x.Adjustments))
	io.Varuint32(&count67)
	if io.Reading() {
		if uint64(count67) > uint64(^uint(0)>>1) {
			io.InvalidValue(count67, "collection length overflows int")
			return
		}
		x.Adjustments = make([]BiomeElementData, int(count67))
	}
	for index68 := range x.Adjustments {
		x.Adjustments[index68].Marshal(io)
	}
}

// Marshal reads or writes BiomeSurfaceMaterialData using its canonical wire layout.
func (x *BiomeSurfaceMaterialData) Marshal(io IO) {
	io.Uint32(&x.TopBlock)
	io.Uint32(&x.MidBlock)
	io.Uint32(&x.SeaFloorBlock)
	io.Uint32(&x.FoundationBlock)
	io.Uint32(&x.SeaBlock)
	io.Int32(&x.SeaFloorDepth)
}

// Marshal reads or writes BiomeTagsData using its canonical wire layout.
func (x *BiomeTagsData) Marshal(io IO) {
	if !io.Reading() && uint64(len(x.Tags)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.Tags), "collection length overflows uint32")
		return
	}
	count69 := uint32(len(x.Tags))
	io.Varuint32(&count69)
	if io.Reading() {
		if uint64(count69) > uint64(^uint(0)>>1) {
			io.InvalidValue(count69, "collection length overflows int")
			return
		}
		x.Tags = make([]uint16, int(count69))
	}
	for index70 := range x.Tags {
		io.Uint16(&x.Tags[index70])
	}
}

// Marshal reads or writes BiomeWeightedData using its canonical wire layout.
func (x *BiomeWeightedData) Marshal(io IO) {
	io.Uint16(&x.BiomeIdentifier)
	io.Uint32(&x.Weight)
}

// Marshal reads or writes BiomeWeightedTemperatureData using its canonical wire layout.
func (x *BiomeWeightedTemperatureData) Marshal(io IO) {
	io.Varint32(&x.Temperature)
	io.Uint32(&x.Weight)
}

// Marshal reads or writes BlockPos using its canonical wire layout.
func (x *BlockPos) Marshal(io IO) {
	io.Varint32(&x.X)
	io.Varint32(&x.Y)
	io.Varint32(&x.Z)
}

func marshalBookEditAction(io IO, x *BookEditAction) {
	if io.Reading() {
		var tag uint32
		io.Varuint32(&tag)
		switch int64(tag) {
		case 0:
			var value BookEditActionReplacePage
			value.Marshal(io)
			*x = value
		case 1:
			var value BookEditActionAddPage
			value.Marshal(io)
			*x = value
		case 2:
			var value BookEditActionDeletePage
			value.Marshal(io)
			*x = value
		case 3:
			var value BookEditActionSwapPages
			value.Marshal(io)
			*x = value
		case 4:
			var value BookEditActionFinalize
			value.Marshal(io)
			*x = value
		default:
			io.InvalidValue(tag, "unknown union tag")
		}
		return
	}
	switch value := (*x).(type) {
	case BookEditActionReplacePage:
		tag := uint32(0)
		io.Varuint32(&tag)
		value.Marshal(io)
	case BookEditActionAddPage:
		tag := uint32(1)
		io.Varuint32(&tag)
		value.Marshal(io)
	case BookEditActionDeletePage:
		tag := uint32(2)
		io.Varuint32(&tag)
		value.Marshal(io)
	case BookEditActionSwapPages:
		tag := uint32(3)
		io.Varuint32(&tag)
		value.Marshal(io)
	case BookEditActionFinalize:
		tag := uint32(4)
		io.Varuint32(&tag)
		value.Marshal(io)
	default:
		io.InvalidValue(*x, "unknown union value")
	}
}

// Marshal reads or writes BookEditActionAddPage using its canonical wire layout.
func (x *BookEditActionAddPage) Marshal(io IO) {
	io.Varint32(&x.PageIndex)
	io.String(&x.PageText)
	io.String(&x.PhotoName)
}

// Marshal reads or writes BookEditActionDeletePage using its canonical wire layout.
func (x *BookEditActionDeletePage) Marshal(io IO) {
	io.Varint32(&x.PageIndex)
}

// Marshal reads or writes BookEditActionFinalize using its canonical wire layout.
func (x *BookEditActionFinalize) Marshal(io IO) {
	io.String(&x.Title)
	io.String(&x.Author)
	io.String(&x.XUID)
}

// Marshal reads or writes BookEditActionReplacePage using its canonical wire layout.
func (x *BookEditActionReplacePage) Marshal(io IO) {
	io.Varint32(&x.PageIndex)
	io.String(&x.PageText)
	io.String(&x.PhotoName)
}

// Marshal reads or writes BookEditActionSwapPages using its canonical wire layout.
func (x *BookEditActionSwapPages) Marshal(io IO) {
	io.Varint32(&x.PageIndex)
	io.Varint32(&x.SwapWithIndex)
}

// Marshal reads or writes BoxData using its canonical wire layout.
func (x *BoxData) Marshal(io IO) {
	io.Vec3(&x.BoxBound)
}

// Marshal reads or writes CameraAimAssistActorPriorityPriorityData using its canonical wire layout.
func (x *CameraAimAssistActorPriorityPriorityData) Marshal(io IO) {
	io.Int32(&x.PresetIndex)
	io.Int32(&x.CategoryIndex)
	io.Int32(&x.ActorIndex)
	io.Int32(&x.PriorityValue)
}

// Marshal reads or writes CameraAimAssistCategoryDefinition using its canonical wire layout.
func (x *CameraAimAssistCategoryDefinition) Marshal(io IO) {
	io.String(&x.Name)
	x.Priorities.Marshal(io)
}

// Marshal reads or writes CameraAimAssistCategoryPriorities using its canonical wire layout.
func (x *CameraAimAssistCategoryPriorities) Marshal(io IO) {
	if !io.Reading() && uint64(len(x.Entities)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.Entities), "map length overflows uint32")
		return
	}
	count71 := uint32(len(x.Entities))
	io.Varuint32(&count71)
	if io.Reading() {
		if uint64(count71) > uint64(^uint(0)>>1) {
			io.InvalidValue(count71, "map length overflows int")
			return
		}
		x.Entities = make([]OrderedEntry[string, int32], int(count71))
	}
	for index72 := range x.Entities {
		io.String(&x.Entities[index72].Key)
		io.Int32(&x.Entities[index72].Value)
	}
	if !io.Reading() && uint64(len(x.Blocks)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.Blocks), "map length overflows uint32")
		return
	}
	count73 := uint32(len(x.Blocks))
	io.Varuint32(&count73)
	if io.Reading() {
		if uint64(count73) > uint64(^uint(0)>>1) {
			io.InvalidValue(count73, "map length overflows int")
			return
		}
		x.Blocks = make([]OrderedEntry[string, int32], int(count73))
	}
	for index74 := range x.Blocks {
		io.String(&x.Blocks[index74].Key)
		io.Int32(&x.Blocks[index74].Value)
	}
	if !io.Reading() && uint64(len(x.BlockTags)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.BlockTags), "map length overflows uint32")
		return
	}
	count75 := uint32(len(x.BlockTags))
	io.Varuint32(&count75)
	if io.Reading() {
		if uint64(count75) > uint64(^uint(0)>>1) {
			io.InvalidValue(count75, "map length overflows int")
			return
		}
		x.BlockTags = make([]OrderedEntry[string, int32], int(count75))
	}
	for index76 := range x.BlockTags {
		io.String(&x.BlockTags[index76].Key)
		io.Int32(&x.BlockTags[index76].Value)
	}
	if !io.Reading() && uint64(len(x.EntityTypeFamilies)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.EntityTypeFamilies), "map length overflows uint32")
		return
	}
	count77 := uint32(len(x.EntityTypeFamilies))
	io.Varuint32(&count77)
	if io.Reading() {
		if uint64(count77) > uint64(^uint(0)>>1) {
			io.InvalidValue(count77, "map length overflows int")
			return
		}
		x.EntityTypeFamilies = make([]OrderedEntry[string, int32], int(count77))
	}
	for index78 := range x.EntityTypeFamilies {
		io.String(&x.EntityTypeFamilies[index78].Key)
		io.Int32(&x.EntityTypeFamilies[index78].Value)
	}
	io.Bool(&x.EntityDefault.set)
	if x.EntityDefault.set {
		io.Int32(&x.EntityDefault.val)
	} else if io.Reading() {
		var zero int32
		x.EntityDefault.val = zero
	}
	io.Bool(&x.BlockDefault.set)
	if x.BlockDefault.set {
		io.Int32(&x.BlockDefault.val)
	} else if io.Reading() {
		var zero int32
		x.BlockDefault.val = zero
	}
}

// Marshal reads or writes CameraAimAssistCommandPresetDefinition using its canonical wire layout.
func (x *CameraAimAssistCommandPresetDefinition) Marshal(io IO) {
	io.Bool(&x.PresetId.set)
	if x.PresetId.set {
		io.String(&x.PresetId.val)
	} else if io.Reading() {
		var zero string
		x.PresetId.val = zero
	}
	io.Bool(&x.TargetMode.set)
	if x.TargetMode.set {
		enumValue79 := uint8(x.TargetMode.val)
		io.Uint8(&enumValue79)
		x.TargetMode.val = CameraAimAssistTargetMode(enumValue79)
		switch int64(enumValue79) {
		case 0, 1:
		default:
			io.InvalidValue(enumValue79, "unknown enum value")
		}
	} else if io.Reading() {
		var zero CameraAimAssistTargetMode
		x.TargetMode.val = zero
	}
	io.Bool(&x.ViewAngle.set)
	if x.ViewAngle.set {
		io.Vec2(&x.ViewAngle.val)
	} else if io.Reading() {
		var zero mgl32.Vec2
		x.ViewAngle.val = zero
	}
	io.Bool(&x.Distance.set)
	if x.Distance.set {
		io.Float32(&x.Distance.val)
	} else if io.Reading() {
		var zero float32
		x.Distance.val = zero
	}
}

// Marshal reads or writes CameraAimAssistPresetDefinition using its canonical wire layout.
func (x *CameraAimAssistPresetDefinition) Marshal(io IO) {
	io.String(&x.Identifier)
	x.ExclusionSettings.Marshal(io)
	if !io.Reading() && uint64(len(x.LiquidTargetingList)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.LiquidTargetingList), "collection length overflows uint32")
		return
	}
	count80 := uint32(len(x.LiquidTargetingList))
	io.Varuint32(&count80)
	if io.Reading() {
		if uint64(count80) > uint64(^uint(0)>>1) {
			io.InvalidValue(count80, "collection length overflows int")
			return
		}
		x.LiquidTargetingList = make([]string, int(count80))
	}
	for index81 := range x.LiquidTargetingList {
		io.String(&x.LiquidTargetingList[index81])
	}
	if !io.Reading() && uint64(len(x.ItemSettings)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.ItemSettings), "map length overflows uint32")
		return
	}
	count82 := uint32(len(x.ItemSettings))
	io.Varuint32(&count82)
	if io.Reading() {
		if uint64(count82) > uint64(^uint(0)>>1) {
			io.InvalidValue(count82, "map length overflows int")
			return
		}
		x.ItemSettings = make([]OrderedEntry[string, string], int(count82))
	}
	for index83 := range x.ItemSettings {
		io.String(&x.ItemSettings[index83].Key)
		io.String(&x.ItemSettings[index83].Value)
	}
	io.Bool(&x.DefaultItemSettings.set)
	if x.DefaultItemSettings.set {
		io.String(&x.DefaultItemSettings.val)
	} else if io.Reading() {
		var zero string
		x.DefaultItemSettings.val = zero
	}
	io.Bool(&x.HandSettings.set)
	if x.HandSettings.set {
		io.String(&x.HandSettings.val)
	} else if io.Reading() {
		var zero string
		x.HandSettings.val = zero
	}
}

// Marshal reads or writes CameraAimAssistPresetExclusionDefinition using its canonical wire layout.
func (x *CameraAimAssistPresetExclusionDefinition) Marshal(io IO) {
	if !io.Reading() && uint64(len(x.Blocks)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.Blocks), "collection length overflows uint32")
		return
	}
	count84 := uint32(len(x.Blocks))
	io.Varuint32(&count84)
	if io.Reading() {
		if uint64(count84) > uint64(^uint(0)>>1) {
			io.InvalidValue(count84, "collection length overflows int")
			return
		}
		x.Blocks = make([]string, int(count84))
	}
	for index85 := range x.Blocks {
		io.String(&x.Blocks[index85])
	}
	if !io.Reading() && uint64(len(x.Entities)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.Entities), "collection length overflows uint32")
		return
	}
	count86 := uint32(len(x.Entities))
	io.Varuint32(&count86)
	if io.Reading() {
		if uint64(count86) > uint64(^uint(0)>>1) {
			io.InvalidValue(count86, "collection length overflows int")
			return
		}
		x.Entities = make([]string, int(count86))
	}
	for index87 := range x.Entities {
		io.String(&x.Entities[index87])
	}
	if !io.Reading() && uint64(len(x.BlockTags)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.BlockTags), "collection length overflows uint32")
		return
	}
	count88 := uint32(len(x.BlockTags))
	io.Varuint32(&count88)
	if io.Reading() {
		if uint64(count88) > uint64(^uint(0)>>1) {
			io.InvalidValue(count88, "collection length overflows int")
			return
		}
		x.BlockTags = make([]string, int(count88))
	}
	for index89 := range x.BlockTags {
		io.String(&x.BlockTags[index89])
	}
	if !io.Reading() && uint64(len(x.EntityTypeFamilies)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.EntityTypeFamilies), "collection length overflows uint32")
		return
	}
	count90 := uint32(len(x.EntityTypeFamilies))
	io.Varuint32(&count90)
	if io.Reading() {
		if uint64(count90) > uint64(^uint(0)>>1) {
			io.InvalidValue(count90, "collection length overflows int")
			return
		}
		x.EntityTypeFamilies = make([]string, int(count90))
	}
	for index91 := range x.EntityTypeFamilies {
		io.String(&x.EntityTypeFamilies[index91])
	}
}

// Marshal reads or writes CameraInstructionData using its canonical wire layout.
func (x *CameraInstructionData) Marshal(io IO) {
	io.Bool(&x.Set.set)
	if x.Set.set {
		x.Set.val.Marshal(io)
	} else if io.Reading() {
		var zero CameraInstructionOptionsSetInstruction
		x.Set.val = zero
	}
	io.Bool(&x.Clear.set)
	if x.Clear.set {
		io.Bool(&x.Clear.val)
	} else if io.Reading() {
		var zero bool
		x.Clear.val = zero
	}
	io.Bool(&x.Fade.set)
	if x.Fade.set {
		x.Fade.val.Marshal(io)
	} else if io.Reading() {
		var zero CameraInstructionOptionsFadeInstruction
		x.Fade.val = zero
	}
	io.Bool(&x.Target.set)
	if x.Target.set {
		x.Target.val.Marshal(io)
	} else if io.Reading() {
		var zero CameraInstructionOptionsTargetInstruction
		x.Target.val = zero
	}
	io.Bool(&x.RemoveTarget.set)
	if x.RemoveTarget.set {
		io.Bool(&x.RemoveTarget.val)
	} else if io.Reading() {
		var zero bool
		x.RemoveTarget.val = zero
	}
	io.Bool(&x.FieldOfView.set)
	if x.FieldOfView.set {
		x.FieldOfView.val.Marshal(io)
	} else if io.Reading() {
		var zero CameraInstructionOptionsFovInstruction
		x.FieldOfView.val = zero
	}
	io.Bool(&x.Spline.set)
	if x.Spline.set {
		x.Spline.val.Marshal(io)
	} else if io.Reading() {
		var zero CameraInstructionOptionsSplineInstruction
		x.Spline.val = zero
	}
	io.Bool(&x.AttachToEntity.set)
	if x.AttachToEntity.set {
		x.AttachToEntity.val.Marshal(io)
	} else if io.Reading() {
		var zero CameraInstructionOptionsAttachToEntityInstruction
		x.AttachToEntity.val = zero
	}
	io.Bool(&x.DetachFromEntity.set)
	if x.DetachFromEntity.set {
		io.Bool(&x.DetachFromEntity.val)
	} else if io.Reading() {
		var zero bool
		x.DetachFromEntity.val = zero
	}
}

// Marshal reads or writes CameraInstructionOptionsAttachToEntityInstruction using its canonical wire layout.
func (x *CameraInstructionOptionsAttachToEntityInstruction) Marshal(io IO) {
	io.Int64(&x.EntityActorID)
}

// Marshal reads or writes CameraInstructionOptionsFadeInstruction using its canonical wire layout.
func (x *CameraInstructionOptionsFadeInstruction) Marshal(io IO) {
	io.Bool(&x.Time.set)
	if x.Time.set {
		x.Time.val.Marshal(io)
	} else if io.Reading() {
		var zero CameraInstructionOptionsFadeInstructionTimeOption
		x.Time.val = zero
	}
	io.Bool(&x.Color.set)
	if x.Color.set {
		x.Color.val.Marshal(io)
	} else if io.Reading() {
		var zero CameraInstructionOptionsFadeInstructionColorOption
		x.Color.val = zero
	}
}

// Marshal reads or writes CameraInstructionOptionsFadeInstructionColorOption using its canonical wire layout.
func (x *CameraInstructionOptionsFadeInstructionColorOption) Marshal(io IO) {
	io.Float32(&x.Red)
	io.Float32(&x.Green)
	io.Float32(&x.Blue)
}

// Marshal reads or writes CameraInstructionOptionsFadeInstructionTimeOption using its canonical wire layout.
func (x *CameraInstructionOptionsFadeInstructionTimeOption) Marshal(io IO) {
	io.Float32(&x.FadeInTime)
	io.Float32(&x.HoldTime)
	io.Float32(&x.FadeOutTime)
}

// Marshal reads or writes CameraInstructionOptionsFovInstruction using its canonical wire layout.
func (x *CameraInstructionOptionsFovInstruction) Marshal(io IO) {
	io.Float32(&x.FieldOfView)
	io.Float32(&x.FOVEaseTime)
	io.String(&x.FOVEaseType)
	io.Bool(&x.FieldOfViewClear)
}

// Marshal reads or writes CameraInstructionOptionsSetInstruction using its canonical wire layout.
func (x *CameraInstructionOptionsSetInstruction) Marshal(io IO) {
	io.Uint32(&x.Preset)
	io.Bool(&x.Ease.set)
	if x.Ease.set {
		x.Ease.val.Marshal(io)
	} else if io.Reading() {
		var zero CameraInstructionOptionsSetInstructionEaseOption
		x.Ease.val = zero
	}
	io.Bool(&x.Pos.set)
	if x.Pos.set {
		x.Pos.val.Marshal(io)
	} else if io.Reading() {
		var zero CameraInstructionOptionsSetInstructionPosOption
		x.Pos.val = zero
	}
	io.Bool(&x.Rot.set)
	if x.Rot.set {
		x.Rot.val.Marshal(io)
	} else if io.Reading() {
		var zero CameraInstructionOptionsSetInstructionRotOption
		x.Rot.val = zero
	}
	io.Bool(&x.Facing.set)
	if x.Facing.set {
		x.Facing.val.Marshal(io)
	} else if io.Reading() {
		var zero CameraInstructionOptionsSetInstructionFacingOption
		x.Facing.val = zero
	}
	io.Bool(&x.ViewOffset.set)
	if x.ViewOffset.set {
		x.ViewOffset.val.Marshal(io)
	} else if io.Reading() {
		var zero CameraInstructionOptionsSetInstructionViewOffsetOption
		x.ViewOffset.val = zero
	}
	io.Bool(&x.EntityOffset.set)
	if x.EntityOffset.set {
		x.EntityOffset.val.Marshal(io)
	} else if io.Reading() {
		var zero CameraInstructionOptionsSetInstructionEntityOffsetOption
		x.EntityOffset.val = zero
	}
	io.Bool(&x.Default.set)
	if x.Default.set {
		io.Bool(&x.Default.val)
	} else if io.Reading() {
		var zero bool
		x.Default.val = zero
	}
	io.Bool(&x.RemoveIgnoreStartingValuesComponent)
}

// Marshal reads or writes CameraInstructionOptionsSetInstructionEaseOption using its canonical wire layout.
func (x *CameraInstructionOptionsSetInstructionEaseOption) Marshal(io IO) {
	io.Uint8(&x.Type)
	io.Float32(&x.Time)
}

// Marshal reads or writes CameraInstructionOptionsSetInstructionEntityOffsetOption using its canonical wire layout.
func (x *CameraInstructionOptionsSetInstructionEntityOffsetOption) Marshal(io IO) {
	io.Float32(&x.EntityOffsetX)
	io.Float32(&x.EntityOffsetY)
	io.Float32(&x.EntityOffsetZ)
}

// Marshal reads or writes CameraInstructionOptionsSetInstructionFacingOption using its canonical wire layout.
func (x *CameraInstructionOptionsSetInstructionFacingOption) Marshal(io IO) {
	io.Vec3(&x.Pos)
}

// Marshal reads or writes CameraInstructionOptionsSetInstructionPosOption using its canonical wire layout.
func (x *CameraInstructionOptionsSetInstructionPosOption) Marshal(io IO) {
	io.Vec3(&x.Pos)
}

// Marshal reads or writes CameraInstructionOptionsSetInstructionRotOption using its canonical wire layout.
func (x *CameraInstructionOptionsSetInstructionRotOption) Marshal(io IO) {
	io.Float32(&x.X)
	io.Float32(&x.Y)
}

// Marshal reads or writes CameraInstructionOptionsSetInstructionViewOffsetOption using its canonical wire layout.
func (x *CameraInstructionOptionsSetInstructionViewOffsetOption) Marshal(io IO) {
	io.Float32(&x.X)
	io.Float32(&x.Y)
}

// Marshal reads or writes CameraInstructionOptionsSplineInstruction using its canonical wire layout.
func (x *CameraInstructionOptionsSplineInstruction) Marshal(io IO) {
	io.Float32(&x.TotalTime)
	io.Uint8(&x.Type)
	if !io.Reading() && uint64(len(x.Curve)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.Curve), "collection length overflows uint32")
		return
	}
	count92 := uint32(len(x.Curve))
	io.Varuint32(&count92)
	if io.Reading() {
		if uint64(count92) > uint64(^uint(0)>>1) {
			io.InvalidValue(count92, "collection length overflows int")
			return
		}
		x.Curve = make([]mgl32.Vec3, int(count92))
	}
	for index93 := range x.Curve {
		io.Vec3(&x.Curve[index93])
	}
	if !io.Reading() && uint64(len(x.ProgressKeyFrames)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.ProgressKeyFrames), "collection length overflows uint32")
		return
	}
	count94 := uint32(len(x.ProgressKeyFrames))
	io.Varuint32(&count94)
	if io.Reading() {
		if uint64(count94) > uint64(^uint(0)>>1) {
			io.InvalidValue(count94, "collection length overflows int")
			return
		}
		x.ProgressKeyFrames = make([]CameraInstructionOptionsSplineInstructionSplineProgressOption, int(count94))
	}
	for index95 := range x.ProgressKeyFrames {
		x.ProgressKeyFrames[index95].Marshal(io)
	}
	if !io.Reading() && uint64(len(x.RotationOption)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.RotationOption), "collection length overflows uint32")
		return
	}
	count96 := uint32(len(x.RotationOption))
	io.Varuint32(&count96)
	if io.Reading() {
		if uint64(count96) > uint64(^uint(0)>>1) {
			io.InvalidValue(count96, "collection length overflows int")
			return
		}
		x.RotationOption = make([]CameraInstructionOptionsSplineInstructionSplineRotationOption, int(count96))
	}
	for index97 := range x.RotationOption {
		x.RotationOption[index97].Marshal(io)
	}
	io.String(&x.SplineIdentifier)
	io.Bool(&x.LoadFromJson)
}

// Marshal reads or writes CameraInstructionOptionsSplineInstructionSplineProgressOption using its canonical wire layout.
func (x *CameraInstructionOptionsSplineInstructionSplineProgressOption) Marshal(io IO) {
	io.Float32(&x.KeyFrameValue)
	io.Float32(&x.KeyFrameTime)
	io.String(&x.KeyFrameEasingFunc)
}

// Marshal reads or writes CameraInstructionOptionsSplineInstructionSplineRotationOption using its canonical wire layout.
func (x *CameraInstructionOptionsSplineInstructionSplineRotationOption) Marshal(io IO) {
	io.Vec3(&x.KeyFrameValue)
	io.Float32(&x.KeyFrameTime)
	io.String(&x.KeyFrameEasingFunc)
}

// Marshal reads or writes CameraInstructionOptionsTargetInstruction using its canonical wire layout.
func (x *CameraInstructionOptionsTargetInstruction) Marshal(io IO) {
	io.Bool(&x.TargetCenterOffset.set)
	if x.TargetCenterOffset.set {
		io.Vec3(&x.TargetCenterOffset.val)
	} else if io.Reading() {
		var zero mgl32.Vec3
		x.TargetCenterOffset.val = zero
	}
	io.Int64(&x.TargetActorID)
}

// Marshal reads or writes CameraPreset using its canonical wire layout.
func (x *CameraPreset) Marshal(io IO) {
	io.String(&x.Name)
	io.String(&x.InheritFrom)
	io.Bool(&x.PosX.set)
	if x.PosX.set {
		io.Float32(&x.PosX.val)
	} else if io.Reading() {
		var zero float32
		x.PosX.val = zero
	}
	io.Bool(&x.PosY.set)
	if x.PosY.set {
		io.Float32(&x.PosY.val)
	} else if io.Reading() {
		var zero float32
		x.PosY.val = zero
	}
	io.Bool(&x.PosZ.set)
	if x.PosZ.set {
		io.Float32(&x.PosZ.val)
	} else if io.Reading() {
		var zero float32
		x.PosZ.val = zero
	}
	io.Bool(&x.RotX.set)
	if x.RotX.set {
		io.Float32(&x.RotX.val)
	} else if io.Reading() {
		var zero float32
		x.RotX.val = zero
	}
	io.Bool(&x.RotY.set)
	if x.RotY.set {
		io.Float32(&x.RotY.val)
	} else if io.Reading() {
		var zero float32
		x.RotY.val = zero
	}
	io.Bool(&x.RotationSpeed.set)
	if x.RotationSpeed.set {
		io.Float32(&x.RotationSpeed.val)
	} else if io.Reading() {
		var zero float32
		x.RotationSpeed.val = zero
	}
	io.Bool(&x.SnapToTarget.set)
	if x.SnapToTarget.set {
		io.Bool(&x.SnapToTarget.val)
	} else if io.Reading() {
		var zero bool
		x.SnapToTarget.val = zero
	}
	io.Bool(&x.HorizontalRotationLimit.set)
	if x.HorizontalRotationLimit.set {
		io.Vec2(&x.HorizontalRotationLimit.val)
	} else if io.Reading() {
		var zero mgl32.Vec2
		x.HorizontalRotationLimit.val = zero
	}
	io.Bool(&x.VerticalRotationLimit.set)
	if x.VerticalRotationLimit.set {
		io.Vec2(&x.VerticalRotationLimit.val)
	} else if io.Reading() {
		var zero mgl32.Vec2
		x.VerticalRotationLimit.val = zero
	}
	io.Bool(&x.ContinueTargeting.set)
	if x.ContinueTargeting.set {
		io.Bool(&x.ContinueTargeting.val)
	} else if io.Reading() {
		var zero bool
		x.ContinueTargeting.val = zero
	}
	io.Bool(&x.BlockListeningRadius.set)
	if x.BlockListeningRadius.set {
		io.Float32(&x.BlockListeningRadius.val)
	} else if io.Reading() {
		var zero float32
		x.BlockListeningRadius.val = zero
	}
	io.Bool(&x.ViewOffset.set)
	if x.ViewOffset.set {
		io.Vec2(&x.ViewOffset.val)
	} else if io.Reading() {
		var zero mgl32.Vec2
		x.ViewOffset.val = zero
	}
	io.Bool(&x.EntityOffset.set)
	if x.EntityOffset.set {
		io.Vec3(&x.EntityOffset.val)
	} else if io.Reading() {
		var zero mgl32.Vec3
		x.EntityOffset.val = zero
	}
	io.Bool(&x.Radius.set)
	if x.Radius.set {
		io.Float32(&x.Radius.val)
	} else if io.Reading() {
		var zero float32
		x.Radius.val = zero
	}
	io.Bool(&x.YawLimitMin.set)
	if x.YawLimitMin.set {
		io.Float32(&x.YawLimitMin.val)
	} else if io.Reading() {
		var zero float32
		x.YawLimitMin.val = zero
	}
	io.Bool(&x.YawLimitMax.set)
	if x.YawLimitMax.set {
		io.Float32(&x.YawLimitMax.val)
	} else if io.Reading() {
		var zero float32
		x.YawLimitMax.val = zero
	}
	io.Bool(&x.Listener.set)
	if x.Listener.set {
		enumValue98 := uint8(x.Listener.val)
		io.Uint8(&enumValue98)
		x.Listener.val = CameraPresetAudioListener(enumValue98)
		switch int64(enumValue98) {
		case 0, 1:
		default:
			io.InvalidValue(enumValue98, "unknown enum value")
		}
	} else if io.Reading() {
		var zero CameraPresetAudioListener
		x.Listener.val = zero
	}
	io.Bool(&x.PlayerEffects.set)
	if x.PlayerEffects.set {
		io.Bool(&x.PlayerEffects.val)
	} else if io.Reading() {
		var zero bool
		x.PlayerEffects.val = zero
	}
	io.Bool(&x.AimAssist.set)
	if x.AimAssist.set {
		x.AimAssist.val.Marshal(io)
	} else if io.Reading() {
		var zero CameraAimAssistCommandPresetDefinition
		x.AimAssist.val = zero
	}
	io.Bool(&x.ControlScheme.set)
	if x.ControlScheme.set {
		enumValue99 := uint8(x.ControlScheme.val)
		io.Uint8(&enumValue99)
		x.ControlScheme.val = ControlSchemeScheme(enumValue99)
		switch int64(enumValue99) {
		case 0, 1, 2, 3, 4:
		default:
			io.InvalidValue(enumValue99, "unknown enum value")
		}
	} else if io.Reading() {
		var zero ControlSchemeScheme
		x.ControlScheme.val = zero
	}
}

// Marshal reads or writes CameraPresetsData using its canonical wire layout.
func (x *CameraPresetsData) Marshal(io IO) {
	if !io.Reading() && uint64(len(x.Presets)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.Presets), "collection length overflows uint32")
		return
	}
	count100 := uint32(len(x.Presets))
	io.Varuint32(&count100)
	if io.Reading() {
		if uint64(count100) > uint64(^uint(0)>>1) {
			io.InvalidValue(count100, "collection length overflows int")
			return
		}
		x.Presets = make([]CameraPreset, int(count100))
	}
	for index101 := range x.Presets {
		x.Presets[index101].Marshal(io)
	}
}

// Marshal reads or writes CameraSplineControlPoint using its canonical wire layout.
func (x *CameraSplineControlPoint) Marshal(io IO) {
	io.Vec3(&x.Position)
}

// Marshal reads or writes CameraSplineDefinition using its canonical wire layout.
func (x *CameraSplineDefinition) Marshal(io IO) {
	io.String(&x.Name)
	io.Float32(&x.TotalTime)
	io.String(&x.SplineType)
	if !io.Reading() && uint64(len(x.ControlPoints)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.ControlPoints), "collection length overflows uint32")
		return
	}
	count102 := uint32(len(x.ControlPoints))
	io.Varuint32(&count102)
	if io.Reading() {
		if uint64(count102) > uint64(^uint(0)>>1) {
			io.InvalidValue(count102, "collection length overflows int")
			return
		}
		x.ControlPoints = make([]CameraSplineControlPoint, int(count102))
	}
	for index103 := range x.ControlPoints {
		x.ControlPoints[index103].Marshal(io)
	}
	if !io.Reading() && uint64(len(x.ProgressKeyFrames)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.ProgressKeyFrames), "collection length overflows uint32")
		return
	}
	count104 := uint32(len(x.ProgressKeyFrames))
	io.Varuint32(&count104)
	if io.Reading() {
		if uint64(count104) > uint64(^uint(0)>>1) {
			io.InvalidValue(count104, "collection length overflows int")
			return
		}
		x.ProgressKeyFrames = make([]CameraSplineProgressKeyFrame, int(count104))
	}
	for index105 := range x.ProgressKeyFrames {
		x.ProgressKeyFrames[index105].Marshal(io)
	}
	if !io.Reading() && uint64(len(x.RotationKeyFrames)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.RotationKeyFrames), "collection length overflows uint32")
		return
	}
	count106 := uint32(len(x.RotationKeyFrames))
	io.Varuint32(&count106)
	if io.Reading() {
		if uint64(count106) > uint64(^uint(0)>>1) {
			io.InvalidValue(count106, "collection length overflows int")
			return
		}
		x.RotationKeyFrames = make([]CameraSplineRotationKeyFrame, int(count106))
	}
	for index107 := range x.RotationKeyFrames {
		x.RotationKeyFrames[index107].Marshal(io)
	}
}

// Marshal reads or writes CameraSplineProgressKeyFrame using its canonical wire layout.
func (x *CameraSplineProgressKeyFrame) Marshal(io IO) {
	io.Float32(&x.Progress)
	io.Float32(&x.Time)
	io.Bool(&x.Easing.set)
	if x.Easing.set {
		io.String(&x.Easing.val)
	} else if io.Reading() {
		var zero string
		x.Easing.val = zero
	}
}

// Marshal reads or writes CameraSplineRotationKeyFrame using its canonical wire layout.
func (x *CameraSplineRotationKeyFrame) Marshal(io IO) {
	io.Vec3(&x.Rotation)
	io.Float32(&x.Time)
	io.Bool(&x.Easing.set)
	if x.Easing.set {
		io.String(&x.Easing.val)
	} else if io.Reading() {
		var zero string
		x.Easing.val = zero
	}
}

func marshalCerealDynamicValue(io IO, x *CerealDynamicValue) {
	if io.Reading() {
		var tag int32
		io.Int32(&tag)
		switch int64(tag) {
		case 0:
			var value CerealDynamicValueNone
			value.Marshal(io)
			*x = value
		case 1:
			var value CerealDynamicValueBool
			value.Marshal(io)
			*x = value
		case 2:
			var value CerealDynamicValueInt64
			value.Marshal(io)
			*x = value
		case 3:
			var value CerealDynamicValueDouble
			value.Marshal(io)
			*x = value
		case 4:
			var value CerealDynamicValueString
			value.Marshal(io)
			*x = value
		case 5:
			var value CerealDynamicValueList
			value.Marshal(io)
			*x = value
		case 6:
			var value CerealDynamicValueMap
			value.Marshal(io)
			*x = value
		default:
			io.InvalidValue(tag, "unknown union tag")
		}
		return
	}
	switch value := (*x).(type) {
	case CerealDynamicValueNone:
		tag := int32(0)
		io.Int32(&tag)
		value.Marshal(io)
	case CerealDynamicValueBool:
		tag := int32(1)
		io.Int32(&tag)
		value.Marshal(io)
	case CerealDynamicValueInt64:
		tag := int32(2)
		io.Int32(&tag)
		value.Marshal(io)
	case CerealDynamicValueDouble:
		tag := int32(3)
		io.Int32(&tag)
		value.Marshal(io)
	case CerealDynamicValueString:
		tag := int32(4)
		io.Int32(&tag)
		value.Marshal(io)
	case CerealDynamicValueList:
		tag := int32(5)
		io.Int32(&tag)
		value.Marshal(io)
	case CerealDynamicValueMap:
		tag := int32(6)
		io.Int32(&tag)
		value.Marshal(io)
	default:
		io.InvalidValue(*x, "unknown union value")
	}
}

// Marshal reads or writes CerealDynamicValueBool using its canonical wire layout.
func (x *CerealDynamicValueBool) Marshal(io IO) {
	io.Bool(&x.Value)
}

// Marshal reads or writes CerealDynamicValueDouble using its canonical wire layout.
func (x *CerealDynamicValueDouble) Marshal(io IO) {
	io.Float64(&x.Value)
}

// Marshal reads or writes CerealDynamicValueInt64 using its canonical wire layout.
func (x *CerealDynamicValueInt64) Marshal(io IO) {
	io.Int64(&x.Value)
}

// Marshal reads or writes CerealDynamicValueList using its canonical wire layout.
func (x *CerealDynamicValueList) Marshal(io IO) {
	if !io.Reading() && uint64(len(x.Value)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.Value), "collection length overflows uint32")
		return
	}
	count108 := uint32(len(x.Value))
	io.Varuint32(&count108)
	if io.Reading() {
		if uint64(count108) > uint64(^uint(0)>>1) {
			io.InvalidValue(count108, "collection length overflows int")
			return
		}
		x.Value = make([]CerealDynamicValue, int(count108))
	}
	for index109 := range x.Value {
		marshalCerealDynamicValue(io, &x.Value[index109])
	}
}

// Marshal reads or writes CerealDynamicValueMap using its canonical wire layout.
func (x *CerealDynamicValueMap) Marshal(io IO) {
	if !io.Reading() && uint64(len(x.Value)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.Value), "map length overflows uint32")
		return
	}
	count110 := uint32(len(x.Value))
	io.Varuint32(&count110)
	if io.Reading() {
		if uint64(count110) > uint64(^uint(0)>>1) {
			io.InvalidValue(count110, "map length overflows int")
			return
		}
		x.Value = make([]OrderedEntry[string, CerealDynamicValue], int(count110))
	}
	for index111 := range x.Value {
		io.String(&x.Value[index111].Key)
		marshalCerealDynamicValue(io, &x.Value[index111].Value)
	}
}

// Marshal reads or writes CerealDynamicValueNone using its canonical wire layout.
func (x *CerealDynamicValueNone) Marshal(io IO) {
}

// Marshal reads or writes CerealDynamicValueString using its canonical wire layout.
func (x *CerealDynamicValueString) Marshal(io IO) {
	io.String(&x.Value)
}

// Marshal reads or writes CerealizerExperimentsAnonExperimentToggle using its canonical wire layout.
func (x *CerealizerExperimentsAnonExperimentToggle) Marshal(io IO) {
	io.String(&x.Name)
	io.Bool(&x.Enabled)
}

// Marshal reads or writes CerealizerNetworkItemInstanceDescriptorSerializedData using its canonical wire layout.
func (x *CerealizerNetworkItemInstanceDescriptorSerializedData) Marshal(io IO) {
	io.Varint32(&x.Id)
	io.Uint16(&x.StackSize)
	io.Varuint32(&x.AuxValue)
	io.Varint32(&x.BlockRuntimeId)
	io.Bytes(&x.UserDataBuffer)
}

// Marshal reads or writes CerealizerNetworkItemStackDescriptorSerializedData using its canonical wire layout.
func (x *CerealizerNetworkItemStackDescriptorSerializedData) Marshal(io IO) {
	io.Int16(&x.Id)
	io.Uint16(&x.StackSize)
	io.Varuint32(&x.AuxValue)
	io.Bool(&x.NetIdVariant.set)
	if x.NetIdVariant.set {
		io.Varint32(&x.NetIdVariant.val)
	} else if io.Reading() {
		var zero int32
		x.NetIdVariant.val = zero
	}
	io.Varuint32(&x.BlockRuntimeId)
	io.Bytes(&x.UserDataBuffer)
}

// Marshal reads or writes CerealizerRecipeIngredientSerializedData using its canonical wire layout.
func (x *CerealizerRecipeIngredientSerializedData) Marshal(io IO) {
	if !io.Reading() && uint64(len(x.Descriptor)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.Descriptor), "map length overflows uint32")
		return
	}
	count112 := uint32(len(x.Descriptor))
	io.Varuint32(&count112)
	if io.Reading() {
		if uint64(count112) > uint64(^uint(0)>>1) {
			io.InvalidValue(count112, "map length overflows int")
			return
		}
		x.Descriptor = make([]OrderedEntry[string, string], int(count112))
	}
	for index113 := range x.Descriptor {
		io.String(&x.Descriptor[index113].Key)
		io.String(&x.Descriptor[index113].Value)
	}
	io.Varint32(&x.AuxValue)
	io.Varint32(&x.StackSize)
}

// Marshal reads or writes CerealizerRecipeUnlockingRequirementSerializedData using its canonical wire layout.
func (x *CerealizerRecipeUnlockingRequirementSerializedData) Marshal(io IO) {
	enumValue114 := int32(x.UnlockingContext)
	io.Varint32(&enumValue114)
	x.UnlockingContext = RecipeUnlockingRequirementUnlockingContext(enumValue114)
	switch int64(enumValue114) {
	case 0, 1, 2, 3:
	default:
		io.InvalidValue(enumValue114, "unknown enum value")
	}
	io.Bool(&x.UnlockingIngredients.set)
	if x.UnlockingIngredients.set {
		if !io.Reading() && uint64(len(x.UnlockingIngredients.val)) > uint64(^uint32(0)) {
			io.InvalidValue(len(x.UnlockingIngredients.val), "collection length overflows uint32")
			return
		}
		count115 := uint32(len(x.UnlockingIngredients.val))
		io.Varuint32(&count115)
		if io.Reading() {
			if uint64(count115) > uint64(^uint(0)>>1) {
				io.InvalidValue(count115, "collection length overflows int")
				return
			}
			x.UnlockingIngredients.val = make([]CerealizerRecipeIngredientSerializedData, int(count115))
		}
		for index116 := range x.UnlockingIngredients.val {
			x.UnlockingIngredients.val[index116].Marshal(io)
		}
	} else if io.Reading() {
		var zero []CerealizerRecipeIngredientSerializedData
		x.UnlockingIngredients.val = zero
	}
}

// Marshal reads or writes ChangeEntityScore using its canonical wire layout.
func (x *ChangeEntityScore) Marshal(io IO) {
	io.String(&x.Action)
	x.ScoreboardId.Marshal(io)
	io.String(&x.ObjectiveName)
	io.Int32(&x.ScoreValue)
	x.ActorId.Marshal(io)
}

// Marshal reads or writes ChangeFakePlayerScore using its canonical wire layout.
func (x *ChangeFakePlayerScore) Marshal(io IO) {
	io.String(&x.Action)
	x.ScoreboardId.Marshal(io)
	io.String(&x.ObjectiveName)
	io.Int32(&x.ScoreValue)
	io.String(&x.FakePlayerName)
}

// Marshal reads or writes ChangePlayerScore using its canonical wire layout.
func (x *ChangePlayerScore) Marshal(io IO) {
	io.String(&x.Action)
	x.ScoreboardId.Marshal(io)
	io.String(&x.ObjectiveName)
	io.Int32(&x.ScoreValue)
	x.PlayerUniqueId.Marshal(io)
}

// Marshal reads or writes ChunkPos using its canonical wire layout.
func (x *ChunkPos) Marshal(io IO) {
	io.Varint32(&x.X)
	io.Varint32(&x.Z)
}

// Marshal reads or writes ClientboundDebugRendererDebugMarkerData using its canonical wire layout.
func (x *ClientboundDebugRendererDebugMarkerData) Marshal(io IO) {
	io.String(&x.Text)
	io.Vec3(&x.Position)
	io.RGBA(&x.Color)
	io.Uint64(&x.Duration)
}

// Marshal reads or writes CommandBlockUpdateBlockCommandData using its canonical wire layout.
func (x *CommandBlockUpdateBlockCommandData) Marshal(io IO) {
	x.BlockPosition.Marshal(io)
	io.Varuint32(&x.CommandBlockMode)
	io.Bool(&x.RedstoneMode)
	io.Bool(&x.IsConditional)
}

// Marshal reads or writes CommandBlockUpdateEntityCommandTarget using its canonical wire layout.
func (x *CommandBlockUpdateEntityCommandTarget) Marshal(io IO) {
	x.TargetRuntimeID.Marshal(io)
}

func marshalCommandBlockUpdateTarget(io IO, x *CommandBlockUpdateTarget) {
	if io.Reading() {
		var tag uint32
		io.Varuint32(&tag)
		switch int64(tag) {
		case 0:
			var value CommandBlockUpdateEntityCommandTarget
			value.Marshal(io)
			*x = value
		case 1:
			var value CommandBlockUpdateBlockCommandData
			value.Marshal(io)
			*x = value
		default:
			io.InvalidValue(tag, "unknown union tag")
		}
		return
	}
	switch value := (*x).(type) {
	case CommandBlockUpdateEntityCommandTarget:
		tag := uint32(0)
		io.Varuint32(&tag)
		value.Marshal(io)
	case CommandBlockUpdateBlockCommandData:
		tag := uint32(1)
		io.Varuint32(&tag)
		value.Marshal(io)
	default:
		io.InvalidValue(*x, "unknown union value")
	}
}

// Marshal reads or writes CommandOriginData using its canonical wire layout.
func (x *CommandOriginData) Marshal(io IO) {
	io.String(&x.Type)
	io.UUID(&x.UUID)
	io.String(&x.RequestId)
	io.Int64(&x.PlayerId)
}

// Marshal reads or writes CommandOutputData using its canonical wire layout.
func (x *CommandOutputData) Marshal(io IO) {
	io.String(&x.OutputType)
	io.Uint32(&x.SuccessCount)
	if !io.Reading() && uint64(len(x.OutputMessages)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.OutputMessages), "collection length overflows uint32")
		return
	}
	count117 := uint32(len(x.OutputMessages))
	io.Varuint32(&count117)
	if io.Reading() {
		if uint64(count117) > uint64(^uint(0)>>1) {
			io.InvalidValue(count117, "collection length overflows int")
			return
		}
		x.OutputMessages = make([]CommandOutputMessage, int(count117))
	}
	for index118 := range x.OutputMessages {
		x.OutputMessages[index118].Marshal(io)
	}
	io.Bool(&x.DataSet.set)
	if x.DataSet.set {
		io.String(&x.DataSet.val)
	} else if io.Reading() {
		var zero string
		x.DataSet.val = zero
	}
}

// Marshal reads or writes CommandOutputMessage using its canonical wire layout.
func (x *CommandOutputMessage) Marshal(io IO) {
	io.String(&x.MessageID)
	io.Bool(&x.Successful)
	if !io.Reading() && uint64(len(x.Parameters)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.Parameters), "collection length overflows uint32")
		return
	}
	count119 := uint32(len(x.Parameters))
	io.Varuint32(&count119)
	if io.Reading() {
		if uint64(count119) > uint64(^uint(0)>>1) {
			io.InvalidValue(count119, "collection length overflows int")
			return
		}
		x.Parameters = make([]string, int(count119))
	}
	for index120 := range x.Parameters {
		io.String(&x.Parameters[index120])
	}
}

// Marshal reads or writes ConeData using its canonical wire layout.
func (x *ConeData) Marshal(io IO) {
	io.Vec2(&x.Radii)
	io.Float32(&x.Height)
	io.Uint8(&x.NumSegments)
}

// Marshal reads or writes ContainerMixDataEntry using its canonical wire layout.
func (x *ContainerMixDataEntry) Marshal(io IO) {
	io.Varint32(&x.FromItemId)
	io.Varint32(&x.ReagentItemId)
	io.Varint32(&x.ToItemId)
}

// Marshal reads or writes ContentIdentity using its canonical wire layout.
func (x *ContentIdentity) Marshal(io IO) {
	io.String(&x.Identity)
}

// Marshal reads or writes CreativeGroupInfo using its canonical wire layout.
func (x *CreativeGroupInfo) Marshal(io IO) {
	enumValue121 := uint8(x.CreativeCategory)
	io.Uint8(&enumValue121)
	x.CreativeCategory = CreativeItemCategory(enumValue121)
	switch int64(enumValue121) {
	case 1, 2, 3, 4, 5:
	default:
		io.InvalidValue(enumValue121, "unknown enum value")
	}
	io.String(&x.Name)
	x.GroupIconItem.Marshal(io)
}

// Marshal reads or writes CreativeItemEntry using its canonical wire layout.
func (x *CreativeItemEntry) Marshal(io IO) {
	x.CreativeNetId.Marshal(io)
	x.ItemInstance.Marshal(io)
	io.Varuint32(&x.GroupIndex)
}

// Marshal reads or writes CylinderData using its canonical wire layout.
func (x *CylinderData) Marshal(io IO) {
	io.Vec2(&x.RadiusX)
	io.Vec2(&x.RadiusZ)
	io.Float32(&x.Height)
	io.Uint8(&x.NumSegments)
}

// Marshal reads or writes DataItemByte using its canonical wire layout.
func (x *DataItemByte) Marshal(io IO) {
	enumValue122 := uint8(x.Type)
	io.Uint8(&enumValue122)
	x.Type = DataItemType(enumValue122)
	switch int64(enumValue122) {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8:
	default:
		io.InvalidValue(enumValue122, "unknown enum value")
	}
	io.Int8(&x.Value)
}

// Marshal reads or writes DataItemCompoundTag using its canonical wire layout.
func (x *DataItemCompoundTag) Marshal(io IO) {
	enumValue123 := uint8(x.Type)
	io.Uint8(&enumValue123)
	x.Type = DataItemType(enumValue123)
	switch int64(enumValue123) {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8:
	default:
		io.InvalidValue(enumValue123, "unknown enum value")
	}
	io.NBT(&x.Value)
}

// Marshal reads or writes DataItemEntry using its canonical wire layout.
func (x *DataItemEntry) Marshal(io IO) {
	io.Varuint32(&x.ID)
	marshalDataItemEntryValue(io, &x.Payload)
}

func marshalDataItemEntryValue(io IO, x *DataItemEntryValue) {
	if io.Reading() {
		var tag uint8
		io.Uint8(&tag)
		switch int64(tag) {
		case 0:
			var value DataItemByte
			value.Marshal(io)
			*x = value
		case 1:
			var value DataItemShort
			value.Marshal(io)
			*x = value
		case 2:
			var value DataItemInt
			value.Marshal(io)
			*x = value
		case 3:
			var value DataItemFloat
			value.Marshal(io)
			*x = value
		case 4:
			var value DataItemString
			value.Marshal(io)
			*x = value
		case 5:
			var value DataItemCompoundTag
			value.Marshal(io)
			*x = value
		case 6:
			var value DataItemPos
			value.Marshal(io)
			*x = value
		case 7:
			var value DataItemInt64
			value.Marshal(io)
			*x = value
		case 8:
			var value DataItemVec3
			value.Marshal(io)
			*x = value
		default:
			io.InvalidValue(tag, "unknown union tag")
		}
		return
	}
	switch value := (*x).(type) {
	case DataItemByte:
		tag := uint8(0)
		io.Uint8(&tag)
		value.Marshal(io)
	case DataItemShort:
		tag := uint8(1)
		io.Uint8(&tag)
		value.Marshal(io)
	case DataItemInt:
		tag := uint8(2)
		io.Uint8(&tag)
		value.Marshal(io)
	case DataItemFloat:
		tag := uint8(3)
		io.Uint8(&tag)
		value.Marshal(io)
	case DataItemString:
		tag := uint8(4)
		io.Uint8(&tag)
		value.Marshal(io)
	case DataItemCompoundTag:
		tag := uint8(5)
		io.Uint8(&tag)
		value.Marshal(io)
	case DataItemPos:
		tag := uint8(6)
		io.Uint8(&tag)
		value.Marshal(io)
	case DataItemInt64:
		tag := uint8(7)
		io.Uint8(&tag)
		value.Marshal(io)
	case DataItemVec3:
		tag := uint8(8)
		io.Uint8(&tag)
		value.Marshal(io)
	default:
		io.InvalidValue(*x, "unknown union value")
	}
}

// Marshal reads or writes DataItemFloat using its canonical wire layout.
func (x *DataItemFloat) Marshal(io IO) {
	enumValue124 := uint8(x.Type)
	io.Uint8(&enumValue124)
	x.Type = DataItemType(enumValue124)
	switch int64(enumValue124) {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8:
	default:
		io.InvalidValue(enumValue124, "unknown enum value")
	}
	io.Float32(&x.Value)
}

// Marshal reads or writes DataItemInt using its canonical wire layout.
func (x *DataItemInt) Marshal(io IO) {
	enumValue125 := uint8(x.Type)
	io.Uint8(&enumValue125)
	x.Type = DataItemType(enumValue125)
	switch int64(enumValue125) {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8:
	default:
		io.InvalidValue(enumValue125, "unknown enum value")
	}
	io.Varint32(&x.Value)
}

// Marshal reads or writes DataItemInt64 using its canonical wire layout.
func (x *DataItemInt64) Marshal(io IO) {
	enumValue126 := uint8(x.Type)
	io.Uint8(&enumValue126)
	x.Type = DataItemType(enumValue126)
	switch int64(enumValue126) {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8:
	default:
		io.InvalidValue(enumValue126, "unknown enum value")
	}
	io.Varint64(&x.Value)
}

// Marshal reads or writes DataItemPos using its canonical wire layout.
func (x *DataItemPos) Marshal(io IO) {
	enumValue127 := uint8(x.Type)
	io.Uint8(&enumValue127)
	x.Type = DataItemType(enumValue127)
	switch int64(enumValue127) {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8:
	default:
		io.InvalidValue(enumValue127, "unknown enum value")
	}
	x.Value.Marshal(io)
}

// Marshal reads or writes DataItemShort using its canonical wire layout.
func (x *DataItemShort) Marshal(io IO) {
	enumValue128 := uint8(x.Type)
	io.Uint8(&enumValue128)
	x.Type = DataItemType(enumValue128)
	switch int64(enumValue128) {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8:
	default:
		io.InvalidValue(enumValue128, "unknown enum value")
	}
	io.Int16(&x.Value)
}

// Marshal reads or writes DataItemString using its canonical wire layout.
func (x *DataItemString) Marshal(io IO) {
	enumValue129 := uint8(x.Type)
	io.Uint8(&enumValue129)
	x.Type = DataItemType(enumValue129)
	switch int64(enumValue129) {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8:
	default:
		io.InvalidValue(enumValue129, "unknown enum value")
	}
	io.String(&x.Value)
}

// Marshal reads or writes DataItemVec3 using its canonical wire layout.
func (x *DataItemVec3) Marshal(io IO) {
	enumValue130 := uint8(x.Type)
	io.Uint8(&enumValue130)
	x.Type = DataItemType(enumValue130)
	switch int64(enumValue130) {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8:
	default:
		io.InvalidValue(enumValue130, "unknown enum value")
	}
	io.Vec3(&x.Value)
}

// Marshal reads or writes DimensionDefinition using its canonical wire layout.
func (x *DimensionDefinition) Marshal(io IO) {
	io.Varint32(&x.HeightMaximum)
	io.Varint32(&x.HeightMinimum)
	enumValue131 := int32(x.GeneratorType)
	io.Varint32(&enumValue131)
	x.GeneratorType = GeneratorType(enumValue131)
	switch int64(enumValue131) {
	case 0, 1, 2, 3, 4, 5, 6:
	default:
		io.InvalidValue(enumValue131, "unknown enum value")
	}
	x.DimensionType.Marshal(io)
	io.UUID(&x.PackId)
}

// Marshal reads or writes DimensionType using its canonical wire layout.
func (x *DimensionType) Marshal(io IO) {
	io.Varint32(&x.Value)
}

func marshalDisconnectMessages(io IO, x *DisconnectMessages) {
	if io.Reading() {
		var tag uint32
		io.Varuint32(&tag)
		switch int64(tag) {
		case 0:
			var value DisconnectPacketMessages
			value.Marshal(io)
			*x = value
		case 1:
			var value DisconnectMessagesEmpty1
			value.Marshal(io)
			*x = value
		default:
			io.InvalidValue(tag, "unknown union tag")
		}
		return
	}
	switch value := (*x).(type) {
	case DisconnectPacketMessages:
		tag := uint32(0)
		io.Varuint32(&tag)
		value.Marshal(io)
	case DisconnectMessagesEmpty1:
		tag := uint32(1)
		io.Varuint32(&tag)
		value.Marshal(io)
	default:
		io.InvalidValue(*x, "unknown union value")
	}
}

// Marshal reads or writes DisconnectMessagesEmpty1 using its canonical wire layout.
func (x *DisconnectMessagesEmpty1) Marshal(io IO) {
}

// Marshal reads or writes DisconnectPacketMessages using its canonical wire layout.
func (x *DisconnectPacketMessages) Marshal(io IO) {
	io.String(&x.Message)
	io.String(&x.FilteredMessage)
}

func marshalEAS(io IO, x *EAS) {
	if io.Reading() {
		var tag uint32
		io.Varuint32(&tag)
		switch int64(tag) {
		case 0:
			var value EASBoolAttributeData
			value.Marshal(io)
			*x = value
		case 1:
			var value EASFloatAttributeData
			value.Marshal(io)
			*x = value
		case 2:
			var value EASColorAttributeData
			value.Marshal(io)
			*x = value
		default:
			io.InvalidValue(tag, "unknown union tag")
		}
		return
	}
	switch value := (*x).(type) {
	case EASBoolAttributeData:
		tag := uint32(0)
		io.Varuint32(&tag)
		value.Marshal(io)
	case EASFloatAttributeData:
		tag := uint32(1)
		io.Varuint32(&tag)
		value.Marshal(io)
	case EASColorAttributeData:
		tag := uint32(2)
		io.Varuint32(&tag)
		value.Marshal(io)
	default:
		io.InvalidValue(*x, "unknown union value")
	}
}

// Marshal reads or writes EASAttributeLayerData using its canonical wire layout.
func (x *EASAttributeLayerData) Marshal(io IO) {
	io.String(&x.Name)
	io.Bool(&x.NoiseName.set)
	if x.NoiseName.set {
		io.String(&x.NoiseName.val)
	} else if io.Reading() {
		var zero string
		x.NoiseName.val = zero
	}
	x.Dimension.Marshal(io)
	x.Settings.Marshal(io)
	if !io.Reading() && uint64(len(x.Attributes)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.Attributes), "collection length overflows uint32")
		return
	}
	count132 := uint32(len(x.Attributes))
	io.Varuint32(&count132)
	if io.Reading() {
		if uint64(count132) > uint64(^uint(0)>>1) {
			io.InvalidValue(count132, "collection length overflows int")
			return
		}
		x.Attributes = make([]EASEnvironmentAttributeData, int(count132))
	}
	for index133 := range x.Attributes {
		x.Attributes[index133].Marshal(io)
	}
}

// Marshal reads or writes EASAttributeLayerSettings using its canonical wire layout.
func (x *EASAttributeLayerSettings) Marshal(io IO) {
	io.Int32(&x.Priority)
	io.Float32(&x.Weight)
	io.Bool(&x.Enabled)
	io.Bool(&x.TransitionsPaused)
}

// Marshal reads or writes EASBoolAttributeData using its canonical wire layout.
func (x *EASBoolAttributeData) Marshal(io IO) {
	io.Bool(&x.Value)
	io.String(&x.Operation)
}

// Marshal reads or writes EASColorAttributeData using its canonical wire layout.
func (x *EASColorAttributeData) Marshal(io IO) {
	for index134 := range x.Value {
		io.Int32(&x.Value[index134])
	}
	io.String(&x.Operation)
}

// Marshal reads or writes EASEnvironmentAttributeData using its canonical wire layout.
func (x *EASEnvironmentAttributeData) Marshal(io IO) {
	io.String(&x.AttributeName)
	io.Bool(&x.FromAttribute.set)
	if x.FromAttribute.set {
		marshalEAS(io, &x.FromAttribute.val)
	} else if io.Reading() {
		var zero EAS
		x.FromAttribute.val = zero
	}
	marshalEAS(io, &x.Attribute)
	io.Bool(&x.ToAttribute.set)
	if x.ToAttribute.set {
		marshalEAS(io, &x.ToAttribute.val)
	} else if io.Reading() {
		var zero EAS
		x.ToAttribute.val = zero
	}
	io.Uint32(&x.CurrentTransitionTicks)
	io.Uint32(&x.TotalTransitionTicks)
	io.String(&x.Easing)
	io.Uint32(&x.LocalTransitionTicks)
	io.Bool(&x.NoiseTransition)
}

// Marshal reads or writes EASFloatAttributeData using its canonical wire layout.
func (x *EASFloatAttributeData) Marshal(io IO) {
	io.Float32(&x.Value)
	io.String(&x.Operation)
	io.Bool(&x.ConstraintMin.set)
	if x.ConstraintMin.set {
		io.Float32(&x.ConstraintMin.val)
	} else if io.Reading() {
		var zero float32
		x.ConstraintMin.val = zero
	}
	io.Bool(&x.ConstraintMax.set)
	if x.ConstraintMax.set {
		io.Float32(&x.ConstraintMax.val)
	} else if io.Reading() {
		var zero float32
		x.ConstraintMax.val = zero
	}
}

// Marshal reads or writes ECSProfilingDiagnosticsEntityDiagnosticTimingInfo using its canonical wire layout.
func (x *ECSProfilingDiagnosticsEntityDiagnosticTimingInfo) Marshal(io IO) {
	io.String(&x.DisplayName)
	io.String(&x.Entity)
	io.Uint64(&x.TimeInNS)
	io.Uint8(&x.PercentOfTotal)
}

// Marshal reads or writes ECSProfilingDiagnosticsSystemCategory using its canonical wire layout.
func (x *ECSProfilingDiagnosticsSystemCategory) Marshal(io IO) {
	io.String(&x.CategoryName)
	io.Uint64(&x.SystemIndex)
}

// Marshal reads or writes ECSProfilingDiagnosticsSystemDiagnosticTimingInfo using its canonical wire layout.
func (x *ECSProfilingDiagnosticsSystemDiagnosticTimingInfo) Marshal(io IO) {
	io.String(&x.DisplayName)
	io.Uint64(&x.SystemIndex)
	io.Uint64(&x.TimeInNS)
	io.Uint8(&x.PercentOfTotal)
}

// Marshal reads or writes EduSharedUriResource using its canonical wire layout.
func (x *EduSharedUriResource) Marshal(io IO) {
	io.String(&x.ButtonName)
	io.String(&x.LinkUri)
}

// Marshal reads or writes EducationLevelSettings using its canonical wire layout.
func (x *EducationLevelSettings) Marshal(io IO) {
	io.String(&x.CodeBuilderDefaultURI)
	io.String(&x.CodeBuilderTitle)
	io.Bool(&x.CanResizeCodeBuilder)
	io.Bool(&x.DisableLegacyTitleBar)
	io.String(&x.PostProcessFilter)
	io.String(&x.ScreenshotBorderResourcePath)
	io.Bool(&x.AgentCapabilities.set)
	if x.AgentCapabilities.set {
		x.AgentCapabilities.val.Marshal(io)
	} else if io.Reading() {
		var zero AgentCapabilities
		x.AgentCapabilities.val = zero
	}
	x.LocalSettings.Marshal(io)
	io.Bool(&x.DeprecatedAlwaysFalse)
	io.Bool(&x.ExternalLinkSettings.set)
	if x.ExternalLinkSettings.set {
		x.ExternalLinkSettings.val.Marshal(io)
	} else if io.Reading() {
		var zero ExternalLinkSettings
		x.ExternalLinkSettings.val = zero
	}
}

// Marshal reads or writes EducationLocalLevelSettings using its canonical wire layout.
func (x *EducationLocalLevelSettings) Marshal(io IO) {
	io.Bool(&x.CodeBuilderOverrideUri.set)
	if x.CodeBuilderOverrideUri.set {
		io.String(&x.CodeBuilderOverrideUri.val)
	} else if io.Reading() {
		var zero string
		x.CodeBuilderOverrideUri.val = zero
	}
}

// Marshal reads or writes EllipsoidData using its canonical wire layout.
func (x *EllipsoidData) Marshal(io IO) {
	io.Vec3(&x.Radii)
	io.Uint8(&x.SegmentsPerAxis)
}

// Marshal reads or writes EnchantmentInstance using its canonical wire layout.
func (x *EnchantmentInstance) Marshal(io IO) {
	enumValue135 := uint8(x.EnchantType)
	io.Uint8(&enumValue135)
	x.EnchantType = EnchantType(enumValue135)
	switch int64(enumValue135) {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43:
	default:
		io.InvalidValue(enumValue135, "unknown enum value")
	}
	io.Uint8(&x.EnchantLevel)
}

// Marshal reads or writes EntityNetId using its canonical wire layout.
func (x *EntityNetId) Marshal(io IO) {
	io.Varuint32(&x.RawId)
}

// Marshal reads or writes Experiments using its canonical wire layout.
func (x *Experiments) Marshal(io IO) {
	if !io.Reading() && uint64(len(x.Toggles)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.Toggles), "collection length overflows uint32")
		return
	}
	count136 := uint32(len(x.Toggles))
	io.Uint32(&count136)
	if io.Reading() {
		if uint64(count136) > uint64(^uint(0)>>1) {
			io.InvalidValue(count136, "collection length overflows int")
			return
		}
		x.Toggles = make([]CerealizerExperimentsAnonExperimentToggle, int(count136))
	}
	for index137 := range x.Toggles {
		x.Toggles[index137].Marshal(io)
	}
	io.Bool(&x.ExperimentsEverToggled)
}

// Marshal reads or writes ExternalLinkSettings using its canonical wire layout.
func (x *ExternalLinkSettings) Marshal(io IO) {
	io.String(&x.URL)
	io.String(&x.DisplayName)
}

// Marshal reads or writes FeatureRegistryFeatureBinaryJsonFormat using its canonical wire layout.
func (x *FeatureRegistryFeatureBinaryJsonFormat) Marshal(io IO) {
	io.String(&x.FeatureName)
	io.String(&x.BinaryJsonOutput)
}

// Marshal reads or writes FloatRange using its canonical wire layout.
func (x *FloatRange) Marshal(io IO) {
	io.Float32(&x.Min)
	io.Float32(&x.Max)
}

// Marshal reads or writes FullContainerName using its canonical wire layout.
func (x *FullContainerName) Marshal(io IO) {
	enumValue138 := uint8(x.ContainerName)
	io.Uint8(&enumValue138)
	x.ContainerName = ContainerEnumName(enumValue138)
	switch int64(enumValue138) {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66:
	default:
		io.InvalidValue(enumValue138, "unknown enum value")
	}
	io.Bool(&x.DynamicID.set)
	if x.DynamicID.set {
		io.Uint32(&x.DynamicID.val)
	} else if io.Reading() {
		var zero uint32
		x.DynamicID.val = zero
	}
}

// Marshal reads or writes GameRule using its canonical wire layout.
func (x *GameRule) Marshal(io IO) {
	io.String(&x.RuleName)
	io.Bool(&x.RuleCanBeModified)
	marshalGameRuleRuleValue(io, &x.RuleValue)
}

func marshalGameRuleRuleValue(io IO, x *GameRuleRuleValue) {
	if io.Reading() {
		var tag uint32
		io.Varuint32(&tag)
		switch int64(tag) {
		case 0:
			var value GameRuleRuleValueEmpty0
			value.Marshal(io)
			*x = value
		case 1:
			var value GameRuleRuleValueBool
			value.Marshal(io)
			*x = value
		case 2:
			var value GameRuleRuleValueInt32
			value.Marshal(io)
			*x = value
		case 3:
			var value GameRuleRuleValueFloat
			value.Marshal(io)
			*x = value
		default:
			io.InvalidValue(tag, "unknown union tag")
		}
		return
	}
	switch value := (*x).(type) {
	case GameRuleRuleValueEmpty0:
		tag := uint32(0)
		io.Varuint32(&tag)
		value.Marshal(io)
	case GameRuleRuleValueBool:
		tag := uint32(1)
		io.Varuint32(&tag)
		value.Marshal(io)
	case GameRuleRuleValueInt32:
		tag := uint32(2)
		io.Varuint32(&tag)
		value.Marshal(io)
	case GameRuleRuleValueFloat:
		tag := uint32(3)
		io.Varuint32(&tag)
		value.Marshal(io)
	default:
		io.InvalidValue(*x, "unknown union value")
	}
}

// Marshal reads or writes GameRuleRuleValueBool using its canonical wire layout.
func (x *GameRuleRuleValueBool) Marshal(io IO) {
	io.Bool(&x.Value)
}

// Marshal reads or writes GameRuleRuleValueEmpty0 using its canonical wire layout.
func (x *GameRuleRuleValueEmpty0) Marshal(io IO) {
}

// Marshal reads or writes GameRuleRuleValueFloat using its canonical wire layout.
func (x *GameRuleRuleValueFloat) Marshal(io IO) {
	io.Float32(&x.Value)
}

// Marshal reads or writes GameRuleRuleValueInt32 using its canonical wire layout.
func (x *GameRuleRuleValueInt32) Marshal(io IO) {
	io.Int32(&x.Value)
}

// Marshal reads or writes GameRulesChangedPacketData using its canonical wire layout.
func (x *GameRulesChangedPacketData) Marshal(io IO) {
	if !io.Reading() && uint64(len(x.RulesList)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.RulesList), "collection length overflows uint32")
		return
	}
	count139 := uint32(len(x.RulesList))
	io.Varuint32(&count139)
	if io.Reading() {
		if uint64(count139) > uint64(^uint(0)>>1) {
			io.InvalidValue(count139, "collection length overflows int")
			return
		}
		x.RulesList = make([]GameRule, int(count139))
	}
	for index140 := range x.RulesList {
		x.RulesList[index140].Marshal(io)
	}
}

// Marshal reads or writes InventoryAction using its canonical wire layout.
func (x *InventoryAction) Marshal(io IO) {
	x.Source.Marshal(io)
	io.Varuint32(&x.Slot)
	x.FromItem.Marshal(io)
	x.ToItem.Marshal(io)
}

// Marshal reads or writes InventoryMismatchData using its canonical wire layout.
func (x *InventoryMismatchData) Marshal(io IO) {
	x.Actions.Marshal(io)
}

// Marshal reads or writes InventoryOptions using its canonical wire layout.
func (x *InventoryOptions) Marshal(io IO) {
	enumValue141 := int32(x.LeftInventoryTab)
	io.Varint32(&enumValue141)
	x.LeftInventoryTab = InventoryLeftTabIndex(enumValue141)
	switch int64(enumValue141) {
	case 0, 1, 2, 3, 4, 5, 6:
	default:
		io.InvalidValue(enumValue141, "unknown enum value")
	}
	enumValue142 := int32(x.RightInventoryTab)
	io.Varint32(&enumValue142)
	x.RightInventoryTab = InventoryRightTabIndex(enumValue142)
	switch int64(enumValue142) {
	case 0, 1, 2, 3:
	default:
		io.InvalidValue(enumValue142, "unknown enum value")
	}
	io.Bool(&x.Filtering)
	enumValue143 := int32(x.LayoutInv)
	io.Varint32(&enumValue143)
	x.LayoutInv = InventoryLayout(enumValue143)
	switch int64(enumValue143) {
	case 0, 1, 2, 3:
	default:
		io.InvalidValue(enumValue143, "unknown enum value")
	}
	enumValue144 := int32(x.LayoutCraft)
	io.Varint32(&enumValue144)
	x.LayoutCraft = InventoryLayout(enumValue144)
	switch int64(enumValue144) {
	case 0, 1, 2, 3:
	default:
		io.InvalidValue(enumValue144, "unknown enum value")
	}
}

// Marshal reads or writes InventorySource using its canonical wire layout.
func (x *InventorySource) Marshal(io IO) {
	enumValue145 := uint32(x.SourceType)
	io.Varuint32(&enumValue145)
	x.SourceType = InventorySourceType(enumValue145)
	switch int64(enumValue145) {
	case 0, 1, 2, 3, 99999:
	default:
		io.InvalidValue(enumValue145, "unknown enum value")
	}
	outer146 := true
	io.Bool(&outer146)
	if outer146 {
		io.Bool(&x.ContainerID.set)
		if x.ContainerID.set {
			io.Int8(&x.ContainerID.val)
		} else if io.Reading() {
			var zero int8
			x.ContainerID.val = zero
		}
	} else {
		x.ContainerID = Optional[int8]{}
	}
	outer147 := true
	io.Bool(&outer147)
	if outer147 {
		io.Bool(&x.BitFlags.set)
		if x.BitFlags.set {
			enumValue148 := uint32(x.BitFlags.val)
			io.Varuint32(&enumValue148)
			x.BitFlags.val = InventorySourceInventorySourceFlags(enumValue148)
			switch int64(enumValue148) {
			case 0, 1:
			default:
				io.InvalidValue(enumValue148, "unknown enum value")
			}
		} else if io.Reading() {
			var zero InventorySourceInventorySourceFlags
			x.BitFlags.val = zero
		}
	} else {
		x.BitFlags = Optional[InventorySourceInventorySourceFlags]{}
	}
}

// Marshal reads or writes InventoryTransactionData using its canonical wire layout.
func (x *InventoryTransactionData) Marshal(io IO) {
	io.Bool(&x.Actions.set)
	if x.Actions.set {
		if !io.Reading() && uint64(len(x.Actions.val)) > uint64(^uint32(0)) {
			io.InvalidValue(len(x.Actions.val), "collection length overflows uint32")
			return
		}
		count149 := uint32(len(x.Actions.val))
		io.Varuint32(&count149)
		if io.Reading() {
			if uint64(count149) > uint64(^uint(0)>>1) {
				io.InvalidValue(count149, "collection length overflows int")
				return
			}
			x.Actions.val = make([]InventoryAction, int(count149))
		}
		for index150 := range x.Actions.val {
			x.Actions.val[index150].Marshal(io)
		}
	} else if io.Reading() {
		var zero []InventoryAction
		x.Actions.val = zero
	}
}

func marshalInventoryTransactionTransactionValue(io IO, x *InventoryTransactionTransactionValue) {
	if io.Reading() {
		var tag uint32
		io.Varuint32(&tag)
		switch int64(tag) {
		case 0:
			var value NormalTransactionData
			value.Marshal(io)
			*x = value
		case 1:
			var value InventoryMismatchData
			value.Marshal(io)
			*x = value
		case 2:
			var value ItemUseInventoryTransaction
			value.Marshal(io)
			*x = value
		case 3:
			var value ItemUseOnActorInventoryTransaction
			value.Marshal(io)
			*x = value
		case 4:
			var value ItemReleaseInventoryTransaction
			value.Marshal(io)
			*x = value
		default:
			io.InvalidValue(tag, "unknown union tag")
		}
		return
	}
	switch value := (*x).(type) {
	case NormalTransactionData:
		tag := uint32(0)
		io.Varuint32(&tag)
		value.Marshal(io)
	case InventoryMismatchData:
		tag := uint32(1)
		io.Varuint32(&tag)
		value.Marshal(io)
	case ItemUseInventoryTransaction:
		tag := uint32(2)
		io.Varuint32(&tag)
		value.Marshal(io)
	case ItemUseOnActorInventoryTransaction:
		tag := uint32(3)
		io.Varuint32(&tag)
		value.Marshal(io)
	case ItemReleaseInventoryTransaction:
		tag := uint32(4)
		io.Varuint32(&tag)
		value.Marshal(io)
	default:
		io.InvalidValue(*x, "unknown union value")
	}
}

// Marshal reads or writes ItemData using its canonical wire layout.
func (x *ItemData) Marshal(io IO) {
	io.String(&x.ItemName)
	io.Int16(&x.ItemId)
	io.Bool(&x.IsComponentBased)
	enumValue151 := int32(x.ItemVersion)
	io.Varint32(&enumValue151)
	x.ItemVersion = ItemVersion(enumValue151)
	switch int64(enumValue151) {
	case 0, 1, 2:
	default:
		io.InvalidValue(enumValue151, "unknown enum value")
	}
	io.NBT(&x.ItemComponentData)
}

// Marshal reads or writes ItemEnchantOption using its canonical wire layout.
func (x *ItemEnchantOption) Marshal(io IO) {
	io.Uint8(&x.Cost)
	x.Enchants.Marshal(io)
	io.String(&x.EnchantName)
	x.EnchantNetId.Marshal(io)
}

// Marshal reads or writes ItemEnchants using its canonical wire layout.
func (x *ItemEnchants) Marshal(io IO) {
	io.Int32(&x.Slot)
	for index152 := range x.ItemEnchants {
		if !io.Reading() && uint64(len(x.ItemEnchants[index152])) > uint64(^uint32(0)) {
			io.InvalidValue(len(x.ItemEnchants[index152]), "collection length overflows uint32")
			return
		}
		count153 := uint32(len(x.ItemEnchants[index152]))
		io.Varuint32(&count153)
		if io.Reading() {
			if uint64(count153) > uint64(^uint(0)>>1) {
				io.InvalidValue(count153, "collection length overflows int")
				return
			}
			x.ItemEnchants[index152] = make([]EnchantmentInstance, int(count153))
		}
		for index154 := range x.ItemEnchants[index152] {
			x.ItemEnchants[index152][index154].Marshal(io)
		}
	}
}

// Marshal reads or writes ItemReleaseInventoryTransaction using its canonical wire layout.
func (x *ItemReleaseInventoryTransaction) Marshal(io IO) {
	x.Actions.Marshal(io)
	enumValue155 := int32(x.ActionType)
	io.Varint32(&enumValue155)
	x.ActionType = ItemReleaseInventoryTransactionActionType(enumValue155)
	switch int64(enumValue155) {
	case 0, 1:
	default:
		io.InvalidValue(enumValue155, "unknown enum value")
	}
	io.Varint32(&x.Slot)
	x.Item.Marshal(io)
	io.Vec3(&x.FromPosition)
}

func marshalItemStackRequestCereal(io IO, x *ItemStackRequestCereal) {
	if io.Reading() {
		var tag uint32
		io.Varuint32(&tag)
		switch int64(tag) {
		case 0:
			var value ItemStackRequestCerealTakeActionData
			value.Marshal(io)
			*x = value
		case 1:
			var value ItemStackRequestCerealPlaceActionData
			value.Marshal(io)
			*x = value
		case 2:
			var value ItemStackRequestCerealSwapActionData
			value.Marshal(io)
			*x = value
		case 3:
			var value ItemStackRequestCerealDropActionData
			value.Marshal(io)
			*x = value
		case 4:
			var value ItemStackRequestCerealDestroyActionData
			value.Marshal(io)
			*x = value
		case 5:
			var value ItemStackRequestCerealConsumeActionData
			value.Marshal(io)
			*x = value
		case 6:
			var value ItemStackRequestCerealCreateActionData
			value.Marshal(io)
			*x = value
		case 7:
			var value ItemStackRequestCerealLabTableCombineActionData
			value.Marshal(io)
			*x = value
		case 8:
			var value ItemStackRequestCerealBeaconPaymentActionData
			value.Marshal(io)
			*x = value
		case 9:
			var value ItemStackRequestCerealMineBlockActionData
			value.Marshal(io)
			*x = value
		case 10:
			var value ItemStackRequestCerealCraftRecipeActionData
			value.Marshal(io)
			*x = value
		case 11:
			var value ItemStackRequestCerealCraftRecipeAutoActionData
			value.Marshal(io)
			*x = value
		case 12:
			var value ItemStackRequestCerealCraftCreativeActionData
			value.Marshal(io)
			*x = value
		case 13:
			var value ItemStackRequestCerealCraftRecipeOptionalActionData
			value.Marshal(io)
			*x = value
		case 14:
			var value ItemStackRequestCerealCraftRepairAndDisenchantActionData
			value.Marshal(io)
			*x = value
		case 15:
			var value ItemStackRequestCerealCraftLoomActionData
			value.Marshal(io)
			*x = value
		case 16:
			var value ItemStackRequestCerealCraftNonImplementedActionData
			value.Marshal(io)
			*x = value
		case 17:
			var value ItemStackRequestCerealCraftResultsActionData
			value.Marshal(io)
			*x = value
		default:
			io.InvalidValue(tag, "unknown union tag")
		}
		return
	}
	switch value := (*x).(type) {
	case ItemStackRequestCerealTakeActionData:
		tag := uint32(0)
		io.Varuint32(&tag)
		value.Marshal(io)
	case ItemStackRequestCerealPlaceActionData:
		tag := uint32(1)
		io.Varuint32(&tag)
		value.Marshal(io)
	case ItemStackRequestCerealSwapActionData:
		tag := uint32(2)
		io.Varuint32(&tag)
		value.Marshal(io)
	case ItemStackRequestCerealDropActionData:
		tag := uint32(3)
		io.Varuint32(&tag)
		value.Marshal(io)
	case ItemStackRequestCerealDestroyActionData:
		tag := uint32(4)
		io.Varuint32(&tag)
		value.Marshal(io)
	case ItemStackRequestCerealConsumeActionData:
		tag := uint32(5)
		io.Varuint32(&tag)
		value.Marshal(io)
	case ItemStackRequestCerealCreateActionData:
		tag := uint32(6)
		io.Varuint32(&tag)
		value.Marshal(io)
	case ItemStackRequestCerealLabTableCombineActionData:
		tag := uint32(7)
		io.Varuint32(&tag)
		value.Marshal(io)
	case ItemStackRequestCerealBeaconPaymentActionData:
		tag := uint32(8)
		io.Varuint32(&tag)
		value.Marshal(io)
	case ItemStackRequestCerealMineBlockActionData:
		tag := uint32(9)
		io.Varuint32(&tag)
		value.Marshal(io)
	case ItemStackRequestCerealCraftRecipeActionData:
		tag := uint32(10)
		io.Varuint32(&tag)
		value.Marshal(io)
	case ItemStackRequestCerealCraftRecipeAutoActionData:
		tag := uint32(11)
		io.Varuint32(&tag)
		value.Marshal(io)
	case ItemStackRequestCerealCraftCreativeActionData:
		tag := uint32(12)
		io.Varuint32(&tag)
		value.Marshal(io)
	case ItemStackRequestCerealCraftRecipeOptionalActionData:
		tag := uint32(13)
		io.Varuint32(&tag)
		value.Marshal(io)
	case ItemStackRequestCerealCraftRepairAndDisenchantActionData:
		tag := uint32(14)
		io.Varuint32(&tag)
		value.Marshal(io)
	case ItemStackRequestCerealCraftLoomActionData:
		tag := uint32(15)
		io.Varuint32(&tag)
		value.Marshal(io)
	case ItemStackRequestCerealCraftNonImplementedActionData:
		tag := uint32(16)
		io.Varuint32(&tag)
		value.Marshal(io)
	case ItemStackRequestCerealCraftResultsActionData:
		tag := uint32(17)
		io.Varuint32(&tag)
		value.Marshal(io)
	default:
		io.InvalidValue(*x, "unknown union value")
	}
}

// Marshal reads or writes ItemStackRequestCerealBeaconPaymentActionData using its canonical wire layout.
func (x *ItemStackRequestCerealBeaconPaymentActionData) Marshal(io IO) {
	enumValue156 := uint8(x.ActionType)
	io.Uint8(&enumValue156)
	x.ActionType = ItemStackRequestActionType(enumValue156)
	switch int64(enumValue156) {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19:
	default:
		io.InvalidValue(enumValue156, "unknown enum value")
	}
	io.Varint32(&x.PrimaryEffectId)
	io.Varint32(&x.SecondaryEffectId)
}

// Marshal reads or writes ItemStackRequestCerealConsumeActionData using its canonical wire layout.
func (x *ItemStackRequestCerealConsumeActionData) Marshal(io IO) {
	enumValue157 := uint8(x.ActionType)
	io.Uint8(&enumValue157)
	x.ActionType = ItemStackRequestActionType(enumValue157)
	switch int64(enumValue157) {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19:
	default:
		io.InvalidValue(enumValue157, "unknown enum value")
	}
	io.Uint8(&x.Amount)
	x.Source.Marshal(io)
}

// Marshal reads or writes ItemStackRequestCerealCraftCreativeActionData using its canonical wire layout.
func (x *ItemStackRequestCerealCraftCreativeActionData) Marshal(io IO) {
	enumValue158 := uint8(x.ActionType)
	io.Uint8(&enumValue158)
	x.ActionType = ItemStackRequestActionType(enumValue158)
	switch int64(enumValue158) {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19:
	default:
		io.InvalidValue(enumValue158, "unknown enum value")
	}
	io.Varuint32(&x.CreativeItemNetId)
	io.Uint8(&x.NumberOfRequestedCrafts)
}

// Marshal reads or writes ItemStackRequestCerealCraftLoomActionData using its canonical wire layout.
func (x *ItemStackRequestCerealCraftLoomActionData) Marshal(io IO) {
	enumValue159 := uint8(x.ActionType)
	io.Uint8(&enumValue159)
	x.ActionType = ItemStackRequestActionType(enumValue159)
	switch int64(enumValue159) {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19:
	default:
		io.InvalidValue(enumValue159, "unknown enum value")
	}
	io.String(&x.PatternNameId)
	io.Uint8(&x.NumCrafts)
}

// Marshal reads or writes ItemStackRequestCerealCraftNonImplementedActionData using its canonical wire layout.
func (x *ItemStackRequestCerealCraftNonImplementedActionData) Marshal(io IO) {
	enumValue160 := uint8(x.ActionType)
	io.Uint8(&enumValue160)
	x.ActionType = ItemStackRequestActionType(enumValue160)
	switch int64(enumValue160) {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19:
	default:
		io.InvalidValue(enumValue160, "unknown enum value")
	}
}

// Marshal reads or writes ItemStackRequestCerealCraftRecipeActionData using its canonical wire layout.
func (x *ItemStackRequestCerealCraftRecipeActionData) Marshal(io IO) {
	enumValue161 := uint8(x.ActionType)
	io.Uint8(&enumValue161)
	x.ActionType = ItemStackRequestActionType(enumValue161)
	switch int64(enumValue161) {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19:
	default:
		io.InvalidValue(enumValue161, "unknown enum value")
	}
	x.RecipeNetId.Marshal(io)
	io.Uint8(&x.NumberOfRequestedCrafts)
}

// Marshal reads or writes ItemStackRequestCerealCraftRecipeAutoActionData using its canonical wire layout.
func (x *ItemStackRequestCerealCraftRecipeAutoActionData) Marshal(io IO) {
	enumValue162 := uint8(x.ActionType)
	io.Uint8(&enumValue162)
	x.ActionType = ItemStackRequestActionType(enumValue162)
	switch int64(enumValue162) {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19:
	default:
		io.InvalidValue(enumValue162, "unknown enum value")
	}
	x.RecipeNetId.Marshal(io)
	io.Uint8(&x.NumberOfRequestedCrafts)
	if !io.Reading() && uint64(len(x.Ingredients)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.Ingredients), "collection length overflows uint32")
		return
	}
	count163 := uint32(len(x.Ingredients))
	io.Varuint32(&count163)
	if io.Reading() {
		if uint64(count163) > uint64(^uint(0)>>1) {
			io.InvalidValue(count163, "collection length overflows int")
			return
		}
		x.Ingredients = make([]ItemStackRequestCerealRecipeIngredientData, int(count163))
	}
	for index164 := range x.Ingredients {
		x.Ingredients[index164].Marshal(io)
	}
}

// Marshal reads or writes ItemStackRequestCerealCraftRecipeOptionalActionData using its canonical wire layout.
func (x *ItemStackRequestCerealCraftRecipeOptionalActionData) Marshal(io IO) {
	enumValue165 := uint8(x.ActionType)
	io.Uint8(&enumValue165)
	x.ActionType = ItemStackRequestActionType(enumValue165)
	switch int64(enumValue165) {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19:
	default:
		io.InvalidValue(enumValue165, "unknown enum value")
	}
	x.RecipeNetId.Marshal(io)
	io.Int32(&x.FilteredStringIndex)
}

// Marshal reads or writes ItemStackRequestCerealCraftRepairAndDisenchantActionData using its canonical wire layout.
func (x *ItemStackRequestCerealCraftRepairAndDisenchantActionData) Marshal(io IO) {
	enumValue166 := uint8(x.ActionType)
	io.Uint8(&enumValue166)
	x.ActionType = ItemStackRequestActionType(enumValue166)
	switch int64(enumValue166) {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19:
	default:
		io.InvalidValue(enumValue166, "unknown enum value")
	}
	io.Int32(&x.RecipeNetId)
	io.Uint8(&x.NumberOfRequestedCrafts)
	io.Varint32(&x.RepairCost)
}

// Marshal reads or writes ItemStackRequestCerealCraftResultsActionData using its canonical wire layout.
func (x *ItemStackRequestCerealCraftResultsActionData) Marshal(io IO) {
	enumValue167 := uint8(x.ActionType)
	io.Uint8(&enumValue167)
	x.ActionType = ItemStackRequestActionType(enumValue167)
	switch int64(enumValue167) {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19:
	default:
		io.InvalidValue(enumValue167, "unknown enum value")
	}
	if !io.Reading() && uint64(len(x.CraftResults)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.CraftResults), "collection length overflows uint32")
		return
	}
	count168 := uint32(len(x.CraftResults))
	io.Varuint32(&count168)
	if io.Reading() {
		if uint64(count168) > uint64(^uint(0)>>1) {
			io.InvalidValue(count168, "collection length overflows int")
			return
		}
		x.CraftResults = make([]ItemStackRequestCerealNetworkItemInstanceDescriptorData, int(count168))
	}
	for index169 := range x.CraftResults {
		x.CraftResults[index169].Marshal(io)
	}
	io.Uint8(&x.NumCrafts)
}

// Marshal reads or writes ItemStackRequestCerealCreateActionData using its canonical wire layout.
func (x *ItemStackRequestCerealCreateActionData) Marshal(io IO) {
	enumValue170 := uint8(x.ActionType)
	io.Uint8(&enumValue170)
	x.ActionType = ItemStackRequestActionType(enumValue170)
	switch int64(enumValue170) {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19:
	default:
		io.InvalidValue(enumValue170, "unknown enum value")
	}
	io.Uint8(&x.ResultsIndex)
}

// Marshal reads or writes ItemStackRequestCerealDestroyActionData using its canonical wire layout.
func (x *ItemStackRequestCerealDestroyActionData) Marshal(io IO) {
	enumValue171 := uint8(x.ActionType)
	io.Uint8(&enumValue171)
	x.ActionType = ItemStackRequestActionType(enumValue171)
	switch int64(enumValue171) {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19:
	default:
		io.InvalidValue(enumValue171, "unknown enum value")
	}
	io.Uint8(&x.Amount)
	x.Source.Marshal(io)
}

// Marshal reads or writes ItemStackRequestCerealDropActionData using its canonical wire layout.
func (x *ItemStackRequestCerealDropActionData) Marshal(io IO) {
	enumValue172 := uint8(x.ActionType)
	io.Uint8(&enumValue172)
	x.ActionType = ItemStackRequestActionType(enumValue172)
	switch int64(enumValue172) {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19:
	default:
		io.InvalidValue(enumValue172, "unknown enum value")
	}
	io.Uint8(&x.Amount)
	x.Source.Marshal(io)
	io.Bool(&x.Randomly)
}

// Marshal reads or writes ItemStackRequestCerealEmptyItemDescriptorData using its canonical wire layout.
func (x *ItemStackRequestCerealEmptyItemDescriptorData) Marshal(io IO) {
	enumValue173 := uint8(x.DescriptorType)
	io.Uint8(&enumValue173)
	x.DescriptorType = ItemStackRequestCerealItemDescriptorType(enumValue173)
	switch int64(enumValue173) {
	case 0, 1, 2, 3:
	default:
		io.InvalidValue(enumValue173, "unknown enum value")
	}
}

// Marshal reads or writes ItemStackRequestCerealItemNameDescriptorData using its canonical wire layout.
func (x *ItemStackRequestCerealItemNameDescriptorData) Marshal(io IO) {
	enumValue174 := uint8(x.DescriptorType)
	io.Uint8(&enumValue174)
	x.DescriptorType = ItemStackRequestCerealItemDescriptorType(enumValue174)
	switch int64(enumValue174) {
	case 0, 1, 2, 3:
	default:
		io.InvalidValue(enumValue174, "unknown enum value")
	}
	io.String(&x.FullName)
	io.Varint32(&x.AuxValue)
}

// Marshal reads or writes ItemStackRequestCerealItemTagDescriptorData using its canonical wire layout.
func (x *ItemStackRequestCerealItemTagDescriptorData) Marshal(io IO) {
	enumValue175 := uint8(x.DescriptorType)
	io.Uint8(&enumValue175)
	x.DescriptorType = ItemStackRequestCerealItemDescriptorType(enumValue175)
	switch int64(enumValue175) {
	case 0, 1, 2, 3:
	default:
		io.InvalidValue(enumValue175, "unknown enum value")
	}
	io.String(&x.ItemTag)
}

// Marshal reads or writes ItemStackRequestCerealLabTableCombineActionData using its canonical wire layout.
func (x *ItemStackRequestCerealLabTableCombineActionData) Marshal(io IO) {
	enumValue176 := uint8(x.ActionType)
	io.Uint8(&enumValue176)
	x.ActionType = ItemStackRequestActionType(enumValue176)
	switch int64(enumValue176) {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19:
	default:
		io.InvalidValue(enumValue176, "unknown enum value")
	}
}

// Marshal reads or writes ItemStackRequestCerealMineBlockActionData using its canonical wire layout.
func (x *ItemStackRequestCerealMineBlockActionData) Marshal(io IO) {
	enumValue177 := uint8(x.ActionType)
	io.Uint8(&enumValue177)
	x.ActionType = ItemStackRequestActionType(enumValue177)
	switch int64(enumValue177) {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19:
	default:
		io.InvalidValue(enumValue177, "unknown enum value")
	}
	io.Varint32(&x.Slot)
	io.Varint32(&x.PredictedDurability)
	io.Int32(&x.NetIdVariant)
}

// Marshal reads or writes ItemStackRequestCerealMoLangItemDescriptorData using its canonical wire layout.
func (x *ItemStackRequestCerealMoLangItemDescriptorData) Marshal(io IO) {
	enumValue178 := uint8(x.DescriptorType)
	io.Uint8(&enumValue178)
	x.DescriptorType = ItemStackRequestCerealItemDescriptorType(enumValue178)
	switch int64(enumValue178) {
	case 0, 1, 2, 3:
	default:
		io.InvalidValue(enumValue178, "unknown enum value")
	}
	io.String(&x.TagExpression)
	enumValue179 := int16(x.MolangVersion)
	io.Int16(&enumValue179)
	x.MolangVersion = MoLangVersion(enumValue179)
	switch int64(enumValue179) {
	case -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14:
	default:
		io.InvalidValue(enumValue179, "unknown enum value")
	}
}

// Marshal reads or writes ItemStackRequestCerealNetworkItemInstanceDescriptorData using its canonical wire layout.
func (x *ItemStackRequestCerealNetworkItemInstanceDescriptorData) Marshal(io IO) {
	marshalItemStackRequestCerealRecipeIngredientDataItemDescriptor(io, &x.ItemDescriptor)
	io.Uint16(&x.StackSize)
	io.Varuint32(&x.BlockRuntimeId)
	io.Bytes(&x.UserDataBuffer)
}

// Marshal reads or writes ItemStackRequestCerealPlaceActionData using its canonical wire layout.
func (x *ItemStackRequestCerealPlaceActionData) Marshal(io IO) {
	enumValue180 := uint8(x.ActionType)
	io.Uint8(&enumValue180)
	x.ActionType = ItemStackRequestActionType(enumValue180)
	switch int64(enumValue180) {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19:
	default:
		io.InvalidValue(enumValue180, "unknown enum value")
	}
	io.Uint8(&x.Amount)
	x.Source.Marshal(io)
	x.Destination.Marshal(io)
}

// Marshal reads or writes ItemStackRequestCerealRecipeIngredientData using its canonical wire layout.
func (x *ItemStackRequestCerealRecipeIngredientData) Marshal(io IO) {
	marshalItemStackRequestCerealRecipeIngredientDataItemDescriptor(io, &x.ItemDescriptor)
	io.Uint16(&x.StackSize)
}

func marshalItemStackRequestCerealRecipeIngredientDataItemDescriptor(io IO, x *ItemStackRequestCerealRecipeIngredientDataItemDescriptor) {
	if io.Reading() {
		var tag uint32
		io.Varuint32(&tag)
		switch int64(tag) {
		case 0:
			var value ItemStackRequestCerealEmptyItemDescriptorData
			value.Marshal(io)
			*x = value
		case 1:
			var value ItemStackRequestCerealItemNameDescriptorData
			value.Marshal(io)
			*x = value
		case 2:
			var value ItemStackRequestCerealMoLangItemDescriptorData
			value.Marshal(io)
			*x = value
		case 3:
			var value ItemStackRequestCerealItemTagDescriptorData
			value.Marshal(io)
			*x = value
		default:
			io.InvalidValue(tag, "unknown union tag")
		}
		return
	}
	switch value := (*x).(type) {
	case ItemStackRequestCerealEmptyItemDescriptorData:
		tag := uint32(0)
		io.Varuint32(&tag)
		value.Marshal(io)
	case ItemStackRequestCerealItemNameDescriptorData:
		tag := uint32(1)
		io.Varuint32(&tag)
		value.Marshal(io)
	case ItemStackRequestCerealMoLangItemDescriptorData:
		tag := uint32(2)
		io.Varuint32(&tag)
		value.Marshal(io)
	case ItemStackRequestCerealItemTagDescriptorData:
		tag := uint32(3)
		io.Varuint32(&tag)
		value.Marshal(io)
	default:
		io.InvalidValue(*x, "unknown union value")
	}
}

// Marshal reads or writes ItemStackRequestCerealRequestData using its canonical wire layout.
func (x *ItemStackRequestCerealRequestData) Marshal(io IO) {
	x.ClientRequestId.Marshal(io)
	if !io.Reading() && uint64(len(x.Actions)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.Actions), "collection length overflows uint32")
		return
	}
	count181 := uint32(len(x.Actions))
	io.Varuint32(&count181)
	if io.Reading() {
		if uint64(count181) > uint64(^uint(0)>>1) {
			io.InvalidValue(count181, "collection length overflows int")
			return
		}
		x.Actions = make([]ItemStackRequestCereal, int(count181))
	}
	for index182 := range x.Actions {
		marshalItemStackRequestCereal(io, &x.Actions[index182])
	}
	if !io.Reading() && uint64(len(x.StringsToFilter)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.StringsToFilter), "collection length overflows uint32")
		return
	}
	count183 := uint32(len(x.StringsToFilter))
	io.Varuint32(&count183)
	if io.Reading() {
		if uint64(count183) > uint64(^uint(0)>>1) {
			io.InvalidValue(count183, "collection length overflows int")
			return
		}
		x.StringsToFilter = make([]string, int(count183))
	}
	for index184 := range x.StringsToFilter {
		io.String(&x.StringsToFilter[index184])
	}
	enumValue185 := int32(x.StringsToFilterOrigin)
	io.Int32(&enumValue185)
	x.StringsToFilterOrigin = TextProcessingEventOrigin(enumValue185)
	switch int64(enumValue185) {
	case -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15:
	default:
		io.InvalidValue(enumValue185, "unknown enum value")
	}
}

// Marshal reads or writes ItemStackRequestCerealSlotInfoData using its canonical wire layout.
func (x *ItemStackRequestCerealSlotInfoData) Marshal(io IO) {
	x.FullContainerName.Marshal(io)
	io.Uint8(&x.Slot)
	io.Int32(&x.NetIdVariant)
}

// Marshal reads or writes ItemStackRequestCerealSwapActionData using its canonical wire layout.
func (x *ItemStackRequestCerealSwapActionData) Marshal(io IO) {
	enumValue186 := uint8(x.ActionType)
	io.Uint8(&enumValue186)
	x.ActionType = ItemStackRequestActionType(enumValue186)
	switch int64(enumValue186) {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19:
	default:
		io.InvalidValue(enumValue186, "unknown enum value")
	}
	x.Source.Marshal(io)
	x.Destination.Marshal(io)
}

// Marshal reads or writes ItemStackRequestCerealTakeActionData using its canonical wire layout.
func (x *ItemStackRequestCerealTakeActionData) Marshal(io IO) {
	enumValue187 := uint8(x.ActionType)
	io.Uint8(&enumValue187)
	x.ActionType = ItemStackRequestActionType(enumValue187)
	switch int64(enumValue187) {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19:
	default:
		io.InvalidValue(enumValue187, "unknown enum value")
	}
	io.Uint8(&x.Amount)
	x.Source.Marshal(io)
	x.Destination.Marshal(io)
}

// Marshal reads or writes ItemStackRequestPacketDataRequestData using its canonical wire layout.
func (x *ItemStackRequestPacketDataRequestData) Marshal(io IO) {
	x.ClientRequestId.Marshal(io)
	if !io.Reading() && uint64(len(x.Actions)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.Actions), "collection length overflows uint32")
		return
	}
	count188 := uint32(len(x.Actions))
	io.Varuint32(&count188)
	if io.Reading() {
		if uint64(count188) > uint64(^uint(0)>>1) {
			io.InvalidValue(count188, "collection length overflows int")
			return
		}
		x.Actions = make([]ItemStackRequestCereal, int(count188))
	}
	for index189 := range x.Actions {
		marshalItemStackRequestCereal(io, &x.Actions[index189])
	}
	if !io.Reading() && uint64(len(x.StringsToFilter)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.StringsToFilter), "collection length overflows uint32")
		return
	}
	count190 := uint32(len(x.StringsToFilter))
	io.Varuint32(&count190)
	if io.Reading() {
		if uint64(count190) > uint64(^uint(0)>>1) {
			io.InvalidValue(count190, "collection length overflows int")
			return
		}
		x.StringsToFilter = make([]string, int(count190))
	}
	for index191 := range x.StringsToFilter {
		io.String(&x.StringsToFilter[index191])
	}
	enumValue192 := int32(x.StringsToFilterOrigin)
	io.Int32(&enumValue192)
	x.StringsToFilterOrigin = TextProcessingEventOrigin(enumValue192)
	switch int64(enumValue192) {
	case -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15:
	default:
		io.InvalidValue(enumValue192, "unknown enum value")
	}
}

// Marshal reads or writes ItemStackResponseContainerInfo using its canonical wire layout.
func (x *ItemStackResponseContainerInfo) Marshal(io IO) {
	x.FullContainerName.Marshal(io)
	if !io.Reading() && uint64(len(x.Slots)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.Slots), "collection length overflows uint32")
		return
	}
	count193 := uint32(len(x.Slots))
	io.Varuint32(&count193)
	if io.Reading() {
		if uint64(count193) > uint64(^uint(0)>>1) {
			io.InvalidValue(count193, "collection length overflows int")
			return
		}
		x.Slots = make([]ItemStackResponseSlotInfo, int(count193))
	}
	for index194 := range x.Slots {
		x.Slots[index194].Marshal(io)
	}
}

// Marshal reads or writes ItemStackResponseInfo using its canonical wire layout.
func (x *ItemStackResponseInfo) Marshal(io IO) {
	enumValue195 := uint8(x.Result)
	io.Uint8(&enumValue195)
	x.Result = ItemStackNetResult(enumValue195)
	switch int64(enumValue195) {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67:
	default:
		io.InvalidValue(enumValue195, "unknown enum value")
	}
	x.ClientRequestId.Marshal(io)
	outer196 := true
	io.Bool(&outer196)
	if outer196 {
		io.Bool(&x.Containers.set)
		if x.Containers.set {
			if !io.Reading() && uint64(len(x.Containers.val)) > uint64(^uint32(0)) {
				io.InvalidValue(len(x.Containers.val), "collection length overflows uint32")
				return
			}
			count197 := uint32(len(x.Containers.val))
			io.Varuint32(&count197)
			if io.Reading() {
				if uint64(count197) > uint64(^uint(0)>>1) {
					io.InvalidValue(count197, "collection length overflows int")
					return
				}
				x.Containers.val = make([]ItemStackResponseContainerInfo, int(count197))
			}
			for index198 := range x.Containers.val {
				x.Containers.val[index198].Marshal(io)
			}
		} else if io.Reading() {
			var zero []ItemStackResponseContainerInfo
			x.Containers.val = zero
		}
	} else {
		x.Containers = Optional[[]ItemStackResponseContainerInfo]{}
	}
}

// Marshal reads or writes ItemStackResponseSlotInfo using its canonical wire layout.
func (x *ItemStackResponseSlotInfo) Marshal(io IO) {
	io.Uint8(&x.RequestedSlot)
	io.Uint8(&x.Slot)
	io.Uint8(&x.Amount)
	outer199 := true
	io.Bool(&outer199)
	if outer199 {
		io.Bool(&x.ItemStackNetId.set)
		if x.ItemStackNetId.set {
			x.ItemStackNetId.val.Marshal(io)
		} else if io.Reading() {
			var zero TypedServerNetIdStructItemStackNetIdTagInt32T0
			x.ItemStackNetId.val = zero
		}
	} else {
		x.ItemStackNetId = Optional[TypedServerNetIdStructItemStackNetIdTagInt32T0]{}
	}
	x.CustomName.Marshal(io)
	io.Varint32(&x.DurabilityCorrection)
}

// Marshal reads or writes ItemUseInventoryTransaction using its canonical wire layout.
func (x *ItemUseInventoryTransaction) Marshal(io IO) {
	x.Actions.Marshal(io)
	enumValue200 := int32(x.ActionType)
	io.Varint32(&enumValue200)
	x.ActionType = ItemUseInventoryTransactionActionType(enumValue200)
	switch int64(enumValue200) {
	case 0, 1, 2, 3:
	default:
		io.InvalidValue(enumValue200, "unknown enum value")
	}
	enumValue201 := uint8(x.TriggerType)
	io.Uint8(&enumValue201)
	x.TriggerType = ItemUseInventoryTransactionTriggerType(enumValue201)
	switch int64(enumValue201) {
	case 0, 1, 2:
	default:
		io.InvalidValue(enumValue201, "unknown enum value")
	}
	x.Position.Marshal(io)
	io.Uint8(&x.Face)
	io.Varint32(&x.Slot)
	x.Item.Marshal(io)
	io.Vec3(&x.FromPosition)
	io.Vec3(&x.ClickPosition)
	io.Varuint32(&x.TargetBlockId)
	enumValue202 := uint8(x.ClientInteractPrediction)
	io.Uint8(&enumValue202)
	x.ClientInteractPrediction = ItemUseInventoryTransactionPredictedResult(enumValue202)
	switch int64(enumValue202) {
	case 0, 1:
	default:
		io.InvalidValue(enumValue202, "unknown enum value")
	}
	enumValue203 := uint8(x.ClientCooldownState)
	io.Uint8(&enumValue203)
	x.ClientCooldownState = ItemUseInventoryTransactionClientCooldownState(enumValue203)
	switch int64(enumValue203) {
	case 0, 1:
	default:
		io.InvalidValue(enumValue203, "unknown enum value")
	}
}

// Marshal reads or writes ItemUseOnActorInventoryTransaction using its canonical wire layout.
func (x *ItemUseOnActorInventoryTransaction) Marshal(io IO) {
	x.Actions.Marshal(io)
	x.RuntimeId.Marshal(io)
	enumValue204 := int32(x.ActionType)
	io.Varint32(&enumValue204)
	x.ActionType = ItemUseOnActorInventoryTransactionActionType(enumValue204)
	switch int64(enumValue204) {
	case 0, 1, 2:
	default:
		io.InvalidValue(enumValue204, "unknown enum value")
	}
	io.Varint32(&x.Slot)
	x.Item.Marshal(io)
	io.Vec3(&x.FromPosition)
	io.Vec3(&x.HitPosition)
}

// Marshal reads or writes LegacySetSlot using its canonical wire layout.
func (x *LegacySetSlot) Marshal(io IO) {
	enumValue205 := uint8(x.ContainerEnum)
	io.Uint8(&enumValue205)
	x.ContainerEnum = ContainerEnumName(enumValue205)
	switch int64(enumValue205) {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66:
	default:
		io.InvalidValue(enumValue205, "unknown enum value")
	}
	if !io.Reading() && uint64(len(x.Slots)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.Slots), "collection length overflows uint32")
		return
	}
	count206 := uint32(len(x.Slots))
	io.Varuint32(&count206)
	if io.Reading() {
		if uint64(count206) > uint64(^uint(0)>>1) {
			io.InvalidValue(count206, "collection length overflows int")
			return
		}
		x.Slots = make([]uint8, int(count206))
	}
	for index207 := range x.Slots {
		io.Uint8(&x.Slots[index207])
	}
}

// Marshal reads or writes LegacyTelemetryEventAchievement using its canonical wire layout.
func (x *LegacyTelemetryEventAchievement) Marshal(io IO) {
	enumValue208 := uint8(x.AchievementID)
	io.Uint8(&enumValue208)
	x.AchievementID = MinecraftEventingAchievementIds(enumValue208)
	switch int64(enumValue208) {
	case 7, 10, 20, 21, 29, 30, 37, 38, 39, 40, 50, 52, 53, 54, 56, 58, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 91, 92, 93, 94, 95, 96, 97, 98, 99, 100, 101, 102, 103, 104, 105, 106, 107, 108, 109, 110, 111, 112, 113, 114, 115, 116, 117, 118, 119, 120, 121, 122, 123, 124, 125, 126:
	default:
		io.InvalidValue(enumValue208, "unknown enum value")
	}
}

// Marshal reads or writes LegacyTelemetryEventActorDefinition using its canonical wire layout.
func (x *LegacyTelemetryEventActorDefinition) Marshal(io IO) {
	io.String(&x.EventName)
}

// Marshal reads or writes LegacyTelemetryEventBellUsed using its canonical wire layout.
func (x *LegacyTelemetryEventBellUsed) Marshal(io IO) {
	io.Varint32(&x.ItemId)
}

// Marshal reads or writes LegacyTelemetryEventBossKilled using its canonical wire layout.
func (x *LegacyTelemetryEventBossKilled) Marshal(io IO) {
	io.Varint64(&x.BossActorID)
	io.Varint32(&x.PartySize)
	io.Varint32(&x.BossType)
}

// Marshal reads or writes LegacyTelemetryEventCauldronUsed using its canonical wire layout.
func (x *LegacyTelemetryEventCauldronUsed) Marshal(io IO) {
	io.Varuint32(&x.ContentsColor)
	io.Varint32(&x.ContentsType)
	io.Varint32(&x.FillLevel)
}

// Marshal reads or writes LegacyTelemetryEventCodeBuilderRuntimeAction using its canonical wire layout.
func (x *LegacyTelemetryEventCodeBuilderRuntimeAction) Marshal(io IO) {
	io.String(&x.CodeBuilderRuntimeAction)
}

// Marshal reads or writes LegacyTelemetryEventCodeBuilderScoreboard using its canonical wire layout.
func (x *LegacyTelemetryEventCodeBuilderScoreboard) Marshal(io IO) {
	io.String(&x.ObjectiveName)
	io.Varint32(&x.Score)
}

// Marshal reads or writes LegacyTelemetryEventComposterUsed using its canonical wire layout.
func (x *LegacyTelemetryEventComposterUsed) Marshal(io IO) {
	enumValue209 := uint8(x.BlockInteractionType)
	io.Uint8(&enumValue209)
	x.BlockInteractionType = MinecraftEventingPOIBlockInteractionType(enumValue209)
	switch int64(enumValue209) {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25:
	default:
		io.InvalidValue(enumValue209, "unknown enum value")
	}
	io.Varint32(&x.ItemId)
}

// Marshal reads or writes LegacyTelemetryEventEmpty using its canonical wire layout.
func (x *LegacyTelemetryEventEmpty) Marshal(io IO) {
}

func marshalLegacyTelemetryEventEventData(io IO, x *LegacyTelemetryEventEventData) {
	if io.Reading() {
		var tag uint32
		io.Varuint32(&tag)
		switch int64(tag) {
		case 0:
			var value LegacyTelemetryEventAchievement
			value.Marshal(io)
			*x = value
		case 1:
			var value LegacyTelemetryEventInteraction
			value.Marshal(io)
			*x = value
		case 2:
			var value LegacyTelemetryEventPortalCreated
			value.Marshal(io)
			*x = value
		case 3:
			var value LegacyTelemetryEventPortalUsed
			value.Marshal(io)
			*x = value
		case 4:
			var value LegacyTelemetryEventMobKilled
			value.Marshal(io)
			*x = value
		case 5:
			var value LegacyTelemetryEventCauldronUsed
			value.Marshal(io)
			*x = value
		case 6:
			var value LegacyTelemetryEventPlayerDied
			value.Marshal(io)
			*x = value
		case 7:
			var value LegacyTelemetryEventBossKilled
			value.Marshal(io)
			*x = value
		case 8:
			var value LegacyTelemetryEventSlashCommand
			value.Marshal(io)
			*x = value
		case 9:
			var value LegacyTelemetryEventMobBorn
			value.Marshal(io)
			*x = value
		case 10:
			var value LegacyTelemetryEventPOICauldronUsed
			value.Marshal(io)
			*x = value
		case 11:
			var value LegacyTelemetryEventComposterUsed
			value.Marshal(io)
			*x = value
		case 12:
			var value LegacyTelemetryEventBellUsed
			value.Marshal(io)
			*x = value
		case 13:
			var value LegacyTelemetryEventActorDefinition
			value.Marshal(io)
			*x = value
		case 14:
			var value LegacyTelemetryEventRaidUpdate
			value.Marshal(io)
			*x = value
		case 15:
			var value LegacyTelemetryEventTargetBlockHit
			value.Marshal(io)
			*x = value
		case 16:
			var value LegacyTelemetryEventPiglinBarter
			value.Marshal(io)
			*x = value
		case 17:
			var value LegacyTelemetryEventPlayerWaxedOrUnwaxedCopper
			value.Marshal(io)
			*x = value
		case 18:
			var value LegacyTelemetryEventCodeBuilderRuntimeAction
			value.Marshal(io)
			*x = value
		case 19:
			var value LegacyTelemetryEventCodeBuilderScoreboard
			value.Marshal(io)
			*x = value
		case 20:
			var value LegacyTelemetryEventItemUsed
			value.Marshal(io)
			*x = value
		case 21:
			var value LegacyTelemetryEventEmpty
			value.Marshal(io)
			*x = value
		default:
			io.InvalidValue(tag, "unknown union tag")
		}
		return
	}
	switch value := (*x).(type) {
	case LegacyTelemetryEventAchievement:
		tag := uint32(0)
		io.Varuint32(&tag)
		value.Marshal(io)
	case LegacyTelemetryEventInteraction:
		tag := uint32(1)
		io.Varuint32(&tag)
		value.Marshal(io)
	case LegacyTelemetryEventPortalCreated:
		tag := uint32(2)
		io.Varuint32(&tag)
		value.Marshal(io)
	case LegacyTelemetryEventPortalUsed:
		tag := uint32(3)
		io.Varuint32(&tag)
		value.Marshal(io)
	case LegacyTelemetryEventMobKilled:
		tag := uint32(4)
		io.Varuint32(&tag)
		value.Marshal(io)
	case LegacyTelemetryEventCauldronUsed:
		tag := uint32(5)
		io.Varuint32(&tag)
		value.Marshal(io)
	case LegacyTelemetryEventPlayerDied:
		tag := uint32(6)
		io.Varuint32(&tag)
		value.Marshal(io)
	case LegacyTelemetryEventBossKilled:
		tag := uint32(7)
		io.Varuint32(&tag)
		value.Marshal(io)
	case LegacyTelemetryEventSlashCommand:
		tag := uint32(8)
		io.Varuint32(&tag)
		value.Marshal(io)
	case LegacyTelemetryEventMobBorn:
		tag := uint32(9)
		io.Varuint32(&tag)
		value.Marshal(io)
	case LegacyTelemetryEventPOICauldronUsed:
		tag := uint32(10)
		io.Varuint32(&tag)
		value.Marshal(io)
	case LegacyTelemetryEventComposterUsed:
		tag := uint32(11)
		io.Varuint32(&tag)
		value.Marshal(io)
	case LegacyTelemetryEventBellUsed:
		tag := uint32(12)
		io.Varuint32(&tag)
		value.Marshal(io)
	case LegacyTelemetryEventActorDefinition:
		tag := uint32(13)
		io.Varuint32(&tag)
		value.Marshal(io)
	case LegacyTelemetryEventRaidUpdate:
		tag := uint32(14)
		io.Varuint32(&tag)
		value.Marshal(io)
	case LegacyTelemetryEventTargetBlockHit:
		tag := uint32(15)
		io.Varuint32(&tag)
		value.Marshal(io)
	case LegacyTelemetryEventPiglinBarter:
		tag := uint32(16)
		io.Varuint32(&tag)
		value.Marshal(io)
	case LegacyTelemetryEventPlayerWaxedOrUnwaxedCopper:
		tag := uint32(17)
		io.Varuint32(&tag)
		value.Marshal(io)
	case LegacyTelemetryEventCodeBuilderRuntimeAction:
		tag := uint32(18)
		io.Varuint32(&tag)
		value.Marshal(io)
	case LegacyTelemetryEventCodeBuilderScoreboard:
		tag := uint32(19)
		io.Varuint32(&tag)
		value.Marshal(io)
	case LegacyTelemetryEventItemUsed:
		tag := uint32(20)
		io.Varuint32(&tag)
		value.Marshal(io)
	case LegacyTelemetryEventEmpty:
		tag := uint32(21)
		io.Varuint32(&tag)
		value.Marshal(io)
	default:
		io.InvalidValue(*x, "unknown union value")
	}
}

// Marshal reads or writes LegacyTelemetryEventInteraction using its canonical wire layout.
func (x *LegacyTelemetryEventInteraction) Marshal(io IO) {
	io.Varint64(&x.InteractedEntityID)
	enumValue210 := uint8(x.InteractionType)
	io.Uint8(&enumValue210)
	x.InteractionType = MinecraftEventingInteractionType(enumValue210)
	switch int64(enumValue210) {
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17:
	default:
		io.InvalidValue(enumValue210, "unknown enum value")
	}
	io.Varint32(&x.InteractionActorType)
	io.Varint32(&x.InteractionActorVariant)
	io.Uint8(&x.InteractionActorColor)
}

// Marshal reads or writes LegacyTelemetryEventItemUsed using its canonical wire layout.
func (x *LegacyTelemetryEventItemUsed) Marshal(io IO) {
	io.Int16(&x.ItemId)
	io.Int32(&x.ItemAux)
	io.Int32(&x.UseMethod)
	io.Int32(&x.Count)
}

// Marshal reads or writes LegacyTelemetryEventMobBorn using its canonical wire layout.
func (x *LegacyTelemetryEventMobBorn) Marshal(io IO) {
	io.Varint32(&x.BornBabyEntityType)
	io.Varint32(&x.BornBabyEntityVariant)
	io.Uint8(&x.BornBabyColor)
}

// Marshal reads or writes LegacyTelemetryEventMobKilled using its canonical wire layout.
func (x *LegacyTelemetryEventMobKilled) Marshal(io IO) {
	io.Varint64(&x.InstigatorActorID)
	io.Varint64(&x.TargetActorID)
	enumValue211 := int32(x.InstigatorSChildActorType)
	io.Varint32(&enumValue211)
	x.InstigatorSChildActorType = ActorType(enumValue211)
	switch int64(enumValue211) {
	case 1, 64, 65, 66, 67, 69, 70, 71, 72, 77, 78, 83, 88, 90, 93, 95, 107, 117, 119, 145, 154, 218, 256, 307, 312, 317, 318, 319, 378, 379, 383, 390, 768, 788, 789, 886, 916, 921, 2816, 2849, 2853, 2854, 2857, 2858, 2859, 2861, 2865, 2866, 2869, 2870, 2873, 2875, 2920, 2921, 2930, 2936, 2947, 2956, 2962, 4864, 4874, 4875, 4876, 4877, 4880, 4882, 4892, 4893, 4938, 4977, 4985, 4988, 4989, 4992, 4994, 4996, 5002, 5003, 5006, 5011, 5021, 8960, 8977, 8991, 9068, 9069, 9071, 9072, 9089, 9093, 9109, 21248, 21262, 21270, 21278, 21323, 33024, 33043, 68352, 68388, 68404, 68410, 68478, 70552, 74646, 199424, 199456, 199468, 199471, 199534, 199540, 264960, 264995, 264999, 265000, 265015, 524288, 524372, 524384, 524385, 524386, 524387, 524388, 1116928, 1116962, 1116974, 1116976, 1117072, 1117079, 2118400, 2118423, 2118424, 2118425, 2183962, 2183963, 4194304, 4194372, 4194380, 4194383, 4194385, 4194386, 4194389, 4194390, 4194391, 4194393, 4194395, 4194398, 4194405, 4194406, 4194407, 4194410, 4194445, 4194447, 8388608, 12582985, 12582992, 16777984, 16777999, 16778099:
	default:
		io.InvalidValue(enumValue211, "unknown enum value")
	}
	io.Varint32(&x.DamageSource)
	io.Varint32(&x.TradeTier)
	io.String(&x.TraderName)
}

// Marshal reads or writes LegacyTelemetryEventPOICauldronUsed using its canonical wire layout.
func (x *LegacyTelemetryEventPOICauldronUsed) Marshal(io IO) {
	enumValue212 := uint8(x.BlockInteractionType)
	io.Uint8(&enumValue212)
	x.BlockInteractionType = MinecraftEventingPOIBlockInteractionType(enumValue212)
	switch int64(enumValue212) {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25:
	default:
		io.InvalidValue(enumValue212, "unknown enum value")
	}
	io.Varint32(&x.ItemId)
}

// Marshal reads or writes LegacyTelemetryEventPiglinBarter using its canonical wire layout.
func (x *LegacyTelemetryEventPiglinBarter) Marshal(io IO) {
	io.Varint32(&x.ItemId)
	io.Bool(&x.WasTargetingBarteringPlayer)
}

// Marshal reads or writes LegacyTelemetryEventPlayerDied using its canonical wire layout.
func (x *LegacyTelemetryEventPlayerDied) Marshal(io IO) {
	io.Varint32(&x.InstigatorActorID)
	io.Varint32(&x.InstigatorMobVariant)
	io.Varint32(&x.DamageSource)
	io.Bool(&x.DiedInRaid)
}

// Marshal reads or writes LegacyTelemetryEventPlayerWaxedOrUnwaxedCopper using its canonical wire layout.
func (x *LegacyTelemetryEventPlayerWaxedOrUnwaxedCopper) Marshal(io IO) {
	io.Varint32(&x.PlayerWaxedOrUnwaxedCopperBlockID)
}

// Marshal reads or writes LegacyTelemetryEventPortalCreated using its canonical wire layout.
func (x *LegacyTelemetryEventPortalCreated) Marshal(io IO) {
	io.Varint32(&x.DimensionID)
}

// Marshal reads or writes LegacyTelemetryEventPortalUsed using its canonical wire layout.
func (x *LegacyTelemetryEventPortalUsed) Marshal(io IO) {
	io.Varint32(&x.SourceDimensionID)
	io.Varint32(&x.TargetDimensionID)
}

// Marshal reads or writes LegacyTelemetryEventRaidUpdate using its canonical wire layout.
func (x *LegacyTelemetryEventRaidUpdate) Marshal(io IO) {
	io.Varint32(&x.CurrentWave)
	io.Varint32(&x.TotalWaves)
	io.Bool(&x.Success)
}

// Marshal reads or writes LegacyTelemetryEventSlashCommand using its canonical wire layout.
func (x *LegacyTelemetryEventSlashCommand) Marshal(io IO) {
	io.Varint32(&x.SuccessCount)
	io.Varint32(&x.ErrorCount)
	io.String(&x.CommandName)
	io.String(&x.ErrorList)
}

// Marshal reads or writes LegacyTelemetryEventTargetBlockHit using its canonical wire layout.
func (x *LegacyTelemetryEventTargetBlockHit) Marshal(io IO) {
	io.Varint32(&x.RedstoneLevel)
}

// Marshal reads or writes LevelChunkSubChunkMetadata using its canonical wire layout.
func (x *LevelChunkSubChunkMetadata) Marshal(io IO) {
	io.Uint64(&x.BlobId)
}

// Marshal reads or writes LevelSettings using its canonical wire layout.
func (x *LevelSettings) Marshal(io IO) {
	io.Uint64(&x.Seed)
	x.SpawnSettings.Marshal(io)
	enumValue213 := int32(x.GeneratorType)
	io.Varint32(&enumValue213)
	x.GeneratorType = GeneratorType(enumValue213)
	switch int64(enumValue213) {
	case 0, 1, 2, 3, 4, 5, 6:
	default:
		io.InvalidValue(enumValue213, "unknown enum value")
	}
	enumValue214 := int32(x.GameType)
	io.Varint32(&enumValue214)
	x.GameType = GameType(enumValue214)
	switch int64(enumValue214) {
	case -1, 0, 1, 2, 5, 6:
	default:
		io.InvalidValue(enumValue214, "unknown enum value")
	}
	io.Bool(&x.IsHardcore)
	enumValue215 := int32(x.GameDifficulty)
	io.Varint32(&enumValue215)
	x.GameDifficulty = LegacyDifficulty(enumValue215)
	switch int64(enumValue215) {
	case 0, 1, 2, 3, 4, 5:
	default:
		io.InvalidValue(enumValue215, "unknown enum value")
	}
	x.DefaultSpawnBlockPosition.Marshal(io)
	io.Bool(&x.AchievementsDisabled)
	enumValue216 := int32(x.EditorWorldType)
	io.Varint32(&enumValue216)
	x.EditorWorldType = EditorWorldType(enumValue216)
	switch int64(enumValue216) {
	case 0, 1, 2, 3:
	default:
		io.InvalidValue(enumValue216, "unknown enum value")
	}
	io.Bool(&x.IsCreatedInEditor)
	io.Bool(&x.IsExportedFromEditor)
	io.Varint32(&x.DayCycleStopTime)
	enumValue217 := uint32(x.EducationEditionOffer)
	io.Varuint32(&enumValue217)
	x.EducationEditionOffer = EducationEditionOffer(enumValue217)
	switch int64(enumValue217) {
	case 0, 1, 2:
	default:
		io.InvalidValue(enumValue217, "unknown enum value")
	}
	io.Bool(&x.EducationFeaturesEnabled)
	io.String(&x.EducationProductID)
	io.Float32(&x.RainLevel)
	io.Float32(&x.LightningLevel)
	io.Bool(&x.HasConfirmedPlatformLockedContent)
	io.Bool(&x.MultiplayerGameIntent)
	io.Bool(&x.LANBroadcastIntent)
	enumValue218 := int32(x.XboxLiveBroadcastSetting)
	io.Varint32(&enumValue218)
	x.XboxLiveBroadcastSetting = SocialGamePublishSetting(enumValue218)
	switch int64(enumValue218) {
	case 0, 1, 2, 3, 4:
	default:
		io.InvalidValue(enumValue218, "unknown enum value")
	}
	enumValue219 := int32(x.PlatformBroadcastSetting)
	io.Varint32(&enumValue219)
	x.PlatformBroadcastSetting = SocialGamePublishSetting(enumValue219)
	switch int64(enumValue219) {
	case 0, 1, 2, 3, 4:
	default:
		io.InvalidValue(enumValue219, "unknown enum value")
	}
	io.Bool(&x.CommandsEnabled)
	io.Bool(&x.TexturePacksRequired)
	x.RuleData.Marshal(io)
	x.Experiments.Marshal(io)
	io.Bool(&x.HasBonusChestEnabled)
	io.Bool(&x.StartWithMapEnabled)
	enumValue220 := int8(x.PlayerPermissions)
	io.Int8(&enumValue220)
	x.PlayerPermissions = PlayerPermissionLevel(enumValue220)
	switch int64(enumValue220) {
	case 0, 1, 2, 3:
	default:
		io.InvalidValue(enumValue220, "unknown enum value")
	}
	io.Int32(&x.ServerChunkTickRange)
	io.Bool(&x.HasLockedBehaviorPack)
	io.Bool(&x.HasLockedResourcePack)
	io.Bool(&x.IsFromLockedTemplate)
	io.Bool(&x.UseMsaGamertagsOnly)
	io.Bool(&x.IsFromWorldTemplate)
	io.Bool(&x.IsWorldTemplateOptionLocked)
	io.Bool(&x.OnlySpawnV1Villagers)
	io.Bool(&x.PersonaDisabled)
	io.Bool(&x.CustomSkinsDisabled)
	io.Bool(&x.EmoteChatMuted)
	io.String(&x.BaseGameVersion)
	io.Int32(&x.LimitedWorldWidth)
	io.Int32(&x.LimitedWorldDepth)
	io.Bool(&x.NetherType)
	x.EduSharedUriResource.Marshal(io)
	io.Bool(&x.OverrideForceExperimentalGameplay.set)
	if x.OverrideForceExperimentalGameplay.set {
		io.Bool(&x.OverrideForceExperimentalGameplay.val)
	} else if io.Reading() {
		var zero bool
		x.OverrideForceExperimentalGameplay.val = zero
	}
	enumValue221 := uint8(x.ChatRestrictionLevel)
	io.Uint8(&enumValue221)
	x.ChatRestrictionLevel = ChatRestrictionLevel(enumValue221)
	switch int64(enumValue221) {
	case 0, 1, 2:
	default:
		io.InvalidValue(enumValue221, "unknown enum value")
	}
	io.Bool(&x.DisablePlayerInteractions)
	enumValue222 := int32(x.ServerEditorConnectionPolicy)
	io.Varint32(&enumValue222)
	x.ServerEditorConnectionPolicy = ServerEditorConnectionPolicy(enumValue222)
	switch int64(enumValue222) {
	case 0, 1, 2, 3:
	default:
		io.InvalidValue(enumValue222, "unknown enum value")
	}
	io.Bool(&x.AllowAnonymousBlockDropsInEditorWorlds)
}

// Marshal reads or writes LineData using its canonical wire layout.
func (x *LineData) Marshal(io IO) {
	io.Vec3(&x.LineEndLocation)
}

// Marshal reads or writes LocatorBarWaypoint using its canonical wire layout.
func (x *LocatorBarWaypoint) Marshal(io IO) {
	x.GroupHandle.Marshal(io)
	x.ServerWaypointPayload.Marshal(io)
	enumValue223 := uint8(x.ActionFlag)
	io.Uint8(&enumValue223)
	x.ActionFlag = ServerWaypointGroupAction(enumValue223)
	switch int64(enumValue223) {
	case 0, 1, 2, 3:
	default:
		io.InvalidValue(enumValue223, "unknown enum value")
	}
}

// Marshal reads or writes MapDecoration using its canonical wire layout.
func (x *MapDecoration) Marshal(io IO) {
	enumValue224 := int8(x.ImageType)
	io.Int8(&enumValue224)
	x.ImageType = MapDecorationType(enumValue224)
	switch int64(enumValue224) {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25:
	default:
		io.InvalidValue(enumValue224, "unknown enum value")
	}
	io.Uint8(&x.Rotation)
	io.Uint8(&x.X)
	io.Uint8(&x.Y)
	io.String(&x.Label)
	io.RGBA(&x.Color)
}

// Marshal reads or writes MapInfoRequestPacketAnonClientPixelsProxy using its canonical wire layout.
func (x *MapInfoRequestPacketAnonClientPixelsProxy) Marshal(io IO) {
	io.Uint32(&x.Pixel)
	io.Uint16(&x.Index)
}

// Marshal reads or writes MapItemTrackedActorUniqueId using its canonical wire layout.
func (x *MapItemTrackedActorUniqueId) Marshal(io IO) {
	enumValue225 := int32(x.Type)
	io.Int32(&enumValue225)
	x.Type = MapItemTrackedActorType(enumValue225)
	switch int64(enumValue225) {
	case 0, 1, 2:
	default:
		io.InvalidValue(enumValue225, "unknown enum value")
	}
	io.Bool(&x.EntityID.set)
	if x.EntityID.set {
		x.EntityID.val.Marshal(io)
	} else if io.Reading() {
		var zero ActorUniqueID
		x.EntityID.val = zero
	}
	io.Bool(&x.BlockPosition.set)
	if x.BlockPosition.set {
		x.BlockPosition.val.Marshal(io)
	} else if io.Reading() {
		var zero BlockPos
		x.BlockPosition.val = zero
	}
}

// Marshal reads or writes MaterialReducerDataEntry using its canonical wire layout.
func (x *MaterialReducerDataEntry) Marshal(io IO) {
	io.Varint32(&x.FromItemKey)
	if !io.Reading() && uint64(len(x.ItemIdsAndCounts)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.ItemIdsAndCounts), "collection length overflows uint32")
		return
	}
	count226 := uint32(len(x.ItemIdsAndCounts))
	io.Varuint32(&count226)
	if io.Reading() {
		if uint64(count226) > uint64(^uint(0)>>1) {
			io.InvalidValue(count226, "collection length overflows int")
			return
		}
		x.ItemIdsAndCounts = make([]MaterialReducerEntryOutput, int(count226))
	}
	for index227 := range x.ItemIdsAndCounts {
		x.ItemIdsAndCounts[index227].Marshal(io)
	}
}

// Marshal reads or writes MaterialReducerEntryOutput using its canonical wire layout.
func (x *MaterialReducerEntryOutput) Marshal(io IO) {
	io.Varint32(&x.ItemId)
	io.Varint32(&x.ItemCount)
}

// Marshal reads or writes MemoryMemoryCategoryCounter using its canonical wire layout.
func (x *MemoryMemoryCategoryCounter) Marshal(io IO) {
	enumValue228 := uint8(x.Category)
	io.Uint8(&enumValue228)
	x.Category = MemoryMemoryCategory(enumValue228)
	switch int64(enumValue228) {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 91, 92, 93, 94, 95, 96, 97, 98, 99, 100, 101, 102, 103, 104, 105, 106, 107, 108, 109, 110:
	default:
		io.InvalidValue(enumValue228, "unknown enum value")
	}
	io.Uint64(&x.CurrentBytes)
}

// Marshal reads or writes MissingBlobData using its canonical wire layout.
func (x *MissingBlobData) Marshal(io IO) {
	io.Uint64(&x.BlobId)
	io.Bytes(&x.BlobData)
}

// Marshal reads or writes MoveActorAbsoluteData using its canonical wire layout.
func (x *MoveActorAbsoluteData) Marshal(io IO) {
	x.ActorRuntimeID.Marshal(io)
	io.Uint8(&x.Header)
	io.Vec3(&x.Position)
	io.Uint8(&x.RotationX)
	io.Uint8(&x.RotationY)
	io.Uint8(&x.RotationYHead)
}

// Marshal reads or writes MoveActorDeltaData using its canonical wire layout.
func (x *MoveActorDeltaData) Marshal(io IO) {
	x.ActorRuntimeID.Marshal(io)
	io.Bool(&x.NewPositionX.set)
	if x.NewPositionX.set {
		io.Float32(&x.NewPositionX.val)
	} else if io.Reading() {
		var zero float32
		x.NewPositionX.val = zero
	}
	io.Bool(&x.NewPositionY.set)
	if x.NewPositionY.set {
		io.Float32(&x.NewPositionY.val)
	} else if io.Reading() {
		var zero float32
		x.NewPositionY.val = zero
	}
	io.Bool(&x.NewPositionZ.set)
	if x.NewPositionZ.set {
		io.Float32(&x.NewPositionZ.val)
	} else if io.Reading() {
		var zero float32
		x.NewPositionZ.val = zero
	}
	io.Bool(&x.RotationX.set)
	if x.RotationX.set {
		io.Int8(&x.RotationX.val)
	} else if io.Reading() {
		var zero int8
		x.RotationX.val = zero
	}
	io.Bool(&x.RotationY.set)
	if x.RotationY.set {
		io.Int8(&x.RotationY.val)
	} else if io.Reading() {
		var zero int8
		x.RotationY.val = zero
	}
	io.Bool(&x.RotationYHead.set)
	if x.RotationYHead.set {
		io.Int8(&x.RotationYHead.val)
	} else if io.Reading() {
		var zero int8
		x.RotationYHead.val = zero
	}
	io.Bool(&x.IsOnGround)
	io.Bool(&x.ForceMove)
	io.Bool(&x.ForceMoveLocalEntity)
	io.Bool(&x.ForceCompletion)
}

// Marshal reads or writes MovePlayerTeleportData using its canonical wire layout.
func (x *MovePlayerTeleportData) Marshal(io IO) {
	io.Int32(&x.TeleportationCause)
	io.Int32(&x.SourceActorType)
}

// Marshal reads or writes MultiRecipe using its canonical wire layout.
func (x *MultiRecipe) Marshal(io IO) {
	io.UUID(&x.MultiRecipeUUID)
	x.NetId.Marshal(io)
}

// Marshal reads or writes NetworkPermissions using its canonical wire layout.
func (x *NetworkPermissions) Marshal(io IO) {
	io.Bool(&x.ServerAuthSoundEnabled)
}

// Marshal reads or writes NoiseDescriptor using its canonical wire layout.
func (x *NoiseDescriptor) Marshal(io IO) {
	io.String(&x.Name)
	io.Int32(&x.FirstOctave)
	if !io.Reading() && uint64(len(x.Amplitudes)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.Amplitudes), "collection length overflows uint32")
		return
	}
	count229 := uint32(len(x.Amplitudes))
	io.Varuint32(&count229)
	if io.Reading() {
		if uint64(count229) > uint64(^uint(0)>>1) {
			io.InvalidValue(count229, "collection length overflows int")
			return
		}
		x.Amplitudes = make([]float32, int(count229))
	}
	for index230 := range x.Amplitudes {
		io.Float32(&x.Amplitudes[index230])
	}
}

// Marshal reads or writes NormalTransactionData using its canonical wire layout.
func (x *NormalTransactionData) Marshal(io IO) {
	x.Actions.Marshal(io)
}

// Marshal reads or writes PackIdVersion using its canonical wire layout.
func (x *PackIdVersion) Marshal(io IO) {
	io.UUID(&x.PackUUID)
	x.PackVersion.Marshal(io)
}

// Marshal reads or writes PackIdVersionData using its canonical wire layout.
func (x *PackIdVersionData) Marshal(io IO) {
	io.UUID(&x.PackUUID)
	x.PackVersion.Marshal(io)
}

// Marshal reads or writes PackInfoData using its canonical wire layout.
func (x *PackInfoData) Marshal(io IO) {
	x.PackIdVersion.Marshal(io)
	io.Uint64(&x.PackSize)
	io.String(&x.ContentKey)
	io.String(&x.SubpackName)
	x.ContentIdentity.Marshal(io)
	io.Bool(&x.HasScripts)
	io.Bool(&x.IsAddonPack)
	io.Bool(&x.IsRayTracingCapable)
	io.String(&x.CDNURL)
}

// Marshal reads or writes PackInstanceId using its canonical wire layout.
func (x *PackInstanceId) Marshal(io IO) {
	io.String(&x.PackID)
	io.String(&x.Version)
	io.String(&x.SubPackName)
}

// Marshal reads or writes PackedItemUseLegacyInventoryTransaction using its canonical wire layout.
func (x *PackedItemUseLegacyInventoryTransaction) Marshal(io IO) {
	x.LegacyRequestID.Marshal(io)
	io.Bool(&x.LegacySetItemSlots.set)
	if x.LegacySetItemSlots.set {
		if !io.Reading() && uint64(len(x.LegacySetItemSlots.val)) > uint64(^uint32(0)) {
			io.InvalidValue(len(x.LegacySetItemSlots.val), "collection length overflows uint32")
			return
		}
		count231 := uint32(len(x.LegacySetItemSlots.val))
		io.Varuint32(&count231)
		if io.Reading() {
			if uint64(count231) > uint64(^uint(0)>>1) {
				io.InvalidValue(count231, "collection length overflows int")
				return
			}
			x.LegacySetItemSlots.val = make([]LegacySetSlot, int(count231))
		}
		for index232 := range x.LegacySetItemSlots.val {
			x.LegacySetItemSlots.val[index232].Marshal(io)
		}
	} else if io.Reading() {
		var zero []LegacySetSlot
		x.LegacySetItemSlots.val = zero
	}
	io.Bool(&x.ItemUseTransaction.set)
	if x.ItemUseTransaction.set {
		x.ItemUseTransaction.val.Marshal(io)
	} else if io.Reading() {
		var zero ItemUseInventoryTransaction
		x.ItemUseTransaction.val = zero
	}
}

// Marshal reads or writes PlayerBlockActionData using its canonical wire layout.
func (x *PlayerBlockActionData) Marshal(io IO) {
	enumValue233 := int32(x.PlayerActionType)
	io.Varint32(&enumValue233)
	x.PlayerActionType = PlayerActionType(enumValue233)
	switch int64(enumValue233) {
	case -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39:
	default:
		io.InvalidValue(enumValue233, "unknown enum value")
	}
	x.Position.Marshal(io)
	io.Varint32(&x.Facing)
}

// Marshal reads or writes PlayerInputTick using its canonical wire layout.
func (x *PlayerInputTick) Marshal(io IO) {
	io.Varuint64(&x.InputTick)
}

// Marshal reads or writes PlayerListAddEntry using its canonical wire layout.
func (x *PlayerListAddEntry) Marshal(io IO) {
	io.UUID(&x.UUID)
	x.ActorUniqueID.Marshal(io)
	io.String(&x.PlayerName)
	io.String(&x.XBLXUID)
	io.String(&x.PlatformOnlineID)
	enumValue234 := int32(x.BuildPlatform)
	io.Int32(&enumValue234)
	x.BuildPlatform = BuildPlatform(enumValue234)
	switch int64(enumValue234) {
	case -1, 1, 2, 3, 4, 5, 7, 8, 9, 10, 11, 12, 13, 14, 15:
	default:
		io.InvalidValue(enumValue234, "unknown enum value")
	}
	x.SerializedSkin.Marshal(io)
	io.Bool(&x.IsTeacher)
	io.Bool(&x.IsHost)
	io.Bool(&x.IsSubClient)
	io.RGBA(&x.PlayerColor)
}

func marshalPlayerListEntriesItem(io IO, x *PlayerListEntriesItem) {
	if io.Reading() {
		var tag uint8
		io.Uint8(&tag)
		switch int64(tag) {
		case 0:
			var value PlayerListAddEntry
			value.Marshal(io)
			*x = value
		case 1:
			var value PlayerListRemoveEntry
			value.Marshal(io)
			*x = value
		default:
			io.InvalidValue(tag, "unknown union tag")
		}
		return
	}
	switch value := (*x).(type) {
	case PlayerListAddEntry:
		tag := uint8(0)
		io.Uint8(&tag)
		value.Marshal(io)
	case PlayerListRemoveEntry:
		tag := uint8(1)
		io.Uint8(&tag)
		value.Marshal(io)
	default:
		io.InvalidValue(*x, "unknown union value")
	}
}

// Marshal reads or writes PlayerListRemoveEntry using its canonical wire layout.
func (x *PlayerListRemoveEntry) Marshal(io IO) {
	io.UUID(&x.UUID)
}

// Marshal reads or writes PlayerLocationCoordinatesLocation using its canonical wire layout.
func (x *PlayerLocationCoordinatesLocation) Marshal(io IO) {
	enumValue235 := int32(x.PacketType)
	io.Varint32(&enumValue235)
	x.PacketType = PlayerLocationType(enumValue235)
	switch int64(enumValue235) {
	case 0:
	default:
		io.InvalidValue(enumValue235, "unknown enum value")
	}
	io.Vec3(&x.Position)
}

// Marshal reads or writes PlayerLocationHiddenLocation using its canonical wire layout.
func (x *PlayerLocationHiddenLocation) Marshal(io IO) {
	enumValue236 := int32(x.PacketType)
	io.Varint32(&enumValue236)
	x.PacketType = PlayerLocationType(enumValue236)
	switch int64(enumValue236) {
	case 1:
	default:
		io.InvalidValue(enumValue236, "unknown enum value")
	}
}

func marshalPlayerLocationLocation(io IO, x *PlayerLocationLocation) {
	if io.Reading() {
		var tag uint32
		io.Varuint32(&tag)
		switch int64(tag) {
		case 0:
			var value PlayerLocationCoordinatesLocation
			value.Marshal(io)
			*x = value
		case 1:
			var value PlayerLocationHiddenLocation
			value.Marshal(io)
			*x = value
		default:
			io.InvalidValue(tag, "unknown union tag")
		}
		return
	}
	switch value := (*x).(type) {
	case PlayerLocationCoordinatesLocation:
		tag := uint32(0)
		io.Varuint32(&tag)
		value.Marshal(io)
	case PlayerLocationHiddenLocation:
		tag := uint32(1)
		io.Varuint32(&tag)
		value.Marshal(io)
	default:
		io.InvalidValue(*x, "unknown union value")
	}
}

// Marshal reads or writes PlayerPartyInfo using its canonical wire layout.
func (x *PlayerPartyInfo) Marshal(io IO) {
	io.String(&x.PartyId)
	io.Bool(&x.IsPartyLeader)
}

// Marshal reads or writes PlayerScoreboardId using its canonical wire layout.
func (x *PlayerScoreboardId) Marshal(io IO) {
	io.Varint64(&x.PlayerUniqueId)
}

// Marshal reads or writes PlayerUpdateEntityOverridesClearOverride using its canonical wire layout.
func (x *PlayerUpdateEntityOverridesClearOverride) Marshal(io IO) {
	io.String(&x.Type)
}

// Marshal reads or writes PlayerUpdateEntityOverridesFloatOverride using its canonical wire layout.
func (x *PlayerUpdateEntityOverridesFloatOverride) Marshal(io IO) {
	io.String(&x.Type)
	io.Float32(&x.Value)
}

// Marshal reads or writes PlayerUpdateEntityOverridesIntOverride using its canonical wire layout.
func (x *PlayerUpdateEntityOverridesIntOverride) Marshal(io IO) {
	io.String(&x.Type)
	io.Int32(&x.Value)
}

// Marshal reads or writes PlayerUpdateEntityOverridesRemoveOverride using its canonical wire layout.
func (x *PlayerUpdateEntityOverridesRemoveOverride) Marshal(io IO) {
	io.String(&x.Type)
}

func marshalPlayerUpdateEntityOverridesUpdate(io IO, x *PlayerUpdateEntityOverridesUpdate) {
	if io.Reading() {
		var tag uint8
		io.Uint8(&tag)
		switch int64(tag) {
		case 0:
			var value PlayerUpdateEntityOverridesClearOverride
			value.Marshal(io)
			*x = value
		case 1:
			var value PlayerUpdateEntityOverridesRemoveOverride
			value.Marshal(io)
			*x = value
		case 2:
			var value PlayerUpdateEntityOverridesIntOverride
			value.Marshal(io)
			*x = value
		case 3:
			var value PlayerUpdateEntityOverridesFloatOverride
			value.Marshal(io)
			*x = value
		default:
			io.InvalidValue(tag, "unknown union tag")
		}
		return
	}
	switch value := (*x).(type) {
	case PlayerUpdateEntityOverridesClearOverride:
		tag := uint8(0)
		io.Uint8(&tag)
		value.Marshal(io)
	case PlayerUpdateEntityOverridesRemoveOverride:
		tag := uint8(1)
		io.Uint8(&tag)
		value.Marshal(io)
	case PlayerUpdateEntityOverridesIntOverride:
		tag := uint8(2)
		io.Uint8(&tag)
		value.Marshal(io)
	case PlayerUpdateEntityOverridesFloatOverride:
		tag := uint8(3)
		io.Uint8(&tag)
		value.Marshal(io)
	default:
		io.InvalidValue(*x, "unknown union value")
	}
}

func marshalPlayerVideoCaptureAction(io IO, x *PlayerVideoCaptureAction) {
	if io.Reading() {
		var tag uint8
		io.Uint8(&tag)
		switch int64(tag) {
		case 0:
			var value PlayerVideoCaptureStopVideoCapture
			value.Marshal(io)
			*x = value
		case 1:
			var value PlayerVideoCaptureStartVideoCapture
			value.Marshal(io)
			*x = value
		default:
			io.InvalidValue(tag, "unknown union tag")
		}
		return
	}
	switch value := (*x).(type) {
	case PlayerVideoCaptureStopVideoCapture:
		tag := uint8(0)
		io.Uint8(&tag)
		value.Marshal(io)
	case PlayerVideoCaptureStartVideoCapture:
		tag := uint8(1)
		io.Uint8(&tag)
		value.Marshal(io)
	default:
		io.InvalidValue(*x, "unknown union value")
	}
}

// Marshal reads or writes PlayerVideoCaptureStartVideoCapture using its canonical wire layout.
func (x *PlayerVideoCaptureStartVideoCapture) Marshal(io IO) {
	io.Uint32(&x.FrameRate)
	io.String(&x.FilePrefix)
}

// Marshal reads or writes PlayerVideoCaptureStopVideoCapture using its canonical wire layout.
func (x *PlayerVideoCaptureStopVideoCapture) Marshal(io IO) {
}

// Marshal reads or writes PositionTrackingId using its canonical wire layout.
func (x *PositionTrackingId) Marshal(io IO) {
	io.Varint32(&x.Value)
}

// Marshal reads or writes PotionMixDataEntry using its canonical wire layout.
func (x *PotionMixDataEntry) Marshal(io IO) {
	io.Varint32(&x.FromPotionId)
	io.Varint32(&x.FromItemAux)
	io.Varint32(&x.ReagentItemId)
	io.Varint32(&x.ReagentItemAux)
	io.Varint32(&x.ToPotionId)
	io.Varint32(&x.ToItemAux)
}

// Marshal reads or writes PrimitiveShapeData using its canonical wire layout.
func (x *PrimitiveShapeData) Marshal(io IO) {
	io.Varuint64(&x.NetworkId)
	io.Bool(&x.ShapeType.set)
	if x.ShapeType.set {
		enumValue237 := uint8(x.ShapeType.val)
		io.Uint8(&enumValue237)
		x.ShapeType.val = ScriptModuleMinecraftScriptPrimitiveShapeType(enumValue237)
		switch int64(enumValue237) {
		case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9:
		default:
			io.InvalidValue(enumValue237, "unknown enum value")
		}
	} else if io.Reading() {
		var zero ScriptModuleMinecraftScriptPrimitiveShapeType
		x.ShapeType.val = zero
	}
	io.Bool(&x.Location.set)
	if x.Location.set {
		io.Vec3(&x.Location.val)
	} else if io.Reading() {
		var zero mgl32.Vec3
		x.Location.val = zero
	}
	io.Bool(&x.Scale.set)
	if x.Scale.set {
		io.Float32(&x.Scale.val)
	} else if io.Reading() {
		var zero float32
		x.Scale.val = zero
	}
	io.Bool(&x.Rotation.set)
	if x.Rotation.set {
		io.Vec3(&x.Rotation.val)
	} else if io.Reading() {
		var zero mgl32.Vec3
		x.Rotation.val = zero
	}
	io.Bool(&x.TotalTimeLeft.set)
	if x.TotalTimeLeft.set {
		io.Float32(&x.TotalTimeLeft.val)
	} else if io.Reading() {
		var zero float32
		x.TotalTimeLeft.val = zero
	}
	io.Bool(&x.MaximumRenderDistance.set)
	if x.MaximumRenderDistance.set {
		io.Float32(&x.MaximumRenderDistance.val)
	} else if io.Reading() {
		var zero float32
		x.MaximumRenderDistance.val = zero
	}
	io.Bool(&x.Color.set)
	if x.Color.set {
		io.RGBA(&x.Color.val)
	} else if io.Reading() {
		var zero color.RGBA
		x.Color.val = zero
	}
	io.Bool(&x.DimensionID.set)
	if x.DimensionID.set {
		x.DimensionID.val.Marshal(io)
	} else if io.Reading() {
		var zero DimensionType
		x.DimensionID.val = zero
	}
	io.Bool(&x.AttachedToEntityID.set)
	if x.AttachedToEntityID.set {
		x.AttachedToEntityID.val.Marshal(io)
	} else if io.Reading() {
		var zero ActorUniqueID
		x.AttachedToEntityID.val = zero
	}
	marshalPrimitiveShapeDataExtraShapeData(io, &x.ExtraShapeData)
}

func marshalPrimitiveShapeDataExtraShapeData(io IO, x *PrimitiveShapeDataExtraShapeData) {
	if io.Reading() {
		var tag uint32
		io.Varuint32(&tag)
		switch int64(tag) {
		case 0:
			var value PrimitiveShapeDataExtraShapeDataEmpty0
			value.Marshal(io)
			*x = value
		case 1:
			var value ArrowData
			value.Marshal(io)
			*x = value
		case 2:
			var value TextData
			value.Marshal(io)
			*x = value
		case 3:
			var value BoxData
			value.Marshal(io)
			*x = value
		case 4:
			var value LineData
			value.Marshal(io)
			*x = value
		case 5:
			var value SphereData
			value.Marshal(io)
			*x = value
		case 6:
			var value CylinderData
			value.Marshal(io)
			*x = value
		case 7:
			var value PyramidData
			value.Marshal(io)
			*x = value
		case 8:
			var value EllipsoidData
			value.Marshal(io)
			*x = value
		case 9:
			var value ConeData
			value.Marshal(io)
			*x = value
		default:
			io.InvalidValue(tag, "unknown union tag")
		}
		return
	}
	switch value := (*x).(type) {
	case PrimitiveShapeDataExtraShapeDataEmpty0:
		tag := uint32(0)
		io.Varuint32(&tag)
		value.Marshal(io)
	case ArrowData:
		tag := uint32(1)
		io.Varuint32(&tag)
		value.Marshal(io)
	case TextData:
		tag := uint32(2)
		io.Varuint32(&tag)
		value.Marshal(io)
	case BoxData:
		tag := uint32(3)
		io.Varuint32(&tag)
		value.Marshal(io)
	case LineData:
		tag := uint32(4)
		io.Varuint32(&tag)
		value.Marshal(io)
	case SphereData:
		tag := uint32(5)
		io.Varuint32(&tag)
		value.Marshal(io)
	case CylinderData:
		tag := uint32(6)
		io.Varuint32(&tag)
		value.Marshal(io)
	case PyramidData:
		tag := uint32(7)
		io.Varuint32(&tag)
		value.Marshal(io)
	case EllipsoidData:
		tag := uint32(8)
		io.Varuint32(&tag)
		value.Marshal(io)
	case ConeData:
		tag := uint32(9)
		io.Varuint32(&tag)
		value.Marshal(io)
	default:
		io.InvalidValue(*x, "unknown union value")
	}
}

// Marshal reads or writes PrimitiveShapeDataExtraShapeDataEmpty0 using its canonical wire layout.
func (x *PrimitiveShapeDataExtraShapeDataEmpty0) Marshal(io IO) {
}

// Marshal reads or writes PropertySyncData using its canonical wire layout.
func (x *PropertySyncData) Marshal(io IO) {
	if !io.Reading() && uint64(len(x.IntEntriesList)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.IntEntriesList), "collection length overflows uint32")
		return
	}
	count238 := uint32(len(x.IntEntriesList))
	io.Varuint32(&count238)
	if io.Reading() {
		if uint64(count238) > uint64(^uint(0)>>1) {
			io.InvalidValue(count238, "collection length overflows int")
			return
		}
		x.IntEntriesList = make([]PropertySyncDataPropertySyncIntEntry, int(count238))
	}
	for index239 := range x.IntEntriesList {
		x.IntEntriesList[index239].Marshal(io)
	}
	if !io.Reading() && uint64(len(x.FloatEntriesList)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.FloatEntriesList), "collection length overflows uint32")
		return
	}
	count240 := uint32(len(x.FloatEntriesList))
	io.Varuint32(&count240)
	if io.Reading() {
		if uint64(count240) > uint64(^uint(0)>>1) {
			io.InvalidValue(count240, "collection length overflows int")
			return
		}
		x.FloatEntriesList = make([]PropertySyncDataPropertySyncFloatEntry, int(count240))
	}
	for index241 := range x.FloatEntriesList {
		x.FloatEntriesList[index241].Marshal(io)
	}
}

// Marshal reads or writes PropertySyncDataPropertySyncFloatEntry using its canonical wire layout.
func (x *PropertySyncDataPropertySyncFloatEntry) Marshal(io IO) {
	io.Varuint32(&x.PropertyIndex)
	io.Float32(&x.Data)
}

// Marshal reads or writes PropertySyncDataPropertySyncIntEntry using its canonical wire layout.
func (x *PropertySyncDataPropertySyncIntEntry) Marshal(io IO) {
	io.Varuint32(&x.PropertyIndex)
	io.Varint32(&x.Data)
}

// Marshal reads or writes PyramidData using its canonical wire layout.
func (x *PyramidData) Marshal(io IO) {
	io.Float32(&x.Width)
	io.Bool(&x.Depth.set)
	if x.Depth.set {
		io.Float32(&x.Depth.val)
	} else if io.Reading() {
		var zero float32
		x.Depth.val = zero
	}
	io.Float32(&x.Height)
}

// Marshal reads or writes RemoveScore using its canonical wire layout.
func (x *RemoveScore) Marshal(io IO) {
	io.String(&x.Action)
	x.ScoreboardId.Marshal(io)
	io.Bool(&x.ObjectiveName.set)
	if x.ObjectiveName.set {
		io.String(&x.ObjectiveName.val)
	} else if io.Reading() {
		var zero string
		x.ObjectiveName.val = zero
	}
}

// Marshal reads or writes ResourcePackClientResponseCancel using its canonical wire layout.
func (x *ResourcePackClientResponseCancel) Marshal(io IO) {
	io.String(&x.ResponseType)
}

// Marshal reads or writes ResourcePackClientResponseDownloading using its canonical wire layout.
func (x *ResourcePackClientResponseDownloading) Marshal(io IO) {
	io.String(&x.ResponseType)
	if !io.Reading() && uint64(len(x.DownloadingPacks)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.DownloadingPacks), "collection length overflows uint32")
		return
	}
	count242 := uint32(len(x.DownloadingPacks))
	io.Varuint32(&count242)
	if io.Reading() {
		if uint64(count242) > uint64(^uint(0)>>1) {
			io.InvalidValue(count242, "collection length overflows int")
			return
		}
		x.DownloadingPacks = make([]string, int(count242))
	}
	for index243 := range x.DownloadingPacks {
		io.String(&x.DownloadingPacks[index243])
	}
}

// Marshal reads or writes ResourcePackClientResponseDownloadingFinished using its canonical wire layout.
func (x *ResourcePackClientResponseDownloadingFinished) Marshal(io IO) {
	io.String(&x.ResponseType)
}

// Marshal reads or writes ResourcePackClientResponseResourcePackStackFinished using its canonical wire layout.
func (x *ResourcePackClientResponseResourcePackStackFinished) Marshal(io IO) {
	io.String(&x.ResponseType)
}

func marshalResourcePackClientResponseResponse(io IO, x *ResourcePackClientResponseResponse) {
	if io.Reading() {
		var tag int8
		io.Int8(&tag)
		switch int64(tag) {
		case 1:
			var value ResourcePackClientResponseCancel
			value.Marshal(io)
			*x = value
		case 2:
			var value ResourcePackClientResponseDownloading
			value.Marshal(io)
			*x = value
		case 3:
			var value ResourcePackClientResponseDownloadingFinished
			value.Marshal(io)
			*x = value
		case 4:
			var value ResourcePackClientResponseResourcePackStackFinished
			value.Marshal(io)
			*x = value
		default:
			io.InvalidValue(tag, "unknown union tag")
		}
		return
	}
	switch value := (*x).(type) {
	case ResourcePackClientResponseCancel:
		tag := int8(1)
		io.Int8(&tag)
		value.Marshal(io)
	case ResourcePackClientResponseDownloading:
		tag := int8(2)
		io.Int8(&tag)
		value.Marshal(io)
	case ResourcePackClientResponseDownloadingFinished:
		tag := int8(3)
		io.Int8(&tag)
		value.Marshal(io)
	case ResourcePackClientResponseResourcePackStackFinished:
		tag := int8(4)
		io.Int8(&tag)
		value.Marshal(io)
	default:
		io.InvalidValue(*x, "unknown union value")
	}
}

// Marshal reads or writes ScoreboardId using its canonical wire layout.
func (x *ScoreboardId) Marshal(io IO) {
	io.Varint64(&x.ScoreboardId)
}

// Marshal reads or writes ScoreboardIdentityPacketInfo using its canonical wire layout.
func (x *ScoreboardIdentityPacketInfo) Marshal(io IO) {
	x.ScoreboardId.Marshal(io)
	io.Bool(&x.PlayerUniqueId.set)
	if x.PlayerUniqueId.set {
		io.Varint64(&x.PlayerUniqueId.val)
	} else if io.Reading() {
		var zero int64
		x.PlayerUniqueId.val = zero
	}
}

// Marshal reads or writes SemVersion using its canonical wire layout.
func (x *SemVersion) Marshal(io IO) {
	io.String(&x.Version)
}

// Marshal reads or writes SemVersionData using its canonical wire layout.
func (x *SemVersionData) Marshal(io IO) {
	io.String(&x.Version)
}

// Marshal reads or writes SerializedAbilitiesData using its canonical wire layout.
func (x *SerializedAbilitiesData) Marshal(io IO) {
	io.Int64(&x.TargetPlayerRawId)
	enumValue244 := int8(x.PlayerPermissions)
	io.Int8(&enumValue244)
	x.PlayerPermissions = PlayerPermissionLevel(enumValue244)
	switch int64(enumValue244) {
	case 0, 1, 2, 3:
	default:
		io.InvalidValue(enumValue244, "unknown enum value")
	}
	enumValue245 := uint8(x.CommandPermissions)
	io.Uint8(&enumValue245)
	x.CommandPermissions = CommandPermissionLevel(enumValue245)
	switch int64(enumValue245) {
	case 0, 1, 2, 3, 4, 5:
	default:
		io.InvalidValue(enumValue245, "unknown enum value")
	}
	if !io.Reading() && uint64(len(x.Layers)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.Layers), "collection length overflows uint32")
		return
	}
	count246 := uint32(len(x.Layers))
	io.Varuint32(&count246)
	if io.Reading() {
		if uint64(count246) > uint64(^uint(0)>>1) {
			io.InvalidValue(count246, "collection length overflows int")
			return
		}
		x.Layers = make([]SerializedAbilitiesDataSerializedLayer, int(count246))
	}
	for index247 := range x.Layers {
		x.Layers[index247].Marshal(io)
	}
}

// Marshal reads or writes SerializedAbilitiesDataSerializedLayer using its canonical wire layout.
func (x *SerializedAbilitiesDataSerializedLayer) Marshal(io IO) {
	io.Uint16(&x.SerializedLayer)
	io.Uint32(&x.AbilitiesSet)
	io.Uint32(&x.AbilityValues)
	io.Float32(&x.FlySpeed)
	io.Float32(&x.VerticalFlySpeed)
	io.Float32(&x.WalkSpeed)
}

// Marshal reads or writes SerializedNoiseBlockSpecifier using its canonical wire layout.
func (x *SerializedNoiseBlockSpecifier) Marshal(io IO) {
	io.String(&x.Noise)
	io.Float32(&x.Threshold)
	x.Range.Marshal(io)
	io.Uint32(&x.Block)
}

// Marshal reads or writes SerializedPersonaPieceHandle using its canonical wire layout.
func (x *SerializedPersonaPieceHandle) Marshal(io IO) {
	io.String(&x.PieceId)
	enumValue248 := uint32(x.PieceType)
	io.Uint32(&enumValue248)
	x.PieceType = PersonaPieceType(enumValue248)
	switch int64(enumValue248) {
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27:
	default:
		io.InvalidValue(enumValue248, "unknown enum value")
	}
	io.UUID(&x.PackId)
	io.Bool(&x.IsDefaultPiece)
	io.String(&x.ProductId)
}

// Marshal reads or writes SerializedSkinRef using its canonical wire layout.
func (x *SerializedSkinRef) Marshal(io IO) {
	io.String(&x.ID)
	io.String(&x.PlayFabID)
	io.String(&x.ResourcePatch)
	x.ImageData.Marshal(io)
	if !io.Reading() && uint64(len(x.AnimatedImageData)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.AnimatedImageData), "collection length overflows uint32")
		return
	}
	count249 := uint32(len(x.AnimatedImageData))
	io.Varuint32(&count249)
	if io.Reading() {
		if uint64(count249) > uint64(^uint(0)>>1) {
			io.InvalidValue(count249, "collection length overflows int")
			return
		}
		x.AnimatedImageData = make([]AnimatedImageData, int(count249))
	}
	for index250 := range x.AnimatedImageData {
		x.AnimatedImageData[index250].Marshal(io)
	}
	x.CapeImageData.Marshal(io)
	io.String(&x.GeometryData)
	io.String(&x.GeometryDataMinEngineVersion)
	io.String(&x.AnimationData)
	io.String(&x.CapeID)
	io.String(&x.FullID)
	enumValue251 := uint8(x.ArmSize)
	io.Uint8(&enumValue251)
	x.ArmSize = PersonaArmSizeType(enumValue251)
	switch int64(enumValue251) {
	case 0, 1:
	default:
		io.InvalidValue(enumValue251, "unknown enum value")
	}
	io.RGBA(&x.SkinColor)
	if !io.Reading() && uint64(len(x.PersonaPieces)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.PersonaPieces), "collection length overflows uint32")
		return
	}
	count252 := uint32(len(x.PersonaPieces))
	io.Varuint32(&count252)
	if io.Reading() {
		if uint64(count252) > uint64(^uint(0)>>1) {
			io.InvalidValue(count252, "collection length overflows int")
			return
		}
		x.PersonaPieces = make([]SerializedPersonaPieceHandle, int(count252))
	}
	for index253 := range x.PersonaPieces {
		x.PersonaPieces[index253].Marshal(io)
	}
	if !io.Reading() && uint64(len(x.PieceTintColors)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.PieceTintColors), "map length overflows uint32")
		return
	}
	count254 := uint32(len(x.PieceTintColors))
	io.Varuint32(&count254)
	if io.Reading() {
		if uint64(count254) > uint64(^uint(0)>>1) {
			io.InvalidValue(count254, "map length overflows int")
			return
		}
		x.PieceTintColors = make([]OrderedEntry[string, TintMapColor], int(count254))
	}
	for index255 := range x.PieceTintColors {
		io.String(&x.PieceTintColors[index255].Key)
		x.PieceTintColors[index255].Value.Marshal(io)
	}
	io.Bool(&x.IsPremium)
	io.Bool(&x.IsPersona)
	io.Bool(&x.IsPersonaCapeOnClassicSkin)
	io.Bool(&x.IsPrimaryUser)
	io.Bool(&x.OverridesPlayerAppearance)
	io.String(&x.TrustedSkinFlag)
	io.String(&x.ProfileHash)
}

// Marshal reads or writes ServerBlockProperty using its canonical wire layout.
func (x *ServerBlockProperty) Marshal(io IO) {
	io.String(&x.BlockName)
	io.NBT(&x.BlockDefinition)
}

// Marshal reads or writes ServerConfigurationClientStoreEntryPointConfiguration using its canonical wire layout.
func (x *ServerConfigurationClientStoreEntryPointConfiguration) Marshal(io IO) {
	io.String(&x.StoreId)
	io.String(&x.StoreName)
}

// Marshal reads or writes ServerConfigurationGatheringsConfigurationJoinInfo using its canonical wire layout.
func (x *ServerConfigurationGatheringsConfigurationJoinInfo) Marshal(io IO) {
	io.UUID(&x.ExperienceId)
	io.String(&x.ExperienceName)
	io.Bool(&x.WorldId.set)
	if x.WorldId.set {
		io.UUID(&x.WorldId.val)
	} else if io.Reading() {
		var zero uuid.UUID
		x.WorldId.val = zero
	}
	io.Bool(&x.WorldName.set)
	if x.WorldName.set {
		io.String(&x.WorldName.val)
	} else if io.Reading() {
		var zero string
		x.WorldName.val = zero
	}
	io.String(&x.CreatorId)
	io.Bool(&x.TargetId.set)
	if x.TargetId.set {
		io.UUID(&x.TargetId.val)
	} else if io.Reading() {
		var zero uuid.UUID
		x.TargetId.val = zero
	}
	io.Bool(&x.ScenarioId.set)
	if x.ScenarioId.set {
		io.String(&x.ScenarioId.val)
	} else if io.Reading() {
		var zero string
		x.ScenarioId.val = zero
	}
	io.Bool(&x.ServerId.set)
	if x.ServerId.set {
		io.String(&x.ServerId.val)
	} else if io.Reading() {
		var zero string
		x.ServerId.val = zero
	}
}

// Marshal reads or writes ServerConfigurationPresenceConfiguration using its canonical wire layout.
func (x *ServerConfigurationPresenceConfiguration) Marshal(io IO) {
	io.Bool(&x.RichPresenceId.set)
	if x.RichPresenceId.set {
		io.String(&x.RichPresenceId.val)
	} else if io.Reading() {
		var zero string
		x.RichPresenceId.val = zero
	}
}

// Marshal reads or writes ServerConfigurationServerConfigurationJoinInfo using its canonical wire layout.
func (x *ServerConfigurationServerConfigurationJoinInfo) Marshal(io IO) {
	io.Bool(&x.Gathering.set)
	if x.Gathering.set {
		x.Gathering.val.Marshal(io)
	} else if io.Reading() {
		var zero ServerConfigurationGatheringsConfigurationJoinInfo
		x.Gathering.val = zero
	}
	io.Bool(&x.ClientStoreEntryPoint.set)
	if x.ClientStoreEntryPoint.set {
		x.ClientStoreEntryPoint.val.Marshal(io)
	} else if io.Reading() {
		var zero ServerConfigurationClientStoreEntryPointConfiguration
		x.ClientStoreEntryPoint.val = zero
	}
	io.Bool(&x.Presence.set)
	if x.Presence.set {
		x.Presence.val.Marshal(io)
	} else if io.Reading() {
		var zero ServerConfigurationPresenceConfiguration
		x.Presence.val = zero
	}
}

// Marshal reads or writes ServerSoundHandle using its canonical wire layout.
func (x *ServerSoundHandle) Marshal(io IO) {
	io.Uint64(&x.ServerSoundHandle)
}

// Marshal reads or writes ServerWaypoint using its canonical wire layout.
func (x *ServerWaypoint) Marshal(io IO) {
	io.Uint32(&x.UpdateFlag)
	io.Bool(&x.IsVisible.set)
	if x.IsVisible.set {
		io.Bool(&x.IsVisible.val)
	} else if io.Reading() {
		var zero bool
		x.IsVisible.val = zero
	}
	io.Bool(&x.WorldPosition.set)
	if x.WorldPosition.set {
		x.WorldPosition.val.Marshal(io)
	} else if io.Reading() {
		var zero WorldPosition
		x.WorldPosition.val = zero
	}
	io.Bool(&x.TexturePath.set)
	if x.TexturePath.set {
		io.String(&x.TexturePath.val)
	} else if io.Reading() {
		var zero string
		x.TexturePath.val = zero
	}
	io.Bool(&x.IconSize.set)
	if x.IconSize.set {
		io.Vec2(&x.IconSize.val)
	} else if io.Reading() {
		var zero mgl32.Vec2
		x.IconSize.val = zero
	}
	io.Bool(&x.Color.set)
	if x.Color.set {
		io.RGBA(&x.Color.val)
	} else if io.Reading() {
		var zero color.RGBA
		x.Color.val = zero
	}
	io.Bool(&x.ClientPositionAuthority.set)
	if x.ClientPositionAuthority.set {
		io.Bool(&x.ClientPositionAuthority.val)
	} else if io.Reading() {
		var zero bool
		x.ClientPositionAuthority.val = zero
	}
	io.Bool(&x.ActorUniqueID.set)
	if x.ActorUniqueID.set {
		x.ActorUniqueID.val.Marshal(io)
	} else if io.Reading() {
		var zero ActorUniqueID
		x.ActorUniqueID.val = zero
	}
}

func marshalServerboundPackSettingChangePackSettingValue(io IO, x *ServerboundPackSettingChangePackSettingValue) {
	if io.Reading() {
		var tag uint32
		io.Varuint32(&tag)
		switch int64(tag) {
		case 0:
			var value ServerboundPackSettingChangePackSettingValueFloat
			value.Marshal(io)
			*x = value
		case 1:
			var value ServerboundPackSettingChangePackSettingValueBool
			value.Marshal(io)
			*x = value
		case 2:
			var value ServerboundPackSettingChangePackSettingValueString
			value.Marshal(io)
			*x = value
		default:
			io.InvalidValue(tag, "unknown union tag")
		}
		return
	}
	switch value := (*x).(type) {
	case ServerboundPackSettingChangePackSettingValueFloat:
		tag := uint32(0)
		io.Varuint32(&tag)
		value.Marshal(io)
	case ServerboundPackSettingChangePackSettingValueBool:
		tag := uint32(1)
		io.Varuint32(&tag)
		value.Marshal(io)
	case ServerboundPackSettingChangePackSettingValueString:
		tag := uint32(2)
		io.Varuint32(&tag)
		value.Marshal(io)
	default:
		io.InvalidValue(*x, "unknown union value")
	}
}

// Marshal reads or writes ServerboundPackSettingChangePackSettingValueBool using its canonical wire layout.
func (x *ServerboundPackSettingChangePackSettingValueBool) Marshal(io IO) {
	io.Bool(&x.Value)
}

// Marshal reads or writes ServerboundPackSettingChangePackSettingValueFloat using its canonical wire layout.
func (x *ServerboundPackSettingChangePackSettingValueFloat) Marshal(io IO) {
	io.Float32(&x.Value)
}

// Marshal reads or writes ServerboundPackSettingChangePackSettingValueString using its canonical wire layout.
func (x *ServerboundPackSettingChangePackSettingValueString) Marshal(io IO) {
	io.String(&x.Value)
}

func marshalSetScoreScoreInfoItem(io IO, x *SetScoreScoreInfoItem) {
	if io.Reading() {
		var tag uint8
		io.Uint8(&tag)
		switch int64(tag) {
		case 0:
			var value RemoveScore
			value.Marshal(io)
			*x = value
		case 1:
			var value ChangePlayerScore
			value.Marshal(io)
			*x = value
		case 2:
			var value ChangeEntityScore
			value.Marshal(io)
			*x = value
		case 3:
			var value ChangeFakePlayerScore
			value.Marshal(io)
			*x = value
		default:
			io.InvalidValue(tag, "unknown union tag")
		}
		return
	}
	switch value := (*x).(type) {
	case RemoveScore:
		tag := uint8(0)
		io.Uint8(&tag)
		value.Marshal(io)
	case ChangePlayerScore:
		tag := uint8(1)
		io.Uint8(&tag)
		value.Marshal(io)
	case ChangeEntityScore:
		tag := uint8(2)
		io.Uint8(&tag)
		value.Marshal(io)
	case ChangeFakePlayerScore:
		tag := uint8(3)
		io.Uint8(&tag)
		value.Marshal(io)
	default:
		io.InvalidValue(*x, "unknown union value")
	}
}

// Marshal reads or writes ShapedRecipe using its canonical wire layout.
func (x *ShapedRecipe) Marshal(io IO) {
	io.String(&x.RecipeId)
	io.Varint32(&x.Width)
	io.Varint32(&x.Height)
	if !io.Reading() && uint64(len(x.Ingredients)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.Ingredients), "collection length overflows uint32")
		return
	}
	count256 := uint32(len(x.Ingredients))
	io.Varuint32(&count256)
	if io.Reading() {
		if uint64(count256) > uint64(^uint(0)>>1) {
			io.InvalidValue(count256, "collection length overflows int")
			return
		}
		x.Ingredients = make([]CerealizerRecipeIngredientSerializedData, int(count256))
	}
	for index257 := range x.Ingredients {
		x.Ingredients[index257].Marshal(io)
	}
	if !io.Reading() && uint64(len(x.Results)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.Results), "collection length overflows uint32")
		return
	}
	count258 := uint32(len(x.Results))
	io.Varuint32(&count258)
	if io.Reading() {
		if uint64(count258) > uint64(^uint(0)>>1) {
			io.InvalidValue(count258, "collection length overflows int")
			return
		}
		x.Results = make([]CerealizerNetworkItemInstanceDescriptorSerializedData, int(count258))
	}
	for index259 := range x.Results {
		x.Results[index259].Marshal(io)
	}
	io.UUID(&x.UUID)
	io.String(&x.Tag)
	io.Varint32(&x.Priority)
	io.Bool(&x.AssumeSymmetry)
	io.Bool(&x.UnlockingRequirement.set)
	if x.UnlockingRequirement.set {
		x.UnlockingRequirement.val.Marshal(io)
	} else if io.Reading() {
		var zero CerealizerRecipeUnlockingRequirementSerializedData
		x.UnlockingRequirement.val = zero
	}
	x.NetId.Marshal(io)
}

// Marshal reads or writes ShapelessRecipe using its canonical wire layout.
func (x *ShapelessRecipe) Marshal(io IO) {
	io.String(&x.RecipeId)
	if !io.Reading() && uint64(len(x.Ingredients)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.Ingredients), "collection length overflows uint32")
		return
	}
	count260 := uint32(len(x.Ingredients))
	io.Varuint32(&count260)
	if io.Reading() {
		if uint64(count260) > uint64(^uint(0)>>1) {
			io.InvalidValue(count260, "collection length overflows int")
			return
		}
		x.Ingredients = make([]CerealizerRecipeIngredientSerializedData, int(count260))
	}
	for index261 := range x.Ingredients {
		x.Ingredients[index261].Marshal(io)
	}
	if !io.Reading() && uint64(len(x.Results)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.Results), "collection length overflows uint32")
		return
	}
	count262 := uint32(len(x.Results))
	io.Varuint32(&count262)
	if io.Reading() {
		if uint64(count262) > uint64(^uint(0)>>1) {
			io.InvalidValue(count262, "collection length overflows int")
			return
		}
		x.Results = make([]CerealizerNetworkItemInstanceDescriptorSerializedData, int(count262))
	}
	for index263 := range x.Results {
		x.Results[index263].Marshal(io)
	}
	io.UUID(&x.UUID)
	io.String(&x.Tag)
	io.Varint32(&x.Priority)
	io.Bool(&x.UnlockingRequirement.set)
	if x.UnlockingRequirement.set {
		x.UnlockingRequirement.val.Marshal(io)
	} else if io.Reading() {
		var zero CerealizerRecipeUnlockingRequirementSerializedData
		x.UnlockingRequirement.val = zero
	}
	x.NetId.Marshal(io)
}

// Marshal reads or writes SkinImage using its canonical wire layout.
func (x *SkinImage) Marshal(io IO) {
	io.Uint32(&x.Width)
	io.Uint32(&x.Height)
	if !io.Reading() && uint64(len(x.ImageBytes)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.ImageBytes), "collection length overflows uint32")
		return
	}
	count264 := uint32(len(x.ImageBytes))
	io.Varuint32(&count264)
	if io.Reading() {
		if uint64(count264) > uint64(^uint(0)>>1) {
			io.InvalidValue(count264, "collection length overflows int")
			return
		}
		x.ImageBytes = make([]uint8, int(count264))
	}
	for index265 := range x.ImageBytes {
		io.Uint8(&x.ImageBytes[index265])
	}
}

// Marshal reads or writes SmithingTransformRecipe using its canonical wire layout.
func (x *SmithingTransformRecipe) Marshal(io IO) {
	io.String(&x.RecipeId)
	x.TemplateIngredient.Marshal(io)
	x.BaseIngredient.Marshal(io)
	x.AdditionIngredient.Marshal(io)
	x.Result.Marshal(io)
	io.String(&x.Tag)
	x.NetId.Marshal(io)
}

// Marshal reads or writes SmithingTrimRecipe using its canonical wire layout.
func (x *SmithingTrimRecipe) Marshal(io IO) {
	io.String(&x.RecipeId)
	x.TemplateIngredient.Marshal(io)
	x.BaseIngredient.Marshal(io)
	x.AdditionIngredient.Marshal(io)
	io.String(&x.Tag)
	x.NetId.Marshal(io)
}

// Marshal reads or writes SocialEventsServerTelemetryData using its canonical wire layout.
func (x *SocialEventsServerTelemetryData) Marshal(io IO) {
	io.String(&x.ServerId)
	io.String(&x.ScenarioId)
	io.String(&x.WorldId)
	io.String(&x.OwnerId)
}

func marshalSoundDataEvent(io IO, x *SoundDataEvent) {
	if io.Reading() {
		var tag uint32
		io.Varuint32(&tag)
		switch int64(tag) {
		case 0:
			var value SoundDataEventStop
			value.Marshal(io)
			*x = value
		case 1:
			var value SoundDataEventSetVolume
			value.Marshal(io)
			*x = value
		case 2:
			var value SoundDataEventSetPitch
			value.Marshal(io)
			*x = value
		case 3:
			var value SoundDataEventFade
			value.Marshal(io)
			*x = value
		case 4:
			var value SoundDataEventSeekTo
			value.Marshal(io)
			*x = value
		case 5:
			var value SoundDataEventPause
			value.Marshal(io)
			*x = value
		case 6:
			var value SoundDataEventResume
			value.Marshal(io)
			*x = value
		default:
			io.InvalidValue(tag, "unknown union tag")
		}
		return
	}
	switch value := (*x).(type) {
	case SoundDataEventStop:
		tag := uint32(0)
		io.Varuint32(&tag)
		value.Marshal(io)
	case SoundDataEventSetVolume:
		tag := uint32(1)
		io.Varuint32(&tag)
		value.Marshal(io)
	case SoundDataEventSetPitch:
		tag := uint32(2)
		io.Varuint32(&tag)
		value.Marshal(io)
	case SoundDataEventFade:
		tag := uint32(3)
		io.Varuint32(&tag)
		value.Marshal(io)
	case SoundDataEventSeekTo:
		tag := uint32(4)
		io.Varuint32(&tag)
		value.Marshal(io)
	case SoundDataEventPause:
		tag := uint32(5)
		io.Varuint32(&tag)
		value.Marshal(io)
	case SoundDataEventResume:
		tag := uint32(6)
		io.Varuint32(&tag)
		value.Marshal(io)
	default:
		io.InvalidValue(*x, "unknown union value")
	}
}

// Marshal reads or writes SoundDataEventFade using its canonical wire layout.
func (x *SoundDataEventFade) Marshal(io IO) {
	io.Float32(&x.Duration)
	io.Float32(&x.TargetVolume)
}

// Marshal reads or writes SoundDataEventPause using its canonical wire layout.
func (x *SoundDataEventPause) Marshal(io IO) {
}

// Marshal reads or writes SoundDataEventResume using its canonical wire layout.
func (x *SoundDataEventResume) Marshal(io IO) {
}

// Marshal reads or writes SoundDataEventSeekTo using its canonical wire layout.
func (x *SoundDataEventSeekTo) Marshal(io IO) {
	io.Float32(&x.Seconds)
}

// Marshal reads or writes SoundDataEventSetPitch using its canonical wire layout.
func (x *SoundDataEventSetPitch) Marshal(io IO) {
	io.Float32(&x.Pitch)
}

// Marshal reads or writes SoundDataEventSetVolume using its canonical wire layout.
func (x *SoundDataEventSetVolume) Marshal(io IO) {
	io.Float32(&x.Volume)
}

// Marshal reads or writes SoundDataEventStop using its canonical wire layout.
func (x *SoundDataEventStop) Marshal(io IO) {
}

// Marshal reads or writes SpawnSettings using its canonical wire layout.
func (x *SpawnSettings) Marshal(io IO) {
	enumValue266 := int16(x.SpawnBiomeType)
	io.Int16(&enumValue266)
	x.SpawnBiomeType = SpawnBiomeType(enumValue266)
	switch int64(enumValue266) {
	case 0, 1:
	default:
		io.InvalidValue(enumValue266, "unknown enum value")
	}
	io.String(&x.UserDefinedBiomeName)
	io.Varint32(&x.Dimension)
}

// Marshal reads or writes SphereData using its canonical wire layout.
func (x *SphereData) Marshal(io IO) {
	io.Uint8(&x.NumSegments)
}

// Marshal reads or writes StructureEditorData using its canonical wire layout.
func (x *StructureEditorData) Marshal(io IO) {
	x.StructureName.Marshal(io)
	io.String(&x.DataField)
	io.Bool(&x.ShouldIncludePlayers)
	io.Bool(&x.ShouldShowBoundingBox)
	enumValue267 := int32(x.StructureBlockType)
	io.Varint32(&enumValue267)
	x.StructureBlockType = StructureBlockType(enumValue267)
	switch int64(enumValue267) {
	case 0, 1, 2, 3, 4, 5:
	default:
		io.InvalidValue(enumValue267, "unknown enum value")
	}
	x.StructureSettings.Marshal(io)
	enumValue268 := uint8(x.RedstoneSaveMode)
	io.Uint8(&enumValue268)
	x.RedstoneSaveMode = StructureRedstoneSaveMode(enumValue268)
	switch int64(enumValue268) {
	case 0, 1:
	default:
		io.InvalidValue(enumValue268, "unknown enum value")
	}
}

// Marshal reads or writes StructureSettings using its canonical wire layout.
func (x *StructureSettings) Marshal(io IO) {
	io.String(&x.StructurePaletteName)
	io.Bool(&x.ShouldIgnoreEntities)
	io.Bool(&x.ShouldIgnoreBlocks)
	io.Bool(&x.ShouldAllowNonTickingPlayerAndTickingAreaChunks)
	x.StructureSize.Marshal(io)
	x.StructureOffset.Marshal(io)
	x.LastEditPlayer.Marshal(io)
	enumValue269 := uint8(x.Rotation)
	io.Uint8(&enumValue269)
	x.Rotation = Rotation(enumValue269)
	switch int64(enumValue269) {
	case 0, 1, 2, 3:
	default:
		io.InvalidValue(enumValue269, "unknown enum value")
	}
	enumValue270 := uint8(x.Mirror)
	io.Uint8(&enumValue270)
	x.Mirror = Mirror(enumValue270)
	switch int64(enumValue270) {
	case 0, 1, 2, 3:
	default:
		io.InvalidValue(enumValue270, "unknown enum value")
	}
	enumValue271 := uint8(x.AnimationMode)
	io.Uint8(&enumValue271)
	x.AnimationMode = AnimationMode(enumValue271)
	switch int64(enumValue271) {
	case 0, 1, 2:
	default:
		io.InvalidValue(enumValue271, "unknown enum value")
	}
	io.Float32(&x.AnimationSeconds)
	io.Float32(&x.IntegrityValue)
	io.Uint32(&x.IntegritySeed)
	io.Vec3(&x.RotationPivot)
}

// Marshal reads or writes SubChunkHeightmapData using its canonical wire layout.
func (x *SubChunkHeightmapData) Marshal(io IO) {
	enumValue272 := uint8(x.HeightMapType)
	io.Uint8(&enumValue272)
	x.HeightMapType = SubChunkHeightMapDataType(enumValue272)
	switch int64(enumValue272) {
	case 0, 1, 2, 3:
	default:
		io.InvalidValue(enumValue272, "unknown enum value")
	}
	io.Bool(&x.SubchunkHeightMap.set)
	if x.SubchunkHeightMap.set {
		for index273 := range x.SubchunkHeightMap.val {
			for index274 := range x.SubchunkHeightMap.val[index273] {
				io.Int8(&x.SubchunkHeightMap.val[index273][index274])
			}
		}
	} else if io.Reading() {
		var zero [16][16]int8
		x.SubchunkHeightMap.val = zero
	}
	enumValue275 := uint8(x.RenderHeightMapType)
	io.Uint8(&enumValue275)
	x.RenderHeightMapType = SubChunkHeightMapDataType(enumValue275)
	switch int64(enumValue275) {
	case 0, 1, 2, 3, 4:
	default:
		io.InvalidValue(enumValue275, "unknown enum value")
	}
	io.Bool(&x.SubchunkRenderHeightMap.set)
	if x.SubchunkRenderHeightMap.set {
		for index276 := range x.SubchunkRenderHeightMap.val {
			for index277 := range x.SubchunkRenderHeightMap.val[index276] {
				io.Int8(&x.SubchunkRenderHeightMap.val[index276][index277])
			}
		}
	} else if io.Reading() {
		var zero [16][16]int8
		x.SubchunkRenderHeightMap.val = zero
	}
}

// Marshal reads or writes SubChunkPos using its canonical wire layout.
func (x *SubChunkPos) Marshal(io IO) {
	io.Int32(&x.SubchunkPositionX)
	io.Int32(&x.SubchunkPositionY)
	io.Int32(&x.SubchunkPositionZ)
}

// Marshal reads or writes SubChunkSubChunkPacketData using its canonical wire layout.
func (x *SubChunkSubChunkPacketData) Marshal(io IO) {
	x.SubChunkPosOffset.Marshal(io)
	enumValue278 := uint8(x.SubChunkRequestResult)
	io.Uint8(&enumValue278)
	x.SubChunkRequestResult = SubChunkSubChunkRequestResult(enumValue278)
	switch int64(enumValue278) {
	case 1, 2, 3, 4, 5, 6:
	default:
		io.InvalidValue(enumValue278, "unknown enum value")
	}
	io.Bool(&x.SerializedSubChunk.set)
	if x.SerializedSubChunk.set {
		io.String(&x.SerializedSubChunk.val)
	} else if io.Reading() {
		var zero string
		x.SerializedSubChunk.val = zero
	}
	x.HeightMapData.Marshal(io)
	io.Bool(&x.BlobId.set)
	if x.BlobId.set {
		io.Uint64(&x.BlobId.val)
	} else if io.Reading() {
		var zero uint64
		x.BlobId.val = zero
	}
}

// Marshal reads or writes SubChunkSubChunkPosOffset using its canonical wire layout.
func (x *SubChunkSubChunkPosOffset) Marshal(io IO) {
	io.Int8(&x.SubchunkOffsetX)
	io.Int8(&x.SubchunkOffsetY)
	io.Int8(&x.SubchunkOffsetZ)
}

// Marshal reads or writes SyncWorldClockStateData using its canonical wire layout.
func (x *SyncWorldClockStateData) Marshal(io IO) {
	io.Varuint64(&x.ClockId)
	io.Varint32(&x.Time)
	io.Bool(&x.IsPaused)
}

// Marshal reads or writes SyncWorldClocksAddTimeMarkerData using its canonical wire layout.
func (x *SyncWorldClocksAddTimeMarkerData) Marshal(io IO) {
	io.Varuint64(&x.ClockId)
	if !io.Reading() && uint64(len(x.TimeMarkers)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.TimeMarkers), "collection length overflows uint32")
		return
	}
	count279 := uint32(len(x.TimeMarkers))
	io.Varuint32(&count279)
	if io.Reading() {
		if uint64(count279) > uint64(^uint(0)>>1) {
			io.InvalidValue(count279, "collection length overflows int")
			return
		}
		x.TimeMarkers = make([]TimeMarkerData, int(count279))
	}
	for index280 := range x.TimeMarkers {
		x.TimeMarkers[index280].Marshal(io)
	}
}

func marshalSyncWorldClocksData(io IO, x *SyncWorldClocksData) {
	if io.Reading() {
		var tag uint32
		io.Varuint32(&tag)
		switch int64(tag) {
		case 0:
			var value SyncWorldClocksSyncStateData
			value.Marshal(io)
			*x = value
		case 1:
			var value SyncWorldClocksInitializeRegistryData
			value.Marshal(io)
			*x = value
		case 2:
			var value SyncWorldClocksAddTimeMarkerData
			value.Marshal(io)
			*x = value
		case 3:
			var value SyncWorldClocksRemoveTimeMarkerData
			value.Marshal(io)
			*x = value
		default:
			io.InvalidValue(tag, "unknown union tag")
		}
		return
	}
	switch value := (*x).(type) {
	case SyncWorldClocksSyncStateData:
		tag := uint32(0)
		io.Varuint32(&tag)
		value.Marshal(io)
	case SyncWorldClocksInitializeRegistryData:
		tag := uint32(1)
		io.Varuint32(&tag)
		value.Marshal(io)
	case SyncWorldClocksAddTimeMarkerData:
		tag := uint32(2)
		io.Varuint32(&tag)
		value.Marshal(io)
	case SyncWorldClocksRemoveTimeMarkerData:
		tag := uint32(3)
		io.Varuint32(&tag)
		value.Marshal(io)
	default:
		io.InvalidValue(*x, "unknown union value")
	}
}

// Marshal reads or writes SyncWorldClocksInitializeRegistryData using its canonical wire layout.
func (x *SyncWorldClocksInitializeRegistryData) Marshal(io IO) {
	if !io.Reading() && uint64(len(x.ClockData)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.ClockData), "collection length overflows uint32")
		return
	}
	count281 := uint32(len(x.ClockData))
	io.Varuint32(&count281)
	if io.Reading() {
		if uint64(count281) > uint64(^uint(0)>>1) {
			io.InvalidValue(count281, "collection length overflows int")
			return
		}
		x.ClockData = make([]WorldClockData, int(count281))
	}
	for index282 := range x.ClockData {
		x.ClockData[index282].Marshal(io)
	}
}

// Marshal reads or writes SyncWorldClocksRemoveTimeMarkerData using its canonical wire layout.
func (x *SyncWorldClocksRemoveTimeMarkerData) Marshal(io IO) {
	io.Varuint64(&x.ClockId)
	if !io.Reading() && uint64(len(x.TimeMarkerIds)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.TimeMarkerIds), "collection length overflows uint32")
		return
	}
	count283 := uint32(len(x.TimeMarkerIds))
	io.Varuint32(&count283)
	if io.Reading() {
		if uint64(count283) > uint64(^uint(0)>>1) {
			io.InvalidValue(count283, "collection length overflows int")
			return
		}
		x.TimeMarkerIds = make([]uint64, int(count283))
	}
	for index284 := range x.TimeMarkerIds {
		io.Varuint64(&x.TimeMarkerIds[index284])
	}
}

// Marshal reads or writes SyncWorldClocksSyncStateData using its canonical wire layout.
func (x *SyncWorldClocksSyncStateData) Marshal(io IO) {
	if !io.Reading() && uint64(len(x.ClockData)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.ClockData), "collection length overflows uint32")
		return
	}
	count285 := uint32(len(x.ClockData))
	io.Varuint32(&count285)
	if io.Reading() {
		if uint64(count285) > uint64(^uint(0)>>1) {
			io.InvalidValue(count285, "collection length overflows int")
			return
		}
		x.ClockData = make([]SyncWorldClockStateData, int(count285))
	}
	for index286 := range x.ClockData {
		x.ClockData[index286].Marshal(io)
	}
}

// Marshal reads or writes SyncedAttribute using its canonical wire layout.
func (x *SyncedAttribute) Marshal(io IO) {
	io.String(&x.AttributeName)
	io.Float32(&x.MinValue)
	io.Float32(&x.CurrentValue)
	io.Float32(&x.MaxValue)
}

// Marshal reads or writes SyncedPlayerMovementSettings using its canonical wire layout.
func (x *SyncedPlayerMovementSettings) Marshal(io IO) {
	io.Varint32(&x.RewindHistorySize)
	io.Bool(&x.ServerAuthoritativeBlockBreaking)
}

// Marshal reads or writes SynchedActorDataCopyableDataList using its canonical wire layout.
func (x *SynchedActorDataCopyableDataList) Marshal(io IO) {
	if !io.Reading() && uint64(len(x.Data)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.Data), "collection length overflows uint32")
		return
	}
	count287 := uint32(len(x.Data))
	io.Varuint32(&count287)
	if io.Reading() {
		if uint64(count287) > uint64(^uint(0)>>1) {
			io.InvalidValue(count287, "collection length overflows int")
			return
		}
		x.Data = make([]DataItemEntry, int(count287))
	}
	for index288 := range x.Data {
		x.Data[index288].Marshal(io)
	}
}

// Marshal reads or writes TextAuthorAndMessage using its canonical wire layout.
func (x *TextAuthorAndMessage) Marshal(io IO) {
	io.String(&x.PlayerName)
	io.String(&x.Message)
}

func marshalTextBody(io IO, x *TextBody) {
	if io.Reading() {
		var tag uint8
		io.Uint8(&tag)
		switch int64(tag) {
		case 0:
			var value TextMessageOnly
			value.Marshal(io)
			*x = value
		case 1:
			var value TextAuthorAndMessage
			value.Marshal(io)
			*x = value
		case 2:
			var value TextMessageAndParams
			value.Marshal(io)
			*x = value
		case 3:
			var value TextBodyPopup
			value.Marshal(io)
			*x = value
		case 4:
			var value TextBodyJukeboxPopup
			value.Marshal(io)
			*x = value
		case 5:
			var value TextBodyTip
			value.Marshal(io)
			*x = value
		case 6:
			var value TextBodySystemMessage
			value.Marshal(io)
			*x = value
		case 7:
			var value TextBodyWhisper
			value.Marshal(io)
			*x = value
		case 8:
			var value TextBodyAnnouncement
			value.Marshal(io)
			*x = value
		case 9:
			var value TextBodyTextObjectWhisper
			value.Marshal(io)
			*x = value
		case 10:
			var value TextBodyTextObject
			value.Marshal(io)
			*x = value
		case 11:
			var value TextBodyTextObjectAnnouncement
			value.Marshal(io)
			*x = value
		default:
			io.InvalidValue(tag, "unknown union tag")
		}
		return
	}
	switch value := (*x).(type) {
	case TextMessageOnly:
		tag := uint8(0)
		io.Uint8(&tag)
		value.Marshal(io)
	case TextAuthorAndMessage:
		tag := uint8(1)
		io.Uint8(&tag)
		value.Marshal(io)
	case TextMessageAndParams:
		tag := uint8(2)
		io.Uint8(&tag)
		value.Marshal(io)
	case TextBodyPopup:
		tag := uint8(3)
		io.Uint8(&tag)
		value.Marshal(io)
	case TextBodyJukeboxPopup:
		tag := uint8(4)
		io.Uint8(&tag)
		value.Marshal(io)
	case TextBodyTip:
		tag := uint8(5)
		io.Uint8(&tag)
		value.Marshal(io)
	case TextBodySystemMessage:
		tag := uint8(6)
		io.Uint8(&tag)
		value.Marshal(io)
	case TextBodyWhisper:
		tag := uint8(7)
		io.Uint8(&tag)
		value.Marshal(io)
	case TextBodyAnnouncement:
		tag := uint8(8)
		io.Uint8(&tag)
		value.Marshal(io)
	case TextBodyTextObjectWhisper:
		tag := uint8(9)
		io.Uint8(&tag)
		value.Marshal(io)
	case TextBodyTextObject:
		tag := uint8(10)
		io.Uint8(&tag)
		value.Marshal(io)
	case TextBodyTextObjectAnnouncement:
		tag := uint8(11)
		io.Uint8(&tag)
		value.Marshal(io)
	default:
		io.InvalidValue(*x, "unknown union value")
	}
}

// Marshal reads or writes TextBodyAnnouncement using its canonical wire layout.
func (x *TextBodyAnnouncement) Marshal(io IO) {
	x.Value.Marshal(io)
}

// Marshal reads or writes TextBodyJukeboxPopup using its canonical wire layout.
func (x *TextBodyJukeboxPopup) Marshal(io IO) {
	x.Value.Marshal(io)
}

// Marshal reads or writes TextBodyPopup using its canonical wire layout.
func (x *TextBodyPopup) Marshal(io IO) {
	x.Value.Marshal(io)
}

// Marshal reads or writes TextBodySystemMessage using its canonical wire layout.
func (x *TextBodySystemMessage) Marshal(io IO) {
	x.Value.Marshal(io)
}

// Marshal reads or writes TextBodyTextObject using its canonical wire layout.
func (x *TextBodyTextObject) Marshal(io IO) {
	x.Value.Marshal(io)
}

// Marshal reads or writes TextBodyTextObjectAnnouncement using its canonical wire layout.
func (x *TextBodyTextObjectAnnouncement) Marshal(io IO) {
	x.Value.Marshal(io)
}

// Marshal reads or writes TextBodyTextObjectWhisper using its canonical wire layout.
func (x *TextBodyTextObjectWhisper) Marshal(io IO) {
	x.Value.Marshal(io)
}

// Marshal reads or writes TextBodyTip using its canonical wire layout.
func (x *TextBodyTip) Marshal(io IO) {
	x.Value.Marshal(io)
}

// Marshal reads or writes TextBodyWhisper using its canonical wire layout.
func (x *TextBodyWhisper) Marshal(io IO) {
	x.Value.Marshal(io)
}

// Marshal reads or writes TextData using its canonical wire layout.
func (x *TextData) Marshal(io IO) {
	io.String(&x.Text)
	io.Bool(&x.UseRotation)
	io.Bool(&x.BackgroundColor.set)
	if x.BackgroundColor.set {
		io.RGBA(&x.BackgroundColor.val)
	} else if io.Reading() {
		var zero color.RGBA
		x.BackgroundColor.val = zero
	}
	io.Bool(&x.DepthTest)
	io.Bool(&x.ShowBackface)
	io.Bool(&x.ShowTextBackface)
}

// Marshal reads or writes TextMessageAndParams using its canonical wire layout.
func (x *TextMessageAndParams) Marshal(io IO) {
	io.String(&x.Message)
	if !io.Reading() && uint64(len(x.ParameterList)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.ParameterList), "collection length overflows uint32")
		return
	}
	count289 := uint32(len(x.ParameterList))
	io.Varuint32(&count289)
	if io.Reading() {
		if uint64(count289) > uint64(^uint(0)>>1) {
			io.InvalidValue(count289, "collection length overflows int")
			return
		}
		x.ParameterList = make([]string, int(count289))
	}
	for index290 := range x.ParameterList {
		io.String(&x.ParameterList[index290])
	}
}

// Marshal reads or writes TextMessageOnly using its canonical wire layout.
func (x *TextMessageOnly) Marshal(io IO) {
	io.String(&x.Message)
}

// Marshal reads or writes TimeMarkerData using its canonical wire layout.
func (x *TimeMarkerData) Marshal(io IO) {
	io.Varuint64(&x.Id)
	io.String(&x.Name)
	io.Varint32(&x.Time)
	io.Bool(&x.Period.set)
	if x.Period.set {
		io.Int32(&x.Period.val)
	} else if io.Reading() {
		var zero int32
		x.Period.val = zero
	}
}

// Marshal reads or writes TintMapColor using its canonical wire layout.
func (x *TintMapColor) Marshal(io IO) {
	for index291 := range x.Colors {
		io.RGBA(&x.Colors[index291])
	}
}

// Marshal reads or writes TrimMaterial using its canonical wire layout.
func (x *TrimMaterial) Marshal(io IO) {
	io.String(&x.MaterialId)
	io.String(&x.Color)
	io.String(&x.ItemName)
}

// Marshal reads or writes TrimPattern using its canonical wire layout.
func (x *TrimPattern) Marshal(io IO) {
	io.String(&x.ItemName)
	io.String(&x.PatternId)
}

// Marshal reads or writes TypedClientNetIdStructItemStackLegacyRequestIdTagInt32T0 using its canonical wire layout.
func (x *TypedClientNetIdStructItemStackLegacyRequestIdTagInt32T0) Marshal(io IO) {
	io.Varint32(&x.ID)
}

// Marshal reads or writes TypedClientNetIdStructItemStackRequestIdTagInt32T0 using its canonical wire layout.
func (x *TypedClientNetIdStructItemStackRequestIdTagInt32T0) Marshal(io IO) {
	io.Varint32(&x.ID)
}

// Marshal reads or writes TypedServerNetIdStructCreativeItemNetIdTag using its canonical wire layout.
func (x *TypedServerNetIdStructCreativeItemNetIdTag) Marshal(io IO) {
	io.Varuint32(&x.ID)
}

// Marshal reads or writes TypedServerNetIdStructItemStackNetIdTagInt32T0 using its canonical wire layout.
func (x *TypedServerNetIdStructItemStackNetIdTagInt32T0) Marshal(io IO) {
	io.Varint32(&x.ID)
}

// Marshal reads or writes TypedServerNetIdStructRecipeNetIdTag using its canonical wire layout.
func (x *TypedServerNetIdStructRecipeNetIdTag) Marshal(io IO) {
	io.Varuint32(&x.RawId)
}

// Marshal reads or writes UpdateSubChunkBlocksChangedInfo using its canonical wire layout.
func (x *UpdateSubChunkBlocksChangedInfo) Marshal(io IO) {
	if !io.Reading() && uint64(len(x.BlocksChangedStandards)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.BlocksChangedStandards), "collection length overflows uint32")
		return
	}
	count292 := uint32(len(x.BlocksChangedStandards))
	io.Varuint32(&count292)
	if io.Reading() {
		if uint64(count292) > uint64(^uint(0)>>1) {
			io.InvalidValue(count292, "collection length overflows int")
			return
		}
		x.BlocksChangedStandards = make([]UpdateSubChunkNetworkBlockInfo, int(count292))
	}
	for index293 := range x.BlocksChangedStandards {
		x.BlocksChangedStandards[index293].Marshal(io)
	}
	if !io.Reading() && uint64(len(x.BlocksChangedExtras)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.BlocksChangedExtras), "collection length overflows uint32")
		return
	}
	count294 := uint32(len(x.BlocksChangedExtras))
	io.Varuint32(&count294)
	if io.Reading() {
		if uint64(count294) > uint64(^uint(0)>>1) {
			io.InvalidValue(count294, "collection length overflows int")
			return
		}
		x.BlocksChangedExtras = make([]UpdateSubChunkNetworkBlockInfo, int(count294))
	}
	for index295 := range x.BlocksChangedExtras {
		x.BlocksChangedExtras[index295].Marshal(io)
	}
}

// Marshal reads or writes UpdateSubChunkNetworkBlockInfo using its canonical wire layout.
func (x *UpdateSubChunkNetworkBlockInfo) Marshal(io IO) {
	x.Pos.Marshal(io)
	io.Varuint32(&x.RuntimeId)
	io.Varuint32(&x.UpdateFlags)
	io.Varuint64(&x.SyncMessageEntityUniqueID)
	io.Varuint32(&x.SyncMessageMessage)
}

// Marshal reads or writes VoxelShapesRegistryHandle using its canonical wire layout.
func (x *VoxelShapesRegistryHandle) Marshal(io IO) {
	io.Uint16(&x.Value)
}

// Marshal reads or writes VoxelShapesSerializableCells using its canonical wire layout.
func (x *VoxelShapesSerializableCells) Marshal(io IO) {
	io.Uint8(&x.XSize)
	io.Uint8(&x.YSize)
	io.Uint8(&x.ZSize)
	if !io.Reading() && uint64(len(x.Storage)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.Storage), "collection length overflows uint32")
		return
	}
	count296 := uint32(len(x.Storage))
	io.Varuint32(&count296)
	if io.Reading() {
		if uint64(count296) > uint64(^uint(0)>>1) {
			io.InvalidValue(count296, "collection length overflows int")
			return
		}
		x.Storage = make([]uint8, int(count296))
	}
	for index297 := range x.Storage {
		io.Uint8(&x.Storage[index297])
	}
}

// Marshal reads or writes VoxelShapesSerializableVoxelShape using its canonical wire layout.
func (x *VoxelShapesSerializableVoxelShape) Marshal(io IO) {
	x.Cells.Marshal(io)
	if !io.Reading() && uint64(len(x.XCoordinates)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.XCoordinates), "collection length overflows uint32")
		return
	}
	count298 := uint32(len(x.XCoordinates))
	io.Varuint32(&count298)
	if io.Reading() {
		if uint64(count298) > uint64(^uint(0)>>1) {
			io.InvalidValue(count298, "collection length overflows int")
			return
		}
		x.XCoordinates = make([]float32, int(count298))
	}
	for index299 := range x.XCoordinates {
		io.Float32(&x.XCoordinates[index299])
	}
	if !io.Reading() && uint64(len(x.YCoordinates)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.YCoordinates), "collection length overflows uint32")
		return
	}
	count300 := uint32(len(x.YCoordinates))
	io.Varuint32(&count300)
	if io.Reading() {
		if uint64(count300) > uint64(^uint(0)>>1) {
			io.InvalidValue(count300, "collection length overflows int")
			return
		}
		x.YCoordinates = make([]float32, int(count300))
	}
	for index301 := range x.YCoordinates {
		io.Float32(&x.YCoordinates[index301])
	}
	if !io.Reading() && uint64(len(x.ZCoordinates)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.ZCoordinates), "collection length overflows uint32")
		return
	}
	count302 := uint32(len(x.ZCoordinates))
	io.Varuint32(&count302)
	if io.Reading() {
		if uint64(count302) > uint64(^uint(0)>>1) {
			io.InvalidValue(count302, "collection length overflows int")
			return
		}
		x.ZCoordinates = make([]float32, int(count302))
	}
	for index303 := range x.ZCoordinates {
		io.Float32(&x.ZCoordinates[index303])
	}
}

// Marshal reads or writes WaypointGroupWaypointHandle using its canonical wire layout.
func (x *WaypointGroupWaypointHandle) Marshal(io IO) {
	io.UUID(&x.UUID)
}

// Marshal reads or writes WebSocketPacketData using its canonical wire layout.
func (x *WebSocketPacketData) Marshal(io IO) {
	io.String(&x.WebsocketServerURI)
}

// Marshal reads or writes WorldClockData using its canonical wire layout.
func (x *WorldClockData) Marshal(io IO) {
	io.Varuint64(&x.Id)
	io.String(&x.Name)
	io.Varint32(&x.Time)
	io.Bool(&x.IsPaused)
	if !io.Reading() && uint64(len(x.TimeMarkers)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.TimeMarkers), "collection length overflows uint32")
		return
	}
	count304 := uint32(len(x.TimeMarkers))
	io.Varuint32(&count304)
	if io.Reading() {
		if uint64(count304) > uint64(^uint(0)>>1) {
			io.InvalidValue(count304, "collection length overflows int")
			return
		}
		x.TimeMarkers = make([]TimeMarkerData, int(count304))
	}
	for index305 := range x.TimeMarkers {
		x.TimeMarkers[index305].Marshal(io)
	}
}

// Marshal reads or writes WorldPosition using its canonical wire layout.
func (x *WorldPosition) Marshal(io IO) {
	io.Vec3(&x.Position)
	x.DimensionType.Marshal(io)
}
