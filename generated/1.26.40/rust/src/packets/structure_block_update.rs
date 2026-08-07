// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct StructureBlockUpdate {
    pub block_position: BlockPos,
    pub structure_data: StructureEditorData,
    pub trigger: bool,
    pub is_waterlogged: bool,
}

impl StructureBlockUpdate {
    pub const ID: u32 = 90;
}
