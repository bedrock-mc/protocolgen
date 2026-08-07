// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct SetActorMotion {
    pub target_runtime_id: ActorRuntimeID,
    pub motion: glam::Vec3,
    pub tick: PlayerInputTick,
}

impl SetActorMotion {
    pub const ID: u32 = 40;
}
