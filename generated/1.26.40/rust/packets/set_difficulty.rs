// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct SetDifficulty {
    pub difficulty: u32,
}

pub const SETDIFFICULTY_DIFFICULTY_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}}"#;

impl SetDifficulty {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("SetDifficultyPacket.Difficulty", SETDIFFICULTY_DIFFICULTY_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("SetDifficultyPacket.Difficulty", SETDIFFICULTY_DIFFICULTY_SHAPE);
    }
}
