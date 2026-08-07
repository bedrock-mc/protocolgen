// Code generated from canonical protocol manifest v2. DO NOT EDIT.

use crate::enums::*;

use crate::wire;

#[derive(Clone, Debug, Default, PartialEq, Eq)]
pub struct Nbt(pub Vec<u8>);

#[derive(Clone, Debug, Default, PartialEq)]
pub struct ActorDataBoundingBoxComponent {
    pub actor_data_bounding_box: [wire::F32LE; 3],
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct ActorDataFlagComponent {
    pub actor_flag_bitset_data: Bitset131,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct ActorLink {
    pub target_a: ActorUniqueID,
    pub target_b: ActorUniqueID,
    pub type_: ActorLinkType,
    pub immediate: bool,
    pub passenger_initiated: bool,
    pub vehicle_angular_velocity: wire::F32LE,
}

#[derive(Clone, Copy, Debug, Default, PartialEq, Eq, Hash)]
pub struct ActorRuntimeID(pub u64);

impl wire::WireCodec for ActorRuntimeID {
    fn encode<W: std::io::Write>(&self, writer: &mut W) -> std::io::Result<()> {
        <wire::VarULong as wire::WireCodec>::encode(&wire::VarULong(self.0), writer)
    }

    fn decode<R: std::io::Read>(reader: &mut R) -> std::io::Result<Self> {
        <wire::VarULong as wire::WireCodec>::decode(reader).map(|value| Self(value.0))
    }
}

#[derive(Clone, Copy, Debug, Default, PartialEq, Eq, Hash)]
pub struct ActorUniqueID(pub i64);

impl wire::WireCodec for ActorUniqueID {
    fn encode<W: std::io::Write>(&self, writer: &mut W) -> std::io::Result<()> {
        <wire::ZigZag64 as wire::WireCodec>::encode(&wire::ZigZag64(self.0), writer)
    }

    fn decode<R: std::io::Read>(reader: &mut R) -> std::io::Result<Self> {
        <wire::ZigZag64 as wire::WireCodec>::decode(reader).map(|value| Self(value.0))
    }
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct AdventureSettings {
    pub no_pvm: bool,
    pub no_mvp: bool,
    pub immutable_world: bool,
    pub show_name_tags: bool,
    pub auto_jump: bool,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct AnimatedImageData {
    pub skin_image: SkinImage,
    pub animated_texture_type: PersonaAnimatedTextureType,
    pub frames: wire::F32LE,
    pub animation_expression: PersonaAnimationExpression,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct ArmorSlotAndDamagePair {
    pub armor_slot: LegacyArmorSlot,
    pub damage: wire::I16LE,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct AttributeData {
    pub min_value: wire::F32LE,
    pub max_value: wire::F32LE,
    pub current_value: wire::F32LE,
    pub default_min_value: wire::F32LE,
    pub default_max_value: wire::F32LE,
    pub default_value: wire::F32LE,
    pub name: String,
    pub modifiers: Vec<AttributeModifier>,
}

#[derive(Clone, Debug, PartialEq)]
pub enum AttributeLayerSyncPacketData {
    UpdateAttributeLayersData {
        attribute_layers: Vec<EASAttributeLayerData>,
    },
    UpdateAttributeLayerSettingsData {
        attribute_layer_name: String,
        attribute_layer_dimension: DimensionType,
        attributes_layer_settings: EASAttributeLayerSettings,
    },
    UpdateEnvironmentAttributesData {
        attribute_layer_name: String,
        attribute_layer_dimension: DimensionType,
        attributes: Vec<EASEnvironmentAttributeData>,
    },
    RemoveEnvironmentAttributesData {
        attribute_layer_name: String,
        attribute_layer_dimension: DimensionType,
        attributes: Vec<String>,
    },
}

impl AttributeLayerSyncPacketData {
    pub fn discriminant(&self) -> u32 {
        match self {
            Self::UpdateAttributeLayersData { .. } => 0,
            Self::UpdateAttributeLayerSettingsData { .. } => 1,
            Self::UpdateEnvironmentAttributesData { .. } => 2,
            Self::RemoveEnvironmentAttributesData { .. } => 3,
        }
    }
}

impl Default for AttributeLayerSyncPacketData {
    fn default() -> Self {
        Self::UpdateAttributeLayersData {
            attribute_layers: Default::default(),
        }
    }
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct AttributeModifier {
    pub id: String,
    pub name: String,
    pub amount: wire::F32LE,
    pub operation: wire::I32LE,
    pub operand: wire::I32LE,
    pub is_serializable: bool,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct AvailableCommandsChainedSubcommandData {
    pub name: String,
    pub sub_command_values: Vec<AvailableCommandsChainedSubcommandRelationship>,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct AvailableCommandsChainedSubcommandRelationship {
    pub sub_command_first_value: wire::VarUInt,
    pub sub_command_second_value: wire::VarUInt,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct AvailableCommandsConstrainedValueData {
    pub enum_value_symbol: wire::U32LE,
    pub enum_symbol: wire::U32LE,
    pub constraint_indices: Vec<wire::U8>,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct AvailableCommandsEnumData {
    pub name: String,
    pub values: Vec<wire::U32LE>,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct AvailableCommandsOverloadData {
    pub is_chaining: bool,
    pub parameter_data: Vec<AvailableCommandsParamData>,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct AvailableCommandsPacketCommandData {
    pub name: String,
    pub description: String,
    pub flags: wire::U16LE,
    pub permission_level: String,
    pub alias_enum: wire::I32LE,
    pub command_data_chained_subcommand_indexes: Vec<wire::U32LE>,
    pub overloads: Vec<AvailableCommandsOverloadData>,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct AvailableCommandsParamData {
    pub name: String,
    pub parse_symbol: wire::U32LE,
    pub is_optional: bool,
    pub options: wire::U8,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct AvailableCommandsSoftEnumData {
    pub enum_name: String,
    pub enum_options: Vec<String>,
}

#[derive(Clone, Debug, PartialEq)]
pub enum BedrockDDUI {
    DataStoreUpdate(BedrockDDUIDataStoreUpdate),
    DataStoreChange {
        data_store_name: String,
        property: String,
        update_count: wire::U32LE,
        the_new_property_value: CerealDynamicValue,
    },
    DataStoreRemoval {
        data_store_name: String,
    },
}

impl BedrockDDUI {
    pub fn discriminant(&self) -> u32 {
        match self {
            Self::DataStoreUpdate(..) => 0,
            Self::DataStoreChange { .. } => 1,
            Self::DataStoreRemoval { .. } => 2,
        }
    }
}

impl Default for BedrockDDUI {
    fn default() -> Self {
        Self::DataStoreUpdate(Default::default())
    }
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct BedrockDDUIDataStoreUpdate {
    pub data_store_name: String,
    pub property: String,
    pub path: String,
    pub data: BedrockDDUIDataStoreUpdateData,
    pub property_update_count: wire::U32LE,
    pub path_update_count: wire::U32LE,
}

#[derive(Clone, Debug, PartialEq)]
pub enum BedrockDDUIDataStoreUpdateData {
    Double(wire::F64LE),
    Bool(bool),
    String(String),
}

impl BedrockDDUIDataStoreUpdateData {
    pub fn discriminant(&self) -> u32 {
        match self {
            Self::Double(..) => 0,
            Self::Bool(..) => 1,
            Self::String(..) => 2,
        }
    }
}

impl Default for BedrockDDUIDataStoreUpdateData {
    fn default() -> Self {
        Self::Double(Default::default())
    }
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct BedrockProfileWhiskerDiagnosticsScopeDataSummary {
    pub label: String,
    pub indentation: String,
    pub total_high_cost_ns: wire::U64LE,
    pub total_mid_cost_ns: wire::U64LE,
    pub total_low_cost_ns: wire::U64LE,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct BedrockSafetyRedactableString {
    pub unredacted: String,
    pub redacted: String,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct BiomeCappedSurfaceData {
    pub floor_blocks: Vec<wire::U32LE>,
    pub ceiling_blocks: Vec<wire::U32LE>,
    /// Wire presence: optional value is preceded by a presence marker.
    pub sea_block: Option<wire::U32LE>,
    /// Wire presence: optional value is preceded by a presence marker.
    pub foundation_block: Option<wire::U32LE>,
    /// Wire presence: optional value is preceded by a presence marker.
    pub beach_block: Option<wire::U32LE>,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct BiomeClimateData {
    pub temperature: wire::F32LE,
    pub downfall: wire::F32LE,
    pub snow_accumulation_min: wire::F32LE,
    pub snow_accumulation_max: wire::F32LE,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct BiomeConditionalTransformationData {
    pub transforms_into: Vec<BiomeWeightedData>,
    pub condition_json: wire::U16LE,
    pub min_passing_neighbors: wire::U32LE,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct BiomeConsolidatedFeatureData {
    pub scatter: BiomeScatterParamData,
    pub feature: wire::U16LE,
    pub identifier: wire::U16LE,
    pub pass: wire::U16LE,
    pub can_use_internal_feature: bool,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct BiomeConsolidatedFeaturesData {
    pub features: Vec<BiomeConsolidatedFeatureData>,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct BiomeCoordinateData {
    pub min_value_type: wire::ZigZag32,
    pub min_value: wire::U16LE,
    pub max_value_type: wire::ZigZag32,
    pub max_value: wire::U16LE,
    pub grid_offset: wire::U32LE,
    pub grid_step_size: wire::U32LE,
    pub distribution: RandomDistributionType,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct BiomeDefinitionChunkGenData {
    /// Wire presence: optional value is preceded by a presence marker.
    pub climate: Option<BiomeClimateData>,
    /// Wire presence: optional value is preceded by a presence marker.
    pub consolidated_features: Option<BiomeConsolidatedFeaturesData>,
    /// Wire presence: optional value is preceded by a presence marker.
    pub mountain_params: Option<BiomeMountainParamsData>,
    /// Wire presence: optional value is preceded by a presence marker.
    pub surface_material_adjustments: Option<BiomeSurfaceMaterialAdjustmentData>,
    /// Wire presence: optional value is preceded by a presence marker.
    pub overworld_gen_rules: Option<BiomeOverworldGenRulesData>,
    /// Wire presence: optional value is preceded by a presence marker.
    pub multinoise_gen_rules: Option<BiomeMultinoiseGenRulesData>,
    /// Wire presence: optional value is preceded by a presence marker.
    pub legacy_world_gen_rules: Option<BiomeLegacyWorldGenRulesData>,
    /// Wire presence: optional value is preceded by a presence marker.
    pub replacement_biomes: Option<BiomeReplacementsData>,
    /// Wire presence: optional value is preceded by a presence marker.
    pub village_type: Option<VillageType>,
    /// Wire presence: optional value is preceded by a presence marker.
    pub surface_builder_data: Option<BiomeSurfaceBuilderData>,
    /// Wire presence: optional value is preceded by a presence marker.
    pub subsurface_builder_data: Option<BiomeSurfaceBuilderData>,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct BiomeDefinitionData {
    pub id: wire::U16LE,
    pub temperature: wire::F32LE,
    pub downfall: wire::F32LE,
    pub foliage_snow: wire::F32LE,
    pub depth: wire::F32LE,
    pub scale: wire::F32LE,
    pub map_water_color_argb: wire::I32LE,
    pub rain: bool,
    /// Wire presence: optional value is preceded by a presence marker.
    pub tags: Option<BiomeTagsData>,
    /// Wire presence: optional value is preceded by a presence marker.
    pub chunk_gen_data: Option<BiomeDefinitionChunkGenData>,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct BiomeElementData {
    pub noise_freq_scale: wire::F32LE,
    pub noise_lower_bound: wire::F32LE,
    pub noise_upper_bound: wire::F32LE,
    pub height_min_type: wire::ZigZag32,
    pub height_min: wire::U16LE,
    pub height_max_type: wire::ZigZag32,
    pub height_max: wire::U16LE,
    pub adjusted_materials: BiomeSurfaceMaterialData,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct BiomeLegacyWorldGenRulesData {
    pub legacy_pre_hills_edge: Vec<BiomeConditionalTransformationData>,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct BiomeMesaSurfaceData {
    pub clay_material: wire::U32LE,
    pub hard_clay_material: wire::U32LE,
    pub bryce_pillars: bool,
    pub has_forest: bool,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct BiomeMountainParamsData {
    pub steep_block: wire::U32LE,
    pub north_slopes: bool,
    pub south_slopes: bool,
    pub west_slopes: bool,
    pub east_slopes: bool,
    pub top_slide_enabled: bool,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct BiomeMultinoiseGenRulesData {
    pub temperature: wire::F32LE,
    pub humidity: wire::F32LE,
    pub altitude: wire::F32LE,
    pub weirdness: wire::F32LE,
    pub weight: wire::F32LE,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct BiomeNoiseGradientSurfaceData {
    pub non_replaceable_blocks: Vec<wire::U32LE>,
    pub gradient_blocks: Vec<SerializedNoiseBlockSpecifier>,
    pub noise: NoiseDescriptor,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct BiomeOverworldGenRulesData {
    pub hills_transformations: Vec<BiomeWeightedData>,
    pub mutate_transformations: Vec<BiomeWeightedData>,
    pub river_transformations: Vec<BiomeWeightedData>,
    pub shore_transformations: Vec<BiomeWeightedData>,
    pub pre_hills_edge: Vec<BiomeConditionalTransformationData>,
    pub post_shore_edge: Vec<BiomeConditionalTransformationData>,
    pub climate: Vec<BiomeWeightedTemperatureData>,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct BiomeReplacementData {
    pub replacement_biome: wire::U16LE,
    pub dimension: wire::U16LE,
    pub target_biomes: Vec<wire::U16LE>,
    pub amount: wire::F32LE,
    pub noise_frequency_scale: wire::F32LE,
    pub replacement_index: wire::U32LE,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct BiomeReplacementsData {
    pub biome_replacements: Vec<BiomeReplacementData>,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct BiomeScatterParamData {
    pub coordinates: Vec<BiomeCoordinateData>,
    pub eval_order: CoordinateEvaluationOrder,
    pub chance_percent_type: wire::ZigZag32,
    pub chance_percent: wire::U16LE,
    pub chance_numerator: wire::I32LE,
    pub chance_denominator: wire::I32LE,
    pub iterations_type: wire::ZigZag32,
    pub iterations: wire::U16LE,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct BiomeStringList {
    pub strings: Vec<String>,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct BiomeSurfaceBuilderData {
    /// Wire presence: optional value is preceded by a presence marker.
    pub surface_materials: Option<BiomeSurfaceMaterialData>,
    pub has_default_overworld_surface: bool,
    pub has_swamp_surface: bool,
    pub has_frozen_ocean_surface: bool,
    pub has_the_end_surface: bool,
    /// Wire presence: optional value is preceded by a presence marker.
    pub mesa_surface: Option<BiomeMesaSurfaceData>,
    /// Wire presence: optional value is preceded by a presence marker.
    pub capped_surface: Option<BiomeCappedSurfaceData>,
    /// Wire presence: optional value is preceded by a presence marker.
    pub noise_gradient_surface: Option<BiomeNoiseGradientSurfaceData>,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct BiomeSurfaceMaterialAdjustmentData {
    pub adjustments: Vec<BiomeElementData>,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct BiomeSurfaceMaterialData {
    pub top_block: wire::U32LE,
    pub mid_block: wire::U32LE,
    pub sea_floor_block: wire::U32LE,
    pub foundation_block: wire::U32LE,
    pub sea_block: wire::U32LE,
    pub sea_floor_depth: wire::I32LE,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct BiomeTagsData {
    pub tags: Vec<wire::U16LE>,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct BiomeWeightedData {
    pub biome_identifier: wire::U16LE,
    pub weight: wire::U32LE,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct BiomeWeightedTemperatureData {
    pub temperature: wire::ZigZag32,
    pub weight: wire::U32LE,
}

/// Stores the 131-bit value used by the wire bitset encoding.
#[derive(Clone, Debug, Default, PartialEq, Eq)]
pub struct Bitset131(pub [u64; 3]);

#[derive(Clone, Debug, Default, PartialEq)]
pub struct BlockPos {
    pub x: wire::ZigZag32,
    pub y: wire::ZigZag32,
    pub z: wire::ZigZag32,
}

#[derive(Clone, Debug, PartialEq)]
pub enum BookEditAction {
    ReplacePage {
        page_index: wire::ZigZag32,
        page_text: String,
        photo_name: String,
    },
    AddPage {
        page_index: wire::ZigZag32,
        page_text: String,
        photo_name: String,
    },
    DeletePage {
        page_index: wire::ZigZag32,
    },
    SwapPages {
        page_index: wire::ZigZag32,
        swap_with_index: wire::ZigZag32,
    },
    Finalize {
        title: String,
        author: String,
        xuid: String,
    },
}

impl BookEditAction {
    pub fn discriminant(&self) -> u32 {
        match self {
            Self::ReplacePage { .. } => 0,
            Self::AddPage { .. } => 1,
            Self::DeletePage { .. } => 2,
            Self::SwapPages { .. } => 3,
            Self::Finalize { .. } => 4,
        }
    }
}

impl Default for BookEditAction {
    fn default() -> Self {
        Self::ReplacePage {
            page_index: Default::default(),
            page_text: Default::default(),
            photo_name: Default::default(),
        }
    }
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct CameraAimAssistActorPriorityPriorityData {
    pub preset_index: wire::I32LE,
    pub category_index: wire::I32LE,
    pub actor_index: wire::I32LE,
    pub priority_value: wire::I32LE,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct CameraAimAssistCategoryDefinition {
    pub name: String,
    pub priorities: CameraAimAssistCategoryPriorities,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct CameraAimAssistCategoryPriorities {
    pub entities: Vec<(String, wire::I32LE)>,
    pub blocks: Vec<(String, wire::I32LE)>,
    pub block_tags: Vec<(String, wire::I32LE)>,
    pub entity_type_families: Vec<(String, wire::I32LE)>,
    /// Wire presence: optional value is preceded by a presence marker.
    pub entity_default: Option<wire::I32LE>,
    /// Wire presence: optional value is preceded by a presence marker.
    pub block_default: Option<wire::I32LE>,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct CameraAimAssistCommandPresetDefinition {
    /// Wire presence: optional value is preceded by a presence marker.
    pub preset_id: Option<String>,
    /// Wire presence: optional value is preceded by a presence marker.
    pub target_mode: Option<CameraAimAssistTargetMode>,
    /// Wire presence: optional value is preceded by a presence marker.
    pub view_angle: Option<glam::Vec2>,
    /// Wire presence: optional value is preceded by a presence marker.
    pub distance: Option<wire::F32LE>,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct CameraAimAssistPresetDefinition {
    pub identifier: String,
    pub exclusion_settings: CameraAimAssistPresetExclusionDefinition,
    pub liquid_targeting_list: Vec<String>,
    pub item_settings: Vec<(String, String)>,
    /// Wire presence: optional value is preceded by a presence marker.
    pub default_item_settings: Option<String>,
    /// Wire presence: optional value is preceded by a presence marker.
    pub hand_settings: Option<String>,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct CameraAimAssistPresetExclusionDefinition {
    pub blocks: Vec<String>,
    pub entities: Vec<String>,
    pub block_tags: Vec<String>,
    pub entity_type_families: Vec<String>,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct CameraInstructionData {
    /// Wire presence: optional value is preceded by a presence marker.
    pub set: Option<CameraInstructionOptionsSetInstruction>,
    /// Wire presence: optional value is preceded by a presence marker.
    pub clear: Option<bool>,
    /// Wire presence: optional value is preceded by a presence marker.
    pub fade: Option<CameraInstructionOptionsFadeInstruction>,
    /// Wire presence: optional value is preceded by a presence marker.
    pub target: Option<CameraInstructionOptionsTargetInstruction>,
    /// Wire presence: optional value is preceded by a presence marker.
    pub remove_target: Option<bool>,
    /// Wire presence: optional value is preceded by a presence marker.
    pub field_of_view: Option<CameraInstructionOptionsFovInstruction>,
    /// Wire presence: optional value is preceded by a presence marker.
    pub spline: Option<CameraInstructionOptionsSplineInstruction>,
    /// Wire presence: optional value is preceded by a presence marker.
    pub attach_to_entity: Option<CameraInstructionOptionsAttachToEntityInstruction>,
    /// Wire presence: optional value is preceded by a presence marker.
    pub detach_from_entity: Option<bool>,
}

#[derive(Clone, Copy, Debug, Default, PartialEq, Eq, Hash)]
pub struct CameraInstructionOptionsAttachToEntityInstruction(pub i64);

impl wire::WireCodec for CameraInstructionOptionsAttachToEntityInstruction {
    fn encode<W: std::io::Write>(&self, writer: &mut W) -> std::io::Result<()> {
        <wire::I64LE as wire::WireCodec>::encode(&wire::I64LE(self.0), writer)
    }

    fn decode<R: std::io::Read>(reader: &mut R) -> std::io::Result<Self> {
        <wire::I64LE as wire::WireCodec>::decode(reader).map(|value| Self(value.0))
    }
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct CameraInstructionOptionsFadeInstruction {
    /// Wire presence: optional value is preceded by a presence marker.
    pub time: Option<CameraInstructionOptionsFadeInstructionTimeOption>,
    /// Wire presence: optional value is preceded by a presence marker.
    pub color: Option<CameraInstructionOptionsFadeInstructionColorOption>,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct CameraInstructionOptionsFadeInstructionColorOption {
    pub red: wire::F32LE,
    pub green: wire::F32LE,
    pub blue: wire::F32LE,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct CameraInstructionOptionsFadeInstructionTimeOption {
    pub fade_in_time: wire::F32LE,
    pub hold_time: wire::F32LE,
    pub fade_out_time: wire::F32LE,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct CameraInstructionOptionsFovInstruction {
    pub field_of_view: wire::F32LE,
    pub fov_ease_time: wire::F32LE,
    pub fov_ease_type: String,
    pub field_of_view_clear: bool,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct CameraInstructionOptionsSetInstruction {
    pub preset: wire::U32LE,
    /// Wire presence: optional value is preceded by a presence marker.
    pub ease: Option<CameraInstructionOptionsSetInstructionEaseOption>,
    /// Wire presence: optional value is preceded by a presence marker.
    pub pos: Option<CameraInstructionOptionsSetInstructionPosOption>,
    /// Wire presence: optional value is preceded by a presence marker.
    pub rot: Option<CameraInstructionOptionsSetInstructionRotOption>,
    /// Wire presence: optional value is preceded by a presence marker.
    pub facing: Option<CameraInstructionOptionsSetInstructionFacingOption>,
    /// Wire presence: optional value is preceded by a presence marker.
    pub view_offset: Option<CameraInstructionOptionsSetInstructionViewOffsetOption>,
    /// Wire presence: optional value is preceded by a presence marker.
    pub entity_offset: Option<CameraInstructionOptionsSetInstructionEntityOffsetOption>,
    /// Wire presence: optional value is preceded by a presence marker.
    pub default: Option<bool>,
    pub remove_ignore_starting_values_component: bool,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct CameraInstructionOptionsSetInstructionEaseOption {
    pub type_: wire::U8,
    pub time: wire::F32LE,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct CameraInstructionOptionsSetInstructionEntityOffsetOption {
    pub entity_offset_x: wire::F32LE,
    pub entity_offset_y: wire::F32LE,
    pub entity_offset_z: wire::F32LE,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct CameraInstructionOptionsSetInstructionFacingOption {
    pub pos: glam::Vec3,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct CameraInstructionOptionsSetInstructionPosOption {
    pub pos: glam::Vec3,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct CameraInstructionOptionsSetInstructionRotOption {
    pub x: wire::F32LE,
    pub y: wire::F32LE,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct CameraInstructionOptionsSetInstructionViewOffsetOption {
    pub x: wire::F32LE,
    pub y: wire::F32LE,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct CameraInstructionOptionsSplineInstruction {
    pub total_time: wire::F32LE,
    pub type_: wire::U8,
    pub curve: Vec<glam::Vec3>,
    pub progress_key_frames: Vec<CameraInstructionOptionsSplineInstructionSplineProgressOption>,
    pub rotation_option: Vec<CameraInstructionOptionsSplineInstructionSplineRotationOption>,
    pub spline_identifier: String,
    pub load_from_json: bool,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct CameraInstructionOptionsSplineInstructionSplineProgressOption {
    pub key_frame_value: wire::F32LE,
    pub key_frame_time: wire::F32LE,
    pub key_frame_easing_func: String,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct CameraInstructionOptionsSplineInstructionSplineRotationOption {
    pub key_frame_value: glam::Vec3,
    pub key_frame_time: wire::F32LE,
    pub key_frame_easing_func: String,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct CameraInstructionOptionsTargetInstruction {
    /// Wire presence: optional value is preceded by a presence marker.
    pub target_center_offset: Option<glam::Vec3>,
    pub target_actor_id: wire::I64LE,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct CameraPreset {
    pub name: String,
    pub inherit_from: String,
    /// Wire presence: optional value is preceded by a presence marker.
    pub pos_x: Option<wire::F32LE>,
    /// Wire presence: optional value is preceded by a presence marker.
    pub pos_y: Option<wire::F32LE>,
    /// Wire presence: optional value is preceded by a presence marker.
    pub pos_z: Option<wire::F32LE>,
    /// Wire presence: optional value is preceded by a presence marker.
    pub rot_x: Option<wire::F32LE>,
    /// Wire presence: optional value is preceded by a presence marker.
    pub rot_y: Option<wire::F32LE>,
    /// Wire presence: optional value is preceded by a presence marker.
    pub rotation_speed: Option<wire::F32LE>,
    /// Wire presence: optional value is preceded by a presence marker.
    pub snap_to_target: Option<bool>,
    /// Wire presence: optional value is preceded by a presence marker.
    pub horizontal_rotation_limit: Option<glam::Vec2>,
    /// Wire presence: optional value is preceded by a presence marker.
    pub vertical_rotation_limit: Option<glam::Vec2>,
    /// Wire presence: optional value is preceded by a presence marker.
    pub continue_targeting: Option<bool>,
    /// Wire presence: optional value is preceded by a presence marker.
    pub block_listening_radius: Option<wire::F32LE>,
    /// Wire presence: optional value is preceded by a presence marker.
    pub view_offset: Option<glam::Vec2>,
    /// Wire presence: optional value is preceded by a presence marker.
    pub entity_offset: Option<glam::Vec3>,
    /// Wire presence: optional value is preceded by a presence marker.
    pub radius: Option<wire::F32LE>,
    /// Wire presence: optional value is preceded by a presence marker.
    pub yaw_limit_min: Option<wire::F32LE>,
    /// Wire presence: optional value is preceded by a presence marker.
    pub yaw_limit_max: Option<wire::F32LE>,
    /// Wire presence: optional value is preceded by a presence marker.
    pub listener: Option<CameraPresetAudioListener>,
    /// Wire presence: optional value is preceded by a presence marker.
    pub player_effects: Option<bool>,
    /// Wire presence: optional value is preceded by a presence marker.
    pub aim_assist: Option<CameraAimAssistCommandPresetDefinition>,
    /// Wire presence: optional value is preceded by a presence marker.
    pub control_scheme: Option<ControlSchemeScheme>,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct CameraPresetsData {
    pub presets: Vec<CameraPreset>,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct CameraSplineControlPoint {
    pub position: glam::Vec3,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct CameraSplineDefinition {
    pub name: String,
    pub total_time: wire::F32LE,
    pub spline_type: String,
    pub control_points: Vec<CameraSplineControlPoint>,
    pub progress_key_frames: Vec<CameraSplineProgressKeyFrame>,
    pub rotation_key_frames: Vec<CameraSplineRotationKeyFrame>,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct CameraSplineProgressKeyFrame {
    pub progress: wire::F32LE,
    pub time: wire::F32LE,
    /// Wire presence: optional value is preceded by a presence marker.
    pub easing: Option<String>,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct CameraSplineRotationKeyFrame {
    pub rotation: glam::Vec3,
    pub time: wire::F32LE,
    /// Wire presence: optional value is preceded by a presence marker.
    pub easing: Option<String>,
}

#[derive(Clone, Debug, PartialEq, Default)]
pub enum CerealDynamicValue {
    #[default]
    None,
    Bool(bool),
    Int64(wire::I64LE),
    Double(wire::F64LE),
    String(String),
    List(Vec<CerealDynamicValue>),
    Map(Vec<(String, CerealDynamicValue)>),
}

impl CerealDynamicValue {
    pub fn discriminant(&self) -> i32 {
        match self {
            Self::None => 0,
            Self::Bool(..) => 1,
            Self::Int64(..) => 2,
            Self::Double(..) => 3,
            Self::String(..) => 4,
            Self::List(..) => 5,
            Self::Map(..) => 6,
        }
    }
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct CerealizerExperimentsAnonExperimentToggle {
    pub name: String,
    pub enabled: bool,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct CerealizerNetworkItemInstanceDescriptorSerializedData {
    pub id: wire::ZigZag32,
    pub stack_size: wire::U16LE,
    pub aux_value: wire::VarUInt,
    pub block_runtime_id: wire::ZigZag32,
    pub user_data_buffer: bytes::Bytes,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct CerealizerNetworkItemStackDescriptorSerializedData {
    pub id: wire::I16LE,
    pub stack_size: wire::U16LE,
    pub aux_value: wire::VarUInt,
    /// Wire presence: optional value is preceded by a presence marker.
    pub net_id_variant: Option<wire::ZigZag32>,
    pub block_runtime_id: wire::VarUInt,
    pub user_data_buffer: bytes::Bytes,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct CerealizerRecipeIngredientSerializedData {
    pub descriptor: Vec<(String, String)>,
    pub aux_value: wire::ZigZag32,
    pub stack_size: wire::ZigZag32,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct CerealizerRecipeUnlockingRequirementSerializedData {
    pub unlocking_context: RecipeUnlockingRequirementUnlockingContext,
    /// Wire presence: optional value is preceded by a presence marker.
    pub unlocking_ingredients: Option<Vec<CerealizerRecipeIngredientSerializedData>>,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct ChunkPos {
    pub x: wire::ZigZag32,
    pub z: wire::ZigZag32,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct ClientboundDebugRendererDebugMarkerData {
    pub text: String,
    pub position: glam::Vec3,
    pub color: MceColor,
    pub duration: wire::U64LE,
}

#[derive(Clone, Debug, PartialEq)]
pub enum CommandBlockUpdateTarget {
    EntityCommandTarget {
        target_runtime_id: ActorRuntimeID,
    },
    BlockCommandData {
        block_position: BlockPos,
        command_block_mode: wire::VarUInt,
        redstone_mode: bool,
        is_conditional: bool,
    },
}

impl CommandBlockUpdateTarget {
    pub fn discriminant(&self) -> u32 {
        match self {
            Self::EntityCommandTarget { .. } => 0,
            Self::BlockCommandData { .. } => 1,
        }
    }
}

impl Default for CommandBlockUpdateTarget {
    fn default() -> Self {
        Self::EntityCommandTarget {
            target_runtime_id: Default::default(),
        }
    }
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct CommandOriginData {
    pub type_: String,
    pub uuid: uuid::Uuid,
    pub request_id: String,
    pub player_id: wire::I64LE,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct CommandOutputData {
    pub output_type: String,
    pub success_count: wire::U32LE,
    pub output_messages: Vec<CommandOutputMessage>,
    /// Wire presence: optional value is preceded by a presence marker.
    pub data_set: Option<String>,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct CommandOutputMessage {
    pub message_id: String,
    pub successful: bool,
    pub parameters: Vec<String>,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct ContainerMixDataEntry {
    pub from_item_id: wire::ZigZag32,
    pub reagent_item_id: wire::ZigZag32,
    pub to_item_id: wire::ZigZag32,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct ContentIdentity {
    pub identity: String,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct CreativeGroupInfo {
    pub creative_category: CreativeItemCategory,
    pub name: String,
    pub group_icon_item: CerealizerNetworkItemInstanceDescriptorSerializedData,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct CreativeItemEntry {
    pub creative_net_id: TypedServerNetIdStructCreativeItemNetIdTag,
    pub item_instance: CerealizerNetworkItemInstanceDescriptorSerializedData,
    pub group_index: wire::VarUInt,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct DataItemEntry {
    pub id: wire::VarUInt,
    pub payload: DataItemEntryValue,
}

#[derive(Clone, Debug, PartialEq)]
pub enum DataItemEntryValue {
    DataItemByte {
        value: wire::I8,
    },
    DataItemShort {
        value: wire::I16LE,
    },
    DataItemInt {
        value: wire::ZigZag32,
    },
    DataItemFloat {
        value: wire::F32LE,
    },
    DataItemString {
        value: String,
    },
    DataItemCompoundTag {
        value: Nbt,
    },
    DataItemPos {
        value: BlockPos,
    },
    DataItemInt64 {
        value: wire::ZigZag64,
    },
    DataItemVec3 {
        value: glam::Vec3,
    },
}

impl DataItemEntryValue {
    pub fn discriminant(&self) -> u8 {
        match self {
            Self::DataItemByte { .. } => 0,
            Self::DataItemShort { .. } => 1,
            Self::DataItemInt { .. } => 2,
            Self::DataItemFloat { .. } => 3,
            Self::DataItemString { .. } => 4,
            Self::DataItemCompoundTag { .. } => 5,
            Self::DataItemPos { .. } => 6,
            Self::DataItemInt64 { .. } => 7,
            Self::DataItemVec3 { .. } => 8,
        }
    }
}

impl Default for DataItemEntryValue {
    fn default() -> Self {
        Self::DataItemByte {
            value: Default::default(),
        }
    }
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct DimensionDefinition {
    pub height_maximum: wire::ZigZag32,
    pub height_minimum: wire::ZigZag32,
    pub generator_type: GeneratorType,
    pub dimension_type: DimensionType,
    pub pack_id: uuid::Uuid,
}

#[derive(Clone, Copy, Debug, Default, PartialEq, Eq, Hash)]
pub struct DimensionType(pub i32);

impl wire::WireCodec for DimensionType {
    fn encode<W: std::io::Write>(&self, writer: &mut W) -> std::io::Result<()> {
        <wire::ZigZag32 as wire::WireCodec>::encode(&wire::ZigZag32(self.0), writer)
    }

    fn decode<R: std::io::Read>(reader: &mut R) -> std::io::Result<Self> {
        <wire::ZigZag32 as wire::WireCodec>::decode(reader).map(|value| Self(value.0))
    }
}

#[derive(Clone, Debug, PartialEq)]
pub enum DisconnectMessages {
    DisconnectPacketMessages {
        message: String,
        filtered_message: String,
    },
    /// Naming overlay required: source placeholder `Empty1`.
    Empty1,
}

impl DisconnectMessages {
    pub fn discriminant(&self) -> u32 {
        match self {
            Self::DisconnectPacketMessages { .. } => 0,
            Self::Empty1 => 1,
        }
    }
}

impl Default for DisconnectMessages {
    fn default() -> Self {
        Self::DisconnectPacketMessages {
            message: Default::default(),
            filtered_message: Default::default(),
        }
    }
}

#[derive(Clone, Debug, PartialEq)]
pub enum EAS {
    BoolAttributeData {
        value: bool,
        operation: String,
    },
    FloatAttributeData {
        value: wire::F32LE,
        operation: String,
        /// Wire presence: optional value is preceded by a presence marker.
        constraint_min: Option<wire::F32LE>,
        /// Wire presence: optional value is preceded by a presence marker.
        constraint_max: Option<wire::F32LE>,
    },
    ColorAttributeData {
        value: [wire::I32LE; 4],
        operation: String,
    },
}

impl EAS {
    pub fn discriminant(&self) -> u32 {
        match self {
            Self::BoolAttributeData { .. } => 0,
            Self::FloatAttributeData { .. } => 1,
            Self::ColorAttributeData { .. } => 2,
        }
    }
}

impl Default for EAS {
    fn default() -> Self {
        Self::BoolAttributeData {
            value: Default::default(),
            operation: Default::default(),
        }
    }
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct EASAttributeLayerData {
    pub name: String,
    /// Wire presence: optional value is preceded by a presence marker.
    pub noise_name: Option<String>,
    pub dimension: DimensionType,
    pub settings: EASAttributeLayerSettings,
    pub attributes: Vec<EASEnvironmentAttributeData>,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct EASAttributeLayerSettings {
    pub priority: wire::I32LE,
    pub weight: wire::F32LE,
    pub enabled: bool,
    pub transitions_paused: bool,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct EASEnvironmentAttributeData {
    pub attribute_name: String,
    /// Wire presence: optional value is preceded by a presence marker.
    pub from_attribute: Option<EAS>,
    pub attribute: EAS,
    /// Wire presence: optional value is preceded by a presence marker.
    pub to_attribute: Option<EAS>,
    pub current_transition_ticks: wire::U32LE,
    pub total_transition_ticks: wire::U32LE,
    pub easing: String,
    pub local_transition_ticks: wire::U32LE,
    pub noise_transition: bool,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct ECSProfilingDiagnosticsEntityDiagnosticTimingInfo {
    pub display_name: String,
    pub entity: String,
    pub time_in_ns: wire::U64LE,
    pub percent_of_total: wire::U8,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct ECSProfilingDiagnosticsSystemCategory {
    pub category_name: String,
    pub system_index: wire::U64LE,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct ECSProfilingDiagnosticsSystemDiagnosticTimingInfo {
    pub display_name: String,
    pub system_index: wire::U64LE,
    pub time_in_ns: wire::U64LE,
    pub percent_of_total: wire::U8,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct EduSharedUriResource {
    pub button_name: String,
    pub link_uri: String,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct EducationLevelSettings {
    pub code_builder_default_uri: String,
    pub code_builder_title: String,
    pub can_resize_code_builder: bool,
    pub disable_legacy_title_bar: bool,
    pub post_process_filter: String,
    pub screenshot_border_resource_path: String,
    /// Wire presence: optional value is preceded by a presence marker.
    pub agent_capabilities: Option<bool>,
    pub local_settings: EducationLocalLevelSettings,
    pub deprecated_always_false: bool,
    /// Wire presence: optional value is preceded by a presence marker.
    pub external_link_settings: Option<ExternalLinkSettings>,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct EducationLocalLevelSettings {
    /// Wire presence: optional value is preceded by a presence marker.
    pub code_builder_override_uri: Option<String>,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct EnchantmentInstance {
    pub enchant_type: EnchantType,
    pub enchant_level: wire::U8,
}

#[derive(Clone, Copy, Debug, Default, PartialEq, Eq, Hash)]
pub struct EntityNetId(pub u32);

impl wire::WireCodec for EntityNetId {
    fn encode<W: std::io::Write>(&self, writer: &mut W) -> std::io::Result<()> {
        <wire::VarUInt as wire::WireCodec>::encode(&wire::VarUInt(self.0), writer)
    }

    fn decode<R: std::io::Read>(reader: &mut R) -> std::io::Result<Self> {
        <wire::VarUInt as wire::WireCodec>::decode(reader).map(|value| Self(value.0))
    }
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct Experiments {
    pub toggles: Vec<CerealizerExperimentsAnonExperimentToggle>,
    pub experiments_ever_toggled: bool,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct ExternalLinkSettings {
    pub url: String,
    pub display_name: String,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct FeatureRegistryFeatureBinaryJsonFormat {
    pub feature_name: String,
    pub binary_json_output: String,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct FloatRange {
    pub min: wire::F32LE,
    pub max: wire::F32LE,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct FullContainerName {
    pub container_name: ContainerEnumName,
    /// Wire presence: optional value is preceded by a presence marker.
    pub dynamic_id: Option<wire::U32LE>,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct GameRule {
    pub rule_name: String,
    pub rule_can_be_modified: bool,
    pub rule_value: GameRuleRuleValue,
}

#[derive(Clone, Debug, PartialEq, Default)]
pub enum GameRuleRuleValue {
    /// Naming overlay required: source placeholder `Empty0`.
    #[default]
    Empty0,
    Bool(bool),
    Int32(wire::I32LE),
    Float(wire::F32LE),
}

impl GameRuleRuleValue {
    pub fn discriminant(&self) -> u32 {
        match self {
            Self::Empty0 => 0,
            Self::Bool(..) => 1,
            Self::Int32(..) => 2,
            Self::Float(..) => 3,
        }
    }
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct GameRulesChangedPacketData {
    pub rules_list: Vec<GameRule>,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct InventoryAction {
    pub source: InventorySource,
    pub slot: wire::VarUInt,
    pub from_item: CerealizerNetworkItemStackDescriptorSerializedData,
    pub to_item: CerealizerNetworkItemStackDescriptorSerializedData,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct InventoryOptions {
    pub left_inventory_tab: InventoryLeftTabIndex,
    pub right_inventory_tab: InventoryRightTabIndex,
    pub filtering: bool,
    pub layout_inv: InventoryLayout,
    pub layout_craft: InventoryLayout,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct InventorySource {
    pub source_type: InventorySourceType,
    /// Wire presence: optional value is preceded by a presence marker.
    pub container_id: Option<wire::I8>,
    /// Wire presence: optional value is preceded by a presence marker.
    pub bit_flags: Option<InventorySourceInventorySourceFlags>,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct InventoryTransactionData {
    /// Wire presence: optional value is preceded by a presence marker.
    pub actions: Option<Vec<InventoryAction>>,
}

#[derive(Clone, Debug, PartialEq)]
pub enum InventoryTransactionTransactionValue {
    NormalTransactionData {
        actions: InventoryTransactionData,
    },
    InventoryMismatchData {
        actions: InventoryTransactionData,
    },
    ItemUseInventoryTransaction(Box<ItemUseInventoryTransaction>),
    ItemUseOnActorInventoryTransaction {
        actions: InventoryTransactionData,
        runtime_id: ActorRuntimeID,
        action_type: ItemUseOnActorInventoryTransactionActionType,
        slot: wire::ZigZag32,
        item: CerealizerNetworkItemStackDescriptorSerializedData,
        from_position: glam::Vec3,
        hit_position: glam::Vec3,
    },
    ItemReleaseInventoryTransaction {
        actions: InventoryTransactionData,
        action_type: ItemReleaseInventoryTransactionActionType,
        slot: wire::ZigZag32,
        item: CerealizerNetworkItemStackDescriptorSerializedData,
        from_position: glam::Vec3,
    },
}

impl InventoryTransactionTransactionValue {
    pub fn discriminant(&self) -> u32 {
        match self {
            Self::NormalTransactionData { .. } => 0,
            Self::InventoryMismatchData { .. } => 1,
            Self::ItemUseInventoryTransaction(..) => 2,
            Self::ItemUseOnActorInventoryTransaction { .. } => 3,
            Self::ItemReleaseInventoryTransaction { .. } => 4,
        }
    }
}

impl Default for InventoryTransactionTransactionValue {
    fn default() -> Self {
        Self::NormalTransactionData {
            actions: Default::default(),
        }
    }
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct ItemData {
    pub item_name: String,
    pub item_id: wire::I16LE,
    pub is_component_based: bool,
    pub item_version: ItemVersion,
    pub item_component_data: Nbt,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct ItemEnchantOption {
    pub cost: wire::U8,
    pub enchants: ItemEnchants,
    pub enchant_name: String,
    pub enchant_net_id: TypedServerNetIdStructRecipeNetIdTag,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct ItemEnchants {
    pub slot: wire::I32LE,
    pub item_enchants: [Vec<EnchantmentInstance>; 3],
}

#[derive(Clone, Debug, PartialEq)]
pub enum ItemStackRequestCereal {
    TakeActionData {
        action_type: ItemStackRequestActionType,
        amount: wire::U8,
        source: ItemStackRequestCerealSlotInfoData,
        destination: ItemStackRequestCerealSlotInfoData,
    },
    PlaceActionData {
        action_type: ItemStackRequestActionType,
        amount: wire::U8,
        source: ItemStackRequestCerealSlotInfoData,
        destination: ItemStackRequestCerealSlotInfoData,
    },
    SwapActionData {
        action_type: ItemStackRequestActionType,
        source: ItemStackRequestCerealSlotInfoData,
        destination: ItemStackRequestCerealSlotInfoData,
    },
    DropActionData {
        action_type: ItemStackRequestActionType,
        amount: wire::U8,
        source: ItemStackRequestCerealSlotInfoData,
        randomly: bool,
    },
    DestroyActionData {
        action_type: ItemStackRequestActionType,
        amount: wire::U8,
        source: ItemStackRequestCerealSlotInfoData,
    },
    ConsumeActionData {
        action_type: ItemStackRequestActionType,
        amount: wire::U8,
        source: ItemStackRequestCerealSlotInfoData,
    },
    CreateActionData {
        action_type: ItemStackRequestActionType,
        results_index: wire::U8,
    },
    LabTableCombineActionData {
        action_type: ItemStackRequestActionType,
    },
    BeaconPaymentActionData {
        action_type: ItemStackRequestActionType,
        primary_effect_id: wire::ZigZag32,
        secondary_effect_id: wire::ZigZag32,
    },
    MineBlockActionData {
        action_type: ItemStackRequestActionType,
        slot: wire::ZigZag32,
        predicted_durability: wire::ZigZag32,
        net_id_variant: wire::I32LE,
    },
    CraftRecipeActionData {
        action_type: ItemStackRequestActionType,
        recipe_net_id: TypedServerNetIdStructRecipeNetIdTag,
        number_of_requested_crafts: wire::U8,
    },
    CraftRecipeAutoActionData {
        action_type: ItemStackRequestActionType,
        recipe_net_id: TypedServerNetIdStructRecipeNetIdTag,
        number_of_requested_crafts: wire::U8,
        ingredients: Vec<ItemStackRequestCerealRecipeIngredientData>,
    },
    CraftCreativeActionData {
        action_type: ItemStackRequestActionType,
        creative_item_net_id: wire::VarUInt,
        number_of_requested_crafts: wire::U8,
    },
    CraftRecipeOptionalActionData {
        action_type: ItemStackRequestActionType,
        recipe_net_id: TypedServerNetIdStructRecipeNetIdTag,
        filtered_string_index: wire::I32LE,
    },
    CraftRepairAndDisenchantActionData {
        action_type: ItemStackRequestActionType,
        recipe_net_id: wire::I32LE,
        number_of_requested_crafts: wire::U8,
        repair_cost: wire::ZigZag32,
    },
    CraftLoomActionData {
        action_type: ItemStackRequestActionType,
        pattern_name_id: String,
        num_crafts: wire::U8,
    },
    CraftNonImplementedActionData {
        action_type: ItemStackRequestActionType,
    },
    CraftResultsActionData {
        action_type: ItemStackRequestActionType,
        craft_results: Vec<ItemStackRequestCerealNetworkItemInstanceDescriptorData>,
        num_crafts: wire::U8,
    },
}

impl ItemStackRequestCereal {
    pub fn discriminant(&self) -> u32 {
        match self {
            Self::TakeActionData { .. } => 0,
            Self::PlaceActionData { .. } => 1,
            Self::SwapActionData { .. } => 2,
            Self::DropActionData { .. } => 3,
            Self::DestroyActionData { .. } => 4,
            Self::ConsumeActionData { .. } => 5,
            Self::CreateActionData { .. } => 6,
            Self::LabTableCombineActionData { .. } => 7,
            Self::BeaconPaymentActionData { .. } => 8,
            Self::MineBlockActionData { .. } => 9,
            Self::CraftRecipeActionData { .. } => 10,
            Self::CraftRecipeAutoActionData { .. } => 11,
            Self::CraftCreativeActionData { .. } => 12,
            Self::CraftRecipeOptionalActionData { .. } => 13,
            Self::CraftRepairAndDisenchantActionData { .. } => 14,
            Self::CraftLoomActionData { .. } => 15,
            Self::CraftNonImplementedActionData { .. } => 16,
            Self::CraftResultsActionData { .. } => 17,
        }
    }
}

impl Default for ItemStackRequestCereal {
    fn default() -> Self {
        Self::TakeActionData {
            action_type: Default::default(),
            amount: Default::default(),
            source: Default::default(),
            destination: Default::default(),
        }
    }
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct ItemStackRequestCerealNetworkItemInstanceDescriptorData {
    pub item_descriptor: ItemStackRequestCerealRecipeIngredientDataItemDescriptor,
    pub stack_size: wire::U16LE,
    pub block_runtime_id: wire::VarUInt,
    pub user_data_buffer: bytes::Bytes,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct ItemStackRequestCerealRecipeIngredientData {
    pub item_descriptor: ItemStackRequestCerealRecipeIngredientDataItemDescriptor,
    pub stack_size: wire::U16LE,
}

#[derive(Clone, Debug, PartialEq)]
pub enum ItemStackRequestCerealRecipeIngredientDataItemDescriptor {
    EmptyItemDescriptorData {
        descriptor_type: ItemStackRequestCerealItemDescriptorType,
    },
    ItemNameDescriptorData {
        descriptor_type: ItemStackRequestCerealItemDescriptorType,
        full_name: String,
        aux_value: wire::ZigZag32,
    },
    MolangItemDescriptorData {
        descriptor_type: ItemStackRequestCerealItemDescriptorType,
        tag_expression: String,
        molang_version: MoLangVersion,
    },
    ItemTagDescriptorData {
        descriptor_type: ItemStackRequestCerealItemDescriptorType,
        item_tag: String,
    },
}

impl ItemStackRequestCerealRecipeIngredientDataItemDescriptor {
    pub fn discriminant(&self) -> u32 {
        match self {
            Self::EmptyItemDescriptorData { .. } => 0,
            Self::ItemNameDescriptorData { .. } => 1,
            Self::MolangItemDescriptorData { .. } => 2,
            Self::ItemTagDescriptorData { .. } => 3,
        }
    }
}

impl Default for ItemStackRequestCerealRecipeIngredientDataItemDescriptor {
    fn default() -> Self {
        Self::EmptyItemDescriptorData {
            descriptor_type: Default::default(),
        }
    }
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct ItemStackRequestCerealRequestData {
    pub client_request_id: TypedClientNetIdStructItemStackRequestIdTagInt32T0,
    pub actions: Vec<ItemStackRequestCereal>,
    pub strings_to_filter: Vec<String>,
    pub strings_to_filter_origin: TextProcessingEventOrigin,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct ItemStackRequestCerealSlotInfoData {
    pub full_container_name: FullContainerName,
    pub slot: wire::U8,
    pub net_id_variant: wire::I32LE,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct ItemStackRequestPacketDataRequestData {
    pub client_request_id: TypedClientNetIdStructItemStackRequestIdTagInt32T0,
    pub actions: Vec<ItemStackRequestCereal>,
    pub strings_to_filter: Vec<String>,
    pub strings_to_filter_origin: TextProcessingEventOrigin,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct ItemStackResponseContainerInfo {
    pub full_container_name: FullContainerName,
    pub slots: Vec<ItemStackResponseSlotInfo>,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct ItemStackResponseInfo {
    pub result: ItemStackNetResult,
    pub client_request_id: TypedClientNetIdStructItemStackRequestIdTagInt32T0,
    /// Wire presence: optional value is preceded by a presence marker.
    pub containers: Option<Vec<ItemStackResponseContainerInfo>>,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct ItemStackResponseSlotInfo {
    pub requested_slot: wire::U8,
    pub slot: wire::U8,
    pub amount: wire::U8,
    /// Wire presence: optional value is preceded by a presence marker.
    pub item_stack_net_id: Option<TypedServerNetIdStructItemStackNetIdTagInt32T0>,
    pub custom_name: BedrockSafetyRedactableString,
    pub durability_correction: wire::ZigZag32,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct ItemUseInventoryTransaction {
    pub actions: InventoryTransactionData,
    pub action_type: ItemUseInventoryTransactionActionType,
    pub trigger_type: ItemUseInventoryTransactionTriggerType,
    pub position: BlockPos,
    pub face: wire::U8,
    pub slot: wire::ZigZag32,
    pub item: CerealizerNetworkItemStackDescriptorSerializedData,
    pub from_position: glam::Vec3,
    pub click_position: glam::Vec3,
    pub target_block_id: wire::VarUInt,
    pub client_interact_prediction: ItemUseInventoryTransactionPredictedResult,
    pub client_cooldown_state: ItemUseInventoryTransactionClientCooldownState,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct LegacySetSlot {
    pub container_enum: ContainerEnumName,
    pub slots: Vec<wire::U8>,
}

#[derive(Clone, Debug, PartialEq)]
pub enum LegacyTelemetryEventEventData {
    Achievement {
        achievement_id: MinecraftEventingAchievementIds,
    },
    Interaction {
        interacted_entity_id: wire::ZigZag64,
        interaction_type: MinecraftEventingInteractionType,
        interaction_actor_type: wire::ZigZag32,
        interaction_actor_variant: wire::ZigZag32,
        interaction_actor_color: wire::U8,
    },
    PortalCreated {
        dimension_id: wire::ZigZag32,
    },
    PortalUsed {
        source_dimension_id: wire::ZigZag32,
        target_dimension_id: wire::ZigZag32,
    },
    MobKilled {
        instigator_actor_id: wire::ZigZag64,
        target_actor_id: wire::ZigZag64,
        instigator_child_actor_type: ActorType,
        damage_source: wire::ZigZag32,
        trade_tier: wire::ZigZag32,
        trader_name: String,
    },
    CauldronUsed {
        contents_color: wire::VarUInt,
        contents_type: wire::ZigZag32,
        fill_level: wire::ZigZag32,
    },
    PlayerDied {
        instigator_actor_id: wire::ZigZag32,
        instigator_mob_variant: wire::ZigZag32,
        damage_source: wire::ZigZag32,
        died_in_raid: bool,
    },
    BossKilled {
        boss_actor_id: wire::ZigZag64,
        party_size: wire::ZigZag32,
        boss_type: wire::ZigZag32,
    },
    SlashCommand {
        success_count: wire::ZigZag32,
        error_count: wire::ZigZag32,
        command_name: String,
        error_list: String,
    },
    MobBorn {
        born_baby_entity_type: wire::ZigZag32,
        born_baby_entity_variant: wire::ZigZag32,
        born_baby_color: wire::U8,
    },
    PoiCauldronUsed {
        block_interaction_type: MinecraftEventingPOIBlockInteractionType,
        item_id: wire::ZigZag32,
    },
    ComposterUsed {
        block_interaction_type: MinecraftEventingPOIBlockInteractionType,
        item_id: wire::ZigZag32,
    },
    BellUsed {
        item_id: wire::ZigZag32,
    },
    ActorDefinition {
        event_name: String,
    },
    RaidUpdate {
        current_wave: wire::ZigZag32,
        total_waves: wire::ZigZag32,
        success: bool,
    },
    TargetBlockHit {
        redstone_level: wire::ZigZag32,
    },
    PiglinBarter {
        item_id: wire::ZigZag32,
        was_targeting_bartering_player: bool,
    },
    PlayerWaxedOrUnwaxedCopper {
        player_waxed_or_unwaxed_copper_block_id: wire::ZigZag32,
    },
    CodeBuilderRuntimeAction {
        code_builder_runtime_action: String,
    },
    CodeBuilderScoreboard {
        objective_name: String,
        score: wire::ZigZag32,
    },
    ItemUsed {
        item_id: wire::I16LE,
        item_aux: wire::I32LE,
        use_method: wire::I32LE,
        count: wire::I32LE,
    },
    Empty,
}

impl LegacyTelemetryEventEventData {
    pub fn discriminant(&self) -> u32 {
        match self {
            Self::Achievement { .. } => 0,
            Self::Interaction { .. } => 1,
            Self::PortalCreated { .. } => 2,
            Self::PortalUsed { .. } => 3,
            Self::MobKilled { .. } => 4,
            Self::CauldronUsed { .. } => 5,
            Self::PlayerDied { .. } => 6,
            Self::BossKilled { .. } => 7,
            Self::SlashCommand { .. } => 8,
            Self::MobBorn { .. } => 9,
            Self::PoiCauldronUsed { .. } => 10,
            Self::ComposterUsed { .. } => 11,
            Self::BellUsed { .. } => 12,
            Self::ActorDefinition { .. } => 13,
            Self::RaidUpdate { .. } => 14,
            Self::TargetBlockHit { .. } => 15,
            Self::PiglinBarter { .. } => 16,
            Self::PlayerWaxedOrUnwaxedCopper { .. } => 17,
            Self::CodeBuilderRuntimeAction { .. } => 18,
            Self::CodeBuilderScoreboard { .. } => 19,
            Self::ItemUsed { .. } => 20,
            Self::Empty => 21,
        }
    }
}

impl Default for LegacyTelemetryEventEventData {
    fn default() -> Self {
        Self::Achievement {
            achievement_id: Default::default(),
        }
    }
}

#[derive(Clone, Copy, Debug, Default, PartialEq, Eq, Hash)]
pub struct LevelChunkSubChunkMetadata(pub u64);

impl wire::WireCodec for LevelChunkSubChunkMetadata {
    fn encode<W: std::io::Write>(&self, writer: &mut W) -> std::io::Result<()> {
        <wire::U64LE as wire::WireCodec>::encode(&wire::U64LE(self.0), writer)
    }

    fn decode<R: std::io::Read>(reader: &mut R) -> std::io::Result<Self> {
        <wire::U64LE as wire::WireCodec>::decode(reader).map(|value| Self(value.0))
    }
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct LevelSettings {
    pub seed: wire::U64LE,
    pub spawn_settings: SpawnSettings,
    pub generator_type: GeneratorType,
    pub game_type: GameType,
    pub is_hardcore: bool,
    pub game_difficulty: LegacyDifficulty,
    pub default_spawn_block_position: BlockPos,
    pub achievements_disabled: bool,
    pub editor_world_type: EditorWorldType,
    pub is_created_in_editor: bool,
    pub is_exported_from_editor: bool,
    pub day_cycle_stop_time: wire::ZigZag32,
    pub education_edition_offer: EducationEditionOffer,
    pub education_features_enabled: bool,
    pub education_product_id: String,
    pub rain_level: wire::F32LE,
    pub lightning_level: wire::F32LE,
    pub has_confirmed_platform_locked_content: bool,
    pub multiplayer_game_intent: bool,
    pub lan_broadcast_intent: bool,
    pub xbox_live_broadcast_setting: SocialGamePublishSetting,
    pub platform_broadcast_setting: SocialGamePublishSetting,
    pub commands_enabled: bool,
    pub texture_packs_required: bool,
    pub rule_data: GameRulesChangedPacketData,
    pub experiments: Experiments,
    pub has_bonus_chest_enabled: bool,
    pub start_with_map_enabled: bool,
    pub player_permissions: PlayerPermissionLevel,
    pub server_chunk_tick_range: wire::I32LE,
    pub has_locked_behavior_pack: bool,
    pub has_locked_resource_pack: bool,
    pub is_from_locked_template: bool,
    pub use_msa_gamertags_only: bool,
    pub is_from_world_template: bool,
    pub is_world_template_option_locked: bool,
    pub only_spawn_v1_villagers: bool,
    pub persona_disabled: bool,
    pub custom_skins_disabled: bool,
    pub emote_chat_muted: bool,
    pub base_game_version: String,
    pub limited_world_width: wire::I32LE,
    pub limited_world_depth: wire::I32LE,
    pub nether_type: bool,
    pub edu_shared_uri_resource: EduSharedUriResource,
    /// Wire presence: optional value is preceded by a presence marker.
    pub override_force_experimental_gameplay: Option<bool>,
    pub chat_restriction_level: ChatRestrictionLevel,
    pub disable_player_interactions: bool,
    pub server_editor_connection_policy: ServerEditorConnectionPolicy,
    pub allow_anonymous_block_drops_in_editor_worlds: bool,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct LocatorBarWaypoint {
    pub group_handle: WaypointGroupWaypointHandle,
    pub server_waypoint_payload: ServerWaypoint,
    pub action_flag: ServerWaypointGroupAction,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct MapDecoration {
    pub image_type: MapDecorationType,
    pub rotation: wire::U8,
    pub x: wire::U8,
    pub y: wire::U8,
    pub label: String,
    pub color: MceColor,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct MapInfoRequestPacketAnonClientPixelsProxy {
    pub pixel: wire::U32LE,
    pub index: wire::U16LE,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct MapItemTrackedActorUniqueId {
    pub type_: MapItemTrackedActorType,
    /// Wire presence: optional value is preceded by a presence marker.
    pub entity_id: Option<ActorUniqueID>,
    /// Wire presence: optional value is preceded by a presence marker.
    pub block_position: Option<BlockPos>,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct MaterialReducerDataEntry {
    pub from_item_key: wire::ZigZag32,
    pub item_ids_and_counts: Vec<MaterialReducerEntryOutput>,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct MaterialReducerEntryOutput {
    pub item_id: wire::ZigZag32,
    pub item_count: wire::ZigZag32,
}

#[derive(Clone, Copy, Debug, Default, PartialEq, Eq, Hash)]
pub struct MceColor(pub i32);

impl wire::WireCodec for MceColor {
    fn encode<W: std::io::Write>(&self, writer: &mut W) -> std::io::Result<()> {
        <wire::I32LE as wire::WireCodec>::encode(&wire::I32LE(self.0), writer)
    }

    fn decode<R: std::io::Read>(reader: &mut R) -> std::io::Result<Self> {
        <wire::I32LE as wire::WireCodec>::decode(reader).map(|value| Self(value.0))
    }
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct MemoryMemoryCategoryCounter {
    pub category: MemoryMemoryCategory,
    pub current_bytes: wire::U64LE,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct MissingBlobData {
    pub blob_id: wire::U64LE,
    pub blob_data: bytes::Bytes,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct MoveActorAbsoluteData {
    pub actor_runtime_id: ActorRuntimeID,
    pub header: wire::U8,
    pub position: glam::Vec3,
    pub rotation_x: wire::U8,
    pub rotation_y: wire::U8,
    pub rotation_y_head: wire::U8,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct MoveActorDeltaData {
    pub actor_runtime_id: ActorRuntimeID,
    /// Wire presence: optional value is preceded by a presence marker.
    pub new_position_x: Option<wire::F32LE>,
    /// Wire presence: optional value is preceded by a presence marker.
    pub new_position_y: Option<wire::F32LE>,
    /// Wire presence: optional value is preceded by a presence marker.
    pub new_position_z: Option<wire::F32LE>,
    /// Wire presence: optional value is preceded by a presence marker.
    pub rotation_x: Option<wire::I8>,
    /// Wire presence: optional value is preceded by a presence marker.
    pub rotation_y: Option<wire::I8>,
    /// Wire presence: optional value is preceded by a presence marker.
    pub rotation_y_head: Option<wire::I8>,
    pub is_on_ground: bool,
    pub force_move: bool,
    pub force_move_local_entity: bool,
    pub force_completion: bool,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct MovePlayerTeleportData {
    pub teleportation_cause: wire::I32LE,
    pub source_actor_type: wire::I32LE,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct MultiRecipe {
    pub multi_recipe_uuid: uuid::Uuid,
    pub net_id: TypedServerNetIdStructRecipeNetIdTag,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct NetworkPermissions {
    pub server_auth_sound_enabled: bool,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct NoiseDescriptor {
    pub name: String,
    pub first_octave: wire::I32LE,
    pub amplitudes: Vec<wire::F32LE>,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct PackIdVersion {
    pub pack_uuid: uuid::Uuid,
    pub pack_version: SemVersion,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct PackIdVersionData {
    pub pack_uuid: uuid::Uuid,
    pub pack_version: SemVersionData,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct PackInfoData {
    pub pack_id_version: PackIdVersionData,
    pub pack_size: wire::U64LE,
    pub content_key: String,
    pub subpack_name: String,
    pub content_identity: ContentIdentity,
    pub has_scripts: bool,
    pub is_addon_pack: bool,
    pub is_ray_tracing_capable: bool,
    pub cdn_url: String,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct PackInstanceId {
    pub pack_id: String,
    pub version: String,
    pub sub_pack_name: String,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct PackedItemUseLegacyInventoryTransaction {
    pub legacy_request_id: TypedClientNetIdStructItemStackLegacyRequestIdTagInt32T0,
    /// Wire presence: optional value is preceded by a presence marker.
    pub legacy_set_item_slots: Option<Vec<LegacySetSlot>>,
    /// Wire presence: optional value is preceded by a presence marker.
    pub item_use_transaction: Option<ItemUseInventoryTransaction>,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct PlayerBlockActionData {
    pub player_action_type: PlayerActionType,
    pub position: BlockPos,
    pub facing: wire::ZigZag32,
}

#[derive(Clone, Copy, Debug, Default, PartialEq, Eq, Hash)]
pub struct PlayerInputTick(pub u64);

impl wire::WireCodec for PlayerInputTick {
    fn encode<W: std::io::Write>(&self, writer: &mut W) -> std::io::Result<()> {
        <wire::VarULong as wire::WireCodec>::encode(&wire::VarULong(self.0), writer)
    }

    fn decode<R: std::io::Read>(reader: &mut R) -> std::io::Result<Self> {
        <wire::VarULong as wire::WireCodec>::decode(reader).map(|value| Self(value.0))
    }
}

#[derive(Clone, Debug, PartialEq)]
pub enum PlayerListEntriesItem {
    Add {
        uuid: uuid::Uuid,
        actor_unique_id: ActorUniqueID,
        player_name: String,
        xbl_xuid: String,
        platform_online_id: String,
        build_platform: BuildPlatform,
        serialized_skin: Box<SerializedSkinRef>,
        is_teacher: bool,
        is_host: bool,
        is_sub_client: bool,
        player_color: MceColor,
    },
    Remove {
        uuid: uuid::Uuid,
    },
}

impl PlayerListEntriesItem {
    pub fn discriminant(&self) -> u8 {
        match self {
            Self::Add { .. } => 0,
            Self::Remove { .. } => 1,
        }
    }
}

impl Default for PlayerListEntriesItem {
    fn default() -> Self {
        Self::Add {
            uuid: Default::default(),
            actor_unique_id: Default::default(),
            player_name: Default::default(),
            xbl_xuid: Default::default(),
            platform_online_id: Default::default(),
            build_platform: Default::default(),
            serialized_skin: Default::default(),
            is_teacher: Default::default(),
            is_host: Default::default(),
            is_sub_client: Default::default(),
            player_color: Default::default(),
        }
    }
}

#[derive(Clone, Debug, PartialEq)]
pub enum PlayerLocationLocation {
    PlayerLocationCoordinates {
        packet_type: PlayerLocationType,
        position: glam::Vec3,
    },
    PlayerLocationHide {
        packet_type: PlayerLocationType,
    },
}

impl PlayerLocationLocation {
    pub fn discriminant(&self) -> u32 {
        match self {
            Self::PlayerLocationCoordinates { .. } => 0,
            Self::PlayerLocationHide { .. } => 1,
        }
    }
}

impl Default for PlayerLocationLocation {
    fn default() -> Self {
        Self::PlayerLocationCoordinates {
            packet_type: Default::default(),
            position: Default::default(),
        }
    }
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct PlayerPartyInfo {
    pub party_id: String,
    pub is_party_leader: bool,
}

#[derive(Clone, Copy, Debug, Default, PartialEq, Eq, Hash)]
pub struct PlayerScoreboardId(pub i64);

impl wire::WireCodec for PlayerScoreboardId {
    fn encode<W: std::io::Write>(&self, writer: &mut W) -> std::io::Result<()> {
        <wire::ZigZag64 as wire::WireCodec>::encode(&wire::ZigZag64(self.0), writer)
    }

    fn decode<R: std::io::Read>(reader: &mut R) -> std::io::Result<Self> {
        <wire::ZigZag64 as wire::WireCodec>::decode(reader).map(|value| Self(value.0))
    }
}

#[derive(Clone, Debug, PartialEq)]
pub enum PlayerUpdateEntityOverridesUpdate {
    ClearOverride {
        type_: String,
    },
    RemoveOverride {
        type_: String,
    },
    IntOverride {
        type_: String,
        value: wire::I32LE,
    },
    FloatOverride {
        type_: String,
        value: wire::F32LE,
    },
}

impl PlayerUpdateEntityOverridesUpdate {
    pub fn discriminant(&self) -> u8 {
        match self {
            Self::ClearOverride { .. } => 0,
            Self::RemoveOverride { .. } => 1,
            Self::IntOverride { .. } => 2,
            Self::FloatOverride { .. } => 3,
        }
    }
}

impl Default for PlayerUpdateEntityOverridesUpdate {
    fn default() -> Self {
        Self::ClearOverride {
            type_: Default::default(),
        }
    }
}

#[derive(Clone, Debug, PartialEq, Default)]
pub enum PlayerVideoCaptureAction {
    #[default]
    StopVideoCapture,
    StartVideoCapture {
        frame_rate: wire::U32LE,
        file_prefix: String,
    },
}

impl PlayerVideoCaptureAction {
    pub fn discriminant(&self) -> u8 {
        match self {
            Self::StopVideoCapture => 0,
            Self::StartVideoCapture { .. } => 1,
        }
    }
}

#[derive(Clone, Copy, Debug, Default, PartialEq, Eq, Hash)]
pub struct PositionTrackingId(pub i32);

impl wire::WireCodec for PositionTrackingId {
    fn encode<W: std::io::Write>(&self, writer: &mut W) -> std::io::Result<()> {
        <wire::ZigZag32 as wire::WireCodec>::encode(&wire::ZigZag32(self.0), writer)
    }

    fn decode<R: std::io::Read>(reader: &mut R) -> std::io::Result<Self> {
        <wire::ZigZag32 as wire::WireCodec>::decode(reader).map(|value| Self(value.0))
    }
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct PotionMixDataEntry {
    pub from_potion_id: wire::ZigZag32,
    pub from_item_aux: wire::ZigZag32,
    pub reagent_item_id: wire::ZigZag32,
    pub reagent_item_aux: wire::ZigZag32,
    pub to_potion_id: wire::ZigZag32,
    pub to_item_aux: wire::ZigZag32,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct PrimitiveShapeData {
    pub network_id: wire::VarULong,
    /// Wire presence: optional value is preceded by a presence marker.
    pub shape_type: Option<ScriptModuleMinecraftScriptPrimitiveShapeType>,
    /// Wire presence: optional value is preceded by a presence marker.
    pub location: Option<glam::Vec3>,
    /// Wire presence: optional value is preceded by a presence marker.
    pub scale: Option<wire::F32LE>,
    /// Wire presence: optional value is preceded by a presence marker.
    pub rotation: Option<glam::Vec3>,
    /// Wire presence: optional value is preceded by a presence marker.
    pub total_time_left: Option<wire::F32LE>,
    /// Wire presence: optional value is preceded by a presence marker.
    pub maximum_render_distance: Option<wire::F32LE>,
    /// Wire presence: optional value is preceded by a presence marker.
    pub color: Option<MceColor>,
    /// Wire presence: optional value is preceded by a presence marker.
    pub dimension_id: Option<DimensionType>,
    /// Wire presence: optional value is preceded by a presence marker.
    pub attached_to_entity_id: Option<ActorUniqueID>,
    pub extra_shape_data: PrimitiveShapeDataExtraShapeData,
}

#[derive(Clone, Debug, PartialEq, Default)]
pub enum PrimitiveShapeDataExtraShapeData {
    /// Naming overlay required: source placeholder `Empty0`.
    #[default]
    Empty0,
    ArrowData {
        /// Wire presence: optional value is preceded by a presence marker.
        arrow_end_location: Option<glam::Vec3>,
        /// Wire presence: optional value is preceded by a presence marker.
        arrow_head_length: Option<wire::F32LE>,
        /// Wire presence: optional value is preceded by a presence marker.
        arrow_head_radius: Option<wire::F32LE>,
        /// Wire presence: optional value is preceded by a presence marker.
        num_segments: Option<wire::U8>,
    },
    TextData {
        text: String,
        use_rotation: bool,
        /// Wire presence: optional value is preceded by a presence marker.
        background_color: Option<MceColor>,
        depth_test: bool,
        show_backface: bool,
        show_text_backface: bool,
    },
    BoxData {
        box_bound: glam::Vec3,
    },
    LineData {
        line_end_location: glam::Vec3,
    },
    SphereData {
        num_segments: wire::U8,
    },
    CylinderData {
        radius_x: glam::Vec2,
        radius_z: glam::Vec2,
        height: wire::F32LE,
        num_segments: wire::U8,
    },
    PyramidData {
        width: wire::F32LE,
        /// Wire presence: optional value is preceded by a presence marker.
        depth: Option<wire::F32LE>,
        height: wire::F32LE,
    },
    EllipsoidData {
        radii: glam::Vec3,
        segments_per_axis: wire::U8,
    },
    ConeData {
        radii: glam::Vec2,
        height: wire::F32LE,
        num_segments: wire::U8,
    },
}

impl PrimitiveShapeDataExtraShapeData {
    pub fn discriminant(&self) -> u32 {
        match self {
            Self::Empty0 => 0,
            Self::ArrowData { .. } => 1,
            Self::TextData { .. } => 2,
            Self::BoxData { .. } => 3,
            Self::LineData { .. } => 4,
            Self::SphereData { .. } => 5,
            Self::CylinderData { .. } => 6,
            Self::PyramidData { .. } => 7,
            Self::EllipsoidData { .. } => 8,
            Self::ConeData { .. } => 9,
        }
    }
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct PropertySyncData {
    pub int_entries_list: Vec<PropertySyncDataPropertySyncIntEntry>,
    pub float_entries_list: Vec<PropertySyncDataPropertySyncFloatEntry>,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct PropertySyncDataPropertySyncFloatEntry {
    pub property_index: wire::VarUInt,
    pub data: wire::F32LE,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct PropertySyncDataPropertySyncIntEntry {
    pub property_index: wire::VarUInt,
    pub data: wire::ZigZag32,
}

#[derive(Clone, Debug, PartialEq)]
pub enum ResourcePackClientResponseResponse {
    Cancel {
        response_type: String,
    },
    Downloading {
        response_type: String,
        downloading_packs: Vec<String>,
    },
    DownloadingFinished {
        response_type: String,
    },
    ResourcePackStackFinished {
        response_type: String,
    },
}

impl ResourcePackClientResponseResponse {
    pub fn discriminant(&self) -> i8 {
        match self {
            Self::Cancel { .. } => 1,
            Self::Downloading { .. } => 2,
            Self::DownloadingFinished { .. } => 3,
            Self::ResourcePackStackFinished { .. } => 4,
        }
    }
}

impl Default for ResourcePackClientResponseResponse {
    fn default() -> Self {
        Self::Cancel {
            response_type: Default::default(),
        }
    }
}

#[derive(Clone, Copy, Debug, Default, PartialEq, Eq, Hash)]
pub struct ScoreboardId(pub i64);

impl wire::WireCodec for ScoreboardId {
    fn encode<W: std::io::Write>(&self, writer: &mut W) -> std::io::Result<()> {
        <wire::ZigZag64 as wire::WireCodec>::encode(&wire::ZigZag64(self.0), writer)
    }

    fn decode<R: std::io::Read>(reader: &mut R) -> std::io::Result<Self> {
        <wire::ZigZag64 as wire::WireCodec>::decode(reader).map(|value| Self(value.0))
    }
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct ScoreboardIdentityPacketInfo {
    pub scoreboard_id: ScoreboardId,
    /// Wire presence: optional value is preceded by a presence marker.
    pub player_unique_id: Option<wire::ZigZag64>,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct SemVersion {
    pub version: String,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct SemVersionData {
    pub version: String,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct SerializedAbilitiesData {
    pub target_player_raw_id: wire::I64LE,
    pub player_permissions: PlayerPermissionLevel,
    pub command_permissions: CommandPermissionLevel,
    pub layers: Vec<SerializedAbilitiesDataSerializedLayer>,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct SerializedAbilitiesDataSerializedLayer {
    pub serialized_layer: wire::U16LE,
    pub abilities_set: wire::U32LE,
    pub ability_values: wire::U32LE,
    pub fly_speed: wire::F32LE,
    pub vertical_fly_speed: wire::F32LE,
    pub walk_speed: wire::F32LE,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct SerializedNoiseBlockSpecifier {
    pub noise: String,
    pub threshold: wire::F32LE,
    pub range: FloatRange,
    pub block: wire::U32LE,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct SerializedPersonaPieceHandle {
    pub piece_id: String,
    pub piece_type: PersonaPieceType,
    pub pack_id: uuid::Uuid,
    pub is_default_piece: bool,
    pub product_id: String,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct SerializedSkinRef {
    pub id: String,
    pub play_fab_id: String,
    pub resource_patch: String,
    pub image_data: SkinImage,
    pub animated_image_data: Vec<AnimatedImageData>,
    pub cape_image_data: SkinImage,
    pub geometry_data: String,
    pub geometry_data_min_engine_version: String,
    pub animation_data: String,
    pub cape_id: String,
    pub full_id: String,
    pub arm_size: PersonaArmSizeType,
    pub skin_color: MceColor,
    pub persona_pieces: Vec<SerializedPersonaPieceHandle>,
    pub piece_tint_colors: Vec<(String, TintMapColor)>,
    pub is_premium: bool,
    pub is_persona: bool,
    pub is_persona_cape_on_classic_skin: bool,
    pub is_primary_user: bool,
    pub overrides_player_appearance: bool,
    pub trusted_skin_flag: String,
    pub profile_hash: String,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct ServerBlockProperty {
    pub block_name: String,
    pub block_definition: Nbt,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct ServerConfigurationClientStoreEntryPointConfiguration {
    pub store_id: String,
    pub store_name: String,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct ServerConfigurationGatheringsConfigurationJoinInfo {
    pub experience_id: uuid::Uuid,
    pub experience_name: String,
    /// Wire presence: optional value is preceded by a presence marker.
    pub world_id: Option<uuid::Uuid>,
    /// Wire presence: optional value is preceded by a presence marker.
    pub world_name: Option<String>,
    pub creator_id: String,
    /// Wire presence: optional value is preceded by a presence marker.
    pub target_id: Option<uuid::Uuid>,
    /// Wire presence: optional value is preceded by a presence marker.
    pub scenario_id: Option<String>,
    /// Wire presence: optional value is preceded by a presence marker.
    pub server_id: Option<String>,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct ServerConfigurationPresenceConfiguration {
    /// Wire presence: optional value is preceded by a presence marker.
    pub rich_presence_id: Option<String>,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct ServerConfigurationServerConfigurationJoinInfo {
    /// Wire presence: optional value is preceded by a presence marker.
    pub gathering: Option<ServerConfigurationGatheringsConfigurationJoinInfo>,
    /// Wire presence: optional value is preceded by a presence marker.
    pub client_store_entry_point: Option<ServerConfigurationClientStoreEntryPointConfiguration>,
    /// Wire presence: optional value is preceded by a presence marker.
    pub presence: Option<ServerConfigurationPresenceConfiguration>,
}

#[derive(Clone, Copy, Debug, Default, PartialEq, Eq, Hash)]
pub struct ServerSoundHandle(pub u64);

impl wire::WireCodec for ServerSoundHandle {
    fn encode<W: std::io::Write>(&self, writer: &mut W) -> std::io::Result<()> {
        <wire::U64LE as wire::WireCodec>::encode(&wire::U64LE(self.0), writer)
    }

    fn decode<R: std::io::Read>(reader: &mut R) -> std::io::Result<Self> {
        <wire::U64LE as wire::WireCodec>::decode(reader).map(|value| Self(value.0))
    }
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct ServerWaypoint {
    pub update_flag: wire::U32LE,
    /// Wire presence: optional value is preceded by a presence marker.
    pub is_visible: Option<bool>,
    /// Wire presence: optional value is preceded by a presence marker.
    pub world_position: Option<WorldPosition>,
    /// Wire presence: optional value is preceded by a presence marker.
    pub texture_path: Option<String>,
    /// Wire presence: optional value is preceded by a presence marker.
    pub icon_size: Option<glam::Vec2>,
    /// Wire presence: optional value is preceded by a presence marker.
    pub color: Option<MceColor>,
    /// Wire presence: optional value is preceded by a presence marker.
    pub client_position_authority: Option<bool>,
    /// Wire presence: optional value is preceded by a presence marker.
    pub actor_unique_id: Option<ActorUniqueID>,
}

#[derive(Clone, Debug, PartialEq)]
pub enum ServerboundPackSettingChangePackSettingValue {
    Float(wire::F32LE),
    Bool(bool),
    String(String),
}

impl ServerboundPackSettingChangePackSettingValue {
    pub fn discriminant(&self) -> u32 {
        match self {
            Self::Float(..) => 0,
            Self::Bool(..) => 1,
            Self::String(..) => 2,
        }
    }
}

impl Default for ServerboundPackSettingChangePackSettingValue {
    fn default() -> Self {
        Self::Float(Default::default())
    }
}

#[derive(Clone, Debug, PartialEq)]
pub enum SetScoreScoreInfoItem {
    RemoveScore {
        action: String,
        scoreboard_id: ScoreboardId,
        /// Wire presence: optional value is preceded by a presence marker.
        objective_name: Option<String>,
    },
    ChangePlayerScore {
        action: String,
        scoreboard_id: ScoreboardId,
        objective_name: String,
        score_value: wire::I32LE,
        player_unique_id: PlayerScoreboardId,
    },
    ChangeEntityScore {
        action: String,
        scoreboard_id: ScoreboardId,
        objective_name: String,
        score_value: wire::I32LE,
        actor_id: ActorUniqueID,
    },
    ChangeFakePlayerScore {
        action: String,
        scoreboard_id: ScoreboardId,
        objective_name: String,
        score_value: wire::I32LE,
        fake_player_name: String,
    },
}

impl SetScoreScoreInfoItem {
    pub fn discriminant(&self) -> u8 {
        match self {
            Self::RemoveScore { .. } => 0,
            Self::ChangePlayerScore { .. } => 1,
            Self::ChangeEntityScore { .. } => 2,
            Self::ChangeFakePlayerScore { .. } => 3,
        }
    }
}

impl Default for SetScoreScoreInfoItem {
    fn default() -> Self {
        Self::RemoveScore {
            action: Default::default(),
            scoreboard_id: Default::default(),
            objective_name: Default::default(),
        }
    }
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct ShapedRecipe {
    pub recipe_id: String,
    pub width: wire::ZigZag32,
    pub height: wire::ZigZag32,
    pub ingredients: Vec<CerealizerRecipeIngredientSerializedData>,
    pub results: Vec<CerealizerNetworkItemInstanceDescriptorSerializedData>,
    pub uuid: uuid::Uuid,
    pub tag: String,
    pub priority: wire::ZigZag32,
    pub assume_symmetry: bool,
    /// Wire presence: optional value is preceded by a presence marker.
    pub unlocking_requirement: Option<CerealizerRecipeUnlockingRequirementSerializedData>,
    pub net_id: TypedServerNetIdStructRecipeNetIdTag,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct ShapelessRecipe {
    pub recipe_id: String,
    pub ingredients: Vec<CerealizerRecipeIngredientSerializedData>,
    pub results: Vec<CerealizerNetworkItemInstanceDescriptorSerializedData>,
    pub uuid: uuid::Uuid,
    pub tag: String,
    pub priority: wire::ZigZag32,
    /// Wire presence: optional value is preceded by a presence marker.
    pub unlocking_requirement: Option<CerealizerRecipeUnlockingRequirementSerializedData>,
    pub net_id: TypedServerNetIdStructRecipeNetIdTag,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct SkinImage {
    pub width: wire::U32LE,
    pub height: wire::U32LE,
    pub image_bytes: Vec<wire::U8>,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct SmithingTransformRecipe {
    pub recipe_id: String,
    pub template_ingredient: CerealizerRecipeIngredientSerializedData,
    pub base_ingredient: CerealizerRecipeIngredientSerializedData,
    pub addition_ingredient: CerealizerRecipeIngredientSerializedData,
    pub result: CerealizerNetworkItemInstanceDescriptorSerializedData,
    pub tag: String,
    pub net_id: TypedServerNetIdStructRecipeNetIdTag,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct SmithingTrimRecipe {
    pub recipe_id: String,
    pub template_ingredient: CerealizerRecipeIngredientSerializedData,
    pub base_ingredient: CerealizerRecipeIngredientSerializedData,
    pub addition_ingredient: CerealizerRecipeIngredientSerializedData,
    pub tag: String,
    pub net_id: TypedServerNetIdStructRecipeNetIdTag,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct SocialEventsServerTelemetryData {
    pub server_id: String,
    pub scenario_id: String,
    pub world_id: String,
    pub owner_id: String,
}

#[derive(Clone, Debug, PartialEq, Default)]
pub enum SoundDataEvent {
    #[default]
    Stop,
    SetVolume {
        volume: wire::F32LE,
    },
    SetPitch {
        pitch: wire::F32LE,
    },
    Fade {
        duration: wire::F32LE,
        target_volume: wire::F32LE,
    },
    SeekTo {
        seconds: wire::F32LE,
    },
    Pause,
    Resume,
}

impl SoundDataEvent {
    pub fn discriminant(&self) -> u32 {
        match self {
            Self::Stop => 0,
            Self::SetVolume { .. } => 1,
            Self::SetPitch { .. } => 2,
            Self::Fade { .. } => 3,
            Self::SeekTo { .. } => 4,
            Self::Pause => 5,
            Self::Resume => 6,
        }
    }
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct SpawnSettings {
    pub spawn_biome_type: SpawnBiomeType,
    pub user_defined_biome_name: String,
    pub dimension: wire::ZigZag32,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct StructureEditorData {
    pub structure_name: BedrockSafetyRedactableString,
    pub data_field: String,
    pub should_include_players: bool,
    pub should_show_bounding_box: bool,
    pub structure_block_type: StructureBlockType,
    pub structure_settings: StructureSettings,
    pub redstone_save_mode: StructureRedstoneSaveMode,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct StructureSettings {
    pub structure_palette_name: String,
    pub should_ignore_entities: bool,
    pub should_ignore_blocks: bool,
    pub should_allow_non_ticking_player_and_ticking_area_chunks: bool,
    pub structure_size: BlockPos,
    pub structure_offset: BlockPos,
    pub last_edit_player: ActorUniqueID,
    pub rotation: Rotation,
    pub mirror: Mirror,
    pub animation_mode: AnimationMode,
    pub animation_seconds: wire::F32LE,
    pub integrity_value: wire::F32LE,
    pub integrity_seed: wire::U32LE,
    pub rotation_pivot: glam::Vec3,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct SubChunkHeightmapData {
    pub height_map_type: SubChunkHeightMapDataType,
    /// Wire presence: optional value is preceded by a presence marker.
    pub subchunk_height_map: Option<[[wire::I8; 16]; 16]>,
    pub render_height_map_type: SubChunkHeightMapDataType,
    /// Wire presence: optional value is preceded by a presence marker.
    pub subchunk_render_height_map: Option<[[wire::I8; 16]; 16]>,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct SubChunkPos {
    pub subchunk_position_x: wire::I32LE,
    pub subchunk_position_y: wire::I32LE,
    pub subchunk_position_z: wire::I32LE,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct SubChunkSubChunkPacketData {
    pub sub_chunk_pos_offset: SubChunkSubChunkPosOffset,
    pub sub_chunk_request_result: SubChunkSubChunkRequestResult,
    /// Wire presence: optional value is preceded by a presence marker.
    pub serialized_sub_chunk: Option<String>,
    pub height_map_data: SubChunkHeightmapData,
    /// Wire presence: optional value is preceded by a presence marker.
    pub blob_id: Option<wire::U64LE>,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct SubChunkSubChunkPosOffset {
    pub subchunk_offset_x: wire::I8,
    pub subchunk_offset_y: wire::I8,
    pub subchunk_offset_z: wire::I8,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct SyncWorldClockStateData {
    pub clock_id: wire::VarULong,
    pub time: wire::ZigZag32,
    pub is_paused: bool,
}

#[derive(Clone, Debug, PartialEq)]
pub enum SyncWorldClocksData {
    SyncStateData {
        clock_data: Vec<SyncWorldClockStateData>,
    },
    InitializeRegistryData {
        clock_data: Vec<WorldClockData>,
    },
    AddTimeMarkerData {
        clock_id: wire::VarULong,
        time_markers: Vec<TimeMarkerData>,
    },
    RemoveTimeMarkerData {
        clock_id: wire::VarULong,
        time_marker_ids: Vec<wire::VarULong>,
    },
}

impl SyncWorldClocksData {
    pub fn discriminant(&self) -> u32 {
        match self {
            Self::SyncStateData { .. } => 0,
            Self::InitializeRegistryData { .. } => 1,
            Self::AddTimeMarkerData { .. } => 2,
            Self::RemoveTimeMarkerData { .. } => 3,
        }
    }
}

impl Default for SyncWorldClocksData {
    fn default() -> Self {
        Self::SyncStateData {
            clock_data: Default::default(),
        }
    }
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct SyncedAttribute {
    pub attribute_name: String,
    pub min_value: wire::F32LE,
    pub current_value: wire::F32LE,
    pub max_value: wire::F32LE,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct SyncedPlayerMovementSettings {
    pub rewind_history_size: wire::ZigZag32,
    pub server_authoritative_block_breaking: bool,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct SynchedActorDataCopyableDataList {
    pub data: Vec<DataItemEntry>,
}

#[derive(Clone, Debug, PartialEq)]
pub enum TextBody {
    Raw {
        message: String,
    },
    Chat {
        player_name: String,
        message: String,
    },
    Translate {
        message: String,
        parameter_list: Vec<String>,
    },
    Popup {
        message: String,
        parameter_list: Vec<String>,
    },
    JukeboxPopup {
        message: String,
        parameter_list: Vec<String>,
    },
    Tip {
        message: String,
    },
    SystemMessage {
        message: String,
    },
    Whisper {
        player_name: String,
        message: String,
    },
    Announcement {
        player_name: String,
        message: String,
    },
    TextObjectWhisper {
        message: String,
    },
    TextObject {
        message: String,
    },
    TextObjectAnnouncement {
        message: String,
    },
}

impl TextBody {
    pub fn discriminant(&self) -> u8 {
        match self {
            Self::Raw { .. } => 0,
            Self::Chat { .. } => 1,
            Self::Translate { .. } => 2,
            Self::Popup { .. } => 3,
            Self::JukeboxPopup { .. } => 4,
            Self::Tip { .. } => 5,
            Self::SystemMessage { .. } => 6,
            Self::Whisper { .. } => 7,
            Self::Announcement { .. } => 8,
            Self::TextObjectWhisper { .. } => 9,
            Self::TextObject { .. } => 10,
            Self::TextObjectAnnouncement { .. } => 11,
        }
    }
}

impl Default for TextBody {
    fn default() -> Self {
        Self::Raw {
            message: Default::default(),
        }
    }
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct TimeMarkerData {
    pub id: wire::VarULong,
    pub name: String,
    pub time: wire::ZigZag32,
    /// Wire presence: optional value is preceded by a presence marker.
    pub period: Option<wire::I32LE>,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct TintMapColor {
    pub colors: [MceColor; 4],
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct TrimMaterial {
    pub material_id: String,
    pub color: String,
    pub item_name: String,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct TrimPattern {
    pub item_name: String,
    pub pattern_id: String,
}

#[derive(Clone, Copy, Debug, Default, PartialEq, Eq, Hash)]
pub struct TypedClientNetIdStructItemStackLegacyRequestIdTagInt32T0(pub i32);

impl wire::WireCodec for TypedClientNetIdStructItemStackLegacyRequestIdTagInt32T0 {
    fn encode<W: std::io::Write>(&self, writer: &mut W) -> std::io::Result<()> {
        <wire::ZigZag32 as wire::WireCodec>::encode(&wire::ZigZag32(self.0), writer)
    }

    fn decode<R: std::io::Read>(reader: &mut R) -> std::io::Result<Self> {
        <wire::ZigZag32 as wire::WireCodec>::decode(reader).map(|value| Self(value.0))
    }
}

#[derive(Clone, Copy, Debug, Default, PartialEq, Eq, Hash)]
pub struct TypedClientNetIdStructItemStackRequestIdTagInt32T0(pub i32);

impl wire::WireCodec for TypedClientNetIdStructItemStackRequestIdTagInt32T0 {
    fn encode<W: std::io::Write>(&self, writer: &mut W) -> std::io::Result<()> {
        <wire::ZigZag32 as wire::WireCodec>::encode(&wire::ZigZag32(self.0), writer)
    }

    fn decode<R: std::io::Read>(reader: &mut R) -> std::io::Result<Self> {
        <wire::ZigZag32 as wire::WireCodec>::decode(reader).map(|value| Self(value.0))
    }
}

#[derive(Clone, Copy, Debug, Default, PartialEq, Eq, Hash)]
pub struct TypedServerNetIdStructCreativeItemNetIdTag(pub u32);

impl wire::WireCodec for TypedServerNetIdStructCreativeItemNetIdTag {
    fn encode<W: std::io::Write>(&self, writer: &mut W) -> std::io::Result<()> {
        <wire::VarUInt as wire::WireCodec>::encode(&wire::VarUInt(self.0), writer)
    }

    fn decode<R: std::io::Read>(reader: &mut R) -> std::io::Result<Self> {
        <wire::VarUInt as wire::WireCodec>::decode(reader).map(|value| Self(value.0))
    }
}

#[derive(Clone, Copy, Debug, Default, PartialEq, Eq, Hash)]
pub struct TypedServerNetIdStructItemStackNetIdTagInt32T0(pub i32);

impl wire::WireCodec for TypedServerNetIdStructItemStackNetIdTagInt32T0 {
    fn encode<W: std::io::Write>(&self, writer: &mut W) -> std::io::Result<()> {
        <wire::ZigZag32 as wire::WireCodec>::encode(&wire::ZigZag32(self.0), writer)
    }

    fn decode<R: std::io::Read>(reader: &mut R) -> std::io::Result<Self> {
        <wire::ZigZag32 as wire::WireCodec>::decode(reader).map(|value| Self(value.0))
    }
}

#[derive(Clone, Copy, Debug, Default, PartialEq, Eq, Hash)]
pub struct TypedServerNetIdStructRecipeNetIdTag(pub u32);

impl wire::WireCodec for TypedServerNetIdStructRecipeNetIdTag {
    fn encode<W: std::io::Write>(&self, writer: &mut W) -> std::io::Result<()> {
        <wire::VarUInt as wire::WireCodec>::encode(&wire::VarUInt(self.0), writer)
    }

    fn decode<R: std::io::Read>(reader: &mut R) -> std::io::Result<Self> {
        <wire::VarUInt as wire::WireCodec>::decode(reader).map(|value| Self(value.0))
    }
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct UpdateSubChunkBlocksChangedInfo {
    pub blocks_changed_standards: Vec<UpdateSubChunkNetworkBlockInfo>,
    pub blocks_changed_extras: Vec<UpdateSubChunkNetworkBlockInfo>,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct UpdateSubChunkNetworkBlockInfo {
    pub pos: BlockPos,
    pub runtime_id: wire::VarUInt,
    pub update_flags: wire::VarUInt,
    pub sync_message_entity_unique_id: wire::VarULong,
    pub sync_message_message: wire::VarUInt,
}

#[derive(Clone, Copy, Debug, Default, PartialEq, Eq, Hash)]
pub struct VoxelShapesRegistryHandle(pub u16);

impl wire::WireCodec for VoxelShapesRegistryHandle {
    fn encode<W: std::io::Write>(&self, writer: &mut W) -> std::io::Result<()> {
        <wire::U16LE as wire::WireCodec>::encode(&wire::U16LE(self.0), writer)
    }

    fn decode<R: std::io::Read>(reader: &mut R) -> std::io::Result<Self> {
        <wire::U16LE as wire::WireCodec>::decode(reader).map(|value| Self(value.0))
    }
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct VoxelShapesSerializableCells {
    pub x_size: wire::U8,
    pub y_size: wire::U8,
    pub z_size: wire::U8,
    pub storage: Vec<wire::U8>,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct VoxelShapesSerializableVoxelShape {
    pub cells: VoxelShapesSerializableCells,
    pub x_coordinates: Vec<wire::F32LE>,
    pub y_coordinates: Vec<wire::F32LE>,
    pub z_coordinates: Vec<wire::F32LE>,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct WaypointGroupWaypointHandle {
    pub uuid: uuid::Uuid,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct WebSocketPacketData {
    pub websocket_server_uri: String,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct WorldClockData {
    pub id: wire::VarULong,
    pub name: String,
    pub time: wire::ZigZag32,
    pub is_paused: bool,
    pub time_markers: Vec<TimeMarkerData>,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct WorldPosition {
    pub position: glam::Vec3,
    pub dimension_type: DimensionType,
}
