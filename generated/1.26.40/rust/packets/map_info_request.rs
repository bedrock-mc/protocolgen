// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct MapInfoRequest {
    pub map_unique_id: ActorUniqueID,
    pub client_pixels_list: Vec<MapInfoRequestPacketAnonClientPixelsProxy>,
}

pub const MAPINFOREQUEST_MAP_UNIQUE_ID_SHAPE: &str = r#"{"kind":"struct","semantic":"ActorUniqueID","type_id":"ActorUniqueID","fields":[{"ordinal":0,"name":"Actor Unique ID","semantic":"Actor Unique ID","encode":{"kind":"primitive","primitive":{"code":"zigzag_i64","width":64,"signed":true,"zigzag":true,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}"#;
pub const MAPINFOREQUEST_CLIENT_PIXELS_LIST_SHAPE: &str = r#"{"kind":"array","prefix":{"kind":"primitive","primitive":{"code":"u32le","width":32,"signed":false,"zigzag":false,"endianness":"little"}},"element":{"kind":"struct","semantic":"MapInfoRequestPacketAnon::ClientPixelsProxy","type_id":"MapInfoRequestPacketAnon::ClientPixelsProxy","fields":[{"ordinal":0,"name":"pixel","semantic":"pixel","encode":{"kind":"primitive","primitive":{"code":"u32le","width":32,"signed":false,"zigzag":false,"endianness":"little"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":1,"name":"index","semantic":"index","encode":{"kind":"primitive","primitive":{"code":"u16le","width":16,"signed":false,"zigzag":false,"endianness":"little"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}}"#;

impl MapInfoRequest {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("MapInfoRequestPacket.Map Unique ID", MAPINFOREQUEST_MAP_UNIQUE_ID_SHAPE);
        encoder.field("MapInfoRequestPacket.Client Pixels List", MAPINFOREQUEST_CLIENT_PIXELS_LIST_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("MapInfoRequestPacket.Map Unique ID", MAPINFOREQUEST_MAP_UNIQUE_ID_SHAPE);
        decoder.field("MapInfoRequestPacket.Client Pixels List", MAPINFOREQUEST_CLIENT_PIXELS_LIST_SHAPE);
    }
}
