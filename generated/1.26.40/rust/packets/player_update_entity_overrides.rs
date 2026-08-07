// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct PlayerUpdateEntityOverrides {
    pub target_id: ActorUniqueID,
    pub property_index: u32,
    pub update: PlayerUpdateEntityOverridesUpdate,
}

pub const PLAYERUPDATEENTITYOVERRIDES_TARGET_ID_SHAPE: &str = r#"{"kind":"struct","semantic":"ActorUniqueID","type_id":"ActorUniqueID","fields":[{"ordinal":0,"name":"Actor Unique ID","semantic":"Actor Unique ID","encode":{"kind":"primitive","primitive":{"code":"zigzag_i64","width":64,"signed":true,"zigzag":true,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}"#;
pub const PLAYERUPDATEENTITYOVERRIDES_PROPERTY_INDEX_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}}"#;
pub const PLAYERUPDATEENTITYOVERRIDES_UPDATE_SHAPE: &str = r#"{"kind":"union","variants":[{"value":0,"name":"PlayerUpdateEntityOverridesPacketPayload::ClearOverride","encode":{"kind":"struct","semantic":"PlayerUpdateEntityOverridesPacketPayload::ClearOverride","type_id":"PlayerUpdateEntityOverridesPacketPayload::ClearOverride","fields":[{"ordinal":0,"name":"Type","semantic":"Type","encode":{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}},{"value":1,"name":"PlayerUpdateEntityOverridesPacketPayload::RemoveOverride","encode":{"kind":"struct","semantic":"PlayerUpdateEntityOverridesPacketPayload::RemoveOverride","type_id":"PlayerUpdateEntityOverridesPacketPayload::RemoveOverride","fields":[{"ordinal":0,"name":"Type","semantic":"Type","encode":{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}},{"value":2,"name":"PlayerUpdateEntityOverridesPacketPayload::IntOverride","encode":{"kind":"struct","semantic":"PlayerUpdateEntityOverridesPacketPayload::IntOverride","type_id":"PlayerUpdateEntityOverridesPacketPayload::IntOverride","fields":[{"ordinal":0,"name":"Type","semantic":"Type","encode":{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":1,"name":"Value","semantic":"Value","encode":{"kind":"primitive","primitive":{"code":"i32le","width":32,"signed":true,"zigzag":false,"endianness":"little"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}},{"value":3,"name":"PlayerUpdateEntityOverridesPacketPayload::FloatOverride","encode":{"kind":"struct","semantic":"PlayerUpdateEntityOverridesPacketPayload::FloatOverride","type_id":"PlayerUpdateEntityOverridesPacketPayload::FloatOverride","fields":[{"ordinal":0,"name":"Type","semantic":"Type","encode":{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":1,"name":"Value","semantic":"Value","encode":{"kind":"primitive","primitive":{"code":"f32le","width":32,"signed":false,"zigzag":false,"endianness":"little"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}}],"control":{"kind":"primitive","primitive":{"code":"u8","width":8,"signed":false,"zigzag":false,"endianness":"none"}}}"#;

impl PlayerUpdateEntityOverrides {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("PlayerUpdateEntityOverridesPacket.Target ID", PLAYERUPDATEENTITYOVERRIDES_TARGET_ID_SHAPE);
        encoder.field("PlayerUpdateEntityOverridesPacket.Property Index", PLAYERUPDATEENTITYOVERRIDES_PROPERTY_INDEX_SHAPE);
        encoder.field("PlayerUpdateEntityOverridesPacket.Update", PLAYERUPDATEENTITYOVERRIDES_UPDATE_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("PlayerUpdateEntityOverridesPacket.Target ID", PLAYERUPDATEENTITYOVERRIDES_TARGET_ID_SHAPE);
        decoder.field("PlayerUpdateEntityOverridesPacket.Property Index", PLAYERUPDATEENTITYOVERRIDES_PROPERTY_INDEX_SHAPE);
        decoder.field("PlayerUpdateEntityOverridesPacket.Update", PLAYERUPDATEENTITYOVERRIDES_UPDATE_SHAPE);
    }
}
