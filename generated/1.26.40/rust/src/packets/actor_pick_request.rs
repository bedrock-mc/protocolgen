// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct ActorPickRequest {
    pub actor_id: i64,
    pub max_slots: u8,
    pub with_data: bool,
}

impl ActorPickRequest {
    pub const ID: u32 = 35;
}
