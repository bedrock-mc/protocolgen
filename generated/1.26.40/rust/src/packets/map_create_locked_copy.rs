// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct MapCreateLockedCopy {
    pub original_map_id: ActorUniqueID,
    pub new_map_id: ActorUniqueID,
}

impl MapCreateLockedCopy {
    pub const ID: u32 = 131;
}
