// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct Transfer {
    pub server_address: String,
    pub server_port: u16,
    pub reload_world: bool,
    pub gatherings_configuration: Option<ServerConfigurationGatheringsConfigurationJoinInfo>,
}

impl Transfer {
    pub const ID: u32 = 85;
}
