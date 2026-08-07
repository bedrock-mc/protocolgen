// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct PlayerAction {
    pub player_runtime_id: ActorRuntimeID,
    pub action: PlayerActionType,
    pub block_position: BlockPos,
    pub result_pos: BlockPos,
    pub face: i32,
}

impl PlayerAction {
    pub const ID: u32 = 36;
}
