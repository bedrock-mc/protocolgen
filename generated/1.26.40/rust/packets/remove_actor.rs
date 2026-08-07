// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct RemoveActor {
    pub target_actor_id: ActorUniqueID,
}

pub const REMOVEACTOR_TARGET_ACTOR_ID_SHAPE: &str = r#"{"kind":"struct","semantic":"ActorUniqueID","type_id":"ActorUniqueID","fields":[{"ordinal":0,"name":"Actor Unique ID","semantic":"Actor Unique ID","encode":{"kind":"primitive","primitive":{"code":"zigzag_i64","width":64,"signed":true,"zigzag":true,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}"#;

impl RemoveActor {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("RemoveActorPacket.Target Actor ID", REMOVEACTOR_TARGET_ACTOR_ID_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("RemoveActorPacket.Target Actor ID", REMOVEACTOR_TARGET_ACTOR_ID_SHAPE);
    }
}
