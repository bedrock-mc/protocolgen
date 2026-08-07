// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct ContainerClose {
    pub container_id: u8,
    pub container_type: u8,
    pub server_initiated_close: bool,
}

impl ContainerClose {
    pub const ID: u32 = 47;
}
