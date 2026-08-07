// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct ChangeDimension {
    pub dimension_id: DimensionType,
    pub position: glam::Vec3,
    pub respawn: bool,
    pub loading_screen_id: Option<u32>,
}

impl ChangeDimension {
    pub const ID: u32 = 61;
}
