// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct StopSound {
    pub sound_name: String,
    pub stop_all_sounds: bool,
    pub stop_music_legacy: bool,
}

impl StopSound {
    pub const ID: u32 = 87;
}
