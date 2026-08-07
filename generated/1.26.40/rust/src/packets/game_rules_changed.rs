// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct GameRulesChanged {
    pub rule_data: GameRulesChangedPacketData,
}

impl GameRulesChanged {
    pub const ID: u32 = 72;
}
