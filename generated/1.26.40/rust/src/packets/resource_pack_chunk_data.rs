// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct ResourcePackChunkData {
    pub resource_name: String,
    pub chunk_id: u32,
    pub byte_offset: u64,
    pub chunk_data: bytes::Bytes,
}

impl ResourcePackChunkData {
    pub const ID: u32 = 83;
}
