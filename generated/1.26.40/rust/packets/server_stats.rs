// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct ServerStats {
    pub server_time: f32,
    pub network_time: f32,
}

pub const SERVERSTATS_SERVER_TIME_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"f32le","width":32,"signed":false,"zigzag":false,"endianness":"little"}}"#;
pub const SERVERSTATS_NETWORK_TIME_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"f32le","width":32,"signed":false,"zigzag":false,"endianness":"little"}}"#;

impl ServerStats {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("ServerStatsPacket.ServerTime", SERVERSTATS_SERVER_TIME_SHAPE);
        encoder.field("ServerStatsPacket.NetworkTime", SERVERSTATS_NETWORK_TIME_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("ServerStatsPacket.ServerTime", SERVERSTATS_SERVER_TIME_SHAPE);
        decoder.field("ServerStatsPacket.NetworkTime", SERVERSTATS_NETWORK_TIME_SHAPE);
    }
}
