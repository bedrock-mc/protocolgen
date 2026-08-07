// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct ServerToClientHandshake {
    pub handshake_web_token: String,
}

impl ServerToClientHandshake {
    pub const ID: u32 = 3;
}
