// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct PlayerStartItemCooldown {
    pub item_category: String,
    pub duration_ticks: i32,
}

pub const PLAYERSTARTITEMCOOLDOWN_ITEM_CATEGORY_SHAPE: &str = r#"{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"}"#;
pub const PLAYERSTARTITEMCOOLDOWN_DURATION_TICKS_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"}}"#;

impl PlayerStartItemCooldown {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("PlayerStartItemCooldownPacket.Item Category", PLAYERSTARTITEMCOOLDOWN_ITEM_CATEGORY_SHAPE);
        encoder.field("PlayerStartItemCooldownPacket.Duration Ticks", PLAYERSTARTITEMCOOLDOWN_DURATION_TICKS_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("PlayerStartItemCooldownPacket.Item Category", PLAYERSTARTITEMCOOLDOWN_ITEM_CATEGORY_SHAPE);
        decoder.field("PlayerStartItemCooldownPacket.Duration Ticks", PLAYERSTARTITEMCOOLDOWN_DURATION_TICKS_SHAPE);
    }
}
