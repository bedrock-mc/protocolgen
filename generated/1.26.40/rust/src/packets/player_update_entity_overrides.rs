// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct PlayerUpdateEntityOverrides {
    pub target_id: ActorUniqueID,
    pub property_index: u32,
    pub update: PlayerUpdateEntityOverridesUpdate,
}

impl PlayerUpdateEntityOverrides {
    pub const ID: u32 = 325;
}
