// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct UpdateClientInputLocks {
    pub input_lock_component_data: u32,
}

pub const UPDATECLIENTINPUTLOCKS_INPUT_LOCK_COMPONENT_DATA_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}}"#;

impl UpdateClientInputLocks {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("UpdateClientInputLocksPacket.Input Lock ComponentData", UPDATECLIENTINPUTLOCKS_INPUT_LOCK_COMPONENT_DATA_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("UpdateClientInputLocksPacket.Input Lock ComponentData", UPDATECLIENTINPUTLOCKS_INPUT_LOCK_COMPONENT_DATA_SHAPE);
    }
}
