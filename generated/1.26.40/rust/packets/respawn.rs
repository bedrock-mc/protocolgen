// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct Respawn {
    pub position: glam::Vec3,
    pub state: PlayerRespawnState,
    pub player_runtime_id: ActorRuntimeID,
}

pub const RESPAWN_POSITION_SHAPE: &str = r#"{"kind":"struct","semantic":"Vec3","type_id":"Vec3","fields":[{"ordinal":0,"name":"X","semantic":"X","encode":{"kind":"primitive","primitive":{"code":"f32le","width":32,"signed":false,"zigzag":false,"endianness":"little"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":1,"name":"Y","semantic":"Y","encode":{"kind":"primitive","primitive":{"code":"f32le","width":32,"signed":false,"zigzag":false,"endianness":"little"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":2,"name":"Z","semantic":"Z","encode":{"kind":"primitive","primitive":{"code":"f32le","width":32,"signed":false,"zigzag":false,"endianness":"little"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}"#;
pub const RESPAWN_STATE_SHAPE: &str = r#"{"kind":"enum","semantic":"PlayerRespawnState","type_id":"enums/PlayerRespawnState","primitive":{"code":"u8","width":8,"signed":false,"zigzag":false,"endianness":"none"},"variants":[{"value":0,"name":"SearchingForSpawn","encode":{"kind":"void"}},{"value":1,"name":"ReadyToSpawn","encode":{"kind":"void"}},{"value":2,"name":"ClientReadyToSpawn","encode":{"kind":"void"}}]}"#;
pub const RESPAWN_PLAYER_RUNTIME_ID_SHAPE: &str = r#"{"kind":"struct","semantic":"ActorRuntimeID","type_id":"ActorRuntimeID","fields":[{"ordinal":0,"name":"Actor Runtime ID","semantic":"Actor Runtime ID","encode":{"kind":"primitive","primitive":{"code":"var_u64","width":64,"signed":false,"zigzag":false,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}"#;

impl Respawn {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("RespawnPacket.Position", RESPAWN_POSITION_SHAPE);
        encoder.field("RespawnPacket.State", RESPAWN_STATE_SHAPE);
        encoder.field("RespawnPacket.Player Runtime Id", RESPAWN_PLAYER_RUNTIME_ID_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("RespawnPacket.Position", RESPAWN_POSITION_SHAPE);
        decoder.field("RespawnPacket.State", RESPAWN_STATE_SHAPE);
        decoder.field("RespawnPacket.Player Runtime Id", RESPAWN_PLAYER_RUNTIME_ID_SHAPE);
    }
}
