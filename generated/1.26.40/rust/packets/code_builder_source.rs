// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct CodeBuilderSource {
    pub operation: CodeBuilderStorageQueryOptionsOperation,
    pub category: CodeBuilderStorageQueryOptionsCategory,
    pub code_status: CodeBuilderExecutionStateCodeStatus,
}

pub const CODEBUILDERSOURCE_OPERATION_SHAPE: &str = r#"{"kind":"enum","semantic":"CodeBuilderStorageQueryOptions::Operation","type_id":"enums/CodeBuilderStorageQueryOptions::Operation","primitive":{"code":"u8","width":8,"signed":false,"zigzag":false,"endianness":"none"},"variants":[{"value":0,"name":"None","encode":{"kind":"void"}},{"value":1,"name":"Get","encode":{"kind":"void"}},{"value":2,"name":"Set","encode":{"kind":"void"}},{"value":3,"name":"Reset","encode":{"kind":"void"}}]}"#;
pub const CODEBUILDERSOURCE_CATEGORY_SHAPE: &str = r#"{"kind":"enum","semantic":"CodeBuilderStorageQueryOptions::Category","type_id":"enums/CodeBuilderStorageQueryOptions::Category","primitive":{"code":"u8","width":8,"signed":false,"zigzag":false,"endianness":"none"},"variants":[{"value":0,"name":"None","encode":{"kind":"void"}},{"value":1,"name":"CodeStatus","encode":{"kind":"void"}},{"value":2,"name":"Instantiation","encode":{"kind":"void"}}]}"#;
pub const CODEBUILDERSOURCE_CODE_STATUS_SHAPE: &str = r#"{"kind":"enum","semantic":"CodeBuilderExecutionState::CodeStatus","type_id":"enums/CodeBuilderExecutionState::CodeStatus","primitive":{"code":"u8","width":8,"signed":false,"zigzag":false,"endianness":"none"},"variants":[{"value":0,"name":"None","encode":{"kind":"void"}},{"value":1,"name":"NotStarted","encode":{"kind":"void"}},{"value":2,"name":"InProgress","encode":{"kind":"void"}},{"value":3,"name":"Paused","encode":{"kind":"void"}},{"value":4,"name":"Error","encode":{"kind":"void"}},{"value":5,"name":"Succeeded","encode":{"kind":"void"}}]}"#;

impl CodeBuilderSource {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("CodeBuilderSourcePacket.Operation", CODEBUILDERSOURCE_OPERATION_SHAPE);
        encoder.field("CodeBuilderSourcePacket.Category", CODEBUILDERSOURCE_CATEGORY_SHAPE);
        encoder.field("CodeBuilderSourcePacket.CodeStatus", CODEBUILDERSOURCE_CODE_STATUS_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("CodeBuilderSourcePacket.Operation", CODEBUILDERSOURCE_OPERATION_SHAPE);
        decoder.field("CodeBuilderSourcePacket.Category", CODEBUILDERSOURCE_CATEGORY_SHAPE);
        decoder.field("CodeBuilderSourcePacket.CodeStatus", CODEBUILDERSOURCE_CODE_STATUS_SHAPE);
    }
}
