// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct EmoteList {
    pub runtime_id: ActorRuntimeID,
    pub emote_piece_ids: Vec<[u8; 16]>,
}

pub const EMOTELIST_RUNTIME_ID_SHAPE: &str = r#"{"kind":"struct","semantic":"ActorRuntimeID","type_id":"ActorRuntimeID","fields":[{"ordinal":0,"name":"Actor Runtime ID","semantic":"Actor Runtime ID","encode":{"kind":"primitive","primitive":{"code":"var_u64","width":64,"signed":false,"zigzag":false,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}"#;
pub const EMOTELIST_EMOTE_PIECE_IDS_SHAPE: &str = r#"{"kind":"array","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"element":{"kind":"primitive","primitive":{"code":"uuid","width":128,"signed":false,"zigzag":false,"endianness":"none"}}}"#;

impl EmoteList {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("EmoteListPacket.Runtime id", EMOTELIST_RUNTIME_ID_SHAPE);
        encoder.field("EmoteListPacket.Emote piece ids", EMOTELIST_EMOTE_PIECE_IDS_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("EmoteListPacket.Runtime id", EMOTELIST_RUNTIME_ID_SHAPE);
        decoder.field("EmoteListPacket.Emote piece ids", EMOTELIST_EMOTE_PIECE_IDS_SHAPE);
    }
}
