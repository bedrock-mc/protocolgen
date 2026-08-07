// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct MoveActorAbsolute {
    pub move_data: MoveActorAbsoluteData,
}

pub const MOVEACTORABSOLUTE_MOVE_DATA_SHAPE: &str = r#"{"kind":"struct","semantic":"MoveActorAbsoluteData","type_id":"MoveActorAbsoluteData","fields":[{"ordinal":0,"name":"ActorRuntimeID","semantic":"ActorRuntimeID","encode":{"kind":"struct","semantic":"ActorRuntimeID","type_id":"ActorRuntimeID","fields":[{"ordinal":0,"name":"Actor Runtime ID","semantic":"Actor Runtime ID","encode":{"kind":"primitive","primitive":{"code":"var_u64","width":64,"signed":false,"zigzag":false,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":1,"name":"Header","semantic":"Header","encode":{"kind":"primitive","primitive":{"code":"u8","width":8,"signed":false,"zigzag":false,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":2,"name":"Position","semantic":"Position","encode":{"kind":"struct","semantic":"Vec3","type_id":"Vec3","fields":[{"ordinal":0,"name":"X","semantic":"X","encode":{"kind":"primitive","primitive":{"code":"f32le","width":32,"signed":false,"zigzag":false,"endianness":"little"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":1,"name":"Y","semantic":"Y","encode":{"kind":"primitive","primitive":{"code":"f32le","width":32,"signed":false,"zigzag":false,"endianness":"little"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":2,"name":"Z","semantic":"Z","encode":{"kind":"primitive","primitive":{"code":"f32le","width":32,"signed":false,"zigzag":false,"endianness":"little"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":3,"name":"Rotation X","semantic":"Rotation X","encode":{"kind":"primitive","primitive":{"code":"u8","width":8,"signed":false,"zigzag":false,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":4,"name":"Rotation Y","semantic":"Rotation Y","encode":{"kind":"primitive","primitive":{"code":"u8","width":8,"signed":false,"zigzag":false,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":5,"name":"Rotation Y Head","semantic":"Rotation Y Head","encode":{"kind":"primitive","primitive":{"code":"u8","width":8,"signed":false,"zigzag":false,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}"#;

impl MoveActorAbsolute {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("MoveActorAbsolutePacket.Move Data", MOVEACTORABSOLUTE_MOVE_DATA_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("MoveActorAbsolutePacket.Move Data", MOVEACTORABSOLUTE_MOVE_DATA_SHAPE);
    }
}
