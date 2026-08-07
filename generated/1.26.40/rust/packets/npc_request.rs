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

pub const NPCREQUEST_NPC_RUNTIME_ID_SHAPE: &str = r#"{"kind":"struct","semantic":"ActorRuntimeID","type_id":"ActorRuntimeID","fields":[{"ordinal":0,"name":"Actor Runtime ID","semantic":"Actor Runtime ID","encode":{"kind":"primitive","primitive":{"code":"var_u64","width":64,"signed":false,"zigzag":false,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}"#;
pub const NPCREQUEST_REQUEST_TYPE_SHAPE: &str = r#"{"kind":"enum","semantic":"NpcRequestPacketPayload::RequestType","type_id":"enums/NpcRequestPacketPayload::RequestType","primitive":{"code":"u8","width":8,"signed":false,"zigzag":false,"endianness":"none"},"variants":[{"value":0,"name":"SetActions","encode":{"kind":"void"}},{"value":1,"name":"ExecuteAction","encode":{"kind":"void"}},{"value":2,"name":"ExecuteClosingCommands","encode":{"kind":"void"}},{"value":3,"name":"SetName","encode":{"kind":"void"}},{"value":4,"name":"SetSkin","encode":{"kind":"void"}},{"value":5,"name":"SetInteractText","encode":{"kind":"void"}},{"value":6,"name":"ExecuteOpeningCommands","encode":{"kind":"void"}}]}"#;
pub const NPCREQUEST_ACTIONS_SHAPE: &str = r#"{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"}"#;
pub const NPCREQUEST_ACTION_INDEX_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"u8","width":8,"signed":false,"zigzag":false,"endianness":"none"}}"#;
pub const NPCREQUEST_SCENE_NAME_SHAPE: &str = r#"{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"}"#;

impl NpcRequest {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("NpcRequestPacket.NPC Runtime ID", NPCREQUEST_NPC_RUNTIME_ID_SHAPE);
        encoder.field("NpcRequestPacket.Request Type", NPCREQUEST_REQUEST_TYPE_SHAPE);
        encoder.field("NpcRequestPacket.Actions", NPCREQUEST_ACTIONS_SHAPE);
        encoder.field("NpcRequestPacket.Action Index", NPCREQUEST_ACTION_INDEX_SHAPE);
        encoder.field("NpcRequestPacket.Scene Name", NPCREQUEST_SCENE_NAME_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("NpcRequestPacket.NPC Runtime ID", NPCREQUEST_NPC_RUNTIME_ID_SHAPE);
        decoder.field("NpcRequestPacket.Request Type", NPCREQUEST_REQUEST_TYPE_SHAPE);
        decoder.field("NpcRequestPacket.Actions", NPCREQUEST_ACTIONS_SHAPE);
        decoder.field("NpcRequestPacket.Action Index", NPCREQUEST_ACTION_INDEX_SHAPE);
        decoder.field("NpcRequestPacket.Scene Name", NPCREQUEST_SCENE_NAME_SHAPE);
    }
}
