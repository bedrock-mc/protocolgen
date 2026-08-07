// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct ResourcePackChunkRequest {
    pub resource_name: String,
    pub chunk: i32,
}

impl ResourcePackChunkRequest {
    pub const ID: u32 = 84;
}
