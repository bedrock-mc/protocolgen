// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct CodeBuilderSource {
    pub operation: CodeBuilderStorageQueryOptionsOperation,
    pub category: CodeBuilderStorageQueryOptionsCategory,
    pub code_status: CodeBuilderExecutionStateCodeStatus,
}

impl CodeBuilderSource {
    pub const ID: u32 = 178;
}
