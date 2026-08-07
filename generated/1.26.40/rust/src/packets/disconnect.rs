// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct Disconnect {
    pub reason: ConnectionDisconnectFailReason,
    pub messages: DisconnectMessages,
}

impl Disconnect {
    pub const ID: u32 = 5;
}
