// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct SyncActorProperty {
    pub property_data: Nbt,
}

pub const SYNCACTORPROPERTY_PROPERTY_DATA_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"nbt_le","width":0,"signed":false,"zigzag":false,"endianness":"little"}}"#;

impl SyncActorProperty {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("SyncActorPropertyPacket.Property Data", SYNCACTORPROPERTY_PROPERTY_DATA_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("SyncActorPropertyPacket.Property Data", SYNCACTORPROPERTY_PROPERTY_DATA_SHAPE);
    }
}
