// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct ServerToClientHandshake {
    pub handshake_web_token: String,
}

pub const SERVERTOCLIENTHANDSHAKE_HANDSHAKE_WEB_TOKEN_SHAPE: &str = r#"{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"}"#;

impl ServerToClientHandshake {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("ServerToClientHandshakePacket.Handshake WebToken", SERVERTOCLIENTHANDSHAKE_HANDSHAKE_WEB_TOKEN_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("ServerToClientHandshakePacket.Handshake WebToken", SERVERTOCLIENTHANDSHAKE_HANDSHAKE_WEB_TOKEN_SHAPE);
    }
}
