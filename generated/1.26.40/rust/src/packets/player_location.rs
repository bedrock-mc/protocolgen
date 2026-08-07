// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct PlayerLocation {
    pub target_actor_id: ActorUniqueID,
    pub location: PlayerLocationLocation,
}

impl PlayerLocation {
    pub const ID: u32 = 326;
}
