// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct Camera {
    pub camera_id: ActorUniqueID,
    pub target_player_id: ActorUniqueID,
}

impl Camera {
    pub const ID: u32 = 73;
}
