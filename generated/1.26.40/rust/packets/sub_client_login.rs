// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct SubClientLogin {
    pub sub_client_connection_request: String,
}

pub const SUBCLIENTLOGIN_SUB_CLIENT_CONNECTION_REQUEST_SHAPE: &str = r#"{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"}"#;

impl SubClientLogin {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("SubClientLoginPacket.Sub Client Connection Request", SUBCLIENTLOGIN_SUB_CLIENT_CONNECTION_REQUEST_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("SubClientLoginPacket.Sub Client Connection Request", SUBCLIENTLOGIN_SUB_CLIENT_CONNECTION_REQUEST_SHAPE);
    }
}
