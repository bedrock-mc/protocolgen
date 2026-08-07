// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct PositionTrackingDBClientRequest {
    pub action: PositionTrackingDBClientRequestAction,
    pub id: PositionTrackingId,
}

pub const POSITIONTRACKINGDBCLIENTREQUEST_ACTION_SHAPE: &str = r#"{"kind":"enum","semantic":"PositionTrackingDBClientRequestPacketPayload::Action","type_id":"enums/PositionTrackingDBClientRequestPacketPayload::Action","primitive":{"code":"u8","width":8,"signed":false,"zigzag":false,"endianness":"none"},"variants":[{"value":0,"name":"Query","encode":{"kind":"void"}}]}"#;
pub const POSITIONTRACKINGDBCLIENTREQUEST_ID_SHAPE: &str = r#"{"kind":"struct","semantic":"PositionTrackingId","type_id":"PositionTrackingId","fields":[{"ordinal":0,"name":"Value","semantic":"Value","encode":{"kind":"primitive","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}"#;

impl PositionTrackingDBClientRequest {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("PositionTrackingDBClientRequestPacket.Action", POSITIONTRACKINGDBCLIENTREQUEST_ACTION_SHAPE);
        encoder.field("PositionTrackingDBClientRequestPacket.Id", POSITIONTRACKINGDBCLIENTREQUEST_ID_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("PositionTrackingDBClientRequestPacket.Action", POSITIONTRACKINGDBCLIENTREQUEST_ACTION_SHAPE);
        decoder.field("PositionTrackingDBClientRequestPacket.Id", POSITIONTRACKINGDBCLIENTREQUEST_ID_SHAPE);
    }
}
