// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct ServerboundDataDrivenScreenClosed {
    pub form_id: u32,
    pub close_reason: String,
}

impl ServerboundDataDrivenScreenClosed {
    pub const ID: u32 = 343;
}
