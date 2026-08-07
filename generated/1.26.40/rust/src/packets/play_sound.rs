// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct PlaySound {
    pub name: String,
    pub position: BlockPos,
    pub volume: f32,
    pub pitch: f32,
    pub loop_count: i32,
    pub server_sound_handle: Option<ServerSoundHandle>,
}

impl PlaySound {
    pub const ID: u32 = 86;
}
