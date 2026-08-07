// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct CommandBlockUpdate {
    pub target: CommandBlockUpdateTarget,
    pub command: String,
    pub last_output: String,
    pub name: String,
    pub filtered_name: String,
    pub track_output: bool,
    pub tick_delay: i32,
    pub execute_on_first_tick: bool,
}

impl CommandBlockUpdate {
    pub const ID: u32 = 78;
}
