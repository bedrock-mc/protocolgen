// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct InventorySlot {
    pub container_id: u8,
    pub slot: u32,
    pub full_container_name: Option<FullContainerName>,
    pub storage_item: Option<CerealizerNetworkItemStackDescriptorSerializedData>,
    pub item: CerealizerNetworkItemStackDescriptorSerializedData,
}

impl InventorySlot {
    pub const ID: u32 = 50;
}
