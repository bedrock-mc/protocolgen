// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct HurtArmor {
    pub cause: i32,
    pub damage: i32,
    pub armor_slots: u64,
}

pub const HURTARMOR_CAUSE_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"}}"#;
pub const HURTARMOR_DAMAGE_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"}}"#;
pub const HURTARMOR_ARMOR_SLOTS_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"var_u64","width":64,"signed":false,"zigzag":false,"endianness":"none"}}"#;

impl HurtArmor {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("HurtArmorPacket.Cause", HURTARMOR_CAUSE_SHAPE);
        encoder.field("HurtArmorPacket.Damage", HURTARMOR_DAMAGE_SHAPE);
        encoder.field("HurtArmorPacket.Armor Slots", HURTARMOR_ARMOR_SLOTS_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("HurtArmorPacket.Cause", HURTARMOR_CAUSE_SHAPE);
        decoder.field("HurtArmorPacket.Damage", HURTARMOR_DAMAGE_SHAPE);
        decoder.field("HurtArmorPacket.Armor Slots", HURTARMOR_ARMOR_SLOTS_SHAPE);
    }
}
