// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct UpdateBlockSynced {
    pub block_position: BlockPos,
    pub block_runtime_id: u32,
    pub flags: u32,
    pub layer: u32,
    pub unique_actor_id: u64,
    pub actor_sync_message: u64,
}

pub const UPDATEBLOCKSYNCED_BLOCK_POSITION_SHAPE: &str = r#"{"kind":"struct","semantic":"BlockPos","type_id":"BlockPos","fields":[{"ordinal":0,"name":"X","semantic":"X","encode":{"kind":"primitive","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":1,"name":"Y","semantic":"Y","encode":{"kind":"primitive","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":2,"name":"Z","semantic":"Z","encode":{"kind":"primitive","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}"#;
pub const UPDATEBLOCKSYNCED_BLOCK_RUNTIME_ID_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}}"#;
pub const UPDATEBLOCKSYNCED_FLAGS_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}}"#;
pub const UPDATEBLOCKSYNCED_LAYER_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}}"#;
pub const UPDATEBLOCKSYNCED_UNIQUE_ACTOR_ID_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"var_u64","width":64,"signed":false,"zigzag":false,"endianness":"none"}}"#;
pub const UPDATEBLOCKSYNCED_ACTOR_SYNC_MESSAGE_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"var_u64","width":64,"signed":false,"zigzag":false,"endianness":"none"}}"#;

impl UpdateBlockSynced {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("UpdateBlockSyncedPacket.Block Position", UPDATEBLOCKSYNCED_BLOCK_POSITION_SHAPE);
        encoder.field("UpdateBlockSyncedPacket.Block Runtime ID", UPDATEBLOCKSYNCED_BLOCK_RUNTIME_ID_SHAPE);
        encoder.field("UpdateBlockSyncedPacket.Flags", UPDATEBLOCKSYNCED_FLAGS_SHAPE);
        encoder.field("UpdateBlockSyncedPacket.Layer", UPDATEBLOCKSYNCED_LAYER_SHAPE);
        encoder.field("UpdateBlockSyncedPacket.Unique Actor Id", UPDATEBLOCKSYNCED_UNIQUE_ACTOR_ID_SHAPE);
        encoder.field("UpdateBlockSyncedPacket.Actor Sync Message", UPDATEBLOCKSYNCED_ACTOR_SYNC_MESSAGE_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("UpdateBlockSyncedPacket.Block Position", UPDATEBLOCKSYNCED_BLOCK_POSITION_SHAPE);
        decoder.field("UpdateBlockSyncedPacket.Block Runtime ID", UPDATEBLOCKSYNCED_BLOCK_RUNTIME_ID_SHAPE);
        decoder.field("UpdateBlockSyncedPacket.Flags", UPDATEBLOCKSYNCED_FLAGS_SHAPE);
        decoder.field("UpdateBlockSyncedPacket.Layer", UPDATEBLOCKSYNCED_LAYER_SHAPE);
        decoder.field("UpdateBlockSyncedPacket.Unique Actor Id", UPDATEBLOCKSYNCED_UNIQUE_ACTOR_ID_SHAPE);
        decoder.field("UpdateBlockSyncedPacket.Actor Sync Message", UPDATEBLOCKSYNCED_ACTOR_SYNC_MESSAGE_SHAPE);
    }
}
