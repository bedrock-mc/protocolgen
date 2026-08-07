// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct ClientCacheStatus {
    pub is_cache_supported: bool,
}

pub const CLIENTCACHESTATUS_IS_CACHE_SUPPORTED_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"bool","width":1,"signed":false,"zigzag":false,"endianness":"none"}}"#;

impl ClientCacheStatus {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("ClientCacheStatusPacket.Is cache supported?", CLIENTCACHESTATUS_IS_CACHE_SUPPORTED_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("ClientCacheStatusPacket.Is cache supported?", CLIENTCACHESTATUS_IS_CACHE_SUPPORTED_SHAPE);
    }
}
