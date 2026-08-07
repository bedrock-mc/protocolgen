// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct ClientboundDataDrivenUIShowScreen {
    pub screen_id: String,
    pub form_id: u32,
    pub data_instance_id: Option<u32>,
}

impl ClientboundDataDrivenUIShowScreen {
    pub const ID: u32 = 333;
}
