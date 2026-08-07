// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct ToastRequest {
    pub title: String,
    pub content: String,
}

impl ToastRequest {
    pub const ID: u32 = 186;
}
