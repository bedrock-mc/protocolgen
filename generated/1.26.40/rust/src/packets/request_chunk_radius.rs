// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct RequestChunkRadius {
    pub chunk_radius: i32,
    pub max_chunk_radius: u8,
}

impl RequestChunkRadius {
    pub const ID: u32 = 69;
}
