// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct CodeBuilder {
    pub url: String,
    pub should_open_code_builder: bool,
}

pub const CODEBUILDER_URL_SHAPE: &str = r#"{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"}"#;
pub const CODEBUILDER_SHOULD_OPEN_CODE_BUILDER_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"bool","width":1,"signed":false,"zigzag":false,"endianness":"none"}}"#;

impl CodeBuilder {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("CodeBuilderPacket.URL", CODEBUILDER_URL_SHAPE);
        encoder.field("CodeBuilderPacket.Should open code builder", CODEBUILDER_SHOULD_OPEN_CODE_BUILDER_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("CodeBuilderPacket.URL", CODEBUILDER_URL_SHAPE);
        decoder.field("CodeBuilderPacket.Should open code builder", CODEBUILDER_SHOULD_OPEN_CODE_BUILDER_SHAPE);
    }
}
