// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct UpdatePlayerGameType {
    pub player_game_type: GameType,
    pub target_player: ActorUniqueID,
    pub tick: PlayerInputTick,
}

pub const UPDATEPLAYERGAMETYPE_PLAYER_GAME_TYPE_SHAPE: &str = r#"{"kind":"enum","semantic":"GameType","type_id":"enums/GameType","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"},"variants":[{"value":-1,"name":"Undefined","encode":{"kind":"void"}},{"value":0,"name":"Survival","encode":{"kind":"void"}},{"value":1,"name":"Creative","encode":{"kind":"void"}},{"value":2,"name":"Adventure","encode":{"kind":"void"}},{"value":5,"name":"Default","encode":{"kind":"void"}},{"value":6,"name":"Spectator","encode":{"kind":"void"}}]}"#;
pub const UPDATEPLAYERGAMETYPE_TARGET_PLAYER_SHAPE: &str = r#"{"kind":"struct","semantic":"ActorUniqueID","type_id":"ActorUniqueID","fields":[{"ordinal":0,"name":"Actor Unique ID","semantic":"Actor Unique ID","encode":{"kind":"primitive","primitive":{"code":"zigzag_i64","width":64,"signed":true,"zigzag":true,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}"#;
pub const UPDATEPLAYERGAMETYPE_TICK_SHAPE: &str = r#"{"kind":"struct","semantic":"PlayerInputTick","type_id":"PlayerInputTick","fields":[{"ordinal":0,"name":"Input tick","semantic":"Input tick","encode":{"kind":"primitive","primitive":{"code":"var_u64","width":64,"signed":false,"zigzag":false,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}"#;

impl UpdatePlayerGameType {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("UpdatePlayerGameTypePacket.Player Game Type", UPDATEPLAYERGAMETYPE_PLAYER_GAME_TYPE_SHAPE);
        encoder.field("UpdatePlayerGameTypePacket.Target player", UPDATEPLAYERGAMETYPE_TARGET_PLAYER_SHAPE);
        encoder.field("UpdatePlayerGameTypePacket.Tick", UPDATEPLAYERGAMETYPE_TICK_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("UpdatePlayerGameTypePacket.Player Game Type", UPDATEPLAYERGAMETYPE_PLAYER_GAME_TYPE_SHAPE);
        decoder.field("UpdatePlayerGameTypePacket.Target player", UPDATEPLAYERGAMETYPE_TARGET_PLAYER_SHAPE);
        decoder.field("UpdatePlayerGameTypePacket.Tick", UPDATEPLAYERGAMETYPE_TICK_SHAPE);
    }
}
