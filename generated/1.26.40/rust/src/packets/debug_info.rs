// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct DebugInfo {
    pub actor_id: ActorUniqueID,
    pub data: Vec<u8>,
}

impl DebugInfo {
    pub const ID: u32 = 155;
}
