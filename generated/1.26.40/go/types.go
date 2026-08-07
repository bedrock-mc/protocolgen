// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import (
	"image/color"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/google/uuid"
)

// OrderedEntry preserves the source order and duplicate keys of a wire map.
type OrderedEntry[K, V any] struct {
	Key   K
	Value V
}

type ActorDataBoundingBoxComponent struct {
	ActorDataBoundingBox [3]float32
}

type ActorDataFlagComponent struct {
	ActorFlagBitsetData []byte
}

type ActorLink struct {
	TargetA                ActorUniqueID
	TargetB                ActorUniqueID
	Type                   ActorLinkType
	Immediate              bool
	PassengerInitiated     bool
	VehicleAngularVelocity float32
}

type ActorRuntimeID struct {
	ActorRuntimeID uint64
}

type ActorUniqueID struct {
	ActorUniqueID int64
}

type AdventureSettings struct {
	NoPvM          bool
	NoMvP          bool
	ImmutableWorld bool
	ShowNameTags   bool
	AutoJump       bool
}

type AgentCapabilities struct {
	CanModifyBlocks *bool
}

type AnimatedImageData struct {
	SkinImage           SkinImage
	AnimatedTextureType PersonaAnimatedTextureType
	Frames              float32
	AnimationExpression PersonaAnimationExpression
}

type ArmorSlotAndDamagePair struct {
	ArmorSlot LegacyArmorSlot
	Damage    int16
}

type ArrowData struct {
	ArrowEndLocation *mgl32.Vec3
	ArrowHeadLength  *float32
	ArrowHeadRadius  *float32
	NumSegments      *uint8
}

func (ArrowData) isPrimitiveShapeDataExtraShapeData() {}

type AttributeData struct {
	MinValue        float32
	MaxValue        float32
	CurrentValue    float32
	DefaultMinValue float32
	DefaultMaxValue float32
	DefaultValue    float32
	Name            string
	Modifiers       []AttributeModifier
}

type AttributeLayerSyncPacketData interface {
	isAttributeLayerSyncPacketData()
}

type AttributeLayerSyncPacketDataRemoveEnvironmentAttributesData struct {
	AttributeLayerName      string
	AttributeLayerDimension DimensionType
	Attributes              []string
}

func (AttributeLayerSyncPacketDataRemoveEnvironmentAttributesData) isAttributeLayerSyncPacketData() {}

type AttributeLayerSyncPacketDataUpdateAttributeLayerSettingsData struct {
	AttributeLayerName      string
	AttributeLayerDimension DimensionType
	AttributesLayerSettings EASAttributeLayerSettings
}

func (AttributeLayerSyncPacketDataUpdateAttributeLayerSettingsData) isAttributeLayerSyncPacketData() {
}

type AttributeLayerSyncPacketDataUpdateAttributeLayersData struct {
	AttributeLayers []EASAttributeLayerData
}

func (AttributeLayerSyncPacketDataUpdateAttributeLayersData) isAttributeLayerSyncPacketData() {}

type AttributeLayerSyncPacketDataUpdateEnvironmentAttributesData struct {
	AttributeLayerName      string
	AttributeLayerDimension DimensionType
	Attributes              []EASEnvironmentAttributeData
}

func (AttributeLayerSyncPacketDataUpdateEnvironmentAttributesData) isAttributeLayerSyncPacketData() {}

type AttributeModifier struct {
	Id             string
	Name           string
	Amount         float32
	Operation      int32
	Operand        int32
	IsSerializable bool
}

type AvailableCommandsChainedSubcommandData struct {
	Name             string
	SubCommandValues []AvailableCommandsChainedSubcommandRelationship
}

type AvailableCommandsChainedSubcommandRelationship struct {
	SubCommandFirstValue  uint32
	SubCommandSecondValue uint32
}

type AvailableCommandsConstrainedValueData struct {
	EnumValueSymbol   uint32
	EnumSymbol        uint32
	ConstraintIndices []uint8
}

type AvailableCommandsEnumData struct {
	Name   string
	Values []uint32
}

type AvailableCommandsOverloadData struct {
	IsChaining    bool
	ParameterData []AvailableCommandsParamData
}

type AvailableCommandsPacketCommandData struct {
	Name                                string
	Description                         string
	Flags                               uint16
	PermissionLevel                     string
	AliasEnum                           int32
	CommandDataChainedSubcommandIndexes []uint32
	Overloads                           []AvailableCommandsOverloadData
}

type AvailableCommandsParamData struct {
	Name        string
	ParseSymbol uint32
	IsOptional  bool
	Options     uint8
}

type AvailableCommandsSoftEnumData struct {
	EnumName    string
	EnumOptions []string
}

type BedrockDDUI interface {
	isBedrockDDUI()
}

type BedrockDDUIDataStoreChange struct {
	DataStoreName       string
	Property            string
	UpdateCount         uint32
	TheNewPropertyValue CerealDynamicValue
}

func (BedrockDDUIDataStoreChange) isBedrockDDUI() {}

type BedrockDDUIDataStoreRemoval struct {
	DataStoreName string
}

func (BedrockDDUIDataStoreRemoval) isBedrockDDUI() {}

type BedrockDDUIDataStoreUpdate struct {
	DataStoreName       string
	Property            string
	Path                string
	Data                BedrockDDUIDataStoreUpdateData
	PropertyUpdateCount uint32
	PathUpdateCount     uint32
}

func (BedrockDDUIDataStoreUpdate) isBedrockDDUI() {}

type BedrockDDUIDataStoreUpdateData interface {
	isBedrockDDUIDataStoreUpdateData()
}

type BedrockDDUIDataStoreUpdateDataBool struct {
	Value bool
}

func (BedrockDDUIDataStoreUpdateDataBool) isBedrockDDUIDataStoreUpdateData() {}

type BedrockDDUIDataStoreUpdateDataDouble struct {
	Value float64
}

func (BedrockDDUIDataStoreUpdateDataDouble) isBedrockDDUIDataStoreUpdateData() {}

type BedrockDDUIDataStoreUpdateDataString struct {
	Value string
}

func (BedrockDDUIDataStoreUpdateDataString) isBedrockDDUIDataStoreUpdateData() {}

type BedrockProfileWhiskerDiagnosticsScopeDataSummary struct {
	Label           string
	Indentation     string
	TotalHighCostNS uint64
	TotalMidCostNS  uint64
	TotalLowCostNS  uint64
}

type BedrockSafetyRedactableString struct {
	Unredacted string
	Redacted   *string
}

type BiomeCappedSurfaceData struct {
	FloorBlocks     []uint32
	CeilingBlocks   []uint32
	SeaBlock        *uint32
	FoundationBlock *uint32
	BeachBlock      *uint32
}

type BiomeClimateData struct {
	Temperature         float32
	Downfall            float32
	SnowAccumulationMin float32
	SnowAccumulationMax float32
}

type BiomeConditionalTransformationData struct {
	TransformsInto      []BiomeWeightedData
	ConditionJson       uint16
	MinPassingNeighbors uint32
}

type BiomeConsolidatedFeatureData struct {
	Scatter               BiomeScatterParamData
	Feature               uint16
	Identifier            uint16
	Pass                  uint16
	CanUseInternalFeature bool
}

type BiomeConsolidatedFeaturesData struct {
	Features []BiomeConsolidatedFeatureData
}

type BiomeCoordinateData struct {
	MinValueType int32
	MinValue     uint16
	MaxValueType int32
	MaxValue     uint16
	GridOffset   uint32
	GridStepSize uint32
	Distribution RandomDistributionType
}

type BiomeDefinitionChunkGenData struct {
	Climate                    *BiomeClimateData
	ConsolidatedFeatures       *BiomeConsolidatedFeaturesData
	MountainParams             *BiomeMountainParamsData
	SurfaceMaterialAdjustments *BiomeSurfaceMaterialAdjustmentData
	OverworldGenRules          *BiomeOverworldGenRulesData
	MultinoiseGenRules         *BiomeMultinoiseGenRulesData
	LegacyWorldGenRules        *BiomeLegacyWorldGenRulesData
	ReplacementBiomes          *BiomeReplacementsData
	VillageType                *VillageType
	SurfaceBuilderData         *BiomeSurfaceBuilderData
	SubsurfaceBuilderData      *BiomeSurfaceBuilderData
}

type BiomeDefinitionData struct {
	Id                uint16
	Temperature       float32
	Downfall          float32
	FoliageSnow       float32
	Depth             float32
	Scale             float32
	MapWaterColorARGB int32
	Rain              bool
	Tags              *BiomeTagsData
	ChunkGenData      *BiomeDefinitionChunkGenData
}

type BiomeElementData struct {
	NoiseFreqScale    float32
	NoiseLowerBound   float32
	NoiseUpperBound   float32
	HeightMinType     int32
	HeightMin         uint16
	HeightMaxType     int32
	HeightMax         uint16
	AdjustedMaterials BiomeSurfaceMaterialData
}

type BiomeLegacyWorldGenRulesData struct {
	LegacyPreHillsEdge []BiomeConditionalTransformationData
}

type BiomeMesaSurfaceData struct {
	ClayMaterial     uint32
	HardClayMaterial uint32
	BrycePillars     bool
	HasForest        bool
}

type BiomeMountainParamsData struct {
	SteepBlock      uint32
	NorthSlopes     bool
	SouthSlopes     bool
	WestSlopes      bool
	EastSlopes      bool
	TopSlideEnabled bool
}

type BiomeMultinoiseGenRulesData struct {
	Temperature float32
	Humidity    float32
	Altitude    float32
	Weirdness   float32
	Weight      float32
}

type BiomeNoiseGradientSurfaceData struct {
	NonReplaceableBlocks []uint32
	GradientBlocks       []SerializedNoiseBlockSpecifier
	Noise                NoiseDescriptor
}

type BiomeOverworldGenRulesData struct {
	HillsTransformations  []BiomeWeightedData
	MutateTransformations []BiomeWeightedData
	RiverTransformations  []BiomeWeightedData
	ShoreTransformations  []BiomeWeightedData
	PreHillsEdge          []BiomeConditionalTransformationData
	PostShoreEdge         []BiomeConditionalTransformationData
	Climate               []BiomeWeightedTemperatureData
}

type BiomeReplacementData struct {
	ReplacementBiome    uint16
	Dimension           uint16
	TargetBiomes        []uint16
	Amount              float32
	NoiseFrequencyScale float32
	ReplacementIndex    uint32
}

type BiomeReplacementsData struct {
	BiomeReplacements []BiomeReplacementData
}

type BiomeScatterParamData struct {
	Coordinates       []BiomeCoordinateData
	EvalOrder         CoordinateEvaluationOrder
	ChancePercentType int32
	ChancePercent     uint16
	ChanceNumerator   int32
	ChanceDenominator int32
	IterationsType    int32
	Iterations        uint16
}

type BiomeStringList struct {
	Strings []string
}

type BiomeSurfaceBuilderData struct {
	SurfaceMaterials           *BiomeSurfaceMaterialData
	HasDefaultOverworldSurface bool
	HasSwampSurface            bool
	HasFrozenOceanSurface      bool
	HasTheEndSurface           bool
	MesaSurface                *BiomeMesaSurfaceData
	CappedSurface              *BiomeCappedSurfaceData
	NoiseGradientSurface       *BiomeNoiseGradientSurfaceData
}

type BiomeSurfaceMaterialAdjustmentData struct {
	Adjustments []BiomeElementData
}

type BiomeSurfaceMaterialData struct {
	TopBlock        uint32
	MidBlock        uint32
	SeaFloorBlock   uint32
	FoundationBlock uint32
	SeaBlock        uint32
	SeaFloorDepth   int32
}

type BiomeTagsData struct {
	Tags []uint16
}

type BiomeWeightedData struct {
	BiomeIdentifier uint16
	Weight          uint32
}

type BiomeWeightedTemperatureData struct {
	Temperature int32
	Weight      uint32
}

type BlockPos struct {
	X int32
	Y int32
	Z int32
}

type BookEditAction interface {
	isBookEditAction()
}

type BookEditActionAddPage struct {
	PageIndex int32
	PageText  string
	PhotoName string
}

func (BookEditActionAddPage) isBookEditAction() {}

type BookEditActionDeletePage struct {
	PageIndex int32
}

func (BookEditActionDeletePage) isBookEditAction() {}

type BookEditActionFinalize struct {
	Title  string
	Author string
	XUID   string
}

func (BookEditActionFinalize) isBookEditAction() {}

type BookEditActionReplacePage struct {
	PageIndex int32
	PageText  string
	PhotoName string
}

func (BookEditActionReplacePage) isBookEditAction() {}

type BookEditActionSwapPages struct {
	PageIndex     int32
	SwapWithIndex int32
}

func (BookEditActionSwapPages) isBookEditAction() {}

type BoxData struct {
	BoxBound mgl32.Vec3
}

func (BoxData) isPrimitiveShapeDataExtraShapeData() {}

type CameraAimAssistActorPriorityPriorityData struct {
	PresetIndex   int32
	CategoryIndex int32
	ActorIndex    int32
	PriorityValue int32
}

type CameraAimAssistCategoryDefinition struct {
	Name       string
	Priorities CameraAimAssistCategoryPriorities
}

type CameraAimAssistCategoryPriorities struct {
	Entities           []OrderedEntry[string, int32]
	Blocks             []OrderedEntry[string, int32]
	BlockTags          []OrderedEntry[string, int32]
	EntityTypeFamilies []OrderedEntry[string, int32]
	EntityDefault      *int32
	BlockDefault       *int32
}

type CameraAimAssistCommandPresetDefinition struct {
	PresetId   *string
	TargetMode *CameraAimAssistTargetMode
	ViewAngle  *mgl32.Vec2
	Distance   *float32
}

type CameraAimAssistPresetDefinition struct {
	Identifier          string
	ExclusionSettings   CameraAimAssistPresetExclusionDefinition
	LiquidTargetingList []string
	ItemSettings        []OrderedEntry[string, string]
	DefaultItemSettings *string
	HandSettings        *string
}

type CameraAimAssistPresetExclusionDefinition struct {
	Blocks             []string
	Entities           []string
	BlockTags          []string
	EntityTypeFamilies []string
}

type CameraInstructionData struct {
	Set              *CameraInstructionOptionsSetInstruction
	Clear            *bool
	Fade             *CameraInstructionOptionsFadeInstruction
	Target           *CameraInstructionOptionsTargetInstruction
	RemoveTarget     *bool
	FieldOfView      *CameraInstructionOptionsFovInstruction
	Spline           *CameraInstructionOptionsSplineInstruction
	AttachToEntity   *CameraInstructionOptionsAttachToEntityInstruction
	DetachFromEntity *bool
}

type CameraInstructionOptionsAttachToEntityInstruction struct {
	EntityActorID int64
}

type CameraInstructionOptionsFadeInstruction struct {
	Time  *CameraInstructionOptionsFadeInstructionTimeOption
	Color *CameraInstructionOptionsFadeInstructionColorOption
}

type CameraInstructionOptionsFadeInstructionColorOption struct {
	Red   float32
	Green float32
	Blue  float32
}

type CameraInstructionOptionsFadeInstructionTimeOption struct {
	FadeInTime  float32
	HoldTime    float32
	FadeOutTime float32
}

type CameraInstructionOptionsFovInstruction struct {
	FieldOfView      float32
	FOVEaseTime      float32
	FOVEaseType      string
	FieldOfViewClear bool
}

type CameraInstructionOptionsSetInstruction struct {
	Preset                              uint32
	Ease                                *CameraInstructionOptionsSetInstructionEaseOption
	Pos                                 *CameraInstructionOptionsSetInstructionPosOption
	Rot                                 *CameraInstructionOptionsSetInstructionRotOption
	Facing                              *CameraInstructionOptionsSetInstructionFacingOption
	ViewOffset                          *CameraInstructionOptionsSetInstructionViewOffsetOption
	EntityOffset                        *CameraInstructionOptionsSetInstructionEntityOffsetOption
	Default                             *bool
	RemoveIgnoreStartingValuesComponent bool
}

type CameraInstructionOptionsSetInstructionEaseOption struct {
	Type uint8
	Time float32
}

type CameraInstructionOptionsSetInstructionEntityOffsetOption struct {
	EntityOffsetX float32
	EntityOffsetY float32
	EntityOffsetZ float32
}

type CameraInstructionOptionsSetInstructionFacingOption struct {
	Pos mgl32.Vec3
}

type CameraInstructionOptionsSetInstructionPosOption struct {
	Pos mgl32.Vec3
}

type CameraInstructionOptionsSetInstructionRotOption struct {
	X float32
	Y float32
}

type CameraInstructionOptionsSetInstructionViewOffsetOption struct {
	X float32
	Y float32
}

type CameraInstructionOptionsSplineInstruction struct {
	TotalTime         float32
	Type              uint8
	Curve             []mgl32.Vec3
	ProgressKeyFrames []CameraInstructionOptionsSplineInstructionSplineProgressOption
	RotationOption    []CameraInstructionOptionsSplineInstructionSplineRotationOption
	SplineIdentifier  string
	LoadFromJson      bool
}

type CameraInstructionOptionsSplineInstructionSplineProgressOption struct {
	KeyFrameValue      float32
	KeyFrameTime       float32
	KeyFrameEasingFunc string
}

type CameraInstructionOptionsSplineInstructionSplineRotationOption struct {
	KeyFrameValue      mgl32.Vec3
	KeyFrameTime       float32
	KeyFrameEasingFunc string
}

type CameraInstructionOptionsTargetInstruction struct {
	TargetCenterOffset *mgl32.Vec3
	TargetActorID      int64
}

type CameraPreset struct {
	Name                    string
	InheritFrom             string
	PosX                    *float32
	PosY                    *float32
	PosZ                    *float32
	RotX                    *float32
	RotY                    *float32
	RotationSpeed           *float32
	SnapToTarget            *bool
	HorizontalRotationLimit *mgl32.Vec2
	VerticalRotationLimit   *mgl32.Vec2
	ContinueTargeting       *bool
	BlockListeningRadius    *float32
	ViewOffset              *mgl32.Vec2
	EntityOffset            *mgl32.Vec3
	Radius                  *float32
	YawLimitMin             *float32
	YawLimitMax             *float32
	Listener                *CameraPresetAudioListener
	PlayerEffects           *bool
	AimAssist               *CameraAimAssistCommandPresetDefinition
	ControlScheme           *ControlSchemeScheme
}

type CameraPresetsData struct {
	Presets []CameraPreset
}

type CameraSplineControlPoint struct {
	Position mgl32.Vec3
}

type CameraSplineDefinition struct {
	Name              string
	TotalTime         float32
	SplineType        string
	ControlPoints     []CameraSplineControlPoint
	ProgressKeyFrames []CameraSplineProgressKeyFrame
	RotationKeyFrames []CameraSplineRotationKeyFrame
}

type CameraSplineProgressKeyFrame struct {
	Progress float32
	Time     float32
	Easing   *string
}

type CameraSplineRotationKeyFrame struct {
	Rotation mgl32.Vec3
	Time     float32
	Easing   *string
}

type CerealDynamicValue interface {
	isCerealDynamicValue()
}

type CerealDynamicValueBool struct {
	Value bool
}

func (CerealDynamicValueBool) isCerealDynamicValue() {}

type CerealDynamicValueDouble struct {
	Value float64
}

func (CerealDynamicValueDouble) isCerealDynamicValue() {}

type CerealDynamicValueInt64 struct {
	Value int64
}

func (CerealDynamicValueInt64) isCerealDynamicValue() {}

type CerealDynamicValueList struct {
	Value []any
}

func (CerealDynamicValueList) isCerealDynamicValue() {}

type CerealDynamicValueMap struct {
	Value []OrderedEntry[string, any]
}

func (CerealDynamicValueMap) isCerealDynamicValue() {}

type CerealDynamicValueNone struct {
}

func (CerealDynamicValueNone) isCerealDynamicValue() {}

type CerealDynamicValueString struct {
	Value string
}

func (CerealDynamicValueString) isCerealDynamicValue() {}

type CerealizerExperimentsAnonExperimentToggle struct {
	Name    string
	Enabled bool
}

type CerealizerNetworkItemInstanceDescriptorSerializedData struct {
	Id             int32
	StackSize      uint16
	AuxValue       uint32
	BlockRuntimeId int32
	UserDataBuffer string
}

type CerealizerNetworkItemStackDescriptorSerializedData struct {
	Id             int16
	StackSize      uint16
	AuxValue       uint32
	NetIdVariant   *int32
	BlockRuntimeId uint32
	UserDataBuffer string
}

type CerealizerRecipeIngredientSerializedData struct {
	Descriptor []OrderedEntry[string, string]
	AuxValue   int32
	StackSize  int32
}

type CerealizerRecipeUnlockingRequirementSerializedData struct {
	UnlockingContext     RecipeUnlockingRequirementUnlockingContext
	UnlockingIngredients *[]CerealizerRecipeIngredientSerializedData
}

type ChangeEntityScore struct {
	Action        string
	ScoreboardId  ScoreboardId
	ObjectiveName string
	ScoreValue    int32
	ActorId       ActorUniqueID
}

func (ChangeEntityScore) isSetScoreScoreInfoItem() {}

type ChangeFakePlayerScore struct {
	Action         string
	ScoreboardId   ScoreboardId
	ObjectiveName  string
	ScoreValue     int32
	FakePlayerName string
}

func (ChangeFakePlayerScore) isSetScoreScoreInfoItem() {}

type ChangePlayerScore struct {
	Action         string
	ScoreboardId   ScoreboardId
	ObjectiveName  string
	ScoreValue     int32
	PlayerUniqueId PlayerScoreboardId
}

func (ChangePlayerScore) isSetScoreScoreInfoItem() {}

type ChunkPos struct {
	X int32
	Z int32
}

type ClientboundDebugRendererDebugMarkerData struct {
	Text     string
	Position mgl32.Vec3
	Color    color.RGBA
	Duration uint64
}

type CommandBlockUpdateBlockCommandData struct {
	BlockPosition    BlockPos
	CommandBlockMode uint32
	RedstoneMode     bool
	IsConditional    bool
}

func (CommandBlockUpdateBlockCommandData) isCommandBlockUpdateTarget() {}

type CommandBlockUpdateEntityCommandTarget struct {
	TargetRuntimeID ActorRuntimeID
}

func (CommandBlockUpdateEntityCommandTarget) isCommandBlockUpdateTarget() {}

type CommandBlockUpdateTarget interface {
	isCommandBlockUpdateTarget()
}

type CommandOriginData struct {
	Type      string
	UUID      uuid.UUID
	RequestId string
	PlayerId  int64
}

type CommandOutputData struct {
	OutputType     string
	SuccessCount   uint32
	OutputMessages []CommandOutputMessage
	DataSet        *string
}

type CommandOutputMessage struct {
	MessageID  string
	Successful bool
	Parameters []string
}

type ConeData struct {
	Radii       mgl32.Vec2
	Height      float32
	NumSegments uint8
}

func (ConeData) isPrimitiveShapeDataExtraShapeData() {}

type ContainerMixDataEntry struct {
	FromItemId    int32
	ReagentItemId int32
	ToItemId      int32
}

type ContentIdentity struct {
	Identity string
}

type CreativeGroupInfo struct {
	CreativeCategory CreativeItemCategory
	Name             string
	GroupIconItem    CerealizerNetworkItemInstanceDescriptorSerializedData
}

type CreativeItemEntry struct {
	CreativeNetId TypedServerNetIdStructCreativeItemNetIdTag
	ItemInstance  CerealizerNetworkItemInstanceDescriptorSerializedData
	GroupIndex    uint32
}

type CylinderData struct {
	RadiusX     mgl32.Vec2
	RadiusZ     mgl32.Vec2
	Height      float32
	NumSegments uint8
}

func (CylinderData) isPrimitiveShapeDataExtraShapeData() {}

type DataItemByte struct {
	Type  DataItemType
	Value int8
}

func (DataItemByte) isDataItemEntryValue() {}

type DataItemCompoundTag struct {
	Type  DataItemType
	Value []byte
}

func (DataItemCompoundTag) isDataItemEntryValue() {}

type DataItemEntry struct {
	ID      uint32
	Payload DataItemEntryValue
}

type DataItemEntryValue interface {
	isDataItemEntryValue()
}

type DataItemFloat struct {
	Type  DataItemType
	Value float32
}

func (DataItemFloat) isDataItemEntryValue() {}

type DataItemInt struct {
	Type  DataItemType
	Value int32
}

func (DataItemInt) isDataItemEntryValue() {}

type DataItemInt64 struct {
	Type  DataItemType
	Value int64
}

func (DataItemInt64) isDataItemEntryValue() {}

type DataItemPos struct {
	Type  DataItemType
	Value BlockPos
}

func (DataItemPos) isDataItemEntryValue() {}

type DataItemShort struct {
	Type  DataItemType
	Value int16
}

func (DataItemShort) isDataItemEntryValue() {}

type DataItemString struct {
	Type  DataItemType
	Value string
}

func (DataItemString) isDataItemEntryValue() {}

type DataItemVec3 struct {
	Type  DataItemType
	Value mgl32.Vec3
}

func (DataItemVec3) isDataItemEntryValue() {}

type DimensionDefinitionGroupDimensionDefinition struct {
	HeightMaximum int32
	HeightMinimum int32
	GeneratorType GeneratorType
	DimensionType DimensionType
	PackId        uuid.UUID
}

type DimensionType struct {
	Value int32
}

type DisconnectMessages interface {
	isDisconnectMessages()
}

type DisconnectMessagesEmpty1 struct {
}

func (DisconnectMessagesEmpty1) isDisconnectMessages() {}

type DisconnectPacketMessages struct {
	Message         string
	FilteredMessage string
}

func (DisconnectPacketMessages) isDisconnectMessages() {}

type EAS interface {
	isEAS()
}

type EASAttributeLayerData struct {
	Name       string
	NoiseName  *string
	Dimension  DimensionType
	Settings   EASAttributeLayerSettings
	Attributes []EASEnvironmentAttributeData
}

type EASAttributeLayerSettings struct {
	Priority          int32
	Weight            float32
	Enabled           bool
	TransitionsPaused bool
}

type EASBoolAttributeData struct {
	Value     bool
	Operation string
}

func (EASBoolAttributeData) isEAS() {}

type EASColorAttributeData struct {
	Value     [4]int32
	Operation string
}

func (EASColorAttributeData) isEAS() {}

type EASEnvironmentAttributeData struct {
	AttributeName          string
	FromAttribute          *EAS
	Attribute              EAS
	ToAttribute            *EAS
	CurrentTransitionTicks uint32
	TotalTransitionTicks   uint32
	Easing                 string
	LocalTransitionTicks   uint32
	NoiseTransition        bool
}

type EASFloatAttributeData struct {
	Value         float32
	Operation     string
	ConstraintMin *float32
	ConstraintMax *float32
}

func (EASFloatAttributeData) isEAS() {}

type ECSProfilingDiagnosticsEntityDiagnosticTimingInfo struct {
	DisplayName    string
	Entity         string
	TimeInNS       uint64
	PercentOfTotal uint8
}

type ECSProfilingDiagnosticsSystemCategory struct {
	CategoryName string
	SystemIndex  uint64
}

type ECSProfilingDiagnosticsSystemDiagnosticTimingInfo struct {
	DisplayName    string
	SystemIndex    uint64
	TimeInNS       uint64
	PercentOfTotal uint8
}

type EduSharedUriResource struct {
	ButtonName string
	LinkUri    string
}

type EducationLevelSettings struct {
	CodeBuilderDefaultURI        string
	CodeBuilderTitle             string
	CanResizeCodeBuilder         bool
	DisableLegacyTitleBar        bool
	PostProcessFilter            string
	ScreenshotBorderResourcePath string
	AgentCapabilities            *AgentCapabilities
	LocalSettings                EducationLocalLevelSettings
	DeprecatedAlwaysFalse        bool
	ExternalLinkSettings         *ExternalLinkSettings
}

type EducationLocalLevelSettings struct {
	CodeBuilderOverrideUri *string
}

type EllipsoidData struct {
	Radii           mgl32.Vec3
	SegmentsPerAxis uint8
}

func (EllipsoidData) isPrimitiveShapeDataExtraShapeData() {}

type EnchantmentInstance struct {
	EnchantType  EnchantType
	EnchantLevel uint8
}

type EntityNetId struct {
	RawId uint32
}

type Experiments struct {
	Toggles                []CerealizerExperimentsAnonExperimentToggle
	ExperimentsEverToggled bool
}

type ExternalLinkSettings struct {
	URL         string
	DisplayName string
}

type FeatureRegistryFeatureBinaryJsonFormat struct {
	FeatureName      string
	BinaryJsonOutput string
}

type FloatRange struct {
	Min float32
	Max float32
}

type FullContainerName struct {
	ContainerName ContainerEnumName
	DynamicID     *uint32
}

type GameRule struct {
	RuleName          string
	RuleCanBeModified bool
	RuleValue         GameRuleRuleValue
}

type GameRuleRuleValue interface {
	isGameRuleRuleValue()
}

type GameRuleRuleValueBool struct {
	Value bool
}

func (GameRuleRuleValueBool) isGameRuleRuleValue() {}

type GameRuleRuleValueEmpty0 struct {
}

func (GameRuleRuleValueEmpty0) isGameRuleRuleValue() {}

type GameRuleRuleValueFloat struct {
	Value float32
}

func (GameRuleRuleValueFloat) isGameRuleRuleValue() {}

type GameRuleRuleValueInt32 struct {
	Value int32
}

func (GameRuleRuleValueInt32) isGameRuleRuleValue() {}

type GameRulesChangedPacketData struct {
	RulesList []GameRule
}

type InventoryAction struct {
	Source   InventorySource
	Slot     uint32
	FromItem CerealizerNetworkItemStackDescriptorSerializedData
	ToItem   CerealizerNetworkItemStackDescriptorSerializedData
}

type InventoryMismatchData struct {
	Actions InventoryTransactionData
}

func (InventoryMismatchData) isInventoryTransactionTransactionValue() {}

type InventoryOptions struct {
	LeftInventoryTab  InventoryLeftTabIndex
	RightInventoryTab InventoryRightTabIndex
	Filtering         bool
	LayoutInv         InventoryLayout
	LayoutCraft       InventoryLayout
}

type InventorySource struct {
	SourceType  InventorySourceType
	ContainerID **int8
	BitFlags    **InventorySourceInventorySourceFlags
}

type InventoryTransactionData struct {
	Actions *[]InventoryAction
}

type InventoryTransactionTransactionValue interface {
	isInventoryTransactionTransactionValue()
}

type ItemData struct {
	ItemName          string
	ItemId            int16
	IsComponentBased  bool
	ItemVersion       ItemVersion
	ItemComponentData []byte
}

type ItemEnchantOption struct {
	Cost         uint8
	Enchants     ItemEnchants
	EnchantName  string
	EnchantNetId TypedServerNetIdStructRecipeNetIdTag
}

type ItemEnchants struct {
	Slot         int32
	ItemEnchants [3][]EnchantmentInstance
}

type ItemReleaseInventoryTransaction struct {
	Actions      InventoryTransactionData
	ActionType   ItemReleaseInventoryTransactionActionType
	Slot         int32
	Item         CerealizerNetworkItemStackDescriptorSerializedData
	FromPosition mgl32.Vec3
}

func (ItemReleaseInventoryTransaction) isInventoryTransactionTransactionValue() {}

type ItemStackRequestCereal interface {
	isItemStackRequestCereal()
}

type ItemStackRequestCerealBeaconPaymentActionData struct {
	ActionType        ItemStackRequestActionType
	PrimaryEffectId   int32
	SecondaryEffectId int32
}

func (ItemStackRequestCerealBeaconPaymentActionData) isItemStackRequestCereal() {}

type ItemStackRequestCerealConsumeActionData struct {
	ActionType ItemStackRequestActionType
	Amount     uint8
	Source     ItemStackRequestCerealSlotInfoData
}

func (ItemStackRequestCerealConsumeActionData) isItemStackRequestCereal() {}

type ItemStackRequestCerealCraftCreativeActionData struct {
	ActionType              ItemStackRequestActionType
	CreativeItemNetId       uint32
	NumberOfRequestedCrafts uint8
}

func (ItemStackRequestCerealCraftCreativeActionData) isItemStackRequestCereal() {}

type ItemStackRequestCerealCraftLoomActionData struct {
	ActionType    ItemStackRequestActionType
	PatternNameId string
	NumCrafts     uint8
}

func (ItemStackRequestCerealCraftLoomActionData) isItemStackRequestCereal() {}

type ItemStackRequestCerealCraftNonImplementedActionData struct {
	ActionType ItemStackRequestActionType
}

func (ItemStackRequestCerealCraftNonImplementedActionData) isItemStackRequestCereal() {}

type ItemStackRequestCerealCraftRecipeActionData struct {
	ActionType              ItemStackRequestActionType
	RecipeNetId             TypedServerNetIdStructRecipeNetIdTag
	NumberOfRequestedCrafts uint8
}

func (ItemStackRequestCerealCraftRecipeActionData) isItemStackRequestCereal() {}

type ItemStackRequestCerealCraftRecipeAutoActionData struct {
	ActionType              ItemStackRequestActionType
	RecipeNetId             TypedServerNetIdStructRecipeNetIdTag
	NumberOfRequestedCrafts uint8
	Ingredients             []ItemStackRequestCerealRecipeIngredientData
}

func (ItemStackRequestCerealCraftRecipeAutoActionData) isItemStackRequestCereal() {}

type ItemStackRequestCerealCraftRecipeOptionalActionData struct {
	ActionType          ItemStackRequestActionType
	RecipeNetId         TypedServerNetIdStructRecipeNetIdTag
	FilteredStringIndex int32
}

func (ItemStackRequestCerealCraftRecipeOptionalActionData) isItemStackRequestCereal() {}

type ItemStackRequestCerealCraftRepairAndDisenchantActionData struct {
	ActionType              ItemStackRequestActionType
	RecipeNetId             int32
	NumberOfRequestedCrafts uint8
	RepairCost              int32
}

func (ItemStackRequestCerealCraftRepairAndDisenchantActionData) isItemStackRequestCereal() {}

type ItemStackRequestCerealCraftResultsActionData struct {
	ActionType   ItemStackRequestActionType
	CraftResults []ItemStackRequestCerealNetworkItemInstanceDescriptorData
	NumCrafts    uint8
}

func (ItemStackRequestCerealCraftResultsActionData) isItemStackRequestCereal() {}

type ItemStackRequestCerealCreateActionData struct {
	ActionType   ItemStackRequestActionType
	ResultsIndex uint8
}

func (ItemStackRequestCerealCreateActionData) isItemStackRequestCereal() {}

type ItemStackRequestCerealDestroyActionData struct {
	ActionType ItemStackRequestActionType
	Amount     uint8
	Source     ItemStackRequestCerealSlotInfoData
}

func (ItemStackRequestCerealDestroyActionData) isItemStackRequestCereal() {}

type ItemStackRequestCerealDropActionData struct {
	ActionType ItemStackRequestActionType
	Amount     uint8
	Source     ItemStackRequestCerealSlotInfoData
	Randomly   bool
}

func (ItemStackRequestCerealDropActionData) isItemStackRequestCereal() {}

type ItemStackRequestCerealEmptyItemDescriptorData struct {
	DescriptorType ItemStackRequestCerealItemDescriptorType
}

func (ItemStackRequestCerealEmptyItemDescriptorData) isItemStackRequestCerealRecipeIngredientDataItemDescriptor() {
}

type ItemStackRequestCerealItemNameDescriptorData struct {
	DescriptorType ItemStackRequestCerealItemDescriptorType
	FullName       string
	AuxValue       int32
}

func (ItemStackRequestCerealItemNameDescriptorData) isItemStackRequestCerealRecipeIngredientDataItemDescriptor() {
}

type ItemStackRequestCerealItemTagDescriptorData struct {
	DescriptorType ItemStackRequestCerealItemDescriptorType
	ItemTag        string
}

func (ItemStackRequestCerealItemTagDescriptorData) isItemStackRequestCerealRecipeIngredientDataItemDescriptor() {
}

type ItemStackRequestCerealLabTableCombineActionData struct {
	ActionType ItemStackRequestActionType
}

func (ItemStackRequestCerealLabTableCombineActionData) isItemStackRequestCereal() {}

type ItemStackRequestCerealMineBlockActionData struct {
	ActionType          ItemStackRequestActionType
	Slot                int32
	PredictedDurability int32
	NetIdVariant        int32
}

func (ItemStackRequestCerealMineBlockActionData) isItemStackRequestCereal() {}

type ItemStackRequestCerealMoLangItemDescriptorData struct {
	DescriptorType ItemStackRequestCerealItemDescriptorType
	TagExpression  string
	MolangVersion  MoLangVersion
}

func (ItemStackRequestCerealMoLangItemDescriptorData) isItemStackRequestCerealRecipeIngredientDataItemDescriptor() {
}

type ItemStackRequestCerealNetworkItemInstanceDescriptorData struct {
	ItemDescriptor ItemStackRequestCerealRecipeIngredientDataItemDescriptor
	StackSize      uint16
	BlockRuntimeId uint32
	UserDataBuffer string
}

type ItemStackRequestCerealPlaceActionData struct {
	ActionType  ItemStackRequestActionType
	Amount      uint8
	Source      ItemStackRequestCerealSlotInfoData
	Destination ItemStackRequestCerealSlotInfoData
}

func (ItemStackRequestCerealPlaceActionData) isItemStackRequestCereal() {}

type ItemStackRequestCerealRecipeIngredientData struct {
	ItemDescriptor ItemStackRequestCerealRecipeIngredientDataItemDescriptor
	StackSize      uint16
}

type ItemStackRequestCerealRecipeIngredientDataItemDescriptor interface {
	isItemStackRequestCerealRecipeIngredientDataItemDescriptor()
}

type ItemStackRequestCerealRequestData struct {
	ClientRequestId       TypedClientNetIdStructItemStackRequestIdTagInt32T0
	Actions               []ItemStackRequestCereal
	StringsToFilter       []string
	StringsToFilterOrigin TextProcessingEventOrigin
}

type ItemStackRequestCerealSlotInfoData struct {
	FullContainerName FullContainerName
	Slot              uint8
	NetIdVariant      int32
}

type ItemStackRequestCerealSwapActionData struct {
	ActionType  ItemStackRequestActionType
	Source      ItemStackRequestCerealSlotInfoData
	Destination ItemStackRequestCerealSlotInfoData
}

func (ItemStackRequestCerealSwapActionData) isItemStackRequestCereal() {}

type ItemStackRequestCerealTakeActionData struct {
	ActionType  ItemStackRequestActionType
	Amount      uint8
	Source      ItemStackRequestCerealSlotInfoData
	Destination ItemStackRequestCerealSlotInfoData
}

func (ItemStackRequestCerealTakeActionData) isItemStackRequestCereal() {}

type ItemStackRequestPacketDataRequestData struct {
	ClientRequestId       TypedClientNetIdStructItemStackRequestIdTagInt32T0
	Actions               []ItemStackRequestCereal
	StringsToFilter       []string
	StringsToFilterOrigin TextProcessingEventOrigin
}

type ItemStackResponseContainerInfo struct {
	FullContainerName FullContainerName
	Slots             []ItemStackResponseSlotInfo
}

type ItemStackResponseInfo struct {
	Result          ItemStackNetResult
	ClientRequestId TypedClientNetIdStructItemStackRequestIdTagInt32T0
	Containers      **[]ItemStackResponseContainerInfo
}

type ItemStackResponseSlotInfo struct {
	RequestedSlot        uint8
	Slot                 uint8
	Amount               uint8
	ItemStackNetId       **TypedServerNetIdStructItemStackNetIdTagInt32T0
	CustomName           BedrockSafetyRedactableString
	DurabilityCorrection int32
}

type ItemUseInventoryTransaction struct {
	Actions                  InventoryTransactionData
	ActionType               ItemUseInventoryTransactionActionType
	TriggerType              ItemUseInventoryTransactionTriggerType
	Position                 BlockPos
	Face                     uint8
	Slot                     int32
	Item                     CerealizerNetworkItemStackDescriptorSerializedData
	FromPosition             mgl32.Vec3
	ClickPosition            mgl32.Vec3
	TargetBlockId            uint32
	ClientInteractPrediction ItemUseInventoryTransactionPredictedResult
	ClientCooldownState      ItemUseInventoryTransactionClientCooldownState
}

func (ItemUseInventoryTransaction) isInventoryTransactionTransactionValue() {}

type ItemUseOnActorInventoryTransaction struct {
	Actions      InventoryTransactionData
	RuntimeId    ActorRuntimeID
	ActionType   ItemUseOnActorInventoryTransactionActionType
	Slot         int32
	Item         CerealizerNetworkItemStackDescriptorSerializedData
	FromPosition mgl32.Vec3
	HitPosition  mgl32.Vec3
}

func (ItemUseOnActorInventoryTransaction) isInventoryTransactionTransactionValue() {}

type LegacySetSlot struct {
	ContainerEnum ContainerEnumName
	Slots         []uint8
}

type LegacyTelemetryEventAchievement struct {
	AchievementID MinecraftEventingAchievementIds
}

func (LegacyTelemetryEventAchievement) isLegacyTelemetryEventEventData() {}

type LegacyTelemetryEventActorDefinition struct {
	EventName string
}

func (LegacyTelemetryEventActorDefinition) isLegacyTelemetryEventEventData() {}

type LegacyTelemetryEventBellUsed struct {
	ItemId int32
}

func (LegacyTelemetryEventBellUsed) isLegacyTelemetryEventEventData() {}

type LegacyTelemetryEventBossKilled struct {
	BossActorID int64
	PartySize   int32
	BossType    int32
}

func (LegacyTelemetryEventBossKilled) isLegacyTelemetryEventEventData() {}

type LegacyTelemetryEventCauldronUsed struct {
	ContentsColor uint32
	ContentsType  int32
	FillLevel     int32
}

func (LegacyTelemetryEventCauldronUsed) isLegacyTelemetryEventEventData() {}

type LegacyTelemetryEventCodeBuilderRuntimeAction struct {
	CodeBuilderRuntimeAction string
}

func (LegacyTelemetryEventCodeBuilderRuntimeAction) isLegacyTelemetryEventEventData() {}

type LegacyTelemetryEventCodeBuilderScoreboard struct {
	ObjectiveName string
	Score         int32
}

func (LegacyTelemetryEventCodeBuilderScoreboard) isLegacyTelemetryEventEventData() {}

type LegacyTelemetryEventComposterUsed struct {
	BlockInteractionType MinecraftEventingPOIBlockInteractionType
	ItemId               int32
}

func (LegacyTelemetryEventComposterUsed) isLegacyTelemetryEventEventData() {}

type LegacyTelemetryEventEmpty struct {
}

func (LegacyTelemetryEventEmpty) isLegacyTelemetryEventEventData() {}

type LegacyTelemetryEventEventData interface {
	isLegacyTelemetryEventEventData()
}

type LegacyTelemetryEventInteraction struct {
	InteractedEntityID      int64
	InteractionType         MinecraftEventingInteractionType
	InteractionActorType    int32
	InteractionActorVariant int32
	InteractionActorColor   uint8
}

func (LegacyTelemetryEventInteraction) isLegacyTelemetryEventEventData() {}

type LegacyTelemetryEventItemUsed struct {
	ItemId    int16
	ItemAux   int32
	UseMethod int32
	Count     int32
}

func (LegacyTelemetryEventItemUsed) isLegacyTelemetryEventEventData() {}

type LegacyTelemetryEventMobBorn struct {
	BornBabyEntityType    int32
	BornBabyEntityVariant int32
	BornBabyColor         uint8
}

func (LegacyTelemetryEventMobBorn) isLegacyTelemetryEventEventData() {}

type LegacyTelemetryEventMobKilled struct {
	InstigatorActorID         int64
	TargetActorID             int64
	InstigatorSChildActorType ActorType
	DamageSource              int32
	TradeTier                 int32
	TraderName                string
}

func (LegacyTelemetryEventMobKilled) isLegacyTelemetryEventEventData() {}

type LegacyTelemetryEventPOICauldronUsed struct {
	BlockInteractionType MinecraftEventingPOIBlockInteractionType
	ItemId               int32
}

func (LegacyTelemetryEventPOICauldronUsed) isLegacyTelemetryEventEventData() {}

type LegacyTelemetryEventPiglinBarter struct {
	ItemId                      int32
	WasTargetingBarteringPlayer bool
}

func (LegacyTelemetryEventPiglinBarter) isLegacyTelemetryEventEventData() {}

type LegacyTelemetryEventPlayerDied struct {
	InstigatorActorID    int32
	InstigatorMobVariant int32
	DamageSource         int32
	DiedInRaid           bool
}

func (LegacyTelemetryEventPlayerDied) isLegacyTelemetryEventEventData() {}

type LegacyTelemetryEventPlayerWaxedOrUnwaxedCopper struct {
	PlayerWaxedOrUnwaxedCopperBlockID int32
}

func (LegacyTelemetryEventPlayerWaxedOrUnwaxedCopper) isLegacyTelemetryEventEventData() {}

type LegacyTelemetryEventPortalCreated struct {
	DimensionID int32
}

func (LegacyTelemetryEventPortalCreated) isLegacyTelemetryEventEventData() {}

type LegacyTelemetryEventPortalUsed struct {
	SourceDimensionID int32
	TargetDimensionID int32
}

func (LegacyTelemetryEventPortalUsed) isLegacyTelemetryEventEventData() {}

type LegacyTelemetryEventRaidUpdate struct {
	CurrentWave int32
	TotalWaves  int32
	Success     bool
}

func (LegacyTelemetryEventRaidUpdate) isLegacyTelemetryEventEventData() {}

type LegacyTelemetryEventSlashCommand struct {
	SuccessCount int32
	ErrorCount   int32
	CommandName  string
	ErrorList    string
}

func (LegacyTelemetryEventSlashCommand) isLegacyTelemetryEventEventData() {}

type LegacyTelemetryEventTargetBlockHit struct {
	RedstoneLevel int32
}

func (LegacyTelemetryEventTargetBlockHit) isLegacyTelemetryEventEventData() {}

type LevelChunkSubChunkMetadata struct {
	BlobId uint64
}

type LevelSettings struct {
	Seed                                   uint64
	SpawnSettings                          SpawnSettings
	GeneratorType                          GeneratorType
	GameType                               GameType
	IsHardcore                             bool
	GameDifficulty                         LegacyDifficulty
	DefaultSpawnBlockPosition              BlockPos
	AchievementsDisabled                   bool
	EditorWorldType                        EditorWorldType
	IsCreatedInEditor                      bool
	IsExportedFromEditor                   bool
	DayCycleStopTime                       int32
	EducationEditionOffer                  EducationEditionOffer
	EducationFeaturesEnabled               bool
	EducationProductID                     string
	RainLevel                              float32
	LightningLevel                         float32
	HasConfirmedPlatformLockedContent      bool
	MultiplayerGameIntent                  bool
	LANBroadcastIntent                     bool
	XboxLiveBroadcastSetting               SocialGamePublishSetting
	PlatformBroadcastSetting               SocialGamePublishSetting
	CommandsEnabled                        bool
	TexturePacksRequired                   bool
	RuleData                               GameRulesChangedPacketData
	Experiments                            Experiments
	HasBonusChestEnabled                   bool
	StartWithMapEnabled                    bool
	PlayerPermissions                      PlayerPermissionLevel
	ServerChunkTickRange                   int32
	HasLockedBehaviorPack                  bool
	HasLockedResourcePack                  bool
	IsFromLockedTemplate                   bool
	UseMsaGamertagsOnly                    bool
	IsFromWorldTemplate                    bool
	IsWorldTemplateOptionLocked            bool
	OnlySpawnV1Villagers                   bool
	PersonaDisabled                        bool
	CustomSkinsDisabled                    bool
	EmoteChatMuted                         bool
	BaseGameVersion                        string
	LimitedWorldWidth                      int32
	LimitedWorldDepth                      int32
	NetherType                             bool
	EduSharedUriResource                   EduSharedUriResource
	OverrideForceExperimentalGameplay      *bool
	ChatRestrictionLevel                   ChatRestrictionLevel
	DisablePlayerInteractions              bool
	ServerEditorConnectionPolicy           ServerEditorConnectionPolicy
	AllowAnonymousBlockDropsInEditorWorlds bool
}

type LineData struct {
	LineEndLocation mgl32.Vec3
}

func (LineData) isPrimitiveShapeDataExtraShapeData() {}

type LocatorBarWaypoint struct {
	GroupHandle           WaypointGroupWaypointHandle
	ServerWaypointPayload ServerWaypoint
	ActionFlag            ServerWaypointGroupAction
}

type MapDecoration struct {
	ImageType MapDecorationType
	Rotation  uint8
	X         uint8
	Y         uint8
	Label     string
	Color     color.RGBA
}

type MapInfoRequestPacketAnonClientPixelsProxy struct {
	Pixel uint32
	Index uint16
}

type MapItemTrackedActorUniqueId struct {
	Type          MapItemTrackedActorType
	EntityID      *ActorUniqueID
	BlockPosition *BlockPos
}

type MaterialReducerDataEntry struct {
	FromItemKey      int32
	ItemIdsAndCounts []MaterialReducerEntryOutput
}

type MaterialReducerEntryOutput struct {
	ItemId    int32
	ItemCount int32
}

type MemoryMemoryCategoryCounter struct {
	Category     MemoryMemoryCategory
	CurrentBytes uint64
}

type MissingBlobData struct {
	BlobId   uint64
	BlobData string
}

type MoveActorAbsoluteData struct {
	ActorRuntimeID ActorRuntimeID
	Header         uint8
	Position       mgl32.Vec3
	RotationX      uint8
	RotationY      uint8
	RotationYHead  uint8
}

type MoveActorDeltaData struct {
	ActorRuntimeID       ActorRuntimeID
	NewPositionX         *float32
	NewPositionY         *float32
	NewPositionZ         *float32
	RotationX            *int8
	RotationY            *int8
	RotationYHead        *int8
	IsOnGround           bool
	ForceMove            bool
	ForceMoveLocalEntity bool
	ForceCompletion      bool
}

type MovePlayerTeleportData struct {
	TeleportationCause int32
	SourceActorType    int32
}

type MultiRecipe struct {
	MultiRecipeUUID uuid.UUID
	NetId           TypedServerNetIdStructRecipeNetIdTag
}

type NetworkPermissions struct {
	ServerAuthSoundEnabled bool
}

type NoiseDescriptor struct {
	Name        string
	FirstOctave int32
	Amplitudes  []float32
}

type NormalTransactionData struct {
	Actions InventoryTransactionData
}

func (NormalTransactionData) isInventoryTransactionTransactionValue() {}

type PackIdVersion struct {
	PackUUID    uuid.UUID
	PackVersion SemVersion
}

type PackIdVersionData struct {
	PackUUID    uuid.UUID
	PackVersion SemVersionData
}

type PackInfoData struct {
	PackIdVersion       PackIdVersionData
	PackSize            uint64
	ContentKey          string
	SubpackName         string
	ContentIdentity     ContentIdentity
	HasScripts          bool
	IsAddonPack         bool
	IsRayTracingCapable bool
	CDNURL              string
}

type PackInstanceId struct {
	PackID      string
	Version     string
	SubPackName string
}

type PackedItemUseLegacyInventoryTransaction struct {
	LegacyRequestID    TypedClientNetIdStructItemStackLegacyRequestIdTagInt32T0
	LegacySetItemSlots *[]LegacySetSlot
	ItemUseTransaction *ItemUseInventoryTransaction
}

type PlayerBlockActionData struct {
	PlayerActionType PlayerActionType
	Position         BlockPos
	Facing           int32
}

type PlayerInputTick struct {
	InputTick uint64
}

type PlayerListAddEntry struct {
	UUID             uuid.UUID
	ActorUniqueID    ActorUniqueID
	PlayerName       string
	XBLXUID          string
	PlatformOnlineID string
	BuildPlatform    BuildPlatform
	SerializedSkin   SerializedSkinRef
	IsTeacher        bool
	IsHost           bool
	IsSubClient      bool
	PlayerColor      color.RGBA
}

func (PlayerListAddEntry) isPlayerListEntriesItem() {}

type PlayerListEntriesItem interface {
	isPlayerListEntriesItem()
}

type PlayerListRemoveEntry struct {
	UUID uuid.UUID
}

func (PlayerListRemoveEntry) isPlayerListEntriesItem() {}

type PlayerLocationCoordinatesLocation struct {
	PacketType PlayerLocationType
	Position   mgl32.Vec3
}

func (PlayerLocationCoordinatesLocation) isPlayerLocationLocation() {}

type PlayerLocationHiddenLocation struct {
	PacketType PlayerLocationType
}

func (PlayerLocationHiddenLocation) isPlayerLocationLocation() {}

type PlayerLocationLocation interface {
	isPlayerLocationLocation()
}

type PlayerPartyInfo struct {
	PartyId       string
	IsPartyLeader bool
}

type PlayerScoreboardId struct {
	PlayerUniqueId int64
}

type PlayerUpdateEntityOverridesClearOverride struct {
	Type string
}

func (PlayerUpdateEntityOverridesClearOverride) isPlayerUpdateEntityOverridesUpdate() {}

type PlayerUpdateEntityOverridesFloatOverride struct {
	Type  string
	Value float32
}

func (PlayerUpdateEntityOverridesFloatOverride) isPlayerUpdateEntityOverridesUpdate() {}

type PlayerUpdateEntityOverridesIntOverride struct {
	Type  string
	Value int32
}

func (PlayerUpdateEntityOverridesIntOverride) isPlayerUpdateEntityOverridesUpdate() {}

type PlayerUpdateEntityOverridesRemoveOverride struct {
	Type string
}

func (PlayerUpdateEntityOverridesRemoveOverride) isPlayerUpdateEntityOverridesUpdate() {}

type PlayerUpdateEntityOverridesUpdate interface {
	isPlayerUpdateEntityOverridesUpdate()
}

type PlayerVideoCaptureAction interface {
	isPlayerVideoCaptureAction()
}

type PlayerVideoCaptureStartVideoCapture struct {
	FrameRate  uint32
	FilePrefix string
}

func (PlayerVideoCaptureStartVideoCapture) isPlayerVideoCaptureAction() {}

type PlayerVideoCaptureStopVideoCapture struct {
}

func (PlayerVideoCaptureStopVideoCapture) isPlayerVideoCaptureAction() {}

type PositionTrackingId struct {
	Value int32
}

type PotionMixDataEntry struct {
	FromPotionId   int32
	FromItemAux    int32
	ReagentItemId  int32
	ReagentItemAux int32
	ToPotionId     int32
	ToItemAux      int32
}

type PrimitiveShapeData struct {
	NetworkId             uint64
	ShapeType             *ScriptModuleMinecraftScriptPrimitiveShapeType
	Location              *mgl32.Vec3
	Scale                 *float32
	Rotation              *mgl32.Vec3
	TotalTimeLeft         *float32
	MaximumRenderDistance *float32
	Color                 *color.RGBA
	DimensionID           *DimensionType
	AttachedToEntityID    *ActorUniqueID
	ExtraShapeData        PrimitiveShapeDataExtraShapeData
}

type PrimitiveShapeDataExtraShapeData interface {
	isPrimitiveShapeDataExtraShapeData()
}

type PrimitiveShapeDataExtraShapeDataEmpty0 struct {
}

func (PrimitiveShapeDataExtraShapeDataEmpty0) isPrimitiveShapeDataExtraShapeData() {}

type PropertySyncData struct {
	IntEntriesList   []PropertySyncDataPropertySyncIntEntry
	FloatEntriesList []PropertySyncDataPropertySyncFloatEntry
}

type PropertySyncDataPropertySyncFloatEntry struct {
	PropertyIndex uint32
	Data          float32
}

type PropertySyncDataPropertySyncIntEntry struct {
	PropertyIndex uint32
	Data          int32
}

type PyramidData struct {
	Width  float32
	Depth  *float32
	Height float32
}

func (PyramidData) isPrimitiveShapeDataExtraShapeData() {}

type RemoveScore struct {
	Action        string
	ScoreboardId  ScoreboardId
	ObjectiveName *string
}

func (RemoveScore) isSetScoreScoreInfoItem() {}

type ResourcePackClientResponseCancel struct {
	ResponseType string
}

func (ResourcePackClientResponseCancel) isResourcePackClientResponseResponse() {}

type ResourcePackClientResponseDownloading struct {
	ResponseType     string
	DownloadingPacks []string
}

func (ResourcePackClientResponseDownloading) isResourcePackClientResponseResponse() {}

type ResourcePackClientResponseDownloadingFinished struct {
	ResponseType string
}

func (ResourcePackClientResponseDownloadingFinished) isResourcePackClientResponseResponse() {}

type ResourcePackClientResponseResourcePackStackFinished struct {
	ResponseType string
}

func (ResourcePackClientResponseResourcePackStackFinished) isResourcePackClientResponseResponse() {}

type ResourcePackClientResponseResponse interface {
	isResourcePackClientResponseResponse()
}

type ScoreboardId struct {
	ScoreboardId int64
}

type ScoreboardIdentityPacketInfo struct {
	ScoreboardId   ScoreboardId
	PlayerUniqueId *int64
}

type SemVersion struct {
	Version string
}

type SemVersionData struct {
	Version string
}

type SerializedAbilitiesData struct {
	TargetPlayerRawId  int64
	PlayerPermissions  PlayerPermissionLevel
	CommandPermissions CommandPermissionLevel
	Layers             []SerializedAbilitiesDataSerializedLayer
}

type SerializedAbilitiesDataSerializedLayer struct {
	SerializedLayer  uint16
	AbilitiesSet     uint32
	AbilityValues    uint32
	FlySpeed         float32
	VerticalFlySpeed float32
	WalkSpeed        float32
}

type SerializedNoiseBlockSpecifier struct {
	Noise     string
	Threshold float32
	Range     FloatRange
	Block     uint32
}

type SerializedPersonaPieceHandle struct {
	PieceId        string
	PieceType      PersonaPieceType
	PackId         uuid.UUID
	IsDefaultPiece bool
	ProductId      string
}

type SerializedSkinRef struct {
	ID                           string
	PlayFabID                    string
	ResourcePatch                string
	ImageData                    SkinImage
	AnimatedImageData            []AnimatedImageData
	CapeImageData                SkinImage
	GeometryData                 string
	GeometryDataMinEngineVersion string
	AnimationData                string
	CapeID                       string
	FullID                       string
	ArmSize                      PersonaArmSizeType
	SkinColor                    color.RGBA
	PersonaPieces                []SerializedPersonaPieceHandle
	PieceTintColors              []OrderedEntry[string, TintMapColor]
	IsPremium                    bool
	IsPersona                    bool
	IsPersonaCapeOnClassicSkin   bool
	IsPrimaryUser                bool
	OverridesPlayerAppearance    bool
	TrustedSkinFlag              string
	ProfileHash                  string
}

type ServerBlockProperty struct {
	BlockName       string
	BlockDefinition []byte
}

type ServerConfigurationClientStoreEntryPointConfiguration struct {
	StoreId   string
	StoreName string
}

type ServerConfigurationGatheringsConfigurationJoinInfo struct {
	ExperienceId   uuid.UUID
	ExperienceName string
	WorldId        *uuid.UUID
	WorldName      *string
	CreatorId      string
	TargetId       *uuid.UUID
	ScenarioId     *string
	ServerId       *string
}

type ServerConfigurationPresenceConfiguration struct {
	RichPresenceId *string
}

type ServerConfigurationServerConfigurationJoinInfo struct {
	Gathering             *ServerConfigurationGatheringsConfigurationJoinInfo
	ClientStoreEntryPoint *ServerConfigurationClientStoreEntryPointConfiguration
	Presence              *ServerConfigurationPresenceConfiguration
}

type ServerSoundHandle struct {
	ServerSoundHandle uint64
}

type ServerWaypoint struct {
	UpdateFlag              uint32
	IsVisible               *bool
	WorldPosition           *WorldPosition
	TexturePath             *string
	IconSize                *mgl32.Vec2
	Color                   *color.RGBA
	ClientPositionAuthority *bool
	ActorUniqueID           *ActorUniqueID
}

type ServerboundPackSettingChangePackSettingValue interface {
	isServerboundPackSettingChangePackSettingValue()
}

type ServerboundPackSettingChangePackSettingValueBool struct {
	Value bool
}

func (ServerboundPackSettingChangePackSettingValueBool) isServerboundPackSettingChangePackSettingValue() {
}

type ServerboundPackSettingChangePackSettingValueFloat struct {
	Value float32
}

func (ServerboundPackSettingChangePackSettingValueFloat) isServerboundPackSettingChangePackSettingValue() {
}

type ServerboundPackSettingChangePackSettingValueString struct {
	Value string
}

func (ServerboundPackSettingChangePackSettingValueString) isServerboundPackSettingChangePackSettingValue() {
}

type SetScoreScoreInfoItem interface {
	isSetScoreScoreInfoItem()
}

type ShapedRecipe struct {
	RecipeId             string
	Width                int32
	Height               int32
	Ingredients          []CerealizerRecipeIngredientSerializedData
	Results              []CerealizerNetworkItemInstanceDescriptorSerializedData
	UUID                 uuid.UUID
	Tag                  string
	Priority             int32
	AssumeSymmetry       bool
	UnlockingRequirement *CerealizerRecipeUnlockingRequirementSerializedData
	NetId                TypedServerNetIdStructRecipeNetIdTag
}

type ShapelessRecipe struct {
	RecipeId             string
	Ingredients          []CerealizerRecipeIngredientSerializedData
	Results              []CerealizerNetworkItemInstanceDescriptorSerializedData
	UUID                 uuid.UUID
	Tag                  string
	Priority             int32
	UnlockingRequirement *CerealizerRecipeUnlockingRequirementSerializedData
	NetId                TypedServerNetIdStructRecipeNetIdTag
}

type SkinImage struct {
	Width      uint32
	Height     uint32
	ImageBytes []uint8
}

type SmithingTransformRecipe struct {
	RecipeId           string
	TemplateIngredient CerealizerRecipeIngredientSerializedData
	BaseIngredient     CerealizerRecipeIngredientSerializedData
	AdditionIngredient CerealizerRecipeIngredientSerializedData
	Result             CerealizerNetworkItemInstanceDescriptorSerializedData
	Tag                string
	NetId              TypedServerNetIdStructRecipeNetIdTag
}

type SmithingTrimRecipe struct {
	RecipeId           string
	TemplateIngredient CerealizerRecipeIngredientSerializedData
	BaseIngredient     CerealizerRecipeIngredientSerializedData
	AdditionIngredient CerealizerRecipeIngredientSerializedData
	Tag                string
	NetId              TypedServerNetIdStructRecipeNetIdTag
}

type SocialEventsServerTelemetryData struct {
	ServerId   string
	ScenarioId string
	WorldId    string
	OwnerId    string
}

type SoundDataEvent interface {
	isSoundDataEvent()
}

type SoundDataEventFade struct {
	Duration     float32
	TargetVolume float32
}

func (SoundDataEventFade) isSoundDataEvent() {}

type SoundDataEventPause struct {
}

func (SoundDataEventPause) isSoundDataEvent() {}

type SoundDataEventResume struct {
}

func (SoundDataEventResume) isSoundDataEvent() {}

type SoundDataEventSeekTo struct {
	Seconds float32
}

func (SoundDataEventSeekTo) isSoundDataEvent() {}

type SoundDataEventSetPitch struct {
	Pitch float32
}

func (SoundDataEventSetPitch) isSoundDataEvent() {}

type SoundDataEventSetVolume struct {
	Volume float32
}

func (SoundDataEventSetVolume) isSoundDataEvent() {}

type SoundDataEventStop struct {
}

func (SoundDataEventStop) isSoundDataEvent() {}

type SpawnSettings struct {
	SpawnBiomeType       SpawnBiomeType
	UserDefinedBiomeName string
	Dimension            int32
}

type SphereData struct {
	NumSegments uint8
}

func (SphereData) isPrimitiveShapeDataExtraShapeData() {}

type StructureEditorData struct {
	StructureName         BedrockSafetyRedactableString
	DataField             string
	ShouldIncludePlayers  bool
	ShouldShowBoundingBox bool
	StructureBlockType    StructureBlockType
	StructureSettings     StructureSettings
	RedstoneSaveMode      StructureRedstoneSaveMode
}

type StructureSettings struct {
	StructurePaletteName                            string
	ShouldIgnoreEntities                            bool
	ShouldIgnoreBlocks                              bool
	ShouldAllowNonTickingPlayerAndTickingAreaChunks bool
	StructureSize                                   BlockPos
	StructureOffset                                 BlockPos
	LastEditPlayer                                  ActorUniqueID
	Rotation                                        Rotation
	Mirror                                          Mirror
	AnimationMode                                   AnimationMode
	AnimationSeconds                                float32
	IntegrityValue                                  float32
	IntegritySeed                                   uint32
	RotationPivot                                   mgl32.Vec3
}

type SubChunkHeightmapData struct {
	HeightMapType           SubChunkHeightMapDataType
	SubchunkHeightMap       *[16][16]int8
	RenderHeightMapType     SubChunkHeightMapDataType
	SubchunkRenderHeightMap *[16][16]int8
}

type SubChunkPos struct {
	SubchunkPositionX int32
	SubchunkPositionY int32
	SubchunkPositionZ int32
}

type SubChunkSubChunkPacketData struct {
	SubChunkPosOffset     SubChunkSubChunkPosOffset
	SubChunkRequestResult SubChunkSubChunkRequestResult
	SerializedSubChunk    *string
	HeightMapData         SubChunkHeightmapData
	BlobId                *uint64
}

type SubChunkSubChunkPosOffset struct {
	SubchunkOffsetX int8
	SubchunkOffsetY int8
	SubchunkOffsetZ int8
}

type SyncWorldClockStateData struct {
	ClockId  uint64
	Time     int32
	IsPaused bool
}

type SyncWorldClocksAddTimeMarkerData struct {
	ClockId     uint64
	TimeMarkers []TimeMarkerData
}

func (SyncWorldClocksAddTimeMarkerData) isSyncWorldClocksData() {}

type SyncWorldClocksData interface {
	isSyncWorldClocksData()
}

type SyncWorldClocksInitializeRegistryData struct {
	ClockData []WorldClockData
}

func (SyncWorldClocksInitializeRegistryData) isSyncWorldClocksData() {}

type SyncWorldClocksRemoveTimeMarkerData struct {
	ClockId       uint64
	TimeMarkerIds []uint64
}

func (SyncWorldClocksRemoveTimeMarkerData) isSyncWorldClocksData() {}

type SyncWorldClocksSyncStateData struct {
	ClockData []SyncWorldClockStateData
}

func (SyncWorldClocksSyncStateData) isSyncWorldClocksData() {}

type SyncedAttribute struct {
	AttributeName string
	MinValue      float32
	CurrentValue  float32
	MaxValue      float32
}

type SyncedPlayerMovementSettings struct {
	RewindHistorySize                int32
	ServerAuthoritativeBlockBreaking bool
}

type SynchedActorDataCopyableDataList struct {
	Data []DataItemEntry
}

type TextAuthorAndMessage struct {
	PlayerName string
	Message    string
}

func (TextAuthorAndMessage) isTextBody() {}

type TextBody interface {
	isTextBody()
}

type TextData struct {
	Text             string
	UseRotation      bool
	BackgroundColor  *color.RGBA
	DepthTest        bool
	ShowBackface     bool
	ShowTextBackface bool
}

func (TextData) isPrimitiveShapeDataExtraShapeData() {}

type TextMessageAndParams struct {
	Message       string
	ParameterList []string
}

func (TextMessageAndParams) isTextBody() {}

type TextMessageOnly struct {
	Message string
}

func (TextMessageOnly) isTextBody() {}

type TimeMarkerData struct {
	Id     uint64
	Name   string
	Time   int32
	Period *int32
}

type TintMapColor struct {
	Colors [4]color.RGBA
}

type TrimMaterial struct {
	MaterialId string
	Color      string
	ItemName   string
}

type TrimPattern struct {
	ItemName  string
	PatternId string
}

type TypedClientNetIdStructItemStackLegacyRequestIdTagInt32T0 struct {
	ID int32
}

type TypedClientNetIdStructItemStackRequestIdTagInt32T0 struct {
	ID int32
}

type TypedServerNetIdStructCreativeItemNetIdTag struct {
	ID uint32
}

type TypedServerNetIdStructItemStackNetIdTagInt32T0 struct {
	ID int32
}

type TypedServerNetIdStructRecipeNetIdTag struct {
	RawId uint32
}

type UpdateSubChunkBlocksChangedInfo struct {
	BlocksChangedStandards []UpdateSubChunkNetworkBlockInfo
	BlocksChangedExtras    []UpdateSubChunkNetworkBlockInfo
}

type UpdateSubChunkNetworkBlockInfo struct {
	Pos                       BlockPos
	RuntimeId                 uint32
	UpdateFlags               uint32
	SyncMessageEntityUniqueID uint64
	SyncMessageMessage        uint32
}

type VoxelShapesRegistryHandle struct {
	Value uint16
}

type VoxelShapesSerializableCells struct {
	XSize   uint8
	YSize   uint8
	ZSize   uint8
	Storage []uint8
}

type VoxelShapesSerializableVoxelShape struct {
	Cells        VoxelShapesSerializableCells
	XCoordinates []float32
	YCoordinates []float32
	ZCoordinates []float32
}

type WaypointGroupWaypointHandle struct {
	UUID uuid.UUID
}

type WebSocketPacketData struct {
	WebsocketServerURI string
}

type WorldClockData struct {
	Id          uint64
	Name        string
	Time        int32
	IsPaused    bool
	TimeMarkers []TimeMarkerData
}

type WorldPosition struct {
	Position      mgl32.Vec3
	DimensionType DimensionType
}
