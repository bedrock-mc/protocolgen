// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct SimpleEvent {
    pub r#type: SimpleEventSubtype,
}

pub const SIMPLEEVENT_R_TYPE_SHAPE: &str = r#"{"kind":"enum","semantic":"SimpleEventPacketPayload::Subtype","type_id":"enums/SimpleEventPacketPayload::Subtype","primitive":{"code":"u16le","width":16,"signed":false,"zigzag":false,"endianness":"little"},"variants":[{"value":0,"name":"UninitializedSubtype","encode":{"kind":"void"}},{"value":1,"name":"EnableCommands","encode":{"kind":"void"}},{"value":2,"name":"DisableCommands","encode":{"kind":"void"}},{"value":3,"name":"UnlockWorldTemplateSettings","encode":{"kind":"void"}}]}"#;

impl SimpleEvent {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("SimpleEventPacket.Type", SIMPLEEVENT_R_TYPE_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("SimpleEventPacket.Type", SIMPLEEVENT_R_TYPE_SHAPE);
    }
}
