// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct PositionTrackingDBServerBroadcast {
    pub action: PositionTrackingDBServerBroadcastAction,
    pub id: PositionTrackingId,
    pub position_tracking_data: Vec<u8>,
}

pub const POSITIONTRACKINGDBSERVERBROADCAST_ACTION_SHAPE: &str = r#"{"kind":"enum","semantic":"PositionTrackingDBServerBroadcastPacketPayload::Action","type_id":"enums/PositionTrackingDBServerBroadcastPacketPayload::Action","primitive":{"code":"u8","width":8,"signed":false,"zigzag":false,"endianness":"none"},"variants":[{"value":0,"name":"Update","encode":{"kind":"void"}},{"value":1,"name":"Destroy","encode":{"kind":"void"}},{"value":2,"name":"NotFound","encode":{"kind":"void"}}]}"#;
pub const POSITIONTRACKINGDBSERVERBROADCAST_ID_SHAPE: &str = r#"{"kind":"struct","semantic":"PositionTrackingId","type_id":"PositionTrackingId","fields":[{"ordinal":0,"name":"Value","semantic":"Value","encode":{"kind":"primitive","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}"#;
pub const POSITIONTRACKINGDBSERVERBROADCAST_POSITION_TRACKING_DATA_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"nbt_le","width":0,"signed":false,"zigzag":false,"endianness":"little"}}"#;

impl PositionTrackingDBServerBroadcast {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("PositionTrackingDBServerBroadcastPacket.Action", POSITIONTRACKINGDBSERVERBROADCAST_ACTION_SHAPE);
        encoder.field("PositionTrackingDBServerBroadcastPacket.Id", POSITIONTRACKINGDBSERVERBROADCAST_ID_SHAPE);
        encoder.field("PositionTrackingDBServerBroadcastPacket.Position tracking data", POSITIONTRACKINGDBSERVERBROADCAST_POSITION_TRACKING_DATA_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("PositionTrackingDBServerBroadcastPacket.Action", POSITIONTRACKINGDBSERVERBROADCAST_ACTION_SHAPE);
        decoder.field("PositionTrackingDBServerBroadcastPacket.Id", POSITIONTRACKINGDBSERVERBROADCAST_ID_SHAPE);
        decoder.field("PositionTrackingDBServerBroadcastPacket.Position tracking data", POSITIONTRACKINGDBSERVERBROADCAST_POSITION_TRACKING_DATA_SHAPE);
    }
}
