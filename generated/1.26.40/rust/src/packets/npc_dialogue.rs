// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct NpcDialogue {
    pub npc_id_raw_id: u64,
    pub npc_dialogue_action_type: NpcDialogueNpcDialogueActionType,
    pub dialogue: String,
    pub scene_name: String,
    pub npc_name: String,
    pub action_json: String,
}

impl NpcDialogue {
    pub const ID: u32 = 169;
}
