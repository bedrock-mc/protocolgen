// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct EduUriResource {
    pub edu_shared_uri_resource: EduSharedUriResource,
}

pub const EDUURIRESOURCE_EDU_SHARED_URI_RESOURCE_SHAPE: &str = r#"{"kind":"struct","semantic":"EduSharedUriResource","type_id":"EduSharedUriResource","fields":[{"ordinal":0,"name":"Button Name","semantic":"Button Name","encode":{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":1,"name":"Link Uri","semantic":"Link Uri","encode":{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}"#;

impl EduUriResource {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("EduUriResourcePacket.Edu Shared URI Resource", EDUURIRESOURCE_EDU_SHARED_URI_RESOURCE_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("EduUriResourcePacket.Edu Shared URI Resource", EDUURIRESOURCE_EDU_SHARED_URI_RESOURCE_SHAPE);
    }
}
