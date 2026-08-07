// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct ContainerOpen {
    pub container_id: u8,
    pub container_type: u8,
    pub position: BlockPos,
    pub target_actor_id: ActorUniqueID,
}

pub const CONTAINEROPEN_CONTAINER_ID_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"u8","width":8,"signed":false,"zigzag":false,"endianness":"none"}}"#;
pub const CONTAINEROPEN_CONTAINER_TYPE_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"u8","width":8,"signed":false,"zigzag":false,"endianness":"none"}}"#;
pub const CONTAINEROPEN_POSITION_SHAPE: &str = r#"{"kind":"struct","semantic":"BlockPos","type_id":"BlockPos","fields":[{"ordinal":0,"name":"X","semantic":"X","encode":{"kind":"primitive","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":1,"name":"Y","semantic":"Y","encode":{"kind":"primitive","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":2,"name":"Z","semantic":"Z","encode":{"kind":"primitive","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}"#;
pub const CONTAINEROPEN_TARGET_ACTOR_ID_SHAPE: &str = r#"{"kind":"struct","semantic":"ActorUniqueID","type_id":"ActorUniqueID","fields":[{"ordinal":0,"name":"Actor Unique ID","semantic":"Actor Unique ID","encode":{"kind":"primitive","primitive":{"code":"zigzag_i64","width":64,"signed":true,"zigzag":true,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}"#;

impl ContainerOpen {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("ContainerOpenPacket.Container Id", CONTAINEROPEN_CONTAINER_ID_SHAPE);
        encoder.field("ContainerOpenPacket.Container Type", CONTAINEROPEN_CONTAINER_TYPE_SHAPE);
        encoder.field("ContainerOpenPacket.Position", CONTAINEROPEN_POSITION_SHAPE);
        encoder.field("ContainerOpenPacket.Target Actor ID", CONTAINEROPEN_TARGET_ACTOR_ID_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("ContainerOpenPacket.Container Id", CONTAINEROPEN_CONTAINER_ID_SHAPE);
        decoder.field("ContainerOpenPacket.Container Type", CONTAINEROPEN_CONTAINER_TYPE_SHAPE);
        decoder.field("ContainerOpenPacket.Position", CONTAINEROPEN_POSITION_SHAPE);
        decoder.field("ContainerOpenPacket.Target Actor ID", CONTAINEROPEN_TARGET_ACTOR_ID_SHAPE);
    }
}
