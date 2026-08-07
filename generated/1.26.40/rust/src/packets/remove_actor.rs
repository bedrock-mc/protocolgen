// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct RemoveActor {
    pub target_actor_id: ActorUniqueID,
}

impl RemoveActor {
    pub const ID: u32 = 14;
}
