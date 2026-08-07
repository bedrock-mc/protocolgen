// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct SetTitle {
    pub title_type: SetTitleTitleType,
    pub title_text: String,
    pub fade_in_time: i32,
    pub stay_time: i32,
    pub fade_out_time: i32,
    pub xuid: String,
    pub platform_online_id: String,
    pub filtered_title_message: String,
}

pub const SETTITLE_TITLE_TYPE_SHAPE: &str = r#"{"kind":"enum","semantic":"SetTitlePacketPayload::TitleType","type_id":"enums/SetTitlePacketPayload::TitleType","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"},"variants":[{"value":0,"name":"Clear","encode":{"kind":"void"}},{"value":1,"name":"Reset","encode":{"kind":"void"}},{"value":2,"name":"Title","encode":{"kind":"void"}},{"value":3,"name":"Subtitle","encode":{"kind":"void"}},{"value":4,"name":"Actionbar","encode":{"kind":"void"}},{"value":5,"name":"Times","encode":{"kind":"void"}},{"value":6,"name":"TitleTextObject","encode":{"kind":"void"}},{"value":7,"name":"SubtitleTextObject","encode":{"kind":"void"}},{"value":8,"name":"ActionbarTextObject","encode":{"kind":"void"}}]}"#;
pub const SETTITLE_TITLE_TEXT_SHAPE: &str = r#"{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"}"#;
pub const SETTITLE_FADE_IN_TIME_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"}}"#;
pub const SETTITLE_STAY_TIME_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"}}"#;
pub const SETTITLE_FADE_OUT_TIME_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"}}"#;
pub const SETTITLE_XUID_SHAPE: &str = r#"{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"}"#;
pub const SETTITLE_PLATFORM_ONLINE_ID_SHAPE: &str = r#"{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"}"#;
pub const SETTITLE_FILTERED_TITLE_MESSAGE_SHAPE: &str = r#"{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"}"#;

impl SetTitle {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("SetTitlePacket.Title Type", SETTITLE_TITLE_TYPE_SHAPE);
        encoder.field("SetTitlePacket.Title Text", SETTITLE_TITLE_TEXT_SHAPE);
        encoder.field("SetTitlePacket.Fade In Time", SETTITLE_FADE_IN_TIME_SHAPE);
        encoder.field("SetTitlePacket.Stay Time", SETTITLE_STAY_TIME_SHAPE);
        encoder.field("SetTitlePacket.Fade Out Time", SETTITLE_FADE_OUT_TIME_SHAPE);
        encoder.field("SetTitlePacket.Xuid", SETTITLE_XUID_SHAPE);
        encoder.field("SetTitlePacket.Platform Online Id", SETTITLE_PLATFORM_ONLINE_ID_SHAPE);
        encoder.field("SetTitlePacket.Filtered Title Message", SETTITLE_FILTERED_TITLE_MESSAGE_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("SetTitlePacket.Title Type", SETTITLE_TITLE_TYPE_SHAPE);
        decoder.field("SetTitlePacket.Title Text", SETTITLE_TITLE_TEXT_SHAPE);
        decoder.field("SetTitlePacket.Fade In Time", SETTITLE_FADE_IN_TIME_SHAPE);
        decoder.field("SetTitlePacket.Stay Time", SETTITLE_STAY_TIME_SHAPE);
        decoder.field("SetTitlePacket.Fade Out Time", SETTITLE_FADE_OUT_TIME_SHAPE);
        decoder.field("SetTitlePacket.Xuid", SETTITLE_XUID_SHAPE);
        decoder.field("SetTitlePacket.Platform Online Id", SETTITLE_PLATFORM_ONLINE_ID_SHAPE);
        decoder.field("SetTitlePacket.Filtered Title Message", SETTITLE_FILTERED_TITLE_MESSAGE_SHAPE);
    }
}
