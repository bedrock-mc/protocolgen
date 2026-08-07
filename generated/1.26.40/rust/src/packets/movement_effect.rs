// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct MovementEffect {
    pub target_runtime_id: ActorRuntimeID,
    pub effect_id: MovementEffectType,
    pub effect_duration: i32,
    pub tick: PlayerInputTick,
}

impl MovementEffect {
    pub const ID: u32 = 318;
}
