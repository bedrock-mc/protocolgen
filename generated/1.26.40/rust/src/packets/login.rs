// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct Login {
    pub client_network_version: i32,
    pub connection_request: Vec<u8>,
}

impl Login {
    pub const ID: u32 = 1;
}
