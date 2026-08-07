// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct RemoveObjective {
    pub objective_name: String,
}

pub const REMOVEOBJECTIVE_OBJECTIVE_NAME_SHAPE: &str = r#"{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"}"#;

impl RemoveObjective {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("RemoveObjectivePacket.Objective Name", REMOVEOBJECTIVE_OBJECTIVE_NAME_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("RemoveObjectivePacket.Objective Name", REMOVEOBJECTIVE_OBJECTIVE_NAME_SHAPE);
    }
}
