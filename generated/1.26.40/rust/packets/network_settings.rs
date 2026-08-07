// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct NetworkSettings {
    pub compression_threshold: u16,
    pub compression_algorithm: PacketCompressionAlgorithm,
    pub client_throttle_enabled: bool,
    pub client_throttle_threshold: u8,
    pub client_throttle_scalar: f32,
}

pub const NETWORKSETTINGS_COMPRESSION_THRESHOLD_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"u16le","width":16,"signed":false,"zigzag":false,"endianness":"little"}}"#;
pub const NETWORKSETTINGS_COMPRESSION_ALGORITHM_SHAPE: &str = r#"{"kind":"enum","semantic":"PacketCompressionAlgorithm","type_id":"enums/PacketCompressionAlgorithm","primitive":{"code":"u16le","width":16,"signed":false,"zigzag":false,"endianness":"little"},"variants":[{"value":0,"name":"ZLib","encode":{"kind":"void"}},{"value":1,"name":"Snappy","encode":{"kind":"void"}},{"value":65535,"name":"None","encode":{"kind":"void"}}]}"#;
pub const NETWORKSETTINGS_CLIENT_THROTTLE_ENABLED_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"bool","width":1,"signed":false,"zigzag":false,"endianness":"none"}}"#;
pub const NETWORKSETTINGS_CLIENT_THROTTLE_THRESHOLD_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"u8","width":8,"signed":false,"zigzag":false,"endianness":"none"}}"#;
pub const NETWORKSETTINGS_CLIENT_THROTTLE_SCALAR_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"f32le","width":32,"signed":false,"zigzag":false,"endianness":"little"}}"#;

impl NetworkSettings {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("NetworkSettingsPacket.Compression Threshold", NETWORKSETTINGS_COMPRESSION_THRESHOLD_SHAPE);
        encoder.field("NetworkSettingsPacket.CompressionAlgorithm", NETWORKSETTINGS_COMPRESSION_ALGORITHM_SHAPE);
        encoder.field("NetworkSettingsPacket.Client Throttle Enabled", NETWORKSETTINGS_CLIENT_THROTTLE_ENABLED_SHAPE);
        encoder.field("NetworkSettingsPacket.Client Throttle Threshold", NETWORKSETTINGS_CLIENT_THROTTLE_THRESHOLD_SHAPE);
        encoder.field("NetworkSettingsPacket.Client Throttle Scalar", NETWORKSETTINGS_CLIENT_THROTTLE_SCALAR_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("NetworkSettingsPacket.Compression Threshold", NETWORKSETTINGS_COMPRESSION_THRESHOLD_SHAPE);
        decoder.field("NetworkSettingsPacket.CompressionAlgorithm", NETWORKSETTINGS_COMPRESSION_ALGORITHM_SHAPE);
        decoder.field("NetworkSettingsPacket.Client Throttle Enabled", NETWORKSETTINGS_CLIENT_THROTTLE_ENABLED_SHAPE);
        decoder.field("NetworkSettingsPacket.Client Throttle Threshold", NETWORKSETTINGS_CLIENT_THROTTLE_THRESHOLD_SHAPE);
        decoder.field("NetworkSettingsPacket.Client Throttle Scalar", NETWORKSETTINGS_CLIENT_THROTTLE_SCALAR_SHAPE);
    }
}
