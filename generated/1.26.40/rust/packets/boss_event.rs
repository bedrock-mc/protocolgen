// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct BossEvent {
    pub target_actor_id: ActorUniqueID,
    pub player_id: ActorUniqueID,
    pub event_type: BossEventUpdateType,
    pub name: String,
    pub filtered_name: String,
    pub health_percent: f32,
    pub color: BossBarColor,
    pub overlay: BossBarOverlay,
}

pub const BOSSEVENT_TARGET_ACTOR_ID_SHAPE: &str = r#"{"kind":"struct","semantic":"ActorUniqueID","type_id":"ActorUniqueID","fields":[{"ordinal":0,"name":"Actor Unique ID","semantic":"Actor Unique ID","encode":{"kind":"primitive","primitive":{"code":"zigzag_i64","width":64,"signed":true,"zigzag":true,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}"#;
pub const BOSSEVENT_PLAYER_ID_SHAPE: &str = r#"{"kind":"struct","semantic":"ActorUniqueID","type_id":"ActorUniqueID","fields":[{"ordinal":0,"name":"Actor Unique ID","semantic":"Actor Unique ID","encode":{"kind":"primitive","primitive":{"code":"zigzag_i64","width":64,"signed":true,"zigzag":true,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}"#;
pub const BOSSEVENT_EVENT_TYPE_SHAPE: &str = r#"{"kind":"enum","semantic":"BossEventUpdateType","type_id":"enums/BossEventUpdateType","primitive":{"code":"u8","width":8,"signed":false,"zigzag":false,"endianness":"none"},"variants":[{"value":0,"name":"Add","encode":{"kind":"void"}},{"value":1,"name":"PlayerAdded","encode":{"kind":"void"}},{"value":2,"name":"Remove","encode":{"kind":"void"}},{"value":3,"name":"PlayerRemoved","encode":{"kind":"void"}},{"value":4,"name":"Update_Percent","encode":{"kind":"void"}},{"value":5,"name":"Update_Name","encode":{"kind":"void"}},{"value":6,"name":"Update_Properties","encode":{"kind":"void"}},{"value":7,"name":"Update_Style","encode":{"kind":"void"}},{"value":8,"name":"Query","encode":{"kind":"void"}}]}"#;
pub const BOSSEVENT_NAME_SHAPE: &str = r#"{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"}"#;
pub const BOSSEVENT_FILTERED_NAME_SHAPE: &str = r#"{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"}"#;
pub const BOSSEVENT_HEALTH_PERCENT_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"f32le","width":32,"signed":false,"zigzag":false,"endianness":"little"}}"#;
pub const BOSSEVENT_COLOR_SHAPE: &str = r#"{"kind":"enum","semantic":"BossBarColor","type_id":"enums/BossBarColor","primitive":{"code":"u8","width":8,"signed":false,"zigzag":false,"endianness":"none"},"variants":[{"value":0,"name":"PINK","encode":{"kind":"void"}},{"value":1,"name":"BLUE","encode":{"kind":"void"}},{"value":2,"name":"RED","encode":{"kind":"void"}},{"value":3,"name":"GREEN","encode":{"kind":"void"}},{"value":4,"name":"YELLOW","encode":{"kind":"void"}},{"value":5,"name":"PURPLE","encode":{"kind":"void"}},{"value":6,"name":"REBECCA_PURPLE","encode":{"kind":"void"}},{"value":7,"name":"WHITE","encode":{"kind":"void"}}]}"#;
pub const BOSSEVENT_OVERLAY_SHAPE: &str = r#"{"kind":"enum","semantic":"BossBarOverlay","type_id":"enums/BossBarOverlay","primitive":{"code":"u8","width":8,"signed":false,"zigzag":false,"endianness":"none"},"variants":[{"value":0,"name":"PROGRESS","encode":{"kind":"void"}},{"value":1,"name":"NOTCHED_6","encode":{"kind":"void"}},{"value":2,"name":"NOTCHED_10","encode":{"kind":"void"}},{"value":3,"name":"NOTCHED_12","encode":{"kind":"void"}},{"value":4,"name":"NOTCHED_20","encode":{"kind":"void"}}]}"#;

impl BossEvent {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("BossEventPacket.Target Actor ID", BOSSEVENT_TARGET_ACTOR_ID_SHAPE);
        encoder.field("BossEventPacket.Player ID", BOSSEVENT_PLAYER_ID_SHAPE);
        encoder.field("BossEventPacket.Event Type", BOSSEVENT_EVENT_TYPE_SHAPE);
        encoder.field("BossEventPacket.Name", BOSSEVENT_NAME_SHAPE);
        encoder.field("BossEventPacket.FilteredName", BOSSEVENT_FILTERED_NAME_SHAPE);
        encoder.field("BossEventPacket.Health Percent", BOSSEVENT_HEALTH_PERCENT_SHAPE);
        encoder.field("BossEventPacket.Color", BOSSEVENT_COLOR_SHAPE);
        encoder.field("BossEventPacket.Overlay", BOSSEVENT_OVERLAY_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("BossEventPacket.Target Actor ID", BOSSEVENT_TARGET_ACTOR_ID_SHAPE);
        decoder.field("BossEventPacket.Player ID", BOSSEVENT_PLAYER_ID_SHAPE);
        decoder.field("BossEventPacket.Event Type", BOSSEVENT_EVENT_TYPE_SHAPE);
        decoder.field("BossEventPacket.Name", BOSSEVENT_NAME_SHAPE);
        decoder.field("BossEventPacket.FilteredName", BOSSEVENT_FILTERED_NAME_SHAPE);
        decoder.field("BossEventPacket.Health Percent", BOSSEVENT_HEALTH_PERCENT_SHAPE);
        decoder.field("BossEventPacket.Color", BOSSEVENT_COLOR_SHAPE);
        decoder.field("BossEventPacket.Overlay", BOSSEVENT_OVERLAY_SHAPE);
    }
}
