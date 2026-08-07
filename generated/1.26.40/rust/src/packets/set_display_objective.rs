// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct SetDisplayObjective {
    pub display_slot_name: String,
    pub objective_name: String,
    pub objective_display_name: String,
    pub criteria_name: String,
    pub sort_order: i32,
}

impl SetDisplayObjective {
    pub const ID: u32 = 107;
}
