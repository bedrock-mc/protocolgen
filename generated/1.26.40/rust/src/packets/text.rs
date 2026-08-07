// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct Text {
    pub localize: bool,
    pub body: TextBody,
    pub sender_s_xuid: String,
    pub platform_id: String,
    pub filtered_message: Option<String>,
}

impl Text {
    pub const ID: u32 = 9;
}
