// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct ClientCacheMissResponse {
    pub missing_blobs: Vec<MissingBlobData>,
}

pub const CLIENTCACHEMISSRESPONSE_MISSING_BLOBS_SHAPE: &str = r#"{"kind":"array","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"element":{"kind":"struct","semantic":"MissingBlobData","type_id":"MissingBlobData","fields":[{"ordinal":0,"name":"Blob Id","semantic":"Blob Id","encode":{"kind":"primitive","primitive":{"code":"u64le","width":64,"signed":false,"zigzag":false,"endianness":"little"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":1,"name":"Blob Data","semantic":"Blob Data","encode":{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}}"#;

impl ClientCacheMissResponse {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("ClientCacheMissResponsePacket.Missing Blobs", CLIENTCACHEMISSRESPONSE_MISSING_BLOBS_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("ClientCacheMissResponsePacket.Missing Blobs", CLIENTCACHEMISSRESPONSE_MISSING_BLOBS_SHAPE);
    }
}
