// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct PlayerArmorDamage {
    pub armor_slot_and_damage_pairs: Vec<ArmorSlotAndDamagePair>,
}

pub const PLAYERARMORDAMAGE_ARMOR_SLOT_AND_DAMAGE_PAIRS_SHAPE: &str = r#"{"kind":"array","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"element":{"kind":"struct","semantic":"ArmorSlotAndDamagePair","type_id":"ArmorSlotAndDamagePair","fields":[{"ordinal":0,"name":"Armor Slot","semantic":"Armor Slot","encode":{"kind":"enum","semantic":"SharedTypes::Legacy::ArmorSlot","type_id":"enums/SharedTypes::Legacy::ArmorSlot","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"},"variants":[{"value":0,"name":"Head","encode":{"kind":"void"}},{"value":1,"name":"Torso","encode":{"kind":"void"}},{"value":2,"name":"Legs","encode":{"kind":"void"}},{"value":3,"name":"Feet","encode":{"kind":"void"}},{"value":4,"name":"Body","encode":{"kind":"void"}}]},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":1,"name":"Damage","semantic":"Damage","encode":{"kind":"primitive","primitive":{"code":"i16le","width":16,"signed":true,"zigzag":false,"endianness":"little"}},"symmetry":"symmetric","provenance":{"pins":["endstone","mojang"]}}]}}"#;

impl PlayerArmorDamage {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("PlayerArmorDamagePacket.Armor Slot and Damage Pairs", PLAYERARMORDAMAGE_ARMOR_SLOT_AND_DAMAGE_PAIRS_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("PlayerArmorDamagePacket.Armor Slot and Damage Pairs", PLAYERARMORDAMAGE_ARMOR_SLOT_AND_DAMAGE_PAIRS_SHAPE);
    }
}
