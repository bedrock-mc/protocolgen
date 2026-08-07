// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct ServerSettingsResponse {
    pub form_id: u32,
    pub form_ui_json: String,
}

pub const SERVERSETTINGSRESPONSE_FORM_ID_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}}"#;
pub const SERVERSETTINGSRESPONSE_FORM_UI_JSON_SHAPE: &str = r#"{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"}"#;

impl ServerSettingsResponse {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("ServerSettingsResponsePacket.Form ID", SERVERSETTINGSRESPONSE_FORM_ID_SHAPE);
        encoder.field("ServerSettingsResponsePacket.Form UI JSON", SERVERSETTINGSRESPONSE_FORM_UI_JSON_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("ServerSettingsResponsePacket.Form ID", SERVERSETTINGSRESPONSE_FORM_ID_SHAPE);
        decoder.field("ServerSettingsResponsePacket.Form UI JSON", SERVERSETTINGSRESPONSE_FORM_UI_JSON_SHAPE);
    }
}
