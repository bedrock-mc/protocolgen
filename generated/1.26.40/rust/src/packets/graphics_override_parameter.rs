// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct GraphicsOverrideParameter {
    pub parameter_keyframe_values: Vec<(f32, glam::Vec3)>,
    pub float_value: Option<f32>,
    pub vec3_value: Option<glam::Vec3>,
    pub biome_identifier: String,
    pub player_identifier: Option<String>,
    pub identifier_for_parameter: GraphicsOverrideParameterType,
    pub reset_parameter: bool,
}

impl GraphicsOverrideParameter {
    pub const ID: u32 = 331;
}
