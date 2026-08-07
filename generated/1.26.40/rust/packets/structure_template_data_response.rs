// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct StructureTemplateDataResponse {
    pub structure_name: String,
    pub structure_s_nbt: Vec<u8>,
    pub response_type: StructureTemplateResponseType,
}

pub const STRUCTURETEMPLATEDATARESPONSE_STRUCTURE_NAME_SHAPE: &str = r#"{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"}"#;
pub const STRUCTURETEMPLATEDATARESPONSE_STRUCTURE_S_NBT_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"nbt_le","width":0,"signed":false,"zigzag":false,"endianness":"little"}}"#;
pub const STRUCTURETEMPLATEDATARESPONSE_RESPONSE_TYPE_SHAPE: &str = r#"{"kind":"enum","semantic":"StructureTemplateResponseType","type_id":"enums/StructureTemplateResponseType","primitive":{"code":"u8","width":8,"signed":false,"zigzag":false,"endianness":"none"},"variants":[{"value":0,"name":"None","encode":{"kind":"void"}},{"value":1,"name":"Export","encode":{"kind":"void"}},{"value":2,"name":"Query","encode":{"kind":"void"}}]}"#;

impl StructureTemplateDataResponse {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("StructureTemplateDataResponsePacket.Structure Name", STRUCTURETEMPLATEDATARESPONSE_STRUCTURE_NAME_SHAPE);
        encoder.field("StructureTemplateDataResponsePacket.Structure's NBT", STRUCTURETEMPLATEDATARESPONSE_STRUCTURE_S_NBT_SHAPE);
        encoder.field("StructureTemplateDataResponsePacket.Response Type", STRUCTURETEMPLATEDATARESPONSE_RESPONSE_TYPE_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("StructureTemplateDataResponsePacket.Structure Name", STRUCTURETEMPLATEDATARESPONSE_STRUCTURE_NAME_SHAPE);
        decoder.field("StructureTemplateDataResponsePacket.Structure's NBT", STRUCTURETEMPLATEDATARESPONSE_STRUCTURE_S_NBT_SHAPE);
        decoder.field("StructureTemplateDataResponsePacket.Response Type", STRUCTURETEMPLATEDATARESPONSE_RESPONSE_TYPE_SHAPE);
    }
}
