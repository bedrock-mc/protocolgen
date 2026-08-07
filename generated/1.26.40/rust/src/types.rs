// Code generated from canonical protocol manifest v2. DO NOT EDIT.

use crate::enums::*;

use crate::wire;

// Domain: actor

#[derive(Clone, Debug, Default, PartialEq)]
pub struct ActorDataBoundingBoxComponent {
    pub actor_data_bounding_box: [wire::F32LE; 3],
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct ActorDataFlagComponent {
    pub actor_flag_bitset_data: Bitset131,
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

// Domain: attribute

/// AttributeModifier temporarily buffs/debuffs a given attribute until the modifier is used. In
/// vanilla, these are mainly used for effects.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct AttributeModifier {
    /// ID is the unique ID of the modifier. It is used to identify the modifier in the packet.
    pub id: String,
    /// Name is the name of the attribute that is modified.
    pub name: String,
    /// Amount is the amount of difference between the current value of the attribute and the new value.
    pub amount: wire::F32LE,
    /// Operation is the operation that is performed on the attribute. It can be addition, multiply
    /// base, multiply total or cap.
    pub operation: wire::I32LE,
    /// Operand ... TODO: Figure out what this field is used for.
    pub operand: wire::I32LE,
    /// Serializable ... TODO: Figure out what this field is used for.
    pub is_serializable: bool,
}

// Domain: attribute_layer

/// AttributeData represents a polymorphic attribute value.
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
pub enum AttributeLayerSyncData {
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

impl AttributeLayerSyncData {
    pub fn discriminant(&self) -> u32 {
        match self {
            Self::UpdateAttributeLayersData { .. } => 0,
            Self::UpdateAttributeLayerSettingsData { .. } => 1,
            Self::UpdateEnvironmentAttributesData { .. } => 2,
            Self::RemoveEnvironmentAttributesData { .. } => 3,
        }
    }
}

impl Default for AttributeLayerSyncData {
    fn default() -> Self {
        Self::UpdateAttributeLayersData {
            attribute_layers: Default::default(),
        }
    }
}

// Domain: bedrock_profile

#[derive(Clone, Debug, Default, PartialEq)]
pub struct BedrockProfileWhiskerDiagnosticsScopeDataSummary {
    pub label: String,
    pub indentation: String,
    pub total_high_cost_ns: wire::U64LE,
    pub total_mid_cost_ns: wire::U64LE,
    pub total_low_cost_ns: wire::U64LE,
}

// Domain: bedrock_safety

#[derive(Clone, Debug, Default, PartialEq)]
pub struct BedrockSafetyRedactableString {
    pub unredacted: String,
    pub redacted: String,
}

// Domain: biome

/// BiomeCappedSurface specifies the materials to use for the capped surface of a biome, such as in
/// the Nether.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct BiomeCappedSurfaceData {
    /// FloorBlocks is a list of runtime IDs to use for the floor blocks.
    pub floor_blocks: Vec<wire::U32LE>,
    /// CeilingBlocks is a list of runtime IDs to use for the ceiling blocks.
    pub ceiling_blocks: Vec<wire::U32LE>,
    /// SeaBlock is an optional runtime ID to use for the sea block.
    /// Wire presence: optional value is preceded by a presence marker.
    pub sea_block: Option<wire::U32LE>,
    /// FoundationBlock is an optional runtime ID to use for the foundation block.
    /// Wire presence: optional value is preceded by a presence marker.
    pub foundation_block: Option<wire::U32LE>,
    /// BeachBlock is an optional runtime ID to use for the beach block.
    /// Wire presence: optional value is preceded by a presence marker.
    pub beach_block: Option<wire::U32LE>,
}

/// BiomeClimate represents the climate of a biome, mainly for ambience but also defines certain
/// behaviours.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct BiomeClimateData {
    /// Temperature is the temperature of the biome, used for weather, biome behaviours and sky colour.
    pub temperature: wire::F32LE,
    /// Downfall is the amount that precipitation affects colours and block changes.
    pub downfall: wire::F32LE,
    /// SnowAccumulationMin is the minimum amount of snow that can accumulate in the biome, every 0.125
    /// is another layer of snow.
    pub snow_accumulation_min: wire::F32LE,
    /// SnowAccumulationMax is the maximum amount of snow that can accumulate in the biome, every 0.125
    /// is another layer of snow.
    pub snow_accumulation_max: wire::F32LE,
}

/// BiomeConditionalTransformation is the legacy method of transforming biomes.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct BiomeConditionalTransformationData {
    pub transforms_into: Vec<BiomeWeightedData>,
    /// ConditionJSON is an index of the condition JSON data in the string list.
    pub condition_json: wire::U16LE,
    pub min_passing_neighbors: wire::U32LE,
}

/// BiomeConsolidatedFeature represents a feature that is consolidated into a single feature for the
/// biome.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct BiomeConsolidatedFeatureData {
    /// Scatter defines how the feature is scattered in the biome.
    pub scatter: BiomeScatterParamData,
    /// Feature is the index of the feature's name in the string list.
    pub feature: wire::U16LE,
    /// Identifier is the index of the feature's identifier in the string list.
    pub identifier: wire::U16LE,
    /// Pass is the index of the feature's pass in the string list.
    pub pass: wire::U16LE,
    pub can_use_internal_feature: bool,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct BiomeConsolidatedFeaturesData {
    pub features: Vec<BiomeConsolidatedFeatureData>,
}

/// BiomeCoordinate specifies coordinate rules for where features can be scattered in the biome.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct BiomeCoordinateData {
    /// MinValueType is the type of expression operation to use for the minimum value, and is one of the
    /// BiomeExpressionOp constants above.
    pub min_value_type: wire::ZigZag32,
    /// MinValue is the index of the minimum value expression in the string list.
    pub min_value: wire::U16LE,
    /// MaxValueType is the type of expression operation to use for the maximum value, and is one of the
    pub max_value_type: wire::ZigZag32,
    /// MaxValue is the index of the maximum value expression in the string list.
    pub max_value: wire::U16LE,
    /// GridOffset is the offset of the grid, used for fixed grid and jittered grid distributions.
    pub grid_offset: wire::U32LE,
    /// GridStepSize is the step size of the grid, used for fixed grid and jittered grid distributions.
    pub grid_step_size: wire::U32LE,
    /// Distribution is the type of distribution to use for the coordinate, and is one of the
    /// BiomeRandomDistributionType constants above.
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

/// BiomeDefinition represents a biome definition in the game. This can be a vanilla biome or a
/// completely custom biome.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct BiomeDefinitionData {
    pub id: wire::U16LE,
    /// Temperature is the temperature of the biome, used for weather, biome behaviours and sky colour.
    pub temperature: wire::F32LE,
    /// Downfall is the amount that precipitation affects colours and block changes.
    pub downfall: wire::F32LE,
    /// FoliageSnow is the progression factor for foliage turning white due to snow.
    pub foliage_snow: wire::F32LE,
    /// Depth is the depth of the biome.
    pub depth: wire::F32LE,
    /// Scale is the scale of the biome.
    pub scale: wire::F32LE,
    pub map_water_color_argb: wire::I32LE,
    /// Rain is true if the biome has rain, false if it is a dry biome.
    pub rain: bool,
    /// Tags are a list of indices of tags in the string list. These are used to group biomes together
    /// for biome generation and other purposes.
    /// Wire presence: optional value is preceded by a presence marker.
    pub tags: Option<BiomeTagsData>,
    /// Wire presence: optional value is preceded by a presence marker.
    pub chunk_gen_data: Option<BiomeDefinitionChunkGenData>,
}

/// BiomeElementData are set rules to adjust the surface materials of the biome.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct BiomeElementData {
    pub noise_freq_scale: wire::F32LE,
    /// NoiseLowerBound is the minimum noise value required to be selected.
    pub noise_lower_bound: wire::F32LE,
    /// NoiseUpperBound is the maximum noise value required to be selected.
    pub noise_upper_bound: wire::F32LE,
    /// HeightMinType is the type of expression operation to use for the minimum height, and is one of
    /// the BiomeExpressionOp constants above.
    pub height_min_type: wire::ZigZag32,
    /// HeightMin is the index of the minimum height expression in the string list.
    pub height_min: wire::U16LE,
    /// HeightMaxType is the type of expression operation to use for the maximum height, and is one of
    /// the BiomeExpressionOp constants above.
    pub height_max_type: wire::ZigZag32,
    /// HeightMax is the index of the maximum height expression in the string list.
    pub height_max: wire::U16LE,
    /// AdjustedMaterials is the materials to use for the surface layers of the biome if selected.
    pub adjusted_materials: BiomeSurfaceMaterialData,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct BiomeLegacyWorldGenRulesData {
    pub legacy_pre_hills_edge: Vec<BiomeConditionalTransformationData>,
}

/// BiomeMesaSurface specifies the materials to use for the mesa biome.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct BiomeMesaSurfaceData {
    /// ClayMaterial is the runtime ID of the block to use for clay layers.
    pub clay_material: wire::U32LE,
    /// HardClayMaterial is the runtime ID of the block to use for hard clay layers.
    pub hard_clay_material: wire::U32LE,
    /// BrycePillars is true if the biome has bryce pillars, which are tall spire-like structures.
    pub bryce_pillars: bool,
    /// HasForest is true if the biome has a forest.
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

/// BiomeNoiseGradientSurface specifies noise-gradient surface block data for a biome.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct BiomeNoiseGradientSurfaceData {
    /// NonReplaceableBlocks is a list of block runtime IDs that may not be replaced.
    pub non_replaceable_blocks: Vec<wire::U32LE>,
    /// GradientBlocks is a list of noise block specifiers used by the gradient.
    pub gradient_blocks: Vec<SerializedNoiseBlockSpecifier>,
    /// Noise is the noise descriptor used by the gradient.
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

/// BiomeReplacementData represents data for biome replacements.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct BiomeReplacementData {
    pub replacement_biome: wire::U16LE,
    /// Dimension is the dimension ID where the replacement applies.
    pub dimension: wire::U16LE,
    /// TargetBiomes is a list of target biome IDs for the replacement.
    pub target_biomes: Vec<wire::U16LE>,
    /// Amount is the amount of replacement to apply.
    pub amount: wire::F32LE,
    /// NoiseFrequencyScale ...
    pub noise_frequency_scale: wire::F32LE,
    /// ReplacementIndex is the index of the replacement.
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

/// BiomeSurfaceBuilder specifies the materials and special surface rules to use for a biome
/// surface.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct BiomeSurfaceBuilderData {
    /// SurfaceMaterials is a set of materials to use for the surface layers of the biome.
    /// Wire presence: optional value is preceded by a presence marker.
    pub surface_materials: Option<BiomeSurfaceMaterialData>,
    /// HasDefaultOverworldSurface is true if the biome has a default overworld surface.
    pub has_default_overworld_surface: bool,
    /// HasSwampSurface is true if the biome has a swamp surface.
    pub has_swamp_surface: bool,
    /// HasFrozenOceanSurface is true if the biome has a frozen ocean surface.
    pub has_frozen_ocean_surface: bool,
    /// HasEndSurface is true if the biome has an end surface.
    pub has_the_end_surface: bool,
    /// MesaSurface is optional information to specify the biome's mesa surface.
    /// Wire presence: optional value is preceded by a presence marker.
    pub mesa_surface: Option<BiomeMesaSurfaceData>,
    /// CappedSurface is optional information to specify the biome's capped surface, i.e. in the Nether.
    /// Wire presence: optional value is preceded by a presence marker.
    pub capped_surface: Option<BiomeCappedSurfaceData>,
    /// NoiseGradientSurface is optional information to specify noise-gradient surface data.
    /// Wire presence: optional value is preceded by a presence marker.
    pub noise_gradient_surface: Option<BiomeNoiseGradientSurfaceData>,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct BiomeSurfaceMaterialAdjustmentData {
    pub adjustments: Vec<BiomeElementData>,
}

/// BiomeSurfaceMaterial specifies the materials to use for the surface layers of the biome.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct BiomeSurfaceMaterialData {
    /// TopBlock is the runtime ID of the block to use for the top layer.
    pub top_block: wire::U32LE,
    /// MidBlock is the runtime ID to use for the middle layers.
    pub mid_block: wire::U32LE,
    /// SeaFloorBlock is the runtime ID to use for the sea floor.
    pub sea_floor_block: wire::U32LE,
    /// FoundationBlock is the runtime ID to use for the foundation layers.
    pub foundation_block: wire::U32LE,
    /// SeaBlock is the runtime ID to use for the sea layers.
    pub sea_block: wire::U32LE,
    /// SeaFloorDepth is the depth of the sea floor, in blocks.
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

/// FloatRange is an inclusive minimum/maximum pair of float32 values.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct FloatRange {
    /// Min is the minimum value of the range.
    pub min: wire::F32LE,
    /// Max is the maximum value of the range.
    pub max: wire::F32LE,
}

/// NoiseDescriptor describes the gradient noise used by a BiomeNoiseGradientSurface.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct NoiseDescriptor {
    /// Name is the string used to initialise the noise.
    pub name: String,
    /// FirstOctave is the first octave used by the noise.
    pub first_octave: wire::I32LE,
    /// Amplitudes is a list of amplitude values used by the noise. It must contain between 1 and 100
    /// entries.
    pub amplitudes: Vec<wire::F32LE>,
}

// Domain: block_pos

/// BlockPos is the position of a block. It is composed of three integers, and is typically written
/// as either 3 varint32s or a varint32, varuint32 and varint32.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct BlockPos {
    pub x: wire::ZigZag32,
    pub y: wire::ZigZag32,
    pub z: wire::ZigZag32,
}

// Domain: camera

/// CameraAimAssistActorPriorityData represents priority data for aim assist actor targeting.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct CameraAimAssistActorPriorityData {
    /// PresetIndex is the index of the aim assist preset.
    pub preset_index: wire::I32LE,
    /// CategoryIndex is the index of the aim assist category.
    pub category_index: wire::I32LE,
    /// ActorIndex is the index of the actor.
    pub actor_index: wire::I32LE,
    /// Priority is the priority value for this actor.
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

/// CameraEase represents an easing function that can be used by a CameraInstructionSet.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct CameraEase {
    /// Type is the type of easing function used. This is one of the constants above.
    pub type_: wire::U8,
    /// Duration is the time in seconds that the easing function should take.
    pub time: wire::F32LE,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct CameraEntityOffset {
    pub entity_offset_x: wire::F32LE,
    pub entity_offset_y: wire::F32LE,
    pub entity_offset_z: wire::F32LE,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct CameraFacing {
    pub pos: glam::Vec3,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct CameraFadeColor {
    pub red: wire::F32LE,
    pub green: wire::F32LE,
    pub blue: wire::F32LE,
}

/// CameraFadeTimeData represents the time data for a CameraInstructionFade.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct CameraFadeTimeData {
    /// FadeInDuration is the time in seconds for the screen to fully fade in.
    pub fade_in_time: wire::F32LE,
    /// WaitDuration is time in seconds to wait before fading out.
    pub hold_time: wire::F32LE,
    /// FadeOutDuration is the time in seconds for the screen to fully fade out.
    pub fade_out_time: wire::F32LE,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct CameraInstructionData {
    /// Wire presence: optional value is preceded by a presence marker.
    pub set: Option<CameraInstructionSet>,
    /// Wire presence: optional value is preceded by a presence marker.
    pub clear: Option<bool>,
    /// Wire presence: optional value is preceded by a presence marker.
    pub fade: Option<CameraInstructionFade>,
    /// Wire presence: optional value is preceded by a presence marker.
    pub target: Option<CameraInstructionTargetData>,
    /// Wire presence: optional value is preceded by a presence marker.
    pub remove_target: Option<bool>,
    /// Wire presence: optional value is preceded by a presence marker.
    pub field_of_view: Option<CameraInstructionFieldOfView>,
    /// Wire presence: optional value is preceded by a presence marker.
    pub spline: Option<CameraSplineInstruction>,
    /// Wire presence: optional value is preceded by a presence marker.
    pub attach_to_entity: Option<CameraInstructionTarget>,
    /// Wire presence: optional value is preceded by a presence marker.
    pub detach_from_entity: Option<bool>,
}

/// CameraInstructionFade represents a camera instruction that fades the screen to a specified
/// colour.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct CameraInstructionFade {
    /// TimeData is the time data for the fade, which includes the fade in duration, wait duration and
    /// fade out duration.
    /// Wire presence: optional value is preceded by a presence marker.
    pub time: Option<CameraFadeTimeData>,
    /// Colour is the colour of the screen to fade to. This only uses the red, green and blue
    /// components.
    /// Wire presence: optional value is preceded by a presence marker.
    pub color: Option<CameraFadeColor>,
}

/// CameraInstructionFieldOfView represents a camera instruction that updates the field of view.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct CameraInstructionFieldOfView {
    /// FieldOfView is the field of view of the camera.
    pub field_of_view: wire::F32LE,
    pub fov_ease_time: wire::F32LE,
    pub fov_ease_type: String,
    pub field_of_view_clear: bool,
}

/// CameraInstructionSet represents a camera instruction that sets the camera to a specified preset
/// and can be extended with easing functions and translations to the camera's position and
/// rotation.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct CameraInstructionSet {
    /// Preset is the index of the preset in the CameraPresets packet sent to the player.
    pub preset: wire::U32LE,
    /// Ease represents the easing function that is used by the instruction.
    /// Wire presence: optional value is preceded by a presence marker.
    pub ease: Option<CameraEase>,
    /// Position represents the position of the camera.
    /// Wire presence: optional value is preceded by a presence marker.
    pub pos: Option<CameraPosition>,
    /// Rotation represents the rotation of the camera.
    /// Wire presence: optional value is preceded by a presence marker.
    pub rot: Option<CameraRotation>,
    /// Facing is a vector that the camera will always face towards during the duration of the
    /// instruction.
    /// Wire presence: optional value is preceded by a presence marker.
    pub facing: Option<CameraFacing>,
    /// ViewOffset is an offset based on a pivot point to the player, causing the camera to be shifted
    /// in a certain direction.
    /// Wire presence: optional value is preceded by a presence marker.
    pub view_offset: Option<CameraViewOffset>,
    /// EntityOffset is an offset from the entity that the camera should be rendered at.
    /// Wire presence: optional value is preceded by a presence marker.
    pub entity_offset: Option<CameraEntityOffset>,
    /// Default determines whether the camera is a default camera or not.
    /// Wire presence: optional value is preceded by a presence marker.
    pub default: Option<bool>,
    /// IgnoreStartingValuesComponent behavior is currently unknown.
    pub remove_ignore_starting_values_component: bool,
}

/// CameraInstructionTarget represents a camera instruction that targets a specific entity.
#[derive(Clone, Copy, Debug, Default, PartialEq, Eq, Hash)]
pub struct CameraInstructionTarget(pub i64);

impl wire::WireCodec for CameraInstructionTarget {
    fn encode<W: std::io::Write>(&self, writer: &mut W) -> std::io::Result<()> {
        <wire::I64LE as wire::WireCodec>::encode(&wire::I64LE(self.0), writer)
    }

    fn decode<R: std::io::Read>(reader: &mut R) -> std::io::Result<Self> {
        <wire::I64LE as wire::WireCodec>::decode(reader).map(|value| Self(value.0))
    }
}

/// CameraInstructionTarget represents a camera instruction that targets a specific entity.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct CameraInstructionTargetData {
    /// CenterOffset is the offset from the center of the entity that the camera should target.
    /// Wire presence: optional value is preceded by a presence marker.
    pub target_center_offset: Option<glam::Vec3>,
    /// EntityUniqueID is the unique ID of the entity that the camera should target.
    pub target_actor_id: wire::I64LE,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct CameraPosition {
    pub pos: glam::Vec3,
}

/// CameraPreset represents a basic preset that can be extended upon by more complex instructions.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct CameraPreset {
    /// Name is the name of the preset. Each preset must have their own unique name.
    pub name: String,
    /// Parent is the name of the preset that this preset extends upon. This can be left empty.
    pub inherit_from: String,
    /// PosX is the default X position of the camera.
    /// Wire presence: optional value is preceded by a presence marker.
    pub pos_x: Option<wire::F32LE>,
    /// PosY is the default Y position of the camera.
    /// Wire presence: optional value is preceded by a presence marker.
    pub pos_y: Option<wire::F32LE>,
    /// PosZ is the default Z position of the camera.
    /// Wire presence: optional value is preceded by a presence marker.
    pub pos_z: Option<wire::F32LE>,
    /// RotX is the default pitch of the camera.
    /// Wire presence: optional value is preceded by a presence marker.
    pub rot_x: Option<wire::F32LE>,
    /// RotY is the default yaw of the camera.
    /// Wire presence: optional value is preceded by a presence marker.
    pub rot_y: Option<wire::F32LE>,
    /// RotationSpeed is the speed at which the camera should rotate.
    /// Wire presence: optional value is preceded by a presence marker.
    pub rotation_speed: Option<wire::F32LE>,
    /// SnapToTarget determines whether the camera should snap to the target entity or not.
    /// Wire presence: optional value is preceded by a presence marker.
    pub snap_to_target: Option<bool>,
    /// HorizontalRotationLimit is the horizontal rotation limit of the camera.
    /// Wire presence: optional value is preceded by a presence marker.
    pub horizontal_rotation_limit: Option<glam::Vec2>,
    /// VerticalRotationLimit is the vertical rotation limit of the camera.
    /// Wire presence: optional value is preceded by a presence marker.
    pub vertical_rotation_limit: Option<glam::Vec2>,
    /// ContinueTargeting determines whether the camera should continue targeting when using aim assist.
    /// Wire presence: optional value is preceded by a presence marker.
    pub continue_targeting: Option<bool>,
    /// TrackingRadius is the radius around the camera that the aim assist should track targets.
    /// Wire presence: optional value is preceded by a presence marker.
    pub block_listening_radius: Option<wire::F32LE>,
    /// ViewOffset is only used in a follow_orbit camera and controls an offset based on a pivot point
    /// to the player, causing it to be shifted in a certain direction.
    /// Wire presence: optional value is preceded by a presence marker.
    pub view_offset: Option<glam::Vec2>,
    /// EntityOffset controls the offset from the entity that the camera should be rendered at.
    /// Wire presence: optional value is preceded by a presence marker.
    pub entity_offset: Option<glam::Vec3>,
    /// Radius is only used in a follow_orbit camera and controls how far away from the player the
    /// camera should be rendered.
    /// Wire presence: optional value is preceded by a presence marker.
    pub radius: Option<wire::F32LE>,
    /// MinYawLimit is the minimum yaw limit of the camera.
    /// Wire presence: optional value is preceded by a presence marker.
    pub yaw_limit_min: Option<wire::F32LE>,
    /// MaxYawLimit is the maximum yaw limit of the camera.
    /// Wire presence: optional value is preceded by a presence marker.
    pub yaw_limit_max: Option<wire::F32LE>,
    /// AudioListener defines where the audio should be played from when using this preset. This is one
    /// of the constants above.
    /// Wire presence: optional value is preceded by a presence marker.
    pub listener: Option<CameraPresetAudioListener>,
    /// PlayerEffects is currently unknown.
    /// Wire presence: optional value is preceded by a presence marker.
    pub player_effects: Option<bool>,
    /// AimAssist defines the aim assist to use when using this preset.
    /// Wire presence: optional value is preceded by a presence marker.
    pub aim_assist: Option<CameraAimAssistCommandPresetDefinition>,
    /// ControlScheme is the control scheme that the client should use in this camera. It is one of the
    /// following: - ControlSchemeLockedPlayerRelativeStrafe is the default behaviour, this cannot be
    /// set when the client is in a custom camera. - ControlSchemeCameraRelative makes movement relative
    /// to the camera's transform, with the client's rotation being relative to the client's movement. -
    /// ControlSchemeCameraRelativeStrafe makes movement relative to the camera's transform, with the
    /// client's rotation being locked. - ControlSchemePlayerRelative makes movement relative to the
    /// player's transform, meaning holding left/right will make the player turn in a circle. -
    /// ControlSchemePlayerRelativeStrafe makes movement the same as the default behaviour, but can be
    /// used in a custom camera.
    /// Wire presence: optional value is preceded by a presence marker.
    pub control_scheme: Option<ControlScheme>,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct CameraPresetList {
    pub presets: Vec<CameraPreset>,
}

/// CameraProgressOption represents a progress keyframe option for camera spline instructions.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct CameraProgressOption {
    pub key_frame_value: wire::F32LE,
    pub key_frame_time: wire::F32LE,
    pub key_frame_easing_func: String,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct CameraRotation {
    pub x: wire::F32LE,
    pub y: wire::F32LE,
}

/// CameraRotationOption represents a rotation option for camera spline instructions.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct CameraRotationOption {
    pub key_frame_value: glam::Vec3,
    pub key_frame_time: wire::F32LE,
    pub key_frame_easing_func: String,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct CameraSplineControlPoint {
    pub position: glam::Vec3,
}

/// CameraSplineDefinition represents a named camera spline definition.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct CameraSplineDefinition {
    /// Name is the name of the spline definition.
    pub name: String,
    /// TotalTime is the total time for the spline animation.
    pub total_time: wire::F32LE,
    /// SplineType is the optional spline interpolation type.
    pub spline_type: String,
    /// ControlPoints is a list of points that define the spline curve.
    pub control_points: Vec<CameraSplineControlPoint>,
    /// ProgressKeyFrames is a list of progress key frames for the spline.
    pub progress_key_frames: Vec<CameraSplineProgressKeyFrame>,
    /// RotationKeyFrames is a list of rotation key frames for the spline.
    pub rotation_key_frames: Vec<CameraSplineRotationKeyFrame>,
}

/// CameraSplineInstruction represents a camera instruction that creates a spline path for the
/// camera to follow.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct CameraSplineInstruction {
    /// TotalTime is the total time for the spline animation.
    pub total_time: wire::F32LE,
    pub type_: wire::U8,
    /// Curve is a list of points that define the spline curve.
    pub curve: Vec<glam::Vec3>,
    /// ProgressKeyFrames is a list of progress key frames for the spline.
    pub progress_key_frames: Vec<CameraProgressOption>,
    pub rotation_option: Vec<CameraRotationOption>,
    /// SplineIdentifier is an optional identifier for referencing the spline by name.
    pub spline_identifier: String,
    /// LoadFromJson optionally determines whether the spline should be loaded from a JSON definition.
    pub load_from_json: bool,
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

#[derive(Clone, Debug, Default, PartialEq)]
pub struct CameraViewOffset {
    pub x: wire::F32LE,
    pub y: wire::F32LE,
}

// Domain: chunk_pos

/// ChunkPos is the position of a chunk. It is composed of two integers and is written as two
/// varint32s.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct ChunkPos {
    pub x: wire::ZigZag32,
    pub z: wire::ZigZag32,
}

/// SubChunkPos is the position of a sub-chunk. The X and Z coordinates are the coordinates of the
/// chunk, and the Y coordinate is the absolute sub-chunk index.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct SubChunkPos {
    pub subchunk_position_x: wire::I32LE,
    pub subchunk_position_y: wire::I32LE,
    pub subchunk_position_z: wire::I32LE,
}

// Domain: clock

/// SyncWorldClockStateData represents the state data for synchronising a world clock.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct SyncWorldClockStateData {
    /// ClockID is the unique identifier for the clock.
    pub clock_id: wire::VarULong,
    /// Time is the current time of the clock.
    pub time: wire::ZigZag32,
    /// Paused indicates if the clock is paused.
    pub is_paused: bool,
}

/// TimeMarkerData represents a time marker within a world clock.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct TimeMarkerData {
    /// ID is the unique identifier for the time marker.
    pub id: wire::VarULong,
    /// Name is the name of the time marker.
    pub name: String,
    /// Time is the time at which the marker is set.
    pub time: wire::ZigZag32,
    /// Period is the optional period for the time marker.
    /// Wire presence: optional value is preceded by a presence marker.
    pub period: Option<wire::I32LE>,
}

/// WorldClockData represents a complete world clock with its time markers.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct WorldClockData {
    /// ID is the unique identifier for the clock.
    pub id: wire::VarULong,
    /// Name is the name of the clock.
    pub name: String,
    /// Time is the current time of the clock.
    pub time: wire::ZigZag32,
    /// Paused indicates if the clock is paused.
    pub is_paused: bool,
    /// TimeMarkers is a list of time markers for this clock.
    pub time_markers: Vec<TimeMarkerData>,
}

// Domain: command

/// ChainedSubcommand represents a subcommand that can have chained commands, such as /execute which
/// allows you to run another command as another entity or at a different position etc.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct ChainedSubcommand {
    /// Name is the name of the chained subcommand and shows up in the list as a regular subcommand
    /// enum.
    pub name: String,
    /// Values contains the index and parameter type of the chained subcommand.
    pub sub_command_values: Vec<ChainedSubcommandValue>,
}

/// ChainedSubcommandValue represents the value for a chained subcommand argument.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct ChainedSubcommandValue {
    /// Index is the index of the argument in the ChainedSubcommandValues slice from the
    /// AvailableCommands packet. This is then used to set the type specified by the Value field below.
    pub sub_command_first_value: wire::VarUInt,
    /// Value is a combination of the flags above and specified the type of argument. Unlike regular
    /// parameter types, this should NOT contain any of the special flags (valid, enum, suffixed or soft
    /// enum) but only the basic types.
    pub sub_command_second_value: wire::VarUInt,
}

/// Command holds the data that a command requires to be shown to a player client-side. The command
/// is shown in the /help command and auto-completed using this data.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct Command {
    /// Name is the name of the command. The command may be executed using this name, and will be shown
    /// in the /help list with it. It currently seems that the client crashes if the Name contains
    /// uppercase letters.
    pub name: String,
    /// Description is the description of the command. It is shown in the /help list and when starting
    /// to write a command.
    pub description: String,
    /// Flags is a combination of flags not currently known. Leaving the Flags field empty appears to
    /// work.
    pub flags: wire::U16LE,
    /// PermissionLevel is the command permission level that the player required to execute this
    /// command. The field no longer seems to serve a purpose, as the client does not handle the
    /// execution of commands anymore: The permissions should be checked server-side.
    pub permission_level: String,
    pub alias_enum: wire::I32LE,
    pub command_data_chained_subcommand_indexes: Vec<wire::U32LE>,
    /// Overloads is a list of command overloads that specify the ways in which a command may be
    /// executed. The overloads may be completely different.
    pub overloads: Vec<CommandOverload>,
}

#[derive(Clone, Debug, PartialEq)]
pub enum CommandBlockUpdateData {
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

impl CommandBlockUpdateData {
    pub fn discriminant(&self) -> u32 {
        match self {
            Self::EntityCommandTarget { .. } => 0,
            Self::BlockCommandData { .. } => 1,
        }
    }
}

impl Default for CommandBlockUpdateData {
    fn default() -> Self {
        Self::EntityCommandTarget {
            target_runtime_id: Default::default(),
        }
    }
}

/// CommandEnum represents an enum in a command usage. The enum typically has a type and a set of
/// options that are valid. A value that is not one of the options results in a failure during
/// execution.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct CommandEnum {
    /// Type is the type of the command enum. The type will show up in the command usage as the type of
    /// the argument if it has a certain amount of arguments, or when Options is set to true in the
    /// command holding the enum.
    pub name: String,
    /// ValueIndices holds a list of indices that point to the EnumValues slice in the
    /// AvailableCommandsPacket. These represent the options of the enum.
    pub values: Vec<wire::U32LE>,
}

/// CommandEnumConstraint is sent in the AvailableCommands packet to limit what values of an enum
/// may be used taking in account things such as whether cheats are enabled.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct CommandEnumConstraint {
    /// EnumValueIndex points to an enum value in the AvailableCommands packet that this constraint
    /// should apply to.
    pub enum_value_symbol: wire::U32LE,
    /// EnumIndex points to an enum in the AvailableCommands packet to which this constraint should
    /// apply to.
    pub enum_symbol: wire::U32LE,
    /// Constraints holds a slice of constraints as present in the constants above.
    pub constraint_indices: Vec<wire::U8>,
}

/// CommandOrigin holds data that identifies the origin of the requesting of a command. It holds
/// several fields that may be used to get specific information. When sent in a CommandRequest
/// packet, the same CommandOrigin should be sent in a CommandOutput packet.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct CommandOriginData {
    pub type_: String,
    /// UUID is a unique identifier for every instantiation of a command.
    pub uuid: uuid::Uuid,
    /// RequestID is an ID that identifies the request of the client. The server should send a
    /// CommandOrigin with the same request ID to ensure it can be matched with the request by the
    /// caller of the command. This is especially important for websocket servers and it seems that this
    /// field is only non-empty for these websocket servers.
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

/// CommandOutputMessage represents a message sent by a command that holds the output of one of the
/// commands executed.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct CommandOutputMessage {
    pub message_id: String,
    pub successful: bool,
    /// Parameters is a list of parameters that serve to supply the message sent with additional
    /// information, such as the position that a player was teleported to or the effect that was applied
    /// to an entity. These parameters only apply for the Minecraft built-in command output.
    pub parameters: Vec<String>,
}

/// CommandOverload represents an overload of a command. This overload can be compared to function
/// overloading in languages such as java. It represents a single usage of the command. A command
/// may have multiple different overloads, which are handled differently.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct CommandOverload {
    /// Chaining determines if the parameters use chained subcommands or not.
    pub is_chaining: bool,
    /// Parameters is a list of command parameters that are part of the overload. These parameters
    /// specify the usage of the command when this overload is applied.
    pub parameter_data: Vec<CommandParameter>,
}

/// CommandParameter represents a single parameter of a command overload, which accepts a certain
/// type of input values. It has a name and a type which show up client-side when a player is
/// entering the command.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct CommandParameter {
    /// Name is the name of the command parameter. It shows up in the usage like <$Name: $Type>, with
    /// the exception of enum types, which show up simply as a list of options if the list is short
    /// enough and Options is set to false.
    pub name: String,
    pub parse_symbol: wire::U32LE,
    pub is_optional: bool,
    /// Options holds a combinations of options that additionally apply to the command parameter. The
    /// list of options can be found above.
    pub options: wire::U8,
}

/// DynamicEnum is an enum variant that can have its options changed during runtime, without sending
/// a new AvailableCommands packet.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct DynamicEnum {
    /// Type is the type of the command enum. The type will show up in the command usage as the type of
    /// the argument if it has a certain amount of arguments, or when Options is set to true in the
    /// command holding the enum.
    pub enum_name: String,
    /// Values is a slice of possible options for the enum.
    pub enum_options: Vec<String>,
}

// Domain: container

#[derive(Clone, Debug, Default, PartialEq)]
pub struct ContainerMixDataEntry {
    pub from_item_id: wire::ZigZag32,
    pub reagent_item_id: wire::ZigZag32,
    pub to_item_id: wire::ZigZag32,
}

/// FullContainerName contains information required to identify a container in a
/// StackRequestSlotInfo.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct FullContainerName {
    /// ContainerID is the ID of the container that the slot was in.
    pub container_name: ContainerEnumName,
    /// DynamicContainerID is the ID of the container if it is dynamic. If the container is not dynamic,
    /// this field should be left empty. A non-optional value of 0 is assumed to be non-empty.
    /// Wire presence: optional value is preceded by a presence marker.
    pub dynamic_id: Option<wire::U32LE>,
}

// Domain: creative

/// CreativeGroup represents a group of items in the creative inventory. Each group has a category,
/// name and an icon that represents the group.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct CreativeGroupInfo {
    /// Category is the category the group falls under. It is one of the constants above.
    pub creative_category: CreativeItemCategory,
    /// Name is the locale name of the group, i.e. "itemGroup.name.planks".
    pub name: String,
    /// Icon is the item that represents the group in the creative inventory.
    pub group_icon_item: NetworkItemInstanceDescriptorSerializedData,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct CreativeItemEntry {
    pub creative_net_id: CreativeItemNetID,
    pub item_instance: NetworkItemInstanceDescriptorSerializedData,
    pub group_index: wire::VarUInt,
}

#[derive(Clone, Copy, Debug, Default, PartialEq, Eq, Hash)]
pub struct CreativeItemNetID(pub u32);

impl wire::WireCodec for CreativeItemNetID {
    fn encode<W: std::io::Write>(&self, writer: &mut W) -> std::io::Result<()> {
        <wire::VarUInt as wire::WireCodec>::encode(&wire::VarUInt(self.0), writer)
    }

    fn decode<R: std::io::Read>(reader: &mut R) -> std::io::Result<Self> {
        <wire::VarUInt as wire::WireCodec>::decode(reader).map(|value| Self(value.0))
    }
}

// Domain: education

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

// Domain: enchant

/// EnchantmentInstance represents a single enchantment instance with the type of the enchantment
/// and its level.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct EnchantmentInstance {
    pub enchant_type: EnchantType,
    pub enchant_level: wire::U8,
}

// Domain: entity

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

// Domain: entity_link

/// EntityLink is a link between two entities, typically being one entity riding another.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct EntityLink {
    pub target_a: ActorUniqueID,
    pub target_b: ActorUniqueID,
    /// Type is one of the types above. It specifies the way the entity is linked to another entity.
    pub type_: ActorLinkType,
    /// Immediate is set to immediately dismount an entity from another. This should be set when the
    /// mount of an entity is killed.
    pub immediate: bool,
    pub passenger_initiated: bool,
    /// VehicleAngularVelocity is the angular velocity of the vehicle that the rider is riding.
    pub vehicle_angular_velocity: wire::F32LE,
}

// Domain: events

/// Event represents an object that holds data specific to an event. The data it holds depends on
/// the type.
#[derive(Clone, Debug, PartialEq)]
pub enum EventData {
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

impl EventData {
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

impl Default for EventData {
    fn default() -> Self {
        Self::Achievement {
            achievement_id: Default::default(),
        }
    }
}

// Domain: experiment

#[derive(Clone, Debug, Default, PartialEq)]
pub struct ExperimentToggle {
    pub name: String,
    pub enabled: bool,
}

// Domain: game_rule

/// GameRule contains game rule data.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct GameRule {
    /// Name is the name of the game rule.
    pub rule_name: String,
    /// CanBeModifiedByPlayer specifies if the game rule can be modified by the player through the
    /// in-game UI.
    pub rule_can_be_modified: bool,
    /// Value is the new value of the game rule. This is either a bool, uint32 or float32, or nil for
    /// the null variant, which carries no value at all.
    pub rule_value: GameRuleValue,
}

// Domain: generated

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

/// Stores the 131-bit value used by the wire bitset encoding.
#[derive(Clone, Debug, Default, PartialEq, Eq)]
pub struct Bitset131(pub [u64; 3]);

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
        value: wire::NetworkNbt,
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

#[derive(Clone, Debug, PartialEq)]
pub enum DisconnectMessages {
    DisconnectPacketMessages {
        message: String,
        filtered_message: String,
    },
    /// Naming overlay required: source placeholder `Empty1`.
    Empty,
}

impl DisconnectMessages {
    pub fn discriminant(&self) -> u32 {
        match self {
            Self::DisconnectPacketMessages { .. } => 0,
            Self::Empty => 1,
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

#[derive(Clone, Debug, PartialEq, Default)]
pub enum GameRuleValue {
    /// Naming overlay required: source placeholder `Empty0`.
    #[default]
    Empty,
    Bool(bool),
    Int32(wire::I32LE),
    Float(wire::F32LE),
}

impl GameRuleValue {
    pub fn discriminant(&self) -> u32 {
        match self {
            Self::Empty => 0,
            Self::Bool(..) => 1,
            Self::Int32(..) => 2,
            Self::Float(..) => 3,
        }
    }
}

#[derive(Clone, Debug, PartialEq)]
pub enum InventoryTransactionValue {
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
        item: NetworkItemStackDescriptorSerializedData,
        from_position: glam::Vec3,
        hit_position: glam::Vec3,
    },
    ItemReleaseInventoryTransaction {
        actions: InventoryTransactionData,
        action_type: ItemReleaseInventoryTransactionActionType,
        slot: wire::ZigZag32,
        item: NetworkItemStackDescriptorSerializedData,
        from_position: glam::Vec3,
    },
}

impl InventoryTransactionValue {
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

impl Default for InventoryTransactionValue {
    fn default() -> Self {
        Self::NormalTransactionData {
            actions: Default::default(),
        }
    }
}

#[derive(Clone, Debug, PartialEq, Default)]
pub enum PrimitiveShapeExtraShapeData {
    /// Naming overlay required: source placeholder `Empty0`.
    #[default]
    Empty,
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
        /// Text is the text of the debug text shape.
        text: String,
        /// UseRotation is if the text should use the provided rotation, meaning it will be static and does
        /// not follow the camera. Use false for default behaviour.
        use_rotation: bool,
        /// BackgroundColour is the RGBA colour to use for the text background. This is a translucent black
        /// colour by default.
        /// Wire presence: optional value is preceded by a presence marker.
        background_color: Option<MceColor>,
        /// DepthTest is whether the text should show through walls. Use true for default behaviour.
        depth_test: bool,
        /// ShowBackface is if the background should render on the back side of the shape. This only has a
        /// visible effect when UseRotation is true since you cannot see the back side of the text
        /// otherwise. Use true for default behaviour.
        show_backface: bool,
        /// ShowBackfaceText is if the text should render on the back side of the shape. This only has a
        /// visible effect when UseRotation is true since you cannot see the back side of the text
        /// otherwise. Use true for default behaviour.
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

impl PrimitiveShapeExtraShapeData {
    pub fn discriminant(&self) -> u32 {
        match self {
            Self::Empty => 0,
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
pub enum SetScoreInfoItem {
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

impl SetScoreInfoItem {
    pub fn discriminant(&self) -> u8 {
        match self {
            Self::RemoveScore { .. } => 0,
            Self::ChangePlayerScore { .. } => 1,
            Self::ChangeEntityScore { .. } => 2,
            Self::ChangeFakePlayerScore { .. } => 3,
        }
    }
}

impl Default for SetScoreInfoItem {
    fn default() -> Self {
        Self::RemoveScore {
            action: Default::default(),
            scoreboard_id: Default::default(),
            objective_name: Default::default(),
        }
    }
}

// Domain: inventory

/// InventoryAction represents a single action that took place during an inventory transaction. On
/// itself, this inventory action is always unbalanced: It must be combined with other actions in an
/// inventory transaction to form a balanced transaction.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct InventoryAction {
    pub source: InventorySource,
    pub slot: wire::VarUInt,
    pub from_item: NetworkItemStackDescriptorSerializedData,
    pub to_item: NetworkItemStackDescriptorSerializedData,
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

/// InventoryTransactionData represents an object that holds data specific to an inventory
/// transaction type. The data it holds depends on the type.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct InventoryTransactionData {
    /// Wire presence: optional value is preceded by a presence marker.
    pub actions: Option<Vec<InventoryAction>>,
}

// Domain: item

#[derive(Clone, Debug, Default, PartialEq)]
pub struct ItemData {
    pub item_name: String,
    pub item_id: wire::I16LE,
    pub is_component_based: bool,
    pub item_version: ItemVersion,
    pub item_component_data: wire::NetworkNbt,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct ItemEnchantOption {
    pub cost: wire::U8,
    pub enchants: ItemEnchants,
    pub enchant_name: String,
    pub enchant_net_id: RecipeNetID,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct ItemEnchants {
    pub slot: wire::I32LE,
    pub item_enchants: [Vec<EnchantmentInstance>; 3],
}

/// ItemInstance represents a unique instance of an item stack. These instances carry a specific
/// network ID that is persistent for the stack.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct ItemInstance {
    pub item_descriptor: ItemDescriptor,
    pub stack_size: wire::U16LE,
    pub block_runtime_id: wire::VarUInt,
    pub user_data_buffer: bytes::Bytes,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct ItemUseInventoryTransaction {
    pub actions: InventoryTransactionData,
    pub action_type: ItemUseInventoryTransactionActionType,
    pub trigger_type: ItemUseInventoryTransactionTriggerType,
    pub position: BlockPos,
    pub face: wire::U8,
    pub slot: wire::ZigZag32,
    pub item: NetworkItemStackDescriptorSerializedData,
    pub from_position: glam::Vec3,
    pub click_position: glam::Vec3,
    pub target_block_id: wire::VarUInt,
    pub client_interact_prediction: ItemUseInventoryTransactionPredictedResult,
    pub client_cooldown_state: ItemUseInventoryTransactionClientCooldownState,
}

// Domain: item_descriptor

/// ItemDescriptor represents a type of item descriptor. This is one of the concrete types below. It
/// is an alias of Marshaler.
#[derive(Clone, Debug, PartialEq)]
pub enum ItemDescriptor {
    EmptyItemDescriptorData {
        descriptor_type: ItemDescriptorType,
    },
    ItemNameDescriptorData {
        descriptor_type: ItemDescriptorType,
        full_name: String,
        aux_value: wire::ZigZag32,
    },
    MolangItemDescriptorData {
        descriptor_type: ItemDescriptorType,
        tag_expression: String,
        molang_version: MoLangVersion,
    },
    ItemTagDescriptorData {
        descriptor_type: ItemDescriptorType,
        item_tag: String,
    },
}

impl ItemDescriptor {
    pub fn discriminant(&self) -> u32 {
        match self {
            Self::EmptyItemDescriptorData { .. } => 0,
            Self::ItemNameDescriptorData { .. } => 1,
            Self::MolangItemDescriptorData { .. } => 2,
            Self::ItemTagDescriptorData { .. } => 3,
        }
    }
}

impl Default for ItemDescriptor {
    fn default() -> Self {
        Self::EmptyItemDescriptorData {
            descriptor_type: Default::default(),
        }
    }
}

/// ItemDescriptor represents a type of item descriptor. This is one of the concrete types below. It
/// is an alias of Marshaler.
#[derive(Clone, Debug, PartialEq)]
pub enum StackRequestAction {
    TakeActionData {
        action_type: ItemStackRequestActionType,
        amount: wire::U8,
        source: StackRequestSlotInfo,
        destination: StackRequestSlotInfo,
    },
    PlaceActionData {
        action_type: ItemStackRequestActionType,
        amount: wire::U8,
        source: StackRequestSlotInfo,
        destination: StackRequestSlotInfo,
    },
    SwapActionData {
        action_type: ItemStackRequestActionType,
        /// Source and Destination point to the source slot from which Count of the item stack were taken
        /// and the destination slot to which this item was moved.
        source: StackRequestSlotInfo,
        /// Source and Destination point to the source slot from which Count of the item stack were taken
        /// and the destination slot to which this item was moved.
        destination: StackRequestSlotInfo,
    },
    DropActionData {
        action_type: ItemStackRequestActionType,
        amount: wire::U8,
        /// Source is the source slot from which items were dropped to the ground.
        source: StackRequestSlotInfo,
        /// Randomly seems to be set to false in most cases. I'm not entirely sure what this does, but this
        /// is what vanilla calls this field.
        randomly: bool,
    },
    DestroyActionData {
        action_type: ItemStackRequestActionType,
        amount: wire::U8,
        /// Source is the source slot from which items came that were destroyed by moving them into the
        /// creative inventory.
        source: StackRequestSlotInfo,
    },
    ConsumeActionData {
        action_type: ItemStackRequestActionType,
        amount: wire::U8,
        source: StackRequestSlotInfo,
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
        /// PredictedDurability is the durability of the item that the client assumes to be present at the
        /// time.
        predicted_durability: wire::ZigZag32,
        net_id_variant: wire::I32LE,
    },
    CraftRecipeActionData {
        action_type: ItemStackRequestActionType,
        recipe_net_id: RecipeNetID,
        number_of_requested_crafts: wire::U8,
    },
    CraftRecipeAutoActionData {
        action_type: ItemStackRequestActionType,
        recipe_net_id: RecipeNetID,
        number_of_requested_crafts: wire::U8,
        /// Ingredients is a slice of ItemDescriptorCount that contains the ingredients that were used to
        /// craft the recipe. It is not exactly clear what this is used for, but it is sent by the vanilla
        /// client.
        ingredients: Vec<RecipeIngredient>,
    },
    CraftCreativeActionData {
        action_type: ItemStackRequestActionType,
        creative_item_net_id: wire::VarUInt,
        number_of_requested_crafts: wire::U8,
    },
    CraftRecipeOptionalActionData {
        action_type: ItemStackRequestActionType,
        recipe_net_id: RecipeNetID,
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
        craft_results: Vec<ItemInstance>,
        num_crafts: wire::U8,
    },
}

impl StackRequestAction {
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

impl Default for StackRequestAction {
    fn default() -> Self {
        Self::TakeActionData {
            action_type: Default::default(),
            amount: Default::default(),
            source: Default::default(),
            destination: Default::default(),
        }
    }
}

// Domain: item_stack

#[derive(Clone, Copy, Debug, Default, PartialEq, Eq, Hash)]
pub struct ItemStackLegacyRequestID(pub i32);

impl wire::WireCodec for ItemStackLegacyRequestID {
    fn encode<W: std::io::Write>(&self, writer: &mut W) -> std::io::Result<()> {
        <wire::ZigZag32 as wire::WireCodec>::encode(&wire::ZigZag32(self.0), writer)
    }

    fn decode<R: std::io::Read>(reader: &mut R) -> std::io::Result<Self> {
        <wire::ZigZag32 as wire::WireCodec>::decode(reader).map(|value| Self(value.0))
    }
}

#[derive(Clone, Copy, Debug, Default, PartialEq, Eq, Hash)]
pub struct ItemStackNetID(pub i32);

impl wire::WireCodec for ItemStackNetID {
    fn encode<W: std::io::Write>(&self, writer: &mut W) -> std::io::Result<()> {
        <wire::ZigZag32 as wire::WireCodec>::encode(&wire::ZigZag32(self.0), writer)
    }

    fn decode<R: std::io::Read>(reader: &mut R) -> std::io::Result<Self> {
        <wire::ZigZag32 as wire::WireCodec>::decode(reader).map(|value| Self(value.0))
    }
}

/// ItemStackRequest represents a single request present in an ItemStackRequest packet sent by the
/// client to change an item in an inventory. Item stack requests are either approved or rejected by
/// the server using the ItemStackResponse packet.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct ItemStackRequestData {
    pub client_request_id: ItemStackRequestID,
    /// Actions is a list of actions performed by the client. The actual type of the actions depends on
    /// which ID was present, and is one of the concrete types below.
    pub actions: Vec<StackRequestAction>,
    pub strings_to_filter: Vec<String>,
    pub strings_to_filter_origin: TextProcessingEventOrigin,
}

#[derive(Clone, Copy, Debug, Default, PartialEq, Eq, Hash)]
pub struct ItemStackRequestID(pub i32);

impl wire::WireCodec for ItemStackRequestID {
    fn encode<W: std::io::Write>(&self, writer: &mut W) -> std::io::Result<()> {
        <wire::ZigZag32 as wire::WireCodec>::encode(&wire::ZigZag32(self.0), writer)
    }

    fn decode<R: std::io::Read>(reader: &mut R) -> std::io::Result<Self> {
        <wire::ZigZag32 as wire::WireCodec>::decode(reader).map(|value| Self(value.0))
    }
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct ItemStackRequestPacketData {
    pub client_request_id: ItemStackRequestID,
    pub actions: Vec<StackRequestAction>,
    pub strings_to_filter: Vec<String>,
    pub strings_to_filter_origin: TextProcessingEventOrigin,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct ItemStackResponseContainerInfo {
    pub full_container_name: FullContainerName,
    pub slots: Vec<ItemStackResponseSlotInfo>,
}

/// ItemStackResponse is a response to an individual ItemStackRequest.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct ItemStackResponseInfo {
    pub result: ItemStackNetResult,
    pub client_request_id: ItemStackRequestID,
    /// Wire presence: optional value is preceded by a presence marker.
    pub containers: Option<Vec<ItemStackResponseContainerInfo>>,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct ItemStackResponseSlotInfo {
    pub requested_slot: wire::U8,
    pub slot: wire::U8,
    pub amount: wire::U8,
    /// Wire presence: optional value is preceded by a presence marker.
    pub item_stack_net_id: Option<ItemStackNetID>,
    pub custom_name: BedrockSafetyRedactableString,
    pub durability_correction: wire::ZigZag32,
}

/// StackRequestSlotInfo holds information on a specific slot client-side.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct StackRequestSlotInfo {
    pub full_container_name: FullContainerName,
    /// Slot is the index of the slot within the container with the ContainerID above.
    pub slot: wire::U8,
    pub net_id_variant: wire::I32LE,
}

// Domain: map

/// MapDecoration is a fixed decoration on a map: Its position or other properties do not change
/// automatically client-side.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct MapDecoration {
    pub image_type: MapDecorationType,
    /// Rotation is the rotation of the map decoration. It is byte due to the 16 fixed directions that
    /// the map decoration may face.
    pub rotation: wire::U8,
    /// X is the offset on the X axis in pixels of the decoration.
    pub x: wire::U8,
    /// Y is the offset on the Y axis in pixels of the decoration.
    pub y: wire::U8,
    /// Label is the name of the map decoration. This name may be of any value.
    pub label: String,
    pub color: MceColor,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct MapItemTrackedActorUniqueId {
    pub type_: MapItemTrackedActorType,
    /// Wire presence: optional value is preceded by a presence marker.
    pub entity_id: Option<ActorUniqueID>,
    /// Wire presence: optional value is preceded by a presence marker.
    pub block_position: Option<BlockPos>,
}

/// PixelRequest is the request for the colour of a pixel in a MapInfoRequest packet.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct PixelRequest {
    pub pixel: wire::U32LE,
    pub index: wire::U16LE,
}

// Domain: memory_category

/// MemoryCategoryCounter represents a memory usage counter for a specific category.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct MemoryCategoryCounter {
    /// Category is the memory category. It is one of the MemoryCategory constants above.
    pub category: MemoryCategory,
    pub current_bytes: wire::U64LE,
}

// Domain: misc

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

#[derive(Clone, Debug, PartialEq)]
pub enum BedrockDDUI {
    DataStoreUpdate(BedrockDDUIDataStoreUpdate),
    DataStoreChange {
        data_store_name: String,
        property: String,
        update_count: wire::U32LE,
        the_new_property_value: DynamicValue,
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
pub struct ContentIdentity {
    pub identity: String,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct DataItemEntry {
    pub id: wire::VarUInt,
    pub payload: DataItemEntryValue,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct DebugMarkerData {
    pub text: String,
    pub position: glam::Vec3,
    pub color: MceColor,
    pub duration: wire::U64LE,
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

#[derive(Clone, Debug, PartialEq, Default)]
pub enum DynamicValue {
    #[default]
    None,
    Bool(bool),
    Int64(wire::I64LE),
    Double(wire::F64LE),
    String(String),
    List(Vec<DynamicValue>),
    Map(Vec<(String, DynamicValue)>),
}

impl DynamicValue {
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
pub struct Experiments {
    pub toggles: Vec<ExperimentToggle>,
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
pub struct GameRulesChangedData {
    pub rules_list: Vec<GameRule>,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct HeightmapData {
    pub height_map_type: HeightMapDataType,
    /// Wire presence: optional value is preceded by a presence marker.
    pub subchunk_height_map: Option<[[wire::I8; 16]; 16]>,
    pub render_height_map_type: HeightMapDataType,
    /// Wire presence: optional value is preceded by a presence marker.
    pub subchunk_render_height_map: Option<[[wire::I8; 16]; 16]>,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct LegacySetSlot {
    pub container_enum: ContainerEnumName,
    pub slots: Vec<wire::U8>,
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
    pub rule_data: GameRulesChangedData,
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
pub struct NetworkItemInstanceDescriptorSerializedData {
    pub id: wire::ZigZag32,
    pub stack_size: wire::U16LE,
    pub aux_value: wire::VarUInt,
    pub block_runtime_id: wire::ZigZag32,
    pub user_data_buffer: bytes::Bytes,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct NetworkItemStackDescriptorSerializedData {
    pub id: wire::I16LE,
    pub stack_size: wire::U16LE,
    pub aux_value: wire::VarUInt,
    /// Wire presence: optional value is preceded by a presence marker.
    pub net_id_variant: Option<wire::ZigZag32>,
    pub block_runtime_id: wire::VarUInt,
    pub user_data_buffer: bytes::Bytes,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct NetworkPermissions {
    pub server_auth_sound_enabled: bool,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct PackedItemUseLegacyInventoryTransaction {
    pub legacy_request_id: ItemStackLegacyRequestID,
    /// Wire presence: optional value is preceded by a presence marker.
    pub legacy_set_item_slots: Option<Vec<LegacySetSlot>>,
    /// Wire presence: optional value is preceded by a presence marker.
    pub item_use_transaction: Option<ItemUseInventoryTransaction>,
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
    pub block_definition: wire::NetworkNbt,
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

#[derive(Clone, Debug, Default, PartialEq)]
pub struct SocialEventsServerTelemetryData {
    pub server_id: String,
    pub scenario_id: String,
    pub world_id: String,
    pub owner_id: String,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct SpawnSettings {
    pub spawn_biome_type: SpawnBiomeType,
    pub user_defined_biome_name: String,
    pub dimension: wire::ZigZag32,
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

#[derive(Clone, Debug, Default, PartialEq)]
pub struct TintMapColor {
    pub colors: [MceColor; 4],
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

#[derive(Clone, Debug, Default, PartialEq)]
pub struct WebSocketData {
    pub websocket_server_uri: String,
}

// Domain: pack

#[derive(Clone, Debug, Default, PartialEq)]
pub struct PackIDVersionData {
    pub pack_uuid: uuid::Uuid,
    pub pack_version: SemVersionData,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct PackIdVersion {
    pub pack_uuid: uuid::Uuid,
    pub pack_version: SemVersion,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct PackInfoData {
    pub pack_id_version: PackIDVersionData,
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

// Domain: player

/// PlayerBlockAction ...
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
pub enum PlayerListData {
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

impl PlayerListData {
    pub fn discriminant(&self) -> u8 {
        match self {
            Self::Add { .. } => 0,
            Self::Remove { .. } => 1,
        }
    }
}

impl Default for PlayerListData {
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
pub enum PlayerLocationData {
    PlayerLocationCoordinates {
        packet_type: PlayerLocationType,
        position: glam::Vec3,
    },
    PlayerLocationHide {
        packet_type: PlayerLocationType,
    },
}

impl PlayerLocationData {
    pub fn discriminant(&self) -> u32 {
        match self {
            Self::PlayerLocationCoordinates { .. } => 0,
            Self::PlayerLocationHide { .. } => 1,
        }
    }
}

impl Default for PlayerLocationData {
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
pub enum PlayerUpdateEntityOverridesData {
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

impl PlayerUpdateEntityOverridesData {
    pub fn discriminant(&self) -> u8 {
        match self {
            Self::ClearOverride { .. } => 0,
            Self::RemoveOverride { .. } => 1,
            Self::IntOverride { .. } => 2,
            Self::FloatOverride { .. } => 3,
        }
    }
}

impl Default for PlayerUpdateEntityOverridesData {
    fn default() -> Self {
        Self::ClearOverride {
            type_: Default::default(),
        }
    }
}

#[derive(Clone, Debug, PartialEq, Default)]
pub enum PlayerVideoCaptureData {
    #[default]
    StopVideoCapture,
    StartVideoCapture {
        frame_rate: wire::U32LE,
        file_prefix: String,
    },
}

impl PlayerVideoCaptureData {
    pub fn discriminant(&self) -> u8 {
        match self {
            Self::StopVideoCapture => 0,
            Self::StartVideoCapture { .. } => 1,
        }
    }
}

// Domain: position_tracking

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

// Domain: recipe

/// MultiRecipe serves as an 'enable' switch for multi-shape recipes.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct MultiRecipe {
    pub multi_recipe_uuid: uuid::Uuid,
    pub net_id: RecipeNetID,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct RecipeIngredient {
    pub item_descriptor: ItemDescriptor,
    pub stack_size: wire::U16LE,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct RecipeIngredientSerializedData {
    pub descriptor: Vec<(String, String)>,
    pub aux_value: wire::ZigZag32,
    pub stack_size: wire::ZigZag32,
}

#[derive(Clone, Copy, Debug, Default, PartialEq, Eq, Hash)]
pub struct RecipeNetID(pub u32);

impl wire::WireCodec for RecipeNetID {
    fn encode<W: std::io::Write>(&self, writer: &mut W) -> std::io::Result<()> {
        <wire::VarUInt as wire::WireCodec>::encode(&wire::VarUInt(self.0), writer)
    }

    fn decode<R: std::io::Read>(reader: &mut R) -> std::io::Result<Self> {
        <wire::VarUInt as wire::WireCodec>::decode(reader).map(|value| Self(value.0))
    }
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct RecipeUnlockRequirementSerializedData {
    pub unlocking_context: RecipeUnlockingRequirementUnlockingContext,
    /// Wire presence: optional value is preceded by a presence marker.
    pub unlocking_ingredients: Option<Vec<RecipeIngredientSerializedData>>,
}

/// ShapedRecipe is a recipe that has a specific shape that must be used to craft the output of the
/// recipe. Trying to craft the item in any other shape will not work. The ShapedRecipe is of the
/// same structure as the ShapedChemistryRecipe.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct ShapedRecipe {
    /// RecipeID is a unique ID of the recipe. This ID must be unique amongst all other types of recipes
    /// too, but its functionality is not exactly known.
    pub recipe_id: String,
    /// Width is the width of the recipe's shape.
    pub width: wire::ZigZag32,
    /// Height is the height of the recipe's shape.
    pub height: wire::ZigZag32,
    pub ingredients: Vec<RecipeIngredientSerializedData>,
    pub results: Vec<NetworkItemInstanceDescriptorSerializedData>,
    /// UUID is a UUID identifying the recipe. Since the CraftingEvent packet no longer exists, this can
    /// always be empty.
    pub uuid: uuid::Uuid,
    pub tag: String,
    /// Priority ...
    pub priority: wire::ZigZag32,
    /// AssumeSymmetry specifies if the recipe is symmetrical. If this is set to true, the recipe will
    /// be mirrored along the diagonal axis. This means that the recipe will be the same if rotated 180
    /// degrees.
    pub assume_symmetry: bool,
    /// Wire presence: optional value is preceded by a presence marker.
    pub unlocking_requirement: Option<RecipeUnlockRequirementSerializedData>,
    pub net_id: RecipeNetID,
}

/// ShapelessRecipe is a recipe that has no particular shape. Its functionality is shared with the
/// RecipeShulkerBox and RecipeShapelessChemistry types.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct ShapelessRecipe {
    /// RecipeID is a unique ID of the recipe. This ID must be unique amongst all other types of recipes
    /// too, but its functionality is not exactly known.
    pub recipe_id: String,
    pub ingredients: Vec<RecipeIngredientSerializedData>,
    pub results: Vec<NetworkItemInstanceDescriptorSerializedData>,
    /// UUID is a UUID identifying the recipe. Since the CraftingEvent packet no longer exists, this can
    /// always be empty.
    pub uuid: uuid::Uuid,
    pub tag: String,
    /// Priority ...
    pub priority: wire::ZigZag32,
    /// Wire presence: optional value is preceded by a presence marker.
    pub unlocking_requirement: Option<RecipeUnlockRequirementSerializedData>,
    pub net_id: RecipeNetID,
}

/// SmithingTransformRecipe is a recipe specifically used for smithing tables. It has three input
/// items and adds them together, resulting in a new item.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct SmithingTransformRecipe {
    /// RecipeID is a unique ID of the recipe. This ID must be unique amongst all other types of recipes
    /// too, but its functionality is not exactly known.
    pub recipe_id: String,
    pub template_ingredient: RecipeIngredientSerializedData,
    pub base_ingredient: RecipeIngredientSerializedData,
    pub addition_ingredient: RecipeIngredientSerializedData,
    /// Result is the resulting item from the two items being added together.
    pub result: NetworkItemInstanceDescriptorSerializedData,
    pub tag: String,
    pub net_id: RecipeNetID,
}

/// SmithingTrimRecipe is a recipe specifically used for applying armour trims to an armour piece
/// inside a smithing table.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct SmithingTrimRecipe {
    /// RecipeID is a unique ID of the recipe. This ID must be unique amongst all other types of recipes
    /// too, but its functionality is not exactly known.
    pub recipe_id: String,
    pub template_ingredient: RecipeIngredientSerializedData,
    pub base_ingredient: RecipeIngredientSerializedData,
    pub addition_ingredient: RecipeIngredientSerializedData,
    pub tag: String,
    pub net_id: RecipeNetID,
}

// Domain: resource_pack

#[derive(Clone, Debug, PartialEq)]
pub enum ResourcePackClientResponseData {
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

impl ResourcePackClientResponseData {
    pub fn discriminant(&self) -> i8 {
        match self {
            Self::Cancel { .. } => 1,
            Self::Downloading { .. } => 2,
            Self::DownloadingFinished { .. } => 3,
            Self::ResourcePackStackFinished { .. } => 4,
        }
    }
}

impl Default for ResourcePackClientResponseData {
    fn default() -> Self {
        Self::Cancel {
            response_type: Default::default(),
        }
    }
}

// Domain: scoreboard

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

// Domain: shape

/// PrimitiveShape defines a single shape to be rendered on the client. Each shape has a unique
/// NetworkID and a set of optional parameters depending on its type.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct PrimitiveShape {
    /// NetworkID is the network ID of the shape.
    pub network_id: wire::VarULong,
    /// DimensionID is the optional dimension ID where the shape is rendered.
    /// Wire presence: optional value is preceded by a presence marker.
    pub shape_type: Option<ScriptModuleMinecraftScriptPrimitiveShapeType>,
    /// Location is the location of the shape.
    /// Wire presence: optional value is preceded by a presence marker.
    pub location: Option<glam::Vec3>,
    /// Scale is the scale of the shape.
    /// Wire presence: optional value is preceded by a presence marker.
    pub scale: Option<wire::F32LE>,
    /// Rotation is the rotation of the shape.
    /// Wire presence: optional value is preceded by a presence marker.
    pub rotation: Option<glam::Vec3>,
    /// TotalTimeLeft is the total time left of the shape.
    /// Wire presence: optional value is preceded by a presence marker.
    pub total_time_left: Option<wire::F32LE>,
    /// Rotation is the rotation of the shape.
    /// Wire presence: optional value is preceded by a presence marker.
    pub maximum_render_distance: Option<wire::F32LE>,
    /// TotalTimeLeft is the total time left of the shape.
    /// Wire presence: optional value is preceded by a presence marker.
    pub color: Option<MceColor>,
    /// DimensionID is the optional dimension ID where the shape is rendered.
    /// Wire presence: optional value is preceded by a presence marker.
    pub dimension_id: Option<DimensionType>,
    /// AttachedToEntityID is the optional unique ID of the entity the shape is attached to. Mojang's
    /// documentation describes it as a runtime ID, but the field is an ActorUniqueID and the client
    /// resolves it as one.
    /// Wire presence: optional value is preceded by a presence marker.
    pub attached_to_entity_id: Option<ActorUniqueID>,
    /// ExtraShapeData holding data specific to the type of shape (such as text string for the text
    /// shape).
    pub extra_shape_data: PrimitiveShapeExtraShapeData,
}

// Domain: skin

#[derive(Clone, Debug, Default, PartialEq)]
pub struct SkinImage {
    pub width: wire::U32LE,
    pub height: wire::U32LE,
    pub image_bytes: Vec<wire::U8>,
}

// Domain: sound

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

// Domain: structure

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

/// StructureSettings is a struct holding settings of a structure block. Its fields may be changed
/// using the in-game UI on the client-side.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct StructureSettings {
    pub structure_palette_name: String,
    pub should_ignore_entities: bool,
    pub should_ignore_blocks: bool,
    pub should_allow_non_ticking_player_and_ticking_area_chunks: bool,
    pub structure_size: BlockPos,
    pub structure_offset: BlockPos,
    pub last_edit_player: ActorUniqueID,
    /// Rotation is the rotation that the structure block should obtain. See the constants above for
    /// available options.
    pub rotation: Rotation,
    /// Mirror specifies the way the structure should be mirrored. It is either no mirror at all, mirror
    /// on the x/z axis or both.
    pub mirror: Mirror,
    /// AnimationMode ...
    pub animation_mode: AnimationMode,
    pub animation_seconds: wire::F32LE,
    pub integrity_value: wire::F32LE,
    pub integrity_seed: wire::U32LE,
    pub rotation_pivot: glam::Vec3,
}

// Domain: sub_chunk

#[derive(Clone, Debug, Default, PartialEq)]
pub struct SubChunkData {
    pub sub_chunk_pos_offset: SubChunkPosOffset,
    pub sub_chunk_request_result: SubChunkRequestResult,
    /// Wire presence: optional value is preceded by a presence marker.
    pub serialized_sub_chunk: Option<String>,
    pub height_map_data: HeightmapData,
    /// Wire presence: optional value is preceded by a presence marker.
    pub blob_id: Option<wire::U64LE>,
}

#[derive(Clone, Copy, Debug, Default, PartialEq, Eq, Hash)]
pub struct SubChunkMetadata(pub u64);

impl wire::WireCodec for SubChunkMetadata {
    fn encode<W: std::io::Write>(&self, writer: &mut W) -> std::io::Result<()> {
        <wire::U64LE as wire::WireCodec>::encode(&wire::U64LE(self.0), writer)
    }

    fn decode<R: std::io::Read>(reader: &mut R) -> std::io::Result<Self> {
        <wire::U64LE as wire::WireCodec>::decode(reader).map(|value| Self(value.0))
    }
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct SubChunkPosOffset {
    pub subchunk_offset_x: wire::I8,
    pub subchunk_offset_y: wire::I8,
    pub subchunk_offset_z: wire::I8,
}

// Domain: sync_world_clocks

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

// Domain: text

#[derive(Clone, Debug, PartialEq)]
pub enum TextData {
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

impl TextData {
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

impl Default for TextData {
    fn default() -> Self {
        Self::Raw {
            message: Default::default(),
        }
    }
}

// Domain: trim

/// TrimMaterial represents a material that can be used when applying an armour trim.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct TrimMaterial {
    /// MaterialID is the identifier of the material, for example 'netherite'.
    pub material_id: String,
    /// Colour is the colour code used for text formatting, for example '§j'.
    pub color: String,
    /// ItemName is the identifier of the item that represents the material, for example,
    /// 'minecraft:netherite_ingot'.
    pub item_name: String,
}

/// TrimPattern represents a pattern that can be applied to an armour piece in combination with a
/// TrimMaterial.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct TrimPattern {
    /// ItemName is the identifier of the item that represents the pattern, for example
    /// 'minecraft:wayfinder_armor_trim_smithing_template'.
    pub item_name: String,
    /// PatternID is the identifier of the pattern, for example, 'wayfinder'.
    pub pattern_id: String,
}

// Domain: voxel

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

// Domain: waypoint

/// LocatorBarWaypoint represents a waypoint entry in the locator bar packet.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct LocatorBarWaypoint {
    /// GroupHandle is the UUID handle for the waypoint group.
    pub group_handle: WaypointGroupWaypointHandle,
    pub server_waypoint_payload: ServerWaypoint,
    pub action_flag: ServerWaypointGroupAction,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct WaypointGroupWaypointHandle {
    pub uuid: uuid::Uuid,
}

// Domain: world

/// DimensionDefinition contains information specifying dimension-specific properties, used for
/// data-driven dimensions. These include the range (the height min/max), generator variant, and
/// more.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct DimensionDefinition {
    pub height_maximum: wire::ZigZag32,
    pub height_minimum: wire::ZigZag32,
    pub generator_type: GeneratorType,
    /// DimensionType is the numeric identifier of the dimension. This cannot override a vanilla
    /// dimension (0-2), but custom dimensions should start from 1000 like vanilla.
    pub dimension_type: DimensionType,
    /// PackID is the UUID of the behaviour pack which has added the dimension.
    pub pack_id: uuid::Uuid,
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct WorldPosition {
    pub position: glam::Vec3,
    pub dimension_type: DimensionType,
}
