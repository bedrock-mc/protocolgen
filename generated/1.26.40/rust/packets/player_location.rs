// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct PlayerLocation {
    pub target_actor_id: ActorUniqueID,
    pub location: PlayerLocationLocation,
}

pub const PLAYERLOCATION_TARGET_ACTOR_ID_SHAPE: &str = r#"{"kind":"struct","semantic":"ActorUniqueID","type_id":"ActorUniqueID","fields":[{"ordinal":0,"name":"Actor Unique ID","semantic":"Actor Unique ID","encode":{"kind":"primitive","primitive":{"code":"zigzag_i64","width":64,"signed":true,"zigzag":true,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}"#;
pub const PLAYERLOCATION_LOCATION_SHAPE: &str = r#"{"kind":"union","variants":[{"value":0,"name":"PLAYER_LOCATION_COORDINATES","encode":{"kind":"struct","semantic":"PlayerLocationPacketPayload::CoordinatesLocation","type_id":"PlayerLocationPacketPayload::CoordinatesLocation","fields":[{"ordinal":0,"name":"Packet Type","semantic":"Packet Type","encode":{"kind":"enum","semantic":"PlayerLocationPacketPayload::Type","type_id":"enums/PlayerLocationPacketPayload::Type","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"},"variants":[{"value":0,"name":"PLAYER_LOCATION_COORDINATES","encode":{"kind":"void"}}]},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":1,"name":"Position","semantic":"Position","encode":{"kind":"struct","semantic":"Vec3","type_id":"Vec3","fields":[{"ordinal":0,"name":"X","semantic":"X","encode":{"kind":"primitive","primitive":{"code":"f32le","width":32,"signed":false,"zigzag":false,"endianness":"little"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":1,"name":"Y","semantic":"Y","encode":{"kind":"primitive","primitive":{"code":"f32le","width":32,"signed":false,"zigzag":false,"endianness":"little"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":2,"name":"Z","semantic":"Z","encode":{"kind":"primitive","primitive":{"code":"f32le","width":32,"signed":false,"zigzag":false,"endianness":"little"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}},{"value":1,"name":"PLAYER_LOCATION_HIDE","encode":{"kind":"struct","semantic":"PlayerLocationPacketPayload::HiddenLocation","type_id":"PlayerLocationPacketPayload::HiddenLocation","fields":[{"ordinal":0,"name":"Packet Type","semantic":"Packet Type","encode":{"kind":"enum","semantic":"PlayerLocationPacketPayload::Type","type_id":"enums/PlayerLocationPacketPayload::Type","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"},"variants":[{"value":1,"name":"PLAYER_LOCATION_HIDE","encode":{"kind":"void"}}]},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}}],"control":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}}}"#;

impl PlayerLocation {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("PlayerLocationPacket.Target Actor ID", PLAYERLOCATION_TARGET_ACTOR_ID_SHAPE);
        encoder.field("PlayerLocationPacket.Location", PLAYERLOCATION_LOCATION_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("PlayerLocationPacket.Target Actor ID", PLAYERLOCATION_TARGET_ACTOR_ID_SHAPE);
        decoder.field("PlayerLocationPacket.Location", PLAYERLOCATION_LOCATION_SHAPE);
    }
}
