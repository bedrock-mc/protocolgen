// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct UpdatePlayerGameType {
    pub player_game_type: GameType,
    pub target_player: ActorUniqueID,
    pub tick: PlayerInputTick,
}

impl UpdatePlayerGameType {
    pub const ID: u32 = 151;
}
