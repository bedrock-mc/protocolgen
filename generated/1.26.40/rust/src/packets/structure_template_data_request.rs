// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct StructureTemplateDataRequest {
    pub structure_name: String,
    pub structure_position: BlockPos,
    pub structure_settings: StructureSettings,
    pub requested_operation: StructureTemplateRequestOperation,
}

impl StructureTemplateDataRequest {
    pub const ID: u32 = 132;
}
