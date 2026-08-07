// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct ClientCacheMissResponse {
    pub missing_blobs: Vec<MissingBlobData>,
}

impl ClientCacheMissResponse {
    pub const ID: u32 = 136;
}
