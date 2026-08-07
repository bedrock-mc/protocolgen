// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct CommandRequest {
    pub command: String,
    pub origin: CommandOriginData,
    pub is_internal: bool,
    pub version: String,
}

impl CommandRequest {
    pub const ID: u32 = 77;
}
