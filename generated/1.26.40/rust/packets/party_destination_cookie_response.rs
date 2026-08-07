// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct PartyDestinationCookieResponse {
    pub cookie: String,
    pub accepted: bool,
}

pub const PARTYDESTINATIONCOOKIERESPONSE_COOKIE_SHAPE: &str = r#"{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"}"#;
pub const PARTYDESTINATIONCOOKIERESPONSE_ACCEPTED_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"bool","width":1,"signed":false,"zigzag":false,"endianness":"none"}}"#;

impl PartyDestinationCookieResponse {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("PartyDestinationCookieResponsePacket.cookie", PARTYDESTINATIONCOOKIERESPONSE_COOKIE_SHAPE);
        encoder.field("PartyDestinationCookieResponsePacket.accepted", PARTYDESTINATIONCOOKIERESPONSE_ACCEPTED_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("PartyDestinationCookieResponsePacket.cookie", PARTYDESTINATIONCOOKIERESPONSE_COOKIE_SHAPE);
        decoder.field("PartyDestinationCookieResponsePacket.accepted", PARTYDESTINATIONCOOKIERESPONSE_ACCEPTED_SHAPE);
    }
}
