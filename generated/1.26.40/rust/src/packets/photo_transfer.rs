// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct PhotoTransfer {
    pub photo_name: String,
    pub photo_data: Vec<u8>,
    pub book_id: String,
    pub r#type: PhotoType,
    pub source_type: PhotoType,
    pub owner_id: i64,
    pub new_photo_name: String,
}

impl PhotoTransfer {
    pub const ID: u32 = 99;
}
