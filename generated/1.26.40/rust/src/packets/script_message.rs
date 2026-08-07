// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct ScriptMessage {
    pub message_id: String,
    pub message_value: Vec<u8>,
}

impl ScriptMessage {
    pub const ID: u32 = 177;
}
