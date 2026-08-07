// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct CameraShake {
    pub intensity: f32,
    pub seconds: f32,
    pub shake_type: CameraShakeType,
    pub shake_action: CameraShakeAction,
}

impl CameraShake {
    pub const ID: u32 = 159;
}
