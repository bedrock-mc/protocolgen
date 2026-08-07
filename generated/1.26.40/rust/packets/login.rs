// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct Login {
    pub client_network_version: i32,
    pub connection_request: String,
}

pub const LOGIN_CLIENT_NETWORK_VERSION_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"i32be","width":32,"signed":true,"zigzag":false,"endianness":"big"}}"#;
pub const LOGIN_CONNECTION_REQUEST_SHAPE: &str = r#"{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"}"#;

impl Login {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("LoginPacket.Client Network Version", LOGIN_CLIENT_NETWORK_VERSION_SHAPE);
        encoder.field("LoginPacket.Connection Request", LOGIN_CONNECTION_REQUEST_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("LoginPacket.Client Network Version", LOGIN_CLIENT_NETWORK_VERSION_SHAPE);
        decoder.field("LoginPacket.Connection Request", LOGIN_CONNECTION_REQUEST_SHAPE);
    }
}
