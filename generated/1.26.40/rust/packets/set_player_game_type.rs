// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct SetPlayerGameType {
    pub player_game_type: GameType,
}

pub const SETPLAYERGAMETYPE_PLAYER_GAME_TYPE_SHAPE: &str = r#"{"kind":"enum","semantic":"GameType","type_id":"enums/GameType","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"},"variants":[{"value":-1,"name":"Undefined","encode":{"kind":"void"}},{"value":0,"name":"Survival","encode":{"kind":"void"}},{"value":1,"name":"Creative","encode":{"kind":"void"}},{"value":2,"name":"Adventure","encode":{"kind":"void"}},{"value":5,"name":"Default","encode":{"kind":"void"}},{"value":6,"name":"Spectator","encode":{"kind":"void"}}]}"#;

impl SetPlayerGameType {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("SetPlayerGameTypePacket.Player Game Type", SETPLAYERGAMETYPE_PLAYER_GAME_TYPE_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("SetPlayerGameTypePacket.Player Game Type", SETPLAYERGAMETYPE_PLAYER_GAME_TYPE_SHAPE);
    }
}
