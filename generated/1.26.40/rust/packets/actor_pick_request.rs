// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct ActorPickRequest {
    pub actor_id: i64,
    pub max_slots: u8,
    pub with_data: bool,
}

pub const ACTORPICKREQUEST_ACTOR_ID_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"i64le","width":64,"signed":true,"zigzag":false,"endianness":"little"}}"#;
pub const ACTORPICKREQUEST_MAX_SLOTS_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"u8","width":8,"signed":false,"zigzag":false,"endianness":"none"}}"#;
pub const ACTORPICKREQUEST_WITH_DATA_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"bool","width":1,"signed":false,"zigzag":false,"endianness":"none"}}"#;

impl ActorPickRequest {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("ActorPickRequestPacket.Actor ID", ACTORPICKREQUEST_ACTOR_ID_SHAPE);
        encoder.field("ActorPickRequestPacket.Max Slots", ACTORPICKREQUEST_MAX_SLOTS_SHAPE);
        encoder.field("ActorPickRequestPacket.With Data", ACTORPICKREQUEST_WITH_DATA_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("ActorPickRequestPacket.Actor ID", ACTORPICKREQUEST_ACTOR_ID_SHAPE);
        decoder.field("ActorPickRequestPacket.Max Slots", ACTORPICKREQUEST_MAX_SLOTS_SHAPE);
        decoder.field("ActorPickRequestPacket.With Data", ACTORPICKREQUEST_WITH_DATA_SHAPE);
    }
}
