// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct SetLocalPlayerAsInitialized {
    pub player_id: ActorRuntimeID,
}

pub const SETLOCALPLAYERASINITIALIZED_PLAYER_ID_SHAPE: &str = r#"{"kind":"struct","semantic":"ActorRuntimeID","type_id":"ActorRuntimeID","fields":[{"ordinal":0,"name":"Actor Runtime ID","semantic":"Actor Runtime ID","encode":{"kind":"primitive","primitive":{"code":"var_u64","width":64,"signed":false,"zigzag":false,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}"#;

impl SetLocalPlayerAsInitialized {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("SetLocalPlayerAsInitializedPacket.Player ID", SETLOCALPLAYERASINITIALIZED_PLAYER_ID_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("SetLocalPlayerAsInitializedPacket.Player ID", SETLOCALPLAYERASINITIALIZED_PLAYER_ID_SHAPE);
    }
}
