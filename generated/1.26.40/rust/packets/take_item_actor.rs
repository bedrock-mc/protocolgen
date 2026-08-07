// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct TakeItemActor {
    pub item_runtime_id: ActorRuntimeID,
    pub actor_runtime_id: ActorRuntimeID,
}

pub const TAKEITEMACTOR_ITEM_RUNTIME_ID_SHAPE: &str = r#"{"kind":"struct","semantic":"ActorRuntimeID","type_id":"ActorRuntimeID","fields":[{"ordinal":0,"name":"Actor Runtime ID","semantic":"Actor Runtime ID","encode":{"kind":"primitive","primitive":{"code":"var_u64","width":64,"signed":false,"zigzag":false,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}"#;
pub const TAKEITEMACTOR_ACTOR_RUNTIME_ID_SHAPE: &str = r#"{"kind":"struct","semantic":"ActorRuntimeID","type_id":"ActorRuntimeID","fields":[{"ordinal":0,"name":"Actor Runtime ID","semantic":"Actor Runtime ID","encode":{"kind":"primitive","primitive":{"code":"var_u64","width":64,"signed":false,"zigzag":false,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}"#;

impl TakeItemActor {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("TakeItemActorPacket.Item Runtime ID", TAKEITEMACTOR_ITEM_RUNTIME_ID_SHAPE);
        encoder.field("TakeItemActorPacket.Actor Runtime ID", TAKEITEMACTOR_ACTOR_RUNTIME_ID_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("TakeItemActorPacket.Item Runtime ID", TAKEITEMACTOR_ITEM_RUNTIME_ID_SHAPE);
        decoder.field("TakeItemActorPacket.Actor Runtime ID", TAKEITEMACTOR_ACTOR_RUNTIME_ID_SHAPE);
    }
}
