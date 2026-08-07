// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct ContainerClose {
    pub container_id: u8,
    pub container_type: u8,
    pub server_initiated_close: bool,
}

pub const CONTAINERCLOSE_CONTAINER_ID_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"u8","width":8,"signed":false,"zigzag":false,"endianness":"none"}}"#;
pub const CONTAINERCLOSE_CONTAINER_TYPE_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"u8","width":8,"signed":false,"zigzag":false,"endianness":"none"}}"#;
pub const CONTAINERCLOSE_SERVER_INITIATED_CLOSE_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"bool","width":1,"signed":false,"zigzag":false,"endianness":"none"}}"#;

impl ContainerClose {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("ContainerClosePacket.Container Id", CONTAINERCLOSE_CONTAINER_ID_SHAPE);
        encoder.field("ContainerClosePacket.Container Type", CONTAINERCLOSE_CONTAINER_TYPE_SHAPE);
        encoder.field("ContainerClosePacket.Server Initiated Close", CONTAINERCLOSE_SERVER_INITIATED_CLOSE_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("ContainerClosePacket.Container Id", CONTAINERCLOSE_CONTAINER_ID_SHAPE);
        decoder.field("ContainerClosePacket.Container Type", CONTAINERCLOSE_CONTAINER_TYPE_SHAPE);
        decoder.field("ContainerClosePacket.Server Initiated Close", CONTAINERCLOSE_SERVER_INITIATED_CLOSE_SHAPE);
    }
}
