// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct ShowCredits {
    pub player_runtime_id: ActorRuntimeID,
    pub credits_state: i32,
}

impl ShowCredits {
    pub const ID: u32 = 75;
}
