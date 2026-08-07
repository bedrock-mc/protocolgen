// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct ServerPresenceInfo {
    pub presence_configuration: Option<ServerConfigurationPresenceConfiguration>,
}

impl ServerPresenceInfo {
    pub const ID: u32 = 347;
}
