// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct SetActorData {
    pub target_runtime_id: ActorRuntimeID,
    pub actor_data: SynchedActorDataCopyableDataList,
    pub synched_properties: PropertySyncData,
    pub tick: PlayerInputTick,
}

impl SetActorData {
    pub const ID: u32 = 39;
}
