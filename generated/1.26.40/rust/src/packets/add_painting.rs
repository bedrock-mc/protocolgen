// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct AddPainting {
    pub target_actor_id: ActorUniqueID,
    pub target_runtime_id: ActorRuntimeID,
    pub position: glam::Vec3,
    pub direction: i32,
    pub motif: String,
}

impl AddPainting {
    pub const ID: u32 = 22;
}
