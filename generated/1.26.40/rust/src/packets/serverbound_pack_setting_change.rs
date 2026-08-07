// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct ServerboundPackSettingChange {
    pub pack_id: uuid::Uuid,
    pub pack_setting_name: String,
    pub pack_setting_value: ServerboundPackSettingChangePackSettingValue,
}

impl ServerboundPackSettingChange {
    pub const ID: u32 = 329;
}
