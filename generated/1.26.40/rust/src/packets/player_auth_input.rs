// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct PlayerAuthInput {
    pub player_rotation: glam::Vec2,
    pub position: glam::Vec3,
    pub move_vector: glam::Vec2,
    pub player_head_rotation: f32,
    pub input_data: Option<Vec<PlayerAuthInputInputData>>,
    pub input_mode: InputMode,
    pub play_mode: ClientPlayMode,
    pub new_interaction_model: NewInteractionModel,
    pub interact_rotation: glam::Vec2,
    pub client_tick: PlayerInputTick,
    pub pos_delta: glam::Vec3,
    pub item_use_transaction: Option<Option<PackedItemUseLegacyInventoryTransaction>>,
    pub item_stack_request: Option<Option<ItemStackRequestCerealRequestData>>,
    pub player_block_actions: Option<Option<Vec<PlayerBlockActionData>>>,
    pub vehicle_rotation: Option<Option<glam::Vec2>>,
    pub client_predicted_vehicle: Option<Option<ActorUniqueID>>,
    pub analog_move_vector: glam::Vec2,
    pub camera_orientation: glam::Vec3,
    pub raw_move_vector: glam::Vec2,
}

impl PlayerAuthInput {
    pub const ID: u32 = 144;
}
