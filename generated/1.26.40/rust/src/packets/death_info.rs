// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct DeathInfo {
    pub death_cause_attack_name: String,
    pub death_cause_message_list: Vec<String>,
}

impl DeathInfo {
    pub const ID: u32 = 189;
}
