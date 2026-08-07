// Code generated from canonical protocol manifest v2. DO NOT EDIT.

use crate::enums::*;

#[derive(Clone, Debug, Default, PartialEq, Eq)]
pub struct Nbt(pub Vec<u8>);

#[derive(Clone, Debug, PartialEq)]
pub struct ActorDataBoundingBoxComponent {
    pub actor_data_bounding_box: [f32; 3],
}

#[derive(Clone, Debug, PartialEq)]
pub struct ActorDataFlagComponent {
    pub actor_flag_bitset_data: Vec<u8>,
}

#[derive(Clone, Debug, PartialEq)]
pub struct ActorLink {
    pub target_a: ActorUniqueID,
    pub target_b: ActorUniqueID,
    pub r#type: ActorLinkType,
    pub immediate: bool,
    pub passenger_initiated: bool,
    pub vehicle_angular_velocity: f32,
}

#[derive(Clone, Debug, PartialEq)]
pub struct ActorRuntimeID {
    pub actor_runtime_id: u64,
}

#[derive(Clone, Debug, PartialEq)]
pub struct ActorUniqueID {
    pub actor_unique_id: i64,
}

#[derive(Clone, Debug, PartialEq)]
pub struct AdventureSettings {
    pub no_pv_m: bool,
    pub no_mv_p: bool,
    pub immutable_world: bool,
    pub show_name_tags: bool,
    pub auto_jump: bool,
}

#[derive(Clone, Debug, PartialEq)]
pub struct AgentCapabilities {
    pub can_modify_blocks: Option<bool>,
}

#[derive(Clone, Debug, PartialEq)]
pub struct AnimatedImageData {
    pub skin_image: SkinImage,
    pub animated_texture_type: PersonaAnimatedTextureType,
    pub frames: f32,
    pub animation_expression: PersonaAnimationExpression,
}

#[derive(Clone, Debug, PartialEq)]
pub struct ArmorSlotAndDamagePair {
    pub armor_slot: LegacyArmorSlot,
    pub damage: i16,
}

#[derive(Clone, Debug, PartialEq)]
pub struct ArrowData {
    pub arrow_end_location: Option<glam::Vec3>,
    pub arrow_head_length: Option<f32>,
    pub arrow_head_radius: Option<f32>,
    pub num_segments: Option<u8>,
}

#[derive(Clone, Debug, PartialEq)]
pub struct AttributeData {
    pub min_value: f32,
    pub max_value: f32,
    pub current_value: f32,
    pub default_min_value: f32,
    pub default_max_value: f32,
    pub default_value: f32,
    pub name: String,
    pub modifiers: Vec<AttributeModifier>,
}

#[derive(Clone, Debug, PartialEq)]
pub enum AttributeLayerSyncPacketData {
    UpdateAttributeLayersData(AttributeLayerSyncPacketDataUpdateAttributeLayersData),
    UpdateAttributeLayerSettingsData(AttributeLayerSyncPacketDataUpdateAttributeLayerSettingsData),
    UpdateEnvironmentAttributesData(AttributeLayerSyncPacketDataUpdateEnvironmentAttributesData),
    RemoveEnvironmentAttributesData(AttributeLayerSyncPacketDataRemoveEnvironmentAttributesData),
}

#[derive(Clone, Debug, PartialEq)]
pub struct AttributeLayerSyncPacketDataRemoveEnvironmentAttributesData {
    pub attribute_layer_name: String,
    pub attribute_layer_dimension: DimensionType,
    pub attributes: Vec<String>,
}

#[derive(Clone, Debug, PartialEq)]
pub struct AttributeLayerSyncPacketDataUpdateAttributeLayerSettingsData {
    pub attribute_layer_name: String,
    pub attribute_layer_dimension: DimensionType,
    pub attributes_layer_settings: EASAttributeLayerSettings,
}

#[derive(Clone, Debug, PartialEq)]
pub struct AttributeLayerSyncPacketDataUpdateAttributeLayersData {
    pub attribute_layers: Vec<EASAttributeLayerData>,
}

#[derive(Clone, Debug, PartialEq)]
pub struct AttributeLayerSyncPacketDataUpdateEnvironmentAttributesData {
    pub attribute_layer_name: String,
    pub attribute_layer_dimension: DimensionType,
    pub attributes: Vec<EASEnvironmentAttributeData>,
}

#[derive(Clone, Debug, PartialEq)]
pub struct AttributeModifier {
    pub id: String,
    pub name: String,
    pub amount: f32,
    pub operation: i32,
    pub operand: i32,
    pub is_serializable: bool,
}

#[derive(Clone, Debug, PartialEq)]
pub struct AvailableCommandsChainedSubcommandData {
    pub name: String,
    pub sub_command_values: Vec<AvailableCommandsChainedSubcommandRelationship>,
}

#[derive(Clone, Debug, PartialEq)]
pub struct AvailableCommandsChainedSubcommandRelationship {
    pub sub_command_first_value: u32,
    pub sub_command_second_value: u32,
}

#[derive(Clone, Debug, PartialEq)]
pub struct AvailableCommandsConstrainedValueData {
    pub enum_value_symbol: u32,
    pub enum_symbol: u32,
    pub constraint_indices: Vec<u8>,
}

#[derive(Clone, Debug, PartialEq)]
pub struct AvailableCommandsEnumData {
    pub name: String,
    pub values: Vec<u32>,
}

#[derive(Clone, Debug, PartialEq)]
pub struct AvailableCommandsOverloadData {
    pub is_chaining: bool,
    pub parameter_data: Vec<AvailableCommandsParamData>,
}

#[derive(Clone, Debug, PartialEq)]
pub struct AvailableCommandsPacketCommandData {
    pub name: String,
    pub description: String,
    pub flags: u16,
    pub permission_level: String,
    pub alias_enum: i32,
    pub command_data_chained_subcommand_indexes: Vec<u32>,
    pub overloads: Vec<AvailableCommandsOverloadData>,
}

#[derive(Clone, Debug, PartialEq)]
pub struct AvailableCommandsParamData {
    pub name: String,
    pub parse_symbol: u32,
    pub is_optional: bool,
    pub options: u8,
}

#[derive(Clone, Debug, PartialEq)]
pub struct AvailableCommandsSoftEnumData {
    pub enum_name: String,
    pub enum_options: Vec<String>,
}

#[derive(Clone, Debug, PartialEq)]
pub enum BedrockDDUI {
    DataStoreUpdate(BedrockDDUIDataStoreUpdate),
    DataStoreChange(BedrockDDUIDataStoreChange),
    DataStoreRemoval(BedrockDDUIDataStoreRemoval),
}

#[derive(Clone, Debug, PartialEq)]
pub struct BedrockDDUIDataStoreChange {
    pub data_store_name: String,
    pub property: String,
    pub update_count: u32,
    pub the_new_property_value: CerealDynamicValue,
}

#[derive(Clone, Debug, PartialEq)]
pub struct BedrockDDUIDataStoreRemoval {
    pub data_store_name: String,
}

#[derive(Clone, Debug, PartialEq)]
pub struct BedrockDDUIDataStoreUpdate {
    pub data_store_name: String,
    pub property: String,
    pub path: String,
    pub data: BedrockDDUIDataStoreUpdateData,
    pub property_update_count: u32,
    pub path_update_count: u32,
}

#[derive(Clone, Debug, PartialEq)]
pub enum BedrockDDUIDataStoreUpdateData {
    Double(f64),
    Bool(bool),
    String(String),
}

#[derive(Clone, Debug, PartialEq)]
pub struct BedrockProfileWhiskerDiagnosticsScopeDataSummary {
    pub label: String,
    pub indentation: String,
    pub total_high_cost_ns: u64,
    pub total_mid_cost_ns: u64,
    pub total_low_cost_ns: u64,
}

#[derive(Clone, Debug, PartialEq)]
pub struct BedrockSafetyRedactableString {
    pub unredacted: String,
    pub redacted: Option<String>,
}

#[derive(Clone, Debug, PartialEq)]
pub struct BiomeCappedSurfaceData {
    pub floor_blocks: Vec<u32>,
    pub ceiling_blocks: Vec<u32>,
    pub sea_block: Option<u32>,
    pub foundation_block: Option<u32>,
    pub beach_block: Option<u32>,
}

#[derive(Clone, Debug, PartialEq)]
pub struct BiomeClimateData {
    pub temperature: f32,
    pub downfall: f32,
    pub snow_accumulation_min: f32,
    pub snow_accumulation_max: f32,
}

#[derive(Clone, Debug, PartialEq)]
pub struct BiomeConditionalTransformationData {
    pub transforms_into: Vec<BiomeWeightedData>,
    pub condition_json: u16,
    pub min_passing_neighbors: u32,
}

#[derive(Clone, Debug, PartialEq)]
pub struct BiomeConsolidatedFeatureData {
    pub scatter: BiomeScatterParamData,
    pub feature: u16,
    pub identifier: u16,
    pub pass: u16,
    pub can_use_internal_feature: bool,
}

#[derive(Clone, Debug, PartialEq)]
pub struct BiomeConsolidatedFeaturesData {
    pub features: Vec<BiomeConsolidatedFeatureData>,
}

#[derive(Clone, Debug, PartialEq)]
pub struct BiomeCoordinateData {
    pub min_value_type: i32,
    pub min_value: u16,
    pub max_value_type: i32,
    pub max_value: u16,
    pub grid_offset: u32,
    pub grid_step_size: u32,
    pub distribution: RandomDistributionType,
}

#[derive(Clone, Debug, PartialEq)]
pub struct BiomeDefinitionChunkGenData {
    pub climate: Option<BiomeClimateData>,
    pub consolidated_features: Option<BiomeConsolidatedFeaturesData>,
    pub mountain_params: Option<BiomeMountainParamsData>,
    pub surface_material_adjustments: Option<BiomeSurfaceMaterialAdjustmentData>,
    pub overworld_gen_rules: Option<BiomeOverworldGenRulesData>,
    pub multinoise_gen_rules: Option<BiomeMultinoiseGenRulesData>,
    pub legacy_world_gen_rules: Option<BiomeLegacyWorldGenRulesData>,
    pub replacement_biomes: Option<BiomeReplacementsData>,
    pub village_type: Option<VillageType>,
    pub surface_builder_data: Option<BiomeSurfaceBuilderData>,
    pub subsurface_builder_data: Option<BiomeSurfaceBuilderData>,
}

#[derive(Clone, Debug, PartialEq)]
pub struct BiomeDefinitionData {
    pub id: u16,
    pub temperature: f32,
    pub downfall: f32,
    pub foliage_snow: f32,
    pub depth: f32,
    pub scale: f32,
    pub map_water_color_argb: i32,
    pub rain: bool,
    pub tags: Option<BiomeTagsData>,
    pub chunk_gen_data: Option<BiomeDefinitionChunkGenData>,
}

#[derive(Clone, Debug, PartialEq)]
pub struct BiomeElementData {
    pub noise_freq_scale: f32,
    pub noise_lower_bound: f32,
    pub noise_upper_bound: f32,
    pub height_min_type: i32,
    pub height_min: u16,
    pub height_max_type: i32,
    pub height_max: u16,
    pub adjusted_materials: BiomeSurfaceMaterialData,
}

#[derive(Clone, Debug, PartialEq)]
pub struct BiomeLegacyWorldGenRulesData {
    pub legacy_pre_hills_edge: Vec<BiomeConditionalTransformationData>,
}

#[derive(Clone, Debug, PartialEq)]
pub struct BiomeMesaSurfaceData {
    pub clay_material: u32,
    pub hard_clay_material: u32,
    pub bryce_pillars: bool,
    pub has_forest: bool,
}

#[derive(Clone, Debug, PartialEq)]
pub struct BiomeMountainParamsData {
    pub steep_block: u32,
    pub north_slopes: bool,
    pub south_slopes: bool,
    pub west_slopes: bool,
    pub east_slopes: bool,
    pub top_slide_enabled: bool,
}

#[derive(Clone, Debug, PartialEq)]
pub struct BiomeMultinoiseGenRulesData {
    pub temperature: f32,
    pub humidity: f32,
    pub altitude: f32,
    pub weirdness: f32,
    pub weight: f32,
}

#[derive(Clone, Debug, PartialEq)]
pub struct BiomeNoiseGradientSurfaceData {
    pub non_replaceable_blocks: Vec<u32>,
    pub gradient_blocks: Vec<SerializedNoiseBlockSpecifier>,
    pub noise: NoiseDescriptor,
}

#[derive(Clone, Debug, PartialEq)]
pub struct BiomeOverworldGenRulesData {
    pub hills_transformations: Vec<BiomeWeightedData>,
    pub mutate_transformations: Vec<BiomeWeightedData>,
    pub river_transformations: Vec<BiomeWeightedData>,
    pub shore_transformations: Vec<BiomeWeightedData>,
    pub pre_hills_edge: Vec<BiomeConditionalTransformationData>,
    pub post_shore_edge: Vec<BiomeConditionalTransformationData>,
    pub climate: Vec<BiomeWeightedTemperatureData>,
}

#[derive(Clone, Debug, PartialEq)]
pub struct BiomeReplacementData {
    pub replacement_biome: u16,
    pub dimension: u16,
    pub target_biomes: Vec<u16>,
    pub amount: f32,
    pub noise_frequency_scale: f32,
    pub replacement_index: u32,
}

#[derive(Clone, Debug, PartialEq)]
pub struct BiomeReplacementsData {
    pub biome_replacements: Vec<BiomeReplacementData>,
}

#[derive(Clone, Debug, PartialEq)]
pub struct BiomeScatterParamData {
    pub coordinates: Vec<BiomeCoordinateData>,
    pub eval_order: CoordinateEvaluationOrder,
    pub chance_percent_type: i32,
    pub chance_percent: u16,
    pub chance_numerator: i32,
    pub chance_denominator: i32,
    pub iterations_type: i32,
    pub iterations: u16,
}

#[derive(Clone, Debug, PartialEq)]
pub struct BiomeStringList {
    pub strings: Vec<String>,
}

#[derive(Clone, Debug, PartialEq)]
pub struct BiomeSurfaceBuilderData {
    pub surface_materials: Option<BiomeSurfaceMaterialData>,
    pub has_default_overworld_surface: bool,
    pub has_swamp_surface: bool,
    pub has_frozen_ocean_surface: bool,
    pub has_the_end_surface: bool,
    pub mesa_surface: Option<BiomeMesaSurfaceData>,
    pub capped_surface: Option<BiomeCappedSurfaceData>,
    pub noise_gradient_surface: Option<BiomeNoiseGradientSurfaceData>,
}

#[derive(Clone, Debug, PartialEq)]
pub struct BiomeSurfaceMaterialAdjustmentData {
    pub adjustments: Vec<BiomeElementData>,
}

#[derive(Clone, Debug, PartialEq)]
pub struct BiomeSurfaceMaterialData {
    pub top_block: u32,
    pub mid_block: u32,
    pub sea_floor_block: u32,
    pub foundation_block: u32,
    pub sea_block: u32,
    pub sea_floor_depth: i32,
}

#[derive(Clone, Debug, PartialEq)]
pub struct BiomeTagsData {
    pub tags: Vec<u16>,
}

#[derive(Clone, Debug, PartialEq)]
pub struct BiomeWeightedData {
    pub biome_identifier: u16,
    pub weight: u32,
}

#[derive(Clone, Debug, PartialEq)]
pub struct BiomeWeightedTemperatureData {
    pub temperature: i32,
    pub weight: u32,
}

#[derive(Clone, Debug, PartialEq)]
pub struct BlockPos {
    pub x: i32,
    pub y: i32,
    pub z: i32,
}

#[derive(Clone, Debug, PartialEq)]
pub enum BookEditAction {
    ReplacePage(BookEditActionReplacePage),
    AddPage(BookEditActionAddPage),
    DeletePage(BookEditActionDeletePage),
    SwapPages(BookEditActionSwapPages),
    Finalize(BookEditActionFinalize),
}

#[derive(Clone, Debug, PartialEq)]
pub struct BookEditActionAddPage {
    pub page_index: i32,
    pub page_text: String,
    pub photo_name: String,
}

#[derive(Clone, Debug, PartialEq)]
pub struct BookEditActionDeletePage {
    pub page_index: i32,
}

#[derive(Clone, Debug, PartialEq)]
pub struct BookEditActionFinalize {
    pub title: String,
    pub author: String,
    pub xuid: String,
}

#[derive(Clone, Debug, PartialEq)]
pub struct BookEditActionReplacePage {
    pub page_index: i32,
    pub page_text: String,
    pub photo_name: String,
}

#[derive(Clone, Debug, PartialEq)]
pub struct BookEditActionSwapPages {
    pub page_index: i32,
    pub swap_with_index: i32,
}

#[derive(Clone, Debug, PartialEq)]
pub struct BoxData {
    pub box_bound: glam::Vec3,
}

#[derive(Clone, Debug, PartialEq)]
pub struct CameraAimAssistActorPriorityPriorityData {
    pub preset_index: i32,
    pub category_index: i32,
    pub actor_index: i32,
    pub priority_value: i32,
}

#[derive(Clone, Debug, PartialEq)]
pub struct CameraAimAssistCategoryDefinition {
    pub name: String,
    pub priorities: CameraAimAssistCategoryPriorities,
}

#[derive(Clone, Debug, PartialEq)]
pub struct CameraAimAssistCategoryPriorities {
    pub entities: Vec<(String, i32)>,
    pub blocks: Vec<(String, i32)>,
    pub block_tags: Vec<(String, i32)>,
    pub entity_type_families: Vec<(String, i32)>,
    pub entity_default: Option<i32>,
    pub block_default: Option<i32>,
}

#[derive(Clone, Debug, PartialEq)]
pub struct CameraAimAssistCommandPresetDefinition {
    pub preset_id: Option<String>,
    pub target_mode: Option<CameraAimAssistTargetMode>,
    pub view_angle: Option<glam::Vec2>,
    pub distance: Option<f32>,
}

#[derive(Clone, Debug, PartialEq)]
pub struct CameraAimAssistPresetDefinition {
    pub identifier: String,
    pub exclusion_settings: CameraAimAssistPresetExclusionDefinition,
    pub liquid_targeting_list: Vec<String>,
    pub item_settings: Vec<(String, String)>,
    pub default_item_settings: Option<String>,
    pub hand_settings: Option<String>,
}

#[derive(Clone, Debug, PartialEq)]
pub struct CameraAimAssistPresetExclusionDefinition {
    pub blocks: Vec<String>,
    pub entities: Vec<String>,
    pub block_tags: Vec<String>,
    pub entity_type_families: Vec<String>,
}

#[derive(Clone, Debug, PartialEq)]
pub struct CameraInstructionData {
    pub set: Option<CameraInstructionOptionsSetInstruction>,
    pub clear: Option<bool>,
    pub fade: Option<CameraInstructionOptionsFadeInstruction>,
    pub target: Option<CameraInstructionOptionsTargetInstruction>,
    pub remove_target: Option<bool>,
    pub field_of_view: Option<CameraInstructionOptionsFovInstruction>,
    pub spline: Option<CameraInstructionOptionsSplineInstruction>,
    pub attach_to_entity: Option<CameraInstructionOptionsAttachToEntityInstruction>,
    pub detach_from_entity: Option<bool>,
}

#[derive(Clone, Debug, PartialEq)]
pub struct CameraInstructionOptionsAttachToEntityInstruction {
    pub entity_actor_id: i64,
}

#[derive(Clone, Debug, PartialEq)]
pub struct CameraInstructionOptionsFadeInstruction {
    pub time: Option<CameraInstructionOptionsFadeInstructionTimeOption>,
    pub color: Option<CameraInstructionOptionsFadeInstructionColorOption>,
}

#[derive(Clone, Debug, PartialEq)]
pub struct CameraInstructionOptionsFadeInstructionColorOption {
    pub red: f32,
    pub green: f32,
    pub blue: f32,
}

#[derive(Clone, Debug, PartialEq)]
pub struct CameraInstructionOptionsFadeInstructionTimeOption {
    pub fade_in_time: f32,
    pub hold_time: f32,
    pub fade_out_time: f32,
}

#[derive(Clone, Debug, PartialEq)]
pub struct CameraInstructionOptionsFovInstruction {
    pub field_of_view: f32,
    pub fov_ease_time: f32,
    pub fov_ease_type: String,
    pub field_of_view_clear: bool,
}

#[derive(Clone, Debug, PartialEq)]
pub struct CameraInstructionOptionsSetInstruction {
    pub preset: u32,
    pub ease: Option<CameraInstructionOptionsSetInstructionEaseOption>,
    pub pos: Option<CameraInstructionOptionsSetInstructionPosOption>,
    pub rot: Option<CameraInstructionOptionsSetInstructionRotOption>,
    pub facing: Option<CameraInstructionOptionsSetInstructionFacingOption>,
    pub view_offset: Option<CameraInstructionOptionsSetInstructionViewOffsetOption>,
    pub entity_offset: Option<CameraInstructionOptionsSetInstructionEntityOffsetOption>,
    pub default: Option<bool>,
    pub remove_ignore_starting_values_component: bool,
}

#[derive(Clone, Debug, PartialEq)]
pub struct CameraInstructionOptionsSetInstructionEaseOption {
    pub r#type: u8,
    pub time: f32,
}

#[derive(Clone, Debug, PartialEq)]
pub struct CameraInstructionOptionsSetInstructionEntityOffsetOption {
    pub entity_offset_x: f32,
    pub entity_offset_y: f32,
    pub entity_offset_z: f32,
}

#[derive(Clone, Debug, PartialEq)]
pub struct CameraInstructionOptionsSetInstructionFacingOption {
    pub pos: glam::Vec3,
}

#[derive(Clone, Debug, PartialEq)]
pub struct CameraInstructionOptionsSetInstructionPosOption {
    pub pos: glam::Vec3,
}

#[derive(Clone, Debug, PartialEq)]
pub struct CameraInstructionOptionsSetInstructionRotOption {
    pub x: f32,
    pub y: f32,
}

#[derive(Clone, Debug, PartialEq)]
pub struct CameraInstructionOptionsSetInstructionViewOffsetOption {
    pub x: f32,
    pub y: f32,
}

#[derive(Clone, Debug, PartialEq)]
pub struct CameraInstructionOptionsSplineInstruction {
    pub total_time: f32,
    pub r#type: u8,
    pub curve: Vec<glam::Vec3>,
    pub progress_key_frames: Vec<CameraInstructionOptionsSplineInstructionSplineProgressOption>,
    pub rotation_option: Vec<CameraInstructionOptionsSplineInstructionSplineRotationOption>,
    pub spline_identifier: String,
    pub load_from_json: bool,
}

#[derive(Clone, Debug, PartialEq)]
pub struct CameraInstructionOptionsSplineInstructionSplineProgressOption {
    pub key_frame_value: f32,
    pub key_frame_time: f32,
    pub key_frame_easing_func: String,
}

#[derive(Clone, Debug, PartialEq)]
pub struct CameraInstructionOptionsSplineInstructionSplineRotationOption {
    pub key_frame_value: glam::Vec3,
    pub key_frame_time: f32,
    pub key_frame_easing_func: String,
}

#[derive(Clone, Debug, PartialEq)]
pub struct CameraInstructionOptionsTargetInstruction {
    pub target_center_offset: Option<glam::Vec3>,
    pub target_actor_id: i64,
}

#[derive(Clone, Debug, PartialEq)]
pub struct CameraPreset {
    pub name: String,
    pub inherit_from: String,
    pub pos_x: Option<f32>,
    pub pos_y: Option<f32>,
    pub pos_z: Option<f32>,
    pub rot_x: Option<f32>,
    pub rot_y: Option<f32>,
    pub rotation_speed: Option<f32>,
    pub snap_to_target: Option<bool>,
    pub horizontal_rotation_limit: Option<glam::Vec2>,
    pub vertical_rotation_limit: Option<glam::Vec2>,
    pub continue_targeting: Option<bool>,
    pub block_listening_radius: Option<f32>,
    pub view_offset: Option<glam::Vec2>,
    pub entity_offset: Option<glam::Vec3>,
    pub radius: Option<f32>,
    pub yaw_limit_min: Option<f32>,
    pub yaw_limit_max: Option<f32>,
    pub listener: Option<CameraPresetAudioListener>,
    pub player_effects: Option<bool>,
    pub aim_assist: Option<CameraAimAssistCommandPresetDefinition>,
    pub control_scheme: Option<ControlSchemeScheme>,
}

#[derive(Clone, Debug, PartialEq)]
pub struct CameraPresetsData {
    pub presets: Vec<CameraPreset>,
}

#[derive(Clone, Debug, PartialEq)]
pub struct CameraSplineControlPoint {
    pub position: glam::Vec3,
}

#[derive(Clone, Debug, PartialEq)]
pub struct CameraSplineDefinition {
    pub name: String,
    pub total_time: f32,
    pub spline_type: String,
    pub control_points: Vec<CameraSplineControlPoint>,
    pub progress_key_frames: Vec<CameraSplineProgressKeyFrame>,
    pub rotation_key_frames: Vec<CameraSplineRotationKeyFrame>,
}

#[derive(Clone, Debug, PartialEq)]
pub struct CameraSplineProgressKeyFrame {
    pub progress: f32,
    pub time: f32,
    pub easing: Option<String>,
}

#[derive(Clone, Debug, PartialEq)]
pub struct CameraSplineRotationKeyFrame {
    pub rotation: glam::Vec3,
    pub time: f32,
    pub easing: Option<String>,
}

#[derive(Clone, Debug, PartialEq)]
pub enum CerealDynamicValue {
    None,
    Bool(bool),
    Int64(i64),
    Double(f64),
    String(String),
    List(Vec<Vec<u8>>),
    Map(Vec<(String, Vec<u8>)>),
}

#[derive(Clone, Debug, PartialEq)]
pub struct CerealizerExperimentsAnonExperimentToggle {
    pub name: String,
    pub enabled: bool,
}

#[derive(Clone, Debug, PartialEq)]
pub struct CerealizerNetworkItemInstanceDescriptorSerializedData {
    pub id: i32,
    pub stack_size: u16,
    pub aux_value: u32,
    pub block_runtime_id: i32,
    pub user_data_buffer: String,
}

#[derive(Clone, Debug, PartialEq)]
pub struct CerealizerNetworkItemStackDescriptorSerializedData {
    pub id: i16,
    pub stack_size: u16,
    pub aux_value: u32,
    pub net_id_variant: Option<i32>,
    pub block_runtime_id: u32,
    pub user_data_buffer: String,
}

#[derive(Clone, Debug, PartialEq)]
pub struct CerealizerRecipeIngredientSerializedData {
    pub descriptor: Vec<(String, String)>,
    pub aux_value: i32,
    pub stack_size: i32,
}

#[derive(Clone, Debug, PartialEq)]
pub struct CerealizerRecipeUnlockingRequirementSerializedData {
    pub unlocking_context: RecipeUnlockingRequirementUnlockingContext,
    pub unlocking_ingredients: Option<Vec<CerealizerRecipeIngredientSerializedData>>,
}

#[derive(Clone, Debug, PartialEq)]
pub struct ChangeEntityScore {
    pub action: String,
    pub scoreboard_id: ScoreboardId,
    pub objective_name: String,
    pub score_value: i32,
    pub actor_id: ActorUniqueID,
}

#[derive(Clone, Debug, PartialEq)]
pub struct ChangeFakePlayerScore {
    pub action: String,
    pub scoreboard_id: ScoreboardId,
    pub objective_name: String,
    pub score_value: i32,
    pub fake_player_name: String,
}

#[derive(Clone, Debug, PartialEq)]
pub struct ChangePlayerScore {
    pub action: String,
    pub scoreboard_id: ScoreboardId,
    pub objective_name: String,
    pub score_value: i32,
    pub player_unique_id: PlayerScoreboardId,
}

#[derive(Clone, Debug, PartialEq)]
pub struct ChunkPos {
    pub x: i32,
    pub z: i32,
}

#[derive(Clone, Debug, PartialEq)]
pub struct ClientboundDebugRendererDebugMarkerData {
    pub text: String,
    pub position: glam::Vec3,
    pub color: MceColor,
    pub duration: u64,
}

#[derive(Clone, Debug, PartialEq)]
pub struct CommandBlockUpdateBlockCommandData {
    pub block_position: BlockPos,
    pub command_block_mode: u32,
    pub redstone_mode: bool,
    pub is_conditional: bool,
}

#[derive(Clone, Debug, PartialEq)]
pub struct CommandBlockUpdateEntityCommandTarget {
    pub target_runtime_id: ActorRuntimeID,
}

#[derive(Clone, Debug, PartialEq)]
pub enum CommandBlockUpdateTarget {
    EntityCommandTarget(CommandBlockUpdateEntityCommandTarget),
    BlockCommandData(CommandBlockUpdateBlockCommandData),
}

#[derive(Clone, Debug, PartialEq)]
pub struct CommandOriginData {
    pub r#type: String,
    pub uuid: uuid::Uuid,
    pub request_id: String,
    pub player_id: i64,
}

#[derive(Clone, Debug, PartialEq)]
pub struct CommandOutputData {
    pub output_type: String,
    pub success_count: u32,
    pub output_messages: Vec<CommandOutputMessage>,
    pub data_set: Option<String>,
}

#[derive(Clone, Debug, PartialEq)]
pub struct CommandOutputMessage {
    pub message_id: String,
    pub successful: bool,
    pub parameters: Vec<String>,
}

#[derive(Clone, Debug, PartialEq)]
pub struct ConeData {
    pub radii: glam::Vec2,
    pub height: f32,
    pub num_segments: u8,
}

#[derive(Clone, Debug, PartialEq)]
pub struct ContainerMixDataEntry {
    pub from_item_id: i32,
    pub reagent_item_id: i32,
    pub to_item_id: i32,
}

#[derive(Clone, Debug, PartialEq)]
pub struct ContentIdentity {
    pub identity: String,
}

#[derive(Clone, Debug, PartialEq)]
pub struct CreativeGroupInfo {
    pub creative_category: CreativeItemCategory,
    pub name: String,
    pub group_icon_item: CerealizerNetworkItemInstanceDescriptorSerializedData,
}

#[derive(Clone, Debug, PartialEq)]
pub struct CreativeItemEntry {
    pub creative_net_id: TypedServerNetIdStructCreativeItemNetIdTag,
    pub item_instance: CerealizerNetworkItemInstanceDescriptorSerializedData,
    pub group_index: u32,
}

#[derive(Clone, Debug, PartialEq)]
pub struct CylinderData {
    pub radius_x: glam::Vec2,
    pub radius_z: glam::Vec2,
    pub height: f32,
    pub num_segments: u8,
}

#[derive(Clone, Debug, PartialEq)]
pub struct DataItemByte {
    pub r#type: DataItemType,
    pub value: i8,
}

#[derive(Clone, Debug, PartialEq)]
pub struct DataItemCompoundTag {
    pub r#type: DataItemType,
    pub value: Nbt,
}

#[derive(Clone, Debug, PartialEq)]
pub struct DataItemEntry {
    pub id: u32,
    pub payload: DataItemEntryValue,
}

#[derive(Clone, Debug, PartialEq)]
pub enum DataItemEntryValue {
    DataItemByte(DataItemByte),
    DataItemShort(DataItemShort),
    DataItemInt(DataItemInt),
    DataItemFloat(DataItemFloat),
    DataItemString(DataItemString),
    DataItemCompoundTag(DataItemCompoundTag),
    DataItemPos(DataItemPos),
    DataItemInt64(DataItemInt64),
    DataItemVec3(DataItemVec3),
}

#[derive(Clone, Debug, PartialEq)]
pub struct DataItemFloat {
    pub r#type: DataItemType,
    pub value: f32,
}

#[derive(Clone, Debug, PartialEq)]
pub struct DataItemInt {
    pub r#type: DataItemType,
    pub value: i32,
}

#[derive(Clone, Debug, PartialEq)]
pub struct DataItemInt64 {
    pub r#type: DataItemType,
    pub value: i64,
}

#[derive(Clone, Debug, PartialEq)]
pub struct DataItemPos {
    pub r#type: DataItemType,
    pub value: BlockPos,
}

#[derive(Clone, Debug, PartialEq)]
pub struct DataItemShort {
    pub r#type: DataItemType,
    pub value: i16,
}

#[derive(Clone, Debug, PartialEq)]
pub struct DataItemString {
    pub r#type: DataItemType,
    pub value: String,
}

#[derive(Clone, Debug, PartialEq)]
pub struct DataItemVec3 {
    pub r#type: DataItemType,
    pub value: glam::Vec3,
}

#[derive(Clone, Debug, PartialEq)]
pub struct DimensionDefinitionGroupDimensionDefinition {
    pub height_maximum: i32,
    pub height_minimum: i32,
    pub generator_type: GeneratorType,
    pub dimension_type: DimensionType,
    pub pack_id: uuid::Uuid,
}

#[derive(Clone, Debug, PartialEq)]
pub struct DimensionType {
    pub value: i32,
}

#[derive(Clone, Debug, PartialEq)]
pub enum DisconnectMessages {
    DisconnectPacketMessages(DisconnectPacketMessages),
    Empty1,
}

#[derive(Clone, Debug, PartialEq)]
pub struct DisconnectPacketMessages {
    pub message: String,
    pub filtered_message: String,
}

#[derive(Clone, Debug, PartialEq)]
pub enum EAS {
    BoolAttributeData(EASBoolAttributeData),
    FloatAttributeData(EASFloatAttributeData),
    ColorAttributeData(EASColorAttributeData),
}

#[derive(Clone, Debug, PartialEq)]
pub struct EASAttributeLayerData {
    pub name: String,
    pub noise_name: Option<String>,
    pub dimension: DimensionType,
    pub settings: EASAttributeLayerSettings,
    pub attributes: Vec<EASEnvironmentAttributeData>,
}

#[derive(Clone, Debug, PartialEq)]
pub struct EASAttributeLayerSettings {
    pub priority: i32,
    pub weight: f32,
    pub enabled: bool,
    pub transitions_paused: bool,
}

#[derive(Clone, Debug, PartialEq)]
pub struct EASBoolAttributeData {
    pub value: bool,
    pub operation: String,
}

#[derive(Clone, Debug, PartialEq)]
pub struct EASColorAttributeData {
    pub value: [i32; 4],
    pub operation: String,
}

#[derive(Clone, Debug, PartialEq)]
pub struct EASEnvironmentAttributeData {
    pub attribute_name: String,
    pub from_attribute: Option<EAS>,
    pub attribute: EAS,
    pub to_attribute: Option<EAS>,
    pub current_transition_ticks: u32,
    pub total_transition_ticks: u32,
    pub easing: String,
    pub local_transition_ticks: u32,
    pub noise_transition: bool,
}

#[derive(Clone, Debug, PartialEq)]
pub struct EASFloatAttributeData {
    pub value: f32,
    pub operation: String,
    pub constraint_min: Option<f32>,
    pub constraint_max: Option<f32>,
}

#[derive(Clone, Debug, PartialEq)]
pub struct ECSProfilingDiagnosticsEntityDiagnosticTimingInfo {
    pub display_name: String,
    pub entity: String,
    pub time_in_ns: u64,
    pub percent_of_total: u8,
}

#[derive(Clone, Debug, PartialEq)]
pub struct ECSProfilingDiagnosticsSystemCategory {
    pub category_name: String,
    pub system_index: u64,
}

#[derive(Clone, Debug, PartialEq)]
pub struct ECSProfilingDiagnosticsSystemDiagnosticTimingInfo {
    pub display_name: String,
    pub system_index: u64,
    pub time_in_ns: u64,
    pub percent_of_total: u8,
}

#[derive(Clone, Debug, PartialEq)]
pub struct EduSharedUriResource {
    pub button_name: String,
    pub link_uri: String,
}

#[derive(Clone, Debug, PartialEq)]
pub struct EducationLevelSettings {
    pub code_builder_default_uri: String,
    pub code_builder_title: String,
    pub can_resize_code_builder: bool,
    pub disable_legacy_title_bar: bool,
    pub post_process_filter: String,
    pub screenshot_border_resource_path: String,
    pub agent_capabilities: Option<AgentCapabilities>,
    pub local_settings: EducationLocalLevelSettings,
    pub deprecated_always_false: bool,
    pub external_link_settings: Option<ExternalLinkSettings>,
}

#[derive(Clone, Debug, PartialEq)]
pub struct EducationLocalLevelSettings {
    pub code_builder_override_uri: Option<String>,
}

#[derive(Clone, Debug, PartialEq)]
pub struct EllipsoidData {
    pub radii: glam::Vec3,
    pub segments_per_axis: u8,
}

#[derive(Clone, Debug, PartialEq)]
pub struct EnchantmentInstance {
    pub enchant_type: EnchantType,
    pub enchant_level: u8,
}

#[derive(Clone, Debug, PartialEq)]
pub struct EntityNetId {
    pub raw_id: u32,
}

#[derive(Clone, Debug, PartialEq)]
pub struct Experiments {
    pub toggles: Vec<CerealizerExperimentsAnonExperimentToggle>,
    pub experiments_ever_toggled: bool,
}

#[derive(Clone, Debug, PartialEq)]
pub struct ExternalLinkSettings {
    pub url: String,
    pub display_name: String,
}

#[derive(Clone, Debug, PartialEq)]
pub struct FeatureRegistryFeatureBinaryJsonFormat {
    pub feature_name: String,
    pub binary_json_output: String,
}

#[derive(Clone, Debug, PartialEq)]
pub struct FloatRange {
    pub min: f32,
    pub max: f32,
}

#[derive(Clone, Debug, PartialEq)]
pub struct FullContainerName {
    pub container_name: ContainerEnumName,
    pub dynamic_id: Option<u32>,
}

#[derive(Clone, Debug, PartialEq)]
pub struct GameRule {
    pub rule_name: String,
    pub rule_can_be_modified: bool,
    pub rule_value: GameRuleRuleValue,
}

#[derive(Clone, Debug, PartialEq)]
pub enum GameRuleRuleValue {
    Empty0,
    Bool(bool),
    Int32(i32),
    Float(f32),
}

#[derive(Clone, Debug, PartialEq)]
pub struct GameRulesChangedPacketData {
    pub rules_list: Vec<GameRule>,
}

#[derive(Clone, Debug, PartialEq)]
pub struct InventoryAction {
    pub source: InventorySource,
    pub slot: u32,
    pub from_item: CerealizerNetworkItemStackDescriptorSerializedData,
    pub to_item: CerealizerNetworkItemStackDescriptorSerializedData,
}

#[derive(Clone, Debug, PartialEq)]
pub struct InventoryMismatchData {
    pub actions: InventoryTransactionData,
}

#[derive(Clone, Debug, PartialEq)]
pub struct InventoryOptions {
    pub left_inventory_tab: InventoryLeftTabIndex,
    pub right_inventory_tab: InventoryRightTabIndex,
    pub filtering: bool,
    pub layout_inv: InventoryLayout,
    pub layout_craft: InventoryLayout,
}

#[derive(Clone, Debug, PartialEq)]
pub struct InventorySource {
    pub source_type: InventorySourceType,
    pub container_id: Option<Option<i8>>,
    pub bit_flags: Option<Option<InventorySourceInventorySourceFlags>>,
}

#[derive(Clone, Debug, PartialEq)]
pub struct InventoryTransactionData {
    pub actions: Option<Vec<InventoryAction>>,
}

#[derive(Clone, Debug, PartialEq)]
pub enum InventoryTransactionTransactionValue {
    NormalTransactionData(NormalTransactionData),
    InventoryMismatchData(InventoryMismatchData),
    ItemUseInventoryTransaction(ItemUseInventoryTransaction),
    ItemUseOnActorInventoryTransaction(ItemUseOnActorInventoryTransaction),
    ItemReleaseInventoryTransaction(ItemReleaseInventoryTransaction),
}

#[derive(Clone, Debug, PartialEq)]
pub struct ItemData {
    pub item_name: String,
    pub item_id: i16,
    pub is_component_based: bool,
    pub item_version: ItemVersion,
    pub item_component_data: Nbt,
}

#[derive(Clone, Debug, PartialEq)]
pub struct ItemEnchantOption {
    pub cost: u8,
    pub enchants: ItemEnchants,
    pub enchant_name: String,
    pub enchant_net_id: TypedServerNetIdStructRecipeNetIdTag,
}

#[derive(Clone, Debug, PartialEq)]
pub struct ItemEnchants {
    pub slot: i32,
    pub item_enchants: [Vec<EnchantmentInstance>; 3],
}

#[derive(Clone, Debug, PartialEq)]
pub struct ItemReleaseInventoryTransaction {
    pub actions: InventoryTransactionData,
    pub action_type: ItemReleaseInventoryTransactionActionType,
    pub slot: i32,
    pub item: CerealizerNetworkItemStackDescriptorSerializedData,
    pub from_position: glam::Vec3,
}

#[derive(Clone, Debug, PartialEq)]
pub enum ItemStackRequestCereal {
    TakeActionData(ItemStackRequestCerealTakeActionData),
    PlaceActionData(ItemStackRequestCerealPlaceActionData),
    SwapActionData(ItemStackRequestCerealSwapActionData),
    DropActionData(ItemStackRequestCerealDropActionData),
    DestroyActionData(ItemStackRequestCerealDestroyActionData),
    ConsumeActionData(ItemStackRequestCerealConsumeActionData),
    CreateActionData(ItemStackRequestCerealCreateActionData),
    LabTableCombineActionData(ItemStackRequestCerealLabTableCombineActionData),
    BeaconPaymentActionData(ItemStackRequestCerealBeaconPaymentActionData),
    MineBlockActionData(ItemStackRequestCerealMineBlockActionData),
    CraftRecipeActionData(ItemStackRequestCerealCraftRecipeActionData),
    CraftRecipeAutoActionData(ItemStackRequestCerealCraftRecipeAutoActionData),
    CraftCreativeActionData(ItemStackRequestCerealCraftCreativeActionData),
    CraftRecipeOptionalActionData(ItemStackRequestCerealCraftRecipeOptionalActionData),
    CraftRepairAndDisenchantActionData(ItemStackRequestCerealCraftRepairAndDisenchantActionData),
    CraftLoomActionData(ItemStackRequestCerealCraftLoomActionData),
    CraftNonImplementedActionData(ItemStackRequestCerealCraftNonImplementedActionData),
    CraftResultsActionData(ItemStackRequestCerealCraftResultsActionData),
}

#[derive(Clone, Debug, PartialEq)]
pub struct ItemStackRequestCerealBeaconPaymentActionData {
    pub action_type: ItemStackRequestActionType,
    pub primary_effect_id: i32,
    pub secondary_effect_id: i32,
}

#[derive(Clone, Debug, PartialEq)]
pub struct ItemStackRequestCerealConsumeActionData {
    pub action_type: ItemStackRequestActionType,
    pub amount: u8,
    pub source: ItemStackRequestCerealSlotInfoData,
}

#[derive(Clone, Debug, PartialEq)]
pub struct ItemStackRequestCerealCraftCreativeActionData {
    pub action_type: ItemStackRequestActionType,
    pub creative_item_net_id: u32,
    pub number_of_requested_crafts: u8,
}

#[derive(Clone, Debug, PartialEq)]
pub struct ItemStackRequestCerealCraftLoomActionData {
    pub action_type: ItemStackRequestActionType,
    pub pattern_name_id: String,
    pub num_crafts: u8,
}

#[derive(Clone, Debug, PartialEq)]
pub struct ItemStackRequestCerealCraftNonImplementedActionData {
    pub action_type: ItemStackRequestActionType,
}

#[derive(Clone, Debug, PartialEq)]
pub struct ItemStackRequestCerealCraftRecipeActionData {
    pub action_type: ItemStackRequestActionType,
    pub recipe_net_id: TypedServerNetIdStructRecipeNetIdTag,
    pub number_of_requested_crafts: u8,
}

#[derive(Clone, Debug, PartialEq)]
pub struct ItemStackRequestCerealCraftRecipeAutoActionData {
    pub action_type: ItemStackRequestActionType,
    pub recipe_net_id: TypedServerNetIdStructRecipeNetIdTag,
    pub number_of_requested_crafts: u8,
    pub ingredients: Vec<ItemStackRequestCerealRecipeIngredientData>,
}

#[derive(Clone, Debug, PartialEq)]
pub struct ItemStackRequestCerealCraftRecipeOptionalActionData {
    pub action_type: ItemStackRequestActionType,
    pub recipe_net_id: TypedServerNetIdStructRecipeNetIdTag,
    pub filtered_string_index: i32,
}

#[derive(Clone, Debug, PartialEq)]
pub struct ItemStackRequestCerealCraftRepairAndDisenchantActionData {
    pub action_type: ItemStackRequestActionType,
    pub recipe_net_id: i32,
    pub number_of_requested_crafts: u8,
    pub repair_cost: i32,
}

#[derive(Clone, Debug, PartialEq)]
pub struct ItemStackRequestCerealCraftResultsActionData {
    pub action_type: ItemStackRequestActionType,
    pub craft_results: Vec<ItemStackRequestCerealNetworkItemInstanceDescriptorData>,
    pub num_crafts: u8,
}

#[derive(Clone, Debug, PartialEq)]
pub struct ItemStackRequestCerealCreateActionData {
    pub action_type: ItemStackRequestActionType,
    pub results_index: u8,
}

#[derive(Clone, Debug, PartialEq)]
pub struct ItemStackRequestCerealDestroyActionData {
    pub action_type: ItemStackRequestActionType,
    pub amount: u8,
    pub source: ItemStackRequestCerealSlotInfoData,
}

#[derive(Clone, Debug, PartialEq)]
pub struct ItemStackRequestCerealDropActionData {
    pub action_type: ItemStackRequestActionType,
    pub amount: u8,
    pub source: ItemStackRequestCerealSlotInfoData,
    pub randomly: bool,
}

#[derive(Clone, Debug, PartialEq)]
pub struct ItemStackRequestCerealEmptyItemDescriptorData {
    pub descriptor_type: ItemStackRequestCerealItemDescriptorType,
}

#[derive(Clone, Debug, PartialEq)]
pub struct ItemStackRequestCerealItemNameDescriptorData {
    pub descriptor_type: ItemStackRequestCerealItemDescriptorType,
    pub full_name: String,
    pub aux_value: i32,
}

#[derive(Clone, Debug, PartialEq)]
pub struct ItemStackRequestCerealItemTagDescriptorData {
    pub descriptor_type: ItemStackRequestCerealItemDescriptorType,
    pub item_tag: String,
}

#[derive(Clone, Debug, PartialEq)]
pub struct ItemStackRequestCerealLabTableCombineActionData {
    pub action_type: ItemStackRequestActionType,
}

#[derive(Clone, Debug, PartialEq)]
pub struct ItemStackRequestCerealMineBlockActionData {
    pub action_type: ItemStackRequestActionType,
    pub slot: i32,
    pub predicted_durability: i32,
    pub net_id_variant: i32,
}

#[derive(Clone, Debug, PartialEq)]
pub struct ItemStackRequestCerealMoLangItemDescriptorData {
    pub descriptor_type: ItemStackRequestCerealItemDescriptorType,
    pub tag_expression: String,
    pub molang_version: MoLangVersion,
}

#[derive(Clone, Debug, PartialEq)]
pub struct ItemStackRequestCerealNetworkItemInstanceDescriptorData {
    pub item_descriptor: ItemStackRequestCerealRecipeIngredientDataItemDescriptor,
    pub stack_size: u16,
    pub block_runtime_id: u32,
    pub user_data_buffer: String,
}

#[derive(Clone, Debug, PartialEq)]
pub struct ItemStackRequestCerealPlaceActionData {
    pub action_type: ItemStackRequestActionType,
    pub amount: u8,
    pub source: ItemStackRequestCerealSlotInfoData,
    pub destination: ItemStackRequestCerealSlotInfoData,
}

#[derive(Clone, Debug, PartialEq)]
pub struct ItemStackRequestCerealRecipeIngredientData {
    pub item_descriptor: ItemStackRequestCerealRecipeIngredientDataItemDescriptor,
    pub stack_size: u16,
}

#[derive(Clone, Debug, PartialEq)]
pub enum ItemStackRequestCerealRecipeIngredientDataItemDescriptor {
    EmptyItemDescriptorData(ItemStackRequestCerealEmptyItemDescriptorData),
    ItemNameDescriptorData(ItemStackRequestCerealItemNameDescriptorData),
    MolangItemDescriptorData(ItemStackRequestCerealMoLangItemDescriptorData),
    ItemTagDescriptorData(ItemStackRequestCerealItemTagDescriptorData),
}

#[derive(Clone, Debug, PartialEq)]
pub struct ItemStackRequestCerealRequestData {
    pub client_request_id: TypedClientNetIdStructItemStackRequestIdTagInt32T0,
    pub actions: Vec<ItemStackRequestCereal>,
    pub strings_to_filter: Vec<String>,
    pub strings_to_filter_origin: TextProcessingEventOrigin,
}

#[derive(Clone, Debug, PartialEq)]
pub struct ItemStackRequestCerealSlotInfoData {
    pub full_container_name: FullContainerName,
    pub slot: u8,
    pub net_id_variant: i32,
}

#[derive(Clone, Debug, PartialEq)]
pub struct ItemStackRequestCerealSwapActionData {
    pub action_type: ItemStackRequestActionType,
    pub source: ItemStackRequestCerealSlotInfoData,
    pub destination: ItemStackRequestCerealSlotInfoData,
}

#[derive(Clone, Debug, PartialEq)]
pub struct ItemStackRequestCerealTakeActionData {
    pub action_type: ItemStackRequestActionType,
    pub amount: u8,
    pub source: ItemStackRequestCerealSlotInfoData,
    pub destination: ItemStackRequestCerealSlotInfoData,
}

#[derive(Clone, Debug, PartialEq)]
pub struct ItemStackRequestPacketDataRequestData {
    pub client_request_id: TypedClientNetIdStructItemStackRequestIdTagInt32T0,
    pub actions: Vec<ItemStackRequestCereal>,
    pub strings_to_filter: Vec<String>,
    pub strings_to_filter_origin: TextProcessingEventOrigin,
}

#[derive(Clone, Debug, PartialEq)]
pub struct ItemStackResponseContainerInfo {
    pub full_container_name: FullContainerName,
    pub slots: Vec<ItemStackResponseSlotInfo>,
}

#[derive(Clone, Debug, PartialEq)]
pub struct ItemStackResponseInfo {
    pub result: ItemStackNetResult,
    pub client_request_id: TypedClientNetIdStructItemStackRequestIdTagInt32T0,
    pub containers: Option<Option<Vec<ItemStackResponseContainerInfo>>>,
}

#[derive(Clone, Debug, PartialEq)]
pub struct ItemStackResponseSlotInfo {
    pub requested_slot: u8,
    pub slot: u8,
    pub amount: u8,
    pub item_stack_net_id: Option<Option<TypedServerNetIdStructItemStackNetIdTagInt32T0>>,
    pub custom_name: BedrockSafetyRedactableString,
    pub durability_correction: i32,
}

#[derive(Clone, Debug, PartialEq)]
pub struct ItemUseInventoryTransaction {
    pub actions: InventoryTransactionData,
    pub action_type: ItemUseInventoryTransactionActionType,
    pub trigger_type: ItemUseInventoryTransactionTriggerType,
    pub position: BlockPos,
    pub face: u8,
    pub slot: i32,
    pub item: CerealizerNetworkItemStackDescriptorSerializedData,
    pub from_position: glam::Vec3,
    pub click_position: glam::Vec3,
    pub target_block_id: u32,
    pub client_interact_prediction: ItemUseInventoryTransactionPredictedResult,
    pub client_cooldown_state: ItemUseInventoryTransactionClientCooldownState,
}

#[derive(Clone, Debug, PartialEq)]
pub struct ItemUseOnActorInventoryTransaction {
    pub actions: InventoryTransactionData,
    pub runtime_id: ActorRuntimeID,
    pub action_type: ItemUseOnActorInventoryTransactionActionType,
    pub slot: i32,
    pub item: CerealizerNetworkItemStackDescriptorSerializedData,
    pub from_position: glam::Vec3,
    pub hit_position: glam::Vec3,
}

#[derive(Clone, Debug, PartialEq)]
pub struct LegacySetSlot {
    pub container_enum: ContainerEnumName,
    pub slots: Vec<u8>,
}

#[derive(Clone, Debug, PartialEq)]
pub struct LegacyTelemetryEventAchievement {
    pub achievement_id: MinecraftEventingAchievementIds,
}

#[derive(Clone, Debug, PartialEq)]
pub struct LegacyTelemetryEventActorDefinition {
    pub event_name: String,
}

#[derive(Clone, Debug, PartialEq)]
pub struct LegacyTelemetryEventBellUsed {
    pub item_id: i32,
}

#[derive(Clone, Debug, PartialEq)]
pub struct LegacyTelemetryEventBossKilled {
    pub boss_actor_id: i64,
    pub party_size: i32,
    pub boss_type: i32,
}

#[derive(Clone, Debug, PartialEq)]
pub struct LegacyTelemetryEventCauldronUsed {
    pub contents_color: u32,
    pub contents_type: i32,
    pub fill_level: i32,
}

#[derive(Clone, Debug, PartialEq)]
pub struct LegacyTelemetryEventCodeBuilderRuntimeAction {
    pub code_builder_runtime_action: String,
}

#[derive(Clone, Debug, PartialEq)]
pub struct LegacyTelemetryEventCodeBuilderScoreboard {
    pub objective_name: String,
    pub score: i32,
}

#[derive(Clone, Debug, PartialEq)]
pub struct LegacyTelemetryEventComposterUsed {
    pub block_interaction_type: MinecraftEventingPOIBlockInteractionType,
    pub item_id: i32,
}

#[derive(Clone, Debug, PartialEq)]
pub enum LegacyTelemetryEventEventData {
    Achievement(LegacyTelemetryEventAchievement),
    Interaction(LegacyTelemetryEventInteraction),
    PortalCreated(LegacyTelemetryEventPortalCreated),
    PortalUsed(LegacyTelemetryEventPortalUsed),
    MobKilled(LegacyTelemetryEventMobKilled),
    CauldronUsed(LegacyTelemetryEventCauldronUsed),
    PlayerDied(LegacyTelemetryEventPlayerDied),
    BossKilled(LegacyTelemetryEventBossKilled),
    SlashCommand(LegacyTelemetryEventSlashCommand),
    MobBorn(LegacyTelemetryEventMobBorn),
    PoiCauldronUsed(LegacyTelemetryEventPOICauldronUsed),
    ComposterUsed(LegacyTelemetryEventComposterUsed),
    BellUsed(LegacyTelemetryEventBellUsed),
    ActorDefinition(LegacyTelemetryEventActorDefinition),
    RaidUpdate(LegacyTelemetryEventRaidUpdate),
    TargetBlockHit(LegacyTelemetryEventTargetBlockHit),
    PiglinBarter(LegacyTelemetryEventPiglinBarter),
    PlayerWaxedOrUnwaxedCopper(LegacyTelemetryEventPlayerWaxedOrUnwaxedCopper),
    CodeBuilderRuntimeAction(LegacyTelemetryEventCodeBuilderRuntimeAction),
    CodeBuilderScoreboard(LegacyTelemetryEventCodeBuilderScoreboard),
    ItemUsed(LegacyTelemetryEventItemUsed),
    Empty,
}

#[derive(Clone, Debug, PartialEq)]
pub struct LegacyTelemetryEventInteraction {
    pub interacted_entity_id: i64,
    pub interaction_type: MinecraftEventingInteractionType,
    pub interaction_actor_type: i32,
    pub interaction_actor_variant: i32,
    pub interaction_actor_color: u8,
}

#[derive(Clone, Debug, PartialEq)]
pub struct LegacyTelemetryEventItemUsed {
    pub item_id: i16,
    pub item_aux: i32,
    pub use_method: i32,
    pub count: i32,
}

#[derive(Clone, Debug, PartialEq)]
pub struct LegacyTelemetryEventMobBorn {
    pub born_baby_entity_type: i32,
    pub born_baby_entity_variant: i32,
    pub born_baby_color: u8,
}

#[derive(Clone, Debug, PartialEq)]
pub struct LegacyTelemetryEventMobKilled {
    pub instigator_actor_id: i64,
    pub target_actor_id: i64,
    pub instigator_s_child_actor_type: ActorType,
    pub damage_source: i32,
    pub trade_tier: i32,
    pub trader_name: String,
}

#[derive(Clone, Debug, PartialEq)]
pub struct LegacyTelemetryEventPOICauldronUsed {
    pub block_interaction_type: MinecraftEventingPOIBlockInteractionType,
    pub item_id: i32,
}

#[derive(Clone, Debug, PartialEq)]
pub struct LegacyTelemetryEventPiglinBarter {
    pub item_id: i32,
    pub was_targeting_bartering_player: bool,
}

#[derive(Clone, Debug, PartialEq)]
pub struct LegacyTelemetryEventPlayerDied {
    pub instigator_actor_id: i32,
    pub instigator_mob_variant: i32,
    pub damage_source: i32,
    pub died_in_raid: bool,
}

#[derive(Clone, Debug, PartialEq)]
pub struct LegacyTelemetryEventPlayerWaxedOrUnwaxedCopper {
    pub player_waxed_or_unwaxed_copper_block_id: i32,
}

#[derive(Clone, Debug, PartialEq)]
pub struct LegacyTelemetryEventPortalCreated {
    pub dimension_id: i32,
}

#[derive(Clone, Debug, PartialEq)]
pub struct LegacyTelemetryEventPortalUsed {
    pub source_dimension_id: i32,
    pub target_dimension_id: i32,
}

#[derive(Clone, Debug, PartialEq)]
pub struct LegacyTelemetryEventRaidUpdate {
    pub current_wave: i32,
    pub total_waves: i32,
    pub success: bool,
}

#[derive(Clone, Debug, PartialEq)]
pub struct LegacyTelemetryEventSlashCommand {
    pub success_count: i32,
    pub error_count: i32,
    pub command_name: String,
    pub error_list: String,
}

#[derive(Clone, Debug, PartialEq)]
pub struct LegacyTelemetryEventTargetBlockHit {
    pub redstone_level: i32,
}

#[derive(Clone, Debug, PartialEq)]
pub struct LevelChunkSubChunkMetadata {
    pub blob_id: u64,
}

#[derive(Clone, Debug, PartialEq)]
pub struct LevelSettings {
    pub seed: u64,
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
    pub day_cycle_stop_time: i32,
    pub education_edition_offer: EducationEditionOffer,
    pub education_features_enabled: bool,
    pub education_product_id: String,
    pub rain_level: f32,
    pub lightning_level: f32,
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
    pub server_chunk_tick_range: i32,
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
    pub limited_world_width: i32,
    pub limited_world_depth: i32,
    pub nether_type: bool,
    pub edu_shared_uri_resource: EduSharedUriResource,
    pub override_force_experimental_gameplay: Option<bool>,
    pub chat_restriction_level: ChatRestrictionLevel,
    pub disable_player_interactions: bool,
    pub server_editor_connection_policy: ServerEditorConnectionPolicy,
    pub allow_anonymous_block_drops_in_editor_worlds: bool,
}

#[derive(Clone, Debug, PartialEq)]
pub struct LineData {
    pub line_end_location: glam::Vec3,
}

#[derive(Clone, Debug, PartialEq)]
pub struct LocatorBarWaypoint {
    pub group_handle: WaypointGroupWaypointHandle,
    pub server_waypoint_payload: ServerWaypoint,
    pub action_flag: ServerWaypointGroupAction,
}

#[derive(Clone, Debug, PartialEq)]
pub struct MapDecoration {
    pub image_type: MapDecorationType,
    pub rotation: u8,
    pub x: u8,
    pub y: u8,
    pub label: String,
    pub color: MceColor,
}

#[derive(Clone, Debug, PartialEq)]
pub struct MapInfoRequestPacketAnonClientPixelsProxy {
    pub pixel: u32,
    pub index: u16,
}

#[derive(Clone, Debug, PartialEq)]
pub struct MapItemTrackedActorUniqueId {
    pub r#type: MapItemTrackedActorType,
    pub entity_id: Option<ActorUniqueID>,
    pub block_position: Option<BlockPos>,
}

#[derive(Clone, Debug, PartialEq)]
pub struct MaterialReducerDataEntry {
    pub from_item_key: i32,
    pub item_ids_and_counts: Vec<MaterialReducerEntryOutput>,
}

#[derive(Clone, Debug, PartialEq)]
pub struct MaterialReducerEntryOutput {
    pub item_id: i32,
    pub item_count: i32,
}

#[derive(Clone, Debug, PartialEq)]
pub struct MceColor {
    pub color: i32,
}

#[derive(Clone, Debug, PartialEq)]
pub struct MemoryMemoryCategoryCounter {
    pub category: MemoryMemoryCategory,
    pub current_bytes: u64,
}

#[derive(Clone, Debug, PartialEq)]
pub struct MissingBlobData {
    pub blob_id: u64,
    pub blob_data: String,
}

#[derive(Clone, Debug, PartialEq)]
pub struct MoveActorAbsoluteData {
    pub actor_runtime_id: ActorRuntimeID,
    pub header: u8,
    pub position: glam::Vec3,
    pub rotation_x: u8,
    pub rotation_y: u8,
    pub rotation_y_head: u8,
}

#[derive(Clone, Debug, PartialEq)]
pub struct MoveActorDeltaData {
    pub actor_runtime_id: ActorRuntimeID,
    pub new_position_x: Option<f32>,
    pub new_position_y: Option<f32>,
    pub new_position_z: Option<f32>,
    pub rotation_x: Option<i8>,
    pub rotation_y: Option<i8>,
    pub rotation_y_head: Option<i8>,
    pub is_on_ground: bool,
    pub force_move: bool,
    pub force_move_local_entity: bool,
    pub force_completion: bool,
}

#[derive(Clone, Debug, PartialEq)]
pub struct MovePlayerTeleportData {
    pub teleportation_cause: i32,
    pub source_actor_type: i32,
}

#[derive(Clone, Debug, PartialEq)]
pub struct MultiRecipe {
    pub multi_recipe_uuid: uuid::Uuid,
    pub net_id: TypedServerNetIdStructRecipeNetIdTag,
}

#[derive(Clone, Debug, PartialEq)]
pub struct NetworkPermissions {
    pub server_auth_sound_enabled: bool,
}

#[derive(Clone, Debug, PartialEq)]
pub struct NoiseDescriptor {
    pub name: String,
    pub first_octave: i32,
    pub amplitudes: Vec<f32>,
}

#[derive(Clone, Debug, PartialEq)]
pub struct NormalTransactionData {
    pub actions: InventoryTransactionData,
}

#[derive(Clone, Debug, PartialEq)]
pub struct PackIdVersion {
    pub pack_uuid: uuid::Uuid,
    pub pack_version: SemVersion,
}

#[derive(Clone, Debug, PartialEq)]
pub struct PackIdVersionData {
    pub pack_uuid: uuid::Uuid,
    pub pack_version: SemVersionData,
}

#[derive(Clone, Debug, PartialEq)]
pub struct PackInfoData {
    pub pack_id_version: PackIdVersionData,
    pub pack_size: u64,
    pub content_key: String,
    pub subpack_name: String,
    pub content_identity: ContentIdentity,
    pub has_scripts: bool,
    pub is_addon_pack: bool,
    pub is_ray_tracing_capable: bool,
    pub cdn_url: String,
}

#[derive(Clone, Debug, PartialEq)]
pub struct PackInstanceId {
    pub pack_id: String,
    pub version: String,
    pub sub_pack_name: String,
}

#[derive(Clone, Debug, PartialEq)]
pub struct PackedItemUseLegacyInventoryTransaction {
    pub legacy_request_id: TypedClientNetIdStructItemStackLegacyRequestIdTagInt32T0,
    pub legacy_set_item_slots: Option<Vec<LegacySetSlot>>,
    pub item_use_transaction: Option<ItemUseInventoryTransaction>,
}

#[derive(Clone, Debug, PartialEq)]
pub struct PlayerBlockActionData {
    pub player_action_type: PlayerActionType,
    pub position: BlockPos,
    pub facing: i32,
}

#[derive(Clone, Debug, PartialEq)]
pub struct PlayerInputTick {
    pub input_tick: u64,
}

#[derive(Clone, Debug, PartialEq)]
pub struct PlayerListAddEntry {
    pub uuid: uuid::Uuid,
    pub actor_unique_id: ActorUniqueID,
    pub player_name: String,
    pub xbl_xuid: String,
    pub platform_online_id: String,
    pub build_platform: BuildPlatform,
    pub serialized_skin: SerializedSkinRef,
    pub is_teacher: bool,
    pub is_host: bool,
    pub is_sub_client: bool,
    pub player_color: MceColor,
}

#[derive(Clone, Debug, PartialEq)]
pub enum PlayerListEntriesItem {
    Add(PlayerListAddEntry),
    Remove(PlayerListRemoveEntry),
}

#[derive(Clone, Debug, PartialEq)]
pub struct PlayerListRemoveEntry {
    pub uuid: uuid::Uuid,
}

#[derive(Clone, Debug, PartialEq)]
pub struct PlayerLocationCoordinatesLocation {
    pub packet_type: PlayerLocationType,
    pub position: glam::Vec3,
}

#[derive(Clone, Debug, PartialEq)]
pub struct PlayerLocationHiddenLocation {
    pub packet_type: PlayerLocationType,
}

#[derive(Clone, Debug, PartialEq)]
pub enum PlayerLocationLocation {
    PlayerLocationCoordinates(PlayerLocationCoordinatesLocation),
    PlayerLocationHide(PlayerLocationHiddenLocation),
}

#[derive(Clone, Debug, PartialEq)]
pub struct PlayerPartyInfo {
    pub party_id: String,
    pub is_party_leader: bool,
}

#[derive(Clone, Debug, PartialEq)]
pub struct PlayerScoreboardId {
    pub player_unique_id: i64,
}

#[derive(Clone, Debug, PartialEq)]
pub struct PlayerUpdateEntityOverridesClearOverride {
    pub r#type: String,
}

#[derive(Clone, Debug, PartialEq)]
pub struct PlayerUpdateEntityOverridesFloatOverride {
    pub r#type: String,
    pub value: f32,
}

#[derive(Clone, Debug, PartialEq)]
pub struct PlayerUpdateEntityOverridesIntOverride {
    pub r#type: String,
    pub value: i32,
}

#[derive(Clone, Debug, PartialEq)]
pub struct PlayerUpdateEntityOverridesRemoveOverride {
    pub r#type: String,
}

#[derive(Clone, Debug, PartialEq)]
pub enum PlayerUpdateEntityOverridesUpdate {
    ClearOverride(PlayerUpdateEntityOverridesClearOverride),
    RemoveOverride(PlayerUpdateEntityOverridesRemoveOverride),
    IntOverride(PlayerUpdateEntityOverridesIntOverride),
    FloatOverride(PlayerUpdateEntityOverridesFloatOverride),
}

#[derive(Clone, Debug, PartialEq)]
pub enum PlayerVideoCaptureAction {
    StopVideoCapture,
    StartVideoCapture(PlayerVideoCaptureStartVideoCapture),
}

#[derive(Clone, Debug, PartialEq)]
pub struct PlayerVideoCaptureStartVideoCapture {
    pub frame_rate: u32,
    pub file_prefix: String,
}

#[derive(Clone, Debug, PartialEq)]
pub struct PositionTrackingId {
    pub value: i32,
}

#[derive(Clone, Debug, PartialEq)]
pub struct PotionMixDataEntry {
    pub from_potion_id: i32,
    pub from_item_aux: i32,
    pub reagent_item_id: i32,
    pub reagent_item_aux: i32,
    pub to_potion_id: i32,
    pub to_item_aux: i32,
}

#[derive(Clone, Debug, PartialEq)]
pub struct PrimitiveShapeData {
    pub network_id: u64,
    pub shape_type: Option<ScriptModuleMinecraftScriptPrimitiveShapeType>,
    pub location: Option<glam::Vec3>,
    pub scale: Option<f32>,
    pub rotation: Option<glam::Vec3>,
    pub total_time_left: Option<f32>,
    pub maximum_render_distance: Option<f32>,
    pub color: Option<MceColor>,
    pub dimension_id: Option<DimensionType>,
    pub attached_to_entity_id: Option<ActorUniqueID>,
    pub extra_shape_data: PrimitiveShapeDataExtraShapeData,
}

#[derive(Clone, Debug, PartialEq)]
pub enum PrimitiveShapeDataExtraShapeData {
    Empty0,
    ArrowData(ArrowData),
    TextData(TextData),
    BoxData(BoxData),
    LineData(LineData),
    SphereData(SphereData),
    CylinderData(CylinderData),
    PyramidData(PyramidData),
    EllipsoidData(EllipsoidData),
    ConeData(ConeData),
}

#[derive(Clone, Debug, PartialEq)]
pub struct PropertySyncData {
    pub int_entries_list: Vec<PropertySyncDataPropertySyncIntEntry>,
    pub float_entries_list: Vec<PropertySyncDataPropertySyncFloatEntry>,
}

#[derive(Clone, Debug, PartialEq)]
pub struct PropertySyncDataPropertySyncFloatEntry {
    pub property_index: u32,
    pub data: f32,
}

#[derive(Clone, Debug, PartialEq)]
pub struct PropertySyncDataPropertySyncIntEntry {
    pub property_index: u32,
    pub data: i32,
}

#[derive(Clone, Debug, PartialEq)]
pub struct PyramidData {
    pub width: f32,
    pub depth: Option<f32>,
    pub height: f32,
}

#[derive(Clone, Debug, PartialEq)]
pub struct RemoveScore {
    pub action: String,
    pub scoreboard_id: ScoreboardId,
    pub objective_name: Option<String>,
}

#[derive(Clone, Debug, PartialEq)]
pub struct ResourcePackClientResponseCancel {
    pub response_type: String,
}

#[derive(Clone, Debug, PartialEq)]
pub struct ResourcePackClientResponseDownloading {
    pub response_type: String,
    pub downloading_packs: Vec<String>,
}

#[derive(Clone, Debug, PartialEq)]
pub struct ResourcePackClientResponseDownloadingFinished {
    pub response_type: String,
}

#[derive(Clone, Debug, PartialEq)]
pub struct ResourcePackClientResponseResourcePackStackFinished {
    pub response_type: String,
}

#[derive(Clone, Debug, PartialEq)]
pub enum ResourcePackClientResponseResponse {
    Cancel(ResourcePackClientResponseCancel),
    Downloading(ResourcePackClientResponseDownloading),
    DownloadingFinished(ResourcePackClientResponseDownloadingFinished),
    ResourcePackStackFinished(ResourcePackClientResponseResourcePackStackFinished),
}

#[derive(Clone, Debug, PartialEq)]
pub struct ScoreboardId {
    pub scoreboard_id: i64,
}

#[derive(Clone, Debug, PartialEq)]
pub struct ScoreboardIdentityPacketInfo {
    pub scoreboard_id: ScoreboardId,
    pub player_unique_id: Option<i64>,
}

#[derive(Clone, Debug, PartialEq)]
pub struct SemVersion {
    pub version: String,
}

#[derive(Clone, Debug, PartialEq)]
pub struct SemVersionData {
    pub version: String,
}

#[derive(Clone, Debug, PartialEq)]
pub struct SerializedAbilitiesData {
    pub target_player_raw_id: i64,
    pub player_permissions: PlayerPermissionLevel,
    pub command_permissions: CommandPermissionLevel,
    pub layers: Vec<SerializedAbilitiesDataSerializedLayer>,
}

#[derive(Clone, Debug, PartialEq)]
pub struct SerializedAbilitiesDataSerializedLayer {
    pub serialized_layer: u16,
    pub abilities_set: u32,
    pub ability_values: u32,
    pub fly_speed: f32,
    pub vertical_fly_speed: f32,
    pub walk_speed: f32,
}

#[derive(Clone, Debug, PartialEq)]
pub struct SerializedNoiseBlockSpecifier {
    pub noise: String,
    pub threshold: f32,
    pub range: FloatRange,
    pub block: u32,
}

#[derive(Clone, Debug, PartialEq)]
pub struct SerializedPersonaPieceHandle {
    pub piece_id: String,
    pub piece_type: PersonaPieceType,
    pub pack_id: uuid::Uuid,
    pub is_default_piece: bool,
    pub product_id: String,
}

#[derive(Clone, Debug, PartialEq)]
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

#[derive(Clone, Debug, PartialEq)]
pub struct ServerBlockProperty {
    pub block_name: String,
    pub block_definition: Nbt,
}

#[derive(Clone, Debug, PartialEq)]
pub struct ServerConfigurationClientStoreEntryPointConfiguration {
    pub store_id: String,
    pub store_name: String,
}

#[derive(Clone, Debug, PartialEq)]
pub struct ServerConfigurationGatheringsConfigurationJoinInfo {
    pub experience_id: uuid::Uuid,
    pub experience_name: String,
    pub world_id: Option<uuid::Uuid>,
    pub world_name: Option<String>,
    pub creator_id: String,
    pub target_id: Option<uuid::Uuid>,
    pub scenario_id: Option<String>,
    pub server_id: Option<String>,
}

#[derive(Clone, Debug, PartialEq)]
pub struct ServerConfigurationPresenceConfiguration {
    pub rich_presence_id: Option<String>,
}

#[derive(Clone, Debug, PartialEq)]
pub struct ServerConfigurationServerConfigurationJoinInfo {
    pub gathering: Option<ServerConfigurationGatheringsConfigurationJoinInfo>,
    pub client_store_entry_point: Option<ServerConfigurationClientStoreEntryPointConfiguration>,
    pub presence: Option<ServerConfigurationPresenceConfiguration>,
}

#[derive(Clone, Debug, PartialEq)]
pub struct ServerSoundHandle {
    pub server_sound_handle: u64,
}

#[derive(Clone, Debug, PartialEq)]
pub struct ServerWaypoint {
    pub update_flag: u32,
    pub is_visible: Option<bool>,
    pub world_position: Option<WorldPosition>,
    pub texture_path: Option<String>,
    pub icon_size: Option<glam::Vec2>,
    pub color: Option<MceColor>,
    pub client_position_authority: Option<bool>,
    pub actor_unique_id: Option<ActorUniqueID>,
}

#[derive(Clone, Debug, PartialEq)]
pub enum ServerboundPackSettingChangePackSettingValue {
    Float(f32),
    Bool(bool),
    String(String),
}

#[derive(Clone, Debug, PartialEq)]
pub enum SetScoreScoreInfoItem {
    RemoveScore(RemoveScore),
    ChangePlayerScore(ChangePlayerScore),
    ChangeEntityScore(ChangeEntityScore),
    ChangeFakePlayerScore(ChangeFakePlayerScore),
}

#[derive(Clone, Debug, PartialEq)]
pub struct ShapedRecipe {
    pub recipe_id: String,
    pub width: i32,
    pub height: i32,
    pub ingredients: Vec<CerealizerRecipeIngredientSerializedData>,
    pub results: Vec<CerealizerNetworkItemInstanceDescriptorSerializedData>,
    pub uuid: uuid::Uuid,
    pub tag: String,
    pub priority: i32,
    pub assume_symmetry: bool,
    pub unlocking_requirement: Option<CerealizerRecipeUnlockingRequirementSerializedData>,
    pub net_id: TypedServerNetIdStructRecipeNetIdTag,
}

#[derive(Clone, Debug, PartialEq)]
pub struct ShapelessRecipe {
    pub recipe_id: String,
    pub ingredients: Vec<CerealizerRecipeIngredientSerializedData>,
    pub results: Vec<CerealizerNetworkItemInstanceDescriptorSerializedData>,
    pub uuid: uuid::Uuid,
    pub tag: String,
    pub priority: i32,
    pub unlocking_requirement: Option<CerealizerRecipeUnlockingRequirementSerializedData>,
    pub net_id: TypedServerNetIdStructRecipeNetIdTag,
}

#[derive(Clone, Debug, PartialEq)]
pub struct SkinImage {
    pub width: u32,
    pub height: u32,
    pub image_bytes: Vec<u8>,
}

#[derive(Clone, Debug, PartialEq)]
pub struct SmithingTransformRecipe {
    pub recipe_id: String,
    pub template_ingredient: CerealizerRecipeIngredientSerializedData,
    pub base_ingredient: CerealizerRecipeIngredientSerializedData,
    pub addition_ingredient: CerealizerRecipeIngredientSerializedData,
    pub result: CerealizerNetworkItemInstanceDescriptorSerializedData,
    pub tag: String,
    pub net_id: TypedServerNetIdStructRecipeNetIdTag,
}

#[derive(Clone, Debug, PartialEq)]
pub struct SmithingTrimRecipe {
    pub recipe_id: String,
    pub template_ingredient: CerealizerRecipeIngredientSerializedData,
    pub base_ingredient: CerealizerRecipeIngredientSerializedData,
    pub addition_ingredient: CerealizerRecipeIngredientSerializedData,
    pub tag: String,
    pub net_id: TypedServerNetIdStructRecipeNetIdTag,
}

#[derive(Clone, Debug, PartialEq)]
pub struct SocialEventsServerTelemetryData {
    pub server_id: String,
    pub scenario_id: String,
    pub world_id: String,
    pub owner_id: String,
}

#[derive(Clone, Debug, PartialEq)]
pub enum SoundDataEvent {
    Stop,
    SetVolume(SoundDataEventSetVolume),
    SetPitch(SoundDataEventSetPitch),
    Fade(SoundDataEventFade),
    SeekTo(SoundDataEventSeekTo),
    Pause,
    Resume,
}

#[derive(Clone, Debug, PartialEq)]
pub struct SoundDataEventFade {
    pub duration: f32,
    pub target_volume: f32,
}

#[derive(Clone, Debug, PartialEq)]
pub struct SoundDataEventSeekTo {
    pub seconds: f32,
}

#[derive(Clone, Debug, PartialEq)]
pub struct SoundDataEventSetPitch {
    pub pitch: f32,
}

#[derive(Clone, Debug, PartialEq)]
pub struct SoundDataEventSetVolume {
    pub volume: f32,
}

#[derive(Clone, Debug, PartialEq)]
pub struct SpawnSettings {
    pub spawn_biome_type: SpawnBiomeType,
    pub user_defined_biome_name: String,
    pub dimension: i32,
}

#[derive(Clone, Debug, PartialEq)]
pub struct SphereData {
    pub num_segments: u8,
}

#[derive(Clone, Debug, PartialEq)]
pub struct StructureEditorData {
    pub structure_name: BedrockSafetyRedactableString,
    pub data_field: String,
    pub should_include_players: bool,
    pub should_show_bounding_box: bool,
    pub structure_block_type: StructureBlockType,
    pub structure_settings: StructureSettings,
    pub redstone_save_mode: StructureRedstoneSaveMode,
}

#[derive(Clone, Debug, PartialEq)]
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
    pub animation_seconds: f32,
    pub integrity_value: f32,
    pub integrity_seed: u32,
    pub rotation_pivot: glam::Vec3,
}

#[derive(Clone, Debug, PartialEq)]
pub struct SubChunkHeightmapData {
    pub height_map_type: SubChunkHeightMapDataType,
    pub subchunk_height_map: Option<[[i8; 16]; 16]>,
    pub render_height_map_type: SubChunkHeightMapDataType,
    pub subchunk_render_height_map: Option<[[i8; 16]; 16]>,
}

#[derive(Clone, Debug, PartialEq)]
pub struct SubChunkPos {
    pub subchunk_position_x: i32,
    pub subchunk_position_y: i32,
    pub subchunk_position_z: i32,
}

#[derive(Clone, Debug, PartialEq)]
pub struct SubChunkSubChunkPacketData {
    pub sub_chunk_pos_offset: SubChunkSubChunkPosOffset,
    pub sub_chunk_request_result: SubChunkSubChunkRequestResult,
    pub serialized_sub_chunk: Option<String>,
    pub height_map_data: SubChunkHeightmapData,
    pub blob_id: Option<u64>,
}

#[derive(Clone, Debug, PartialEq)]
pub struct SubChunkSubChunkPosOffset {
    pub subchunk_offset_x: i8,
    pub subchunk_offset_y: i8,
    pub subchunk_offset_z: i8,
}

#[derive(Clone, Debug, PartialEq)]
pub struct SyncWorldClockStateData {
    pub clock_id: u64,
    pub time: i32,
    pub is_paused: bool,
}

#[derive(Clone, Debug, PartialEq)]
pub struct SyncWorldClocksAddTimeMarkerData {
    pub clock_id: u64,
    pub time_markers: Vec<TimeMarkerData>,
}

#[derive(Clone, Debug, PartialEq)]
pub enum SyncWorldClocksData {
    SyncStateData(SyncWorldClocksSyncStateData),
    InitializeRegistryData(SyncWorldClocksInitializeRegistryData),
    AddTimeMarkerData(SyncWorldClocksAddTimeMarkerData),
    RemoveTimeMarkerData(SyncWorldClocksRemoveTimeMarkerData),
}

#[derive(Clone, Debug, PartialEq)]
pub struct SyncWorldClocksInitializeRegistryData {
    pub clock_data: Vec<WorldClockData>,
}

#[derive(Clone, Debug, PartialEq)]
pub struct SyncWorldClocksRemoveTimeMarkerData {
    pub clock_id: u64,
    pub time_marker_ids: Vec<u64>,
}

#[derive(Clone, Debug, PartialEq)]
pub struct SyncWorldClocksSyncStateData {
    pub clock_data: Vec<SyncWorldClockStateData>,
}

#[derive(Clone, Debug, PartialEq)]
pub struct SyncedAttribute {
    pub attribute_name: String,
    pub min_value: f32,
    pub current_value: f32,
    pub max_value: f32,
}

#[derive(Clone, Debug, PartialEq)]
pub struct SyncedPlayerMovementSettings {
    pub rewind_history_size: i32,
    pub server_authoritative_block_breaking: bool,
}

#[derive(Clone, Debug, PartialEq)]
pub struct SynchedActorDataCopyableDataList {
    pub data: Vec<DataItemEntry>,
}

#[derive(Clone, Debug, PartialEq)]
pub struct TextAuthorAndMessage {
    pub player_name: String,
    pub message: String,
}

#[derive(Clone, Debug, PartialEq)]
pub enum TextBody {
    Raw(TextMessageOnly),
    Chat(TextAuthorAndMessage),
    Translate(TextMessageAndParams),
    Popup(TextMessageAndParams),
    JukeboxPopup(TextMessageAndParams),
    Tip(TextMessageOnly),
    SystemMessage(TextMessageOnly),
    Whisper(TextAuthorAndMessage),
    Announcement(TextAuthorAndMessage),
    TextObjectWhisper(TextMessageOnly),
    TextObject(TextMessageOnly),
    TextObjectAnnouncement(TextMessageOnly),
}

#[derive(Clone, Debug, PartialEq)]
pub struct TextData {
    pub text: String,
    pub use_rotation: bool,
    pub background_color: Option<MceColor>,
    pub depth_test: bool,
    pub show_backface: bool,
    pub show_text_backface: bool,
}

#[derive(Clone, Debug, PartialEq)]
pub struct TextMessageAndParams {
    pub message: String,
    pub parameter_list: Vec<String>,
}

#[derive(Clone, Debug, PartialEq)]
pub struct TextMessageOnly {
    pub message: String,
}

#[derive(Clone, Debug, PartialEq)]
pub struct TimeMarkerData {
    pub id: u64,
    pub name: String,
    pub time: i32,
    pub period: Option<i32>,
}

#[derive(Clone, Debug, PartialEq)]
pub struct TintMapColor {
    pub colors: [MceColor; 4],
}

#[derive(Clone, Debug, PartialEq)]
pub struct TrimMaterial {
    pub material_id: String,
    pub color: String,
    pub item_name: String,
}

#[derive(Clone, Debug, PartialEq)]
pub struct TrimPattern {
    pub item_name: String,
    pub pattern_id: String,
}

#[derive(Clone, Debug, PartialEq)]
pub struct TypedClientNetIdStructItemStackLegacyRequestIdTagInt32T0 {
    pub id: i32,
}

#[derive(Clone, Debug, PartialEq)]
pub struct TypedClientNetIdStructItemStackRequestIdTagInt32T0 {
    pub id: i32,
}

#[derive(Clone, Debug, PartialEq)]
pub struct TypedServerNetIdStructCreativeItemNetIdTag {
    pub id: u32,
}

#[derive(Clone, Debug, PartialEq)]
pub struct TypedServerNetIdStructItemStackNetIdTagInt32T0 {
    pub id: i32,
}

#[derive(Clone, Debug, PartialEq)]
pub struct TypedServerNetIdStructRecipeNetIdTag {
    pub raw_id: u32,
}

#[derive(Clone, Debug, PartialEq)]
pub struct UpdateSubChunkBlocksChangedInfo {
    pub blocks_changed_standards: Vec<UpdateSubChunkNetworkBlockInfo>,
    pub blocks_changed_extras: Vec<UpdateSubChunkNetworkBlockInfo>,
}

#[derive(Clone, Debug, PartialEq)]
pub struct UpdateSubChunkNetworkBlockInfo {
    pub pos: BlockPos,
    pub runtime_id: u32,
    pub update_flags: u32,
    pub sync_message_entity_unique_id: u64,
    pub sync_message_message: u32,
}

#[derive(Clone, Debug, PartialEq)]
pub struct VoxelShapesRegistryHandle {
    pub value: u16,
}

#[derive(Clone, Debug, PartialEq)]
pub struct VoxelShapesSerializableCells {
    pub x_size: u8,
    pub y_size: u8,
    pub z_size: u8,
    pub storage: Vec<u8>,
}

#[derive(Clone, Debug, PartialEq)]
pub struct VoxelShapesSerializableVoxelShape {
    pub cells: VoxelShapesSerializableCells,
    pub x_coordinates: Vec<f32>,
    pub y_coordinates: Vec<f32>,
    pub z_coordinates: Vec<f32>,
}

#[derive(Clone, Debug, PartialEq)]
pub struct WaypointGroupWaypointHandle {
    pub uuid: uuid::Uuid,
}

#[derive(Clone, Debug, PartialEq)]
pub struct WebSocketPacketData {
    pub websocket_server_uri: String,
}

#[derive(Clone, Debug, PartialEq)]
pub struct WorldClockData {
    pub id: u64,
    pub name: String,
    pub time: i32,
    pub is_paused: bool,
    pub time_markers: Vec<TimeMarkerData>,
}

#[derive(Clone, Debug, PartialEq)]
pub struct WorldPosition {
    pub position: glam::Vec3,
    pub dimension_type: DimensionType,
}
