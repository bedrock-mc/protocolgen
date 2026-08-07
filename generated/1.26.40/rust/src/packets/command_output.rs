// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct CommandOutput {
    pub origin_data: CommandOriginData,
    pub output: CommandOutputData,
}

impl CommandOutput {
    pub const ID: u32 = 79;
}
