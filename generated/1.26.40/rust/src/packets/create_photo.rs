// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct CreatePhoto {
    pub raw_id: u64,
    pub photo_name: String,
    pub photo_item_name: String,
}

impl CreatePhoto {
    pub const ID: u32 = 171;
}
