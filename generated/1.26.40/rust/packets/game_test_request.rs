// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct GameTestRequest {
    pub max_tests_per_batch: i32,
    pub repeat_count: i32,
    pub rotation: Rotation,
    pub stop_on_failure: bool,
    pub test_pos: BlockPos,
    pub tests_per_row: i32,
    pub test_name: String,
}

pub const GAMETESTREQUEST_MAX_TESTS_PER_BATCH_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"}}"#;
pub const GAMETESTREQUEST_REPEAT_COUNT_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"}}"#;
pub const GAMETESTREQUEST_ROTATION_SHAPE: &str = r#"{"kind":"enum","semantic":"Rotation","type_id":"enums/Rotation","primitive":{"code":"u8","width":8,"signed":false,"zigzag":false,"endianness":"none"},"variants":[{"value":0,"name":"None","encode":{"kind":"void"}},{"value":1,"name":"Rotate90","encode":{"kind":"void"}},{"value":2,"name":"Rotate180","encode":{"kind":"void"}},{"value":3,"name":"Rotate270","encode":{"kind":"void"}}]}"#;
pub const GAMETESTREQUEST_STOP_ON_FAILURE_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"bool","width":1,"signed":false,"zigzag":false,"endianness":"none"}}"#;
pub const GAMETESTREQUEST_TEST_POS_SHAPE: &str = r#"{"kind":"struct","semantic":"BlockPos","type_id":"BlockPos","fields":[{"ordinal":0,"name":"X","semantic":"X","encode":{"kind":"primitive","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":1,"name":"Y","semantic":"Y","encode":{"kind":"primitive","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":2,"name":"Z","semantic":"Z","encode":{"kind":"primitive","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}"#;
pub const GAMETESTREQUEST_TESTS_PER_ROW_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"}}"#;
pub const GAMETESTREQUEST_TEST_NAME_SHAPE: &str = r#"{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"}"#;

impl GameTestRequest {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("GameTestRequestPacket.MaxTestsPerBatch", GAMETESTREQUEST_MAX_TESTS_PER_BATCH_SHAPE);
        encoder.field("GameTestRequestPacket.RepeatCount", GAMETESTREQUEST_REPEAT_COUNT_SHAPE);
        encoder.field("GameTestRequestPacket.Rotation", GAMETESTREQUEST_ROTATION_SHAPE);
        encoder.field("GameTestRequestPacket.StopOnFailure", GAMETESTREQUEST_STOP_ON_FAILURE_SHAPE);
        encoder.field("GameTestRequestPacket.TestPos", GAMETESTREQUEST_TEST_POS_SHAPE);
        encoder.field("GameTestRequestPacket.TestsPerRow", GAMETESTREQUEST_TESTS_PER_ROW_SHAPE);
        encoder.field("GameTestRequestPacket.TestName", GAMETESTREQUEST_TEST_NAME_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("GameTestRequestPacket.MaxTestsPerBatch", GAMETESTREQUEST_MAX_TESTS_PER_BATCH_SHAPE);
        decoder.field("GameTestRequestPacket.RepeatCount", GAMETESTREQUEST_REPEAT_COUNT_SHAPE);
        decoder.field("GameTestRequestPacket.Rotation", GAMETESTREQUEST_ROTATION_SHAPE);
        decoder.field("GameTestRequestPacket.StopOnFailure", GAMETESTREQUEST_STOP_ON_FAILURE_SHAPE);
        decoder.field("GameTestRequestPacket.TestPos", GAMETESTREQUEST_TEST_POS_SHAPE);
        decoder.field("GameTestRequestPacket.TestsPerRow", GAMETESTREQUEST_TESTS_PER_ROW_SHAPE);
        decoder.field("GameTestRequestPacket.TestName", GAMETESTREQUEST_TEST_NAME_SHAPE);
    }
}
