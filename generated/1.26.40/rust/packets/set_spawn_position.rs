// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct SetSpawnPosition {
    pub spawn_position_type: SpawnPositionType,
    pub block_position: BlockPos,
    pub dimension_type: DimensionType,
    pub spawn_block_pos: BlockPos,
}

pub const SETSPAWNPOSITION_SPAWN_POSITION_TYPE_SHAPE: &str = r#"{"kind":"enum","semantic":"SpawnPositionType","type_id":"enums/SpawnPositionType","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"},"variants":[{"value":0,"name":"PlayerRespawn","encode":{"kind":"void"}},{"value":1,"name":"WorldSpawn","encode":{"kind":"void"}}]}"#;
pub const SETSPAWNPOSITION_BLOCK_POSITION_SHAPE: &str = r#"{"kind":"struct","semantic":"BlockPos","type_id":"BlockPos","fields":[{"ordinal":0,"name":"X","semantic":"X","encode":{"kind":"primitive","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":1,"name":"Y","semantic":"Y","encode":{"kind":"primitive","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":2,"name":"Z","semantic":"Z","encode":{"kind":"primitive","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}"#;
pub const SETSPAWNPOSITION_DIMENSION_TYPE_SHAPE: &str = r#"{"kind":"struct","semantic":"DimensionType","type_id":"DimensionType","fields":[{"ordinal":0,"name":"value","semantic":"value","encode":{"kind":"primitive","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}"#;
pub const SETSPAWNPOSITION_SPAWN_BLOCK_POS_SHAPE: &str = r#"{"kind":"struct","semantic":"BlockPos","type_id":"BlockPos","fields":[{"ordinal":0,"name":"X","semantic":"X","encode":{"kind":"primitive","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":1,"name":"Y","semantic":"Y","encode":{"kind":"primitive","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":2,"name":"Z","semantic":"Z","encode":{"kind":"primitive","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}"#;

impl SetSpawnPosition {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("SetSpawnPositionPacket.Spawn Position Type", SETSPAWNPOSITION_SPAWN_POSITION_TYPE_SHAPE);
        encoder.field("SetSpawnPositionPacket.Block Position", SETSPAWNPOSITION_BLOCK_POSITION_SHAPE);
        encoder.field("SetSpawnPositionPacket.Dimension type", SETSPAWNPOSITION_DIMENSION_TYPE_SHAPE);
        encoder.field("SetSpawnPositionPacket.Spawn Block Pos", SETSPAWNPOSITION_SPAWN_BLOCK_POS_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("SetSpawnPositionPacket.Spawn Position Type", SETSPAWNPOSITION_SPAWN_POSITION_TYPE_SHAPE);
        decoder.field("SetSpawnPositionPacket.Block Position", SETSPAWNPOSITION_BLOCK_POSITION_SHAPE);
        decoder.field("SetSpawnPositionPacket.Dimension type", SETSPAWNPOSITION_DIMENSION_TYPE_SHAPE);
        decoder.field("SetSpawnPositionPacket.Spawn Block Pos", SETSPAWNPOSITION_SPAWN_BLOCK_POS_SHAPE);
    }
}
