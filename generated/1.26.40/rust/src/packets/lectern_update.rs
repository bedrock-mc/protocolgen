// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct LecternUpdate {
    pub new_page_to_show: u8,
    pub total_pages: u8,
    pub position_of_lectern_to_update: BlockPos,
}

impl LecternUpdate {
    pub const ID: u32 = 125;
}
