// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct AddBehaviorTree {
    pub behavior_tree_structure_json: String,
}

pub const ADDBEHAVIORTREE_BEHAVIOR_TREE_STRUCTURE_JSON_SHAPE: &str = r#"{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"}"#;

impl AddBehaviorTree {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("AddBehaviorTreePacket.Behavior Tree Structure (JSON)", ADDBEHAVIORTREE_BEHAVIOR_TREE_STRUCTURE_JSON_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("AddBehaviorTreePacket.Behavior Tree Structure (JSON)", ADDBEHAVIORTREE_BEHAVIOR_TREE_STRUCTURE_JSON_SHAPE);
    }
}
