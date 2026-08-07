// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct NetworkStackLatency {
    pub creation_time: u64,
    pub is_from_server: bool,
}

pub const NETWORKSTACKLATENCY_CREATION_TIME_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"u64le","width":64,"signed":false,"zigzag":false,"endianness":"little"}}"#;
pub const NETWORKSTACKLATENCY_IS_FROM_SERVER_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"bool","width":1,"signed":false,"zigzag":false,"endianness":"none"}}"#;

impl NetworkStackLatency {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("NetworkStackLatencyPacket.Creation Time", NETWORKSTACKLATENCY_CREATION_TIME_SHAPE);
        encoder.field("NetworkStackLatencyPacket.Is From Server", NETWORKSTACKLATENCY_IS_FROM_SERVER_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("NetworkStackLatencyPacket.Creation Time", NETWORKSTACKLATENCY_CREATION_TIME_SHAPE);
        decoder.field("NetworkStackLatencyPacket.Is From Server", NETWORKSTACKLATENCY_IS_FROM_SERVER_SHAPE);
    }
}
