// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct AddPainting {
    pub target_actor_id: ActorUniqueID,
    pub target_runtime_id: ActorRuntimeID,
    pub position: glam::Vec3,
    pub direction: i32,
    pub motif: String,
}

pub const ADDPAINTING_TARGET_ACTOR_ID_SHAPE: &str = r#"{"kind":"struct","semantic":"ActorUniqueID","type_id":"ActorUniqueID","fields":[{"ordinal":0,"name":"Actor Unique ID","semantic":"Actor Unique ID","encode":{"kind":"primitive","primitive":{"code":"zigzag_i64","width":64,"signed":true,"zigzag":true,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}"#;
pub const ADDPAINTING_TARGET_RUNTIME_ID_SHAPE: &str = r#"{"kind":"struct","semantic":"ActorRuntimeID","type_id":"ActorRuntimeID","fields":[{"ordinal":0,"name":"Actor Runtime ID","semantic":"Actor Runtime ID","encode":{"kind":"primitive","primitive":{"code":"var_u64","width":64,"signed":false,"zigzag":false,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}"#;
pub const ADDPAINTING_POSITION_SHAPE: &str = r#"{"kind":"struct","semantic":"Vec3","type_id":"Vec3","fields":[{"ordinal":0,"name":"X","semantic":"X","encode":{"kind":"primitive","primitive":{"code":"f32le","width":32,"signed":false,"zigzag":false,"endianness":"little"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":1,"name":"Y","semantic":"Y","encode":{"kind":"primitive","primitive":{"code":"f32le","width":32,"signed":false,"zigzag":false,"endianness":"little"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":2,"name":"Z","semantic":"Z","encode":{"kind":"primitive","primitive":{"code":"f32le","width":32,"signed":false,"zigzag":false,"endianness":"little"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}"#;
pub const ADDPAINTING_DIRECTION_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"}}"#;
pub const ADDPAINTING_MOTIF_SHAPE: &str = r#"{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"}"#;

impl AddPainting {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("AddPaintingPacket.Target Actor ID", ADDPAINTING_TARGET_ACTOR_ID_SHAPE);
        encoder.field("AddPaintingPacket.Target Runtime ID", ADDPAINTING_TARGET_RUNTIME_ID_SHAPE);
        encoder.field("AddPaintingPacket.Position", ADDPAINTING_POSITION_SHAPE);
        encoder.field("AddPaintingPacket.Direction", ADDPAINTING_DIRECTION_SHAPE);
        encoder.field("AddPaintingPacket.Motif", ADDPAINTING_MOTIF_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("AddPaintingPacket.Target Actor ID", ADDPAINTING_TARGET_ACTOR_ID_SHAPE);
        decoder.field("AddPaintingPacket.Target Runtime ID", ADDPAINTING_TARGET_RUNTIME_ID_SHAPE);
        decoder.field("AddPaintingPacket.Position", ADDPAINTING_POSITION_SHAPE);
        decoder.field("AddPaintingPacket.Direction", ADDPAINTING_DIRECTION_SHAPE);
        decoder.field("AddPaintingPacket.Motif", ADDPAINTING_MOTIF_SHAPE);
    }
}
