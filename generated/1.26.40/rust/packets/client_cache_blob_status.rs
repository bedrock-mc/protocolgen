// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct ClientCacheBlobStatus {
    pub missing_ids: Vec<u64>,
    pub found_ids: Vec<u64>,
}

pub const CLIENTCACHEBLOBSTATUS_MISSING_IDS_SHAPE: &str = r#"{"kind":"array","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"element":{"kind":"primitive","primitive":{"code":"u64le","width":64,"signed":false,"zigzag":false,"endianness":"little"}}}"#;
pub const CLIENTCACHEBLOBSTATUS_FOUND_IDS_SHAPE: &str = r#"{"kind":"array","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"element":{"kind":"primitive","primitive":{"code":"u64le","width":64,"signed":false,"zigzag":false,"endianness":"little"}}}"#;

impl ClientCacheBlobStatus {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("ClientCacheBlobStatusPacket.Missing Ids", CLIENTCACHEBLOBSTATUS_MISSING_IDS_SHAPE);
        encoder.field("ClientCacheBlobStatusPacket.Found Ids", CLIENTCACHEBLOBSTATUS_FOUND_IDS_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("ClientCacheBlobStatusPacket.Missing Ids", CLIENTCACHEBLOBSTATUS_MISSING_IDS_SHAPE);
        decoder.field("ClientCacheBlobStatusPacket.Found Ids", CLIENTCACHEBLOBSTATUS_FOUND_IDS_SHAPE);
    }
}
