// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct AddItemActor {
    pub target_actor_id: ActorUniqueID,
    pub target_runtime_id: ActorRuntimeID,
    pub item: CerealizerNetworkItemStackDescriptorSerializedData,
    pub position: glam::Vec3,
    pub velocity: glam::Vec3,
    pub entity_data: SynchedActorDataCopyableDataList,
    pub is_from_fishing: bool,
}

impl AddItemActor {
    pub const ID: u32 = 15;
}
