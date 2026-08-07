// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct GameTestResults {
    pub succeeded: bool,
    pub error: String,
    pub test_name: String,
}

pub const GAMETESTRESULTS_SUCCEEDED_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"bool","width":1,"signed":false,"zigzag":false,"endianness":"none"}}"#;
pub const GAMETESTRESULTS_ERROR_SHAPE: &str = r#"{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"}"#;
pub const GAMETESTRESULTS_TEST_NAME_SHAPE: &str = r#"{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"}"#;

impl GameTestResults {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("GameTestResultsPacket.Succeeded", GAMETESTRESULTS_SUCCEEDED_SHAPE);
        encoder.field("GameTestResultsPacket.Error", GAMETESTRESULTS_ERROR_SHAPE);
        encoder.field("GameTestResultsPacket.TestName", GAMETESTRESULTS_TEST_NAME_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("GameTestResultsPacket.Succeeded", GAMETESTRESULTS_SUCCEEDED_SHAPE);
        decoder.field("GameTestResultsPacket.Error", GAMETESTRESULTS_ERROR_SHAPE);
        decoder.field("GameTestResultsPacket.TestName", GAMETESTRESULTS_TEST_NAME_SHAPE);
    }
}
