// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct EditorNetwork {
    pub route_to_manager: bool,
    pub raw_variant_name: String,
    pub raw_variant_data: Vec<u8>,
}

impl EditorNetwork {
    pub const ID: u32 = 190;
}
