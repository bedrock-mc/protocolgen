// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct AvailableActorIdentifiers {
    pub identifier_list: Vec<u8>,
}

pub const AVAILABLEACTORIDENTIFIERS_IDENTIFIER_LIST_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"nbt_le","width":0,"signed":false,"zigzag":false,"endianness":"little"}}"#;

impl AvailableActorIdentifiers {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("AvailableActorIdentifiersPacket.Identifier List", AVAILABLEACTORIDENTIFIERS_IDENTIFIER_LIST_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("AvailableActorIdentifiersPacket.Identifier List", AVAILABLEACTORIDENTIFIERS_IDENTIFIER_LIST_SHAPE);
    }
}
