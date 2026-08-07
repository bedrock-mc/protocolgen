// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct NpcRequest {
    pub npc_runtime_id: ActorRuntimeID,
    pub request_type: NpcRequestRequestType,
    pub actions: String,
    pub action_index: u8,
    pub scene_name: String,
}

impl NpcRequest {
    pub const ID: u32 = 98;
}
