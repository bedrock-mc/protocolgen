// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct ServerboundDataDrivenScreenClosed {
    pub form_id: u32,
    pub close_reason: String,
}

pub const SERVERBOUNDDATADRIVENSCREENCLOSED_FORM_ID_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"u32le","width":32,"signed":false,"zigzag":false,"endianness":"little"}}"#;
pub const SERVERBOUNDDATADRIVENSCREENCLOSED_CLOSE_REASON_SHAPE: &str = r#"{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"}"#;

impl ServerboundDataDrivenScreenClosed {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("ServerboundDataDrivenScreenClosedPacket.FormId", SERVERBOUNDDATADRIVENSCREENCLOSED_FORM_ID_SHAPE);
        encoder.field("ServerboundDataDrivenScreenClosedPacket.CloseReason", SERVERBOUNDDATADRIVENSCREENCLOSED_CLOSE_REASON_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("ServerboundDataDrivenScreenClosedPacket.FormId", SERVERBOUNDDATADRIVENSCREENCLOSED_FORM_ID_SHAPE);
        decoder.field("ServerboundDataDrivenScreenClosedPacket.CloseReason", SERVERBOUNDDATADRIVENSCREENCLOSED_CLOSE_REASON_SHAPE);
    }
}
