// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct MobEquipment {
    pub target_runtime_id: ActorRuntimeID,
    pub item: CerealizerNetworkItemStackDescriptorSerializedData,
    pub slot: u8,
    pub selected_slot: u8,
    pub container_id: u8,
}

impl MobEquipment {
    pub const ID: u32 = 31;
}
