// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct LabTable {
    pub r#type: LabTableType,
    pub position: BlockPos,
    pub reaction: LabTableReactionType,
}

impl LabTable {
    pub const ID: u32 = 109;
}
