// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct SendPartyDestinationCookie {
    pub cookie: String,
    pub intent: String,
    pub destination_name: String,
}

pub const SENDPARTYDESTINATIONCOOKIE_COOKIE_SHAPE: &str = r#"{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"}"#;
pub const SENDPARTYDESTINATIONCOOKIE_INTENT_SHAPE: &str = r#"{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"}"#;
pub const SENDPARTYDESTINATIONCOOKIE_DESTINATION_NAME_SHAPE: &str = r#"{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"}"#;

impl SendPartyDestinationCookie {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("SendPartyDestinationCookiePacket.cookie", SENDPARTYDESTINATIONCOOKIE_COOKIE_SHAPE);
        encoder.field("SendPartyDestinationCookiePacket.intent", SENDPARTYDESTINATIONCOOKIE_INTENT_SHAPE);
        encoder.field("SendPartyDestinationCookiePacket.destination_name", SENDPARTYDESTINATIONCOOKIE_DESTINATION_NAME_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("SendPartyDestinationCookiePacket.cookie", SENDPARTYDESTINATIONCOOKIE_COOKIE_SHAPE);
        decoder.field("SendPartyDestinationCookiePacket.intent", SENDPARTYDESTINATIONCOOKIE_INTENT_SHAPE);
        decoder.field("SendPartyDestinationCookiePacket.destination_name", SENDPARTYDESTINATIONCOOKIE_DESTINATION_NAME_SHAPE);
    }
}
