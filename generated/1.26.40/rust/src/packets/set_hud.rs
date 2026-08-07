// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct SetHud {
    pub hud_element: Vec<HudElement>,
    pub hud_visible: HudVisibility,
}

impl SetHud {
    pub const ID: u32 = 308;
}
