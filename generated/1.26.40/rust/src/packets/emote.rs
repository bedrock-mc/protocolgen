// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct Emote {
    pub actor_runtime_id: ActorRuntimeID,
    pub emote_id: String,
    pub emote_length_ticks: u32,
    pub xuid: String,
    pub platform_id: String,
    pub flags: u8,
}

impl Emote {
    pub const ID: u32 = 138;
}
