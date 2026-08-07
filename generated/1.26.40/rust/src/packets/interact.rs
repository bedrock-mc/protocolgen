// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct Interact {
    pub action: InteractAction,
    pub target_runtime_id: ActorRuntimeID,
    pub position: Option<glam::Vec3>,
}

impl Interact {
    pub const ID: u32 = 33;
}
