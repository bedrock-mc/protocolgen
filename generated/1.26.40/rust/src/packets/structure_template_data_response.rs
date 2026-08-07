// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct StructureTemplateDataResponse {
    pub structure_name: String,
    pub structure_s_nbt: Nbt,
    pub response_type: StructureTemplateResponseType,
}

impl StructureTemplateDataResponse {
    pub const ID: u32 = 133;
}
