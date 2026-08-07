// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct PlayerHotbar {
    pub selected_slot: u32,
    pub container_id: u8,
    pub should_select_slot: bool,
}

pub const PLAYERHOTBAR_SELECTED_SLOT_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}}"#;
pub const PLAYERHOTBAR_CONTAINER_ID_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"u8","width":8,"signed":false,"zigzag":false,"endianness":"none"}}"#;
pub const PLAYERHOTBAR_SHOULD_SELECT_SLOT_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"bool","width":1,"signed":false,"zigzag":false,"endianness":"none"}}"#;

impl PlayerHotbar {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("PlayerHotbarPacket.Selected Slot", PLAYERHOTBAR_SELECTED_SLOT_SHAPE);
        encoder.field("PlayerHotbarPacket.Container ID", PLAYERHOTBAR_CONTAINER_ID_SHAPE);
        encoder.field("PlayerHotbarPacket.Should select slot?", PLAYERHOTBAR_SHOULD_SELECT_SLOT_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("PlayerHotbarPacket.Selected Slot", PLAYERHOTBAR_SELECTED_SLOT_SHAPE);
        decoder.field("PlayerHotbarPacket.Container ID", PLAYERHOTBAR_CONTAINER_ID_SHAPE);
        decoder.field("PlayerHotbarPacket.Should select slot?", PLAYERHOTBAR_SHOULD_SELECT_SLOT_SHAPE);
    }
}
