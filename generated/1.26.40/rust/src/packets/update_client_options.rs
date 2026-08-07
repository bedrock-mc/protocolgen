// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct UpdateClientOptions {
    pub graphics_mode_change: Option<GraphicsMode>,
    pub filter_profanity_change: Option<bool>,
}

impl UpdateClientOptions {
    pub const ID: u32 = 323;
}
