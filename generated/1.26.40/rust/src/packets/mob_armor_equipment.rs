// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct MobArmorEquipment {
    pub target_runtime_id: ActorRuntimeID,
    pub head: CerealizerNetworkItemStackDescriptorSerializedData,
    pub torso: CerealizerNetworkItemStackDescriptorSerializedData,
    pub legs: CerealizerNetworkItemStackDescriptorSerializedData,
    pub feet: CerealizerNetworkItemStackDescriptorSerializedData,
    pub body: CerealizerNetworkItemStackDescriptorSerializedData,
}

impl MobArmorEquipment {
    pub const ID: u32 = 32;
}
