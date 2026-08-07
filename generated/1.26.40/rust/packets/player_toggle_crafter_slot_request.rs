// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct PlayerToggleCrafterSlotRequest {
    pub pos_x: i32,
    pub pos_y: i32,
    pub pos_z: i32,
    pub slot_index: u8,
    pub is_disabled: bool,
}

pub const PLAYERTOGGLECRAFTERSLOTREQUEST_POS_X_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"i32le","width":32,"signed":true,"zigzag":false,"endianness":"little"}}"#;
pub const PLAYERTOGGLECRAFTERSLOTREQUEST_POS_Y_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"i32le","width":32,"signed":true,"zigzag":false,"endianness":"little"}}"#;
pub const PLAYERTOGGLECRAFTERSLOTREQUEST_POS_Z_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"i32le","width":32,"signed":true,"zigzag":false,"endianness":"little"}}"#;
pub const PLAYERTOGGLECRAFTERSLOTREQUEST_SLOT_INDEX_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"u8","width":8,"signed":false,"zigzag":false,"endianness":"none"}}"#;
pub const PLAYERTOGGLECRAFTERSLOTREQUEST_IS_DISABLED_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"bool","width":1,"signed":false,"zigzag":false,"endianness":"none"}}"#;

impl PlayerToggleCrafterSlotRequest {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("PlayerToggleCrafterSlotRequestPacket.Pos X", PLAYERTOGGLECRAFTERSLOTREQUEST_POS_X_SHAPE);
        encoder.field("PlayerToggleCrafterSlotRequestPacket.Pos Y", PLAYERTOGGLECRAFTERSLOTREQUEST_POS_Y_SHAPE);
        encoder.field("PlayerToggleCrafterSlotRequestPacket.Pos Z", PLAYERTOGGLECRAFTERSLOTREQUEST_POS_Z_SHAPE);
        encoder.field("PlayerToggleCrafterSlotRequestPacket.Slot Index", PLAYERTOGGLECRAFTERSLOTREQUEST_SLOT_INDEX_SHAPE);
        encoder.field("PlayerToggleCrafterSlotRequestPacket.Is Disabled", PLAYERTOGGLECRAFTERSLOTREQUEST_IS_DISABLED_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("PlayerToggleCrafterSlotRequestPacket.Pos X", PLAYERTOGGLECRAFTERSLOTREQUEST_POS_X_SHAPE);
        decoder.field("PlayerToggleCrafterSlotRequestPacket.Pos Y", PLAYERTOGGLECRAFTERSLOTREQUEST_POS_Y_SHAPE);
        decoder.field("PlayerToggleCrafterSlotRequestPacket.Pos Z", PLAYERTOGGLECRAFTERSLOTREQUEST_POS_Z_SHAPE);
        decoder.field("PlayerToggleCrafterSlotRequestPacket.Slot Index", PLAYERTOGGLECRAFTERSLOTREQUEST_SLOT_INDEX_SHAPE);
        decoder.field("PlayerToggleCrafterSlotRequestPacket.Is Disabled", PLAYERTOGGLECRAFTERSLOTREQUEST_IS_DISABLED_SHAPE);
    }
}
