// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct MobEffect {
    pub target_runtime_id: ActorRuntimeID,
    pub event_id: MobEffectEvent,
    pub effect_id: i32,
    pub effect_amplifier: i32,
    pub show_particles: bool,
    pub effect_duration_ticks: i32,
    pub tick: PlayerInputTick,
    pub ambient: bool,
}

impl MobEffect {
    pub const ID: u32 = 28;
}
