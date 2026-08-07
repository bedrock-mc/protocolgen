// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct ClientboundDataStore {
    pub updates: Vec<BedrockDDUI>,
}

impl ClientboundDataStore {
    pub const ID: u32 = 330;
}
