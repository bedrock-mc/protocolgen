// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct Animate {
    pub action: AnimateAction,
    pub target_actor_runtime_id: ActorRuntimeID,
    pub data: f32,
    pub swing_source: Option<String>,
}

pub const ANIMATE_ACTION_SHAPE: &str = r#"{"kind":"enum","semantic":"AnimatePacketPayload::Action","type_id":"enums/AnimatePacketPayload::Action","primitive":{"code":"u8","width":8,"signed":false,"zigzag":false,"endianness":"none"},"variants":[{"value":0,"name":"NoAction","encode":{"kind":"void"}},{"value":1,"name":"Swing","encode":{"kind":"void"}},{"value":3,"name":"WakeUp","encode":{"kind":"void"}},{"value":4,"name":"CriticalHit","encode":{"kind":"void"}},{"value":5,"name":"MagicCriticalHit","encode":{"kind":"void"}}]}"#;
pub const ANIMATE_TARGET_ACTOR_RUNTIME_ID_SHAPE: &str = r#"{"kind":"struct","semantic":"ActorRuntimeID","type_id":"ActorRuntimeID","fields":[{"ordinal":0,"name":"Actor Runtime ID","semantic":"Actor Runtime ID","encode":{"kind":"primitive","primitive":{"code":"var_u64","width":64,"signed":false,"zigzag":false,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}"#;
pub const ANIMATE_DATA_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"f32le","width":32,"signed":false,"zigzag":false,"endianness":"little"}}"#;
pub const ANIMATE_SWING_SOURCE_SHAPE: &str = r#"{"kind":"optional","value":{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"}}"#;

impl Animate {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("AnimatePacket.Action", ANIMATE_ACTION_SHAPE);
        encoder.field("AnimatePacket.Target Actor Runtime ID", ANIMATE_TARGET_ACTOR_RUNTIME_ID_SHAPE);
        encoder.field("AnimatePacket.Data", ANIMATE_DATA_SHAPE);
        encoder.field("AnimatePacket.Swing Source", ANIMATE_SWING_SOURCE_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("AnimatePacket.Action", ANIMATE_ACTION_SHAPE);
        decoder.field("AnimatePacket.Target Actor Runtime ID", ANIMATE_TARGET_ACTOR_RUNTIME_ID_SHAPE);
        decoder.field("AnimatePacket.Data", ANIMATE_DATA_SHAPE);
        decoder.field("AnimatePacket.Swing Source", ANIMATE_SWING_SOURCE_SHAPE);
    }
}
