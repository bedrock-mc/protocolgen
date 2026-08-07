// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct ClientMovementPredictionSync {
    pub actor_data_flag: ActorDataFlagComponent,
    pub actor_bounding_box: ActorDataBoundingBoxComponent,
    pub movement_attributes: [f32; 9],
    pub actor_unique_id: ActorUniqueID,
    pub actor_flying_state: bool,
}

pub const CLIENTMOVEMENTPREDICTIONSYNC_ACTOR_DATA_FLAG_SHAPE: &str = r#"{"kind":"struct","semantic":"ActorDataFlagComponent","type_id":"ActorDataFlagComponent","fields":[{"ordinal":0,"name":"Actor Flag Bitset Data","semantic":"Actor Flag Bitset Data","encode":{"kind":"bitset","representation":"bitset","length":131},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}"#;
pub const CLIENTMOVEMENTPREDICTIONSYNC_ACTOR_BOUNDING_BOX_SHAPE: &str = r#"{"kind":"struct","semantic":"ActorDataBoundingBoxComponent","type_id":"ActorDataBoundingBoxComponent","fields":[{"ordinal":0,"name":"Actor Data Bounding Box","semantic":"Actor Data Bounding Box","encode":{"kind":"fixed_array","element":{"kind":"primitive","primitive":{"code":"f32le","width":32,"signed":false,"zigzag":false,"endianness":"little"}},"length":3},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}"#;
pub const CLIENTMOVEMENTPREDICTIONSYNC_MOVEMENT_ATTRIBUTES_SHAPE: &str = r##"{"kind":"fixed_array","semantic":"MovementAttributesComponent","type_id":"MovementAttributesComponent.json#","element":{"kind":"primitive","primitive":{"code":"f32le","width":32,"signed":false,"zigzag":false,"endianness":"little"}},"length":9}"##;
pub const CLIENTMOVEMENTPREDICTIONSYNC_ACTOR_UNIQUE_ID_SHAPE: &str = r#"{"kind":"struct","semantic":"ActorUniqueID","type_id":"ActorUniqueID","fields":[{"ordinal":0,"name":"Actor Unique ID","semantic":"Actor Unique ID","encode":{"kind":"primitive","primitive":{"code":"zigzag_i64","width":64,"signed":true,"zigzag":true,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}"#;
pub const CLIENTMOVEMENTPREDICTIONSYNC_ACTOR_FLYING_STATE_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"bool","width":1,"signed":false,"zigzag":false,"endianness":"none"}}"#;

impl ClientMovementPredictionSync {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("ClientMovementPredictionSyncPacket.Actor Data Flag", CLIENTMOVEMENTPREDICTIONSYNC_ACTOR_DATA_FLAG_SHAPE);
        encoder.field("ClientMovementPredictionSyncPacket.Actor Bounding Box", CLIENTMOVEMENTPREDICTIONSYNC_ACTOR_BOUNDING_BOX_SHAPE);
        encoder.field("ClientMovementPredictionSyncPacket.Movement Attributes", CLIENTMOVEMENTPREDICTIONSYNC_MOVEMENT_ATTRIBUTES_SHAPE);
        encoder.field("ClientMovementPredictionSyncPacket.Actor Unique ID", CLIENTMOVEMENTPREDICTIONSYNC_ACTOR_UNIQUE_ID_SHAPE);
        encoder.field("ClientMovementPredictionSyncPacket.Actor Flying State", CLIENTMOVEMENTPREDICTIONSYNC_ACTOR_FLYING_STATE_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("ClientMovementPredictionSyncPacket.Actor Data Flag", CLIENTMOVEMENTPREDICTIONSYNC_ACTOR_DATA_FLAG_SHAPE);
        decoder.field("ClientMovementPredictionSyncPacket.Actor Bounding Box", CLIENTMOVEMENTPREDICTIONSYNC_ACTOR_BOUNDING_BOX_SHAPE);
        decoder.field("ClientMovementPredictionSyncPacket.Movement Attributes", CLIENTMOVEMENTPREDICTIONSYNC_MOVEMENT_ATTRIBUTES_SHAPE);
        decoder.field("ClientMovementPredictionSyncPacket.Actor Unique ID", CLIENTMOVEMENTPREDICTIONSYNC_ACTOR_UNIQUE_ID_SHAPE);
        decoder.field("ClientMovementPredictionSyncPacket.Actor Flying State", CLIENTMOVEMENTPREDICTIONSYNC_ACTOR_FLYING_STATE_SHAPE);
    }
}
