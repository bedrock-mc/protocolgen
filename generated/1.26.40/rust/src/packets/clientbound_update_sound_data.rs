// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct ClientboundUpdateSoundData {
    pub server_sound_handle: ServerSoundHandle,
    pub stop: Option<SoundDataEvent>,
    pub set_volume: Option<SoundDataEvent>,
    pub set_pitch: Option<SoundDataEvent>,
    pub fade: Option<SoundDataEvent>,
    pub seek_to: Option<SoundDataEvent>,
    pub pause: Option<SoundDataEvent>,
    pub resume: Option<SoundDataEvent>,
}

impl ClientboundUpdateSoundData {
    pub const ID: u32 = 348;
}
