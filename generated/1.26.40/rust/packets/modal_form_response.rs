// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct ModalFormResponse {
    pub form_id: u32,
    pub json_response: Option<String>,
    pub form_cancel_reason: Option<ModalFormCancelReason>,
}

pub const MODALFORMRESPONSE_FORM_ID_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}}"#;
pub const MODALFORMRESPONSE_JSON_RESPONSE_SHAPE: &str = r##"{"kind":"optional","value":{"kind":"string","semantic":"Json::Value","type_id":"Json__Value.json#","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"}}"##;
pub const MODALFORMRESPONSE_FORM_CANCEL_REASON_SHAPE: &str = r#"{"kind":"optional","value":{"kind":"enum","semantic":"ModalFormCancelReason","type_id":"enums/ModalFormCancelReason","primitive":{"code":"u8","width":8,"signed":false,"zigzag":false,"endianness":"none"},"variants":[{"value":0,"name":"UserClosed","encode":{"kind":"void"}},{"value":1,"name":"UserBusy","encode":{"kind":"void"}}]}}"#;

impl ModalFormResponse {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("ModalFormResponsePacket.Form ID", MODALFORMRESPONSE_FORM_ID_SHAPE);
        encoder.field("ModalFormResponsePacket.JSON Response", MODALFORMRESPONSE_JSON_RESPONSE_SHAPE);
        encoder.field("ModalFormResponsePacket.Form Cancel Reason", MODALFORMRESPONSE_FORM_CANCEL_REASON_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("ModalFormResponsePacket.Form ID", MODALFORMRESPONSE_FORM_ID_SHAPE);
        decoder.field("ModalFormResponsePacket.JSON Response", MODALFORMRESPONSE_JSON_RESPONSE_SHAPE);
        decoder.field("ModalFormResponsePacket.Form Cancel Reason", MODALFORMRESPONSE_FORM_CANCEL_REASON_SHAPE);
    }
}
