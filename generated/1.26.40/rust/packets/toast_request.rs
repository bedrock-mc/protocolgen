// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct ToastRequest {
    pub title: String,
    pub content: String,
}

pub const TOASTREQUEST_TITLE_SHAPE: &str = r#"{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"}"#;
pub const TOASTREQUEST_CONTENT_SHAPE: &str = r#"{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"}"#;

impl ToastRequest {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("ToastRequestPacket.Title", TOASTREQUEST_TITLE_SHAPE);
        encoder.field("ToastRequestPacket.Content", TOASTREQUEST_CONTENT_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("ToastRequestPacket.Title", TOASTREQUEST_TITLE_SHAPE);
        decoder.field("ToastRequestPacket.Content", TOASTREQUEST_CONTENT_SHAPE);
    }
}
