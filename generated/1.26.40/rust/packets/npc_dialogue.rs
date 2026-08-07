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

pub const NPCDIALOGUE_NPC_ID_RAW_ID_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"u64le","width":64,"signed":false,"zigzag":false,"endianness":"little"}}"#;
pub const NPCDIALOGUE_NPC_DIALOGUE_ACTION_TYPE_SHAPE: &str = r#"{"kind":"enum","semantic":"NpcDialoguePacketPayload::NpcDialogueActionType","type_id":"enums/NpcDialoguePacketPayload::NpcDialogueActionType","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"},"variants":[{"value":0,"name":"Open","encode":{"kind":"void"}},{"value":1,"name":"Close","encode":{"kind":"void"}}]}"#;
pub const NPCDIALOGUE_DIALOGUE_SHAPE: &str = r#"{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"}"#;
pub const NPCDIALOGUE_SCENE_NAME_SHAPE: &str = r#"{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"}"#;
pub const NPCDIALOGUE_NPC_NAME_SHAPE: &str = r#"{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"}"#;
pub const NPCDIALOGUE_ACTION_JSON_SHAPE: &str = r#"{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"}"#;

impl NpcDialogue {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("NpcDialoguePacket.Npc Id Raw Id", NPCDIALOGUE_NPC_ID_RAW_ID_SHAPE);
        encoder.field("NpcDialoguePacket.Npc Dialogue Action Type", NPCDIALOGUE_NPC_DIALOGUE_ACTION_TYPE_SHAPE);
        encoder.field("NpcDialoguePacket.Dialogue", NPCDIALOGUE_DIALOGUE_SHAPE);
        encoder.field("NpcDialoguePacket.Scene Name", NPCDIALOGUE_SCENE_NAME_SHAPE);
        encoder.field("NpcDialoguePacket.Npc Name", NPCDIALOGUE_NPC_NAME_SHAPE);
        encoder.field("NpcDialoguePacket.Action JSON", NPCDIALOGUE_ACTION_JSON_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("NpcDialoguePacket.Npc Id Raw Id", NPCDIALOGUE_NPC_ID_RAW_ID_SHAPE);
        decoder.field("NpcDialoguePacket.Npc Dialogue Action Type", NPCDIALOGUE_NPC_DIALOGUE_ACTION_TYPE_SHAPE);
        decoder.field("NpcDialoguePacket.Dialogue", NPCDIALOGUE_DIALOGUE_SHAPE);
        decoder.field("NpcDialoguePacket.Scene Name", NPCDIALOGUE_SCENE_NAME_SHAPE);
        decoder.field("NpcDialoguePacket.Npc Name", NPCDIALOGUE_NPC_NAME_SHAPE);
        decoder.field("NpcDialoguePacket.Action JSON", NPCDIALOGUE_ACTION_JSON_SHAPE);
    }
}
