// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct ServerboundLoadingScreen {
    pub loading_screen_packet_type: ServerboundLoadingScreenPacketType,
    pub loading_screen_id: Option<u32>,
}

impl ServerboundLoadingScreen {
    pub const ID: u32 = 312;
}
