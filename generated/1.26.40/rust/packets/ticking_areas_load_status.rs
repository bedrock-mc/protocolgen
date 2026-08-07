// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct TickingAreasLoadStatus {
    pub waiting_for_preload: bool,
}

pub const TICKINGAREASLOADSTATUS_WAITING_FOR_PRELOAD_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"bool","width":1,"signed":false,"zigzag":false,"endianness":"none"}}"#;

impl TickingAreasLoadStatus {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("TickingAreasLoadStatusPacket.Waiting For Preload", TICKINGAREASLOADSTATUS_WAITING_FOR_PRELOAD_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("TickingAreasLoadStatusPacket.Waiting For Preload", TICKINGAREASLOADSTATUS_WAITING_FOR_PRELOAD_SHAPE);
    }
}
