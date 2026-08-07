// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct MobEquipment {
    pub target_runtime_id: ActorRuntimeID,
    pub item: CerealizerNetworkItemStackDescriptorSerializedData,
    pub slot: u8,
    pub selected_slot: u8,
    pub container_id: u8,
}

pub const MOBEQUIPMENT_TARGET_RUNTIME_ID_SHAPE: &str = r#"{"kind":"struct","semantic":"ActorRuntimeID","type_id":"ActorRuntimeID","fields":[{"ordinal":0,"name":"Actor Runtime ID","semantic":"Actor Runtime ID","encode":{"kind":"primitive","primitive":{"code":"var_u64","width":64,"signed":false,"zigzag":false,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}"#;
pub const MOBEQUIPMENT_ITEM_SHAPE: &str = r##"{"kind":"struct","semantic":"cerealizer\u003cNetworkItemStackDescriptor\u003e::SerializedData","type_id":"cerealizer\u003cNetworkItemStackDescriptor\u003e::SerializedData","fields":[{"ordinal":0,"name":"Id","semantic":"Id","encode":{"kind":"primitive","primitive":{"code":"i16le","width":16,"signed":true,"zigzag":false,"endianness":"little"}},"symmetry":"symmetric","provenance":{"pins":["endstone","mojang"]}},{"ordinal":1,"name":"Stack size","semantic":"Stack size","encode":{"kind":"primitive","primitive":{"code":"u16le","width":16,"signed":false,"zigzag":false,"endianness":"little"}},"symmetry":"symmetric","provenance":{"pins":["endstone","mojang"]}},{"ordinal":2,"name":"Aux value","semantic":"Aux value","encode":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone","mojang"]}},{"ordinal":3,"name":"Net Id Variant","semantic":"Net Id Variant","type_id":"ItemStackNetIdVariant.json#","encode":{"kind":"optional","value":{"kind":"primitive","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"}}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":4,"name":"Block Runtime Id","semantic":"Block Runtime Id","encode":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone","mojang"]}},{"ordinal":5,"name":"User Data Buffer","semantic":"User Data Buffer","encode":{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"},"symmetry":"symmetric","provenance":{"pins":["endstone","mojang"]}}]}"##;
pub const MOBEQUIPMENT_SLOT_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"u8","width":8,"signed":false,"zigzag":false,"endianness":"none"}}"#;
pub const MOBEQUIPMENT_SELECTED_SLOT_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"u8","width":8,"signed":false,"zigzag":false,"endianness":"none"}}"#;
pub const MOBEQUIPMENT_CONTAINER_ID_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"u8","width":8,"signed":false,"zigzag":false,"endianness":"none"}}"#;

impl MobEquipment {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("MobEquipmentPacket.Target Runtime ID", MOBEQUIPMENT_TARGET_RUNTIME_ID_SHAPE);
        encoder.field("MobEquipmentPacket.Item", MOBEQUIPMENT_ITEM_SHAPE);
        encoder.field("MobEquipmentPacket.Slot", MOBEQUIPMENT_SLOT_SHAPE);
        encoder.field("MobEquipmentPacket.Selected Slot", MOBEQUIPMENT_SELECTED_SLOT_SHAPE);
        encoder.field("MobEquipmentPacket.Container ID", MOBEQUIPMENT_CONTAINER_ID_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("MobEquipmentPacket.Target Runtime ID", MOBEQUIPMENT_TARGET_RUNTIME_ID_SHAPE);
        decoder.field("MobEquipmentPacket.Item", MOBEQUIPMENT_ITEM_SHAPE);
        decoder.field("MobEquipmentPacket.Slot", MOBEQUIPMENT_SLOT_SHAPE);
        decoder.field("MobEquipmentPacket.Selected Slot", MOBEQUIPMENT_SELECTED_SLOT_SHAPE);
        decoder.field("MobEquipmentPacket.Container ID", MOBEQUIPMENT_CONTAINER_ID_SHAPE);
    }
}
