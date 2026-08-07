// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct ClientCacheBlobStatus {
    pub missing_ids: Vec<u64>,
    pub found_ids: Vec<u64>,
}

impl ClientCacheBlobStatus {
    pub const ID: u32 = 135;
}
