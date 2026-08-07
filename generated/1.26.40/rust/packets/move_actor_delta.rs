// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct MoveActorDelta {
    pub move_data: MoveActorDeltaData,
}

pub const MOVEACTORDELTA_MOVE_DATA_SHAPE: &str = r#"{"kind":"struct","semantic":"MoveActorDeltaData","type_id":"MoveActorDeltaData","fields":[{"ordinal":0,"name":"Actor Runtime ID","semantic":"Actor Runtime ID","encode":{"kind":"struct","semantic":"ActorRuntimeID","type_id":"ActorRuntimeID","fields":[{"ordinal":0,"name":"Actor Runtime ID","semantic":"Actor Runtime ID","encode":{"kind":"primitive","primitive":{"code":"var_u64","width":64,"signed":false,"zigzag":false,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":1,"name":"New Position X","semantic":"New Position X","encode":{"kind":"optional","value":{"kind":"primitive","primitive":{"code":"f32le","width":32,"signed":false,"zigzag":false,"endianness":"little"}}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":2,"name":"New Position Y","semantic":"New Position Y","encode":{"kind":"optional","value":{"kind":"primitive","primitive":{"code":"f32le","width":32,"signed":false,"zigzag":false,"endianness":"little"}}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":3,"name":"New Position Z","semantic":"New Position Z","encode":{"kind":"optional","value":{"kind":"primitive","primitive":{"code":"f32le","width":32,"signed":false,"zigzag":false,"endianness":"little"}}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":4,"name":"Rotation X","semantic":"Rotation X","encode":{"kind":"optional","value":{"kind":"primitive","primitive":{"code":"i8","width":8,"signed":true,"zigzag":false,"endianness":"none"}}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":5,"name":"Rotation Y","semantic":"Rotation Y","encode":{"kind":"optional","value":{"kind":"primitive","primitive":{"code":"i8","width":8,"signed":true,"zigzag":false,"endianness":"none"}}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":6,"name":"Rotation Y Head","semantic":"Rotation Y Head","encode":{"kind":"optional","value":{"kind":"primitive","primitive":{"code":"i8","width":8,"signed":true,"zigzag":false,"endianness":"none"}}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":7,"name":"Is On Ground","semantic":"Is On Ground","encode":{"kind":"primitive","primitive":{"code":"bool","width":1,"signed":false,"zigzag":false,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":8,"name":"Force Move","semantic":"Force Move","encode":{"kind":"primitive","primitive":{"code":"bool","width":1,"signed":false,"zigzag":false,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":9,"name":"Force Move Local Entity","semantic":"Force Move Local Entity","encode":{"kind":"primitive","primitive":{"code":"bool","width":1,"signed":false,"zigzag":false,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":10,"name":"Force Completion","semantic":"Force Completion","encode":{"kind":"primitive","primitive":{"code":"bool","width":1,"signed":false,"zigzag":false,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}"#;

impl MoveActorDelta {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("MoveActorDeltaPacket.Move Data", MOVEACTORDELTA_MOVE_DATA_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("MoveActorDeltaPacket.Move Data", MOVEACTORDELTA_MOVE_DATA_SHAPE);
    }
}
