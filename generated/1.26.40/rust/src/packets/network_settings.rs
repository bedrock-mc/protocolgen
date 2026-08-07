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

impl NetworkSettings {
    pub const ID: u32 = 143;
}
