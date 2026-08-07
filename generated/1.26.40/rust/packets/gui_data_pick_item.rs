// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct GuiDataPickItem {
    pub item_name: String,
    pub item_effect_name: String,
    pub slot: i32,
}

pub const GUIDATAPICKITEM_ITEM_NAME_SHAPE: &str = r#"{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"}"#;
pub const GUIDATAPICKITEM_ITEM_EFFECT_NAME_SHAPE: &str = r#"{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"}"#;
pub const GUIDATAPICKITEM_SLOT_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"i32le","width":32,"signed":true,"zigzag":false,"endianness":"little"}}"#;

impl GuiDataPickItem {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("GuiDataPickItemPacket.Item Name", GUIDATAPICKITEM_ITEM_NAME_SHAPE);
        encoder.field("GuiDataPickItemPacket.Item Effect Name", GUIDATAPICKITEM_ITEM_EFFECT_NAME_SHAPE);
        encoder.field("GuiDataPickItemPacket.Slot", GUIDATAPICKITEM_SLOT_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("GuiDataPickItemPacket.Item Name", GUIDATAPICKITEM_ITEM_NAME_SHAPE);
        decoder.field("GuiDataPickItemPacket.Item Effect Name", GUIDATAPICKITEM_ITEM_EFFECT_NAME_SHAPE);
        decoder.field("GuiDataPickItemPacket.Slot", GUIDATAPICKITEM_SLOT_SHAPE);
    }
}
