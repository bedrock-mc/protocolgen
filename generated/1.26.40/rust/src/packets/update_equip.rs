// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct UpdateEquip {
    pub container_id: u8,
    pub r#type: u8,
    pub size: i32,
    pub entity_unique_id: ActorUniqueID,
    pub data: Nbt,
}

impl UpdateEquip {
    pub const ID: u32 = 81;
}
