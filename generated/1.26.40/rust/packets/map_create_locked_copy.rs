// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct MapCreateLockedCopy {
    pub original_map_id: ActorUniqueID,
    pub new_map_id: ActorUniqueID,
}

pub const MAPCREATELOCKEDCOPY_ORIGINAL_MAP_ID_SHAPE: &str = r#"{"kind":"struct","semantic":"ActorUniqueID","type_id":"ActorUniqueID","fields":[{"ordinal":0,"name":"Actor Unique ID","semantic":"Actor Unique ID","encode":{"kind":"primitive","primitive":{"code":"zigzag_i64","width":64,"signed":true,"zigzag":true,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}"#;
pub const MAPCREATELOCKEDCOPY_NEW_MAP_ID_SHAPE: &str = r#"{"kind":"struct","semantic":"ActorUniqueID","type_id":"ActorUniqueID","fields":[{"ordinal":0,"name":"Actor Unique ID","semantic":"Actor Unique ID","encode":{"kind":"primitive","primitive":{"code":"zigzag_i64","width":64,"signed":true,"zigzag":true,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}"#;

impl MapCreateLockedCopy {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("MapCreateLockedCopyPacket.Original Map Id", MAPCREATELOCKEDCOPY_ORIGINAL_MAP_ID_SHAPE);
        encoder.field("MapCreateLockedCopyPacket.New Map Id", MAPCREATELOCKEDCOPY_NEW_MAP_ID_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("MapCreateLockedCopyPacket.Original Map Id", MAPCREATELOCKEDCOPY_ORIGINAL_MAP_ID_SHAPE);
        decoder.field("MapCreateLockedCopyPacket.New Map Id", MAPCREATELOCKEDCOPY_NEW_MAP_ID_SHAPE);
    }
}
