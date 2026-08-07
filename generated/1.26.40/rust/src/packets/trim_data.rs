// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct TrimData {
    pub trim_pattern_list: Vec<TrimPattern>,
    pub trim_material_list: Vec<TrimMaterial>,
}

impl TrimData {
    pub const ID: u32 = 302;
}
