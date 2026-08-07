// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct EditorNetwork {
    pub route_to_manager: bool,
    pub raw_variant_name: String,
    pub raw_variant_data: String,
}

pub const EDITORNETWORK_ROUTE_TO_MANAGER_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"bool","width":1,"signed":false,"zigzag":false,"endianness":"none"}}"#;
pub const EDITORNETWORK_RAW_VARIANT_NAME_SHAPE: &str = r#"{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"}"#;
pub const EDITORNETWORK_RAW_VARIANT_DATA_SHAPE: &str = r#"{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"}"#;

impl EditorNetwork {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("EditorNetworkPacket.Route To Manager", EDITORNETWORK_ROUTE_TO_MANAGER_SHAPE);
        encoder.field("EditorNetworkPacket.Raw Variant Name", EDITORNETWORK_RAW_VARIANT_NAME_SHAPE);
        encoder.field("EditorNetworkPacket.Raw Variant Data", EDITORNETWORK_RAW_VARIANT_DATA_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("EditorNetworkPacket.Route To Manager", EDITORNETWORK_ROUTE_TO_MANAGER_SHAPE);
        decoder.field("EditorNetworkPacket.Raw Variant Name", EDITORNETWORK_RAW_VARIANT_NAME_SHAPE);
        decoder.field("EditorNetworkPacket.Raw Variant Data", EDITORNETWORK_RAW_VARIANT_DATA_SHAPE);
    }
}
