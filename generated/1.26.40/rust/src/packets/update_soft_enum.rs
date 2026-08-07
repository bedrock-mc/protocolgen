// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct UpdateSoftEnum {
    pub enum_name: String,
    pub values: Vec<String>,
    pub update_type: SoftEnumUpdateType,
}

impl UpdateSoftEnum {
    pub const ID: u32 = 114;
}
