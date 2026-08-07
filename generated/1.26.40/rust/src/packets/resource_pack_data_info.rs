// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct ResourcePackDataInfo {
    pub resource_name: String,
    pub chunk_size: u32,
    pub number_of_chunks: u32,
    pub file_size: u64,
    pub file_hash: String,
    pub is_premium_pack: bool,
    pub pack_type: u8,
}

impl ResourcePackDataInfo {
    pub const ID: u32 = 82;
}
