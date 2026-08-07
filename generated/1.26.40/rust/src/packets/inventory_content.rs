// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct InventoryContent {
    pub container_id: u32,
    pub slots: Vec<CerealizerNetworkItemStackDescriptorSerializedData>,
    pub full_container_name: FullContainerName,
    pub storage_item: CerealizerNetworkItemStackDescriptorSerializedData,
}

impl InventoryContent {
    pub const ID: u32 = 49;
}
