// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct ContainerOpen {
    pub container_id: u8,
    pub container_type: u8,
    pub position: BlockPos,
    pub target_actor_id: ActorUniqueID,
}

impl ContainerOpen {
    pub const ID: u32 = 46;
}
