// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct SpawnParticleEffect {
    pub dimension_id: u8,
    pub actor_id: ActorUniqueID,
    pub position: glam::Vec3,
    pub effect_name: String,
    pub molang_variables: Option<String>,
}

impl SpawnParticleEffect {
    pub const ID: u32 = 118;
}
