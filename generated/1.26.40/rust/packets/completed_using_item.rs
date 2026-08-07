// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct CompletedUsingItem {
    pub item_id: i16,
    pub item_use_method: i32,
}

pub const COMPLETEDUSINGITEM_ITEM_ID_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"i16le","width":16,"signed":true,"zigzag":false,"endianness":"little"}}"#;
pub const COMPLETEDUSINGITEM_ITEM_USE_METHOD_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"i32le","width":32,"signed":true,"zigzag":false,"endianness":"little"}}"#;

impl CompletedUsingItem {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("CompletedUsingItemPacket.Item Id", COMPLETEDUSINGITEM_ITEM_ID_SHAPE);
        encoder.field("CompletedUsingItemPacket.Item Use Method", COMPLETEDUSINGITEM_ITEM_USE_METHOD_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("CompletedUsingItemPacket.Item Id", COMPLETEDUSINGITEM_ITEM_ID_SHAPE);
        decoder.field("CompletedUsingItemPacket.Item Use Method", COMPLETEDUSINGITEM_ITEM_USE_METHOD_SHAPE);
    }
}
