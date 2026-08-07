// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct ServerboundPackSettingChange {
    pub pack_id: uuid::Uuid,
    pub pack_setting_name: String,
    pub pack_setting_value: ServerboundPackSettingChangePackSettingValue,
}

pub const SERVERBOUNDPACKSETTINGCHANGE_PACK_ID_SHAPE: &str = r##"{"kind":"primitive","semantic":"mce::UUID","type_id":"mce__UUID.json#","primitive":{"code":"uuid","width":128,"signed":false,"zigzag":false,"endianness":"none"}}"##;
pub const SERVERBOUNDPACKSETTINGCHANGE_PACK_SETTING_NAME_SHAPE: &str = r#"{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"}"#;
pub const SERVERBOUNDPACKSETTINGCHANGE_PACK_SETTING_VALUE_SHAPE: &str = r#"{"kind":"union","variants":[{"value":0,"name":"float","encode":{"kind":"primitive","primitive":{"code":"f32le","width":32,"signed":false,"zigzag":false,"endianness":"little"}}},{"value":1,"name":"bool","encode":{"kind":"primitive","primitive":{"code":"bool","width":1,"signed":false,"zigzag":false,"endianness":"none"}}},{"value":2,"name":"string","encode":{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"}}],"control":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}}}"#;

impl ServerboundPackSettingChange {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("ServerboundPackSettingChangePacket.PackId", SERVERBOUNDPACKSETTINGCHANGE_PACK_ID_SHAPE);
        encoder.field("ServerboundPackSettingChangePacket.PackSettingName", SERVERBOUNDPACKSETTINGCHANGE_PACK_SETTING_NAME_SHAPE);
        encoder.field("ServerboundPackSettingChangePacket.PackSettingValue", SERVERBOUNDPACKSETTINGCHANGE_PACK_SETTING_VALUE_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("ServerboundPackSettingChangePacket.PackId", SERVERBOUNDPACKSETTINGCHANGE_PACK_ID_SHAPE);
        decoder.field("ServerboundPackSettingChangePacket.PackSettingName", SERVERBOUNDPACKSETTINGCHANGE_PACK_SETTING_NAME_SHAPE);
        decoder.field("ServerboundPackSettingChangePacket.PackSettingValue", SERVERBOUNDPACKSETTINGCHANGE_PACK_SETTING_VALUE_SHAPE);
    }
}
